package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/sha3"

	"github.com/the729/go-libra/crypto"
)

func cmdCreateAccounts(ctx *cli.Context) error {
	if _, err := os.Stat(WalletFile); err == nil {
		log.Printf("wallet file (%s) already exists.", WalletFile)
		return nil
	}

	number, _ := strconv.Atoi(ctx.Args().Get(0))
	if number == 0 {
		number = 10
	}
	log.Printf("generating %d accounts...", number)
	wallet := &WalletConfig{}
	for i := 0; i < number; i++ {
		pubkey, prikey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return err
		}
		hasher := sha3.New256()
		hasher.Write(pubkey)
		account := &AccountConfig{
			PrivateKey: crypto.PrivateKey(prikey),
			Address:    hasher.Sum([]byte{}),
		}
		wallet.Accounts = append(wallet.Accounts, account)
	}

	f, err := os.Create(WalletFile)
	if err != nil {
		log.Printf("cannot create wallet file in %s", WalletFile)
		return nil
	}
	defer f.Close()
	err = toml.NewEncoder(f).Encode(wallet)
	if err != nil {
		log.Printf("cannot encode toml file: %v", err)
	}

	return cmdListAccounts(ctx)
}

func cmdListAccounts(ctx *cli.Context) error {
	wallet, err := LoadAccounts(WalletFile)
	if err != nil {
		log.Fatal(err)
	}

	for addr := range wallet.Accounts {
		log.Printf("account: %s\n", addr)
	}

	return nil
}

func cmdMint(ctx *cli.Context) error {
	wallet, err := LoadAccounts(WalletFile)
	if err != nil {
		log.Fatal(err)
	}

	receiver, err := wallet.GetAccount(ctx.Args().Get(0))
	if err != nil {
		return err
	}

	amount, err := strconv.Atoi(ctx.Args().Get(1))
	if err != nil {
		return err
	}
	amountMicro := uint64(amount) * 1000000

	log.Printf("Please visit the following faucet service:")
	log.Printf("http://faucet.testnet.libra.org/?amount=%d&address=%s", amountMicro, hex.EncodeToString(receiver.Address))
	return nil
}