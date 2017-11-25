package main

// Apache License
// Version 2.0, January 2004
// http://www.apache.org/licenses/

// TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

// 1. Definitions.

// "License" shall mean the terms and conditions for use, reproduction,
// and distribution as defined by Sections 1 through 9 of this document.

// "Licensor" shall mean the copyright owner or entity authorized by
// the copyright owner that is granting the License.

// "Legal Entity" shall mean the union of the acting entity and all
// other entities that control, are controlled by, or are under common
// control with that entity. For the purposes of this definition,
// "control" means (i) the power, direct or indirect, to cause the
// direction or management of such entity, whether by contract or
// otherwise, or (ii) ownership of fifty percent (50%) or more of the
// outstanding shares, or (iii) beneficial ownership of such entity.

// "You" (or "Your") shall mean an individual or Legal Entity
// exercising permissions granted by this License.

// "Source" form shall mean the preferred form for making modifications,
// including but not limited to software source code, documentation
// source, and configuration files.

// "Object" form shall mean any form resulting from mechanical
// transformation or translation of a Source form, including but
// not limited to compiled object code, generated documentation,
// and conversions to other media types.

// "Work" shall mean the work of authorship, whether in Source or
// Object form, made available under the License, as indicated by a
// copyright notice that is included in or attached to the work
// (an example is provided in the Appendix below).

// "Derivative Works" shall mean any work, whether in Source or Object
// form, that is based on (or derived from) the Work and for which the
// editorial revisions, annotations, elaborations, or other modifications
// represent, as a whole, an original work of authorship. For the purposes
// of this License, Derivative Works shall not include works that remain
// separable from, or merely link (or bind by name) to the interfaces of,
// the Work and Derivative Works thereof.

// "Contribution" shall mean any work of authorship, including
// the original version of the Work and any modifications or additions
// to that Work or Derivative Works thereof, that is intentionally
// submitted to Licensor for inclusion in the Work by the copyright owner
// or by an individual or Legal Entity authorized to submit on behalf of
// the copyright owner. For the purposes of this definition, "submitted"
// means any form of electronic, verbal, or written communication sent
// to the Licensor or its representatives, including but not limited to
// communication on electronic mailing lists, source code control systems,
// and issue tracking systems that are managed by, or on behalf of, the
// Licensor for the purpose of discussing and improving the Work, but
// excluding communication that is conspicuously marked or otherwise
// designated in writing by the copyright owner as "Not a Contribution."

// "Contributor" shall mean Licensor and any individual or Legal Entity
// on behalf of whom a Contribution has been received by Licensor and
// subsequently incorporated within the Work.

// 2. Grant of Copyright License. Subject to the terms and conditions of
// this License, each Contributor hereby grants to You a perpetual,
// worldwide, non-exclusive, no-charge, royalty-free, irrevocable
// copyright license to reproduce, prepare Derivative Works of,
// publicly display, publicly perform, sublicense, and distribute the
// Work and such Derivative Works in Source or Object form.

// 3. Grant of Patent License. Subject to the terms and conditions of
// this License, each Contributor hereby grants to You a perpetual,
// worldwide, non-exclusive, no-charge, royalty-free, irrevocable
// (except as stated in this section) patent license to make, have made,
// use, offer to sell, sell, import, and otherwise transfer the Work,
// where such license applies only to those patent claims licensable
// by such Contributor that are necessarily infringed by their
// Contribution(s) alone or by combination of their Contribution(s)
// with the Work to which such Contribution(s) was submitted. If You
// institute patent litigation against any entity (including a
// cross-claim or counterclaim in a lawsuit) alleging that the Work
// or a Contribution incorporated within the Work constitutes direct
// or contributory patent infringement, then any patent licenses
// granted to You under this License for that Work shall terminate
// as of the date such litigation is filed.

// 4. Redistribution. You may reproduce and distribute copies of the
// Work or Derivative Works thereof in any medium, with or without
// modifications, and in Source or Object form, provided that You
// meet the following conditions:

// (a) You must give any other recipients of the Work or
// Derivative Works a copy of this License; and

// (b) You must cause any modified files to carry prominent notices
// stating that You changed the files; and

// (c) You must retain, in the Source form of any Derivative Works
// that You distribute, all copyright, patent, trademark, and
// attribution notices from the Source form of the Work,
// excluding those notices that do not pertain to any part of
// the Derivative Works; and

// (d) If the Work includes a "NOTICE" text file as part of its
// distribution, then any Derivative Works that You distribute must
// include a readable copy of the attribution notices contained
// within such NOTICE file, excluding those notices that do not
// pertain to any part of the Derivative Works, in at least one
// of the following places: within a NOTICE text file distributed
// as part of the Derivative Works; within the Source form or
// documentation, if provided along with the Derivative Works; or,
// within a display generated by the Derivative Works, if and
// wherever such third-party notices normally appear. The contents
// of the NOTICE file are for informational purposes only and
// do not modify the License. You may add Your own attribution
// notices within Derivative Works that You distribute, alongside
// or as an addendum to the NOTICE text from the Work, provided
// that such additional attribution notices cannot be construed
// as modifying the License.

// You may add Your own copyright statement to Your modifications and
// may provide additional or different license terms and conditions
// for use, reproduction, or distribution of Your modifications, or
// for any such Derivative Works as a whole, provided Your use,
// reproduction, and distribution of the Work otherwise complies with
// the conditions stated in this License.

// 5. Submission of Contributions. Unless You explicitly state otherwise,
// any Contribution intentionally submitted for inclusion in the Work
// by You to the Licensor shall be under the terms and conditions of
// this License, without any additional terms or conditions.
// Notwithstanding the above, nothing herein shall supersede or modify
// the terms of any separate license agreement you may have executed
// with Licensor regarding such Contributions.

// 6. Trademarks. This License does not grant permission to use the trade
// names, trademarks, service marks, or product names of the Licensor,
// except as required for reasonable and customary use in describing the
// origin of the Work and reproducing the content of the NOTICE file.

// 7. Disclaimer of Warranty. Unless required by applicable law or
// agreed to in writing, Licensor provides the Work (and each
// Contributor provides its Contributions) on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied, including, without limitation, any warranties or conditions
// of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
// PARTICULAR PURPOSE. You are solely responsible for determining the
// appropriateness of using or redistributing the Work and assume any
// risks associated with Your exercise of permissions under this License.

// 8. Limitation of Liability. In no event and under no legal theory,
// whether in tort (including negligence), contract, or otherwise,
// unless required by applicable law (such as deliberate and grossly
// negligent acts) or agreed to in writing, shall any Contributor be
// liable to You for damages, including any direct, indirect, special,
// incidental, or consequential damages of any character arising as a
// result of this License or out of the use or inability to use the
// Work (including but not limited to damages for loss of goodwill,
// work stoppage, computer failure or malfunction, or any and all
// other commercial damages or losses), even if such Contributor
// has been advised of the possibility of such damages.

// 9. Accepting Warranty or Additional Liability. While redistributing
// the Work or Derivative Works thereof, You may choose to offer,
// and charge a fee for, acceptance of support, warranty, indemnity,
// or other liability obligations and/or rights consistent with this
// License. However, in accepting such obligations, You may act only
// on Your own behalf and on Your sole responsibility, not on behalf
// of any other Contributor, and only if You agree to indemnify,
// defend, and hold each Contributor harmless for any liability
// incurred by, or claims asserted against, such Contributor by reason
// of your accepting any such warranty or additional liability.

// END OF TERMS AND CONDITIONS

// APPENDIX: How to apply the Apache License to your work.

// To apply the Apache License to your work, attach the following
// boilerplate notice, with the fields enclosed by brackets "[]"
// replaced with your own identifying information. (Don't include
// the brackets!)  The text should be enclosed in the appropriate
// comment syntax for the file format. We also recommend that a
// file or class name and description of purpose be included on the
// same "printed page" as the copyright notice for easier
// identification within third-party archives.

// Copyright [2017] [Kamalavelan]

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
