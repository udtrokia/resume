package flash_test

import (
	"github.com/find-a-job/flash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ntGen() flash.TreeNode {
	tree := flash.TreeNode{
		Id:    "1",
		Label: "first",
	}

	tree.Push([]flash.TreeNode{
		flash.TreeNode{Id: "2", Label: "second"},
		flash.TreeNode{Id: "3", Label: "third", Children: []flash.TreeNode{
			flash.TreeNode{Id: "4", Label: "fourth"},
			flash.TreeNode{Id: "5", Label: "fifth"},
		}},
	})

	return tree
}

func TestFindNodeById(t *testing.T) {
	tree := ntGen()
	first := tree
	second := flash.TreeNode{Id: "2", Label: "second"}
	third := flash.TreeNode{Id: "3", Label: "third"}
	fourth := flash.TreeNode{Id: "4", Label: "fourth"}
	fifth := flash.TreeNode{Id: "5", Label: "fifth"}

	third.Push([]flash.TreeNode{fourth, fifth})

	ntError := "0_flash.TreeNode查询 went Error"
	assert.Equal(t, flash.FindNodeById(tree, "1"), first, ntError)
	assert.Equal(t, flash.FindNodeById(tree, "2"), second, ntError)
	assert.Equal(t, flash.FindNodeById(tree, "3"), third, ntError)
	assert.Equal(t, flash.FindNodeById(tree, "4"), fourth, ntError)
	assert.Equal(t, flash.FindNodeById(tree, "5"), fifth, ntError)
}
