package main

import "fmt"

// Product es la interfaz común para los productos
type Product interface {
	Use() string
}

// ConcreteProductA es una implementación de Product
type ConcreteProductA struct{}

// Use implementa el método Use de la interfaz Product
func (p *ConcreteProductA) Use() string {
	return "Product A"
}

// ConcreteProductB es otra implementación de Product
type ConcreteProductB struct{}

// Use implementa el método Use de la interfaz Product
func (p *ConcreteProductB) Use() string {
	return "Product B"
}

type ConcreteProductC struct{}

func (p *ConcreteProductC) Use() string {
	return "Product C"
}

// Creator es la interfaz para la fábrica abstracta
type Creator interface {
	CreateProduct() Product
}

// ConcreteCreatorA es una implementación de Creator que crea ConcreteProductA
type ConcreteCreatorA struct{}

// CreateProduct crea una instancia de ConcreteProductA
func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB es otra implementación de Creator que crea ConcreteProductB
type ConcreteCreatorB struct{}

// CreateProduct crea una instancia de ConcreteProductB
func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

type ConcreteCreatorC struct{}

func (c *ConcreteCreatorC) CreateProduct() Product {
	return &ConcreteProductC{}
}

func CreateProductByName(name string) Product {
	if name == "A" {
		return &ConcreteProductA{}
	}

	if name == "B" {
		return &ConcreteProductB{}
	}

	if name == "C" {
		return &ConcreteProductC{}
	}

	return nil
}

func main() {
	// Crear un ConcreteCreatorA
	creatorA := CreateProductByName("A")
	//productA := creatorA.
	fmt.Println(creatorA.Use()) // Output: Product A

	// Crear un ConcreteCreatorB
	creatorB := CreateProductByName("B")
	//productB := creatorB.CreateProduct()
	fmt.Println(creatorB.Use()) // Output: Product B

	creatorC := CreateProductByName("C")
	//productC := creatorC.CreateProduct()
	fmt.Println(creatorC.Use())
}
