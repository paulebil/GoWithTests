package dependencyinjection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string){
	fmt.Fprintf(writer, "Hello, %s", name)
}