package main

import (
	"fmt"

	"example.com/arrays-practice/product"
)

func main() {
	var hobbies = [3]string{"guitar", "woodworking", "photography"} // 1
	fmt.Println(hobbies)

	fmt.Println(hobbies[0]) // 2
	fmt.Println(hobbies[1:])

	sliced := hobbies[0:2] // 3
	fmt.Println(sliced)
	fmt.Println(hobbies[:2])

	fmt.Println(sliced[1:3]) // 4

	var goals = []string{"learn syntax", "be able to utilize"} // 5
	fmt.Println(goals)

	goals[1] = "prepare to commercial use" // 6
	goals = append(goals, "code complex application")
	fmt.Println(goals)

	var products = []product.Product{createProduct("carrot", 2.3), createProduct("potato", 5.23)} // 7
	products = append(products, createProduct("tomato", 10.21))
	fmt.Println(products)
}

func createProduct(title string, price float32) product.Product {
	product, error := product.New(title, price)

	if error != nil {
		panic(error)
	}

	return *product
}
