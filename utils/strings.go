package utils

import "strings"

func RemoveInvalidChars(str string) string {
	str = strings.Replace(str, ",", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, "(", "", -1)

	return str
}
