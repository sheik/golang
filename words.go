package main

import (
	"fmt"
	"strings"
)

type Prefix []string

func (p Prefix) String() string {
	return strings.Join(p, " ")
}

func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) String() string {
	return fmt.Sprintf("Prefix Length: %d\nChain:\n%+v", c.prefixLen, c.chain)
}

func main() {
	fmt.Println(NewChain(2))
}
