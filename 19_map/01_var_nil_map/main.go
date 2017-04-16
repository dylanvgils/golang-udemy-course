package main

import "fmt"

func main() {
	var myGreeting map[string]string

	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
}

/*
  When you try to add a value to this example, you will get a error:
    "panic: assignment to entry in nil map"

  for example add these lines:
    myGreeting["John"] = "Good morning."
    myGreeting["Do"] = "Bonjour."
*/
