package gov


//交通银行开户申请字段
type OpenAccData struct {
	IntentOpenOrgRegion	 string 	 `mapstructure:"intentOpenOrgRegion" form:"intentOpenOrgRegion" json:"intent_open_org_region"`
	ThdId 				 string 	 `mapstructure:"thdId" form:"thdId" json:"thd_id"`
	IntentOpenOrg 		 string 	 `mapstructure:"intentOpenOrg" form:"intentOpenOrg" json:"intent_open_org"`
	IntentOpenOrgName	 string 	 `mapstructure:"intentOpenOrgName" form:"intentOpenOrgName" json:"intent_open_org_name"`
	OpenContactName 	 string 	 `mapstructure:"openContactName" form:"openContactName" json:"open_contact_name"`
	OpenContactIdType 	 string 	 `mapstructure:"openContactIdType" form:"openContactIdType" json:"open_contact_id_type"`
	OpenOntactIdNo 		 string 	 `mapstructure:"openOntactIdNo" form:"openOntactIdNo" json:"open_ontact_id_no"`
	ExpireDate 			 string 	 `mapstructure:"expireDate" form:"expireDate" json:"expire_date"`
	DeadDate 			 string 	 `mapstructure:"deadDate" form:"deadDate" json:"dead_date"`
	OpenContactMobile 	 string 	 `mapstructure:"openContactMobile" form:"openContactMobile" json:"open_contact_mobile"`
	EmailAddr 			 string 	 `mapstructure:"emailAddr" form:"emailAddr" json:"email_addr"`
	OpenContactDocId 	 string 	 `mapstructure:"openContactDocId" form:"openContactDocId" json:"open_contact_doc_id"`
	OpenContactDocidList []BocomDoc	 `mapstructure:"openContactDocidList" form:"openContactDocidList" json:"open_contact_docid_list"`
	ReserveDate 		 string 	 `mapstructure:"reserveDate" form:"reserveDate" json:"reserve_date"`
	ReserveTime 		 string 	 `mapstructure:"reserveTime" form:"reserveTime" json:"reserve_time"`
	AccountCreateType 	 string 	 `mapstructure:"accountCreateType" form:"accountCreateType" json:"account_create_type"`
	MainDocid 			 string 	 `mapstructure:"mainDocid" form:"mainDocid" json:"main_docid"`
	MainDocidList 		 []BocomDoc  `mapstructure:"mainDocidList" form:"mainDocidList" json:"main_docid_list"`
	CorporDocid 		 string 	 `mapstructure:"corporDocid" form:"corporDocid" json:"corpor_docid"`
	CorporDocidList 	 []BocomDoc	 `mapstructure:"corporDocidList" form:"corporDocidList" json:"corpor_docid_list"`
	MainIsLicense 		 string 	 `mapstructure:"mainIsLicense" form:"mainIsLicense" json:"main_is_license"`
	UnifiedNo 			 string 	 `mapstructure:"unifiedNo" form:"unifiedNo" json:"unified_no"`
	UnitName 			 string 	 `mapstructure:"unitName" form:"unitName" json:"unit_name"`
	CertRegisterAddr 	 string 	 `mapstructure:"certRegisterAddr" form:"certRegisterAddr" json:"cert_register_addr"`
	CertCorpName 		 string 	 `mapstructure:"certCorpName" form:"certCorpName" json:"cert_corp_name"`
	CertRegMoney 		 string 	 `mapstructure:"certRegMoney" form:"certRegMoney" json:"cert_reg_money"`
	CertRegCurrency 	 string 	 `mapstructure:"certRegCurrency" form:"CertRegCurrency" json:"cert_reg_currency"`
	CertBusinessScope 	 string 	 `mapstructure:"certBusinessScope" form:"CertBusinessScope" json:"cert_business_scope"`
	CertBeginDate 		 string 	 `mapstructure:"certBeginDate" form:"certBeginDate" json:"cert_begin_date"`
	CertEndDate 		 string 	 `mapstructure:"certEndDate" form:"certEndDate" json:"cert_end_date"`
	CorporMap 			 []CorporMap `mapstructure:"corporMap" form:"corporMap" json:"corpor_map"`
	BusiType 			 string 	 `mapstructure:"BusiType" form:"cusiType" json:"busi_type"`
	CorresponAddr 		 string 	 `mapstructure:"corresponAddr" form:"corresponAddr" json:"correspon_addr"`
	TraceNo 			 string 	 `mapstructure:"traceNo" form:"traceNo" json:"trace_no"`
	Hold1 				 string 	 `mapstructure:"hold1" form:"hold1" json:"hold1"`
	Hold2 				 string 	 `mapstructure:"hold2" form:"hold2" json:"hold2"`

}

//交通银行影像字段
type BocomDoc struct {
	PageNo 		 string  `mapstructure:"pageNo" form:"pageNo" json:"page_no"`
	IsMainId 	 string  `mapstructure:"isMainId" form:"isMainId" json:"is_main_id"`
	BillType 	 string  `mapstructure:"billType" form:"billType" json:"bill_type"`
	BillTypeName string  `mapstructure:"billTypeName" form:"billTypeName" json:"bill_type_name"`
}

//法人证照信息
type CorporMap struct {
	CorpName 		 string  `mapstructure:"corpName" form:"corpName" json:"corp_name"`
	CorporIdNo 		 string  `mapstructure:"corporIdNo" form:"corporIdNo" json:"corpor_id_no"`
	CorporExpireDate string  `mapstructure:"corporExpireDate" form:"corporExpireDate" json:"corpor_expire_date"`
	CorporDeadDate	 string  `mapstructure:"CorporDeadDate" form:"CorporDeadDate" json:"corpor_dead_date"`
	CorporIdType	 string  `mapstructure:"CorporIdType" form:"CorporIdType" json:"corpor_id_type"`
}
