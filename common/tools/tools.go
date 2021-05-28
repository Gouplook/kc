package tools

import (
	"fmt"
	"math"
	"time"

	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/rpcCards/constkey"
)

var (
	ShareLink string //卡项分享链接
	RunMode   string
)

//获取两个经纬度之间的距离
func GetDistance(lng1, lat1, lng2, lat2 float64, param ...string) float64 {
	var unit = "m"
	if len(param) > 0 {
		if param[0] == "km" {
			unit = "km"
		}
	}
	radius := 6371000.0 //单位：米
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	distance := dist * radius
	if unit == "km" {
		distance = distance / 1000
	}
	return distance
}

//Interface2String Interface2String
func Interface2String(inter interface{}) string {
	switch inter.(type) {
	case string:
		return inter.(string)
	}
	return ""
}

//Interface2Int Interface2Int
func Interface2Int(inter interface{}) int {
	switch inter.(type) {
	case int:
		return inter.(int)
	}
	return 0
}

//获取指定时间的开始和结束月份 eg: 2020-10-01 00:00:00~2020-11-01 00:00:00
func GetFirstAndLastOfMonth(timestamp int64) (firstMonth, lastMonth int64) {
	now := time.Unix(timestamp, 0)
	loc, _ := time.LoadLocation("Local")
	currentYear, currentMonth, _ := now.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, loc)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)
	firstMonth = firstOfMonth.Unix()
	lastMonth = lastOfMonth.Unix()
	fmt.Println("firstMonth:", firstOfMonth.String(), "lastMonth:", lastOfMonth.String())
	return
}

//去除整形数组零值
func RemoveArrayZero(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			continue
		}
		newArr = append(newArr, arr[i])
	}
	return newArr
}

//初始化分享链接
func InitShareLink() {
	RunMode = kcgin.AppConfig.String("runmode")
	if RunMode == "prod" {
		ShareLink = constkey.ProdShareLink
	} else {
		ShareLink = constkey.TestShareLink
	}
}

func GetShareLink(itemId, shopId, itemType int) string {
	if itemId == 0 || itemType == 0 {
		return ""
	}
	return fmt.Sprintf(ShareLink, itemId, shopId, itemType)
}
