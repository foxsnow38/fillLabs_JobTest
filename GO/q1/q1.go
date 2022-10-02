package q1

import "fmt"

func Que1(input []string) []string {

	arrayInput := make([]string, len(input)*2)
	copy(arrayInput, input)
	arrayInput = input
	arrayInput = append(arrayInput, string(input[0]+"++")) // for break apeend

	for i := 0; i < len(arrayInput); i++ {
		// main for

		for y := 0; y < len(arrayInput); y++ {
			if arrayInput[y+1] == string(input[0]+"++") {
				// never be same
				break
			}
			// arrange  for  from inside array
			var firstValueInt, secondValueInt int

			for z := 0; z < len(arrayInput[y]); z++ {
				// upper array how mant have a letter

				if string(arrayInput[y][z]) == string("a") {
					firstValueInt++

				}

			}

			for z := 0; z < len(arrayInput[y+1]); z++ {
				// lower array how many have a letter

				if string(arrayInput[y+1][z]) == string("a") {
					secondValueInt++

				}

			}

			if firstValueInt < secondValueInt {
				var firstValueString, secondValueString string
				firstValueString = arrayInput[y]
				secondValueString = arrayInput[y+1]

				arrayInput[y] = secondValueString
				arrayInput[y+1] = firstValueString
				// sort func
			}

		}
	}
	fmt.Println(arrayInput[:len(arrayInput)-1])
	return arrayInput[:len(arrayInput)-1]
}
