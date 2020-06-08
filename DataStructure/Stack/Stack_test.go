package Stack

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Push(1, 2, 3)

	value, err := s.Pop()
	assert.Nil(t, err)
	assert.Equal(t, value, 3)
}

func TestStack_Peek(t *testing.T) {
	s := NewStack(123, true)

	val, _ := s.Peek()
	assert.Equal(t, val, true)
}

func TestStack_Size(t *testing.T) {
	s := NewStack()
	assert.Equal(t, s.Size(), 0)

	s.Push(1, 2, 3)
	assert.Equal(t, s.Size(), 3)
}

func TestStack_IsEmpty(t *testing.T) {
	s := NewStack()
	assert.Equal(t, s.IsEmpty(), true)

	s.Push(1, 2, 3)
	assert.Equal(t, s.IsEmpty(), false)
}

func TestStack_GetValues(t *testing.T) {
	s := NewStack()
	s.Push(1, "abc", true)
	assert.True(t, reflect.DeepEqual(s.GetValues(), []interface{}{1, "abc", true}))
}

func TestStack_Clear(t *testing.T) {
	s := NewStack(1, 2, 3)
	s.Clear()
	assert.Equal(t, s.Size(), 0)
	assert.Equal(t, s.IsEmpty(), true)
}
