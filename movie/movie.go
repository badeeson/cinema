package movie

import "fmt"

func Review(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating is %f\n", name, rating)
}

func FindName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Avengers: Endgame"
	case "tt1825763":
		return "Black Panther"
	}

	return "Not Found."
}
