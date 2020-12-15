/*
 * Copyright 2020 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// quota_rate.go - the quota and rate limit APIs definition supported by the SMS service

package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// QueryQuotaRate - query the quota and rate limit detail of an user
//
// RETURNS:
//     - *QueryQuotaRateResult: the result of the query
//     - error: the return error if any occurs
func QueryQuotaRate(cli bce.Client) (*QueryQuotaRateResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_QUOTA)
	req.SetMethod(http.GET)
	req.SetParam("userQuery", "")
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &QueryQuotaRateResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateQuotaRate - update the quota and rate limit detail of an user
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to update the quota and rate limit
// RETURNS:
//     - *ListBucketsResult: the result bucket list structure
//     - error: nil if ok otherwise the specific error
func UpdateQuotaRate(cli bce.Client, args *UpdateQuotaRateArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_QUOTA)
	req.SetMethod(http.PUT)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}
