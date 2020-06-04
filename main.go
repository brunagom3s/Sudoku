package main

import (
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
  "sudoku/pkg/sudoku"
)

func basic() string {
  return "Cacahuete"
}

func generateSudoku() [9][9]int {
  board:=sudoku.SudokuGenerator()
	return board
}

func main() {
  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "Sudoku",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(generateSudoku)
  app.Run()
}
