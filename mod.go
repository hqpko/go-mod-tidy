package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	replaceMap = map[string]string{
		"golang.org/x/tools":          "github.com/golang/tools",
		"golang.org/x/sys":            "github.com/golang/sys",
		"golang.org/x/sync":           "github.com/golang/sync",
		"golang.org/x/oauth2":         "github.com/golang/oauth2",
		"golang.org/x/net":            "github.com/golang/net",
		"golang.org/x/lint":           "github.com/golang/lint",
		"golang.org/x/text":           "github.com/golang/text",
		"google.golang.org/genproto":  "github.com/google/go-genproto",
		"google.golang.org/grpc":      "github.com/grpc/grpc-go",
		"google.golang.org/appengine": "github.com/golang/appengine",
		"cloud.google.com/go":         "github.com/googleapis/google-cloud-go",
		"google.golang.org/api":       "github.com/googleapis/google-api-go-client",
	}
)

type pack struct {
	name    string
	version string
}

func newPack(name, version string) *pack {
	return &pack{name: name, version: version}
}

type mod struct {
	title   string
	require []*pack
	replace []*pack
}

func newMod() *mod {
	return &mod{require: []*pack{}, replace: []*pack{}}
}

func readMod() (*mod, error) {
	data, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return nil, err
	}

	m := newMod()
	return m, m.parse(string(data))
}

func (m *mod) parse(modStr string) error {
	lines := strings.Split(modStr, "\n")
	if m.title = m.readTitle(lines); m.title == "" {
		return errors.New("no module name")
	}
	m.require = m.readPacks(lines, "require")
	m.replace = m.readPacks(lines, "replace")

	m.moveReplace()

	return nil
}

func (m *mod) addReplace(p *pack) {
	m.replace = append(m.replace, p)
}

func (m *mod) save() {
	s := m.title
	if len(m.require) > 0 {
		s += "\n\nrequire (\n"
		for _, p := range m.require {
			s += fmt.Sprintf("	%s %s\n", p.name, p.version)
		}
		s += ")"
	}
	if len(m.replace) > 0 {
		s += "\n\nreplace (\n"
		for _, p := range m.replace {
			s += fmt.Sprintf("	%s %s => %s %s\n", p.name, p.version, replaceMap[p.name], p.version)
		}
		s += ")"
	}
	_ = ioutil.WriteFile("go.mod", []byte(s), os.ModePerm)
}

func (m *mod) readTitle(lines []string) string {
	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			return line
		}
	}
	return ""
}

func (m *mod) readPacks(lines []string, packType string) []*pack {
	ps := make([]*pack, 0)
	startType := false
	for _, line := range lines {
		if startType {
			if strings.HasPrefix(line, ")") {
				break
			}
			if p := m.readPack(line); p != nil {
				ps = append(ps, p)
			}
		}
		if !startType && strings.HasPrefix(line, packType) {
			startType = true
		}
	}
	return ps
}

func (m *mod) readPack(line string) *pack {
	packInfo := strings.Split(strings.TrimSpace(line), " ")
	if len(packInfo) >= 2 {
		return newPack(packInfo[0], packInfo[1])
	}
	return nil
}

func (m *mod) moveReplace() {
	require := make([]*pack, 0)
	for _, p := range m.require {
		if _, ok := replaceMap[p.name]; ok {
			m.replace = append(m.replace, p)
		} else {
			require = append(require, p)
		}
	}
}
