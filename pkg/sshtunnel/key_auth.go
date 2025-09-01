package sshtunnel

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"
)

// loadSSHKeysFromDir loads private keys from a specified directory
func loadSSHKeysFromDir(sshDir string, passphrase string) ([]ssh.Signer, error) {
	var signers []ssh.Signer

	// Also try to load any file that looks like a private key
	if files, err := os.ReadDir(sshDir); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}

			name := file.Name()
			// Look for files that might be private keys
			if strings.Contains(name, "id_") ||
				strings.Contains(name, "key") ||
				strings.HasSuffix(name, ".pem") {

				keyPath := filepath.Join(sshDir, name)
				if signer, err := loadPrivateKey(keyPath, passphrase); err == nil {
					signers = append(signers, signer)
					log.Printf("Loaded key: %s", name)
				}
			}
		}
	}

	return signers, nil
}

func loadPrivateKey(keyPath, passphrase string) (ssh.Signer, error) {
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return nil, err
	}

	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file %s: %w", keyPath, err)
	}

	// Try without passphrase first
	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		// If it's encrypted, try with passphrase
		if _, ok := err.(*ssh.PassphraseMissingError); ok {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(keyBytes, []byte(passphrase))
			if err != nil {
				return nil, fmt.Errorf("failed to parse encrypted private key %s: %w", keyPath, err)
			}
			return signer, nil
		}
		return nil, fmt.Errorf("failed to parse private key %s: %w", keyPath, err)
	}

	return signer, nil
}

// LoadSSHKeys creates an ssh.AuthMethod from loaded keys
func LoadSSHKeysFromDir(sshDir string, passphrase string) (ssh.AuthMethod, error) {
	if sshDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get home directory: %w", err)
		}
		sshDir = filepath.Join(homeDir, ".ssh")
	}
	signers, err := loadSSHKeysFromDir(sshDir, passphrase)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signers...), nil
}

// Helper function to load a specific key file
func LoadKey(keyPath string) (ssh.AuthMethod, error) {
	signer, err := loadPrivateKey(keyPath, "")
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signer), nil
}

// Helper function to load multiple specific key files
func LoadKeys(keyPaths []string) (ssh.AuthMethod, error) {
	var signers []ssh.Signer

	for _, keyPath := range keyPaths {
		signer, err := loadPrivateKey(keyPath, "")
		if err != nil {
			log.Printf("Warning: failed to load key %s: %v", keyPath, err)
			continue
		}
		signers = append(signers, signer)
	}

	if len(signers) == 0 {
		return nil, fmt.Errorf("no valid keys loaded from provided paths")
	}

	return ssh.PublicKeys(signers...), nil
}
