package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of Passengers", v.NumberOfPassengers)
	fmt.Println("Number of Wheels", v.NumberOfWheels)

}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Electric:", c.IsElectric)
	fmt.Println("Hybrid:", c.IsHybrid)
	c.Vehicle.showDetails()
}
func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2020,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}
	volvoXC90.show()
	fmt.Println()

}
