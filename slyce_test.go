package slyce_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/projekt-go/slyce"
)

func TestMap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	output := slyce.Map(s, func(elem int) int {
		return elem * elem
	})
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("invalid map output. expected %v got %v\n", expected, output)
	}

	expected2 := []string{"1", "4", "9", "16", "25", "36", "49", "64", "81", "100"}
	output2 := slyce.Map(s, func(elem int) string {
		return strconv.Itoa(elem * elem)
	})
	if !reflect.DeepEqual(output2, expected2) {
		t.Errorf("invalid map output. expected %v got %v\n", expected2, output2)
	}
}

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}

	expected := []int{2, 4, 6}
	output := slyce.Filter(s, func(elem int) bool {
		return elem % 2 == 0
	})
	ok := true
	for _, v := range output {
		if v % 2 != 0 {
			ok = false
			break
		}
	}
	if !ok {
		t.Errorf("invalid filter output. expected %v got %v\n", expected, output)
	}
}

func TestCompose(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	output := slyce.Map(slyce.Filter(s, func(x int) bool {
		return x % 2 == 0
	}), func(x int) string {
		return strconv.Itoa(x)
	})

	for _, v := range output {
		if x, _ := strconv.Atoi(v); x % 2 != 0 {
			t.Errorf("invalid composition output. expected no odd values but found %d", x)
		}
	}
}
