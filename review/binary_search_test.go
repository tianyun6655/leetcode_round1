package review

import (
	"testing"
	"fmt"
)

func TestBinarySearch(t *testing.T) {
	source:=[]int{1,2,3,4,5,6,7,8,9}
	fmt.Println(BinarySearch(source,1))
}
