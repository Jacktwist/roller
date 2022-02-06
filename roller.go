////////
// 		Basic Dice roller function. Create a Roller struct that takes a string in
// xdykz format where ie. 4d6k3 means 4 dice with 6 sides and keep 3. Keep is
// optional. It returns a Roll struct which stores the command string given,
// an array of dice rolls and a sum total of the roll. The Roller object itself
// keeps a history of Roll objects

package roller

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"math/rand"
	"time"
)

func init () {
	rand.Seed(time.Now().UnixNano())
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

type Roll struct {
	Command string
	Rolls   []int
	Result  int
}
type Roller struct {
	RollHistory []Roll
}

func (rl *Roller) Roll(s string) {

	var r Roll
	rx := regexp.MustCompile("^[1-9]+[0-9]*d[1-9]+[0-9]*(k[1-9]+[0-9]*)?")
	byteRoll := []byte(s)

	var num, sides, keep []byte
	mode := "num"
	var rolls []int

	// Make sure the string matches the regexp then strip out all the numbers into separate string variables

	if rx.MatchString(string(byteRoll)) {
		for i := range byteRoll {

			if mode == "keep" {
				keep = append(keep, byteRoll[i])
			}
			if mode == "sides" {
				if byteRoll[i] == 'k' {
					mode = "keep"
				} else {
					sides = append(sides, byteRoll[i])
				}
			}
			if mode == "num" {
				if byteRoll[i] == 'd' {
					mode = "sides"
				} else {
					num = append(num, byteRoll[i])
				}
			}
		}
	} else {
		fmt.Println("Regex pattern not matched.")
	}

	// Make our string numbers variables
	numN, _ := strconv.Atoi(string(num))
	sidesN, _ := strconv.Atoi(string(sides))
	keepN, _ := strconv.Atoi(string(keep))

	for i := 0; i < numN; i++ {
		rolls = append(rolls, rand.Intn(sidesN)+1)
	}
	r.Rolls = rolls

	// sort the results and trim the lowest results if needed
	sort.Ints(rolls)
	// Make sure we're not trying to keep more than we're rolling!
	if keepN > 0 && len(rolls)-keepN >= 0 {
		rolls = rolls[len(rolls)-keepN : len(rolls)]
	}
	r.Command = s
	r.Result = sum(rolls)
	rl.RollHistory = append(rl.RollHistory, r)
}

func (rl *Roller) String() string {
	var astring string
	for i := 0; i < len(rl.RollHistory); i++ {
		astring += rl.RollHistory[i].String()
	}
	return astring
}

func (r *Roll) String() string {
	var astring string
	for i := 0; i < len(r.Rolls); i++ {
		astring += fmt.Sprintf("%v ", r.Rolls[i])
	}
	return fmt.Sprintf("%s rolled: %swith a sum of %v\n", r.Command, astring, r.Result)
}
