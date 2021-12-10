// Resource: https://gist.github.com/derianpt/93cbf64a759f037ad8bcf8c52664ce79
package aochelper

import "errors"

type Stack []interface{}

func (s *Stack) Push(element interface{}) {
  *s = append(*s, element)
}

func (s *Stack) Pop() (interface{}, error) {
  if len(*s) > 0 {
    popped := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return popped, nil
  }
  return -1, errors.New("stack is empty")
}

func (s *Stack) Peek() (interface{}, error) {
  if len(*s) > 0 {
    return (*s)[len(*s)-1], nil
  }
  return -1, errors.New("stack is empty")
}
