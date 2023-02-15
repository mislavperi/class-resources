package template

import "fmt"

type Template struct {
}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) BoilWater() {
	fmt.Println("Boiling water")
}

func (t *Template) BrewDrinkInWater(item string) {
	fmt.Printf("Brewing %s in water", item)
}

func (t *Template) PourInCup(item string) {
	fmt.Printf("Pouring %s in cup", item)
}

func (t *Template) AddExtras(items string) {
	fmt.Printf("Add %s", items)
}
