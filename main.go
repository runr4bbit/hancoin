package main

import (
	"github.com/DaegunHan/hancoin/cli"
	"github.com/DaegunHan/hancoin/db"
)

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
