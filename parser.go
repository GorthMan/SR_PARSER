package main

import (
	"fmt"
	"strconv"
)

//Parser Configuration
var grammar = [6][]string{
	{"E", "E", "+", "T"},
	{"E", "T"},
	{"T", "T", "*", "F"},
	{"T", "F"},
	{"F", "(", "E", ")"},
	{"F", "id"}}

var aTable = [12][6]string{
	{"S5", "", "", "S4", "", ""},     // 0
	{"", "S6", "", "", "", "accept"}, // 1
	{"", "R2", "S7", "", "R2", "R2"}, // 2
	{"", "R4", "R4", "", "R4", "R4"}, // 3
	{"S5", "", "", "S4", "", ""},     // 4
	{"", "R6", "R6", "", "R6", "R6"}, // 5
	{"S5", "", "", "S4", "", ""},     // 6
	{"S5", "", "", "S4", "", ""},     // 7
	{"", "S6", "", "", "S11", ""},    // 8
	{"", "R1", "S7", "", "R1", "R1"}, // 9
	{"", "R3", "R3", "", "R3", "R3"}, // 10
	{"", "R5", "R5", "", "R5", "R5"}, // 11

}

var gTable = [12][3]string{
	{"1", "2", "3"}, // 0
	{"", "", ""},    // 1
	{"", "", ""},    // 2
	{"", "", ""},    // 3
	{"8", "2", "3"}, // 4
	{"", "", ""},    // 5
	{"", "9", "3"},  // 6
	{"", "", "10"},  // 7
	{"", "", ""},    // 8
	{"", "", ""},    // 9
	{"", "", ""},    // 10
	{"", "", ""},    // 11
}

//Parser Control
var ps parseStack
var iq []string
var parseNotComplete bool
var choice string
var pt parseTree

func parse(parseIn []string) {
	//Make a new parse stack
	ps = newParseStack()
	ps.push(pstackItem{"", "0"})
	iq = parseIn
	parseNotComplete = true
	//Begin and manage parsing
	for parseNotComplete {
		parse1step()
	}

}

func parse1step() {
	var nextState string
	//Get state from parsestack and grammar from inputqueue
	curState := ps.top().stateSym
	nextSym := iq[0]
	fmt.Printf("Looking up %v %v in actiontable\n", curState, nextSym)
	choice, nextState = aLookup(curState, nextSym)
	println(choice)

	switch choice {
	case "accept":
		println("Parse complete!")
		parseNotComplete = false
	case "error":
		println("Error in parsing")
		parseNotComplete = false
	case "S": //Shift
		//If the next token is an ID, shift it onto parse tree stack
		if iq[0] == "id" {
			pt.parseTreeShift()
		}
		//Remove first item from input queue
		iq = iq[1:]
		//Shift item onto parse stack
		shift(nextSym, nextState)
		//Put item onto parse tree stack

	case "R": //Reduce
		LHS := reduce(nextSym, nextState)
		pt.parseTreeReduce(LHS)
	}
	fmt.Printf("Parse Table after parse: %v\n", ps.String())
}

func aLookup(state string, grammarSym string) (string, string) {

	state_pos, err := strconv.ParseInt(state, 10, 64)
	if err != nil {
		println("ERROR ENCOUNTERED WHEN CONVERTING STRING TO INT!")
		return "error", ""
	}
	sym_pos := symConvert(grammarSym)
	if sym_pos == -1 {
		println("Error: Unknown symbol")
		return "error", ""
	}
	action := aTable[state_pos][sym_pos]
	if action == "" {
		println("ERROR: action table returned an empty string")
		return "error", ""
	}
	//Process action into action and state
	if action == "accept" {
		return action, "0"
	} else {
		choice := action[0]
		newState := action[1]
		return string(choice), string(newState)
	}
}

func goLookup(state string, LHS string) string {
	state_pos, err := strconv.ParseInt(state, 10, 64)
	if err != nil {
		println("ERROR ENCOUNTERED WHEN CONVERTING STRING TO INT!")
		return ""
	}
	var gotoIndex int
	switch LHS {
	case "F":
		gotoIndex = 2
	case "T":
		gotoIndex = 1
	case "E":
		gotoIndex = 0
	}

	mygoto := gTable[state_pos][gotoIndex]
	if mygoto == "" {
		println("Blank goto entry lookup")
		return mygoto
	}
	return mygoto
}

func grammarLookup(state string) (string, int) {
	state_pos, err := strconv.ParseInt(state, 10, 64)
	if err != nil {
		println("ERROR ENCOUNTERED WHEN CONVERTING STRING TO INT!")
		return "", -1
	}
	return grammar[state_pos-1][0], len(grammar[state_pos-1]) - 1
}

func shift(symbol string, state string) {
	ps.push(pstackItem{symbol, state})
}

func reduce(symbol string, state string) string {
	LHS, r := grammarLookup(state)
	ps.popNum(r)
	state = ps.top().stateSym
	state = goLookup(state, LHS)
	ps.push(pstackItem{LHS, state})
	return LHS
}

//Helper function to convert symbol to its position in the aTable list
//-1 = unknown symbol
func symConvert(sym string) int {
	switch sym {
	case "id":
		return 0
	case "+":
		return 1
	case "*":
		return 2
	case "(":
		return 3
	case ")":
		return 4
	case "$":
		return 5
	default:
		return -1
	}

}
