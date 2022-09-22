package tianyan

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/zrb-channel/utils"
)

// BaseInfo
// @param ctx
// @param name
// @date 2022-05-17 19:27:36
func BaseInfo(ctx context.Context, name string) (*BaseInfoResponse, error) {
	params := map[string]string{"keyword": name}

	resp, err := utils.Request(ctx).SetBasicAuth("admin", "123456").SetQueryParams(params).Get("http://zrb8.cn/api/v1/baseinfo")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	result := &BaseResponse[*BaseInfoResponse]{}

	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return nil, err
	}

	if result.ErrorCode != 0 {
		return nil, errors.New(result.Reason)
	}

	return result.Result, nil
}
