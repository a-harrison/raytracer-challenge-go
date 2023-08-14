package playground

import (
	"fmt"

	"github.com/a-harrison/raytracer-challenge-go/tuple"
)

func PrintProjectilePath() {
	startingPosition := tuple.NewPoint(0, 1, 0)
	startingVelocity := tuple.NewVector(1, 1, 0)

	projectile := Projectile{
		Position: startingPosition,
		Velocity: tuple.Normalize(&startingVelocity),
	}

	environment := Environment{
		Gravity: tuple.NewVector(0, -0.1, 0),
		Wind:    tuple.NewVector(-0.01, 0, 0),
	}

	counter := 0
	for true {
		counter++
		projectile = Tick(environment, projectile)
		fmt.Printf("[%v] - Position: %v\n", counter, projectile.Position)

		if projectile.Position.Y() <= 0 {
			break
		}
	}
}
