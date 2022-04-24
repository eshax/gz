package gz

import (
	z "archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
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
	//r.reader.file 是一个集合，里面包括了压缩包里面的所有文件
	for _, file := range fr.Reader.File {

		//判断文件该目录文件是否为文件夹
		if file.FileInfo().IsDir() {
			log.Println("mkdir:", path.Join(folder_path, file.Name))
			if err := os.MkdirAll(path.Join(folder_path, file.Name), 0644); err != nil {
				fmt.Println(err)
			}
			continue
		}

		// 某些情况下, fileinfo 并没有文件夹标识, 所以要根据文件实际的目录创建文件夹
		dir, _ := filepath.Abs(path.Join(folder_path, filepath.Dir(file.Name)))
		if err := os.MkdirAll(dir, 0644); err != nil {
			log.Println("mkdir:", dir)
		}

		//为文件时，打开文件
		r, err := file.Open()

		//文件为空的时候，打印错误
		if err != nil {
			fmt.Println(err)
			continue
		}
		//这里在控制台输出文件的文件名及路径
		fmt.Println("unzip: ", file.Name)

		//在对应的目录中创建相同的文件
		NewFile, err := os.Create(path.Join(folder_path, file.Name))
		if err != nil {
			fmt.Println("create error:", err)
			continue
		}
		//将内容复制
		io.Copy(NewFile, r)
		//关闭文件
		NewFile.Close()
		r.Close()
	}
	return nil
}
