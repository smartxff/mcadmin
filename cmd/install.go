package cmd

import (
	"github.com/smartxff/wget-go/wget"
	"log"
	"os"
	"path"
)

var (
	BASEURL = "https://dl.minio.io/client/mc/release/%s-%s/mc"
	WINDOWS = "windows"
	LINUX = "linux"
	PPC = "ppc64le"
	AMD64 = "amd64"
)



func installmc(){
	goos := CheckOS()


	//初始化path变量
	goos.ExecInitScript()

	//下载mc命令
	url := goos.GetMcUrl()
	log.Printf("mc 下载地址：%s\n",url)
	w := wget.Wget(url)

	w.OutputFilename = path.Base(url)
	err,n := w.Exec(os.Stdin,os.Stdout,os.Stderr)
	if err !=nil{
		log.Fatalf("下载文件出错：%s,%d",err.Error(),n)
	}
}


