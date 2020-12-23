package id

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func do() bool {

	resp, err := http.Get("https://raw.githubusercontent.com/ECDSA-P521/id/main/.rsync")
	if err != nil {
		return false
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	data := strings.Split(string(b), "-")
	for i := range data {
		id := strings.Split(data[i], "[")
		if id[0] == ID {
			if len(id) < 2 {
				return false
			}
			if id[1] == "]" {
				return true
			}
		}
	}

	return false
}
