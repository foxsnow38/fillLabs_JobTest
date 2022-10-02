package q2

import "math"

func Que2(input int) []int {

	stackarray := make([]int, 0) // I create array for none cap
	_, arrayValue := smartDivide(input, stackarray)

	return reverseInts(arrayValue)

}

func smartDivide(input int, stackArray []int) (int, []int) {
	if input != 1 {

		stackArray = append(stackArray, input)
		return smartDivide(int(math.Floor(float64(input/2))), stackArray)
	}
	if input == 1 {
		return input, stackArray
	}

	//thats pretty useless but go want
	return input, stackArray

}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
	// I find this way for reverse
	//https://stackoverflow.com/a/29440669/18456713

}
