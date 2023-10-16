package fonctions

import "fmt"

func Scan() string {
	answer := ""
	fmt.Println("Entrez une lettre ou un mot")
	fmt.Scan(&answer)
	valid := true
	for _, c := range answer {
		if !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') && !(c == '-') {
			valid = false
		}
	}
	if valid {
		if len(answer) > 1 {
			return answer
		} else {
			return answer
		}
	} else {
		return Scan()
	}
}
