package main

import (
	"fmt"

	"github.com/noisemaster/FrinkiacAPIGo"
)

func main() {
	fmt.Println("Let's search for a still of steamed hams")
	link, _ := frinkiac.GetFrinkiacFrame("Steamed Hams")
	fmt.Println(link)
	fmt.Println("What about that with a caption")
	link, _ = frinkiac.GetFrinkiacMeme("Steamed Hams")
	fmt.Println(link)
	fmt.Println("What about a gif from Morbotron")
	link, _ = frinkiac.GetMorbotronGif("Smells good")
	fmt.Println(link)
}
