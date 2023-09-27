package myParser

// (123+3)
import (
	"Project/writeToFile"
	"bufio"

	"strconv"
)

type Parser struct {
	Reader *bufio.Reader
}

func (p Parser) FinderExpression() interface{} {
	c, _ := p.Reader.Peek(1)
	var digits []byte
	switch c[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		p.Reader.Discard(1)

		digits = append(digits, c[0])
	case '+':
		p.Reader.Discard(1)

		intDigits, _ := strconv.Atoi(string(digits))

		return
	}
	digits = 0
}

func RunMyParser() uint64 {
	return Age()
}
