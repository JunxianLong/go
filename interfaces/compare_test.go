package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	var dog1, dog2 Sayer = &Dog{"Tom"}, &Dog{"Tom"}
	var dog3, dog4 Sayer = Dog{"Tom"}, Dog{"Tom"}
	assert.False(t, dog1 == dog2, "引用地址不同")
	assert.True(t, dog3 == dog4, "结构体值相同")
}
