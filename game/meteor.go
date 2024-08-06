package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mkdigo/go-game-tutorial.git/assets"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {
	meteorSprites := assets.MeteorSprites
	image := meteorSprites[int(rand.Float64()*float64(len(meteorSprites)))]

	speed := (rand.Float64() * 13) // Max 13

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.image, op)
}

func (m *Meteor) Collider() React {
	bounds := m.image.Bounds()

	return NewReact(m.position.X, m.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
