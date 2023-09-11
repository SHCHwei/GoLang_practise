package mymath

import (
	"fmt"
	"errors"
)

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0 ; i < 1000 ; i++{
		z -= (z*z-x) / (2 * x)
	}
	return z
}


func Hello(name string)(string, error){
	if name == ""{
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message, nil
}