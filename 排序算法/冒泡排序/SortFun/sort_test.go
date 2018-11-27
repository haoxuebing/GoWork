package SortFun

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	values := []int{7, 5, 3, 9, 8}
	BubbleSort(values)

	if values[0] != 3 && values[4] != 9 {
		t.Error("冒泡排序有误")
	}
}

func TestQuckSort(t *testing.T) {

	values := []int{7, 5, 3, 9, 8}
	QuckSort(values, 0, len(values)-1)

	if values[0] != 3 || values[4] != 9 {
		t.Error("快速排序有误")
	}
}
