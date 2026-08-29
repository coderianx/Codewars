package kata

import "fmt"

func Greet(name string) string {
	if name == "Johnny" {
		return "Hello, my love!"
	}

	return fmt.Sprintf("Hello, %s!", name)
}
