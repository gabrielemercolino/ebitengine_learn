package abstract

import (
	_ "embed"
	"fmt"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	//go:embed abstract_art.kage
	abstractArt []byte
)

type AbstractArt struct {
	time   time.Time
	shader *ebiten.Shader
}

func (g *AbstractArt) Update() error {
	return nil
}

func (g *AbstractArt) Draw(screen *ebiten.Image) {
	shader := g.shader

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	cx, cy := ebiten.CursorPosition()

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"Time":       float32(time.Since(g.time).Nanoseconds()) / 1e9, // 1e9 converts to seconds
		"Cursor":     []float32{float32(cx), float32(cy)},
		"Resolution": []float32{float32(w), float32(h)},
	}

	screen.DrawRectShader(w, h, shader, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f", ebiten.ActualFPS()))
}

func (g *AbstractArt) Layout(_, _ int) (int, int) {
	return ebiten.WindowSize()
}

func Run() {
	// compile the shader
	shader, err := ebiten.NewShader(abstractArt)
	if err != nil {
		panic(err)
	}

	game := &AbstractArt{shader: shader, time: time.Now()}

	ebiten.SetWindowSize(600, 600)
	ebiten.SetWindowTitle("Shader test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
