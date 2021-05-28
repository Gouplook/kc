package common

import (
	"encoding/json"
	"git.900sui.cn/kc/base/common/functions"
	"git.900sui.cn/kc/kcgin"
	"github.com/wenzhenxi/gorsa"
	"strconv"
	"strings"
	"time"
)

// 后台管理员token

// Autoken Autoken
type Autoken struct {
	EncodeStr string //uid加密字符串
	Auid      int
	RoleId    int
}

// GetAdminInfo GetAdminInfo
func (au *Autoken) GetAdminInfo() (*ReplyAdminUserAuthBody, error) {
	if err := au.authDecrypt(); err != nil {
		return nil, err
	}
	authBody := ReplyAdminUserAuthBody{
		Auid:   au.Auid,
		RoleId: au.RoleId,
	}
	return &authBody, nil
}

// GetAuid GetAuid
func (au *Autoken) GetAuid() (int, error) {
	if err := au.authDecrypt(); err != nil {
		return 0, err
	}
	return au.Auid, nil
}

// authDecrypt 解密过程
func (au *Autoken) authDecrypt() (err error) {
	if au.EncodeStr == "" {
		err = GetInterfaceError(ENCODE_IS_NIL)
		return
	}
	//解密过程
	var publicKey = functions.GetPemPublic(kcgin.AppConfig.String("utoken.public_key"))
	var decodeStr string
	decodeStr, err = gorsa.PublicDecrypt(au.EncodeStr, publicKey)
	if err != nil {
		return GetInterfaceError(ENCODE_ERR)
	}
	decodeArr := strings.Split(decodeStr, "|")
	nowTime := time.Now().Local().Unix()
	expTime, _ := strconv.ParseInt(decodeArr[2], 10, 64)
	if expTime < nowTime {
		//已过期
		return GetInterfaceError(ENCODE_DATA_TIMEOUT)
	}
	adminUserBecryptStr := decodeArr[1]
	var becryptReply ReplyAdminUserAuthBody
	if err = json.Unmarshal([]byte(adminUserBecryptStr), &becryptReply); err != nil {
		return
	}
	au.Auid = becryptReply.Auid
	au.RoleId = becryptReply.RoleId
	//\\解密结束
	return nil
}

// ReplyAdminUserAuthBody ReplyAdminUserAuthBody
type ReplyAdminUserAuthBody struct {
	Auid   int
	RoleId int
}
