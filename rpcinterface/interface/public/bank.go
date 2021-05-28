package public

import "context"

//获取网银信息
type ReplyInternetBankList struct {
	EiconBankBranchId string //超级网银行号
	//BankName 		  string
	BankShortName	  string //银行简称
}

//根据超级网银号获取网银信息返回参数
type ReplyInternetBankInfo struct {
	EiconBankBranchId string //超级网银行号
	BankName 		  string //银行名
	BankShortName	  string //银行简称
}


type Bank interface {
	//获取网银信息
	GetInternetBankList(ctx context.Context, args *string,reply *[]ReplyInternetBankList)
	//根据超级网银号获取网银信息返回参数
	InternetBankInfo(ctx context.Context,EiconBankBranchId string,reply *ReplyInternetBankInfo)
}
