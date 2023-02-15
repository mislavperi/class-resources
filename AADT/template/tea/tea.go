package tea

type Template interface {
	BoilWater()
	BrewDrinkInWater(item string)
	PourInCup(item string)
	AddExtras(items string)
}

type Tea struct {
	Template Template
}

func NewTee(template Template) *Tea {
	return &Tea{
		Template: template,
	}
}

func (c *Tea) MakeCoffee() {
	c.Template.BoilWater()
	c.Template.BrewDrinkInWater("tea")
	c.Template.PourInCup("tea")
	c.Template.AddExtras("lemon")
}
