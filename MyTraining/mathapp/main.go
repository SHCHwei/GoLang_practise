package mathapp

import (
	"fmt"
	"log"
	"abc.com/mymath"
)

func mathapp(){
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))

	log.SetPrefix("sqrt: ")
	log.SetFlags(0)
	message, err := mymath.Hello("")

	if err != nil {
        log.Fatal(err)
    }

	fmt.Printf(message)
}


