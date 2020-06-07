package BinaryTree

//      1
//    /   \
//  2       3
// / \    /  \
//4   5  6    7
//      / \
//     8   9

func initTree() *Tree {
	var node9 = &Tree{
		Value: 9,
	}
	var node8 = &Tree{
		Value: 8,
	}
	var node7 = &Tree{
		Value: 7,
	}
	var node6 = &Tree{
		Value: 6,
		Left:  node8,
		Right: node9,
	}
	var node5 = &Tree{
		Value: 5,
	}
	var node4 = &Tree{
		Value: 4,
	}
	var node3 = &Tree{
		Value: 3,
		Left:  node6,
		Right: node7,
	}
	var node2 = &Tree{
		Value: 2,
		Left:  node4,
		Right: node5,
	}
	var root = &Tree{
		Value:  1,
		Left:   node2,
		Right:  node3,
		IsRoot: true,
	}

	return root
}

func ExamplePreOrder() {
	tree := initTree()
	PreOrder(tree)
	//output: 1
	//2
	//4
	//5
	//3
	//6
	//8
	//9
	//7
}

func ExampleInOrder() {
	tree := initTree()
	InOrder(tree)
	//output: 4
	//2
	//5
	//1
	//8
	//6
	//9
	//3
	//7
}

func ExamplePostOrder() {
	tree := initTree()
	PostOrder(tree)
	// Output: 4
	//5
	//2
	//8
	//9
	//6
	//7
	//3
	//1
}
