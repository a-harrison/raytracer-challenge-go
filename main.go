package main

import (
	"fmt"
	"os"

	"github.com/a-harrison/raytracer-challenge-go/playground"
)

func main() {
	if len(os.Args) > 1 {
		switch choice := os.Args[1]; choice {
		case "1":
			playground.PrintProjectilePath()
		default:
			fmt.Println("Unknown option - please select from an option below.")
		}
	} else {
		fmt.Println("Please run program with an option:")
		fmt.Println()
		fmt.Println("./main {option}")
		fmt.Println()
		fmt.Println("Options: ")
		fmt.Println("    '1' - Print Projectile Path")
		fmt.Println("     ")
	}
}
