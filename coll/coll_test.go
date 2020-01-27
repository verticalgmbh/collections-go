package coll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyTrue(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Any(numbers, func(item interface{}) bool {
		number := item.(int)
		return number == 5
	})

	assert.True(t, result)
}

func TestAnyFalse(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Any(numbers, func(item interface{}) bool {
		number := item.(int)
		return number == 17
	})

	assert.False(t, result)
}

func TestAnySingleItem(t *testing.T) {
	result := Any(17, func(item interface{}) bool {
		number := item.(int)
		return number == 17
	})

	assert.True(t, result)
}

func TestAllTrue(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := All(numbers, func(item interface{}) bool {
		number := item.(int)
		return number < 10
	})

	assert.True(t, result)
}

func TestAllFalse(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := All(numbers, func(item interface{}) bool {
		number := item.(int)
		return number < 9
	})

	assert.False(t, result)
}

func TestAllSingleItem(t *testing.T) {
	result := All(5, func(item interface{}) bool {
		number := item.(int)
		return number < 10
	})

	assert.True(t, result)
}

func TestInterfaceConvert(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	converted := Interface(numbers)

	assert.Equal(t, len(numbers), len(converted))
	for index, item := range converted {
		expected := numbers[index]
		assert.Equal(t, expected, item)
	}
}

func TestInterfaceConvertSingle(t *testing.T) {
	converted := Interface(346)

	assert.Equal(t, 1, len(converted))
	assert.Equal(t, 346, converted[0])
}

func TestAddToWhere(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result []int

	AddToWhere(numbers, func(item interface{}) bool {
		number := item.(int)
		return number < 5
	}, &result)

	assert.Equal(t, 4, len(result))
	assert.Equal(t, 1, result[0])
	assert.Equal(t, 2, result[1])
	assert.Equal(t, 3, result[2])
	assert.Equal(t, 4, result[3])
}

func TestAddToWhereSingleItem(t *testing.T) {
	var result []int

	AddToWhere(2, func(item interface{}) bool {
		number := item.(int)
		return number < 5
	}, &result)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, 2, result[0])
}

func TestWhere(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := Where(numbers, func(item interface{}) bool {
		number := item.(int)
		return number < 5
	})

	assert.Equal(t, 4, len(result))
	assert.Equal(t, 1, result[0])
	assert.Equal(t, 2, result[1])
	assert.Equal(t, 3, result[2])
	assert.Equal(t, 4, result[3])
}

func TestWhereSingleItem(t *testing.T) {
	result := Where(2, func(item interface{}) bool {
		number := item.(int)
		return number < 5
	})

	assert.Equal(t, 1, len(result))
	assert.Equal(t, 2, result[0])
}

func TestDoWhere(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result []interface{}

	DoWhere(numbers, func(item interface{}) bool {
		number := item.(int)
		return number < 5
	}, func(item interface{}) {
		result = append(result, item)
	})

	assert.Equal(t, 4, len(result))
	assert.Equal(t, 1, result[0])
	assert.Equal(t, 2, result[1])
	assert.Equal(t, 3, result[2])
	assert.Equal(t, 4, result[3])
}

func TestDoWhereSingleItem(t *testing.T) {
	var result []interface{}
	DoWhere(2, func(item interface{}) bool {
		number := item.(int)
		return number < 5
	}, func(item interface{}) {
		result = append(result, item)
	})

	assert.Equal(t, 1, len(result))
	assert.Equal(t, 2, result[0])
}

func TestFirstOrDefault(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := FirstOrDefault(numbers, func(item interface{}) bool {
		number := item.(int)
		return number > 5
	})

	assert.Equal(t, 6, result)
}

func TestFirstOrDefaultNoItemMatches(t *testing.T) {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	result := FirstOrDefault(numbers, func(item interface{}) bool {
		number := item.(int)
		return number > 10
	})

	assert.Nil(t, result)
}

func TestFirstOrDefaultSingleItem(t *testing.T) {
	result := FirstOrDefault(6, func(item interface{}) bool {
		number := item.(int)
		return number > 5
	})

	assert.Equal(t, 6, result)
}

func TestFirstOrDefaultSingleItemNoMatch(t *testing.T) {
	result := FirstOrDefault(6, func(item interface{}) bool {
		number := item.(int)
		return number > 10
	})

	assert.Nil(t, result)
}
