package main

func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.8 + 32
	return []float64{kelvin, fahrenheit}
}
