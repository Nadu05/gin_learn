package main

import "fmt"

func main() {
	websites := map[string]string{
		"apple":  "https://www.apple.com",
		"google": "https://www.google.com",
		"aws":    "https://www.aws.com",
	}

	//add new
	websites["linkedin"] = "https://www.linkedin.com/"

	//delete
	websites["google"] = "https://www.google.com"
	fmt.Println(websites)
	fmt.Println(websites["apple"])
}
