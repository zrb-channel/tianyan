package tianyan

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/shopspring/decimal"
	"github.com/zrb-channel/utils"
)

// ZhiXingInfo
// @param ctx
// @param taxName
// @param userName
// @date 2022-05-17 19:27:37
func ZhiXingInfo(ctx context.Context, conf *AuthConfig, taxName, userName string) (bool, decimal.Decimal, error) {
	params := map[string]string{
		"name":      taxName,
		"humanName": userName,
	}

	resp, err := utils.Request(ctx).SetBasicAuth(conf.Username, conf.Password).SetQueryParams(params).Get("http://tianyan.zrb8.cn/api/v1/zhixinginfo")
	if err != nil {
		return false, decimal.Decimal{}, err
	}

	if resp.StatusCode() != http.StatusOK {
		return false, decimal.Decimal{}, errors.New(resp.Status())
	}

	result := &BaseResponse[*Pagination[[]*ZhiXingInfoItem]]{}

	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return false, decimal.Decimal{}, err
	}

	if result.ErrorCode != 0 {
		return false, decimal.Decimal{}, errors.New(result.Reason)
	}

	var value decimal.Decimal

	for _, item := range result.Result.Items {
		value = value.Add(item.ExecMoney)
	}

	return result.Result.Total > 0, value, err
}
