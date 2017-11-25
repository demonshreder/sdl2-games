package main

// This is free and unencumbered software released into the public domain.

// Anyone is free to copy, modify, publish, use, compile, sell, or
// distribute this software, either in source code form or as a compiled
// binary, for any purpose, commercial or non-commercial, and by any
// means.

// In jurisdictions that recognize copyright laws, the author or authors
// of this software dedicate any and all copyright interest in the
// software to the public domain. We make this dedication for the benefit
// of the public at large and to the detriment of our heirs and
// successors. We intend this dedication to be an overt act of
// relinquishment in perpetuity of all present and future rights to this
// software under copyright law.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

// For more information, please refer to <http://unlicense.org>
import (
	"math/rand"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/veandco/go-sdl2/sdl"
)

// Codes
// 0 - Edge tiles
// 1 - Play tiles
// 2 - Snake
// 3 - Berry

func main() {
	// Core stuff creation
	sdl.Init(sdl.INIT_VIDEO)
	defer sdl.Quit()
	ttf.Init()
	defer ttf.Quit()
	// img.Init(img.INIT_PNG)
	// defer img.Quit()

	mode := sdl.DisplayMode{}
	sdl.GetCurrentDisplayMode(0, &mode)

	screenHeight := int(mode.H)
	screenWidth := int(mode.W)

	window, _ := sdl.CreateWindow("Basic Snake", 0, 0, screenWidth, screenHeight, sdl.WINDOW_FULLSCREEN)
	defer window.Destroy()

	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	defer renderer.Destroy()

	// Global variables
	running := true

	mapStartX := int32(screenWidth / 6)
	mapStartY := int32(screenHeight / 8)

	tileSize := int32(30)
	tileDef := [][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	tileRect := sdl.Rect{X: mapStartX, Y: mapStartY, H: tileSize, W: tileSize}

	gameState := "start"

	snakeState := "none"
	snakeTiles := [][]int{{4, 4}}
	snakeTilesNext := []int{0, 0}
	snakeTilesPrev := []int{0, 0}

	berryTile := []int{4, 4}

	//Text stuff
	score := 0
	scoreNum := strconv.Itoa(score)
	scoreText := "Score : "
	scoreFont, _ := ttf.OpenFont("ttf/Audiowide-Regular.ttf", 50)
	scoreTextSurface, _ := scoreFont.RenderUTF8_Blended(scoreText, sdl.Color{R: 255, G: 255, B: 255, A: 127})
	scoreTextTexture, _ := renderer.CreateTextureFromSurface(scoreTextSurface)
	defer scoreTextSurface.Free()
	defer scoreTextTexture.Destroy()
	defer scoreFont.Close()
	scoreTextRect := sdl.Rect{X: mapStartX, Y: mapStartY - 60, W: 130, H: 50}

	scoreNumSurface, _ := scoreFont.RenderUTF8_Blended(scoreText, sdl.Color{R: 255, G: 255, B: 255, A: 127})
	scoreNumTexture, _ := renderer.CreateTextureFromSurface(scoreNumSurface)
	defer scoreNumSurface.Free()
	defer scoreNumTexture.Destroy()
	scoreNumRect := sdl.Rect{X: mapStartX + 130, Y: mapStartY - 60, W: 40, H: 50}

	demonFont, _ := ttf.OpenFont("ttf/CinzelDecorative-Black.ttf", 60)
	demon := "DEMONSHREDER's"
	demonSurface, _ := demonFont.RenderUTF8_Blended(demon, sdl.Color{R: 255, G: 255, B: 255, A: 127})
	demonTexture, _ := renderer.CreateTextureFromSurface(demonSurface)
	defer demonSurface.Free()
	defer demonFont.Close()
	defer demonTexture.Destroy()
	// demonRect := sdl.Rect{X: int32(screenWidth / 6), Y: int32(screenHeight / 5), W: int32(screenWidth - 400), H: int32(screenHeight / 3)}
	demonRect := sdl.Rect{X: int32(screenWidth / 6), Y: int32(screenHeight / 6), W: int32(screenWidth / 3), H: int32(screenHeight / 10)}

	basicSnakeFont, _ := ttf.OpenFont("ttf/FontdinerSwanky-Regular.ttf", 240)
	basicSnake := "Basic Snake"
	basicSnakeSurface, _ := basicSnakeFont.RenderUTF8_Blended(basicSnake, sdl.Color{R: 255, G: 255, B: 255, A: 127})
	basicSnakeTexture, _ := renderer.CreateTextureFromSurface(basicSnakeSurface)
	defer basicSnakeSurface.Free()
	defer basicSnakeFont.Close()
	defer basicSnakeTexture.Destroy()
	basicSnakeRect := sdl.Rect{X: int32(screenWidth / 5), Y: int32(screenHeight / 3), W: 750, H: 350}

	gameOver := "Game Over"
	gameOverSurface, _ := basicSnakeFont.RenderUTF8_Blended(gameOver, sdl.Color{R: 255, G: 255, B: 255, A: 127})
	gameOverTexture, _ := renderer.CreateTextureFromSurface(gameOverSurface)
	defer gameOverSurface.Free()
	defer gameOverTexture.Destroy()
	// Snake and Berry spawning
	randSeed := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSeed)

	anyKeyFont, _ := ttf.OpenFont("ttf/FontdinerSwanky-Regular.ttf", 50)
	anyKey := "Press any key to continue"
	anyKeySurface, _ := anyKeyFont.RenderUTF8_Blended(anyKey, sdl.Color{R: 255, G: 255, B: 255, A: 127})
	anyKeyTexture, _ := renderer.CreateTextureFromSurface(anyKeySurface)
	defer anyKeySurface.Free()
	defer anyKeyTexture.Destroy()
	anyKeyRect := sdl.Rect{X: int32(screenWidth / 4), Y: int32(screenHeight - 100), W: 650, H: 75}

	tileLengthOne := len(tileDef) - 2
	tileLengthTwo := len(tileDef[0]) - 2

	for running { // Main game loop

		// keyboard, mouse event polling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
					break
				}
				if t.Keysym.Sym == sdl.K_w && snakeState != "down" {
					snakeState = "up"

				} else if t.Keysym.Sym == sdl.K_d && snakeState != "left" {
					snakeState = "right"

				} else if t.Keysym.Sym == sdl.K_a && snakeState != "right" {
					snakeState = "left"

				} else if t.Keysym.Sym == sdl.K_s && snakeState != "up" {
					snakeState = "down"
				}
				if t.Type == sdl.KEYDOWN && (gameState == "start" || gameState == "over") {
					gameState = "running"

					score = 0
					for i := range snakeTiles {
						tileDef[snakeTiles[i][0]][snakeTiles[i][1]] = 1
					}
					snakeTiles = [][]int{{4, 4}}
					tileDef[berryTile[0]][berryTile[1]] = 1
					snakeState = "none"
					randomGood := false
					for !randomGood {
						berryTile[0] = randGen.Intn(tileLengthOne)
						berryTile[1] = randGen.Intn(tileLengthTwo)

						snakeTiles[0][0] = randGen.Intn(tileLengthOne)
						snakeTiles[0][1] = randGen.Intn(tileLengthTwo)

						if berryTile[0] > 1 && berryTile[1] > 1 && snakeTiles[0][0] > 1 && snakeTiles[0][1] > 1 {
							randomGood = true
						}
					}

					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 2
					tileDef[berryTile[0]][berryTile[1]] = 3

				}
			} // end switch
		} // end event poll loop
		if gameState == "start" {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			renderer.SetDrawColor(255, 255, 255, 255)

			renderer.Copy(demonTexture, nil, &demonRect)
			renderer.Copy(basicSnakeTexture, nil, &basicSnakeRect)
			renderer.Copy(anyKeyTexture, nil, &anyKeyRect)

			renderer.Present()
			// continue
		} else if gameState == "over" {
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()
			renderer.SetDrawColor(255, 255, 255, 255)
			renderer.Copy(gameOverTexture, nil, &basicSnakeRect)
			renderer.Copy(anyKeyTexture, nil, &anyKeyRect)

			renderer.Present()
			// continue
		} else if gameState == "running" {
			// snake chain update starts
			if snakeState != "none" {

				if snakeState == "up" {
					// for some reason it takes the value after change rather before it
					// snakeTilesNext = snakeTiles[0]
					snakeTilesNext[0] = snakeTiles[0][0]
					snakeTilesNext[1] = snakeTiles[0][1]
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 1
					snakeTiles[0][0]--
					if snakeTiles[0][0] < 1 {
						snakeTiles[0][0] = len(tileDef) - 2
					}
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 2
				} else if snakeState == "right" {
					snakeTilesNext[0] = snakeTiles[0][0]
					snakeTilesNext[1] = snakeTiles[0][1]
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 1
					snakeTiles[0][1]++
					if snakeTiles[0][1] > len(tileDef[0])-2 {
						snakeTiles[0][1] = 1
					}
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 2
				} else if snakeState == "left" {
					snakeTilesNext[0] = snakeTiles[0][0]
					snakeTilesNext[1] = snakeTiles[0][1]
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 1
					snakeTiles[0][1]--
					if snakeTiles[0][1] < 1 {
						snakeTiles[0][1] = len(tileDef[0]) - 2
					}
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 2

				} else if snakeState == "down" {
					snakeTilesNext[0] = snakeTiles[0][0]
					snakeTilesNext[1] = snakeTiles[0][1]
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 1
					snakeTiles[0][0]++
					if snakeTiles[0][0] > len(tileDef)-2 {
						snakeTiles[0][0] = 1
					}
					tileDef[snakeTiles[0][0]][snakeTiles[0][1]] = 2

				}
				// Game over condition
				for i := 1; i < len(snakeTiles); i++ {
					if snakeTiles[0][0] == snakeTiles[i][0] && snakeTiles[0][1] == snakeTiles[i][1] {
						gameState = "over"
					}
					snakeTilesPrev = snakeTiles[i]
					tileDef[snakeTilesPrev[0]][snakeTilesPrev[1]] = 1
					snakeTiles[i] = snakeTilesNext
					tileDef[snakeTilesNext[0]][snakeTilesNext[1]] = 2
					snakeTilesNext = snakeTilesPrev
				}
			} // snake chain update done
			if snakeTiles[0][0] == berryTile[0] && snakeTiles[0][1] == berryTile[1] && running {
				berryFound := false
				for !berryFound {
					berryTile[0] = randGen.Intn(tileLengthOne)
					berryTile[1] = randGen.Intn(tileLengthTwo)
					// To remove fringe elements
					if berryTile[0] > 1 && berryTile[1] > 1 {
						berryFound = true
					}
					// To avoid collision with snake tiles
					for _, x := range snakeTiles {
						if berryTile[0] == x[0] && berryTile[1] == x[1] {
							berryFound = false
						}
					}
					tempSnake := []int{0, 0}
					switch snakeState {
					case "up":
						tempSnake[0] = snakeTiles[0][0] + 1
						tempSnake[1] = snakeTiles[0][1]
						break
					case "down":
						tempSnake[0] = snakeTiles[0][0] - 1
						tempSnake[1] = snakeTiles[0][1]
						break
					case "left":
						tempSnake[0] = snakeTiles[0][0]
						tempSnake[1] = snakeTiles[0][1] + 1
						break
					case "right":
						tempSnake[0] = snakeTiles[0][0]
						tempSnake[1] = snakeTiles[0][1] - 1
						break
					}
					snakeTiles = append(snakeTiles, tempSnake)
					tileDef[snakeTiles[len(snakeTiles)-1][0]][snakeTiles[len(snakeTiles)-1][1]] = 2
				}

				tileDef[berryTile[0]][berryTile[1]] = 3
				score += 10
				scoreNum = strconv.Itoa(score)
				scoreNumRect.W = int32(len(scoreNum) * 20)

				// fmt.Println("xxxxxxxxxx")
				// fmt.Println(tileLengthOne, tileLengthTwo)
				// fmt.Println(berryTile)
			}
			sdl.Delay(125)
			// Updating graphics
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.Clear()

			// Score update
			renderer.Copy(scoreTextTexture, nil, &scoreTextRect)
			scoreNumSurface, _ = scoreFont.RenderUTF8_Blended(scoreNum, sdl.Color{R: 255, G: 255, B: 255, A: 127})
			scoreNumTexture, _ = renderer.CreateTextureFromSurface(scoreNumSurface)
			renderer.Copy(scoreNumTexture, nil, &scoreNumRect)
			// Map update
			for k, a := range tileDef {
				tileRect.Y = mapStartY + (tileSize * int32(k))
				for i := range a {
					if tileDef[k][i] == 0 {
						renderer.SetDrawColor(18, 80, 181, 255)
						tileRect.X = mapStartX + (tileSize * int32(i))
						renderer.FillRect(&tileRect)

					} else if tileDef[k][i] == 1 {
						renderer.SetDrawColor(193, 56, 56, 255)
						tileRect.X = mapStartX + (tileSize * int32(i))
						renderer.FillRect(&tileRect)
					} else if tileDef[k][i] == 2 {
						renderer.SetDrawColor(255, 132, 50, 255)
						tileRect.X = mapStartX + (tileSize * int32(i))
						renderer.FillRect(&tileRect)
					} else if tileDef[k][i] == 3 {
						renderer.SetDrawColor(38, 211, 145, 255)
						tileRect.X = mapStartX + (tileSize * int32(i))
						renderer.FillRect(&tileRect)
					}
				}
			} // Tile engine loop end

			//Present frame
			renderer.Present()
		} else {
		} // gameState == "running" condition end

	} // loop end
}
