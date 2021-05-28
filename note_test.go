package main_test

import (
	"git.900sui.cn/kc/rpcCards/common/models"
	"testing"
)

func TestNoteModel(t *testing.T){
	nm := new(models.NoteModel).Init()
	t.Log(nm.GetNotes())
	t.Log(nm.GetRequirements())
}
