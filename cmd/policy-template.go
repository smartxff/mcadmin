package cmd

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

var (
	S3 = "arn:aws:s3:::"
	S3Readonly = []string{"s3:ListBucket","s3:GetObject"}
	S3ReadWrite = []string{"s3:*"}
	Pversion = "2012-10-17"
)


func NewPolicy()*Policy{
	policy := new(Policy)
	policy.Version = Pversion


	statements := make([]*State,1)
	statements[0] = NewState()
	statements[0].Effect = "Allow"

	policy.Statement = statements
	return policy
}

type Policy struct {
	Version      string           `json:"Version"`
	Statement    []*State           `json:"Statement"`
}

func (p *Policy)AddStatement(statement *State){
	p.Statement = append(p.Statement,statement)
}

func NewState()*State{
	return new(State)
}

type State struct {
	Action   []string             `json:"Action"`
	Effect    string              `json:"Effect"`
	Resource  []string            `json:"Resource"`
	Sid        string             `json:"Sid"`
}

func (s *State)AddResource(resourceName string){
	s.Resource = append(s.Resource,resourceName)
}

func (s *State)AddS3Bucket(bucketName string){
	resourceName := S3 + bucketName
	s.AddResource(resourceName)
}
func (s *State)AddS3Buckets(buckets string){
	bucketNames := make([]string,0)
	if buckets == "*"{
		bucketNames = append(bucketNames,"*")
	}else {
		bucketNames = strings.Split(buckets,",")
	}
	for _,v := range bucketNames{
		resourceName := S3 + v
		s.AddResource(resourceName)
	}
}

func (s *State)AddActionRead(){
	s.Action = append(s.Action,S3Readonly...)
}

func (s *State)AddActionRewdWrite(){
	s.Action = append(s.Action,S3ReadWrite...)
}

func (p *Policy)OutPutToFile(filename string){
	sp,err := json.Marshal(p)
	if err !=nil{
		log.Fatalf("policy to json error:%s\n",err.Error())
		return
	}

	file,err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY,0644)
	if err !=nil{
		log.Fatalf("output to file error:%s\n",err.Error())
		return
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	_,err = buf.Write(sp)
	if err !=nil{
		log.Fatalf("write file error: %s\n",err.Error())
		return
	}
	buf.Flush()
	return
}