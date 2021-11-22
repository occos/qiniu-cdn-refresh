package main

import "fmt"

func main() {
	mac := qbox.NewMac("*****", "*****")
	cdnManager := cdn.NewCdnManager(mac)

	// 刷新目录，刷新目录需要联系七牛技术支持开通权限
	// 单次请求链接不可以超过10个，如果超过，请分批发送请求
	dirsToRefresh := []string{
		"https://blog.taycc.com/",
	}
	ret, err := cdnManager.RefreshDirs(dirsToRefresh)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("状态码：", ret.Code)
	fmt.Println("响应ID：", ret.RequestID)
	fmt.Println("刷新状态：", ret.Error)
}
