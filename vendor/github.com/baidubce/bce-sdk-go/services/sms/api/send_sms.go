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

// send_sms.go - the send sms APIs definition supported by the SMS service

package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

// SendSms - send an sms message
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to send an sms message
// RETURNS:
//     - *api.SendSmsResult: the result of sending an sms message
//     - error: the return error if any occurs
func SendSms(cli bce.Client, args *SendSmsArgs) (*SendSmsResult, error) {
	if err := CheckError(len(args.Mobile) > 0, "mobile can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.SignatureId) > 0, "signatureId can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.Template) > 0, "templateId can not be blank"); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_SEND_SMS)
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	if len(args.ClientToken) > 0 {
		req.SetParam(CLIENT_TOKEN, args.ClientToken)
	} else {
		req.SetParam(CLIENT_TOKEN, util.NewUUID())
	}
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &SendSmsResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err

	}
	return result, nil
}
