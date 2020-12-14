package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fileBytes, err := ioutil.ReadFile("report.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")
	nums := convertArray(sliceData)

	var target int = 2020
	sort.Ints(nums)
	part1 := findTwoEntries(nums, target)
	fmt.Println(part1)
	part2 := findThreeEntries(nums, target)
	fmt.Println(part2)

}

func convertArray(stringSlice []string) []int {
	numslice := make([]int, len(stringSlice)-1)
	for i := range numslice {
		numslice[i], err = strconv.Atoi(stringSlice[i])
	}
	return numslice
}

func findTwoEntries(nums []int, target int) int {
	for i, v := range nums {
		search := target - v
		index := sort.SearchInts(nums, search)
		if search == nums[index] {
			fmt.Println(nums[i], nums[index])
			return nums[i] * nums[index]
		}
	}
	return -1
}

func findThreeEntries(nums []int, target int) int {
	for i := range nums {
		for j := 0; j < len(nums); j++ {
			search := target - (nums[i] + nums[j])
			index := sort.SearchInts(nums, search)
			if nums[index] == search {
				fmt.Println(nums[i], nums[j], nums[index])
				return nums[i] * nums[j] * nums[index]
			}
		}
	}
	return -1
}
