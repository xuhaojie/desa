package download

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"autopard.com/desa/common"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
func httpRequest(url string) []byte {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}
	}
	defer response.Body.Close()
	finalURL := response.Request.URL.String()
	fmt.Println(finalURL)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}
	}

	return body
}

func DownloadVmware(buildType BuildType, osType common.OsType, archType common.ArchType, pkgType common.PackageType) error {
	var url string
	switch osType {
	case common.OS_WINDOWS:
		url = "https://www.vmware.com/go/getworkstation-win"
	case common.OS_LINUX:
		url = "https://www.vmware.com/go/getworkstation-linux"
	case common.OS_DARWIN:
		url = "https://www.vmware.com/go/getfusion"
	default:
		return fmt.Errorf("unsupported platform")
	}
	targetUrl, err := common.GetFinalUrl(url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Get target url:", targetUrl)
	fileds := strings.Split(targetUrl, "/")
	file := fileds[len(fileds)-1]
	tmpDir := os.TempDir()

	common.DownloadFileProgress(targetUrl, path.Join(tmpDir, file))
	return nil
}
