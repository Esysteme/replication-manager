// generated by stringer -type=KDEBoundaryMethod; DO NOT EDIT

package stats

import "fmt"

const _KDEBoundaryMethod_name = "BoundaryReflect"

var _KDEBoundaryMethod_index = [...]uint8{0, 15}

func (i KDEBoundaryMethod) String() string {
	if i < 0 || i+1 >= KDEBoundaryMethod(len(_KDEBoundaryMethod_index)) {
		return fmt.Sprintf("KDEBoundaryMethod(%d)", i)
	}
	return _KDEBoundaryMethod_name[_KDEBoundaryMethod_index[i]:_KDEBoundaryMethod_index[i+1]]
}
