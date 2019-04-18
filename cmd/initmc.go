package cmd

import (
	"fmt"
	"runtime"
	"strings"
	"log"
)

var bashScript = `
echo "export PATH=$PWD:\$PATH" >> ~/.bashrc
`

var batScript = `
setx "Path"  "%cd%;%Path%" /m
set path_=%cd%;%Path%
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
	cmd := Command{"sh","-c",bashScript}
	cmd.run()
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
	cmd := Command{"cmd","/c",batScript}
	cmd.run()
}