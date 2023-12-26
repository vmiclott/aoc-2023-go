package day02

type color int

const (
	red color = iota
	green
	blue
)

var colorMap = map[string]color{
	"red":   red,
	"green": green,
	"blue":  blue,
}

type round map[color]int

type game struct {
	id     int
	rounds []round
}

func (game game) isPossible(redLimit int, greenLimit int, blueLimit int) bool {
	for _, round := range game.rounds {
		if round[red] > redLimit || round[green] > greenLimit || round[blue] > blueLimit {
			return false
		}
	}

	return true
}

func (game game) power() int {
	redMin, greenMin, blueMin := 0, 0, 0
	for _, round := range game.rounds {
		if round[red] > redMin {
			redMin = round[red]
		}
		if round[green] > greenMin {
			greenMin = round[green]
		}
		if round[blue] > blueMin {
			blueMin = round[blue]
		}
	}
	return redMin * greenMin * blueMin
}
