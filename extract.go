package gz

import (
	z "archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func Extract(zip_file, file_name, output_path string) (err error) {

	output_path, _ = filepath.Abs(output_path)

	if err := os.MkdirAll(output_path, 0644); err != nil {
		fmt.Println(err)
	}

	fr, err := z.OpenReader(zip_file)
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	//r.reader.file 是一个集合，里面包括了压缩包里面的所有文件
	for _, file := range fr.Reader.File {

		if file.Name != file_name {
			continue
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
		NewFile, err := os.Create(path.Join(output_path, file.Name))
		if err != nil {
			fmt.Println("gz.Extract:", err)
			return err
		}
		//将内容复制
		io.Copy(NewFile, r)
		//关闭文件
		NewFile.Close()
		r.Close()
	}
	return
}
