package main

import (
	"log"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// Categories ...
// Languages ...
// Infrastructures ...
const (
	Categories      = "Categories"
	Languages       = "Languages"
	Infrastructures = "Infrastructures"
)

// Item ...
type Item struct {
	ID     int
	Name   string
	Cursor bool
}

// Screen ...
type Screen struct {
	Items            []Item
	RowCount         int
	CurrentRowNumber int
	ItemName         string
}

func (screen *Screen) moveScreenToTheRight(items []Item, width int, cell []termbox.Cell) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	createQueryRow()
	cursor := false
	for i, item := range items {
		for ii, r := range item.Name {
			termbox.SetCell(ii, i+1, r, termbox.ColorWhite, termbox.ColorBlack)
		}
		if item.Cursor {
			cursor = true
			screen.CurrentRowNumber = i + 1
			for iii := 0; iii < width; iii++ {
				termbox.SetCell(iii, i+1, cell[(i+1)*width+iii].Ch, termbox.ColorWhite, termbox.ColorWhite)
			}
		}
	}
	if !cursor {
		screen.CurrentRowNumber = 0
	}
	screen.RowCount = len(items)
	screen.Items = items
}

func (screen *Screen) moveScreenToTheLeft(items []Item, width int, cell []termbox.Cell) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	createQueryRow()
	for i, item := range items {
		for ii, r := range item.Name {
			termbox.SetCell(ii, i+1, r, termbox.ColorWhite, termbox.ColorBlack)
		}
		if item.Cursor {
			screen.CurrentRowNumber = i + 1
			for iii := 0; iii < width; iii++ {
				termbox.SetCell(iii, i+1, cell[(i+1)*width+iii].Ch, termbox.ColorWhite, termbox.ColorWhite)
			}
		}
	}
	screen.RowCount = len(items)
	screen.Items = items
}

func (screen *Screen) changeCursor() {
	for i := range screen.Items {
		screen.Items[i].Cursor = false
	}
	screen.Items[screen.CurrentRowNumber-1].Cursor = true
}

func createQueryRow() {
	for i, r := range "QUERY>" {
		termbox.SetCell(i, 0, r, termbox.ColorWhite, termbox.ColorBlack)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	screen := new(Screen)

	categories := []Item{
		{1, Languages, false},
		{2, Infrastructures, false},
	}

	languages := []Item{
		{1, "java", false},
		{2, "go", false},
		{3, "ruby", false},
		{4, "python", false},
		{5, "php", false},
	}

	infrastructures := []Item{
		{1, "chef", false},
		{2, "ansible", false},
		{3, "docker", false},
		{4, "terraform", false},
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	createQueryRow()
	for i, category := range categories {
		for ii, r := range category.Name {
			termbox.SetCell(ii, i+1, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}

	screen.Items = categories
	screen.RowCount = len(categories)
	screen.ItemName = Categories
	screen.CurrentRowNumber = 0

	width, _ := termbox.Size()
	cell := termbox.CellBuffer()

	termbox.Flush()

loop:
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}

			// Down
			if ev.Key == termbox.KeyArrowDown {
				if screen.CurrentRowNumber >= screen.RowCount {
					break
				}
				screen.CurrentRowNumber++
				for i := 0; i < width; i++ {
					if screen.CurrentRowNumber > 1 {
						termbox.SetCell(i, screen.CurrentRowNumber-1, cell[(screen.CurrentRowNumber-1)*width+i].Ch, termbox.ColorWhite, termbox.ColorBlack)
					}
					termbox.SetCell(i, screen.CurrentRowNumber, cell[screen.CurrentRowNumber*width+i].Ch, termbox.ColorWhite, termbox.ColorWhite)
				}

				screen.changeCursor()

				// Up
			} else if ev.Key == termbox.KeyArrowUp {
				if 1 >= screen.CurrentRowNumber {
					break
				}
				screen.CurrentRowNumber--
				for i := 0; i < width; i++ {
					termbox.SetCell(i, screen.CurrentRowNumber, cell[screen.CurrentRowNumber*width+i].Ch, termbox.ColorWhite, termbox.ColorWhite)
					termbox.SetCell(i, screen.CurrentRowNumber+1, cell[(screen.CurrentRowNumber+1)*width+i].Ch, termbox.ColorWhite, termbox.ColorBlack)
				}

				screen.changeCursor()

				// Right
			} else if ev.Key == termbox.KeyArrowRight {
				if 1 > screen.CurrentRowNumber {
					break
				}

				switch screen.ItemName {
				case Categories:
					screen.Items[screen.CurrentRowNumber-1].Cursor = true

					runes := []rune{}
					for i := 0; i < width; i++ {
						runes = append(runes, cell[screen.CurrentRowNumber*width+i].Ch)
					}
					switch strings.TrimSpace(string(runes)) {
					case Languages:
						screen.moveScreenToTheRight(languages, width, cell)
						screen.ItemName = Languages
					case Infrastructures:
						screen.moveScreenToTheRight(infrastructures, width, cell)
						screen.ItemName = Infrastructures
					}
				}

				// Left
			} else if ev.Key == termbox.KeyArrowLeft {
				if screen.ItemName == Languages || screen.ItemName == Infrastructures {
					screen.moveScreenToTheLeft(categories, width, cell)
					screen.ItemName = Categories
				}

			} else {
				for i, r := range "QUERY>" {
					termbox.SetCell(i, 0, r, termbox.ColorWhite, termbox.ColorBlack)
				}
				termbox.SetCell(7, 0, ev.Ch, termbox.ColorDefault, termbox.ColorDefault)
			}

			termbox.Flush()
		}
	}
}
