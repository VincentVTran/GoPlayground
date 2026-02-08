package objects

import "fmt"

// Car represents a car with basic attributes.
type Car struct {
	Make  string `json:"make"`  // The manufacturer of the car
	Model string `json:"model"` // The model of the car
	Year  int    `json:"year"`  // The manufacturing year of the car
}

// NewCar creates a new Car instance.
func NewCar(make, model string, year int) Car {
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
	}
}

// Method to display car details
func (c Car) DisplayDetails() {
	fmt.Printf("Car Details:\nMake: %s\nModel: %s\nYear: %d\n", c.Make, c.Model, c.Year)
}

// Method to check if the car is a classic (older than 25 years)
func (c Car) IsClassic() bool {
	currentYear := 2023
	return currentYear-c.Year > 25
}

// Method to update the car's year (uses a pointer receiver to modify the struct)
func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear
}

func DemoCar() {
	car := NewCar("Toyota", "Corolla", 1995)
	car.DisplayDetails()
	car.IsClassic()
	car.UpdateYear(2020)
}
