package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func main() {
	xi := []int{2, 32, 4, 5, 6}
	fmt.Println(xi)

	m := map[string]int{
		"todd": 23,
		"job":  33,
	}
	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	p1.speak()

	sa1 := secretAgent{
		person{
			"james",
			"bond",
		},
		true,
	}
	sa1.speak()
	sa1.person.speak()
}
