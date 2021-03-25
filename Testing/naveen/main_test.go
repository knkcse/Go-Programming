package naveen

import (
	"testing"
)

func TestmySum(t *testing.T) {
	sum := mySum(5, 5, 4)
	if sum != 14 {
		t.Error("Expecting ", 14, "But got ", sum)
	}

}
func TestsumTwo(t *testing.T){
	sum:=sumTwo(5,7)
	if sum!=14{
		t.Error("Expected 12 and got ",sum)
	}
}