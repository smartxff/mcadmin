package cmd

import (
	"github.com/minio/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"log"
)
var mcCmd = cli.Command{
	Name: "mc",
	Usage: "初始化mc配置，用于后续数据传输，及管理",
	Action: mainMc,
}


func mainMc(ctx *cli.Context){
	type Host struct {
		Alias     string
		Server    string
		Accesskey string
		Secretkey string
	}
	host := new(Host)
	var qs = []*survey.Question{
		{
			Name: "server",
			Prompt: &survey.Input{
				Message: "输入存储服务器地址,例如：http://172.16.130.11:9000:",
			},
			Validate:survey.Required,
		},
		{
			Name: "accesskey",
			Prompt: &survey.Input{
				Message: "输入存储服务器的ACCESSKEY，启动minio的时候会输出:",
			},
			Validate:survey.Required,
		},
		{
			Name: "secretkey",
			Prompt: &survey.Input{
				Message: "输入存储服务器的SECRETKEY，启动minio的时候会输出:",
			},
		},
	}
	err := survey.Ask(qs,host)
	if err !=nil{
		log.Fatalf("内部错误:%s\n",err.Error())
		return
	}


	addAlias := []string{"mc","config","host","add","minio",host.Server,host.Accesskey,host.Secretkey}

	Command(addAlias).run()


}
