package main

import (
	"fmt"
	"os"
	"strings"
)

var flag bool = true

type temperature struct {
	name  string
	value float64
}

type (
	celcius    float64
	kelvin     float64
	fahrenheit float64
)

func (c celcius) convert() {
	k := c + 273
	f := c*(9/5) + 32
	fmt.Println(c, " Celcius is ", k, " Kelvin and ", f, " Fahremheit.")
}

func (k kelvin) convert() {
	c := k - 273
	f := (9/5)*(k-273) + 32
	fmt.Println(k, " Kelvin is ", c, " Celcius and ", f, " Fahremheit.")
}

func (f fahrenheit) convert() {
	c := (f - 32) * 5 / 9
	k := (5*f + 2297) / 9
	fmt.Println(f, " Fahremheit is ", c, " Celcius and ", k, " Kelvin.")
}

type convert interface {
	convert()
}

func main() {
	fmt.Println("\nCelcius Kelvin Fahrenheit Converter\n.")
	fmt.Println("You can calculate using Celcius, Fahrenheit or Kelvin")

	var value convert

	for flag {
		var fName string
		var arg float64
		fmt.Println("> Enter name of the type or enter x to exit:")
		_, err := fmt.Scanf("%s", &fName)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		//
		if fName == "x" {
			flag = false
		} else {
			fmt.Println("> Enter a value for the calculation:")
			_, err := fmt.Scanf("%f", &arg)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
			switch strings.ToLower(fName) {
			case "celcius":
				value = celcius(arg)
				value.convert()
			case "fahrenheit":
				value = fahrenheit(arg)
				value.convert()
			case "kelvin":
				value = kelvin(arg)
				value.convert()
			default:
				fmt.Println("Wrong Type!!")
			}

		}
	}
}
