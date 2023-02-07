package main

import "fmt"

func main() {
	magazine1, _ := newPublication("magazine", "examples_magazine", 22, "La capital")
	newspaper, _ := newPublication("newspaper", "example_newspaper", 33, "La nacion")

	pubDetails(magazine1)
	pubDetails(newspaper)

}

func pubDetails(publication iPublication) {
	fmt.Printf("-------------------\n")
	fmt.Printf("%s\n", publication)
	fmt.Printf("Name: %s\n", publication.getName())
	fmt.Printf("Pages: %d\n", publication.getPages())
	fmt.Printf("Publisher: %s\n", publication.getPublisher())
	fmt.Printf("-------------------\n")
}
