package lasagna

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (int, float64) {
	var noodleLayers int = 0
	var sauceLayers int = 0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleLayers++
		}
		if layer == "sauce" {
			sauceLayers++
		}
	}

	return noodleLayers * 50, float64(sauceLayers) * 0.2
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	var scaledPortions = []float64{}
	for _, quantity := range quantities {
		scaledPortions = append(scaledPortions, quantity*float64(numberOfPortions)/2)
	}
	return scaledPortions
}
