package main

import (
	"fmt"
	"log"

	"github.com/holmen1/holmcoin/block"
	"github.com/holmen1/holmcoin/wallet"
)

func init() {
	log.SetPrefix("Holmchain: ")
}

func main() {
	walletMiner := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// Sign
	signature, _ := t.GenerateSignature()

	// Blockchain
	blockchain := block.NewBlockchain(walletMiner.BlockchainAddress())
	blockchain.AddTransaction(
		walletA.BlockchainAddress(),
		walletB.BlockchainAddress(),
		1.0,
		walletA.PublicKey(),
		signature)

	blockchain.Mining()

	blockchain.Print()

	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %.1f\n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("Miner %.1f\n", blockchain.CalculateTotalAmount(walletMiner.BlockchainAddress()))

}
