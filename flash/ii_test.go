package flash_test

import (
	"github.com/find-a-job/flash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func iiGen() flash.IncTreeNode {
	return flash.IncTreeNode{
		Id:   "1",
		Type: "View",
		Name: "view",
		Children: []flash.IncTreeNode{
			flash.IncTreeNode{
				Id:   "2",
				Type: "Button",
				Name: "button",
			},
			flash.IncTreeNode{
				Id:   "3",
				Type: "View",
				Name: "view_1",
				Children: []flash.IncTreeNode{
					flash.IncTreeNode{
						Id:   "4",
						Type: "Button",
						Name: "button_1",
					},
					flash.IncTreeNode{
						Id:   "5",
						Type: "View",
						Name: "view_2",
					},
				},
			},
		},
	}
}

func iiGenSecond() flash.IncTreeNode {
	return flash.IncTreeNode{
		Id:   "1",
		Type: "View",
		Name: "view",
		Children: []flash.IncTreeNode{
			flash.IncTreeNode{
				Id:   "2",
				Type: "Button",
				Name: "button",
			},
			flash.IncTreeNode{
				Id:   "3",
				Type: "View",
				Name: "view_1",
				Children: []flash.IncTreeNode{
					flash.IncTreeNode{
						Id:   "4",
						Type: "Button",
						Name: "button_1",
					},
					flash.IncTreeNode{
						Id:   "5",
						Type: "View",
						Name: "view_3",
					},
				},
			},
		},
	}
}

func TestIncId(t *testing.T) {
	iiError := "7_自增ID went Error"
	itn := iiGen()
	itn2 := iiGenSecond()

	itn.Insert(flash.IncTreeNode{
		Id:   "6",
		Type: "View",
		Name: "view",
	})

	itn2.Insert(flash.IncTreeNode{
		Id:   "6",
		Type: "View",
		Name: "view",
	})

	assert.Equal(t, itn, flash.IncTreeNode{
		Id:   "1",
		Type: "View",
		Name: "view",
		Children: []flash.IncTreeNode{
			flash.IncTreeNode{
				Id:   "2",
				Type: "Button",
				Name: "button",
			},
			flash.IncTreeNode{
				Id:   "3",
				Type: "View",
				Name: "view_1",
				Children: []flash.IncTreeNode{
					flash.IncTreeNode{
						Id:   "4",
						Type: "Button",
						Name: "button_1",
					},
					flash.IncTreeNode{
						Id:   "5",
						Type: "View",
						Name: "view_2",
					},
				},
			},
			flash.IncTreeNode{
				Id:   "6",
				Type: "View",
				Name: "view_3",
			},
		},
	}, iiError)

	assert.Equal(t, itn2, flash.IncTreeNode{
		Id:   "1",
		Type: "View",
		Name: "view",
		Children: []flash.IncTreeNode{
			flash.IncTreeNode{
				Id:   "2",
				Type: "Button",
				Name: "button",
			},
			flash.IncTreeNode{
				Id:   "3",
				Type: "View",
				Name: "view_1",
				Children: []flash.IncTreeNode{
					flash.IncTreeNode{
						Id:   "4",
						Type: "Button",
						Name: "button_1",
					},
					flash.IncTreeNode{
						Id:   "5",
						Type: "View",
						Name: "view_3",
					},
				},
			},
			flash.IncTreeNode{
				Id:   "6",
				Type: "View",
				Name: "view_2",
			},
		},
	}, iiError)
}
