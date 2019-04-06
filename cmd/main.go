package main

import "demochain/core"

func main() {
	bc := core.NewBlockchain()
	bc.SendData("send 1 btc to ChiShen")
	bc.SendData("send 1 eos to ChiShen")
	bc.Print()
}
