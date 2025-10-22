package other

// import (
// 	"fmt"
// 	"math/big"
// 	"slices"
// 	"strconv"
// 	"strings"
// )

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	num1 := convert(l1)
// 	num2 := convert(l2)
// 	sum := new(big.Int).Add(num1, num2)
// 	strRes := sum.String()
// 	var res []string = strings.Split(strRes, "")
// 	slices.Reverse(res)
// 	return createNode(res)
// }

// func convert(l *ListNode) *big.Int {
// 	arr := make([]string, 0)
// 	for {
// 		arr = append(arr, strconv.Itoa(l.Val))
// 		if l.Next != nil {
// 			l = l.Next
// 		} else {
// 			break
// 		}
// 	}
// 	slices.Reverse(arr)
// 	n := new(big.Int)
// 	n.SetString(strings.Join(arr, ""), 10)
// 	return n
// }

// func createNode(res []string) *ListNode {
// 	list := &ListNode{}
// 	copy := list
// 	for i := range res {
// 		list.Val, _ = strconv.Atoi(res[i])
// 		if i == len(res)-1 {
// 			break
// 		}
// 		list.Next = &ListNode{}
// 		list = list.Next
// 	}
// 	fmt.Println(*copy)
// 	return copy
// }
