package common

import (
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/kcgin"
	"github.com/wenzhenxi/gorsa"
	"strconv"
	"strings"
	"time"
)

type Utoken struct {
	UidEncodeStr string //uid加密字符串
}

//获取用户UID
func (u *Utoken) GetUid()(int,error){
	var err error
	if u.UidEncodeStr == ""{
		err = GetInterfaceError(ENCODE_IS_NIL)
		return 0,err
	}
	//解密过程
	var publicKey = functions.GetPemPublic(kcgin.AppConfig.String("utoken.public_key"))
	decodeStr, err := gorsa.PublicDecrypt( u.UidEncodeStr, publicKey )
	if err != nil{
		return 0, GetInterfaceError(ENCODE_ERR)
	}
	decodeArr := strings.Split( decodeStr, "|")
	nowTime := time.Now().Local().Unix()
	expTime, _ := strconv.ParseInt( decodeArr[2], 10, 64 )
	if expTime < nowTime{
		//已过期
		return 0, GetInterfaceError(ENCODE_DATA_TIMEOUT)
	}
	uid, _ := strconv.Atoi( decodeArr[1] )
	//\\解密结束
	return uid,nil
}

