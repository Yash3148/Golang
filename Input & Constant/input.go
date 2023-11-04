// Single Input

package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there, ", name)
}

// output 
// Enter your name: Yoshita
// Hey there,  Yoshita

------------------------------------------------------------------------------------------------------------------------------

// Multiple Input

package main

import "fmt"

func main() {
	var name string
	var is_muggle bool
	fmt.Print("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &name, &is_muggle)
	fmt.Println(name, is_muggle)
}

// output
// Enter your name & are you a muggle: Yoshita false
// Yoshita false

------------------------------------------------------------------------------------------------------------------------------