package util

import (
	cryptorand "crypto/rand"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterBytesLen = len(letterBytes)
)

/*
************************************************
@ 获取伪随机数, 如： 00347135859b4561
@ size:随机数的长度
@ strType:
@ 1: 数字,大小写字母
@ 2: 纯数字
@ 3: 纯小写字母
@ 4: 纯大写字母
@ 5: 数字+小写字母
@ 6: 数字+ 大写字母
@ 7: 小写+大写字母
@*************************************************
*/
func GetRandStr(strType int, size int) string {
	var kinds [][]int
	switch strType {
	case 1:
		kinds = [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}
	case 2:
		kinds = [][]int{[]int{10, 48}}
	case 3:
		kinds = [][]int{[]int{26, 97}}
	case 4:
		kinds = [][]int{[]int{26, 65}}
	case 5:
		kinds = [][]int{[]int{10, 48}, []int{26, 97}}
	case 6:
		kinds = [][]int{[]int{10, 48}, []int{26, 65}}
	case 7:
		kinds = [][]int{[]int{26, 97}, []int{26, 65}}
	default:
		return ""
	}

	kindsLen := len(kinds)
	res := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		// random ikind
		ikind := rand.Intn(kindsLen)
		scope := kinds[ikind][0]
		base := kinds[ikind][1]

		res[i] = uint8(base + rand.Intn(scope))
	}

	return string(res)
}

// GetRandNum 获取一个 min-max(不含max) 区间的随机数
func GetRandNum(min, max int) int {
	if max < min {
		tmp := max
		max = min
		min = tmp
	} else if min == max {
		return min
	}

	rand.Seed(time.Now().UnixNano()) // 随机种子

	num := rand.Intn(max - min)

	return min + num
}

// 真随机
func RealRand(size int) string {
	var buf strings.Builder

	for i := 0; i < size; i++ {
		result, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(letterBytesLen)))
		if err != nil {
			return ""
		}

		index := int(result.Int64())
		buf.WriteString(letterBytes[index : index+1])
	}
	str := buf.String()

	return str
}

func GetGroupUuidString() string {
	return RealRand(16)
}

var randomInt = []byte("0123456789")

func GetRandomIntString(length int) string {
	var result []byte
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, randomInt[r.Intn(len(randomInt))])
	}
	return string(result)
}

func GetRandomUserIdString() string {
	return GetRandomIntString(16)
}
