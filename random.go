package id

import "sync"

var (
	mutex sync.RWMutex
	ok    string = "NON"
	//ID is the ALgo
	ID = "ECDSA-256"
)
