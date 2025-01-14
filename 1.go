package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in, out := getBuffers()
	defer out.Flush()

	var setsCount int
	fmt.Fscan(in, &setsCount)

	for set := 1; set <= setsCount; set++ {
		var count int

		var line string

		fmt.Fscanln(in, &count)

		arrayToValidate := readInts(line)
		validArray := readInts(line)

		sortedArray := make([]int, len(arrayToValidate))
		copy(sortedArray, arrayToValidate)
		sort.Ints(sortedArray)

		opt := validate(sortedArray, validArray, count)
		fmt.Fprintln(out, opt)
	}
}

func readInts(s string) []int {
	var nums []int
	for _, strNum := range strings.Fields(s) {
		num, _ := strconv.Atoi(strNum)
		nums = append(nums, num)
	}
	return nums
}

func validate(sortedArray []int, validArray []int, count int) string {
	if len(sortedArray) != count || len(validArray) != count {
		return "no"
	}

	validArrayStr := make([]string, len(validArray))
	for i, val := range validArray {
		validArrayStr[i] = strconv.Itoa(val)
	}

	splitString := strings.Split(strings.Join(validArrayStr, " "), " ")

	validInts := make([]int, 0, count)
	for _, strNum := range splitString {
		num, err := strconv.Atoi(strNum)
		if err == nil {
			validInts = append(validInts, num)
		}
	}

	if len(validInts) != count {
		return "no"
	}

	if !slices.Equal(sortedArray, validInts) {
		return "no"
	}

	return "yes"
}

func getBuffers() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}
