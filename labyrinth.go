package main

import (
	"fmt"
)

func main() {
	/* Initialization of structure. Root addition. */
	labyrinth := &Node{Val: 100, L: nil, R: nil,
		M: &Meta{ThreadS: []string{"."}, AssocN: map[int][]string{
			100: {"."}},
		},
	}
	fmt.Printf("\nCreate Labyrinth, Adding Values\n")
	fmt.Printf("Working on %v (ROOT)...  %v added. ThreadS Route: %v\n",
		labyrinth.Val, labyrinth.Val, labyrinth.M.ThreadS)

	/* Add values to labyrinth */
	values := []int{19, 7, 24, 150, 140, 276, 199, 2, 144}
	for _, v := range values {
		thrStr := make([]string, 0)  // Create [] with make - Strings
		thrStr = append(thrStr, ".") // Init value of "." for Root
		fmt.Printf("Working on %3d... ", v)
		// Add Node (as key) and thread (as value) at the top level map
		labyrinth.M.AssocN[v] = labyrinth.AddNode(v, thrStr)
	}
	fmt.Printf("Entire Map: %v\n\n", labyrinth.M.AssocN)

	/* Thraverse the entire tree */
	labyrinth.Traverse()

	/* Search for specific values NOT USING Binary search. */
	labyrinth.Search(100)
	labyrinth.Search(1561)
	labyrinth.Search(140)
	labyrinth.Search(199)
	labyrinth.Search(7)
	labyrinth.Search(19)
}

// Node struct as the main tree struct
type Node struct {
	Val int   // Value of Node/Leaf
	L   *Node // Left node
	R   *Node // Right node
	M   *Meta // Meta data embedded struct
}

// Meta struct monitoring the routing information of the node/leaf
type Meta struct {
	ThreadS []string         // Thread with strings
	AssocN  map[int][]string // Associates Node values with thread(s)
}

// Method Search() on *Node type. With receivers argument list of type *Node,
// named n. Method search if node exists, and returns true/false and entire path.
// Then, it goes to the Node, by calling the WalkTheLabyrinth() method.
func (n *Node) Search(key int) {
	// Search the map with Node value as key. If ok, then the Node exists ...
	if v, ok := n.M.AssocN[key]; ok {
		fmt.Printf("Node %v exist.\tThread: %v ... ", key, v)

		// If Node found on map, and it's not the Root, GO TO that Node ->
		if l := len(v); l > 1 {
			v = v[1:]                                 // Cut the first "." (the root)
			fmt.Printf("%v\n", n.WalkTheLabyrinth(v)) // Pass the rest []
		} else {
			fmt.Printf("%v %t (ROOT)\n", n.Val, true) // else, it's the Root Node
		}

	} else {
		// ... if not, it doesn't exist at all. We dont's have to search the tree!!!
		fmt.Printf("Node %v doesn't exist.\n", key)
	}
}

// Method WalkTheLabirynth() on *Note type. With receivers arguments list
// of *Node, named n. Searches for specific node.
func (n *Node) WalkTheLabyrinth(thread []string) bool {
	if l := len(thread); l == 0 { // If length is 0, we reach the Node
		fmt.Printf("%v ", n.Val)
		return true
	}
	// In every recursion call, we re-slice the []thread low bound by 1
	switch thread[0] {
	case "L":
		return n.L.WalkTheLabyrinth(thread[1:]) // Tail Recursive for left *Node
	case "R":
		return n.R.WalkTheLabyrinth(thread[1:]) // Tail Recursive for right *Node
	}
	return false
}

// Method Traverse() on *Node struct type. With receiver argument list
// of type *Node, named n. Dumps the entire tree.
func (n *Node) Traverse() {
	if n == nil {
		fmt.Printf("Receiver is <nil>\n\n")
		return
	}

	switch {
	case n.L != nil && n.R != nil:
		fmt.Printf("Node: %3d  Thread: %v\tChild: (L: %v,  R: %v)\n",
			n.Val, n.M.ThreadS, n.L.Val, n.R.Val)
	case n.L == nil && n.R == nil:
		fmt.Printf("Node: %3d  Thread: %v\tChild: (L: %v,  R: %v)\n",
			n.Val, n.M.ThreadS, "<nil>", "<nil>")
	case n.L == nil && n.R != nil:
		fmt.Printf("Node: %3d  Thread: %v\tChild: (L: %v,  R: %v)\n",
			n.Val, n.M.ThreadS, "<nil>", n.R.Val)
	case n.L != nil && n.R == nil:
		fmt.Printf("Node: %3d  Thread: %v\tChild: (L: %v,  R: %v)\n",
			n.Val, n.M.ThreadS, n.L.Val, "<nil>")
	}

	if n.L != nil || n.R != nil {
		n.L.Traverse()
		n.R.Traverse()
	}
}

// Method AddNode on Node struct type. With receiver Argument List
// of type *Node, named n. Duplicates NOT allowed.
func (n *Node) AddNode(v int, thrStr []string) []string {
	switch {
	/* Found on root */
	case v == n.Val:
		fmt.Printf("Value: %v found on root: (%v). Thread: %v\n",
			v, n.Val, n.M.ThreadS)

	/* Go to left Node */
	case v < n.Val:
		thrStr = append(thrStr, "L") // Update the []thread with Left move
		switch {
		case n.L == nil:
			// add LEFT node and the thread of it, at the currect level's Node fields
			n.L = &Node{Val: v, M: &Meta{ThreadS: thrStr}, L: nil, R: nil}
			fmt.Printf("\t%3d added. Thread: %v\n", n.L.Val, n.L.M.ThreadS)
		case v == n.L.Val:
			fmt.Printf("Value %v already exists (%v)\n", v, n.L.Val)
		default:
			return n.L.AddNode(v, thrStr)
		}

	/* Go to right Node */
	case v > n.Val:
		thrStr = append(thrStr, "R") // Update the []thread with Right move
		switch {
		case n.R == nil:
			// add RIGHT node and the thread of it, at the currect level's Node fields
			n.R = &Node{Val: v, M: &Meta{ThreadS: thrStr}, L: nil, R: nil}
			fmt.Printf("\t%3d added. Thread: %v\n", n.R.Val, n.R.M.ThreadS)
		case v == n.R.Val:
			fmt.Printf("Value %v already exists (%v)\n", v, n.R.Val)
		default:
			return n.R.AddNode(v, thrStr)
		}
	}
	// Return the thread to be added as value to the main map
	return thrStr
}
