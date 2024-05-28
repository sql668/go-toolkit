package cache

import (
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)
// ParseSize 将字符串的内存大小转换为数字的内存大小
func ParseSize(size string) (int64,error){
	re,_ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size),[]byte("")))
	num,err := strconv.ParseInt(strings.Replace(size,unit,"",1),10,64)
	unit = strings.ToUpper(unit)
	var bytenum int64 = 0
	switch unit {
	case "B":
		bytenum = num
	case "KB":
		bytenum = num * KB
	case "MB":
		bytenum = num * MB
	case "GB":
		bytenum = num * GB
	case "TB":
		bytenum = num * TB
	case "PB":
		bytenum = num
	default:
		num = 0
		bytenum = 0

	}

	// 这里可以报错，也可以采用默认值
	if(num == 0){
		// 设置的内存大小不合规范，默认为 100MB
		//num = 100
		//bytenum = 100 * MB
		//unit = "MB"
		//log.Println("仅支持 B、KB、MB、GB、TB、PB")
		return bytenum,errors.New("仅支持 B、KB、MB、GB、TB、PB")
	}

	//sizeStr := strconv.FormatInt(num,10) + unit
	return bytenum,err
}


func GetValSize(val interface{}) int64{
	bytes,_ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}

