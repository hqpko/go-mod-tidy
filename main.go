package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hqpko/gosh"
)

func main() {
	if os.Getenv("GO111MODULE") != "on" {
		fmt.Println("GO111MODULE != on,please\nexport GO111MODULE=on")
		return
	}
	m, err := readMod()
	if err != nil {
		fmt.Println("read mod error:", err.Error())
		return
	}
	m.save()

	for {
		session := gosh.NewSession()
		session.SetHandlerErr(handlerFormatPackagesReplace(m))
		if err := session.Run("go mod tidy"); err == nil {
			fmt.Println("done")
			break
		}
		fmt.Println("\nre tidy ...")
	}
}

func handlerFormatPackagesReplace(m *mod) func(s *gosh.Session, err *exec.ExitError) {
	return func(s *gosh.Session, err *exec.ExitError) {
		failMsg := string(err.Stderr)
		addCount := 0
		for _, line := range strings.Split(failMsg, "\n") {
			if p := getReplacePackage(line); p != nil {
				fmt.Println("add replace:", p.name, p.version)
				m.addReplace(p)
				addCount++
			}
		}
		if addCount == 0 {
			fmt.Println(failMsg)
			os.Exit(-1)
		}
		m.save()
	}
}

// go: golang.org/x/net@v0.0.0-20180906233101-161cd47e91fd: unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
// "golang.org/x/net" start point is 4
const startPoint = 4

func getReplacePackage(s string) *pack {
	if !strings.Contains(s, "unrecognized import path") {
		return nil
	}
	endPoint := strings.Index(s, ": unrecognized")
	if endPoint <= startPoint {
		return nil
	}
	packageInfo := strings.Split(s[startPoint:endPoint], "@")
	if len(packageInfo) != 2 {
		return nil
	}
	if _, ok := replaceMap[packageInfo[0]]; ok {
		return newPack(packageInfo[0], packageInfo[1])
	}

	return nil
}
