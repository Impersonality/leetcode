package dfs

import (
	"fmt"
	"strconv"
	"strings"
)

/*
思路：惭愧，没有认真写题。这题不难，考察的是coding能力
答案：先根遍历树，然后拼接成string，string split成数组，然后递归成树
*/

type Codec struct {
	builder strings.Builder
	input   []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.builder.WriteString("#,")
		return ""
	}
	this.builder.WriteString(fmt.Sprintf("%d,", root.Val))
	this.serialize(root.Left)
	this.serialize(root.Right)
	return this.builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.input = strings.Split(data, ",")
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if this.input[0] == "#" {
		this.input = this.input[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.input[0])
	this.input = this.input[1:]
	return &TreeNode{
		Val:   val,
		Left:  this.deserializeHelper(),
		Right: this.deserializeHelper(),
	}
}
