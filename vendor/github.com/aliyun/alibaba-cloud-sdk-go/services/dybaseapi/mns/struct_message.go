package mns

type Message struct {
	MessageId string 		`json:"MessageId" xml:"MessageId"`
	MessageBodyMD5 string 	`json:"MessageBodyMD5" xml:"MessageBodyMD5"`
	MessageBody string 		`json:"MessageBody" xml:"MessageBody"`
	ReceiptHandle string 	`json:"ReceiptHandle" xml:"ReceiptHandle"`
	EnqueueTime int 		`json:"EnqueueTime" xml:"EnqueueTime"`
	FirstDequeueTime int 	`json:"FirstDequeueTime" xml:"FirstDequeueTime"`
	NextVisibleTime int 	`json:"NextVisibleTime" xml:"NextVisibleTime"`
	DequeueCount int 		`json:"DequeueCount" xml:"DequeueCount"`
	Priority int 			`json:"Priority" xml:"Priority"`
}
