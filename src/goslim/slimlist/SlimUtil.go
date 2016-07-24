package slimlist

import ()

func CSlim_IsCharacter(input byte) bool {
	if (input < 0x80) || (input > 0xBF) {
		return true
	} else {
		return false
	}
}
