package gosteam

import "strconv"

func (sid *SteamID) ToString() string {
	return strconv.FormatUint(uint64(*sid), 10)
}
