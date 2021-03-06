package main

import "fmt"
import "github.com/golang-collections/collections/queue"
import "github.com/golang-collections/collections/stack"
import "github.com/golang-collections/collections/set"

func main1() {
	st := set.New()
	st.Insert(1)
	st.Insert(1)
	st.Insert(2)
	fmt.Println(st.Has(1))
	fmt.Println(st.Has(3))
	st.Remove(1)
	fmt.Println(st.Has(1))
}

func main2() {
	s := stack.New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for s.Len() != 0 {
		val := s.Pop()
		fmt.Print(val, " ")
	}
}


func main3() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for q.Len() != 0 {
		val := q.Dequeue()
		fmt.Print(val, " ")
	}
}



func main4() {

	m := make(map[string]int)
	m["first"] = 1
	m["second"] = 2
	fmt.Println(m["first"])
}

func main() {
	main1()
	main2()
	main3()
	main4()
}
