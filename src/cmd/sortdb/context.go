package main

import (
	"log"
	"net"
	"os"

	"lib/sorteddb"
	"lib/util"
)

type Context struct {
	db           *sorteddb.DB
	httpAddr     *net.TCPAddr
	httpListener net.Listener
	reloadChan   chan int
	waitGroup    util.WaitGroupWrapper
}

func verifyAddress(arg string, address string) *net.TCPAddr {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalf("FATAL: failed to resolve %s address (%s) - %s", arg, address, err)
		os.Exit(1)
	}

	return addr
}

func (c *Context) ReloadLoop() {
	for {
		<-c.reloadChan
		err := c.db.Remap()
		if err != nil {
			log.Fatalf("ERROR remapping DB %q", err)
		}
	}

}
