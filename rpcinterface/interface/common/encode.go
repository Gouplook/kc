package common

import (
	"encoding/json"
	"fmt"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/redis"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 生成二维码信息并返回二维码
// @author liyang<654516092@qq.com>
// @date  2020/8/4 14:09

const (
	//#二维码类型
	//卡包二维码，使用该二维码进行扫码消费确认
	ERCODE_type_cardconsume = 1
	//用户个人二维码，使用该二维码获取用户信息、用户个人动态验证码（可用于消费确认）
	ERCODE_type_user = 2
	//权益卡二维码，使用该二维码进行扫码耗卡确认
	ERCODE_type_equityconsume = 3

	//二维码类型对应的rediskey值
	ERCODE_type_cardconsume_key = "kcqrcode_1000_1001"
	ERCODE_type_user_key = "kcqrcode_1000_1002"
	ERCODE_type_equityconsume_key = "kcqrcode_1000_1003"


	//公共过期key
	ERCODE_common_expire_key = "kcqrcode_common_0001_%d_%d" //%d=二维码类型 %d=二维码值
)

//二维码具体信息结构体
type QrcodeData struct {
	QrcodeType int
	Id int
	Tag string
	Extend []int
	Expire int
	QrcodeSn string
}


type EnCode struct {
	QrcodeType int     //二维码类型 //必须
	Id         int     //标识ID    //必须
	Extend     []int     //扩展值    //非必须
	NoEncode   bool    //二维码信息数据加密 默认加密 非必须
	Expire     int     //过期时间/秒,默认=2天过期 非必须
	RandomLen  int     //二维码code长度 默认=11  非必须,实际长度比设置长度多一位
	random     string  //二维码code
}

//获取二维码
func (e *EnCode) GetQrcodeStr()(encodeStr string){
	data := map[string]interface{}{
	   "QrcodeType":e.QrcodeType,
	   "Id":e.Id,
	   "Tag":"kangcun",
	}
	
	//扩展信息
	if len(e.Extend) == 0{
		e.Extend = make([]int,0)
	}
	data["Extend"] = e.Extend
	//设置二维码默认过期时间
	expire := e.Expire
	if e.Expire == 0{
		data["Expire"]  = 172800  //默认两天过期时间
		expire = 172800
	}else{
		data["Expire"]  = e.Expire
	}
	//\\设置二维码默认过期时间

    //设置二维码code默认长度
	if e.RandomLen == 0{
		e.RandomLen = 11
	}
	//\\设置二维码code默认长度
	e.random = e.genRandom()
	redisKey := e.getQrcodeRelationKey()
	expireKey := fmt.Sprintf(ERCODE_common_expire_key,e.QrcodeType,e.Id)
	bytes,_ := redis.RedisGlobMgr.Get(expireKey)
	if bytes == nil{
		data["QrcodeSn"] = e.random
		jsonBytes,_:= json.Marshal(data)
		if e.NoEncode == false{
			encodeStr = functions.EncodeStr(string(jsonBytes[:]))
		}else{
			encodeStr = string(jsonBytes[:])
		}
		_ = redis.RedisGlobMgr.Set(redisKey,encodeStr,int64(expire))
		_ = redis.RedisGlobMgr.Set(expireKey,encodeStr,int64(expire))
	}else{
		encodeStr = string(bytes.([]byte))
	}
	return
}

//检测key是否已过期,不过期则返回二维码code对应的id值
func (e *EnCode) CheckExpire(qrcodeRandom string)(reply QrcodeData,err error){
	e.random = qrcodeRandom
	redisKey := e.getQrcodeRelationKey()
	data,_:=redis.RedisGlobMgr.Get(redisKey)
	if data == nil{
		err = GetInterfaceError(QRCODE_EXPIRED)
		return
	}
	str := functions.DecodeStr(string(data.([]byte)))
	reply = QrcodeData{}
	_ = json.Unmarshal([]byte(str),&reply)
	return
}

//映射二维码类型与rediskey
//如果增加了二维码类型，则需要在此处做映射
func (e *EnCode) getQrcodeRelationKey() string{
	relationMap := map[int]string{
		ERCODE_type_cardconsume:fmt.Sprintf("%s_%s",ERCODE_type_cardconsume_key,e.random),
		ERCODE_type_user:fmt.Sprintf("%s_%s",ERCODE_type_user_key,e.random),
		ERCODE_type_equityconsume:fmt.Sprintf("%s_%s",ERCODE_type_equityconsume_key,e.random),
	}
	return relationMap[e.QrcodeType]
}

//生成随机数
func (e *EnCode) genRandom() (random string){
	n63,_ := strconv.Atoi("1"+strings.Repeat("0",e.RandomLen))
	random =  strings.Trim(fmt.Sprintf("%"+strconv.Itoa(e.RandomLen)+"v%d",rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(n63)),e.QrcodeType), " ")
	return
}





