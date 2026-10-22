package main

import "fmt"

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Println(err)

		return
	}

	for range departmentCount {
		minTemp, maxTemp := 15, 30
		var employeeCount int

		if _, err := fmt.Scan(&employeeCount); err != nil {
			fmt.Println(err)

			return
		}

		for range employeeCount {
			var action string
			var value int

			if _, err := fmt.Scan(&action, &value); err != nil {
				fmt.Println(err)

				return
			}

			switch action {
			case ">=":
				minTemp = max(minTemp, value)
			case "<=":
				maxTemp = min(maxTemp, value)
			default:
				fmt.Println("Incorrect action: type >=/<=")

				return
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
