package jgrpcclient

// Numeric is an interface support numeric calculation
type Numeric interface {
	Add(a, b int) int
}

// Calculator is a struct implement numeric
type Calculator struct {
	Config map[string]string
}

// Add is a function to calculate a + b
func (c *Calculator) Add(a, b int) int {
	return add(c, a, b)
}
