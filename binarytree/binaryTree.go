// Node = data
// Left pointer
// Right pointer

// Get length of array
// Create node - Root of the tree is the first item in the list - parent node
// -- Node (value, left pointer, right pointer)
// Move to next item in the list and iterate until you reach the end of the list
// ---if value is less then current node value, go to left pointer (continue)
// ---If value is more than current node value go to right pointer (continue)
// ---If node value is empty - create new node with the current value from the list and left and right pointer is empty else go to part 3a

package main

import (
	"fmt"
	"os"
	"strconv"
)

type BinaryNode struct {
	Value    int32
	LeftPtr  *BinaryNode
	RightPtr *BinaryNode
}

// 23 4 78 5 27

func main() {
	args := os.Args[1:]

	list := []int32{}
	for _, item := range args {
		itemInt, err := strconv.ParseInt(item, 0, 32)
		intBase32 := int32(itemInt)

		if err != nil {
			fmt.Println("err: ", err)
		}
		list = append(list, intBase32)
	}

	// Get length of list
	length := len(list)
	var tree BinaryNode

	if length > 0 {
		// Create base case node
		tree = createNode(list[0])
		fmt.Println("tree: ", tree)

		// Loop through the rest of the list and create nodes
		for i := 1; i < length; i++ {
			insertNode(&tree, list[i])
			fmt.Println("tree so far:", tree)
		}
	}
	fmt.Println("================")
	fmt.Println("================")

	updatedTree := deleteNode(&tree, 23)
	fmt.Println("updatedTree: ", updatedTree)

	// fmt.Println("Is node in tree: ", findNodeInTree(&tree, int32(23)))
}

//findNodeInTree finds a node in the binaryTree given a value
func findNodeInTree(binaryTree *BinaryNode, value int32) *BinaryNode {
	// Set current node to parent node
	current := binaryTree
	fmt.Println("value to check: ", value)

	for current != nil {
		if current.Value == value {
			return current
		} else if value <= current.Value {
			fmt.Println("value less: ", current.Value)
			current = current.LeftPtr
		} else {
			fmt.Println("value more: ", current.Value)
			current = current.RightPtr
		}
	}
	return current
}

// search key in the BST and set its parent pointer
// update the parent to the current node
// if the given key is less than the current node, go to the left subtree;
// otherwise, go to the right subtree
// return if the key is not found in the tree

// Case 1: node to be deleted has no children, i.e., it is a leaf node
// if the node to be deleted is not a root node, then set its
// parent left/right child to null
// if the tree has only a root node, set it to null

// Case 2: node to be deleted has two children
// find its inorder successor node
// store successor value
// recursively delete the successor. Note that the successor
// will have at most one child (right child)
// copy value of the successor to the current node

// Case 3: node to be deleted has only one child
// choose a child node
// if the node to be deleted is not a root node, set its parent
// to its child
// if the node to be deleted is a root node, then set the root to the child
// return if the key is not found in the tree
func deleteNode(binaryTree *BinaryNode, value int32) *BinaryNode {

	if binaryTree == nil {
		return binaryTree
	}

	var parent BinaryNode

	// start with root
	current := binaryTree
	fmt.Println("in delete node")

	for current != nil {
		// update the parent to the current node
		parent = *current

		fmt.Println("current val: ", value)
		// If value is found at at current node
		if current.Value == value {
			fmt.Println("value found: ", current.Value)
			//If there is no children...
			if current.LeftPtr == nil && current.RightPtr == nil {
				fmt.Println("no children")
				// if the node to be deleted is not a root node, then set its parent left/right child to null
				if current != binaryTree {
					if parent.LeftPtr == current {
						parent.LeftPtr = nil
					} else {
						parent.RightPtr = nil
					}
				} else {
					binaryTree = nil
				}
			} else if current.LeftPtr != nil && current.RightPtr != nil { //Case 2: node to be deleted has both children
				// if the node to be deleted is not a root node, then set its parent left/right
				fmt.Println("both children left side: ", current.LeftPtr.Value)
				fmt.Println("both children right side: ", current.RightPtr.Value)
				if current != binaryTree {
					// current.LeftPtr.RightPtr = current.LeftPtr
					// current = current.LeftPtr
					current.LeftPtr = deleteNode(current.LeftPtr, value)
					current.RightPtr = deleteNode(current.RightPtr, value)
				} else { // if the node to be deleted is a root node, then set the root to the child
					current.LeftPtr.RightPtr = current.RightPtr
					current = current.LeftPtr
				}
			} else { // Case 3: node to be deleted has only one child
				node := &BinaryNode{}
				fmt.Println("one child: ", current.RightPtr.Value)
				if current.LeftPtr != nil {
					// node = current.LeftPtr
					current = current.LeftPtr
				} else {
					node = current.RightPtr
				}
				if current != binaryTree {

					if parent.LeftPtr == current {
						parent.LeftPtr = node
					} else { // if the node to be deleted is not a root node, set its parent
						// to its child
						parent.RightPtr = node
					}
				} else { // if the node to be deleted is a root node, then set the root to the child
					// return if the key is not found in the tree
					binaryTree = node
				}
			}
		}
		if value <= current.Value {
			current = deleteNode(current.LeftPtr, value)
		} else {
			current = deleteNode(current.RightPtr, value)
		}
	}
	return current
}

func insertNode(node *BinaryNode, value int32) *BinaryNode {
	// If value is less then or equal to current node value, go to left pointer
	tree := node
	if value <= node.Value {
		if node.LeftPtr == nil {
			temp := createNode(value)
			node.LeftPtr = &temp
			fmt.Println("create left node: ", temp)
		} else {
			tree = insertNode(node.LeftPtr, value)
		}
		// Value is more than current node value go to right pointer
	} else {
		if node.RightPtr == nil {
			temp := createNode(value)
			node.RightPtr = &temp
			fmt.Println("create right node: ", temp)
		} else {
			tree = insertNode(node.RightPtr, value)
		}
	}
	return tree
}

func createNode(value int32) BinaryNode {
	return BinaryNode{Value: value}
}
