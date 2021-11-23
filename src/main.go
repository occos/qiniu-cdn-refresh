package main

import (
	"flag"
	"fmt"
	"github.com/qiniu/go-sdk/auth/qbox"
	"github.com/qiniu/go-sdk/cdn"
	"os"
)

func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, `
Name: cdn_refresh
Version: cdn_refresh/0.0.1
Description: 对接七牛云CDN服务，可以对CDN上部署的域名进行目录刷新
Usage: cdn_refresh -a <accessKey> -s <secretKey> -u <urlDir>

Options:
`)
	fmt.Println("------------------------\n")
	flag.PrintDefaults()
	fmt.Println("\n------------------------\n")
}

func FlagInit() (string, string, string) {
	//声明参数
	var accessKey string
	var secretKey string
	var urlDir string
	flag.StringVar(&accessKey, "a", "", "<必填> accessKey：传输秘钥（需要您自行从七牛云上获得）")
	flag.StringVar(&secretKey, "s", "", "<必填> secretKey：验证秘钥（需要您自行从七牛云上获得）")
	flag.StringVar(&urlDir, "u", "", "<必填> urlDir：需要刷新的URL目录地址，必须以\"/\"为结尾；（例如：https://www.baidu.com/aaa/bbb/）")
	flag.Usage = Usage
	flag.Parse()
	if len(accessKey) == 0 || len(secretKey) == 0 || len(urlDir) == 0 {
		fmt.Println("参数不能为空... \033[1;31;8m（输入 --help 查看命令详细用法）\033[0m")
		os.Exit(3)
	}
	return accessKey, secretKey, urlDir
}

func main() {
	//调用参数接收器
	accessKey, secretKey, urlDir := FlagInit()

	mac := qbox.NewMac(accessKey, secretKey)
	cdnManager := cdn.NewCdnManager(mac)

	// 刷新目录，刷新目录需要联系七牛技术支持开通权限
	// 单次请求链接不可以超过10个，如果超过，请分批发送请求
	dirsToRefresh := []string{
		urlDir,
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
