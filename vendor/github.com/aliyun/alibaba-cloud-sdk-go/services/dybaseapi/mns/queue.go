package mns

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/signers"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"net/http"
	"strconv"
)

type Queue struct {
	credential *credentials.StsTokenCredential
	httpClient *http.Client
	isRunning bool
	config *sdk.Config
	signer auth.Signer
}

func NewClientWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken string) (queue *Queue, err error) {
	queue = &Queue{}
	err = queue.InitWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken)
	return
}


func (queue *Queue) InitWithStsToken(regionId, stsAccessKeyId, stsAccessKeySecret, stsToken string) (err error) {
	credential := &credentials.StsTokenCredential{
		AccessKeyId:       stsAccessKeyId,
		AccessKeySecret:   stsAccessKeySecret,
		AccessKeyStsToken: stsToken,
	}
	queue.signer = signers.NewStsTokenSigner(credential)
	config := queue.InitClientConfig()
	return queue.InitWithOptions(config, credential)
}

func (queue *Queue) InitClientConfig() (config *sdk.Config) {
	if queue.config != nil {
		return queue.config
	} else {
		return sdk.NewConfig()
	}
}


func (queue *Queue) InitWithOptions(config *sdk.Config, credential auth.Credential) (err error) {
	queue.isRunning = true
	queue.config = config
	if err != nil {
		return
	}
	queue.httpClient = &http.Client{}

	if config.HttpTransport != nil {
		queue.httpClient.Transport = config.HttpTransport
	}

	if config.Timeout > 0 {
		queue.httpClient.Timeout = config.Timeout
	}

	return
}


func (queue *Queue) DoActionWithSigner(request requests.AcsRequest, response responses.AcsResponse) (err error) {

	// add clientVersion
	request.GetHeaders()["x-sdk-core-version"] = sdk.Version

	if request.GetScheme() == "" {
		request.SetScheme("HTTP")
	}
	// init request params
	err = requests.InitParams(request)
	if err != nil {
		return
	}
	// signature


	httpRequest, err := buildHttpRequest(request, queue.signer)
	if err != nil {
		return
	}
	if queue.config.UserAgent != "" {
		httpRequest.Header.Set("User-Agent", queue.config.UserAgent)
	}
	var httpResponse *http.Response
	for retryTimes := 0; retryTimes <= queue.config.MaxRetryTime; retryTimes++ {
		httpResponse, err = queue.httpClient.Do(httpRequest)

		//var timeout bool
		// receive error
		if err != nil {
			if !queue.config.AutoRetry {
				return
				//} else if timeout = isTimeout(err); !timeout {
				//	// if not timeout error, return
				//	return
			} else if retryTimes >= queue.config.MaxRetryTime {
				// timeout but reached the max retry times, return
				timeoutErrorMsg := fmt.Sprintf(errors.TimeoutErrorMessage, strconv.Itoa(retryTimes+1), strconv.Itoa(retryTimes+1))
				err = errors.NewClientError(errors.TimeoutErrorCode, timeoutErrorMsg, err)
				return
			}
		}
		//  if status code >= 500 or timeout, will trigger retry
		if queue.config.AutoRetry && (err != nil || isServerError(httpResponse)) {
			// rewrite signatureNonce and signature
			httpRequest, err = buildHttpRequest(request, queue.signer)
			if err != nil {
				return
			}
			continue
		}
		break
	}
	err = responses.Unmarshal(response, httpResponse, request.GetAcceptFormat())
	// wrap server errors
	if serverErr, ok := err.(*errors.ServerError); ok {
		var wrapInfo = map[string]string{}
		wrapInfo["StringToSign"] = request.GetStringToSign()
		err = errors.WrapServerError(serverErr, wrapInfo)
	}
	return
}


func isServerError(httpResponse *http.Response) bool {
	return httpResponse.StatusCode >= http.StatusInternalServerError
}

func buildHttpRequest(request requests.AcsRequest, singer auth.Signer) (httpRequest *http.Request, err error) {
	err = signMnsRoaRequest(request, singer)

	if err != nil {
		return
	}
	requestMethod := request.GetMethod()
	requestUrl := request.BuildUrl()
	body := request.GetBodyReader()
	httpRequest, err = http.NewRequest(requestMethod, requestUrl, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}
	// host is a special case
	if host, containsHost := request.GetHeaders()["Host"]; containsHost {
		httpRequest.Host = host
	}
	return
}

func (queue *Queue) Shutdown() {
	queue.isRunning = false
}
