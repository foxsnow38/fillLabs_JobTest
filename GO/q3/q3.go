package q3

import "fmt"

func Que3(input []string) string {

	stringMap := make(map[string]int, len(input))
	biggestKey := ""
	biggestValue := 0
	for i := 0; i < len(input); i++ {
		stringMap[input[i]]++
	}

	for k, v := range stringMap {
		if v > biggestValue {
			biggestKey = string(k)
			biggestValue = v
		}
	}

	fmt.Println(biggestKey)
	return biggestKey

}
