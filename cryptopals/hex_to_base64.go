package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

//Function that turns hex into base64
func hexToBase64(hex_str string)(string, error) {

	x, err := hex.DecodeString(hex_str)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(x), nil
}

const hex_string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

func main() {
	b64, err := hexToBase64(hex_string)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(b64)
}