package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	//imgUrl := "https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=O38GA2SviPLnqfF5&suffix=tar.gz"
	downFileUrl := "https://codeload.github.com/aequitas/macos-menubar-wireguard/tar.gz/1.24"
	// Get the data
	resp, err := http.Get(downFileUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create("cron/wireguard.tar.gz")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	DeCompress("cron/wireguard.tar.gz", "cron/")
}
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		//filename := dest + hdr.Name
		os.MkdirAll(dest+"/"+path.Dir(hdr.Name), os.ModePerm)
		fw, _ := os.OpenFile(dest+"/"+hdr.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
		//file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(fw, tr)
	}
	return nil
}
func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}
func Uncompress(srcFilePath string, destDirPath string) error {
	os.Mkdir(destDirPath, os.ModePerm)
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if hdr.Typeflag != tar.TypeDir {
			os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)

			fw, _ := os.OpenFile(destDirPath+"/"+hdr.Name, os.O_CREATE|os.O_WRONLY, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
