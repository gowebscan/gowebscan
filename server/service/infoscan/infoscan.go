package infoscan

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/addtasklist"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type InfoScanService struct {
	FofaUrlApi string
}

func (s *InfoScanService) FofaInfoScan(requestBody addtasklist.RequestBody) {
	fmt.Println("调用了FofaInfoScan")
	//接口转字符类型
	Target := requestBody.RuleFormRef["TaskTarget"].(string)
	myVar := InfoScanService{
		FofaUrlApi: "https://fofa.info/api/v1/search/all?email=patrickcwong@hotmail.com&key=61a9133c7e9c24849975f2772b529632&qbase64=",
	}
	encodedTarget := base64.StdEncoding.EncodeToString([]byte(Target))
	InfoScan := myVar.FofaUrlApi + encodedTarget
	fmt.Println(InfoScan)
	//发送请求
	req, err := http.NewRequest("GET", InfoScan, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 发送HTTP GET请求，并获取响应结果
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取响应内容并解析JSON格式
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var fofadate map[string]interface{}
	JsonErr := json.Unmarshal([]byte(body), &fofadate)
	if JsonErr != nil {
		fmt.Println(JsonErr)
		return
	}
	results := fofadate["results"].([]interface{})
	//创建表格：
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", "Fofa")
	f.SetCellValue("Fofa", "A1", "URL")
	f.SetCellValue("Fofa", "B1", "IP地址")
	f.SetCellValue("Fofa", "C1", "端口")
	f.SetCellValue("Fofa", "D1", "状态码")
	f.SetCellValue("Fofa", "E1", "标题")
	f.SetCellValue("Fofa", "F1", "指纹信息")
	// 设置宽度为15
	f.SetColWidth("Fofa", "A", "A", 20)
	f.SetColWidth("Fofa", "B", "B", 15)
	f.SetColWidth("Fofa", "C", "C", 9)
	f.SetColWidth("Fofa", "D", "D", 7)
	f.SetColWidth("Fofa", "E", "E", 24)
	f.SetColWidth("Fofa", "F", "F", 20)
	for key := range results {
		//开始遍历results 数据
		TargetUrl := results[key].([]interface{})[0].(string)
		TargetIP := results[key].([]interface{})[1].(string)
		TargetPort := results[key].([]interface{})[2].(string)
		// 表格行
		LineUrl := "A" + strconv.Itoa(key+2)
		LineIP := "B" + strconv.Itoa(key+2)
		LinePort := "C" + strconv.Itoa(key+2)
		LineCode := "D" + strconv.Itoa(key+2)
		LineTitle := "E" + strconv.Itoa(key+2)
		fmt.Println("[ + ] 正在写入数据===》" + "Url：" + TargetUrl + " IP：" + TargetIP + " Port:" + TargetPort)
		f.SetCellValue("Fofa", LineUrl, TargetUrl)
		f.SetCellValue("Fofa", LineIP, TargetIP)
		f.SetCellValue("Fofa", LinePort, TargetPort)
		//获取 状态码 and title ane 指纹
		// 判断URL是否以"http://"或"https://"开头
		if strings.HasPrefix(TargetUrl, "http://") || strings.HasPrefix(TargetUrl, "https://") {
			fmt.Println("[ + ] 校验URL 正确" + TargetUrl)
		} else {
			if strings.Contains(TargetUrl, ":") {
				TargetUrl = "http://" + TargetUrl
				fmt.Println("[ + ] 校验URL 不正确  已修改===》" + TargetUrl)
			} else {
				TargetUrl = "http://" + TargetUrl + ":" + TargetPort
				fmt.Println("[ + ] 校验URL 不正确  已修改===》" + TargetUrl)
			}
		}
		//发送校验请求
		resp, err := http.Get(TargetUrl)
		if err != nil {
			fmt.Printf("%s: %v\n", TargetUrl, err)
			continue
		}
		defer resp.Body.Close()

		title, err := getTitle(resp)
		if err != nil {
			fmt.Printf("%s: %v\n", TargetUrl, err)
			continue
		}
		fmt.Printf("[ + ] 获取 CODE 和 Title %s: %d - %s\n", TargetUrl, resp.StatusCode, title)
		// 写入状态码 and title
		f.SetCellValue("Fofa", LineCode, resp.StatusCode)
		f.SetCellValue("Fofa", LineTitle, title)
	}
	// 获取当前系统时间
	now := time.Now()
	// 格式化时间为字符串
	timeStr := now.Format("200601021504")
	filename := timeStr + ".xlsx"
	// 保存Excel文件
	if err := f.SaveAs("./infofile/" + filename); err != nil {
		panic(err)
	}

	//文件生成成功 上传到upload
	fileupload(filename)

}

func getTitle(resp *http.Response) (string, error) {
	if resp == nil || resp.Body == nil {
		return "", errors.New("response or response body is nil")
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}
	var title string
	var getTitle func(*html.Node)
	getTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			getTitle(c)
		}
	}
	getTitle(doc)
	if title == "" {
		return "", errors.New("no title found")
	}
	return title, nil
}

// 文件上传
func fileupload(filename string) {
	fmt.Println("文件开始上传")
	fmt.Println(filename)
	// 打开文件
	file, err := os.Open("./infofile/" + filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	// 创建表单数据
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "aaa.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 发送HTTP请求
	req, err := http.NewRequest("POST", "http://localhost:8080/api/fileUploadAndDownload/upload", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("x-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiZmE5ODMyNDctNjU0My00MzdkLTkzY2QtNzVmNzI1ZmYzMDY2IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IuWTiOWTiCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJleHAiOjE2ODI3NDMwMjEsImlzcyI6InFtUGx1cyIsIm5iZiI6MTY4MjEzNzIyMX0.acBNAQYC_0_yprPCX3gdUCsRurib5KAkXEGpp2aXZ-Q")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("上传成功")
	fmt.Println(string(respBody))

}
