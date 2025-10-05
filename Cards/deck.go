package main

import "fmt"

type deck []string

func (d deck) print() {
	for idx, card := range d {
		fmt.Printf("%d : %s \n", idx, card)
	}
}
