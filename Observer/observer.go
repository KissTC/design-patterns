package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// producto no disponible
// cuando haya items avise a los observadores que ya hay
type Item struct {
	obsevers  []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.obsevers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(observer Observer) {
	i.obsevers = append(i.obsevers, observer)
}

type EmailClient struct {
	id string
}

func (e *EmailClient) getId() string {
	return e.id
}

func (e *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n", value, e.id)
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}
	secondObserver := &EmailClient{
		id: "34dc",
	}

	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}
