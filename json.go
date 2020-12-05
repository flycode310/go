package main

import (
	"encoding/json"
	"fmt"
)

type RespData struct {
	Data InvoiceResData `json:"data"`
	Message string `json:"message"`
	Status int `json:"status"`
}
//invoice response data
type InvoiceResData struct {
	DepositBankAccount     string   `json:"depositBankAccount,omitempty"`
	InvoiceTitle           string   `json:"invoiceTitle,omitempty"`
	RecipientEmail         string   `json:"recipientEmail,omitempty"`
	InvoiceRemark          string   `json:"invoiceRemark,omitempty"`
	TaxRegistrationNumber  string   `json:"taxRegistrationNumber,omitempty"`
	RegisteredAddressPhone string   `json:"registeredAddressPhone"`
	ApplyTime			   int		`json:"applyTime,omitempty"`
	InvoiceContent		   int		`json:"invoiceContent,omitempty"`
	InvoiceValue		   int		`json:"invoiceValue,omitempty"`
	//Status			       int		`json:"status,omitempty"`
}

var m int32 = 0x12345678
var n int8 = int8(m)

func main(){
	jsonData := "{\"data\":{\"status\":11,\"invoiceTitle\":\"xxxx\"},\"message\":\"成功\",\"status\":0}"
	jsonByte := []byte(jsonData)
	resp := RespData{}
	json.Unmarshal(jsonByte, &resp)

	fmt.Println(resp.Data.InvoiceValue)
	fmt.Println(resp.Data.RegisteredAddressPhone)

	rspByte,_ := json.Marshal(resp)
	rspStr := string(rspByte)
	fmt.Println(rspStr)

	fmt.Println(n)
}
