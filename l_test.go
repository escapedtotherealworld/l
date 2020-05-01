package l

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReflectList(t *testing.T) {
	assert.Panics(t, func() { reflectList("invalid data kind") }, "Successful reflection")
}

func TestExclude(t *testing.T) {
	list := []interface{}{"a", "a", "b", "a", 1, 1, 2, 3, "c"}
	assert.Equal(t, []interface{}{"b", 2, 3, "c"}, Exclude(list, "a", 1), "Incorrect exclusion")
}

func TestInList(t *testing.T) {
	list := []interface{}{"a", "b", "c", 1}
	assert.Equal(t, true, InList(list, 1), "Element not in list")
	assert.Equal(t, false, InList(list, "d"), "Element in list")
}
