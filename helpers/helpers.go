package helpers

import "strings"

// SnakeCaseToCamelCase transform snake to camel case, function copied from
// https://www.socketloop.com/tutorials/golang-underscore-or-snake-case-to-camel-case-example
func SnakeCaseToCamelCase(snake string) (camel string) {

	isToUpper := false

	for k, v := range snake {
		if k == 0 {
			camel = strings.ToUpper(string(snake[0]))
		} else {
			if isToUpper {
				camel += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camel += string(v)
				}
			}
		}
	}
	return
}
