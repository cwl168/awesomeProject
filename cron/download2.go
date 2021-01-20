package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//downloadFile1("cron/1.jpg","https://www.twle.cn/static/i/img1.jpg");
	DownloadFileProgress("cron/1.jpg", "https://www.twle.cn/static/i/img1.jpg")
	//DownloadFileProgress("cron/GeoLite2-City.tar.gz", "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=O38GA2SviPLnqfF5&suffix=tar.gz")
}
func downloadFile1(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

type R struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *R) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f%%", float64(r.Current*10000/r.Total)/100)

	return
}

func DownloadFileProgress(filename, url string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	reader := &R{
		Reader: r.Body,
		Total:  r.ContentLength,
	}
	// 调用/usr/local/go/src/io/io.go文件的copyBuffer方法
	// copyBuffer方法中for 循环调用 src.Read(buf)(由于实现了read接口，调用本代码中的Read方法)
	_, _ = io.Copy(f, reader)
}
