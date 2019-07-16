// +build !windows

package main

import (
	"log"
	"syscall"
)

func init() {
	setUlimit()
}

func setUlimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatal("getting Rlimit failed", err)
	} else {
		rLimit.Max = 1024
		rLimit.Cur = 1024
		if err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
			log.Fatal("setting Rlimit failed", err)
		}
	}
}
