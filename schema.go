package tianyan

import "github.com/shopspring/decimal"

type (
	CommonResponse struct {
		Reason    string `json:"reason"`
		ErrorCode int    `json:"error_code"`
	}

	BaseResponse[T any] struct {
		CommonResponse
		Result T `json:"result,omitempty"`
	}

	Pagination[T any] struct {
		Total int `json:"total"` // 总数
		Items T   `json:"items"`
	}

	ZhiXingInfoItem struct {
		Pname          string          `json:"pname"`          // 被执行人名称
		CaseCode       string          `json:"caseCode"`       // 案号
		PartyCardNum   string          `json:"partyCardNum"`   // 身份证号／组织机构代码
		ExecCourtName  string          `json:"execCourtName"`  // 执行法院
		CaseCreateTime int64           `json:"caseCreateTime"` // 立案日期
		ExecMoney      decimal.Decimal `json:"execMoney"`      // 执行标的（元）
	}

	BaseInfoIndustry struct {
		CategoryMiddle string `json:"categoryMiddle"` // 国民经济行业分类中类
		CategoryBig    string `json:"categoryBig"`    // 国民经济行业分类大类
		Category       string `json:"category"`       // 国民经济行业分类门类
		CategorySmall  string `json:"categorySmall"`  // 国民经济行业分类小类
	}

	BaseInfoResponse struct {
		HistoryNames          string            `json:"historyNames"`          // 曾用名
		CancelDate            int64             `json:"cancelDate,omitempty"`  // 注销日期
		RegStatus             string            `json:"regStatus"`             // 企业状态
		RegCapital            string            `json:"regCapital"`            // 注册资本
		City                  string            `json:"city"`                  // 市
		CompForm              string            `json:"compForm,omitempty"`    // 组成形式，1-个人经营、2-家庭经营
		StaffNumRange         string            `json:"staffNumRange"`         // 人员规模
		BondNum               string            `json:"bondNum,omitempty"`     // 股票号
		HistoryNameList       []string          `json:"historyNameList"`       // 曾用名
		Industry              string            `json:"industry"`              // 行业
		BondName              string            `json:"bondName"`              // 股票名
		RevokeDate            int64             `json:"revokeDate,omitempty"`  // 吊销日期
		Type                  int               `json:"type"`                  // 法人类型，1 人 2 公司
		UpdateTimes           int64             `json:"updateTimes"`           // 更新时间
		LegalPersonName       string            `json:"legalPersonName"`       // 法人
		RevokeReason          string            `json:"revokeReason"`          // 吊销原因
		RegNumber             string            `json:"regNumber"`             // 注册号
		CreditCode            string            `json:"creditCode"`            // 统一社会信用代码
		Property3             string            `json:"property3"`             // 英文名
		UsedBondName          string            `json:"usedBondName"`          // 股票曾用名
		ApprovedTime          int64             `json:"approvedTime"`          // 核准时间
		FromTime              int64             `json:"fromTime"`              // 经营开始时间
		SocialStaffNum        int               `json:"socialStaffNum"`        // 参保人数
		ActualCapitalCurrency string            `json:"actualCapitalCurrency"` // 实收注册资本币种 人民币 美元 欧元 等
		Alias                 string            `json:"alias"`                 // 简称
		CompanyOrgType        string            `json:"companyOrgType"`        // 企业类型
		Id                    int64             `json:"id"`                    // 企业id
		CancelReason          string            `json:"cancelReason"`          /// 注销原因
		OrgNumber             string            `json:"orgNumber"`             // 组织机构代码
		ToTime                int64             `json:"toTime"`                // 经营结束时间
		ActualCapital         string            `json:"actualCapital"`         // 实收注册资金
		EstiblishTime         int64             `json:"estiblishTime"`         // 成立日期
		RegInstitute          string            `json:"regInstitute"`          // 登记机关
		BusinessScope         string            `json:"businessScope"`         // 经营范围
		TaxNumber             string            `json:"taxNumber"`             // 纳税人识别号
		RegLocation           string            `json:"regLocation"`           // 注册地址
		RegCapitalCurrency    string            `json:"regCapitalCurrency"`    // 注册资本币种 人民币 美元 欧元 等
		Tags                  string            `json:"tags"`                  // 企业标签
		District              string            `json:"district"`              // 区
		BondType              string            `json:"bondType"`              // 股票类型
		Name                  string            `json:"name"`                  // 企业名
		PercentileScore       int               `json:"percentileScore"`       // 企业评分
		IndustryAll           *BaseInfoIndustry `json:"industryAll"`           // 国民经济行业分类
		IsMicroEnt            int               `json:"isMicroEnt"`            // 是否是小微企业 0不是 1是
		Base                  string            `json:"base"`                  // 省份简称
	}
)
