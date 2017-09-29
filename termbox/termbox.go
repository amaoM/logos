package main

import (
	"log"
	"strings"

	termbox "github.com/nsf/termbox-go"
)

type Category struct {
	ID     int
	Name   string
	Cursor bool
}

type Language struct {
	ID     int
	Name   string
	Cursor bool
}

type Infrastructure struct {
	ID     int
	Name   string
	Cursor bool
}

type Screen struct {
	Categories      []Category
	Languages       []Language
	Infrastructures []Infrastructure
	Row             int
	ItemName        string
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	screen := new(Screen)

	screen.Categories = []Category{
		{1, "Languages", false},
		{2, "Infrastructures", false},
	}

	screen.Languages = []Language{
		{1, "java", false},
		{2, "go", false},
		{3, "ruby", false},
		{4, "python", false},
		{5, "php", false},
	}

	screen.Infrastructures = []Infrastructure{
		{1, "chef", false},
		{2, "ansible", false},
		{3, "docker", false},
		{4, "terraform", false},
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i, category := range screen.Categories {
		for ii, r := range category.Name {
			termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
		}
	}

	screen.Row = len(screen.Categories)
	screen.ItemName = "Category"
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
			if ev.Key == 65514 && row >= 0 {
				switch screen.ItemName {
				case "Category":
					runes := []rune{}
					for i := 0; i < width; i++ {
						runes = append(runes, cell[row*width+i].Ch)
					}
					switch strings.TrimSpace(string(runes)) {
					case "Languages":
						for i, language := range screen.Languages {
							for ii, r := range language.Name {
								termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
							}
							for iii := len(language.Name); iii < width; iii++ {
								termbox.SetCell(iii, i, 32, termbox.ColorWhite, termbox.ColorBlack)
							}
						}
						screen.Row = len(screen.Languages)
						row = -1
					case "Infrastructures":
						for i, infrastructure := range screen.Infrastructures {
							for ii, r := range infrastructure.Name {
								termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
							}
							for iii := len(infrastructure.Name); iii < width; iii++ {
								termbox.SetCell(iii, i, 32, termbox.ColorWhite, termbox.ColorBlack)
							}
						}
						screen.Row = len(screen.Languages)
						row = -1
					}
				}
			}

			// Left
			if ev.Key == 65515 {

			}

			termbox.Flush()
		}
	}
}
