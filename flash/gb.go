package flash

import "sort"

type Student struct {
	Name  string
	Score int
}

// Sorter
type StudentSorter []Student

func (s StudentSorter) Len() int           { return len(s) }
func (s StudentSorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s StudentSorter) Less(i, j int) bool { return s[i].Score > s[j].Score }

// Group
type Group struct {
	A []Student
	B []Student
	C []Student
}

func GroupBy(ss []Student) Group {
	var a, b, c []Student
	for _, s := range ss {
		if s.Score < 60 {
			c = append(c, s)
			continue
		}

		if s.Score < 80 {
			b = append(b, s)
			continue
		}

		a = append(a, s)
	}

	sort.Sort(StudentSorter(a))
	sort.Sort(StudentSorter(b))
	sort.Sort(StudentSorter(c))

	return Group{A: a, B: b, C: c}
}
