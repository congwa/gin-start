package utils

import (
	"strconv"
	"strings"
	"time"
)

// ParseDuration 将字符串格式的持续时间转换为 time.Duration 类型的值
//
// 参数：
//
//	d: 待解析的字符串，可以是标准的 Go 时间库支持的格式（如 "1h30m"），也可以包含 "d" 来表示天数（如 "1d2h"）
//
// 返回值：
//
//	time.Duration: 解析后得到的持续时间
//	error: 如果解析出错，则返回非零的错误码
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
