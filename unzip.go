package gz

import (
	z "archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
)

/*
	解压缩

	usage:
		var err error
		zip.DeCompress("/data/test.zip", "/data/test")
		err = zip.DeCompress("C:\data\test.zip", "C:\data\test")
		if err != nil {
			fmt.Println(err.Error())
		}
*/
func UnZip(zip_file_name, folder_path string) (err error) {

	folder_path, _ = filepath.Abs(folder_path)

	if err := os.MkdirAll(folder_path, 0644); err != nil {
		fmt.Println(err)
	}

	fr, err := z.OpenReader(zip_file_name)
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	// 创建文件夹
	for _, v := range fr.Reader.File {
		if v.FileInfo().IsDir() {
			if err := os.MkdirAll(path.Join(folder_path, v.Name), 0644); err != nil {
				log.Println("unzip", err)
			}
		}

		// 某些情况下, fileinfo 并没有文件夹标识, 所以要根据文件实际的目录创建文件夹
		dir, _ := filepath.Abs(path.Join(folder_path, filepath.Dir(v.Name)))
		if err := os.MkdirAll(dir, 0644); err != nil {
			log.Println("unzip", err)
		}
	}

	var wg sync.WaitGroup
	p := 0
	pool := 100

	// 提取文件
	for _, obj := range fr.Reader.File {

		// log.Println("unzip: ", i, obj.Name)

		wg.Add(1)
		p++

		go func(file *z.File) {

			defer wg.Done()

			if file.FileInfo().IsDir() {
				return
			}

			r, err := file.Open()
			if err != nil {
				log.Println(err)
				return
			}
			defer r.Close()

			f, err := os.Create(path.Join(folder_path, file.Name))
			if err != nil {
				log.Println("unzip create file error:", err)
				return
			}
			defer f.Close()

			io.Copy(f, r)

		}(obj)

		if p == pool {
			wg.Wait()
			p = 0
		}

	}

	wg.Wait()

	return nil
}
