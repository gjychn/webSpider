package util

import "regexp"

func CheckHttpUrl(url string) bool {
	reg := `http(s?)://[-.\w]+.*`

	r := regexp.MustCompile(reg)
	if r.MatchString(url) {
		return true
	}
	return false
}
