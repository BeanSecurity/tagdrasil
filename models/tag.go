package models

import (
	// "fmt"
)

type TagNode struct {
	ID        int64
	Name      string
	ChildTags []TagNode
}

type User struct {
	ID           int
	FirstName    string
	LastName     string // optional
	UserName     string // optional
	LanguageCode string // optional
}

// func (n *TagNode) GetTagByName(name string) TagNode {
// 	searchStack := Stack{n}

// 	for searchStack.Len() > 0 {

// 		// fmt.Printf("%+v\n", searchStack) // output for debug
// 		tagptr := searchStack.Pop()
// 		// fmt.Printf("%s\n", (*tagptr).Name)
// 		// fmt.Printf("%+v\n", *tagptr) // output for debug
// 		if tagptr.Name == name {
// 			return *tagptr
// 		}
// 		// fmt.Println("CHILDS:")
// 		for i, _ := range tagptr.ChildTags {
// 			// fmt.Printf("%+v\n", tagptr.ChildTags[i].Name) // output for debug
// 			// fmt.Printf("%p\n", &tagptr.ChildTags[i])      // output for debug
// 			searchStack.Push(&tagptr.ChildTags[i])
// 		}
// 		// for _, child := range tagptr.ChildTags {
// 		// 	// fmt.Printf("%+v\n", child.Name) // output for debug
// 		// 	// fmt.Printf("%p\n", &child) // output for debug
// 		// 	searchStack.Push(&child)
// 		// }
// 		// fmt.Println(":CHILDS")
// 	}
// 	return TagNode{}
// }

// func (n *TagNode) GetTagLine(tag TagNode) TagLine {
// 	searchStack := Stack{n}
// 	// var lastTagptr TagNode
// 	depth := 0
// 	hashDepth := map[*TagNode]int{n: 0}
// 	var tagDepthStack Stack
// 	// tagDepthStack.Push(n)

// 	for searchStack.Len() > 0 {
// 		tagptr := searchStack.Pop()
// 		fmt.Printf("%v\n", tagDepthStack) // output for debug

// 		fmt.Printf("tag: %v\n", (*tagptr).Name)          // output for debug
// 		fmt.Printf("Depth: %d\n", depth)                 // output for debug
// 		fmt.Printf("hashDepth: %d\n", hashDepth[tagptr]) // output for debug
// 		fmt.Printf("for: %d\n", depth-hashDepth[tagptr]) // output for debug

// 		// fmt.Printf("\n") // output for debug
// 		for i := 0; i < depth-hashDepth[tagptr]; i++ {
// 			// fmt.Printf("Depth: %d\n", depth) // output for debug
// 			// fmt.Printf("%+v\n", *tagptr)     // output for debug

// 			fmt.Printf("%v\n", i) // output for debug
// 			_ = tagDepthStack.Pop()
// 		}
// 		depth = hashDepth[tagptr]

// 		tagDepthStack.Push(tagptr)
// 		depth++
// 		// if hashDepth[tagptr]-depth > 1
// 		if tagptr.ID == tag.ID {
// 			// fmt.Printf("id: %d\n", tag.ID) // output for debug
// 			// lastTagptr = tagptr
// 			// return *tagptr
// 			var returningTagLine TagLine
// 			for i, _ := range tagDepthStack {
// 				// fmt.Printf("%v\n", *tagDepthStack[i]) // output for debug
// 				returningTagLine = append(returningTagLine, *tagDepthStack[i])
// 			}
// 			fmt.Printf("ruturning line: %v\n", returningTagLine) // output for debug

// 			return returningTagLine
// 		}
// 		for i, _ := range tagptr.ChildTags {

// 			searchStack.Push(&tagptr.ChildTags[i])
// 			hashDepth[&tagptr.ChildTags[i]] = depth
// 			fmt.Printf("hashDepth: %d\n", hashDepth[tagptr]) // output for debug
// 			fmt.Printf("childDepth: %d\n", depth)            // output for debug
// 		}

// 		fmt.Printf("\n") // output for debug
// 	}
// 	return TagLine{}
// }

// type Stack []*TagNode

// func (q *Stack) Push(n *TagNode) {
// 	*q = append(*q, n)
// }

// func (q *Stack) Pop() (n *TagNode) {
// 	x := q.Len() - 1
// 	n = (*q)[x]
// 	*q = (*q)[:x]
// 	return
// }
// func (q *Stack) Len() int {
// 	return len(*q)
// }
