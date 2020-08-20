package utlRandom

import (
	"strconv"
)

type AutoIncrement struct {
	Current int
}

func NewAutoIncrementId() *AutoIncrement {
	return &AutoIncrement{0}
}

func (m *AutoIncrement) GetNextId() int {
	m.Current++
	return m.Current
}

func (m *AutoIncrement) GetNextStringId() string {
	m.Current++
	return strconv.Itoa(m.Current)
}
