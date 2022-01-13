package main

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func encode(id int) string {
	resultUrl := ""
	dividend := id
	reminder := 0

	for dividend > 0 {
		reminder = dividend % 62
		dividend = dividend / 62
		resultUrl += string(alphabet[reminder])
	}

	return resultUrl
}
