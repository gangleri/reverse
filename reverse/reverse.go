package reverse

func Reverse(original string) string {
	newstring := ""
	for i := len(original) - 1; i >= 0; i-- {
		newstring += string(original[i])
	}
	return newstring
}
