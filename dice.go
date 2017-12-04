package main

import (
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/fatih/color"
)

var diceRegex = regexp.MustCompile(`([1-9]+)d{1}([0-9]+)([+|-]{1}[0-9]+)*$`)

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

func areDiceFormatted(dice string) bool {
	matched := diceRegex.MatchString(dice)
	return matched
}

func rollDice(num int, sides int) int {
	total := 0
	brown := color.New(color.FgYellow)
	mwrite := color.New(color.FgMagenta)
	magenta := color.New(color.FgHiMagenta).SprintFunc()

	if num > 1 {
		brown.Printf("\nRolling %d d%d's\n\n", num, sides)
	} else {
		brown.Printf("\nRolling a %dd%d\n\n", num, sides)
	}

	for i := 0; i < num; i++ {
		time.Sleep(200 * time.Millisecond)
		result := random(1, sides)
		total = total + result
		mwrite.Printf("#%03d %v d%d%v %s\n", i+1, color.HiBlackString("-"), sides, color.HiBlackString(":"), magenta(result))
	}

	return total
}

func getDice(dice string) {

	white := color.New(color.FgHiWhite)
	subStr := diceRegex.FindStringSubmatch(dice)

	numRolls, _ := strconv.Atoi(subStr[1])
	numSides, _ := strconv.Atoi(subStr[2])

	if subStr[3] != "" {
		rollTotal := rollDice(numRolls, numSides)
		parameters := make(map[string]interface{}, 8)
		mod, _ := strconv.Atoi(subStr[3])
		parameters["modifier"] = mod
		parameters["rollTotal"] = rollTotal

		expression, _ := govaluate.NewEvaluableExpression("(rollTotal + modifier)")
		result, _ := expression.Evaluate(parameters)
		white.Printf("\nTotal: %d %s is %.0f\n", rollTotal, subStr[3], result)
	} else {
		rollTotal := rollDice(numRolls, numSides)
		white.Printf("\nTotal: %d\n", rollTotal)
	}

}

func main() {

	args := os.Args[1:]

	for i := range args {
		matched := areDiceFormatted(args[i])
		if matched {
			getDice(args[i])
		}
	}

}
