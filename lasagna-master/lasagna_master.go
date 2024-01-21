package lasagna

import "slices"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}

	length := len(layers)

	return length * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodleCount := 0
	sauceCount := 0

	for _, item := range layers {
		if item == "noodles" {
			noodleCount++
		}
		if item == "sauce" {
			sauceCount++
		}
	}

	noodles = noodleCount * 50
	sauce = float64(sauceCount) * 0.2

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	index := slices.Index(myList, "?")
	friendsListLength := len(friendsList)
	secretIngredient := friendsList[friendsListLength-1]
	slices.Replace(myList, index, index+1, secretIngredient)
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	scaleFactor := float64(portions) / 2

	for i := 0; i < len(quantities); i++ {
		scaledQuantities = append(scaledQuantities, quantities[i]*scaleFactor)
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
