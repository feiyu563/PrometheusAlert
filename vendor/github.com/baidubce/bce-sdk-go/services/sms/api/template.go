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

// template.go - the sms template APIs definition supported by the SMS service

package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

// CreateTemplate - create an sms template
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create an sms template
// RETURNS:
//     - *api.CreateTemplateResult: the result of creating an sms template
//     - error: the return error if any occurs
func CreateTemplate(cli bce.Client, args *CreateTemplateArgs) (*CreateTemplateResult, error) {
	if err := CheckError(args != nil, "CreateTemplateArgs can not be nil"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.Content) > 0, "content can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.CountryType) > 0, "countryType can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.Name) > 0, "name can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.SmsType) > 0, "smsType can not be blank"); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_TEMPLATE)
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	req.SetParam(CLIENT_TOKEN, util.NewUUID())
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
	result := &CreateTemplateResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteTemplate - delete an sms template
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to delete an sms template
// RETURNS:
//     - error: the return error if any occurs
func DeleteTemplate(cli bce.Client, args *DeleteTemplateArgs) error {
	if err := CheckError(args != nil, "DeleteTemplateArgs can not be nil"); err != nil {
		return err
	}
	if err := CheckError(len(args.TemplateId) > 0, "templateId can not be blank"); err != nil {
		return err
	}
	return bce.NewRequestBuilder(cli).
		WithMethod(http.DELETE).
		WithURL(REQUEST_URI_TEMPLATE + bce.URI_PREFIX + args.TemplateId).
		Do()
}

// ModifyTemplate - modify an sms template
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to modify an sms template
// RETURNS:
//     - error: the return error if any occurs
func ModifyTemplate(cli bce.Client, args *ModifyTemplateArgs) error {
	if err := CheckError(args != nil, "ModifyTemplateArgs can not be nil"); err != nil {
		return err
	}
	if err := CheckError(len(args.TemplateId) > 0, "templateId can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.Content) > 0, "content can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.CountryType) > 0, "countryType can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.Name) > 0, "name can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.SmsType) > 0, "smsType can not be blank"); err != nil {
		return err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_TEMPLATE + bce.URI_PREFIX + args.TemplateId)
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
	defer func() { resp.Body().Close() }()
	return nil
}

// GetTemplate - modify an sms template
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to modify an sms template
// RETURNS:
//     - error: the return error if any occurs
func GetTemplate(cli bce.Client, args *GetTemplateArgs) (*GetTemplateResult, error) {
	if err := CheckError(args != nil, "GetTemplateResult can not be nil"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.TemplateId) > 0, "templateId can not be blank"); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_TEMPLATE + bce.URI_PREFIX + args.TemplateId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &GetTemplateResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}
