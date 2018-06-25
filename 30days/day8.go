package main

import "fmt"

func main() {
	// Create map and specify number of inputs
	phoneBook := make(map[string]int)
	var num int
	fmt.Scanf("%d", &num)

	for i := 0; i < num; i++ {
		// Set the map up with inputs
		var name string
		var phoneNumber int
		fmt.Scanf("%s", &name)
		fmt.Scanf("%d", &phoneNumber)
		phoneBook[name] = phoneNumber
	}

	// Look up name in map and print if found
	for {
		var lookup_contact string
		// Was getting timeouts without googling how to do this..
		_, err := fmt.Scanf("%s", &lookup_contact)
		if err != nil {
			break
		}
		phoneNumber, exist := phoneBook[lookup_contact]
		if exist {
			fmt.Printf("%s=%d\n", lookup_contact, phoneNumber)
		} else {
			fmt.Println("Not found")
		}
	}
}
