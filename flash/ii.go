package flash

import (
	"strconv"
	"strings"
)

// 7_自增ID
type IncTreeNodeSorter []IncTreeNode

type IncTreeNode struct {
	Id       string
	Type     string
	Name     string
	Children []IncTreeNode
}

func (i *IncTreeNode) findNodeByName(name string) IncTreeNode {
	noRes := IncTreeNode{Id: "-1", Type: "Nil", Name: "nil"}

	if i.Name == name {
		return *i
	}

	if len(i.Children) == 0 {
		return noRes
	}

	for _, n := range i.Children {
		res := n.findNodeByName(name)
		if res.Name != "nil" {
			return res
		}
	}

	return noRes
}

func (i *IncTreeNode) Insert(node IncTreeNode) {
	if _n := i.findNodeByName(node.Name); _n.Name != "nil" {
		idx := strings.IndexByte(_n.Name, '_')
		if idx != -1 {
			_id, _ := strconv.Atoi(string(_n.Name[idx+1]))
			node.Name = _n.Name[:idx+1] + strconv.Itoa(_id+1)
			i.Insert(node)
			return
		}

		node.Name = node.Name + "_1"
		i.Insert(node)
		return
	}

	itns := i.Children
	itns = append(i.Children, node)
	i.Children = itns
	return
}
