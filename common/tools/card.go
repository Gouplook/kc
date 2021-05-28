package tools

type CardStatus struct {
}

//定义未删除状态
func (c *CardStatus) NotDelStatus() int {
	return 0
}

//定义删除状态
func (c *CardStatus) DelStatus() int {
	return 1
}
