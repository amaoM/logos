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

type Item struct{}

type Screen struct {
	Categories       []Category
	Languages        []Language
	Infrastructures  []Infrastructure
	RowCount         int
	CurrentRowNumber int
	ItemName         string
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

	screen.RowCount = len(screen.Categories)
	screen.ItemName = "Categories"
	screen.CurrentRowNumber = -1

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
			if ev.Key == 65516 {
				if screen.CurrentRowNumber >= screen.RowCount-1 {
					break
				}
				screen.CurrentRowNumber++
				for i := 0; i < width; i++ {
					if screen.CurrentRowNumber > 0 {
						termbox.SetCell(i, screen.CurrentRowNumber-1, cell[(screen.CurrentRowNumber-1)*width+i].Ch, termbox.ColorWhite, termbox.ColorBlack)
					}
					termbox.SetCell(i, screen.CurrentRowNumber, cell[screen.CurrentRowNumber*width+i].Ch, termbox.ColorWhite, termbox.ColorWhite)
				}

				switch screen.ItemName {
				case "Categories":
					for i := range screen.Categories {
						screen.Categories[i].Cursor = false
					}
					screen.Categories[screen.CurrentRowNumber].Cursor = true
				case "Languages":
					for i := range screen.Languages {
						screen.Languages[i].Cursor = false
					}
					screen.Languages[screen.CurrentRowNumber].Cursor = true
				case "Infrastructures":
					for i := range screen.Infrastructures {
						screen.Infrastructures[i].Cursor = false
					}
					screen.Infrastructures[screen.CurrentRowNumber].Cursor = true
				}
			}

			// Up
			if ev.Key == 65517 {
				if 0 >= screen.CurrentRowNumber {
					break
				}
				screen.CurrentRowNumber--
				for i := 0; i < width; i++ {
					termbox.SetCell(i, screen.CurrentRowNumber, cell[screen.CurrentRowNumber*width+i].Ch, termbox.ColorWhite, termbox.ColorWhite)
					termbox.SetCell(i, screen.CurrentRowNumber+1, cell[(screen.CurrentRowNumber+1)*width+i].Ch, termbox.ColorWhite, termbox.ColorBlack)
				}

				switch screen.ItemName {
				case "Categories":
					for i := range screen.Categories {
						screen.Categories[i].Cursor = false
					}
					screen.Categories[screen.CurrentRowNumber].Cursor = true
				case "Languages":
					for i := range screen.Languages {
						screen.Languages[i].Cursor = false
					}
					screen.Languages[screen.CurrentRowNumber].Cursor = true
				case "Infrastructures":
					for i := range screen.Infrastructures {
						screen.Infrastructures[i].Cursor = false
					}
					screen.Infrastructures[screen.CurrentRowNumber].Cursor = true
				}
			}

			// Right
			if ev.Key == 65514 {
				if 0 > screen.CurrentRowNumber {
					break
				}

				switch screen.ItemName {
				case "Categories":
					screen.Categories[screen.CurrentRowNumber].Cursor = true

					runes := []rune{}
					for i := 0; i < width; i++ {
						runes = append(runes, cell[screen.CurrentRowNumber*width+i].Ch)
					}

					switch strings.TrimSpace(string(runes)) {
					case "Languages":
						termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
						cursor := false
						for i, language := range screen.Languages {
							for ii, r := range language.Name {
								termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
							}
							if language.Cursor {
								cursor = true
								screen.CurrentRowNumber = i
								for iii := 0; iii < width; iii++ {
									termbox.SetCell(iii, i, cell[i*width+iii].Ch, termbox.ColorWhite, termbox.ColorWhite)
								}
							}
						}
						if !cursor {
							screen.CurrentRowNumber = -1
						}
						screen.ItemName = "Languages"
						screen.RowCount = len(screen.Languages)
					case "Infrastructures":
						termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
						cursor := false
						for i, infrastructure := range screen.Infrastructures {
							for ii, r := range infrastructure.Name {
								termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
							}
							if infrastructure.Cursor {
								cursor = true
								screen.CurrentRowNumber = i
								for iii := 0; iii < width; iii++ {
									termbox.SetCell(iii, i, cell[i*width+iii].Ch, termbox.ColorWhite, termbox.ColorWhite)
								}
							}
						}
						if !cursor {
							screen.CurrentRowNumber = -1
						}
						screen.ItemName = "Infrastructures"
						screen.RowCount = len(screen.Infrastructures)
					}
				}
			}

			// Left
			if ev.Key == 65515 {
				if screen.ItemName == "Languages" || screen.ItemName == "Infrastructures" {
					termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
					for i, category := range screen.Categories {
						for ii, r := range category.Name {
							termbox.SetCell(ii, i, r, termbox.ColorWhite, termbox.ColorBlack)
						}
						if category.Cursor {
							screen.CurrentRowNumber = i
							for iii := 0; iii < width; iii++ {
								termbox.SetCell(iii, i, cell[i*width+iii].Ch, termbox.ColorWhite, termbox.ColorWhite)
							}
						}
					}
					screen.ItemName = "Categories"
					screen.RowCount = len(screen.Categories)
				}
			}

			termbox.Flush()
		}
	}
}
