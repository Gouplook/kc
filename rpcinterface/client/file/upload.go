package file

import (
	"context"
	"git.900sui.cn/kc/rpcinterface/client"
	"git.900sui.cn/kc/rpcinterface/interface/file"
)

type Upload struct {
	client.Baseclient
}

func (upload *Upload) Init() *Upload {
	upload.ServiceName = "rpc_file"
	upload.ServicePath = "Upload"
	return upload
}

// 图片上传
func (upload *Upload) UploadImage(ctx context.Context, args *file.ArgsFile, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "UploadImage", args, reply)
}

// 图片上传-开放平台
func (upload *Upload) OpenPlatFormV1UploadImage(ctx context.Context, args *file.ArgsFile, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "OpenPlatFormV1UploadImage", args, reply)
}

// 文件上传
func (upload *Upload) UploadFile(ctx context.Context, args *file.ArgsFile, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "UploadFile", args, reply)
}

// 根据id查图片
func (upload *Upload) GetImageById(ctx context.Context, id int, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetImageById", id, reply)
}

// 根据ids查图片
func (upload *Upload) GetImageByIds(ctx context.Context, ids []int, reply *map[int]file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetImageByIds", ids, reply)
}

// 根据hash查图片
func (upload *Upload) GetImageByHash(ctx context.Context, hash string, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetImageByHash", hash, reply)
}

// 根据hashs查图片
func (upload *Upload) GetImageByHashs(ctx context.Context, hashs []string, reply *map[string]file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetImageByHashs", hashs, reply)
}

// 根据id查文件
func (upload *Upload) GetFileById(ctx context.Context, id int, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetFileById", id, reply)
}

// 根据ids查文件
func (upload *Upload) GetFileByIds(ctx context.Context, ids []int, reply *map[int]file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetFileByIds", ids, reply)
}

// 根据hash查文件
func (upload *Upload) GetFileByHash(ctx context.Context, hash string, reply *file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetFileByHash", hash, reply)
}

// 根据hashs查文件
func (upload *Upload) GetFileByHashs(ctx context.Context, hashs []string, reply *map[string]file.ReplyFileInfo) error {
	return upload.Call(ctx, "GetFileByHashs", hashs, reply)
}

// 批量图片上传
func (upload *Upload) UploadImages(ctx context.Context, args []*file.ArgsFile, reply *[]file.ReplyFileInfo) error {
	return upload.Call(ctx, "UploadImages", args, reply)
}

// 批量文件上传
func (upload *Upload) UploadFiles(ctx context.Context, args []*file.ArgsFile, reply *[]file.ReplyFileInfo) error {
	return upload.Call(ctx, "UploadFiles", args, reply)
}


//根据url地址获取远程图片
func (upload *Upload) SaveImgFromUrl(ctx context.Context, imgUrl *string, reply *file.ReplyFileInfo) error{
	return upload.Call(ctx, "SaveImgFromUrl", imgUrl, reply)
}