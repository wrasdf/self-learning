package demo_test

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  assert.True(t, true, "True is true!")
  assert.Equal(t, 123, 123, "they should be equal")

}
