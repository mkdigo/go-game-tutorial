package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mkdigo/go-game-tutorial.git/assets"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bounds := image.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position.X -= halfW
	position.Y -= halfH

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0
	l.position.Y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(l.position.X, l.position.Y)

	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() React {
	bounds := l.image.Bounds()

	return NewReact(l.position.X, l.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
