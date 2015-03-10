package main

import (
	"fmt"
	"github.com/chanxuehong/wechat/client"
)

// *DefaultTokenService 实现了 TokenService 接口.
// 当然你也可以不用默认的实现, 这个时候就需要你自己实现 TokenService 接口了,
// 根据你自己的实现, tokenService 不一定要求作为全局变量,
// 但是如果用默认的实现 client.NewDefaultTokenService, 一个 appid 只能有一个实例.
var tokenService = client.NewDefaultTokenService(
	"xxxxxxxxxxxxxxxxxx",               // 公众号 appid
	"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", // 公众号 appsecret
	nil,
)

func main() {
	wechatClient := client.NewClient(tokenService, nil)

	qrcode, err := wechatClient.QRCodeTemporaryCreate(100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(qrcode)
}
