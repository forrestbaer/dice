package main

import (
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/fatih/color"
)

var diceRegex = regexp.MustCompile(`([0-9]+)d{1}([0-9]+)(([+|-]){1}[0-9]+)*$`)
var modRegex = regexp.MustCompile(`([+-])([0-9]+)`)
var multiple = false

func random(min, max int) int {
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

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < num; i++ {
		result := random(1, sides)
		total = total + result
		mwrite.Printf("#%03d %v d%d%v %s\n", i+1, color.HiBlackString("-"), sides, color.HiBlackString(":"), magenta(result))
	}

	return total
}

func getDice(dice string) {

	white := color.New(color.FgHiWhite)
	subStr := diceRegex.FindStringSubmatch(dice)
	modStr := modRegex.FindStringSubmatch(subStr[3])

	numRolls, _ := strconv.Atoi(subStr[1])
	numSides, _ := strconv.Atoi(subStr[2])

	if subStr[3] != "" {
		rollPlusMod := 0
		rollTotal := rollDice(numRolls, numSides)
		modAmmount, _ := strconv.Atoi(modStr[2])
		modifier := modStr[1]

		switch modifier {
		case "-":
			rollPlusMod = rollTotal - modAmmount
		case "+":
			rollPlusMod = rollTotal + modAmmount
		}

		white.Printf("\nTotal: %d %s is %d\n\n", rollTotal, subStr[3], rollPlusMod)
	} else {
		rollTotal := rollDice(numRolls, numSides)
		white.Printf("\nTotal: %d\n\n", rollTotal)
	}

}

func main() {

	args := os.Args[1:]
	black := color.New(color.FgHiBlack)

	if len(args) > 1 {
		multiple = true
	}

	for i := range args {
		matched := areDiceFormatted(args[i])
		if matched {
			if multiple && i > 0 {
				black.Println("--------------------")
			}
			getDice(args[i])
		}
	}

}
