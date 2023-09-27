package conditionParser

import (
	"bufio"
	"fmt"
	"os"
)

type Parser struct {
	Reader *bufio.Reader
}

type Number struct {
	Digits string
}

func (n Number) String() string {
	return fmt.Sprintf("%s", n.Digits)
}

type Sum struct {
	A, B interface{}
}

func (s Sum) String() string {
	return fmt.Sprintf("%s+%s", s.A, s.B)
}

func (p Parser) Numeric() interface{} {

	var digits []byte

	for {
		c, _ := p.Reader.Peek(1)

		switch c[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			p.Reader.Discard(1)

			digits = append(digits, c[0])

		case '+':
			p.Reader.Discard(1)

			e := p.Expression(nil)

			return Sum{
				A: Number{string(digits)},
				B: e,
			}

		default:
			return Number{string(digits)}
		}
	}
}

type Paren struct {
	X interface{}
}

func (p Paren) String() string {
	return fmt.Sprintf("(%s)", p.X)
}

func (p Parser) Paren() interface{} {
	e := p.Expression(nil)

	c, _ := p.Reader.Peek(1)

	if c[0] != ')' {
		return nil
	}

	return Paren{e}
}

func (p Parser) Expression(prev interface{}) interface{} {
	c, _ := p.Reader.Peek(1)

	switch c[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return p.Numeric()
	case '(':
		p.Reader.Discard(1)
		return p.Paren()
	case '+':
		if prev == nil {
			return nil
		}

		p.Reader.Discard(1)

		e := p.Expression(nil)
		return Sum{
			A: prev,
			B: e,
		}

	default:
		return nil
	}
}

func RunPars() {
	reader := bufio.NewReader(os.Stdin)

	parser := Parser{reader}

	e := parser.Expression(nil)

	fmt.Printf("decoded: %#v\n", e)
	fmt.Printf("printed: %s\n", e)
}
