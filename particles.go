package gengine

import "fmt"

type Vector struct {
	x float64
	y float64
}

type Particle struct {
	position     Vector
	velocity     Vector
	acceleration Vector
	force        Vector
	mass         float64
}

// Prints all particles' position to the output.
func PrintParticle(particle *Particle) {
	fmt.Print("{ particle :\n")
	fmt.Printf("  pos (%f,%f)\n", particle.position.x, particle.position.y)
	fmt.Printf("  vect vel (%f,%f)\n", particle.velocity.x, particle.velocity.y)
	fmt.Printf("  vect accel (%f,%f)\n", particle.acceleration.x, particle.acceleration.y)
	fmt.Printf("  vect force (%f,%f)\n", particle.force.x, particle.force.y)
	fmt.Printf("  mass (%f) }\n", particle.mass)
	fmt.Print("}\n")
}

// Initializes all particles with random positions, zero velocities and given mass.
func InitializeParticle(mass float64) Particle {
	particle := Particle{
		position:     Vector{x: 0, y: 0},
		velocity:     Vector{x: 0, y: 0},
		acceleration: Vector{x: 0, y: 0},
		force:        Vector{x: 0, y: 0},
		mass:         mass}
	return particle
}

//Compute force vector applied to a particle
func ComputeForce(particle *Particle, forceX float64, forceY float64) {
	particle.force = Vector{x: particle.mass * forceX, y: particle.mass * forceY}
}

// Compute Acceleration of a particle based on his force vector
func ComputeAcceleration(particle *Particle) {
	particle.acceleration = Vector{x: particle.force.x / particle.mass, y: particle.force.y / particle.mass}
}

// Compute velocity vector of a particle in a given time based on its acceleration
func ComputeVelocity(particle *Particle, time float64) {
	particle.velocity = Vector{
		x: particle.velocity.x + (particle.acceleration.x * time),
		y: particle.velocity.y + particle.acceleration.y*time,
	}
}

// Compute position of a particle in a given time based on its velocity
func ComputePosition(particle *Particle, time float64) {
	particle.position = Vector{
		x: particle.position.x + particle.velocity.x*time,
		y: particle.position.x + particle.velocity.y*time,
	}
}

// Simulate a particle in space give force vector (m/s^2 or N) and delta time
func SimulateParticle(particle *Particle, forceX float64, forceY float64, deltaTime float64) {
	ComputeForce(particle, forceX, forceY)
	ComputeAcceleration(particle)
	ComputeVelocity(particle, deltaTime)
	ComputePosition(particle, deltaTime)
}
