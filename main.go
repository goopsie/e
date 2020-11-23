package main

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playSound(streamer beep.StreamSeeker) {
	log.Print("playing e.mp3")
	err := streamer.Seek(0)
	if err != nil {
		log.Print("seek: ", err)
		return
	}

	speaker.Clear()
	speaker.Play(streamer)
	if streamer.Err() != nil {
		log.Print("streamer: ", streamer.Err())
	}
}

// why am i forced to commit such sin by beep
// it wants close() but bytes reader does not provide
// i make fake :)))
type bufCloser struct {
	*bytes.Reader
}

func (b *bufCloser) Close() error { return nil }

func init() {
	log.SetFlags(log.Lshortfile)
}

func persist() {
	for {
		err := attemptPersist()
		if err != nil {
			log.Print(err)
			time.Sleep(time.Minute)
		} else {
			break
		}
	}
}

func main() {
	// Attempt persistance in the background; since we retry
	// on failures.
	go persist()

	// Initialize speaker and sound code.
	soundBytesB64Buf := bytes.NewBuffer(soundBytesB64)
	r := base64.NewDecoder(base64.StdEncoding, soundBytesB64Buf)
	soundBytes, err := ioutil.ReadAll(r)
	errp(err)
	soundBytesBuf := &bufCloser{bytes.NewReader(soundBytes)}

	streamer, format, err := mp3.Decode(soundBytesBuf)
	errp(err)
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	defer speaker.Close()

	key := keylogInit()
	for {
		waitForE(key)
		playSound(streamer)
	}
}

func errp(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
