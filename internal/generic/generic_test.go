package generic_test

import (
	"testing"
	"github.com/Gusarov2k/generics/internal/generic"
	"github.com/stretchr/testify/assert"
)

func TestGenerateStringInt(t *testing.T) {

	assert.EqualValues(t, "1, 2, 3, 4, 5", generic.GenerateString(1,2,3,4,5) )
}

func TestGenerateStringInt32(t *testing.T) {

	assert.EqualValues(t, "1, 2, 3, 4, 5", generic.GenerateString32(1,2,3,4,5) )
}

func TestGenerateStringInterfaceInt(t *testing.T) {

	assert.EqualValues(t, "1, 2, 3, 4, 5", generic.GenerateStringInter(1,2,3,4,5) )
}

func TestGenerateStringIntInterfaceInt32(t *testing.T) {
	var num int32 = 112
	assert.EqualValues(t, "112", generic.GenerateStringInter(num) )
}
