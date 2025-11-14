package cli

import (
	"fmt"

	"github.com/atsuyaourt/blockchain/internal/blockchain"
)

func (cli *CLI) createWallet(nodeIP string) {
	wallets, _ := blockchain.NewWallets(nodeIP)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeIP)

	fmt.Printf("Your new address: %s\n", address)
}
