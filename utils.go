package main

import (
	"io"
	"os"
)

// dst src is the convention in go for some wierd reason
func copyFile(dst, src string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}
