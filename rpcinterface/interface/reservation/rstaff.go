package reservation

//入参技师被占用时间
type ArgsReservationStaff struct {
	StaffId []int  //技师ID数组
	StartTime int64 //开始时间戳
	EndTime int64 //结束时间戳
}

//返回技师被占用时间
type ReplyReservationStaff struct {
	StaffId int  //技师ID
    ReservationTime int64 //预约到店时间戳
    StartTimePoint string//开始时间节点，格式："12:00"
	EndTimePoint string //结束时间节点，格式："12:00"
	ReservationTimeStr string //预约到店字符串 格式19:35
}