package main

import "fmt"

func main() {

	/* Map declartions
				   var myGreeting = make(map[string]string)
			       myGreeting := map[string]string{}
		           elements := map[string]map[string]string{
	                   "H": map[string]string{
	                       "name": "Hydrogen",
	                       "state": "gas".
	                   }
	               }
	*/
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Buenos dias!",
		3: "Bongiorno",
	}

	fmt.Println(myGreeting)

	//delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value does not exists")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}

}
