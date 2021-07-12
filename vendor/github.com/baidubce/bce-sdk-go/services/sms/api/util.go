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

// util.go - define the utilities for api package of SMS service

package api

import (
	"fmt"
)

const (
	REQUEST_URI_SEND_SMS  = "/api/v3/sendsms"
	REQUEST_URI_SIGNATURE = "/sms/v3/signatureApply"
	REQUEST_URI_TEMPLATE  = "/sms/v3/template"
	REQUEST_URI_QUOTA     = "/sms/v3/quota"
	CLIENT_TOKEN          = "clientToken"
)

func CheckError(condition bool, errMessage string) error {
	if !condition {
		return fmt.Errorf(errMessage)
	}
	return nil
}
