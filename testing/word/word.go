package word

import "unicode"

// IsPalindrome reports whether a reads the  same forward and backward
func IsPalindrome(s string) bool {
	var letters []rune

	for _, l := range s {
		if unicode.IsLetter(l) {
			letters = append(letters, unicode.ToLower(l))

		}
	}

	// bad example, use range
	//for i:=0; i<len(s); i++{
	//	if unicode.IsLetter(rune(s[i])) {
	//		letters = append(letters, unicode.ToLower(rune(s[i])))
	//	}
	//}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
