package txtreader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLinesFromTextFile(t *testing.T) {
	t.Run("returns lines slice from text file", func(t *testing.T) {
		lines, _ := GetLinesFromTextFile("testdata/test.txt")

		assert.Len(t, lines, 3)
		assert.Equal(t, "Hello", lines[0])
		assert.Equal(t, "World", lines[len(lines)-1])
	})

	t.Run("returns err if file was not found", func(t *testing.T) {
		_, err := GetLinesFromTextFile("not_existing_file")

		assert.Error(t, err)
	})
}
