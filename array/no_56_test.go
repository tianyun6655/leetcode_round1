package array

import "testing"

func TestTianyun_merge(t *testing.T) {

	inter:=[]Interval{{2,6},{1,3,},{8,10},{15,18}}

	Tianyun_merge(inter)
}
