package smalltools

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smalltools"
	"io/ioutil"
	"regexp"
	"strings"
)

type KillService struct {
}

func (s *KillService) PorcessesScan(killRequestBody smalltools.KillRequestBody) (arrs []string) {
	fmt.Println("杀软识别操作")
	data, err := ioutil.ReadFile("./scan/杀软.txt") // 读取文件 "123.txt" 的内容并将其保存到 data 变量中
	if err != nil {                               // 如果读取文件时出错，则输出错误信息并退出程序
		fmt.Println(err)
		return
	}
	killcontents := string(data) // 将字节切片转换为字符串类型，并将其保存到 contents 变量中

	//接口转字符类型
	Target := killRequestBody.KillFormData["killtextarea"].(string)
	//fmt.Println(Target)
	lines := strings.Split(Target, "\n")
	// 创建一个空的字符串数组。
	arrs = []string{}
	for _, line := range lines {
		fields := strings.Fields(line) // 将每行数据按空格分割为切片
		if len(fields) >= 2 && fields[1] != "PID" {
			//fmt.Printf("进程 %s 的 PID 是 %s\n", fields[0], fields[1])
			pidname := fields[0]
			pid := fields[1]
			killname, killdescribe := Scan(pidname, killcontents)
			if killname != "" || killdescribe != "" {
				// 创建一个包含四个字段的 map。
				data := map[string]string{
					"killdescribe": killdescribe,
					"killname":     killname,
					"pid":          pid,
					"pidname":      pidname,
				}
				// 将 map 转换为 JSON 格式的字符串。
				jsonStr, err := json.Marshal(data)
				if err != nil {
					fmt.Println(err)
					return arrs
				}

				// 将 JSON 字符串添加到数组中。
				arrs = append(arrs, string(jsonStr))

			}
		}
	}
	return arrs
}
func Scan(pidname string, killcontents string) (killname, killdescribe string) {
	re := regexp.MustCompile(`"([^"]+)":\s*"([^"]+)"`)
	for _, line := range strings.Split(killcontents, "\n") {
		if matches := re.FindStringSubmatch(line); len(matches) >= 3 {
			if matches[1] == pidname {
				//fmt.Println("已识别得到杀软 " + pidname)
				return matches[1], matches[2]
			}
		}
	}
	return "", ""
}
