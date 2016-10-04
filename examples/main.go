package main

import (
	"fmt"

	"github.com/noisemaster/FrinkiacAPIGo"
)

func main() {
	fmt.Println("Let's search for a still of steamed hams")
	fmt.Println(frinkiac.GetFrinkiacFrame("Steamed Hams"))
	fmt.Println("What about that with a caption")
	fmt.Println(frinkiac.GetFrinkiacMeme("Steamed Hams"))
	fmt.Println("What about a gif from Morbotron")
	fmt.Println(frinkiac.GetMorbotronGif("Smells good"))
}
