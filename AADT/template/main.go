package main

import (
	"example.com/template/coffee"
	"example.com/template/tea"
	"example.com/template/template"
)

func main() {
	template := template.NewTemplate()
	coffee := coffee.NewCoffee(template)
	tea := tea.NewTee(template)
	coffee.MakeCoffee()
	tea.MakeCoffee()
}
