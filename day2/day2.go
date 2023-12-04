package day1

import (
	"aoc2023/helpers"
	"regexp"
	"strconv"
	"strings"
)

func splitGame(game string, separator string) (string, string) {
	s := strings.Split(game, separator)
	return s[0], s[1]
}

func One() int {

	var maxCubesByColorForAllRounds = map[string]int{
		"red":   12,
		"blue":  14,
		"green": 13,
	}

	s := helpers.ParseLines("day2/input2.txt")

	res := 0

	re := regexp.MustCompile(`(\d+) (red|blue|green)`)

	for _, v := range s {
		game, detail := splitGame(v, ":")
		gameNumber, _ := strconv.Atoi(strings.Fields(game)[1])
		matches := re.FindAllStringSubmatch(detail, -1)

		var maxCubesByColorForRound = map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, match := range matches {
			cubeNumber, _ := strconv.Atoi(match[1])
			prevMaxCubes := maxCubesByColorForRound[match[2]]
			if cubeNumber > prevMaxCubes {
				maxCubesByColorForRound[match[2]] = cubeNumber
			}
		}

		inRange := true
		for color, maxCubesForRound := range maxCubesByColorForRound {
			if maxCubesForRound > maxCubesByColorForAllRounds[color] {
				inRange = false
				break
			}
		}

		if inRange {
			res += gameNumber
		}

	}

	return res

}

// Game (?P<gameNumber>\d+)
