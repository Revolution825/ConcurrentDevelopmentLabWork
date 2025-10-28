//Copyright (C) 2025 Diarmuid O'Neill

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// Modified By: Diarmuid O'Neill (C00282898@setu.ie)

package main

import (
	"image/color"
	"log"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

const scale int = 5
const width = 1600
const height = 900

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var grid [width][height]uint8 = [width][height]uint8{}
var buffer [width][height]uint8 = [width][height]uint8{}
var count int = 0

func update() error {
	var wg sync.WaitGroup
	jobs := make(chan int, 1000)

	for worker := 1; worker <= 1; worker++ { // Creates 4 workers
		wg.Add(1)
		go concurrentUpdate(jobs, &wg)
	}
	for x := 1; x < width-1; x++ {
		jobs <- x
	}
	close(jobs)

	wg.Wait()
	temp := buffer
	buffer = grid
	grid = temp
	return nil
}

func concurrentUpdate(jobs <-chan int, wg *sync.WaitGroup) error {
	defer wg.Done()
	for job := range jobs {
		x := job
		for y := 1; y < height-1; y++ { // Sends jobs to workers
			buffer[x][y] = 0
			n := grid[x-1][y-1] + grid[x-1][y+0] + grid[x-1][y+1] + grid[x+0][y-1] + grid[x+0][y+1] + grid[x+1][y-1] + grid[x+1][y+0] + grid[x+1][y+1]

			if grid[x][y] == 0 && n == 3 {
				buffer[x][y] = 1
			} else if n < 2 || n > 3 {
				buffer[x][y] = 0
			} else {
				buffer[x][y] = grid[x][y]
			}
		}
	}
	return nil
}

func display(window *ebiten.Image) {
	window.Fill(blue)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					if grid[x][y] == 1 {
						window.Set(x*scale+i, y*scale+j, yellow)
					}
				}
			}
		}
	}
}

func frame(window *ebiten.Image) error {
	count++
	var err error = nil
	if count == 1 {
		err = update()
		count = 0
	}
	if !ebiten.IsDrawingSkipped() {
		display(window)
	}

	return err
}

func main() {
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}

	if err := ebiten.Run(frame, width, height, 1, "Game of Life"); err != nil {
		log.Fatal(err)
	}
}
