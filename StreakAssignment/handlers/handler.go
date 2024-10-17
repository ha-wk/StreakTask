package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"task/findpairs/models"
	"task/findpairs/services"
	"task/findpairs/utils"
)

//finds all pairs of indices whose sum equals the target value
func FindPairs(w http.ResponseWriter, r *http.Request) {
	utils.Log("Received models.Request: %s %s", r.Method, r.URL)

	// Read the models.Request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.SendError(w, "Error reading models.Request body", http.StatusInternalServerError)
		utils.Log("Error reading models.Request body: %v", err)
		return}
	utils.Log("models.Request body: %s", string(body))

	// Decode the models.Request body
	var req models.Request
	if err := json.Unmarshal(body, &req); err != nil {
		utils.SendError(w, "Invalid models.Request payload", http.StatusBadRequest)
		utils.Log("Error decoding models.Request body: %v", err)
		return}

	// Validate input
	if err := services.ValidateRequest(req); err != nil {
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		utils.Log("Validation error: %v", err)
		//return
		return
}

	solutions := services.FindIndexPairs(req.Numbers, req.Target)
	utils.SendResponse(w, map[string]interface{}{"solutions": solutions}, http.StatusOK)
	//utils.SendResponse(w, map[string]interface{}{"solutions": solutions}, http.StatusOK)
	utils.Log("Solutions found: %v", solutions)
}

// // validatemodels.Request checks for valid input in the models.Request
// func validateRequest(req models.Request) error {
// 	if len(req.Numbers) == 0 {
// 		return fmt.Errorf("Numbers array cannot be empty")
// 	}
// 	// Check if the target is an integer (implicit in the struct)
// 	// Check if all elements in Numbers are integers (implicit in the struct)
// 	return nil
// }

// // findIndexPairs returns pairs of indices where the sum equals the target
// func findIndexPairs(numbers []int, target int) [][]int {
// 	pairs := [][]int{}
// 	indexMap := make(map[int]int) // To store index of numbers for quick look-up

// 	for i, num := range numbers {
// 		complement := target - num
// 		if j, found := indexMap[complement]; found {
// 			pairs = append(pairs, []int{j, i}) // Add the pair
// 		}
// 		indexMap[num] = i // Store the index of the number
// 	}

// 	return pairs
// }