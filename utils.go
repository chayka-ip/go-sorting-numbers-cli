package main

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func tryConvertStringToInt(s string) (int, error) {
	n, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		d := int(n)
		return d, nil
	}
	return -1, err
}

func reverseIntSliceInPlace(slice []int) {
	slice_copy := append(make([]int, 0, len(slice)), slice...)
	j := 0
	for i := len(slice) - 1; i >= 0; i-- {
		slice[j] = slice_copy[i]
		j++
	}
}

func generateRandomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func generateRandomIntSlice(numElem int, min int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	out := []int{}

	if (min < 0) || (max < 0) {
		return out
	}

	for {
		elem := generateRandomInt(min, max)
		out = append(out, elem)

		numElem--
		if numElem < 1 {
			return out
		}
	}
}

func getFirstCharFromString(s string) string {
	if len(s) > 0 {
		return s[0:1]
	}
	return ""
}

func getRandomStringSliceValue(s []string) (string, error) {
	n := len(s)
	if n == 0 {
		return "", errors.New("Empty array")
	}

	index := generateRandomInt(0, len(s)-1)
	return s[index], nil
}

func generateRandomBool() bool {
	v := generateRandomInt(0, 1)
	if v == 1 {
		return true
	}
	return false
}

func splitStringBySpaces(s string) []string {
	r := regexp.MustCompile("[^\\s]+")
	return r.FindAllString(s, -1)
}

func stringArrayToIntArray(input []string) []int {
	out := []int{}
	for i := 0; i < len(input); i++ {
		val, err := getNumberFromString(input[i])
		if err == nil {
			out = append(out, val)
		}
	}
	return out
}

func printLines() {
	fmt.Println("__________________________________________")
}
