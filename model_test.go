package main_test

import (
	"context"
	"git.900sui.cn/kc/kcgin/logs"
	"git.900sui.cn/kc/rpcCards/common/logics"
	"git.900sui.cn/kc/rpcCards/common/models"
	"git.900sui.cn/kc/rpcinterface/interface/cards"
	"testing"
)

func TestTag(t *testing.T) {

	//mTag := new(models.TagModel).Init()
	//r := mTag.Insert(map[string]interface{}{
	//	mTag.Field.F_bus_id:1,
	//	mTag.Field.F_name:"完美",
	//})

	//r := mTag.GetBusTagNum(1)

	//r := mTag.GetByBusid(1)

	//mTag.UpdateByTagid(1, map[string]interface{}{
	//	mTag.Field.F_name:"明亮",
	//})
	//
	//r := mTag.GetByTagids([]int{1,2})
	//
	//logs.Info( r )

	//utils.CreateModel("icard_single_backup")
	//utils.CreateModel("icard_goods_backup")
}

func TestSingle(t *testing.T) {
	mS := new(models.SingleModel).Init()
	//r := mS.Insert(map[string]interface{}{
	//	mS.Field.F_name:"染发离子烫",
	//	mS.Field.F_bus_id:1,
	//	mS.Field.F_bind_id:7,
	//	mS.Field.F_price:200,
	//	mS.Field.F_real_price:150,
	//	mS.Field.F_ctime:time.Now().Local().Unix(),
	//	mS.Field.F_service_time:30,
	//	mS.Field.F_sort_desc:"高级离子烫",
	//	mS.Field.F_tag_ids:"1,2",
	//})
	//
	//logs.Info( r )

	r := mS.GetBySingleids([]int{1}, []string{mS.Field.F_single_id, mS.Field.F_name})
	logs.Info(r)

}

func TestSpec(t *testing.T) {
	mSpec := new(logics.SpecLogic)
	r, err := mSpec.AddSpec(1, 1, "高级")
	logs.Info(r, err)
}

func TestSubSpec(t *testing.T) {
	mSpec := new(logics.SpecLogic)
	r := mSpec.GetSubSpec(1, 0)
	logs.Info(r)
}

func TestSubSpecs(t *testing.T) {
	mSpec := new(logics.SpecLogic)
	r, err := mSpec.GetBySpecIds([]int{1, 2, 3, 4})
	logs.Info(r, err)

	r2 := mSpec.GetByParentSpecIds(1, []int{1, 5})
	logs.Info(r2)
}

func TestCheckTags(t *testing.T) {
	mSingle := new(logics.SingleLogic)
	err := mSingle.CheckTagids(1, []int{1, 2})
	logs.Info(err)
}

func TestAddAll(t *testing.T) {
	singImg := struct {
		SubImgIds []int
	}{
		SubImgIds: []int{1, 2, 3, 4},
	}
	if len(singImg.SubImgIds) > 0 {
		mSI := new(models.SingleImgModel).Init()
		var subImg []map[string]interface{}
		for _, imgId := range singImg.SubImgIds {
			subImg = append(subImg, map[string]interface{}{
				mSI.Field.F_single_id: 1,
				mSI.Field.F_img_id:    imgId,
				mSI.Field.F_type:      cards.IMG_TYPE_album,
			})
		}
		mSI.InsertAll(subImg)
	}
}

func TestItemRcard(t *testing.T) {
	r, err := new(logics.ItemLogic).GetItemsBySsids(context.Background(), &cards.ArgsGetItemsBySsids{
		SsIds:    []int{44},
		ItemType: 7,
	})

	logs.Info(r, err)
}
