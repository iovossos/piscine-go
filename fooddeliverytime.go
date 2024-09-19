package piscine

import (
	"strings"
)

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	// Menu with food items and their preparation times
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	// Split the order string into individual items
	orderItems := strings.Split(order, " ")
	totalTime := 0

	// Calculate total preparation time for the order
	for _, item := range orderItems {
		if f, exists := menu[item]; exists {
			totalTime += f.preptime
		} else {
			return 404 // Return 404 if an item is not found in the menu
		}
	}

	return totalTime
}
