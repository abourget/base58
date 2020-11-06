package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/mr-tron/base58/base58"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Pass something as an argument")
	}

	arg1 := os.Args[1]

	asHex, err := hex.DecodeString(arg1)
	if err != nil {
		asBase58, err := base58.Decode(arg1)
		if err != nil {
			fmt.Println("Not recognized as hex nor base58")
			os.Exit(1)
		}

		fmt.Println("Hex:", hex.EncodeToString(asBase58))
		os.Exit(0)
	}
	fmt.Println("Base58:", base58.Encode(asHex))
	os.Exit(0)
}
