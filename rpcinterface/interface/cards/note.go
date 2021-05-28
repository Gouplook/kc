package cards

import "context"

type CardExt struct {
	Notes string `mapstructure:"notes"`//注意事项
	//Requirements string `mapstructure:"requirements"` //要求
	ServiceSubscribe string `mapstructure:"service_subscribe"`
}

type CardNote struct {
	Notes string
}


type GetNotesReplies struct {
	List []string //Notes
}

type GetRequirementsReplies struct{
	List[] string //ServiceSubscribe
}

type Note interface {
	//获取注意事项列表
	GetNotes(ctx context.Context, params *EmptyParams, replies *GetNotesReplies)  error
	//获取要求列表
	GetRequirements(ctx context.Context, params *EmptyParams, replies *GetRequirementsReplies) error
}
