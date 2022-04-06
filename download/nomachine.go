package download

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"autopard.com/desa/common"
)

func DownloadNomachine(buildType BuildType, osType common.OsType, archType common.ArchType, pkgType common.PackageType) error {
	//https://www.nomachine.com/download/linux&id=29&s=Raspberry // https://www.nomachine.com/download/download&id=106&s=Raspberry&hw=Pi2
	var downLoadId int32 = 0
	switch osType {

	case common.OS_LINUX:
		switch archType {
		case common.ARCH_AMD64:
			switch pkgType {

			case common.PACKAGE_RPM:
				downLoadId = 1 //url := "https://www.nomachine.com/download/download&id=1" // https://download.nomachine.com/download/7.9/Linux/nomachine_7.9.2_1_x86_64.rpm
			case common.PACKAGE_ARCHIVE:
				downLoadId = 2 //url := "https://www.nomachine.com/download/download&id=2" // https://download.nomachine.com/download/7.9/Linux/nomachine_7.9.2_1_x86_64.tar.gz
			case common.PACKAGE_DEB:
				downLoadId = 4 //url := //url := "https://www.nomachine.com/download/download&id=4" // https://download.nomachine.com/download/7.9/Linux/nomachine_7.9.2_1_amd64.deb
			default:
				return fmt.Errorf("unsupported package")
			}

		case common.ARCH_X86:
			switch pkgType {
			case common.PACKAGE_ARCHIVE:
				downLoadId = 3 //url := "https://www.nomachine.com/download/download&id=3" // https://download.nomachine.com/download/7.9/Linux/nomachine_7.9.2_1_i686.tar.gz
			case common.PACKAGE_RPM:
				downLoadId = 5 //url := "https://www.nomachine.com/download/download&id=5" // https://download.nomachine.com/download/7.9/Linux/nomachine_7.9.2_1_i686.rpm
			case common.PACKAGE_DEB:
				downLoadId = 6 //url := "https://www.nomachine.com/download/download&id=6" // https://download.nomachine.com/download/7.9/Linux/nomachine_7.9.2_1_i386.deb
			}
		}

	case common.OS_DARWIN:
		downLoadId = 7 //url = "https://www.nomachine.com/download/download&id=7" // https://download.nomachine.com/download/7.9/MacOSX/nomachine_7.9.2_1.dmg
	case common.OS_WINDOWS:
		downLoadId = 8 //url = "https://www.nomachine.com/download/download&id=8" // https://download.nomachine.com/download/7.9/Windows/nomachine_7.9.2_1.exe
	default:
		return errors.New("unsupported platform")
	}
	url := fmt.Sprintf("https://www.nomachine.com/download/download&id=%d", downLoadId)
	fmt.Println(url)
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
		bodyStr := string(data)
		target := "'https://download.nomachine.com/download/"
		start := strings.Index(bodyStr, target)
		if start > 0 {
			start += 1
			end := strings.Index(bodyStr[start:], "');\"")
			if end > 0 {
				targetUrl := bodyStr[start : start+end]
				fmt.Println("Get target url:", targetUrl)
				fileds := strings.Split(targetUrl, "/")
				file := fileds[len(fileds)-1]
				fmt.Println("downloading", file)
				tmpDir := os.TempDir()
				//common.DownloadFile(targetUrl, path.Join(tmpDir, file))
				common.DownloadFileProgress(targetUrl, path.Join(tmpDir, file))
				return nil
			}
		}
	}
	return nil
}
