package demo

import (
	"fmt"
	"math"
	"testing"
)

func TestGetArea(t *testing.T) {
	area := getArea(11, 11)
	if area != 121 {
		fmt.Println("测试失败")
	}
}

func TestNewArgs(t *testing.T) {
	var p interface{}
	p = NewArgs(100.0, 200.0)
	switch p.(type) {
	case *Args:
	default:
		t.Error("测试不通过")
	}
}

func TestArgs_GetAreas(t *testing.T) {
	a := NewArgs(10.0, 10.0)
	areas := a.GetMatrix()
	if math.Abs(areas-100) > 0.00000001 {
		t.Error("测试失败")
	}

}
