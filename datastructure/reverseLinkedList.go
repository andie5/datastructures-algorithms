package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

/*
 * Complete the 'reverse' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts INTEGER_SINGLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Read the list and store 1st -test cases
	// Store length (2nd num)
	// Iterate ll - store in new ds
	// Iterate new ds from end and build new ll

	nodeStart := &SinglyLinkedListNode{}
	if llist.next == nil {
		return llist
	} else {

		temp := []int32{}
		actualData := llist

		for actualData.next != nil {
			temp = append(temp, actualData.data)
			if actualData.next != nil {
				actualData = actualData.next
			}
		}
		temp = append(temp, actualData.data)

		length := len(temp)
		previous := &SinglyLinkedListNode{}

		for j := length - 1; j >= 0; j-- {
			fmt.Println("j: ", j)

			node := SinglyLinkedListNode{data: temp[j]}
			current := node

			if previous != nil {
				fmt.Println("does this print every time")
				previous.next = &current
			}
			previous = &current
			if j == length-1 {
				// nodeStart = &node
				fmt.Println("node: ", node)
				fmt.Println("current: ", current)
				nodeStart = &current
			}
		}
	}
	return nodeStart

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	testsTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tests := int32(testsTemp)

	for testsItr := 0; testsItr < int(tests); testsItr++ {
		llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		llist := SinglyLinkedList{}
		for i := 0; i < int(llistCount); i++ {
			llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			llistItem := int32(llistItemTemp)
			llist.insertNodeIntoSinglyLinkedList(llistItem)
		}

		llist1 := reverse(llist.head)

		printSinglyLinkedList(llist1, " ", writer)
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}