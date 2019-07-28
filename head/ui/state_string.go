// Code generated by "stringer -type=State -trimprefix=State"; DO NOT EDIT.

package ui

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StateInvalid-0]
	_ = x[StateBoot-1]
	_ = x[StateBroken-2]
	_ = x[StateFrontBegin-3]
	_ = x[StateFrontSelect-4]
	_ = x[StateFrontTune-5]
	_ = x[StateFrontAccept-6]
	_ = x[StateFrontTimeout-7]
	_ = x[StateFrontEnd-8]
	_ = x[StateServiceBegin-9]
	_ = x[StateServiceAuth-10]
	_ = x[StateServiceMenu-11]
	_ = x[StateServiceInventory-12]
	_ = x[StateServiceReboot-13]
	_ = x[StateServiceEnd-14]
}

const _State_name = "InvalidBootBrokenFrontBeginFrontSelectFrontTuneFrontAcceptFrontTimeoutFrontEndServiceBeginServiceAuthServiceMenuServiceInventoryServiceRebootServiceEnd"

var _State_index = [...]uint8{0, 7, 11, 17, 27, 38, 47, 58, 70, 78, 90, 101, 112, 128, 141, 151}

func (i State) String() string {
	if i >= State(len(_State_index)-1) {
		return "State(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}