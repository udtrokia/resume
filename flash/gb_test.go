package flash_test

import (
	"github.com/find-a-job/flash"
	"github.com/stretchr/testify/assert"
	"testing"
)

func gbGen() []flash.Student {
	return []flash.Student{
		flash.Student{Name: "张三", Score: 84},
		flash.Student{Name: "李四", Score: 58},
		flash.Student{Name: "王五", Score: 99},
		flash.Student{Name: "赵六", Score: 69},
	}
}

func TestGb(t *testing.T) {
	ss := gbGen()
	assert.Equal(
		t, flash.GroupBy(ss), flash.Group{
			A: []flash.Student{
				flash.Student{Name: "王五", Score: 99},
				flash.Student{Name: "张三", Score: 84},
			},
			B: []flash.Student{
				flash.Student{Name: "赵六", Score: 69},
			},
			C: []flash.Student{
				flash.Student{Name: "李四", Score: 58},
			},
		},
		"4_学生按成绩分组 groupBy went Error",
	)
}
