package utils

import "strconv"

type UtilsInt struct {
}

func (u UtilsInt) Uint64ToString(number uint64) string {
	return strconv.Itoa(int(number))
}

func (u UtilsInt) Uint32ToString(number uint32) string {
	return strconv.Itoa(int(number))
}

func (u UtilsInt) StringToInt64(number string) int64 {
	value, err := strconv.ParseInt(number, 10, 64)

	if err != nil {
		return 0
	}

	return value
}
