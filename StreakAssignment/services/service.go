package services

import (
	"fmt"
	"task/findpairs/models"
)

//checks for valid input in the Request
func ValidateRequest(req models.Request) error {
	if len(req.Numbers) == 0 {
		return fmt.Errorf("Numbers array cannot be empty")
	}
	// Check if the target is an integer (implicit in the struct)
	// Check if all elements in Numbers are integers (implicit in the struct)
	return nil
}

//gives indices where the sum equals the target
func FindIndexPairs(numbers []int, target int) [][]int {
	pairs := [][]int{}
	indexMap := make(map[int]int) 

	for i, num := range numbers {
		complement := target - num
		if j, found := indexMap[complement]; found {
			pairs = append(pairs, []int{j, i}) // Adding pair
		}
		indexMap[num] = i // Storing index of the number
	}

	return pairs
}
