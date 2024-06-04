package lib

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net"

	"github.com/google/uuid"
)

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Error generating random token:", err)
	}

	mac, err := getMacAddr()
	if err != nil {
		log.Fatal("Error getting MAC address:", err)
	}

	hash := sha256.New()
	hash.Write([]byte(base64.StdEncoding.EncodeToString(b) + mac[0]))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
