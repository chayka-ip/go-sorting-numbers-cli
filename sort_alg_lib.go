package main

import "fmt"

var BubbleSortName = "Bubble sort"
var InsertionSortName = "Insertion sort"
var QuickSortName = "Quick Sort"

var SortingAlgorithms = []string{BubbleSortName, InsertionSortName, QuickSortName}
var SortingAlgoMap = map[int]string{1: BubbleSortName, 2: InsertionSortName, 3: QuickSortName}

func SortByAlgorithm(name string, numberArray []int, isAscending bool) {
	switch name {
	case BubbleSortName:
		BubbleSortInPlace(numberArray, isAscending)
	case InsertionSortName:
		InsertionSortInPlace(numberArray, isAscending)
	case QuickSortName:
		QuickSortInPlace(numberArray, isAscending)
	default:
		fmt.Println("This algorithm is not available")
	}
}

func BubbleSortInPlace(numberArray []int, isAscending bool) {
	arrLen := len(numberArray)
	for i := 0; i < arrLen; i++ {
		j_limit := arrLen - i - 1
		hasSwap := false
		for j := 0; j < j_limit; j++ {
			left := numberArray[j]
			right := numberArray[j+1]

			ascSort := isAscending && left > right
			dscSort := !isAscending && left < right
			doSwap := ascSort || dscSort

			if doSwap {
				numberArray[j] = right
				numberArray[j+1] = left
				hasSwap = true
			}
		}
		if !hasSwap {
			return
		}
	}
}

func InsertionSortInPlace(numberArray []int, isAscending bool) {
	arrLen := len(numberArray)

	for i := 1; i < arrLen; i++ {
		j := i - 1
		keyElem := numberArray[i]

		for {
			compElem := numberArray[j]

			ascSort := isAscending && (compElem > keyElem)
			dscSort := !isAscending && (compElem < keyElem)
			doSwap := ascSort || dscSort

			if doSwap {
				numberArray[j+1] = compElem
			} else {
				numberArray[j+1] = keyElem
				break
			}

			j--
			if j < 0 {
				numberArray[0] = keyElem
				break
			}
		}
	}
}

func QuickSortInPlace(numberArray []int, isAscending bool) {
	arrLen := len(numberArray)
	quickSortInPlaceBody(numberArray, 0, arrLen-1)
	if !isAscending {
		// todo: update quicksort to accept isAscending param
		reverseIntSliceInPlace(numberArray)
	}
}

func quickSortInPlaceBody(numberArray []int, start int, end int) {
	arrLen := len(numberArray)
	if arrLen < 2 || (start >= end) {
		return
	}

	i, j := start, end
	pivot := int((start + end) / 2)
	pivotElem := numberArray[pivot]

	for {

		for numberArray[i] < pivotElem {
			i++
		}

		for numberArray[j] > pivotElem {
			j--
		}

		if i <= j {
			iElem := numberArray[i]
			jElem := numberArray[j]
			numberArray[i] = jElem
			numberArray[j] = iElem
			i++
			j--
			break
		}
	}

	quickSortInPlaceBody(numberArray, start, j)
	quickSortInPlaceBody(numberArray, i, end)
}
