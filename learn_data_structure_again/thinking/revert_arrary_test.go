package thinking

import "testing"

func TestGetIndex(t *testing.T) {
	arr := []int{
		4, 5, 6, 7, 0, 1, 2,
	}
	target := 7

	result := getIndex(arr, target, 0, len(arr)-1)
	if result != 3 {
		t.Errorf("should be 3，but get %d", result)
	}
}
