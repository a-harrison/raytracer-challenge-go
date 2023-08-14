package playground 

import (
	"github.com/a-harrison/raytracer-challenge-go/tuple"
)

type Projectile struct {
	Position tuple.Point
	Velocity tuple.Vector
}

type Environment struct {
 	Gravity, Wind tuple.Vector
}

func Tick(env Environment, proj Projectile) Projectile {		
	newPosition, _ := tuple.Add(&proj.Position, &proj.Velocity)
	positionAsPoint := newPosition.(*tuple.Point)

	newVelocity, _ := tuple.Add(&proj.Velocity, &env.Gravity)
	newVelocity, _ = tuple.Add(newVelocity, &env.Wind)
	velocityAsVector := newVelocity.(*tuple.Vector)

	return Projectile { 
		Position: *positionAsPoint,
		Velocity: *velocityAsVector,
	}	
}