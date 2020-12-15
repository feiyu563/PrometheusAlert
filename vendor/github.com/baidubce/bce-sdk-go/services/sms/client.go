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

// client.go - define the client for SMS service

package sms

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/sms/api"
)

const (
	REGION_BJ              = "bj"
	REGION_SU              = "su"
	DEFAULT_SERVICE_DOMAIN = "smsv3." + bce.DEFAULT_REGION + "." + bce.DEFAULT_DOMAIN
)

// Client of SMS service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the SMS service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// SendSms - send an sms message
//
// PARAMS:
//     - args: the arguments to send an sms message
// RETURNS:
//     - *api.SendSmsResult: the result of sending an sms message
//     - error: the return error if any occurs
func (c *Client) SendSms(args *api.SendSmsArgs) (*api.SendSmsResult, error) {
	return api.SendSms(c, args)
}

// CreateSignature - create an sms signature
//
// PARAMS:
//     - args: the arguments to create an sms signature
// RETURNS:
//     - *api.CreateSignatureResult: the result of creating an sms signature
//     - error: the return error if any occurs
func (c *Client) CreateSignature(args *api.CreateSignatureArgs) (*api.CreateSignatureResult, error) {
	return api.CreateSignature(c, args)
}

// DeleteSignature - delete an sms signature
//
// PARAMS:
//     - args: the arguments to delete an sms signature
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) DeleteSignature(args *api.DeleteSignatureArgs) error {
	return api.DeleteSignature(c, args)
}

// ModifySignature - modify an sms signature
//
// PARAMS:
//     - args: the arguments to modify an sms signature
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) ModifySignature(args *api.ModifySignatureArgs) error {
	return api.ModifySignature(c, args)
}

// GetSignature - get the detail of an sms signature
//
// PARAMS:
//     - args: the arguments to get the detail of an sms signature
// RETURNS:
// 	   - *api.GetSignatureResult: the detail of an sms signature
//     - error: the return error if any occurs
func (c *Client) GetSignature(args *api.GetSignatureArgs) (*api.GetSignatureResult, error) {
	return api.GetSignature(c, args)
}

// CreateTemplate - create an sms template
//
// PARAMS:
//     - args: the arguments to create an sms template
// RETURNS:
//     - *api.CreateTemplateResult: the result of creating an sms template
//     - error: the return error if any occurs
func (c *Client) CreateTemplate(args *api.CreateTemplateArgs) (*api.CreateTemplateResult, error) {
	return api.CreateTemplate(c, args)
}

// DeleteTemplate - delete an sms template
//
// PARAMS:
//     - args: the arguments to delete an sms template
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) DeleteTemplate(args *api.DeleteTemplateArgs) error {
	return api.DeleteTemplate(c, args)
}

// ModifyTemplate - modify an sms template
//
// PARAMS:
//     - args: the arguments to modify an sms template
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) ModifyTemplate(args *api.ModifyTemplateArgs) error {
	return api.ModifyTemplate(c, args)
}

// GetTemplate - modify an sms template
//
// PARAMS:
//     - args: the arguments to modify an sms template
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) GetTemplate(args *api.GetTemplateArgs) (*api.GetTemplateResult, error) {
	return api.GetTemplate(c, args)
}

// QueryQuotaAndRateLimit - query the quota and rate limit
//
// RETURNS:
//     - QueryQuotaRateResult: the result of querying the quota and rate limit
//     - error: the return error if any occurs
func (c *Client) QueryQuotaAndRateLimit() (*api.QueryQuotaRateResult, error) {
	return api.QueryQuotaRate(c)
}

// UpdateQuotaAndRateLimit - update the quota or rate limit
// PARAMS:
//     - args: the arguments to update the quota or rate limit
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) UpdateQuotaAndRateLimit(args *api.UpdateQuotaRateArgs) error {
	return api.UpdateQuotaRate(c, args)
}
