package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	sort "github.com/datastructure-algorithms/sorts"
)

func main() {

	args := os.Args
	log.Println(args)
	list := args[1:]

	// convert string arr to ints
	finalList := []int32{}
	for _, item := range list {
		value, _ := strconv.ParseInt(item, 0, 16)
		finalList = append(finalList, int32(value))
	}
	fmt.Println("Start sorts......")
	fmt.Println("Entry list is: ", finalList, "\n")

	sort.InsertionSort(finalList)
	sort.SelectionSort(finalList)
	sort.BubbleSort(finalList)

}
