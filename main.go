package main

func main() {
	LoadSampleData()
	LoadRoutes()
}

func LoadSampleData() {
	cars = append(cars, Car{ID: "1", Make: "Nissan", Model: "300ZX", Year: 1986,
		Transmission: &Transmission{Type: "Manual", Gears: 5}})
	cars = append(cars, Car{ID: "2", Make: "Honda", Model: "S2000", Year: 2005,
		Transmission: &Transmission{Type: "Manual", Gears: 6}})
}
