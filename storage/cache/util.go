package cache

import (
	"regexp"
	"strconv"
	"strings"
)

// ParseSize 将字符串的内存大小转换为数字的内存大小
func ParseSize(size string) (int64,error){
	re,_ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size),[]byte("")))
	num,err := strconv.ParseInt(strings.Replace(size,unit,"",1),10,64)
	unit = strings.ToUpper(unit)
	switch unit {
	case "B":
	case "KB":
	case "MB":
		case "GB":
			case "TB":
	case "PB":
	default:
		num = 0

	}
	return num,err
}
