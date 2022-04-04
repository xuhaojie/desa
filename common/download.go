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

type PackageType int32

const (
	PACKAGE_UNKNOWN PackageType = 0
	PACKAGE_EXE     PackageType = 1
	PACKAGE_MSI     PackageType = 2
	PACKAGE_DEB     PackageType = 3
	PACKAGE_RPM     PackageType = 4
	PACKAGE_ARCHIVE PackageType = 5
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

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer
// interface and we can pass this into io.TeeReader() which will report progress on each
// write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %d B complete", wc.Total)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory. We pass an io.TeeReader
// into Copy() to report progress on the download.
func DownloadFileProgress(url string, filepath string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	tmpFileName := filepath + ".downloading"
	out, err := os.Create(tmpFileName)
	if err != nil {
		return err
	}
	//defer out.Close() 看评论

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	out.Close() // 看评论
	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")
	err = os.Rename(tmpFileName, filepath)
	if err != nil {
		return err
	}

	return nil
}
