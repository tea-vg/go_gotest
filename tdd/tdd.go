package tdd

import (
	"errors"
	"fmt"
)

type ClosedSectionError int

const (
  _ ClosedSectionError = iota
  RangeOrderError
)

type closedSection struct{
  left int
  right int
}

type ClosedSection interface {
  GetLeft() int
  GetRight() int
  Equal(cs_ ClosedSection) bool
  ContainNum(int) bool
  ContainSection(cs_ ClosedSection) bool
  ToString() string
}

func NewClosedSection(left, right int) (ClosedSection, error) {
  if right < left {
    return nil, errors.New("left <= right is required.")
  }

  return &closedSection{left: left, right: right}, nil
}

func (cs *closedSection) GetLeft() int {
  return cs.left
}

func (cs *closedSection) GetRight() int {
  return cs.right
}

func(cs *closedSection) Equal(cs_ ClosedSection) bool {
  return cs.GetLeft() == cs_.GetLeft() && cs.GetRight() == cs_.GetRight()
}

func (cs *closedSection) ContainNum(n int) bool {
  return cs.GetLeft() <= n && n <= cs.GetRight()
}

func (cs *closedSection) ContainSection(cs_ ClosedSection) bool {
  return cs.GetLeft() <= cs_.GetLeft() && cs_.GetRight() <= cs.GetRight()
}

func (cs *closedSection) ToString() string {
  // return "[0, 1]"

  return fmt.Sprintf("[%d, %d]", cs.GetLeft(), cs.GetRight())
}
