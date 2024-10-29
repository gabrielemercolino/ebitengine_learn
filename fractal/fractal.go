package fractal

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed mandelbrot.kage
var mandelbrot []byte

type Fractal struct {
	shader        *ebiten.Shader
	offset        [2]float32
	iterations    float32
	scalingFactor float32
}

func (g *Fractal) Update() error {
	// movement
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.offset[0] += .01 * g.scalingFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.offset[0] -= .01 * g.scalingFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.offset[1] += .01 * g.scalingFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.offset[1] -= .01 * g.scalingFactor
	}

	// scaling
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		g.scalingFactor -= 0.01 * g.scalingFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.scalingFactor += 0.01 * g.scalingFactor
	}

	// iterations
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.iterations++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		if g.iterations > 0 {
			g.iterations--
		}
	}

	return nil
}

func (g *Fractal) Draw(screen *ebiten.Image) {
	shader := g.shader

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	cx, cy := ebiten.CursorPosition()

	op := &ebiten.DrawRectShaderOptions{}
	op.Uniforms = map[string]any{
		"Cursor":        []float32{float32(cx), float32(cy)},
		"Resolution":    []float32{float32(w), float32(h)},
		"Offset":        g.offset,
		"Iterations":    g.iterations,
		"ScalingFactor": g.scalingFactor,
	}

	screen.DrawRectShader(w, h, shader, op)
	msg := fmt.Sprintf(
		"TPS: %.0f\nFPS: %.0f\nOffset[ WASD ]: %.2f:%.2f\nIterations[ UP | DOWN ]: %.0f\nScaling factor[ ALT | SPACE ]: %.5f",
		ebiten.ActualTPS(), ebiten.ActualFPS(), g.offset[0], g.offset[1], g.iterations, g.scalingFactor)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Fractal) Layout(_, _ int) (int, int) {
	return ebiten.WindowSize()
}

func Run() {
	// compile the shader
	shader, err := ebiten.NewShader(mandelbrot)
	if err != nil {
		panic(err)
	}

	game := &Fractal{shader: shader, iterations: 100, scalingFactor: 1}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Shader test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
