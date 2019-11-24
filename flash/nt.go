package flash

// 0_TreeNode
type TreeNode struct {
	Id       string
	Label    string
	Children []TreeNode
}

func (t *TreeNode) Push(nodes []TreeNode) {
	var tns []TreeNode
	for i, e := range nodes {
		if i == 0 {
			tns = append(t.Children, e)
			continue
		}

		tns = append(tns, e)
	}

	t.Children = tns
}

func FindNodeById(root TreeNode, id string) TreeNode {
	noRes := TreeNode{Id: "-1", Label: "no result"}

	if root.Id == id {
		return root
	}

	if len(root.Children) == 0 {
		return noRes
	}

	for _, n := range root.Children {
		res := FindNodeById(n, id)
		if res.Id != "-1" {
			return res
		}
	}

	return noRes
}
