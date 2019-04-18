package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func McAddPolicy(username,buckets,rw,commandName string) {
	passwd := GetRandomString(16)
	var cmd []string
	addUser := []string{"mc","admin","user","add","minio",username,passwd}

	updateUser := []string{"mc","admin","user", "policy","minio",username}

	addpolicy := []string{"mc","admin","policy","add","minio"}

	switch commandName {
	case "add":
		cmd = addUser
	case "update":
		cmd = updateUser

	}
	if commandName == "add"{
		cmd = addUser
	}

	var policyname string

	if buckets == "*"{
		if isReadOnly(rw){
			policyname = "readonly"
		}else{
			policyname = "readwrite"
		}

		cmd = append(cmd,policyname)

		Command(cmd).run()
		return
	}


	policy := NewPolicy()
	policy.Statement[0].AddS3Buckets(buckets)
	if isReadOnly(rw){
		policy.Statement[0].AddActionRead()
	}else{
		policy.Statement[0].AddActionRewdWrite()
	}
	policy.OutPutToFile(buckets)

	addpolicy =append(addpolicy,buckets,buckets)

	Command(addpolicy).run()


	cmd = append(cmd,buckets)

	Command(cmd).run()
	if commandName == "add"{
		fmt.Printf("用户%s的密码为：%s\n",username,passwd)
	}

}

func isReadOnly(rw string)bool{
	if rw == "r"{
		return true
	}
	return false
}


type Command []string

func (c Command)run(){
	log.Println("执行命令：",strings.Join(c," "))
	args := c[1:]
	cmd := exec.Command(c[0],args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

}