package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	// calculate
	average := (math + science + english + indonesia) / 4

	switch {
	case average == 100:
		return "Sempurna"
	case average >= 90:
		return "Sangat Baik"
	case average >= 80:
		return "Baik"
	case average >= 70:
		return "Cukup"
	case average >= 60:
		return "Kurang"
	default:
		return "Sangat kurang"
	}
}
    //debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))  // Output: "Cukup"
	fmt.Println(GetPredicate(100, 100, 100, 100)) // Output: "Sempurna"
	fmt.Println(GetPredicate(85, 75, 90, 80))    // Output: "Baik"
	fmt.Println(GetPredicate(59, 60, 58, 57))    // Output: "Sangat kurang"
}
