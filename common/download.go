package common

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

func Get(url string) error {

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

func GetFinalUrl(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("http.Get => %v", err.Error())
	}

	// Your magic function. The Request in the Response is the last URL the
	// client tried to access.
	finalURL := resp.Request.URL.String()

	return finalURL, nil
}

func DownloadFile(url string, file string) error {

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
