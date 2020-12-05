package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"sort"
)

//私钥
var privateKey = []byte(`
-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBALf18GkOZlRDPMCf
3v4EFuR2COVk8yn55MLRZW8OmyxCrYL1eFcWKuoO57Gwf9fIb1nbtNG/YzIltFZm
0j8Xt+k7ioq49bC7xty/idn7iJ3xCfl8+3uDCPyhjJ47cCLDFQ0ltg7qBKOqV59m
gqS3iSJQCNpJp5h4g5opumxZB27dAgMBAAECgYBtxElkVlxG0SvyADLtvQDv52Jd
hoP1uw9+E7YDs2Jx4YNpDhF1XVvT93rWsutFlWqj4o4dTabh6E+X8phnXz6z4FhV
hFV/PcaEuvitch6+ViO4ueNJbIf1iOduxBOkUBcPy2GyLEQblA73PhRoHr2dzGNV
FTglLZd4zZYLihDSoQJBAODN5KHQeqLx+o0AXAanyg3RBWV0WL0tuKeR+O51wwd4
vf8RVr557kz09RHMnf0NlIqKMM4agBOSrN2O1zFnvskCQQDRfRcGnziGIhU3qrZw
3gXWhZ0V0bRoMYw0ZL113VnSuUcxryNFTOXLZuYh9w0XNWr4aZRNGIoJTJ1EN57r
4NV1AkBw1EYFRTrYH7VOIbkOihZqetdHhW6ofJMqX6ReIvLhBCPwKkasUUxeia+a
4GUHRlKgeh1mxHw11q82gGPXYyepAkBj9FGWQUZRCdYh13xxYh5a+ym2jXaM+Icz
QJX9tP30w56qoCwuFsDWSmSn9B704fIGUSoHlxvV0A6BzCAx3/c5AkEAsquFYnzL
3h6c4GOZXi2kV31sngiBJlZN0ieMYZ6226P3UqPQJ+T3A9jOkB0+smWijqnfLftK
zcew5p4rWcWqHQ==
-----END PRIVATE KEY-----
`)
//公钥
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC39fBpDmZUQzzAn97+BBbkdgjl
ZPMp+eTC0WVvDpssQq2C9XhXFirqDuexsH/XyG9Z27TRv2MyJbRWZtI/F7fpO4qK
uPWwu8bcv4nZ+4id8Qn5fPt7gwj8oYyeO3AiwxUNJbYO6gSjqlefZoKkt4kiUAja
SaeYeIOaKbpsWQdu3QIDAQAB
-----END PUBLIC KEY-----
`)

func main() {
	//加签参数
	params := map[string]string{
		"app_id": "1010",
		"timestamp": "1579058461",
	}

	//ksort排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//拼接参数字符串
	var dataParams string
	for _, k := range keys {
		dataParams = dataParams + k + "=" + params[k] + "&"
	}
	sortStr := dataParams[0 : len(dataParams)-1]

	//生成签名
	sign, err := RsaSignWithSha1Base64(sortStr, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf(sign)

	//校验签名
	pass := RsaVerySignWithSha1Base64(sortStr, sign, publicKey)

	fmt.Println("加签参数：", sortStr)
	fmt.Println("签名：", sign)
	if pass == nil {
		fmt.Println("验签通过")
	} else {
		fmt.Println(pass)
	}
}

//加签
func RsaSignWithSha1Base64(data string, prvKey []byte) (string, error) {
	block, _ := pem.Decode(prvKey)
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

//验签
func RsaVerySignWithSha1Base64(originalData string, signData string, pubKey []byte) error {
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}

	public, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public.Bytes)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}