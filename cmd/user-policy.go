package cmd

import (
	"github.com/minio/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"log"
)

var userCmd = cli.Command{
	Name: "user",
	Usage: "用于用户管理，授权bucket",
	Action: mainUser,
	Subcommands: []cli.Command{
		addUserCmd,
		updateUserCmd,
	},
}

var addUserCmd = cli.Command{
	Name: "add",
	Usage: "新增用户并绑定bucket策略",
	Action: updateUser,
}

var updateUserCmd = cli.Command{
	Name: "update",
	Usage: "更新用户的bucket策略",
	Action: updateUser,
}



func mainUser(cx *cli.Context){
	cli.ShowCommandHelp(cx,cx.Args().First())
}


func updateUser(cx *cli.Context){
	type UserPolicy struct {
		UserName    string           `json:"username"`
		BucketName  string           `json:"bucketname"`
		RW          string           `json:"rw"`
	}

	userPolicy := new(UserPolicy)

	var qs = []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message:"请输入用户名字:\n",
			},
			Validate:survey.Required,
		},
		{
			Name: "bucketname",
			Prompt: &survey.Input{
				Message: "请输入bucketname，比如：mysql，或者多个bucket用逗号隔开：mysql,pgsql。默认为:*,表示所有bucket\n",
				Default: "*",
			},
			Validate:survey.Required,
		},
		{
			Name: "rw",
			Prompt: &survey.Input{
				Message: "请输入用户权限，1：r,2:rw，默认为rw :\n",
				Default: "rw",
			},
			Validate:survey.Required,
		},
	}
	err := survey.Ask(qs,userPolicy)
	if err !=nil{
		log.Fatalf("内部错误:%s\n",err.Error())
		return
	}
	McAddPolicy(userPolicy.UserName,userPolicy.BucketName,userPolicy.RW,cx.Command.Name)
}



