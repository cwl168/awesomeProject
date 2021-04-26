package bool

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewPool(t *testing.T) {
	Test()
}

func TestDuiqi(t *testing.T) {
	fmt.Printf("%d\n", unsafe.Sizeof(struct {
		a int8
		b int64
		c int32
		d int16
	}{}))

	fmt.Printf("%d\n", unsafe.Sizeof(struct {
		a int8
		b int64
		c int16
	}{}))
	fmt.Printf("%d\n", unsafe.Sizeof(struct {
		a int8
		c int16
		b int64
	}{}))
}
