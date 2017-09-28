package main

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

type Category struct {
	ID   int
	Name string
}

type Language struct {
	ID   int
	Name string
}

type Infrastructure struct {
	ID   int
	Name string
}

type Screen struct {
	Categories      []Category
	Languages       []Language
	Infrastructures []Infrastructure
	Row             int
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	screen := new(Screen)

	screen.Categories = []Category{
		{1, "Language"},
		{2, "Tool"},
	}

	screen.Languages = []Language{
		{1, "java"},
		{2, "go"},
		{3, "ruby"},
		{4, "python"},
		{5, "php"},
	}

	screen.Infrastructures = []Infrastructure{
		{1, "chef"},
		{2, "ansible"},
		{3, "docker"},
		{4, "terraform"},
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i, category := range screen.Categories {
		for ii, r := range category.Name {
			termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}

	screen.Row = len(screen.Categories)
	termbox.Flush()

	width, _ := termbox.Size()
	cell := termbox.CellBuffer()
	row := -1

loop:
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break loop
			}

			// Down
			if ev.Key == 65516 {
				if row >= screen.Row-1 {
					break
				}
				row++
				for i := 0; i < width; i++ {
					if row > 0 {
						termbox.SetCell(i, row-1, cell[(row-1)*width+i].Ch, termbox.ColorWhite, termbox.ColorBlack)
					}
					termbox.SetCell(i, row, cell[row*width+i].Ch, termbox.ColorWhite, termbox.ColorWhite)
				}
			}

			// Up
			if ev.Key == 65517 {
				if 0 >= row {
					break
				}
				row--
				for i := 0; i < width; i++ {
					termbox.SetCell(i, row, cell[row*width+i].Ch, termbox.ColorWhite, termbox.ColorWhite)
					termbox.SetCell(i, row+1, cell[(row+1)*width+i].Ch, termbox.ColorWhite, termbox.ColorBlack)
				}
			}

			// Right
			if ev.Key == 65514 {

			}

			// Left
			if ev.Key == 65515 {

			}

			termbox.Flush()
		}
	}
}
