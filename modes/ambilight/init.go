package ambilight

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/Rostislaved/screenshot"
	"log"
	"my-projects/ambilight/config"
)

const (
	numberOfColorsInRGB = 3
	numberOfHeaderByes  = 6
)

type ambilight struct {
	buffer       []byte
	nLeds        int
	tiles        tiles
	update       int
	screenshoter *screenshot.Screenshoter
}

type tiles []tile

type tile struct {
	x      int
	y      int
	width  int
	height int
}

type configuration struct {
	nTilesHorizontal         int
	horizontalHeightFraction float32

	nTilesVertical        int
	verticalWidthFraction float32
}

func New(cfg config.Config) *ambilight {
	nLeds := 2 * (cfg.LEDs.NumberOfHorizontal + cfg.LEDs.NumberOfVertical)

	size := numberOfHeaderByes + nLeds*numberOfColorsInRGB

	buffer := make([]byte, size)
	buffer[0] = 'A' // Magic word
	buffer[1] = 'd'
	buffer[2] = 'a'
	buffer[3] = byte((nLeds - 1) >> 8)             // LED count high byte
	buffer[4] = byte((nLeds - 1) & 0xff)           // LED count low byte
	buffer[5] = byte(buffer[3] ^ buffer[4] ^ 0x55) // Checksum

	//conf := configuration{
	//	nTilesHorizontal:         cfg.LEDs.NumberOfHorizontal,
	//	horizontalHeightFraction: cfg.Ambilight.HorizontalHeightFraction,
	//	nTilesVertical:           cfg.LEDs.NumberOfVertical,
	//	verticalWidthFraction:    cfg.Ambilight.VerticalWidthFraction,
	//}

	tiles, err := calcTiles(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//tiles.decrease()

	screenshoter := screenshot.New()

	return &ambilight{
		buffer:       buffer,
		nLeds:        nLeds,
		tiles:        tiles,
		screenshoter: screenshoter,
	}
}

func (tiles tiles) decrease() {
	// Decrease area of all tiles scaling it to its center
	scalingFactor := 2

	for i := range tiles {
		newWidth := tiles[i].width / scalingFactor
		newHeight := tiles[i].height / scalingFactor
		centerX := tiles[i].x + tiles[i].width/2
		centerY := tiles[i].y + tiles[i].height/2

		tiles[i].x = centerX - newWidth/2
		tiles[i].y = centerY - newHeight/2
		tiles[i].width = newWidth
		tiles[i].height = newHeight
	}
}

func calcTiles(cfg config.Config) (tiles tiles, err error) {
	nTilesHorizontal := cfg.LEDs.NumberOfHorizontal
	horizontalHeightFraction := cfg.Ambilight.HorizontalHeightFraction
	nTilesVertical := cfg.LEDs.NumberOfVertical
	verticalWidthFraction := cfg.Ambilight.VerticalWidthFraction
	verticalOffsetFraction := cfg.Ambilight.VerticalOffsetFraction
	horizontalOffsetFraction := cfg.Ambilight.HorizontalOffsetFraction

	screenWidthInPixels, screenHeightInPixels, err := getScreenSizeInPixels()
	if err != nil {
		return
	}

	horizontalTileWidthInPixels := screenWidthInPixels / nTilesHorizontal
	horizontalTileHeightInPixels := int(float32(screenHeightInPixels) * horizontalHeightFraction)

	verticalTileHeightInPixels := screenHeightInPixels / nTilesVertical
	verticalTileWidthInPixels := int(float32(screenWidthInPixels) * verticalWidthFraction)

	// Левая вертикаль
	for i := 0; i < nTilesVertical; i++ {
		tile := tile{
			x:      0 + int(float32(screenWidthInPixels)*verticalOffsetFraction),
			y:      screenHeightInPixels - (i+1)*verticalTileHeightInPixels,
			width:  verticalTileWidthInPixels,
			height: verticalTileHeightInPixels,
		}

		tiles = append(tiles, tile)
	}

	// Верхняя горизонталь
	for i := 0; i < nTilesHorizontal; i++ {
		tile := tile{
			x:      i * horizontalTileWidthInPixels,
			y:      0 + int(float32(screenHeightInPixels)*horizontalOffsetFraction),
			width:  horizontalTileWidthInPixels,
			height: horizontalTileHeightInPixels,
		}

		tiles = append(tiles, tile)
	}

	// Правая вертикаль
	for i := 0; i < nTilesVertical; i++ {
		tile := tile{
			x:      (screenWidthInPixels - verticalTileWidthInPixels) - int(float32(screenWidthInPixels)*verticalOffsetFraction),
			y:      i * verticalTileHeightInPixels,
			width:  verticalTileWidthInPixels,
			height: verticalTileHeightInPixels,
		}

		tiles = append(tiles, tile)
	}

	// Нижняя горизонталь
	for i := 0; i < nTilesHorizontal; i++ {
		tile := tile{
			x:      screenWidthInPixels - (i+1)*horizontalTileWidthInPixels,
			y:      (screenHeightInPixels - horizontalTileHeightInPixels) - int(float32(screenHeightInPixels)*horizontalOffsetFraction),
			width:  horizontalTileWidthInPixels,
			height: horizontalTileHeightInPixels,
		}

		tiles = append(tiles, tile)
	}

	return
}

func getScreenSizeInPixels() (width, height int, err error) {
	c, err := xgb.NewConn()
	if err != nil {
		return
	}
	defer c.Close()
	screen := xproto.Setup(c).DefaultScreen(c)

	width = int(screen.WidthInPixels)
	height = int(screen.HeightInPixels)

	return
}
