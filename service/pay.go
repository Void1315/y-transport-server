/*
 * @Author: Asahi
 * @Date: 2020-04-28 18:46:56
 * @LastEditors: Asahi
 * @LastEditTime: 2020-04-28 19:15:21
 * @Description: 介绍
 */
package service

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/y-transport-server/pkg/logging"
)

var (
	appId = "2016102200741141" // 沙箱
	// appId = "2021001155683198"
	// 应用私钥（跟csr文件同级目录）
	privateKey   = "MIIEogIBAAKCAQEAnr7NAi3yCX66nidkAtB1gQUUyENSVrzlchcXMgX86POfvwnixYJlMvc1VeG+lBfsAiXrevJjIaMnUCu2JF1bSI0anQOZ+QykSdT7+oy3XEfO6n6hSR65NH25dnSSyUv49kQyxFe8+lZMTLtBx7GZ8GL9VVqsZnxkfsnUTJCtbI6CScXBQs2JEFqT47O7icDsyXVo+bU3izPcutfmUNXNCxzZslQ+uajxR1cNpTQH8lILTJwFHRBko/699JbovYAmi/oocL2N46me79R6s6WexDzm5+n2cqLlEJZdXuNG3d52p5YoUBhIloaOKZqkJHe+SItljiWQYjFGPJHwhkqYYQIDAQABAoIBAHJJDzL4fP4U/KmL1laoWVAvpkyfGxJTICTJNuvOn8veSS2yIk7rl4vfqchQo3He3wyU3DlBc9jtqwnuzkzT4ToUTAqqYxmH3mdBLL+uuvt2vyLXU6pesht2QCVlu5+sUGqLorj0KZtYscm0LqFj3V3RBm2CwAdwmrElAPC2YAR8WKW+7mLYpYbkdFUjWGElM4yIsSs3D3rOFALxh/2vVJUk7q8UR88qEw5jTztEd9St3QnbHCEeANeL+HsHFq6II+gUJ6vnK7B3XzUkKY9J5pW32JNykPWV1iHQxUc1MUr5CIcXQMbbnZfP1/x2y6AZD/y08Cv3DxsLfesNJfXxmRkCgYEA4F/+Xx6d+AqBpOgiYPwJj/eRZl/9doGcPIbAtPobCKAht1KMv0NJGz2Edcld8YEqSWNWoSmbpju1Yrad+rP9DcAPbfwGIeFUfc8VhIOFaPAHlXYPag543pywdFVdJBgB8irNwjLHJnA3DYHqdfamVblcBy//diUwTZrVqNgOQGsCgYEAtR6531ZhMx2aBSpiIVSYgfoo3XpHi3bAOmCd/z3Au/N5mYZMFrcHX2f9pL9PHtL2X5YgFS2TkslRLiX1y5ij41RdUpLCjs/pBw9Br+IrvwcZlHJChrRG9QQfKxG8bzt65jdbATGu6RLH3ptCRmJQdV2OQA3qk1/E4XXuwZIWzWMCgYAqWbQXnB3Q/Zqu16n1iFz9sYreOewLfDdTLVQeVL7Zh/AVPB2Y1EhuPdRMry9uMCUdKQTm8oWQgOChuzpiYaq397Jx3goCRxe9j2LOWkzKG9Qyn9AVAppJ2mVr79jewTFM1kt0BkWUucWKZSUuEtoegBgguyyKzkYpv7noXq+B9QKBgDxBY578ksu5nmL3jbv+89nSCsRcCO9J63vAZu/icHtW7e54clngPFCuPQERMFZV4uc7/6JsDjt4REyxDkvUlTls+Lse3iE65BCYL8c+3ETqLMVpGd7MnQqoe2INvX3X8PQkGF+WUtVxjRLu3iaiUJgrLsb6mj67TrmfTXYSjwl3AoGAZJgi9s/vzQ31eB0Wvl8Mx8TG2SOmiKH6/LvVrydw64mQuBi8ArahF45Da32LxxFOU+CqqnrL8RH7u0W8goZ4MrlAHxwydyBM405WS64NEqAaTadcqih905+2E8fd7NFbUgZmxIpBIBE8IA3pvP9cQaNtsw/5cXHWLRFKoRTbWaM="
	client, _    = alipay.New(appId, privateKey, false)
	aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnr7NAi3yCX66nidkAtB1gQUUyENSVrzlchcXMgX86POfvwnixYJlMvc1VeG+lBfsAiXrevJjIaMnUCu2JF1bSI0anQOZ+QykSdT7+oy3XEfO6n6hSR65NH25dnSSyUv49kQyxFe8+lZMTLtBx7GZ8GL9VVqsZnxkfsnUTJCtbI6CScXBQs2JEFqT47O7icDsyXVo+bU3izPcutfmUNXNCxzZslQ+uajxR1cNpTQH8lILTJwFHRBko/699JbovYAmi/oocL2N46me79R6s6WexDzm5+n2cqLlEJZdXuNG3d52p5YoUBhIloaOKZqkJHe+SItljiWQYjFGPJHwhkqYYQIDAQAB"
	// aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArgkmgqaRXFZtxU/TT0jNiEQSUPa/Evzp21iUGgy8o8GzMOvBlyA9WykbXA6oELlQT/0v5Ank+8IYp1ObcIWeZkFA76LyNM7wgrVeY7hx/gj8ioFL+E3l8u498sS6ZJlJsWyoC++9IyVjqLqk2M0xmMt4ECRX52LdxU3vN117s+AaHgdorIChw++1gx/nH8txV43Vs0npNSJBHrk0hYMhHqLzS5ZzSmdWJW+04ykxEzo5ORLwo7cg1/fbiz/l6/9P2Lpk9RvxwTjqEg+2SeyPLZ5p7zzT1oJhBr25rg7foMNmCtuAuKaKKX3YxSIa0m7MeuCvCqi8OeQUReqE1eaAtQIDAQAB"
)

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
