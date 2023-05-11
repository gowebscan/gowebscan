package vuldetection

import (
	"bufio"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/vuldetection"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"runtime"
)

type AddVulService struct{}

func (s *AddVulService) AddVulScan(requestBody vuldetection.VulRequestBody) {
	fmt.Println("AddVulService")
	//读取响应-》判断是否为已经有的资产 -》
	AddVulName := requestBody.VulFormRef["TaskName"].(string)
	AddVulProperty := requestBody.VulFormRef["TaskRadio"].(string)
	AddVulTarget := requestBody.VulFormRef["TaskTarget"].(string)
	fmt.Println(AddVulName)
	if AddVulProperty == "已有资产" {
		//读取已有资产
	} else if AddVulProperty == "自定义添加" {
		//对 AddVulTarget进行校验 域名 or ip 合法性
		xray(AddVulTarget)
		nuclie(AddVulTarget)
	} else {
		fmt.Println("非法输入")
	}

}

// xray扫描
func xray(AddVulTarget string) {
	fmt.Println("xray扫描")
	if runtime.GOOS == "windows" {
		fmt.Println("当前系统是 Windows" + AddVulTarget)
		// 创建 Cmd 结构体表示要执行的命令
		//cmd := exec.Command("cmd", "/c", "cd", ".\\tools", "&&", "xray.exe", "ws", "--url", AddVulTarget, "--html-output", AddVulTarget)
		cmd := exec.Command("cmd", "/c", "cd", ".\\tools", "&&", "xray.exe", "ws", "--url", AddVulTarget)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if err := cmd.Start(); err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// 读取命令的输出，并解码为 UTF-8 编码
		scanner := bufio.NewScanner(stdout)
		decoder := simplifiedchinese.GBK.NewDecoder()
		for scanner.Scan() {
			line := string(scanner.Bytes()) // 将字节切片转换为字符串类型
			decodedLine, err := decoder.String(line)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			fmt.Println(decodedLine)
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println("Error: ", err)
			return
		}
	} else if runtime.GOOS == "linux" {
		fmt.Println("当前系统是 Linux")
	} else {
		fmt.Println("当前系统不是 Windows 也不是 Linux")
	}
}

func nuclie(AddVulTarget string) {
	fmt.Println("nuclie扫描")
}
