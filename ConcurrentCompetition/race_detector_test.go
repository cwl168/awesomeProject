package ConcurrentCompetition

import (
	"fmt"
	"github.com/fortytw2/leaktest"
	"testing"
)

/*func TestAdd(t *testing.T) {
	addCount()
	UpdateMap()
}


func TestSysncMap(t *testing.T) {
	sysncMap()
}
*/
func TestRaceDetector2(t *testing.T) {
	//race_detector2()
	fmt.Println()
	race_detector3()
}
func TestMapHandle(t *testing.T) {
	MapHandle()
}
func TestSliceHandle(t *testing.T) {
	defer leaktest.Check(t)()
	SliceHandle()
}
