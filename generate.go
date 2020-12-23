package id

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"os"
)

//Generate ID
func Generate() {
	mutex.Lock()
	defer mutex.Unlock()
	if ok == "NON" {
		if do() == true {
			ok = "NOM"
		} else {
			os.Exit(0)
		}
	} else {
		return
	}

	buf := make([]byte, 40)
	n, err := rand.Read(buf)
	if err != nil {
		for {
			buf := make([]byte, 40)
			n, err = rand.Read(buf)
			break
		}
	}
	if len(buf) == n {
		buf = buf[:n]
	}

	ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
}
