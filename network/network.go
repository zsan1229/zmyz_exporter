package network

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

// 解决windows下结果乱码
func decodeGBK(s []byte) string {
	reader := transform.NewReader(strings.NewReader(string(s)), simplifiedchinese.GBK.NewDecoder())
	d, _ := ioutil.ReadAll(reader)
	return string(d)
}

func PingIP(ip string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows 下 ping 默认发送 4 次
		cmd = exec.Command("ping", "-n", "1", ip)
	} else {
		// Linux / macOS 下
		cmd = exec.Command("ping", "-c", "1", ip)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Errorf("ping 执行失败: %v", err))
	}
	var outStr string
	if runtime.GOOS == "windows" {
		outStr = decodeGBK(output)
	} else {
		outStr = string(output)
	}
	fmt.Println(outStr)
}
