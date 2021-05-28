//SingleImgModel
//2020-03-27 15:38:30

package models

import(
		"git.900sui.cn/kc/base/common/functions"
		"git.900sui.cn/kc/base/common/models/base"
	)

//表结构体
type SingleImgModel struct {
	Model *base.Model
	Field SingleImgModelField
}

//表字段
type SingleImgModelField struct{
	T_table	string	`default:"single_img"`
	F_id	string	`default:"id"`
	F_single_id	string	`default:"single_id"`
	F_img_id	string	`default:"img_id"`
	F_type	string	`default:"type"`
}

//初始化
func (s *SingleImgModel) Init() *SingleImgModel{
	functions.ReflectModel(&s.Field)
	s.Model = base.NewModel(s.Field.T_table)
	return s
}

//新增数据
func (s *SingleImgModel) Insert(data map[string]interface{}) (int){
	result,_ := s.Model.Data(data).Insert()
	return result
}

//批量添加
func (s *SingleImgModel) InsertAll( data []map[string]interface{} ) (int)  {
	result,_ := s.Model.InsertAll(data)
	return result
}

//获取单项目的图片数据
func (s *SingleImgModel) GetBySingleId( singleId int ) []map[string]interface{}  {
	return  s.Model.Where(map[string]interface{}{
		s.Field.F_single_id:singleId,
	}).Select()
}
//获取单项目的图片数据
func (s *SingleImgModel) GetBySingleIds( where map[string]interface{} ) []map[string]interface{}  {
	if len(where)==0{
		return  make([]map[string]interface{},0)
	}
	return  s.Model.Where(where).Select()
}
//删除
func (s *SingleImgModel) DelByIds( ids []int ) bool  {
	_, err := s.Model.Where(map[string]interface{}{
		s.Field.F_id:[]interface{}{"IN", ids},
	}).Delete()
	if err != nil{
		return false
	}

	return true
}