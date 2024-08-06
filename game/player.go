package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mkdigo/go-game-tutorial.git/assets"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}

	return &Player{
		image:             image,
		position:          position,
		game:              game,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0

	p.laserLoadingTimer.Update()

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()
		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2
		spawnPos := Vector{
			X: p.position.X + halfW,
			Y: p.position.Y - halfH/2,
		}
		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() React {
	bounds := p.image.Bounds()

	return NewReact(p.position.X, p.position.Y, float64(bounds.Dx()), float64(bounds.Dy()))
}
