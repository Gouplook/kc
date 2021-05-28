/**
 * @Author: Gosin
 * @Date: 2020/2/24 17:17
 */
package common

import "git.900sui.cn/kc/kcgin"

type Paging struct {
	Page     int // 分页
	PageSize int // 分页大小
}

func (p *Paging) GetStart() int {
	if p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.GetPageSize()
}
func (p *Paging) GetPageSize() int {
	maxsize := kcgin.AppConfig.DefaultInt("page.maxsize", 100)
	if p.PageSize > maxsize {
		p.PageSize = maxsize
	}
	if p.PageSize < 1 {
		if pagesize, err := kcgin.AppConfig.Int("page.pagesize"); err == nil {
			return pagesize
		}
	}
	return p.PageSize
}
