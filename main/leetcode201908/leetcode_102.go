package main

func main() {

}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	num := [][]*TreeNode{{root}}
	re := [][]int{{root.Val}}
	for i := 0; i < len(num); i++ {
		var temp []*TreeNode
		var numTemp []int
		for j := 0; j < len(num[i]); j++ {
			if num[i][j].Left != nil {
				temp = append(temp, num[i][j].Left)
				numTemp = append(numTemp, num[i][j].Left.Val)
			}
			if num[i][j].Right != nil {
				temp = append(temp, num[i][j].Right)
				numTemp = append(numTemp, num[i][j].Right.Val)
			}
		}
		if temp != nil {
			num = append(num, temp)
			re = append(re, numTemp)
		}
	}
	return re
}
