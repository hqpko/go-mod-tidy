package main

import (
	"testing"
)

func TestMod(t *testing.T) {
	modStr := `module go-mod-tidy

require (
	github.com/andygrunwald/go-jira v1.5.0 // indirect
	github.com/beevik/guid v0.0.0-20170504223318-d0ea8faecee0 // indirect
	github.com/cloudfoundry/gosigar v1.1.0 // indirect
	github.com/fatih/structs v1.0.0 // indirect
	cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0
)

replace (
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	google.golang.org/grpc v1.14.0 => github.com/grpc/grpc-go v1.14.0
	golang.org/x/sys v0.0.0-20180909124046-d0be0721c37e => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
)`

	m := newMod()
	if e := m.parse(modStr); e != nil {
		t.Error(e)
	}
}
