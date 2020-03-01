package slice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
)

var nums = []int{10, 9, 3, 4, 1, 7, 2, 8, 5, 6}

func TestKoazee(t *testing.T) {
	assert := assert.New(t)

	stream := koazee.StreamOf(nums)
	fmt.Printf("stream: %v\n", stream.Out().Val())

	assert.Equal(10, stream.First().Int())
	assert.Equal(1, stream.At(4).Int())
	assert.Equal(6, stream.Last().Int())

	count, _ := stream.Count()
	assert.Equal(10, count)

	contains, _ := stream.Contains(6)
	assert.Equal(true, contains)

	index, _ := stream.IndexOf(3)
	assert.Equal(2, index)

	assert.Equal(
		[]int{6, 5, 8, 2, 7, 1, 4, 3, 9, 10},
		stream.Reverse().Out().Val())

	assert.Equal(
		[]int{10, 4, 2, 8, 6},
		stream.Filter(
			func(val int) bool {
				return val%2 == 0
			}).
			Out().Val())

	assert.Equal(
		[]int{20, 18, 6, 8, 2, 14, 4, 16, 10, 12},
		stream.Map(
			func(val int) int {
				return val * 2
			}).
			Out().Val())
}
