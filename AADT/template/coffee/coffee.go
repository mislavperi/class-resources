package coffee

type Template interface {
	BoilWater()
	BrewDrinkInWater(item string)
	PourInCup(item string)
	AddExtras(items string)
}

type Coffee struct {
	Template Template
}

func NewCoffee(template Template) *Coffee {
	return &Coffee{
		Template: template,
	}
}

func (c *Coffee) MakeCoffee() {
	c.Template.BoilWater()
	c.Template.BrewDrinkInWater("coffee")
	c.Template.PourInCup("coffee")
	c.Template.AddExtras("sugar and milk")
}
