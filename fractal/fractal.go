package fractal

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	//go:embed mandelbrot.kage
	mandelbrot_fast []byte

	//go:embed mandelbrot_double.kage
	mandelbrot_deep []byte
)

type Shader struct {
	name   string
	shader *ebiten.Shader
}

var (
	shaders       map[int]Shader
	currentShader int
)

type Fractal struct {
	offset        [2]float64
	iterations    uint64
	scalingFactor float64
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

	// toggle between shaders
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		currentShader++
		currentShader %= len(shaders)
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
	g.drawFractal(screen)

	{
		shader := shaders[currentShader]

		versionMsg := shader.name

		msg := fmt.Sprintf(
			"TPS: %.0f\nFPS: %.0f\nVersion: %s\nOffset[ WASD ]: %.2f:%.2f\nIterations[ UP | DOWN ]: %d\nScaling factor[ ALT | SPACE ]: %f",
			ebiten.ActualTPS(), ebiten.ActualFPS(), versionMsg, g.offset[0], g.offset[1], g.iterations, g.scalingFactor)

		ebitenutil.DebugPrint(screen, msg)
	}
}

func (g *Fractal) drawFractal(screen *ebiten.Image) {
	shader := shaders[currentShader]

	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	cx, cy := ebiten.CursorPosition()

	op := &ebiten.DrawRectShaderOptions{
		Uniforms: map[string]any{
			"Cursor":        []float32{float32(cx), float32(cy)},
			"Resolution":    []float32{float32(w), float32(h)},
			"Offset":        [2]float32{float32(g.offset[0]), float32(g.offset[1])},
			"Iterations":    float32(g.iterations),
			"ScalingFactor": float32(g.scalingFactor),
		},
	}
	screen.DrawRectShader(w, h, shader.shader, op)
}

func (g *Fractal) Layout(_, _ int) (int, int) {
	return ebiten.WindowSize()
}

func loadShader(shaderData []byte) *ebiten.Shader {
	shader, err := ebiten.NewShader(shaderData)
	if err != nil {
		log.Fatal(err)
	}

	return shader
}

func Run() {
	shaders = map[int]Shader{
		0: Shader{name: "normal", shader: loadShader(mandelbrot_fast)},
		1: Shader{name: "high res", shader: loadShader(mandelbrot_deep)},
	}

	currentShader = 0

	game := &Fractal{iterations: 100, scalingFactor: 1}

	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Shader test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
