package cli

import (
	"fmt"
	"log"

	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func (cli *CLI) listAddresses(nodeIP string) {
	wallets, err := blockchain.NewWallets(nodeIP)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
