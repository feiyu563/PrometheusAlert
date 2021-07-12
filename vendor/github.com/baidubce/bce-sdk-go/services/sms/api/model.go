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

// model.go - definitions of the request arguments and results data structure model

package api

// SendSmsArgs defines the data structure for sending a SMS request
type SendSmsArgs struct {
	Mobile        string                 `json:"mobile"`
	Template      string                 `json:"template"`
	SignatureId   string                 `json:"signatureId"`
	ContentVar    map[string]interface{} `json:"contentVar"`
	Custom        string                 `json:"custom,omitempty"`
	UserExtId     string                 `json:"userExtId,omitempty"`
	CallbackUrlId string                 `json:"merchantUrlId,omitempty"`
	ClientToken   string                 `json:"clientToken,omitempty"`
}

// SendSmsResult defines the data structure of the result of sending a SMS request
type SendSmsResult struct {
	Code      string            `json:"code"`
	RequestId string            `json:"requestId"`
	Message   string            `json:"message"`
	Data      []SendMessageItem `json:"data"`
}

type SendMessageItem struct {
	Code      string `json:"code"`
	Mobile    string `json:"mobile"`
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

// CreateSignatureArgs defines the data structure for creating a signature
type CreateSignatureArgs struct {
	Content             string `json:"content"`
	ContentType         string `json:"contentType"`
	Description         string `json:"description,omitempty"`
	CountryType         string `json:"countryType"`
	SignatureFileBase64 string `json:"signatureFileBase64,omitempty"`
	SignatureFileFormat string `json:"signatureFileFormat,omitempty"`
}

// CreateSignatureResult defines the data structure of the result of creating a signature
type CreateSignatureResult struct {
	SignatureId string `json:"signatureId"`
	Status      string `json:"status"`
}

// DeleteSignatureArgs defines the input data structure for deleting a signature
type DeleteSignatureArgs struct {
	SignatureId string `json:"signatureId"`
}

// ModifySignatureArgs defines the input data structure for modifying parameters of a signature
type ModifySignatureArgs struct {
	SignatureId         string `json:"signatureId"`
	Content             string `json:"content"`
	ContentType         string `json:"contentType"`
	Description         string `json:"description,omitempty"`
	CountryType         string `json:"countryType"`
	SignatureFileBase64 string `json:"signatureFileBase64,omitempty"`
	SignatureFileFormat string `json:"signatureFileFormat,omitempty"`
}

// GetSignatureArgs defines the input data structure for Getting a signature
type GetSignatureArgs struct {
	SignatureId string `json:"signatureId"`
}

// GetSignatureResult defines the data structure of the result of getting a signature
type GetSignatureResult struct {
	SignatureId string `json:"signatureId"`
	UserId      string `json:"userId"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
	Status      string `json:"status"`
	CountryType string `json:"countryType"`
	Review      string `json:"review"`
}

// CreateTemplateArgs defines the data structure for creating a template
type CreateTemplateArgs struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	SmsType     string `json:"smsType"`
	CountryType string `json:"countryType"`
	Description string `json:"description,omitempty"`
}

// CreateTemplateResult defines the data structure of the result of creating a template
type CreateTemplateResult struct {
	TemplateId string `json:"templateId"`
	Status     string `json:"status"`
}

// DeleteTemplateArgs defines the data structure for deleting a template
type DeleteTemplateArgs struct {
	TemplateId string `json:"templateId"`
}

// ModifyTemplateArgs defines the data structure for modifying a template
type ModifyTemplateArgs struct {
	TemplateId  string `json:"templateId"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	SmsType     string `json:"smsType"`
	CountryType string `json:"countryType"`
	Description string `json:"description,omitempty"`
}

// GetTemplateArgs defines the data structure for getting a template
type GetTemplateArgs struct {
	TemplateId string `json:"templateId"`
}

// GetTemplateResult defines the data structure of the result of getting a template
type GetTemplateResult struct {
	TemplateId  string `json:"templateId"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	CountryType string `json:"countryType"`
	SmsType     string `json:"smsType"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Review      string `json:"review"`
}

// UpdateQuotaRateArgs defines the data structure for updating quota and rate limit
type UpdateQuotaRateArgs struct {
	QuotaPerDay        int `json:"quotaPerDay"`
	QuotaPerMonth      int `json:"quotaPerMonth"`
	RateLimitPerDay    int `json:"rateLimitPerMobilePerSignByDay"`
	RateLimitPerHour   int `json:"rateLimitPerMobilePerSignByHour"`
	RateLimitPerMinute int `json:"rateLimitPerMobilePerSignByMinute"`
}

// QueryQuotaRateResult defines the data structure of querying the user's quota and rate limit
type QueryQuotaRateResult struct {
	QuotaPerDay          int  `json:"quotaPerDay"`
	QuotaRemainToday     int  `json:"quotaRemainToday"`
	QuotaPerMonth        int  `json:"quotaPerMonth"`
	QuotaRemainThisMonth int  `json:"quotaRemainThisMonth"`
	QuotaWhitelist       bool `json:"quotaWhitelist"`
	RateLimitPerDay      int  `json:"rateLimitPerMobilePerSignByDay"`
	RateLimitPerHour     int  `json:"rateLimitPerMobilePerSignByHour"`
	RateLimitPerMinute   int  `json:"rateLimitPerMobilePerSignByMinute"`
	RateLimitWhitelist   bool `json:"rateLimitWhitelist"`
}
