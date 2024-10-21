package app

import (
	"fmt"
)

func ExampleValidateURL() {
	url := "https://ya.ru"

	err := ValidateURL(url)
	if err == nil {
		fmt.Println("URL is valid")
	} else {
		fmt.Println("URL is not valid")
	}

	// Output:
	// URL is valid
}
