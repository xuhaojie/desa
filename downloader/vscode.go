package downloader

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func get(url string) error {

	body := "" // gen_form_body(value_package)

	req, err := http.NewRequest("GET", url, strings.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//	req.Header.Set("Referer", fmt.Sprintf("%s/%s", w.url, current_page))
	//	req.Header.Set("Cookie", cookie)
	client := &http.Client{}
	rsp, err := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		return errors.New(rsp.Status)
	}

	if err != nil {
		return err
	} else {
		defer func() {
			err := rsp.Body.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
		data, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return err
		}
		fmt.Println(data)
	}
	return nil
}

func getFinalUrl(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}

	// Your magic function. The Request in the Response is the last URL the
	// client tried to access.
	finalURL := resp.Request.URL.String()

	return finalURL, nil
}

func genVscodeUrl(build string, os string, arch string, format string) string {
	base := "https://code.visualstudio.com/sha/download"
	var result string
	switch os {
	case "win32":
		if len(arch) > 0 {
			if len(format) > 0 {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s-%s", base, build, os, arch, format)
			} else {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os, arch)
			}
		} else {
			if len(format) > 0 {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os, format)
			} else {
				result = fmt.Sprintf("%s?build=%s&os=%s", base, build, os)
			}
		}
	case "linux":
		if len(arch) > 0 {
			if len(format) > 0 {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s-%s", base, build, os, format, arch)
			} else {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os, arch)
			}
		} else {
			if len(format) > 0 {
				result = fmt.Sprintf("%s?build=%s&os=%s-%s", base, build, os, format)
			} else {
				result = fmt.Sprintf("%s?build=%s&os=%s", base, build, os)
			}
		}
	}

	return result
}

func replaceVscodeDownloadUrl(url string, build string, newbase string) string {
	// newbase = "https://vscode.cdn.azure.cn"
	//https: //vscode.cdn.azure.cn/stable/b4c1bd0a9b03c749ea011b06c6d2676c8091a70c/VSCodeUserSetup-x64-1.57.0.exe
	index := strings.Index(url, "/"+build+"/")
	if index > 0 {
		return newbase + url[index:]
	} else {
		return url
	}
}

func downloadFile(url string, file string) error {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}

	defer resp.Body.Close()

	out, err := os.Create(file)
	wt := bufio.NewWriter(out)

	defer out.Close()

	n, err := io.Copy(wt, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		return err
	}
	wt.Flush()
	return nil
}

func DownloadVscode(build string, os string, arch string, format string) error {
	// "https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=win32"
	// "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=win32-x64-user"
	// "https://code.visualstudio.com/sha/download?build=stable&os=win32-arm64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=win32-arm64-user"
	// "https://code.visualstudio.com/sha/download?build=stable&os=win32-archive"

	// "https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=linux-rpm-x64"

	// "https://code.visualstudio.com/sha/download?build=stable&os=darwin"
	// "https://code.visualstudio.com/sha/download?build=stable&os=darwin-arm64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=darwin-universal"

	// "https://code.visualstudio.com/sha/download?build=stable&os=linux-rpm-x64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64"
	// "https://code.visualstudio.com/sha/download?build=stable&os=linux-deb-x64"

	url := genVscodeUrl(build, os, arch, format)
	// uri := genVscodeUrl("stable", "win32", "x64", "archive")
	// uri := genVscodeUrl("stable", "win32", "x64", "archive")
	// uri := genVscodeUrl("stable", "win32", "x64", "archive")
	// uri := genVscodeUrl("stable", "win32", "x64", "user")
	// uri := genVscodeUrl("stable", "win32", "x64", "")
	// uri := genVscodeUrl("stable", "win32", "", "")
	fmt.Println(url)

	url, err := getFinalUrl(url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(url)
	replacedUrl := replaceVscodeDownloadUrl(url, build, "https://vscode.cdn.azure.cn")
	fmt.Println(replacedUrl)
	arr := strings.Split(replacedUrl, "/")
	file := arr[len(arr)-1]

	downloadFile(replacedUrl, "/tmp/"+file)
	return nil
}
