package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	var answer string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	for _, v := range sd {
		whatIsIt += fmt.Sprintf("%s", (string(v)))
	}
	// fmt.Println("Ascii code from secret")
	// fmt.Println(sd)

	fmt.Println("Answer from secret")
	fmt.Println(whatIsIt)

	for i := len(whatIsIt) - 1; i >= 0; i-- {
		answer += fmt.Sprint(string(whatIsIt[i]))
	}

	fmt.Println("Answer after reverse")
	fmt.Println(answer)
}
