// +build windows
package main

import (
	"time"

	"github.com/kindlyfire/go-keylogger"
)

var lastkeypress int64
var delay int64 = 200

func waitForE(kl keylogger.Keylogger) {
	for {
		if lastkeypress+delay > time.Now().UnixNano()/1000000 { // this library sucks and so does this
			continue
		}
		key := kl.GetKey()
		if !key.Empty {
			if key.Rune == 'e' || key.Rune == 'E' {
				lastkeypress = time.Now().UnixNano() / 1000000 // see above comment
				return
			}
		}
	}
}

func keylogInit() keylogger.Keylogger {
	kl := keylogger.NewKeylogger()
	return kl
}
