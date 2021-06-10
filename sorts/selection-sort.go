package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

// Pseudocode
// Given list:
// 23  6   89  3  56   48

// Get len of list (n)
// Set pointer (pos=0), find min number in list, move number to pos 0 and shuffle all other items by 1, update their positions in the array
// Increment pointer: pos+1 and n = end of list
// find min number in given list
// Move minimum number to pointer position, update their positions in the array
// Repeat until n

func main() {

	args := os.Args
	log.Println(args)
	list := args[1:]

	// Get len of list (n)
	length := len(list)

	// covert string arr to ints
	finalList := []int32{}
	for _, item := range list {
		value, _ := strconv.ParseInt(item, 0, 16)
		finalList = append(finalList, int32(value))
	}

	for i := 0; i < length; i++ {
		min, position, err := getMinAndPosition(finalList, i, length)
		if err != nil {
			return
		}

		if position != i {
			finalList[position] = finalList[i]
			finalList[i] = min
		}
	}
	log.Println("finalList: ", finalList)
}

// getMinAndPosition gets the minimum number and returns the position its found at
func getMinAndPosition(values []int32, current int, length int) (int32, int, error) {
	if len(values) == 0 {
		return int32(0), int(0), errors.New("cannot detect a minimum value in an empty slice")
	}

	positionOfMin := int(current)
	min := values[positionOfMin]

	for i := current; i < length; i++ {
		if values[i] < min {
			min = values[i]
			positionOfMin = int(i)
		}
	}

	return min, positionOfMin, nil
}
