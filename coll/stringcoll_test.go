package coll

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToWhereString(t *testing.T) {
	items := []string{"abc", "abs", "aaa", "arte", "bums", "ball", "bimmel", "bamsch", "bartender"}
	var result []string

	AddToWhereString(items, func(item string) bool {
		return len(item) == 4
	}, &result)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, "arte", result[0])
	assert.Equal(t, "bums", result[1])
	assert.Equal(t, "ball", result[2])
}

func TestWhereString(t *testing.T) {
	items := []string{"abc", "abs", "aaa", "arte", "bums", "ball", "bimmel", "bamsch", "bartender"}

	result := WhereString(items, func(item string) bool {
		return len(item) == 4
	})

	assert.Equal(t, 3, len(result))
	assert.Equal(t, "arte", result[0])
	assert.Equal(t, "bums", result[1])
	assert.Equal(t, "ball", result[2])
}

func TestDoWhereString(t *testing.T) {
	items := []string{"abc", "abs", "aaa", "arte", "bums", "ball", "bimmel", "bamsch", "bartender"}
	var result []string

	DoWhereString(items, func(item string) bool {
		return len(item) == 4
	}, func(item string) {
		result = append(result, item)
	})

	assert.Equal(t, 3, len(result))
	assert.Equal(t, "arte", result[0])
	assert.Equal(t, "bums", result[1])
	assert.Equal(t, "ball", result[2])
}

func TestAnyStringTrue(t *testing.T) {
	names := []string{"abc", "abs", "aaa", "arte", "bums", "ball", "bimmel", "bamsch", "bartender"}

	result := AnyString(names, func(item string) bool {
		return strings.HasPrefix(item, "b")
	})

	assert.True(t, result)
}

func TestAnyStringFalse(t *testing.T) {
	names := []string{"abc", "abs", "aaa", "arte", "bums", "ball", "bimmel", "bamsch", "bartender"}

	result := AnyString(names, func(item string) bool {
		return strings.HasPrefix(item, "t")
	})

	assert.False(t, result)
}

func TestAllStringTrue(t *testing.T) {
	names := []string{"abc", "abs", "aaa", "arte", "bumsa", "ball", "bammel", "bamsch", "bartender"}

	result := AllString(names, func(item string) bool {
		return strings.Contains(item, "a")
	})

	assert.True(t, result)
}

func TestAllStringFalse(t *testing.T) {
	names := []string{"abc", "abs", "aaa", "arte", "bums", "ball", "bimmel", "bamsch", "bartender"}

	result := AllString(names, func(item string) bool {
		return strings.Contains(item, "a")
	})

	assert.False(t, result)
}
