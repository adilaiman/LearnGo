package main

// importing multiple packages
import (
	"fmt"
	"math"

	// import custom package, must state full path from src
	"github.com/adilaiman/LearnGo/03_packages/strutil"
)

func main() {
	// package methods always start with Caps
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Hello"))
}
