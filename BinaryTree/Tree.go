package main

import (
	"fmt"
	"math"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}


func LevelOrderBinaryTree(arr []int) *Tree {
	tree := new(Tree)
	tree.root = levelOrderBinaryTree(arr, 0, len(arr))
	return tree
}

func levelOrderBinaryTree(arr []int, start int, size int) *Node {
	curr := &Node{arr[start], nil, nil}

	left := 2*start + 1
	right := 2*start + 2

	if left < size {
		curr.left = levelOrderBinaryTree(arr, left, size)
	}
	if right < size {
		curr.right = levelOrderBinaryTree(arr, right, size)
	}
	return curr
}

func (t *Tree) Add(value int) {
	t.root = addUtil(t.root, value)
}

func addUtil(n *Node, value int) *Node {
	if n == nil {
		n = new(Node)
		n.value = value
		return n
	}
	if value < n.value {
		n.left = addUtil(n.left, value)
	} else {
		n.right = addUtil(n.right, value)
	}
	return n
}

func (t *Tree) PrintPreOrder() {
	fmt.Print("Pre Order")
	printPreOrder(t.root)
	fmt.Println()
}

func printPreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.value, " ")
	printPreOrder(n.left)
	printPreOrder(n.right)
}

func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
	if node != nil {
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) PrintPostOrder() {
	fmt.Print("Post Order: ")
	printPostOrder(t.root)
	fmt.Println()
}

func printPostOrder(n *Node) {
	if n == nil {
		return
	}
	printPostOrder(n.left)
	printPostOrder(n.right)
	fmt.Print(n.value, " ")
}

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthPostOrder(node.left, index, counter)
		nthPostOrder(node.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Println(node.value)
		}
	}
}

func (t *Tree) PrintInOrder() {
	fmt.Print("In Order")
	printInOrder(t.root)
	fmt.Println()
}

func printInOrder(n *Node) {
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Print(n.value, " ")
	printInOrder(n.right)
}

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthInOrder(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthInOrder(node.right, index, counter)
	}
}

func (t *Tree) PrintBreadthFirst() {
	que := new(queue.Queue)
	var temp *Node

	if t.root != nil {
		que.Enqueue(t.root)
	}
	fmt.Println("Breadth First")
	for que.Len() != 0 {
		temp2 := que.Dequeue()
		temp = temp2.(*Node)

		fmt.Print(temp.value, " ")

		if temp.left != nil {
			que.Enqueue(temp.left)
		}
		if temp.right != nil {
			que.Enqueue(temp.right)
		}
	}
	fmt.Println()
}


func (t *Tree) PrintLevelOrderLineByLine() {
	que1 := new(queue.Queue)
	que2 := new(queue.Queue)
	var temp *Node

	if t.root != nil {
		que1.Enqueue(t.root)
	}

	for que1.Len() != 0 || que2.Len() != 0 {
		for que1.Len() != 0 {
			temp2 := que1.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")

			if temp.left != nil {
				que2.Enqueue(temp.left)
			}
			if temp.right != nil {
				que2.Enqueue(temp.right)
			}
		}
		fmt.Println(" ")		
		for que2.Len() != 0{
			temp2 := que2.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que1.Enqueue(temp.left)
			}
			if temp.right != nil {
				que1.Enqueue(temp.right)
			}
		}
		fmt.Println(" ")		
	}
}


func (t *Tree) PrintLevelOrderLineByLine2() {
	que := new(queue.Queue)
	var temp *Node
	var count int

	if t.root != nil {
		que.Enqueue(t.root)
	}

	for que.Len() != 0 {
		
		count = que.Len()
		
		for count > 0 {
			temp2 := que.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que.Enqueue(temp.left)
			}
			if temp.right != nil {
				que.Enqueue(temp.right)
			}
			count -= 1
		}
		fmt.Println(" ")		
	}
}

func (t *Tree) PrintSpiralTree() {
	stk1 := new(stack.Stack)
	stk2 := new(stack.Stack)
	var temp *Node

	if t.root != nil {
		stk1.Push(t.root)
	}

	for stk1.Len() != 0 || stk2.Len() != 0 {
		for stk1.Len() != 0 {
			temp2 := stk1.Pop()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				stk2.Push(temp.left)
			}
			if temp.right != nil {
				stk2.Push(temp.right)
			}
		}
		fmt.Println(" ")		
		for stk2.Len() != 0{
			temp2 := stk2.Pop()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.right != nil {
				stk1.Push(temp.right)
			}
			if temp.left != nil {
				stk1.Push(temp.left)
			}
			
		}
		fmt.Println(" ")		
	}
}


func (t *Tree) Find(value int) bool {
	var curr *Node = t.root
	for curr != nil {
		if curr.value == value {
			return true
		} else if curr.value > value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}




func (t *Tree) FindMin() (int, bool) {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}

	for node.left != nil {
		node = node.left
	}
	return node.value, true
}

func (t *Tree) FindMax() (int, bool) {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return 0, false
	}

	for node.right != nil {
		node = node.right
	}
	return node.value, true
}

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.right != nil {
		node = node.right
	}
	return node
}

func (t *Tree) FindMinNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.left != nil {
		node = node.left
	}
	return node
}

func FindMax(curr *Node) *Node {
	var node *Node = curr
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.right != nil {
		node = node.right
	}
	return node
}

func FindMin(curr *Node) *Node {
	var node *Node = curr
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}

	for node.left != nil {
		node = node.left
	}
	return node
}

func (t *Tree) Free() {
	t.root = nil
}

func (t *Tree) DeleteNode(value int) {
	t.root = DeleteNode(t.root, value)
}

func DeleteNode(node *Node, value int) *Node {
	var temp *Node = nil
	if node != nil {
		if node.value == value {
			if node.left == nil && node.right == nil {
				return nil
			}
			if node.left == nil {
				temp = node.right
				return temp
			}
			if node.right == nil {
				temp = node.left
				return temp
			}

			maxNode := FindMax(node.left)
			maxValue := maxNode.value
			node.value = maxValue
			node.left = DeleteNode(node.left, maxValue)
		} else {
			if node.value > value {
				node.left = DeleteNode(node.left, value)
			} else {
				node.right = DeleteNode(node.right, value)
			}
		}
	}
	return node
}

func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}

func treeDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lDepth := treeDepth(root.left)
	rDepth := treeDepth(root.right)

	if lDepth > rDepth {
		return lDepth + 1
	}
	return rDepth + 1
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func isEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) &&
			isEqual(node1.left, node2.left) &&
			isEqual(node1.right, node2.right))
	}
}

func (t *Tree) Ancestor(first int, second int) *Node {
	if first > second {
		temp := first
		first = second
		second = temp
	}
	return Ancestor(t.root, first, second)
}

func Ancestor(curr *Node, first int, second int) *Node {
	if curr == nil {
		return nil
	}
	if curr.value > first && curr.value > second {
		return Ancestor(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return Ancestor(curr.right, first, second)
	}
	return curr
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.left = copyTree(curr.left)
		temp.right = copyTree(curr.right)
		return temp
	}
	return nil
}

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func copyMirrorTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.right = copyMirrorTree(curr.left)
		temp.left = copyMirrorTree(curr.right)
		return temp
	}
	return nil
}

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return (1 + numNodes(curr.right) + numNodes(curr.left))
}

func (t *Tree) NumFullNodesBT() int {
	return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *Node) int {
	var count int
	if curr == nil {
		return 0
	}

	count = numFullNodesBT(curr.right) + numFullNodesBT(curr.left)
	if curr.right != nil && curr.left != nil {
		count++
	}
	return count
}

func (t *Tree) MaxLengthPathBT() int {
	return maxLengthPathBT(t.root)
}

func maxLengthPathBT(curr *Node) int {
	var max, leftPath, rightPath, leftMax, rightMax int

	if curr == nil {
		return 0
	}

	leftPath = treeDepth(curr.left)
	rightPath = treeDepth(curr.right)
	max = leftPath + rightPath + 1

	leftMax = maxLengthPathBT(curr.left)
	rightMax = maxLengthPathBT(curr.right)
	if leftMax > max {
		max = leftMax
	}

	if rightMax > max {
		max = rightMax
	}
	return max
}

func (t *Tree) NumLeafNodes() int {
	return numLeafNodes(t.root)
}

func numLeafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	}
	return (numLeafNodes(curr.right) + numLeafNodes(curr.left))
}

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
	var sum, leftSum, rightSum int
	if curr == nil {
		return 0
	}

	rightSum = sumAllBT(curr.right)
	leftSum = sumAllBT(curr.left)
	sum = rightSum + leftSum + curr.value
	return sum
}

func IsBST3(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left != nil && FindMax(root.left).value > root.value {
		return false
	}
	if root.right != nil && FindMin(root.right).value <= root.value {
		return false
	}
	return (IsBST3(root.left) && IsBST3(root.right))
}

func (t *Tree) IsBST() bool {
	return IsBST(t.root, math.MinInt32, math.MaxInt32)
}

func IsBST(curr *Node, min int, max int) bool {
	if curr == nil {
		return true
	}
	if curr.value < min || curr.value > max {
		return false
	}
	return IsBST(curr.left, min, curr.value) && IsBST(curr.right, curr.value, max)
}

func (t *Tree) IsBST2() bool {
	var c int
	return isBST2(t.root, &c)
}

func isBST2(root *Node, count *int) bool {
	var ret bool
	if root != nil {
		ret = isBST2(root.left, count)
		if !ret {
			return false
		}

		if *count > root.value {
			return false
		}

		*count = root.value

		ret = isBST2(root.right, count)
		if !ret {
			return false
		}
	}
	return true
}

func (t *Tree) IsCompleteTree() bool {
	return isCompleteTree(t.root)
}

func isCompleteTree(root *Node) bool {
	que := new(queue.Queue)
	var temp *Node
	var noChild = false

	if root != nil {
		que.Enqueue(root)
	}

	for que.Len() != 0{
		temp = que.Dequeue().(*Node)
		if temp.left != nil {
			if noChild == true {
				return false
			}
			que.Enqueue(temp.left)
		} else {
			noChild = true
		}

		if temp.right != nil {
			if noChild == true {
				return false
			}
			que.Enqueue(temp.right)
		} else {
			noChild = true
		}
	}
	return true
}

func (t *Tree) IsCompleteTree2() bool {
	count := t.NumNodes()
	return isCompleteTreeUtil(t.root, 0, count)
}

func isCompleteTreeUtil(curr *Node, index int, count int) bool {

		if (curr == nil){
			return true
		}
		if (index > count){
			return false
		}

	return isCompleteTreeUtil(curr.left, index * 2 + 1, count) && isCompleteTreeUtil(curr.right, index * 2 + 2, count)
}

func (t *Tree) IsHeap() bool {
	parentValue := -99999999
	return t.IsCompleteTree() && isHeapUtil(t.root, parentValue)
}

func isHeapUtil(curr *Node, parentValue int) bool {

		if (curr == nil){
			return true
		}
		if (curr.value < parentValue){
			return false
		}

	return isHeapUtil(curr.left, curr.value) && isHeapUtil(curr.right, curr.value)
}


func (t *Tree) IsHeap2() bool {
	count := t.NumNodes()
	parentValue := -99999999
	return isHeapUtil2(t.root, 0, count, parentValue)
}

func isHeapUtil2(curr *Node, index int, count int, parentValue int) bool {

		if (curr == nil){
			return true
		}
		if (index > count){
			return false
		}
		if (curr.value < parentValue){
			return false
		}

		return isHeapUtil2(curr.left, index * 2 + 1, count, curr.value) && isHeapUtil2(curr.right, index * 2 + 2, count, curr.value)
}


func (t *Tree) PrintAllPath() {
	stk := new(stack.Stack)
	printAllPath(t.root, stk)
}

func printAllPath(curr *Node, stk *stack.Stack) {
	if curr == nil {
		return
	}
	stk.Push(curr.value)
	if curr.left == nil && curr.right == nil {
		stk.Print()
		fmt.Println()
		stk.Pop()
		return
	}

	printAllPath(curr.right, stk)
	printAllPath(curr.left, stk)
	stk.Pop()
}

func (t *Tree) LCA(first int, second int) (int, bool) {
	ans := LCAUtil(t.root, first, second)
	if ans != nil {
		return ans.value, true
	}
	fmt.Println("NotFoundException")
	return 0, false
}

func LCAUtil(curr *Node, first int, second int) *Node {
	var left, right *Node

	if curr == nil {
		return nil
	}

	if curr.value == first || curr.value == second {
		return curr
	}

	left = LCAUtil(curr.left, first, second)
	right = LCAUtil(curr.right, first, second)

	if left != nil && right != nil {
		return curr
	} else if left != nil {
		return left
	} else {
		return right
	}
}

func (t *Tree) LcaBST(first int, second int) (int, bool) {
	return LcaBST(t.root, first, second)
}

func LcaBST(curr *Node, first int, second int) (int, bool) {
	if curr == nil {
		fmt.Println("NotFoundException")
		return 0, false
	}

	if curr.value > first && curr.value > second {
		return LcaBST(curr.left, first, second)
	}
	if curr.value < first && curr.value < second {
		return LcaBST(curr.right, first, second)
	}
	return curr.value, true
}

func (t *Tree) TrimOutsidedataRange(min int, max int) {
	t.root = trimOutsidedataRange(t.root, min, max)
}

func trimOutsidedataRange(curr *Node, min int, max int) *Node {
	if curr == nil {
		return nil
	}
	curr.left = trimOutsidedataRange(curr.left, min, max)
	curr.right = trimOutsidedataRange(curr.right, min, max)
	if curr.value < min {
		return curr.right
	}
	if curr.value > max {
		return curr.left
	}
	return curr
}

func (t *Tree) PrintDataInRange(min int, max int) {
	printDataInRange(t.root, min, max)
}

func printDataInRange(root *Node, min int, max int) {
	if root == nil {
		return
	}
	printDataInRange(root.left, min, max)
	if root.value >= min && root.value <= max {
		fmt.Print(root.value, " ")
	}
	printDataInRange(root.right, min, max)
}

func (t *Tree) FloorBST(val int) int {
	curr := t.root
	floor := math.MaxInt32

	for curr != nil {
		if curr.value == val {
			floor = curr.value
			break
		} else if curr.value > val {
			curr = curr.left
		} else {
			floor = curr.value
			curr = curr.right
		}
	}
	return floor
}

func (t *Tree) CeilBST(val int) int {
	curr := t.root
	ceil := math.MinInt32

	for curr != nil {
		if curr.value == val {
			ceil = curr.value
			break
		} else if curr.value > val {
			ceil = curr.value
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return ceil
}

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}

	max := curr.value
	left := findMaxBT(curr.left)
	right := findMaxBT(curr.right)
	if left > max {
		max = left
	}

	if right > max {
		max = right
	}
	return max
}

func SearchBT(root *Node, value int) bool {
	var left, right bool
	if root == nil || root.value == value {
		return false
	}

	left = SearchBT(root.left, value)
	if left {
		return true
	}

	right = SearchBT(root.right, value)
	if right {
		return true
	}
	return false
}

func CreateBinaryTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinaryTreeUtil(arr, 0, size-1)
	return t
}

func createBinaryTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	curr := new(Node)
	curr.value = arr[mid]
	curr.left = createBinaryTreeUtil(arr, start, mid-1)
	curr.right = createBinaryTreeUtil(arr, mid+1, end)
	return curr
}

func isBSTArray( preorder[] int, size int) bool {
	stk := new(stack.Stack)
	var value int
	root := -999999;
	for i := 0; i < size; i++ {
		value = preorder[i]

		// If value of the right child is less than root.
		if (value < root){
			return false
		}
		// First left child values will be popped
		// Last popped value will be the root.
		for (stk.Len() > 0 && stk.Peek().(int) < value){
			root = stk.Pop().(int)
		}
		// add current value to the stack.
		stk.Push(value)
	}
	return true
}

// Sort sorts values in place.
func Sort(values []int) {
	t := new(Tree)
	for _, v := range values {
		t.Add(v)
	}
	appendValues(values[:0], t.root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Node) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t := CreateBinaryTree(arr)

	//t := LevelOrderBinaryTree(arr)
	t.PrintPreOrder()
	t.PrintLevelOrderLineByLine()
	t.PrintLevelOrderLineByLine2()
	t.PrintSpiralTree()

	t = new(Tree)
	t.Add(2)
	t.Add(1)
	t.Add(3)
	t.Add(4)
	t.PrintInOrder()
	t.PrintPreOrder()
	t.PrintPostOrder()
	t.PrintBreadthFirst()
	t.NthPreOrder(2)
	t.NthPostOrder(2)
	t.NthInOrder(2)
	fmt.Println("Print All Path")
	t.PrintAllPath()
	fmt.Println(t.Find(10))
	fmt.Println(t.Find(3))
	fmt.Println(t.FindMax())
	fmt.Println(t.FindMaxNode())
	fmt.Println(t.FindMin())
	fmt.Println(t.FindMinNode())
	t.Free()
	t.PrintInOrder()	
}