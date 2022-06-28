package main

import (
	"fmt"
	"time"

	"math/rand"
	//	"bufio"
	//	"io"

	"os"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var letters = []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ")

func makeRandomPassphrase(size int) string {
	rand.Seed(time.Now().UnixNano())
	out := make([]rune, size)
	for i := range out {
		out[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(string(out))
	return string(out)
}

func main() {
	// Generate a mnemonic for memorization or user-friendly seeds

	numberOfAccountToMake := 10000
	secretPassphrase := makeRandomPassphrase(16)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password

	outFile, err := os.Create("accounts.csv")
	check(err)
	defer outFile.Close()

	for i := 0; i < numberOfAccountToMake; i++ {
		entropy, _ := bip39.NewEntropy(128)
		mnemonic, _ := bip39.NewMnemonic(entropy)
		seed := bip39.NewSeed(mnemonic, secretPassphrase)

		masterKey, _ := bip32.NewMasterKey(seed)
		publicKey := masterKey.PublicKey()
		fmt.Fprint(outFile, i)
		outFile.WriteString(",")
		fmt.Fprint(outFile, mnemonic)
		outFile.WriteString(",")
		fmt.Fprint(outFile, masterKey)
		outFile.WriteString(",")
		fmt.Fprint(outFile, publicKey)
		outFile.WriteString(",")
		fmt.Fprint(outFile, secretPassphrase)
		// Display mnemonic and keys
		fmt.Println("Mnemonic: ", mnemonic)
		fmt.Println("Master private key: ", masterKey)
		fmt.Println("Master public key: ", publicKey)
		fmt.Println("secretPassphrase: ", secretPassphrase)
		outFile.WriteString("\n")
		secretPassphrase = makeRandomPassphrase(16)
	}
}
