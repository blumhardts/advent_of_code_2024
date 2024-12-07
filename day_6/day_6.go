package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func (c *Coordinates) Left() Coordinates {
	return Coordinates{c.X - 1, c.Y}
}

func (c *Coordinates) Up() Coordinates {
	return Coordinates{c.X, c.Y + 1}
}

func (c *Coordinates) Right() Coordinates {
	return Coordinates{c.X + 1, c.Y}
}

func (c *Coordinates) Down() Coordinates {
	return Coordinates{c.X, c.Y - 1}
}

type Area struct {
	Map     [][]string
	Visited [][]bool
	Facing  string
	Current Coordinates
}

func (a *Area) At(c Coordinates) (string, bool) {
	if c.X < 0 ||
		c.Y < 0 ||
		c.X >= len(a.Map[0]) ||
		c.Y >= len(a.Map) {
		return "", false
	}

	return a.Map[c.X][c.Y], true
}

func (a *Area) Move() Coordinates {
	switch a.Facing {
	case "<":
		return a.Current.Left()
	case "^":
		return a.Current.Up()
	case ">":
		return a.Current.Right()
	case "v":
		return a.Current.Down()
	default:
		panic(errors.New("facing invalid direction"))
	}
}

func (a *Area) Turn() string {
	switch a.Facing {
	case "<":
		return "^"
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	default:
		panic(errors.New("facing invalid direction"))
	}
}

func (a *Area) FacingObstacle() bool {
	switch a.Facing {
	case "<":
		if v, ok := a.At(a.Current.Left()); ok {
			return v == "#"
		}

		return false
	case "^":
		if v, ok := a.At(a.Current.Up()); ok {
			return v == "#"
		}

		return false
	case ">":
		if v, ok := a.At(a.Current.Right()); ok {
			return v == "#"
		}

		return false
	case "v":
		if v, ok := a.At(a.Current.Down()); ok {
			return v == "#"
		}

		return false
	default:
		panic(errors.New("Facing invalid direction"))
	}
}

func (a *Area) Foo() {

}

func startingRow(row []string) (int, string, bool) {
	for i, v := range row {
		if v == "<" || v == "^" || v == ">" || v == "v" {
			return i, v, true
		}
	}

	return 0, "", false
}

func load() (a Area) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")

		a.Map = append(a.Map, row)
		a.Visited = append(a.Visited, make([]bool, len(row)))

		if j, dir, ok := startingRow(row); ok {
			a.Facing = dir
			a.Current = Coordinates{i, j}
			a.Visited[i][j] = true
		}

		i++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}

func main() {
	area := load()

	fmt.Printf("Start: %d\n", area.Current)
}
