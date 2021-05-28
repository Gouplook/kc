package insurance
//安信保险数据结构

type  AaicBasicData struct {
	BusinessName string `json:"businessName"`
	IsBrand int `json:"isBrand"`
	BrandName string `json:"brandName"`
	DocNo string `json:"docNo"`
	DocExpire string `json:"docExpire"`
	BusinessType int `json:"businessType"`
	BusinessLife string `json:"businessLife"`
	Address string `json:"address"`
	BpremisesType int `json:"bpremisesType"`
	BusinessStoreType int `json:"businessStoreType"`
	RealName string `json:"realName"`
	CardNo string `json:"cardNo"`
	Province int `json:"province"`
	ContactAddress string `json:"contactAddress"`
	Contact string `json:"contact"`
	ContactMobile string `json:"contactMobile"`
	BusinessMode int `json:"businessMode"`
	InternetPhysicalStores int `json:"internetPhysicalStores"`
	PersonCreditPhone string `json:"personCreditPhone"`
	NoShareBillsTimes int `json:"noShareBillsTimes"`
	SumCardScale int `json:"sumCardScale"`
	LastYearTurnover int `json:"lastYearTurnover"`
	LastYearProfit int `json:"lastYearProfit"`
	LastYearBalanceSheet string `json:"lastYearBalanceSheet"`
	LastYearProfitSheet string `json:"lastYearProfitSheet"`
	BusinessRewardRecord int `json:"businessRewardRecord"`
	BusinessPunishRecord int `json:"businessPunishRecord"`
	SesameCredit int `json:"sesameCredit"`
	BusineessIndus int `json:"busineessIndus"`
	CardYImg string `json:"cardYImg"`
	CardNImg string `json:"cardNImg"`
	DocImg string `json:"docImg"`
	MerchantId string `json:"merchantId"`
	CntractId string `json:"cntractId"`
	CntractStartDate string `json:"cntractStartDate"`
	CntractEndDate string `json:"cntractEndDate"`
}

type AaicRiskData struct {
	LyCardScale int `json:"lyCardScale"`
	ShopNum int `json:"shopNum"`
	LyUserNum int `json:"lyUserNum"`
	StaffNum int `json:"staffNum"`
	LyAddShopNum int `json:"lyAddShopNum"`
	PreAddUser int `json:"preAddUser"`
	PreTurnover int `json:"preTurnover"`
	PreDayUsers int `json:"preDayUsers"`
	BusinessArea int `json:"businessArea"`
	CardRate float64 `json:"cardRate"`
	ShareBills int `json:"shareBills"`
	PersonCreditRecord int `json:"personCreditRecord"`
	BusinessCreditRecord int `json:"businessCreditRecord"`
	BusinessDishonesty int `json:"businessDishonesty"`
	BusinessLaw int `json:"businessLaw"`
	BusinessRewardRecord int `json:"businessRewardRecord"`
	RiskType int `json:"riskType"`
	RiskScore float64 `json:"riskScore"`
	RiskDate string `json:"riskDate"`
}

type AaicData struct {
	BasicData AaicBasicData `json:"basicData"`
	RiskData AaicRiskData `json:"riskData"`
}

//安信返回的数据
type AaicInsureData struct {
	Amount float64 `json:"amount"`
	Applicant string `json:"applicant"`
	CntractId string `json:"cntract_id"`
	MerchantId string `json:"merchant_id"`
	Rate float64 `json:"rate"`
	Sign string `json:"sign"`
	TransNo string `json:"trans_no"`
	UnderwirteDate string `json:"underwirte_date"`
	UnderwirteFlag string `json:"underwirte_flag"`
	ExpireTime int64 `json:"expire_time"`
	CntractEndDate string `json:"cntract_end_date"`
	CntractStartDate string `json:"cntract_start_date"`
}