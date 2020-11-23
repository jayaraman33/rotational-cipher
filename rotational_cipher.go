package rotationalcipher

func RotationalCipher(plainText string, key int) string {
	shift := []rune(plainText)
	for i, char := range plainText {
		switch {

		case char >= 'a' && char <= 'z':
			shift[i] = 'a' + (char-'a'+rune(key))%26
		case char >= 'A' && char <= 'Z':
			shift[i] = 'A' + (char-'A'+rune(key))%26
		}
	}
	return string(shift)
}
