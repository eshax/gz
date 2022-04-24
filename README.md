# gz

#### 安装教程

go get gitee.com/eshax/gz@v1.0.1

#### 使用说明

1.  zip压缩

import gitee.com/eshax/gz

func main(){
    if err := gz.Zip("dist/abc", "dist/abc.zip"); err != nil {
        log.Println(err.Error())
    }
}

2. zip解压缩

import gitee.com/eshax/gz

func main(){
    if err := gz.UnZip("dist/abc.zip", "dist/abc"); err != nil {
        log.Println(err.Error())
    }
}

