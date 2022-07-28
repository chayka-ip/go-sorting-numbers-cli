package main

import (
	"bufio"
	"fmt"
	"os"
)

func runSort(algorithmName string, isAscending bool, numbers []int) {
	printLines()
	fmt.Printf("Sorting algorithm: %v \n", algorithmName)
	fmt.Printf("Is accending: %v \n", isAscending)
	fmt.Printf("Initial number array:  %v \n", numbers)
	SortByAlgorithm(algorithmName, numbers, isAscending)
	fmt.Printf("Result:  %v \n", numbers)
	printLines()
}

func lazyMode() {
	minElemValue := 0
	maxElemValue := 15
	numElemtns := 5

	numbers := generateRandomIntSlice(numElemtns, minElemValue, maxElemValue)
	algorithmName, _ := getRandomStringSliceValue(SortingAlgorithms)
	isAscending := generateRandomBool()

	runSort(algorithmName, isAscending, numbers)
}

func nerdMode() {
	sortingName := getSortingAlgoNameFromConsole()
	numbers := getNumberArrayFromConsole()
	isAscending := getIsAscendingFromConsole()
	runSort(sortingName, isAscending, numbers)
}

func getIsAscendingFromConsole() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Is sort ascending? [Y/n]")
	isAsc, _ := reader.ReadString('\n')
	return isYesInput(isAsc)
}

func getSortingAlgoNameFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please, select sorting algorithm:")

		for i := 1; i < 4; i++ {
			fmt.Printf("%v - %v \n", i, SortingAlgoMap[i])
		}

		sortingNumber, _ := reader.ReadString('\n')
		num, err := getNumberFromString(sortingNumber)
		if err == nil {
			sortingName := SortingAlgoMap[num]
			if sortingName != "" {
				fmt.Printf("\nSelected sorting algorithm: %v \n", sortingName)
				return sortingName
			}
		}
	}
}

func getNumberArrayFromConsole() []int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter numbers to sort. Example: 15 0 2 1 4")
		s, _ := reader.ReadString('\n')
		spl := splitStringBySpaces(s)
		out := stringArrayToIntArray(spl)
		if len(out) > 0 {
			fmt.Printf("Your numbers: %v \n", out)
			return out
		}
	}
}

func main() {
	introductionString := "This is a simple program to sort integer numbers with different algorithms.\n"
	fmt.Println(introductionString)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Are you lazy? [Y/n]")
		lazyCheck, _ := reader.ReadString('\n')
		isLazyUser := isYesInput(lazyCheck)
		if isLazyUser {
			lazyMode()
		} else {
			nerdMode()
		}

		fmt.Println("\nWould you like to repeat? [Y/n]")
		repeatCheck, _ := reader.ReadString('\n')
		doRepeat := isYesInput(repeatCheck)
		if !doRepeat {
			break
		}
	}
}
