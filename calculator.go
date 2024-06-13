package main

type Calculator struct{}

type Calculate interface {
	Addition(values ...float64) float64
	Multiplication(values ...float64) float64
	Subtraction(x, y float64) float64
	Division(x, y float64) (float64, error)
}

func (c Calculator) Addition(values ...float64) float64 {

	total := 0.0
	for _, value := range values {
		total += value
	}
	return total
}

func (c Calculator) Subtraction(x, y float64) float64 {
	return x - y
}

func (c Calculator) Multiplication(values ...float64) float64 {

	total := 1.0
	for _, value := range values {
		total *= value
	}
	return total
}

func (c Calculator) Division(x, y float64) (float64, error) {
	if y == 0.0 {
		panic("cannot divide by zero")
	}

	return x / y, nil
}
