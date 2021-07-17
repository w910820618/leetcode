package main

import "strings"

func main() {

}

func replaceSpace(s string) string {
	var res strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteString(string(s[i]))
		}
	}

	return res.String()
}
