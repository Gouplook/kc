/**
 * @Author: Gosin
 * @Date: 2020/4/13 10:26
 */
package common

import (
    "git.900sui.cn/kc/base/common/functions"
    "git.900sui.cn/kc/validata"
)

// 手机号验证
func VerfiyPhone(phone string) error {
    if functions.CheckPhone(phone) {
        return nil
    }
    return GetInterfaceError(PHONE_VERIFY_ERR)
}

// 验证证件号
func VerfiyCardNum(gender string) error {
    if validata.IdCard(gender) {
        return nil
    }
    return GetInterfaceError(MEMBER_CARD_NUM_ERROR)
}