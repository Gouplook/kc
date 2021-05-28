package insurance

/**
 * @author liyang<654516092@qq.com>
 * @date 2020/11/5 10:31
 */

//预付卡已出单标记入参
type ArgsConsumeSign struct {
	RelationId      int  //卡包关联ID
	CardPackageId   int  //卡包ID
	CardPackageSn   string //卡包编号
	CardPackageType int    //卡包类型
	MerchantId  string     //商户编号
	InsuranceChannel int   //保险渠道
}


