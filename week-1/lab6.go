package main

func xorEncrypt(text string, key byte) string {

	result := ""

	for i := 0; i < len(text); i++ {

		result += string(text[i] ^ key)

	}

	return result

}