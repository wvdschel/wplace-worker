package sshtunnel

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

type SOCKSProxy struct {
	sshClient *ssh.Client
	listener  net.Listener
}

func NewSOCKSProxy(sshAddr, username string, auth []ssh.AuthMethod) (*SOCKSProxy, error) {
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Use proper host key verification in production
		Timeout:         30 * time.Second,
	}

	client, err := ssh.Dial("tcp", sshAddr, config)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SSH server: %w", err)
	}

	return &SOCKSProxy{
		sshClient: client,
	}, nil
}

func (s *SOCKSProxy) Start(localAddr string) error {
	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return fmt.Errorf("failed to start SOCKS proxy: %w", err)
	}
	s.listener = listener

	log.Printf("SOCKS proxy listening on %s", localAddr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				select {
				case <-time.After(1 * time.Millisecond):
					// Retry briefly in case of temporary errors
					continue
				default:
					log.Printf("Stopped accepting connections: %v", err)
					return
				}
			}

			go s.handleConnection(conn)
		}
	}()

	return nil
}

func (s *SOCKSProxy) handleConnection(conn net.Conn) {
	defer conn.Close()

	if err := s.handleSOCKSHandshake(conn); err != nil {
		log.Printf("SOCKS handshake failed: %v", err)
		return
	}

	targetAddr, err := s.readSOCKSRequest(conn)
	if err != nil {
		log.Printf("Failed to read SOCKS request: %v", err)
		return
	}

	remoteConn, err := s.sshClient.Dial("tcp", targetAddr)
	if err != nil {
		log.Printf("Failed to connect to %s: %v", targetAddr, err)
		s.sendSOCKSError(conn, 5) // Connection refused
		return
	}
	defer remoteConn.Close()

	// Send success response
	s.sendSOCKSSuccess(conn)

	// Bidirectional copy
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		io.Copy(conn, remoteConn)
	}()

	go func() {
		defer wg.Done()
		io.Copy(remoteConn, conn)
	}()

	wg.Wait()
}

func (s *SOCKSProxy) handleSOCKSHandshake(conn net.Conn) error {
	// Read version
	version, err := readByte(conn)
	if err != nil {
		return err
	}

	if version != 5 {
		return fmt.Errorf("unsupported SOCKS version: %d", version)
	}

	// Read number of methods
	nmethods, err := readByte(conn)
	if err != nil {
		return err
	}

	// Read methods
	methods := make([]byte, nmethods)
	if _, err := io.ReadFull(conn, methods); err != nil {
		return err
	}

	// Reply with no authentication (0x00)
	_, err = conn.Write([]byte{0x05, 0x00})
	return err
}

func (s *SOCKSProxy) readSOCKSRequest(conn net.Conn) (string, error) {
	// Read version
	version, err := readByte(conn)
	if err != nil {
		return "", err
	}

	if version != 5 {
		return "", fmt.Errorf("unsupported SOCKS version in request: %d", version)
	}

	// Read command
	cmd, err := readByte(conn)
	if err != nil {
		return "", err
	}

	if cmd != 1 { // CONNECT
		return "", fmt.Errorf("unsupported command: %d", cmd)
	}

	// Skip RSV
	readByte(conn)

	// Read address type
	addrType, err := readByte(conn)
	if err != nil {
		return "", err
	}

	switch addrType {
	case 1: // IPv4
		return s.readIPv4Address(conn)
	case 3: // Domain name
		return s.readDomainAddress(conn)
	case 4: // IPv6
		return s.readIPv6Address(conn)
	default:
		return "", fmt.Errorf("unsupported address type: %d", addrType)
	}
}

func (s *SOCKSProxy) readIPv4Address(conn net.Conn) (string, error) {
	ip := make([]byte, 4)
	if _, err := io.ReadFull(conn, ip); err != nil {
		return "", err
	}
	port, err := readPort(conn)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", ip[0], ip[1], ip[2], ip[3], port), nil
}

func (s *SOCKSProxy) readDomainAddress(conn net.Conn) (string, error) {
	domainLen, err := readByte(conn)
	if err != nil {
		return "", err
	}
	domain := make([]byte, domainLen)
	if _, err := io.ReadFull(conn, domain); err != nil {
		return "", err
	}
	port, err := readPort(conn)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", string(domain), port), nil
}

func (s *SOCKSProxy) readIPv6Address(conn net.Conn) (string, error) {
	ip := make([]byte, 16)
	if _, err := io.ReadFull(conn, ip); err != nil {
		return "", err
	}
	port, err := readPort(conn)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("[%s]:%d", net.IP(ip).String(), port), nil
}

func (s *SOCKSProxy) sendSOCKSError(conn net.Conn, code byte) {
	conn.Write([]byte{0x05, code, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
}

func (s *SOCKSProxy) sendSOCKSSuccess(conn net.Conn) {
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
}

func (s *SOCKSProxy) Close() error {
	if s.listener != nil {
		s.listener.Close()
	}
	return s.sshClient.Close()
}

func readByte(conn net.Conn) (byte, error) {
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	return buf[0], err
}

func readPort(conn net.Conn) (int, error) {
	portBytes := make([]byte, 2)
	if _, err := io.ReadFull(conn, portBytes); err != nil {
		return 0, err
	}
	return int(portBytes[0])<<8 | int(portBytes[1]), nil
}
