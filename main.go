package main

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const sizeDot = 8

var pointerValid = ebiten.NewImage(sizeDot, sizeDot)
var pointerWait = ebiten.NewImage(sizeDot, sizeDot)

type Game struct {
	imgebiten *ebiten.Image
	Points    []Point
	Px, Py    int
}

type Point struct {
	X, Y int
	Name string
}

func (g *Game) Update() error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.Px, g.Py = ebiten.CursorPosition()
	} else {

		keys := inpututil.AppendPressedKeys(make([]ebiten.Key, 0, 10))

		if containKey(keys, ebiten.KeyEnter) && g.Px != -1 {
			fmt.Printf("Point's %d name: ", len(g.Points)+1)
			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			g.Points = append(g.Points, Point{g.Px, g.Py, strings.Trim(line, "\n")})
			g.Px = -1
			g.Py = -1
		}
	}
	return nil
}

func containKey(samples []ebiten.Key, target ebiten.Key) bool {
	for _, k := range samples {
		if k.String() == target.String() {
			return true
		}
	}
	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.imgebiten, nil)
	for _, p := range g.Points {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(p.X-sizeDot/2), float64(p.Y-sizeDot/2))
		screen.DrawImage(pointerValid, op)
	}
	if g.Px != -1 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.Px-sizeDot/2), float64(g.Py-sizeDot/2))
		screen.DrawImage(pointerWait, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.imgebiten.Bounds().Dx(), g.imgebiten.Bounds().Dy()
}

func main() {
	pointerValid.Fill(color.RGBA{0xff, 0, 0, 0xff})
	pointerWait.Fill(color.RGBA{0, 0xff, 0, 0xff})
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		panic("not enough argument")
	}
	imgebiten, img, err := ebitenutil.NewImageFromFile(args[0])
	if err != nil {
		panic(err)
	}
	x := img.Bounds().Dx()
	y := img.Bounds().Dy()
	fmt.Println("size img:", x, y)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("posgraph")
	ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	g := &Game{imgebiten: imgebiten}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	//// Post Processing
	fmt.Println("")
	// Graph creation
	fmt.Printf("G = nx.empty_graph(%d)  # Our graph has %d nodes\nplt.figure()\n", len(g.Points), len(g.Points))

	// Positions
	fmt.Print("pos = {")
	ptsStr := make([]string, 0)
	for i, p := range g.Points {
		ptsStr = append(ptsStr,
			fmt.Sprintf("%d: [%d, %d]", i, p.X, y-p.Y),
		)
	}
	fmt.Print(strings.Join(ptsStr, ", "))
	fmt.Println("}")

	// Labels
	fmt.Print("labels = {")
	ptsName := make([]string, 0)
	for i, p := range g.Points {
		ptsName = append(ptsName,
			fmt.Sprintf("%d: \"%s\"", i, p.Name),
		)
	}
	fmt.Print(strings.Join(ptsName, ", "))
	fmt.Println("}")
	fmt.Println("nx.draw(G,pos)\nnx.draw_networkx_labels(G,pos, labels, font_size=15)\nplt.show()")

}
