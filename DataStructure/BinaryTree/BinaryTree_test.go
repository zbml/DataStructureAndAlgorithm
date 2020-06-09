package BinaryTree

//      1
//    /   \
//  2       3
// / \    /  \
//4   5  6    7
//      / \
//     8   9

func initTree1() *Tree {
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

//   1
//  / \
// 2  3
// /   \
//4     5
// \
//  6
// /  \
//7    8
func initTree2() *Tree {
	var node7 = &Tree{
		Value: 7,
	}
	var node8 = &Tree{
		Value: 8,
	}
	var node6 = &Tree{
		Value: 6,
		Left:  node7,
		Right: node8,
	}
	var node4 = &Tree{
		Value: 4,
		Right: node6,
	}
	var node5 = &Tree{
		Value: 5,
	}
	var node2 = &Tree{
		Value: 2,
		Left:  node4,
	}
	var node3 = &Tree{
		Value: 3,
		Right: node5,
	}
	var node1 = &Tree{
		Value:  1,
		Left:   node2,
		Right:  node3,
		IsRoot: true,
	}

	return node1
}

func ExamplePreOrderTree1() {
	tree1 := initTree1()
	PreOrder(tree1)
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

func ExamplePreOrderTree2() {
	tree2 := initTree2()
	PreOrder(tree2)
	//output: 1
	//2
	//4
	//6
	//7
	//8
	//3
	//5
}

func ExamplePreOrder2Tree1() {
	tree := initTree1()
	PreOrder2(tree)
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

func ExamplePreOrder2Tree2() {
	tree2 := initTree2()
	PreOrder2(tree2)
	//output: 1
	//2
	//4
	//6
	//7
	//8
	//3
	//5
}

func ExampleInOrder() {
	tree := initTree1()
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

func ExampleInOrder2() {
	tree := initTree1()
	InOrder2(tree)
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
	tree := initTree1()
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

func ExamplePostOrder2() {
	tree := initTree1()
	PostOrder2(tree)
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
