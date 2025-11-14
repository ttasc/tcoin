package cli

import (
	"fmt"
	"log"

	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func (cli *CLI) startNode(nodeIP, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeIP)
	if len(minerAddress) > 0 {
		if blockchain.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	blockchain.StartServer(nodeIP, minerAddress)
}
