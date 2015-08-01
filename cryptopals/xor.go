package main

import (

	"fmt"
	"encoding/hex"
)

func XOR(string_one, string_two string) (string, error) {
	
	buff_one, err := hex.DecodeString(string_one)
	if err != nil {
		return "", err
	}
	buff_two, err := hex.DecodeString(string_two)
	if err != nil {
		return "", err
	}

	res := make([]byte, len(buff_one))
	for i := range(buff_one) {
		res[i] = buff_one[i] ^ buff_two[i]
	}
	return hex.EncodeToString(res), nil
}

const str_one = "1c0111001f010100061a024b53535009181c"

const str_two = "686974207468652062756c6c277320657965"

func main() {
	fmt.Println(XOR(str_one,str_two))
}