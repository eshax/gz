package gz

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
	文件夹压缩

	usage:
		var err error
		err = zip.Compress("/data/test","/data/test.zip")
		log.Println(err.Error())

		if err := zip.Compress("C:\data\test","C:\data\test.zip"); err != nil {
			fmt.Println(err.Error())
		}
*/
func Zip(src_dir string, zip_file_name string) (err error) {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zip_file_name)

	src_dir, _ = filepath.Abs(src_dir)

	// 创建：zip文件
	zipfile, _ := os.Create(zip_file_name)
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(src_dir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == src_dir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		// header.Name = strings.TrimPrefix(path, string(filepath.Separator))
		header.Name = strings.TrimPrefix(strings.ReplaceAll(path, src_dir, ""), string(filepath.Separator))

		log.Println(header.Name, path)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		}

		header.Name = filepath.ToSlash(header.Name)

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return nil
}
