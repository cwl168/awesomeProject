package leak

import (
	"github.com/fortytw2/leaktest"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
	"testing"
	"time"
)

//go test -v
//go test -run TestTest3 -v

// 自动化测试leak  检测goroutine内存泄露
func TestTest3(t *testing.T) {
	defer leaktest.Check(t)()
	Test3()
}

func TestPool(t *testing.T) {
	defer leaktest.Check(t)()

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}
func TestLeak1(t *testing.T) {
	defer leaktest.Check(t)()
	Leak1()
}
func TestLeak2(t *testing.T) {
	defer leaktest.Check(t)()
	Leak2()
}
func TestLeak3(t *testing.T) {
	defer leaktest.Check(t)()
	Leak3()
}
func TestLeak4(t *testing.T) {
	defer leaktest.Check(t)()
	Leak4()

	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

}

func TestLeak5(t *testing.T) {
	defer leaktest.Check(t)()
	Leak5()
}
