package util

import (
	"github.com/tongsq/go-lib/logger"
	"math/big"
	"strconv"
)

func Str2dec(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << uint8(i)
	}
	return
}

func Str2IntBase16(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) & 0xf) << (l - i) * 4
	}
	return
}

func Int2str(i int) string {
	return strconv.Itoa(i)
}

func Str2Int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Error("parse string to int fail", err)
		return 0
	}
	return i
}

func Int64toStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int2byte(i int) byte {
	return uint8(i)
}

func StringTo16BaseByte(s string) []byte {
	var str16 []byte
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			n := i + 2
			x, _ := strconv.ParseUint(s[i:n], 16, 32)
			str16 = append(str16, byte(x))
		}
	}
	return str16
}

func FromBase16(base16 string) *big.Int {
	i, ok := new(big.Int).SetString(base16, 16)
	if !ok {
		panic("bad number: " + base16)
	}
	return i
}

func Add(args ...int) (sum int) {
	for _, v := range args {
		sum = sum + v
	}
	return
}

func Max(args ...int) (m int) {
	for i, v := range args {
		if i == 0 || v > m {
			m = v
		}
	}
	return
}

func Min(args ...int) (m int) {
	for i, v := range args {
		if i == 0 || v < m {
			m = v
		}
	}
	return
}
