package main

import (
	"fmt"

	"github.com/L1z1ng3r-sswe/computer_science/data_structures/binary_search_tree"
)

func main() {
	root := &binary_search_tree.BST{Val: 10}

	root.Add(5)
	root.Add(15)
	root.Add(3)
	root.Add(7)
	root.Add(12)
	root.Add(18)

	fmt.Println("BST after adding values:")
	fmt.Print(root.PrettyVisualise(0))

	searchVal := 7
	node := root.Search(searchVal)
	if node != nil {
		fmt.Printf("\nFound %d in the BST\n", searchVal)
	} else {
		fmt.Printf("\n%d not found in the BST\n", searchVal)
	}

	deleteVal := 5
	root = root.Delete(deleteVal)
	fmt.Printf("\nBST after deleting %d:\n", deleteVal)
	fmt.Print(root.PrettyVisualise(0))
}
