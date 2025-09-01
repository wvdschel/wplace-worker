package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/jrsap/wplace-worker/pkg/sshtunnel"
	"golang.org/x/crypto/ssh"
)

var (
	sshAddr   = flag.String("ssh-addr", "nl1.wim.land:22", "SSH server address")
	username  = flag.String("username", os.Getenv("USER"), "SSH username")
	localPort = flag.Int("local-port", 9999, "Local port to bind SOCKS proxy to")
)

func main() {
	flag.Parse()

	keyAuth, err := sshtunnel.LoadSSHKeysFromDir("", "")
	if err != nil {
		log.Fatal(err)
	}

	// Configure authentication
	auth := []ssh.AuthMethod{
		keyAuth,
	}

	proxy, err := sshtunnel.NewSOCKSProxy(*sshAddr, *username, auth)
	if err != nil {
		log.Fatal("Failed to create SOCKS proxy: ", err)
	}
	defer proxy.Close()

	localAddr := "localhost:" + strconv.Itoa(*localPort)
	if err := proxy.Start(localAddr); err != nil {
		log.Fatal("Failed to start proxy: ", err)
	}

	// Keep the program running
	select {}
}
