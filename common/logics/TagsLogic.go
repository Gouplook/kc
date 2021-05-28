//管理标签
//@author yangzhiwu<578154898@qq.com>
//@date 2020/4/10 15:55
package logics

import (
	"git.900sui.cn/kc/base/common/toolLib"
	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/mapstructure"
	"git.900sui.cn/kc/rpcCards/common/models"
	_const "git.900sui.cn/kc/rpcCards/lang/const"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"git.900sui.cn/kc/rpcinterface/interface/common"
	"time"
)

type TagsLogic struct {
}

//添加标签
func (t *TagsLogic) AddTag(tagData *cards.ArgAddTag) (tagid int, err error) {
	tagid = 0
	//获取busid
	busId, err := tagData.GetBusId()
	if err != nil {
		err = toolLib.CreateKcErr(_const.POWER_ERR)
		return
	}
	mTag := new(models.TagModel).Init()
	//判断当前商家的店铺数量是否大于最大限制数量
	hasTags := mTag.GetByBusid(busId)
	maxNum, _ := kcgin.AppConfig.Int("tag.maxNum")
	if maxNum == 0 {
		maxNum = 100 //没有取到默认100
	}
	if len(hasTags) >= maxNum {
		err = toolLib.CreateKcErr(_const.TAG_MAX_NUM)
		return
	}

	tagid = mTag.Insert(map[string]interface{}{
		mTag.Field.F_name:   tagData.Name,
		mTag.Field.F_bus_id: busId,
		mTag.Field.F_ctime:  time.Now().Local().Unix(),
	})

	return
}

//修改标签
func (t *TagsLogic) EditTag(tagData *cards.ArgEditTag) error {
	//获取busid
	_, err := tagData.GetBusId()
	if err != nil {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	mTag := new(models.TagModel).Init()
	r := mTag.UpdateByTagid(tagData.TagId, map[string]interface{}{
		mTag.Field.F_name: tagData.Name,
	})
	if r == false {
		return toolLib.CreateKcErr(_const.DB_ERR)
	}

	return nil

}

//删除标签
func (t *TagsLogic) DelTag(tagData *cards.ArgDelTag) error {
	//获取busid
	_, err := tagData.GetBusId()
	if err != nil {
		return toolLib.CreateKcErr(_const.POWER_ERR)
	}
	mTag := new(models.TagModel).Init()
	r := mTag.DelByTagid(tagData.TagId)
	if r == false {
		return toolLib.CreateKcErr(_const.DB_ERR)
	}
	return nil
}

//获取商家的标签
func (t *TagsLogic) GetBusTags(busToken common.BsToken) ([]cards.TagInfo, error) {
	busId, err := busToken.GetBusId()
	if err != nil {
		return []cards.TagInfo{}, toolLib.CreateKcErr(_const.POWER_ERR)
	}
	mTag := new(models.TagModel).Init()
	tags := mTag.GetByBusid(busId)
	var busTags = []cards.TagInfo{}
	err = mapstructure.WeakDecode(tags, &busTags)
	if err != nil {
		return []cards.TagInfo{}, toolLib.CreateKcErr(_const.POWER_ERR)
	}
	return busTags, nil
}
