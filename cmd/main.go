package cmd

import (
	"github.com/minio/cli"
	"os"
)

var (
	mcadminFlags = []cli.Flag{}
)

var mcadminHelpTemplate = `NAME:
  {{.Name}} - {{.Usage}}

USAGE:
  {{.Name}}  COMMAND

COMMANDS:
  {{range .VisibleCommands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
  {{end}}{{if .VisibleFlags}}

EXAMPLE:
  1.安装mc客户端
	mcadmin -install
  
  2.配置mc客户端
	mcadmin mc

  3.新增用户，并授予指定权限
	mcadmin user add
  
  4.更新用户权限
	mcadmin user update
{{end}}`

var appCmds = []cli.Command{
mcCmd,
userCmd,
}


func Main(){
	if len(os.Args) >1{
		switch os.Args[1] {
		case "-install":
			installmc()
			return

		}
	}

	app := registerApp()
	app.ExtraInfo = func() map[string]string {
		return make(map[string]string)
	}
	app.RunAndExitOnError()
}


func registerApp() *cli.App{
	for _,cmd := range appCmds{

		registerCmd(cmd)
	}

	app := cli.NewApp()
	app.Usage = "minio授权管理工具"
	app.Commands = commands
	app.Author = "smartxff"
	app.Version = Version
	app.HideHelp = true
	app.HideHelpCommand = true
	app.Flags = mcadminFlags
	app.CustomAppHelpTemplate = mcadminHelpTemplate
	//app.CommandNotFound = commandNotFount
	app.EnableBashCompletion = true

	return app
}

