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

// signature.go - the sms signature APIs definition supported by the SMS service

package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

// CreateSignature - create an sms signature
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create an sms signature
// RETURNS:
//     - *api.CreateSignatureResult: the result of creating an sms signature
//     - error: the return error if any occurs
func CreateSignature(cli bce.Client, args *CreateSignatureArgs) (*CreateSignatureResult, error) {
	if err := CheckError(args != nil, "CreateSignatureArgs can not be nil"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.Content) > 0, "content can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.ContentType) > 0, "contentType can not be blank"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.CountryType) > 0, "countryType can not be blank"); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_SIGNATURE)
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
	result := &CreateSignatureResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSignature - delete an sms signature
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to delete an sms signature
// RETURNS:
//     - error: the return error if any occurs
func DeleteSignature(cli bce.Client, args *DeleteSignatureArgs) error {
	if err := CheckError(args != nil, "DeleteSignatureArgs can not be nil"); err != nil {
		return err
	}
	if err := CheckError(len(args.SignatureId) > 0, "signatureId can not be blank"); err != nil {
		return err
	}
	return bce.NewRequestBuilder(cli).
		WithMethod(http.DELETE).
		WithURL(REQUEST_URI_SIGNATURE + bce.URI_PREFIX + args.SignatureId).
		Do()
}

// ModifySignature - modify an sms signature
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to modify an sms signature
// RETURNS:
//     - error: the return error if any occurs
func ModifySignature(cli bce.Client, args *ModifySignatureArgs) error {
	if err := CheckError(args != nil, "ModifySignatureArgs can not be nil"); err != nil {
		return err
	}
	if err := CheckError(len(args.SignatureId) > 0, "signatureId can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.Content) > 0, "content can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.ContentType) > 0, "contentType can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.CountryType) > 0, "countryType can not be blank"); err != nil {
		return err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_SIGNATURE + bce.URI_PREFIX + args.SignatureId)
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

// GetSignature - get the detail of an sms signature
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to get the detail of an sms signature
// RETURNS:
// 	   - *api.GetSignatureResult: the detail of an sms signature
//     - error: the return error if any occurs
func GetSignature(cli bce.Client, args *GetSignatureArgs) (*GetSignatureResult, error) {
	if err := CheckError(args != nil, "GetSignatureArgs can not be nil"); err != nil {
		return nil, err
	}
	if err := CheckError(len(args.SignatureId) > 0, "signatureId can not be blank"); err != nil {
		return nil, err
	}
	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_SIGNATURE + bce.URI_PREFIX + args.SignatureId)
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &GetSignatureResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}
