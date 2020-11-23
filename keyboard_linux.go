// +build linux
package main

import (
	"log"
	"strings"

	"github.com/MarinX/keylogger"
)

func waitForE(k *keylogger.KeyLogger) {
	events := k.Read()
	for {
		for e := range events {
			switch e.Type {
			// EvKey is used to describe state changes of keyboards, buttons, or other key-like devices.
			// check the input_event.go for more events
			case keylogger.EvKey:

				// if the state of key is pressed
				if e.KeyPress() {
					if strings.ToLower(e.KeyString()) == "e" {
						return
					}
				}
				break
			}
		}
	}
}

func keylogInit() *keylogger.KeyLogger {
	keyboard := keylogger.FindKeyboardDevice()

	log.Println("Found a keyboard at", keyboard)
	// init keylogger with keyboard
	k, err := keylogger.New(keyboard)
	if err != nil {
		log.Fatal(err)
	}

	return k
}
