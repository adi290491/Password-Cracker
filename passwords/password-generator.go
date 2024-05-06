package passwords

var validChars string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVW!@#$%^&*~_-+="

func GeneratePasswords(length int) []string {

	var passwords []string

	generate("", length, &passwords)

	return passwords
}

func generate(prefix string, length int, passwords *[]string) {

	if length == 0 {
		*passwords = append(*passwords, prefix)
		return
	}

	for _, char := range validChars {
		generate(prefix+string(char), length-1, passwords)
	}

}
