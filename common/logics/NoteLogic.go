package logics

import "git.900sui.cn/kc/rpcCards/common/models"

type NodeLogic struct {

}

func (n *NodeLogic)GetNotes()(list []string, err error){
	nm := new(models.NoteModel).Init()
	list, err = nm.GetNotes()
	return
}

func (n *NodeLogic)GetRequirements()(list []string, err error){
	nm := new(models.NoteModel).Init()
	list, err = nm.GetRequirements()
	return
}
