package models

import (
	"git.900sui.cn/kc/kcgin"
	"git.900sui.cn/kc/redis"
)

type NoteModel struct{
}

func (n *NoteModel)Init()*NoteModel{
	return n
}
func (n *NoteModel)GetNotes()(list []string, err error){
	var temp interface{}
	temp, err = redis.RedisGlobMgr.Lrange(kcgin.AppConfig.String("card.note_notes_key"), 0, -1)
	var tempList = temp.([]interface{})
	list = make([]string, len(tempList))
	for index,item := range tempList{
		list[index] = string(item.([]byte))
	}
	return
}

func (n *NoteModel) GetRequirements()(list []string, err error){
	var temp interface{}
	temp, err = redis.RedisGlobMgr.Lrange(kcgin.AppConfig.String("card.note_requirements_key"), 0, -1)
	var tempList = temp.([]interface{})
	list = make([]string, len(tempList))
	for index,item := range tempList{
		list[index] = string(item.([]byte))
	}
	return
}
