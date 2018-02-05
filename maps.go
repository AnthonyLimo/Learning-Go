package main

import "fmt"

func main() {
	presAge := make(map[string]int)
	presAge["Uhuru"] = 53
	presAge["Moi"] = 92

	fmt.Println(len(presAge))
	delete(presAge, "Moi")
	fmt.Println(len(presAge))
}
