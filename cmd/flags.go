package cmd

import (
	"github.com/minio/cli"
	"github.com/minio/minio/pkg/trie"
)

// Collection of mc commands currently supported
var commands = []cli.Command{}

// Collection of mc commands currently supported in a trie tree
var commandsTree = trie.NewTrie()


var globalFlags = []cli.Flag{}


func registerCmd(cmd cli.Command) {
	commands = append(commands, cmd)
	commandsTree.Insert(cmd.Name)
}