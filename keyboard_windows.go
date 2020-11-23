// +build windows
package main

import (
	"time"

	"github.com/kindlyfire/go-keylogger"
)

var lastkeypress time.Time
var delay = time.Millisecond * 100

func waitForE(kl keylogger.Keylogger) {
	for {
		if time.Now().Sub(lastkeypress) < delay { // if E held down only play once (at the start)
			continue
		}
		key := kl.GetKey()
		if !key.Empty {
			if key.Rune == 'e' || key.Rune == 'E' {
				lastkeypress = time.Now()
				return
			}
		}
	}
}

func keylogInit() keylogger.Keylogger {
	kl := keylogger.NewKeylogger()
	return kl
}
