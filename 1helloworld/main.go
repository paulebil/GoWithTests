package main

import (
	"fmt"
	"os"
	"github.com/paulebil/GoWithTests/mocking"
)

const (
	spanish = "Spanish"
	french = "French" 

	englishHelloPrefix = "Hello, "
    spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}

	return greetingPrefix(language) + name
}



func greetingPrefix(language string) (prefix string){
	// using a named return (prefix string)
	// This will create a variable called `prefix` in your function
	// It will be assigned the `zero` value. This depends on the type,
	//  for example `int`s are ) and for `string`s it is ""

	switch language{
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// You can return whaterver it's set to by just calling `return` rather than `return prefix`
	return
}


func main(){
	mocking.Countdown(os.Stdout)
}