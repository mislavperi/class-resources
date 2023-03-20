package converter

type ConverterC struct {

}

type ConverterF struct {

}

func (c *ConverterC) Convert(f int) int {
	return (f - 32) * 5/9
}

func (c *ConverterF) Convert(cl int) int {
	return (cl * 9/5) + 32
}