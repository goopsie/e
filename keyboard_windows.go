// +build windows
package main

import (
	"github.com/kindlyfire/go-keylogger"
)

func waitForE(kl keylogger.Keylogger) {
	for {
		key := kl.GetKey()
		if !key.Empty {
			if key.Rune == 'e' || key.Rune == 'E' {
				return
			}
		}
	}
}

func keylogInit() keylogger.Keylogger {
	kl := keylogger.NewKeylogger()
	return kl
}
