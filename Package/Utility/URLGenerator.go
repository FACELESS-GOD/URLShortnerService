package Utility

import (
	"math/rand"
	"strings"
)

func RandomURLString() string {
	var baseString string = "POIUYTREWQLKJHGFDSAMNBVCXZqwertyuioplkjhgfdsazxcvbnm1234567890"
	chars := strings.Split(baseString, "")
	var randomString string = ""

	for i := 0; i < len(baseString); i++ {
		lastIndex := len(baseString) - 1
		randomindex := rand.Intn(lastIndex)

		randomString = randomString + string(chars[randomindex])
	}

	return randomString
}
