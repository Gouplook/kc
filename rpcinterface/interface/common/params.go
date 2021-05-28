package common

//Input Input
type Input struct {
	Data   map[string]interface{}
	ShopID int
	BusID  int
}

//Output Output
type Output struct {
	Code    int
	Message string
	Data    map[string]interface{}
}

//PaginationInput PaginationInput
type PaginationInput struct {
	Page     int `form:"page" json:"page" binding:"required"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required"`
}

//PaginationOutput PaginationOutput
type PaginationOutput struct {
	Code    int
	Message string
	Data    *Pagination
}

//Pagination Pagination
type Pagination struct {
	TotalNum int
	//List     []interface{}
	IndexImg map[int]string
}
