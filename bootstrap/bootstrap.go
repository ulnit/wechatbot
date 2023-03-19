/*
 * @Description:
 * @Version: 1.0
 * @Autor: Sean
 * @Date: 2023-03-18 21:00:40
 * @LastEditors: Sean
 * @LastEditTime: 2023-03-18 21:10:37
 */
package bootstrap

import (
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/ulnit/wechatbot/handlers"
)

func Run() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler

	// 注册登陆二维码回调
	bot.UUIDCallback = handlers.QrCodeCallBack

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")

	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}