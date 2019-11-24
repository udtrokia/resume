package flash

import "strings"

type Pair struct {
	car int
	cdr int
}

type Ident struct {
	Type     string
	TypeArgs []Ident
}

func (i *Ident) push(_i Ident) {
	i.TypeArgs = append(i.TypeArgs, _i)
}

const (
	LESS    = '<'
	COMMA   = ','
	GREATER = '>'
)

func Parser(raw string) Ident {
	ptr := 0
	stack := Ident{}

	errIdent := Ident{Type: "Error"}
	raw = strings.ReplaceAll(raw, " ", "")

	var gStart, gEnd Pair
	for i, str := range raw {
		if str == LESS {
			if gStart.car == 0 {
				stack.Type = raw[ptr:i]
				ptr = i + 1
			}

			//break for inner-loop
			if gStart.car > 0 {
				idx := strings.IndexByte(raw[i:], GREATER) + len(raw[:i]) + 1

				if idx == -1 {
					return errIdent
				}

				stack.push(Parser(raw[ptr:idx]))

				if raw[idx:] != strings.Repeat(string(GREATER), len(raw[idx:])) {
					stack.push(Parser(raw[idx:]))
				}

				break
			}

			gStart.car += 1
			gStart.cdr = i
			continue
		}

		if str == COMMA {
			stack.push(Parser(raw[ptr:i]))
			ptr = i + 1
			continue
		}

		if str == GREATER {
			gEnd.car += 1
			gEnd.cdr = i

			if gStart.car == gEnd.car {
				stack.push(Parser(raw[ptr:i]))
			}
		}
	}

	if gStart.car == 0 && len(raw[ptr:len(raw)]) > 1 {
		if gEnd.car > gStart.car {
			stack = Ident{Type: raw[ptr : len(raw)-1]}
		} else {
			stack = Ident{Type: raw[ptr:]}
		}
	}

	return stack
}
