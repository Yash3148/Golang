// Single variable declaration
package main

import (
	"fmt"
	"os/user"
)

func main() {
	var city string = "Bhopal"
	fmt.Println(city)
}

// Bhopal

----------------------------------------------------------------------------------------------------------------------------

// Multi variable declaration
package main

import ("fmt") 

func main() {
	var name string = "KodeKloud"
	var user string = "Yoshita"
	fmt.Print("Welcome to ", name, ", ", user)
}

// Welcome to KodeKloud , Yoshita

----------------------------------------------------------------------------------------------------------------------------

package main

import ("fmt")

func main() {
	var name string = "KodeKloud"
	var user string = "Yoshita"
	fmt.Print(name)
	fmt.Print(user)
}

// KodeKloudYoshita

----------------------------------------------------------------------------------------------------------------------------

package main

import ("fmt")

func main() {
	var name string = "KodeKloud"
	var user string = "Yoshita"
	fmt.Println(name)
	fmt.Print(user)
}

// KodeKloud
// Yoshita

----------------------------------------------------------------------------------------------------------------------------

package main

import ("fmt")

func main() {
	var name string = "KodeKloud"
	var user string = "Yoshita"
	fmt.Println(name, "\n")       // \n is called the newline character. It is used to create a new line. Placed within string expressions. When inserted in a string, all the characters after \n are added to a mew line.
	fmt.Print(user)
}


// KodeKloud
// Yoshita

-----------------------------------------------------------------------------------------------------------------------------

package main

import ("fmt")

func main() {
	var name,user string = "KodeKloud" , "Yoshita"
	
	fmt.Println(name)       
	fmt.Print(user)
}

// KodeKloud
// Yoshita

-----------------------------------------------------------------------------------------------------------------------------

package main

import ("fmt")

func main() {
	var (
		name string = "Yoshita"
		age int = 21
	)
	
	fmt.Println(name)       
	fmt.Print(age)
}

// KodeKloud
// 21

-----------------------------------------------------------------------------------------------------------------------------