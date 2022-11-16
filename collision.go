package gengine

func ParticleCollision(a *Particle, b *Particle) bool {
	dx := a.position.x - b.position.x
	dy := a.position.y - b.position.y

	if dx > 0.0 || dy > 0.0 {
		return false
	}

	return true
}
