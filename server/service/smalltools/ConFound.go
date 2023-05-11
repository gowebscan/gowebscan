package smalltools

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/smalltools"
)

type ConFoundService struct{}

func (s *ConFoundService) CodeFound(CodeFoundRequestBody smalltools.CodeFoundRequestBody) {

	fmt.Println("代码混淆service操作")
	shellcode := CodeFoundRequestBody.ConFoundData["shellcode"].(string) // 原始shellcode
	key := []byte("0xAA")                                                // 混淆的 key
	// 将 shellcode 转换为 byte 数组
	shellcodeBytes := []byte(shellcode)
	if v, ok := CodeFoundRequestBody.ConFoundData["checkList"]; ok {
		if s, ok := v.([]interface{}); ok {
			for _, item := range s {
				if str, ok := item.(string); ok {
					switch str {
					case "ADD":
						fmt.Println("混淆 ADD")

					case "XOR":
						fmt.Println("混淆 XOR")
						// 对 shellcode 进行 XOR 混淆
						obfuscatedShellcode := xor(shellcodeBytes, key)
						// 将混淆后的 shellcode 转换为字符串
						obfuscatedShellcodeStr := string(obfuscatedShellcode)
						fmt.Println(obfuscatedShellcodeStr) // 输出混淆后的 shellcode
					case "NOT":
						fmt.Println("混淆 NOT")
					}
				}
			}
		} else {
			fmt.Println("请您选择混淆模式")
		}
	}

}
func xor(data []byte, key []byte) []byte {
	result := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ key[i%len(key)]
	}
	return result
}
