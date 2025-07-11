package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 400
	screenHeight = 600
)

// Player represents the player character.
type Player struct {
	X, Y   float64
	Width  float64
	Height float64
	Image  *ebiten.Image
}

// Game represents the main game state.
type Game struct {
	player *Player
}

// NewGame initializes the game.
func NewGame() *Game {
	playerImage := ebiten.NewImage(30, 50)
	playerImage.Fill(color.White)

	player := &Player{
		X:      screenWidth / 2,
		Y:      screenHeight / 2,
		Width:  30,
		Height: 50,
		Image:  playerImage,
	}

	return &Game{
		player: player,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Will be implemented in PR #2
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Fill the background with light blue
	screen.Fill(color.RGBA{R: 173, G: 216, B: 230, A: 255})

	// Draw the player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.X, g.player.Y)
	screen.DrawImage(g.player.Image, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Doodle Jump")

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}