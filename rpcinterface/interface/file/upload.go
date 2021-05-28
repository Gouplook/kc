/**
 * @Author: Gosin
 * @Date: 2020/3/5 16:30
 */
package file

import (
	"context"
	"mime/multipart"
)

const (
	HeadImg         = 1 // 头像
	BusinessLicense = 2 // 营业执照
	CardImg         = 3 // 证件照
	ProductImg      = 4 // 产品照

	MinImgType = 1
	MaxImgType = 4

	DocFile = 1 // 文档上传

	MinFileType = 1
	MaxFileType = 1
)

type ReplyFileInfo struct {
	Id   int
	Hash string
	Path string
}

type ArgsFile struct {
	Type       int //上传文件分类
	Context    []byte
	FileHeader *multipart.FileHeader
}

type Userinfo interface {
	// 图片上传
	UploadImage(ctx context.Context, args *ArgsFile, reply *ReplyFileInfo) error
	// 根据hash查图片
	GetImageByHash(ctx context.Context, hash string, reply *ReplyFileInfo) error
	// 根据hashs查图片
	GetImageByHashs(ctx context.Context, hashs []string, reply *map[string]ReplyFileInfo) error
	// 根据id查图片
	GetImageById(ctx context.Context, id int, reply *ReplyFileInfo) error
	// 根据ids查图片
	GetImageByIds(ctx context.Context, ids []int, reply *map[int]ReplyFileInfo) error
	// 文件上传
	UploadFile(ctx context.Context, args *ArgsFile, reply *ReplyFileInfo) error
	// 根据hash查文件
	GetFileByHash(ctx context.Context, hash string, reply *ReplyFileInfo) error
	// 根据hash查文件
	GetFileByHashs(ctx context.Context, hashs []string, reply *map[string]ReplyFileInfo) error
	// 根据id查文件
	GetFileById(ctx context.Context, id int, reply *ReplyFileInfo) error
	// 根据id查文件
	GetFileByIds(ctx context.Context, ids []int, reply *map[int]ReplyFileInfo) error
	// 批量图片上传
	UploadImages(ctx context.Context, args []*ArgsFile, reply *[]ReplyFileInfo) error
	// 批量文件上传
	UploadFiles(ctx context.Context, args []*ArgsFile, reply *[]ReplyFileInfo) error
	//根据url地址获取远程图片
	SaveImgFromUrl(ctx context.Context, imgUrl *string, reply *ReplyFileInfo) error
}
