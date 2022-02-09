package main

import (
	"github.com/DaegunHan/hancoin/blockchain"
	"github.com/DaegunHan/hancoin/cli"
	"github.com/DaegunHan/hancoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
