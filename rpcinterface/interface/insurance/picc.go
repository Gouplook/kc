package insurance
//人保数据结构

type PiccBasicInfo struct {
	MerchantId string `json:"merchantId"`
	ComName string `json:"comName"`
	DocType int `json:"docType"`
	DocNo string `json:"docNo"`
	BusinessType int `json:"businessType"`
	BusinessLife int `json:"businessLife"`
	Address string `json:"address"`
	BpremisesType int `json:"bpremisesType"`
	ComIndus int `json:"comIndus"`
	BusinessMode int `json:"businessMode"`
	ContactName string `json:"contactName"`
	ContractMobile string `json:"contractMobile"`
	ContactAddress string `json:"contactAddress"`
	PostCode string `json:"postCode"`
	RealName string `json:"realName"`
	CardNo string `json:"cardNo"`
	Guarantor string `json:"guarantor"`
	Occupation string `json:"occupation"`
	TotalAssets float64 `json:"totalAssets"`
	Industry string `json:"industry"`
	BusinessScope string `json:"businessScope"`
}


type PiccOtherInfo struct {
	ShopNum int `json:"shopNum"`
	NbShopNum int `json:"nbShopNum"`
	CardScale float64 `json:"cardScale"`
	DardAdvanceFunds float64 `json:"dardAdvanceFunds"`
	UserNum int `json:"userNum"`
	StaffNum int `json:"staffNum"`
	MostCard float64 `json:"mostCard"`
	TotalCardScale float64 `json:"totalCardScale"`
	TotalUser int `json:"totalUser"`
	AddShopNum int `json:"addShopNum"`
	CloseShopNum int `json:"closeShopNum"`
	ComplainNum int `json:"complainNum"`
	StarSource int `json:"starSource"`
	StarNum float64 `json:"starNum"`
}

type PiccData struct {
	BasicInfo PiccBasicInfo `json:"basicInfo"`
	OtherInfo PiccOtherInfo `json:"otherInfo"`
}