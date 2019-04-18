package cmd

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

var baseScript = `
echo "PATH=$PWD:$PATH" >> ~/.bashrc
`

var batScript = `
set PATH = %cd%;%PATH%
pause
`

type windows int
type linux int
type goos interface {
	GetMcUrl()string
	ExecInitScript()
}

func CheckOS()goos{
	gos := runtime.GOOS
	switch gos {
	case WINDOWS:
		return new(windows)
	case LINUX:
		 return new(linux)
	}
	log.Fatalf("不支持此操作系统: %s",gos)
	return nil
}

func (s *windows)GetMcUrl()string{
	return fmt.Sprintf(BASEURL,WINDOWS,AMD64)+".exe"
}

func (s *linux)GetMcUrl()string{
	if s.isPPC(){
		return fmt.Sprintf(BASEURL,LINUX,PPC)
	}
	return fmt.Sprintf(BASEURL,LINUX,AMD64)
}

func (s *linux)ExecInitScript(){
	//todo
}


func (s *linux)isPPC()bool{
	carch := runtime.GOARCH
	if strings.Contains(carch,"ppc"){
		return true
	}
	return false
}

func (s *windows)ExecInitScript(){
	//todo
}