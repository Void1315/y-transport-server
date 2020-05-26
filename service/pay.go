/*
 * @Author: Asahi
 * @Date: 2020-04-28 18:46:56
 * @LastEditors: Asahi
 * @LastEditTime: 2020-04-28 19:15:21
 * @Description: 介绍
 */
package service

import (
	"io/ioutil"

	"github.com/smartwalle/alipay/v3"
	"github.com/y-transport-server/pkg/logging"
)

var (
	appId = "2016102200741141" // 沙箱
	// appId = "2021001155683198"
	// 应用私钥（跟csr文件同级目录）
	privateKey   = readPrivateKey() //  读取私钥
	client, _    = alipay.New(appId, privateKey, false)
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnr7NAi3yCX66nidkAtB1gQUUyENSVrzlchcXMgX86POfvwnixYJlMvc1VeG+lBfsAiXrevJjIaMnUCu2JF1bSI0anQOZ+QykSdT7+oy3XEfO6n6hSR65NH25dnSSyUv49kQyxFe8+lZMTLtBx7GZ8GL9VVqsZnxkfsnUTJCtbI6CScXBQs2JEFqT47O7icDsyXVo+bU3izPcutfmUNXNCxzZslQ+uajxR1cNpTQH8lILTJwFHRBko/699JbovYAmi/oocL2N46me79R6s6WexDzm5+n2cqLlEJZdXuNG3d52p5YoUBhIloaOKZqkJHe+SItljiWQYjFGPJHwhkqYYQIDAQAB"
	// aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArgkmgqaRXFZtxU/TT0jNiEQSUPa/Evzp21iUGgy8o8GzMOvBlyA9WykbXA6oELlQT/0v5Ank+8IYp1ObcIWeZkFA76LyNM7wgrVeY7hx/gj8ioFL+E3l8u498sS6ZJlJsWyoC++9IyVjqLqk2M0xmMt4ECRX52LdxU3vN117s+AaHgdorIChw++1gx/nH8txV43Vs0npNSJBHrk0hYMhHqLzS5ZzSmdWJW+04ykxEzo5ORLwo7cg1/fbiz/l6/9P2Lpk9RvxwTjqEg+2SeyPLZ5p7zzT1oJhBr25rg7foMNmCtuAuKaKKX3YxSIa0m7MeuCvCqi8OeQUReqE1eaAtQIDAQAB"
)

func readPrivateKey() string {
	b, err := ioutil.ReadFile("./config/privateKey")
	if err != nil {
		return ""
	}
	return string(b)
}

//网站扫码支付
func WebPageAlipay(uuId string, totalAmount string, returnURL string) (string, error) {
	client.LoadAliPayPublicKey(aliPublicKey)
	pay := alipay.TradePagePay{}
	// 支付成功之后，支付宝将会重定向到该 URL
	pay.ReturnURL = returnURL

	//支付标题
	pay.Subject = "支付宝支付测试"
	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = uuId
	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	//金额
	pay.TotalAmount = totalAmount

	url, err := client.TradePagePay(pay)
	if err != nil {
		return "", err
	}
	payURL := url.String()
	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	// fmt.Println(payURL)
	logging.Info(payURL, "支付url")
	//打开默认浏览器
	// payURL = strings.Replace(payURL, "&", "^&", -1)
	// exec.Command("cmd", "/c", "start", payURL).Start()
	return payURL, nil

}
