# go-txtfile-viewer

一个简单的文本查看工具，通过浏览器查看本地文档，支持 txt 文件和 markdown 文件。

## Get Started

使用方法：

``` bash
# build
$ cd cmd/go-txtfile-viewer
$ go build -ldflags="-s -w"
# usage
$ ./go-txtfile-viewer -h
Usage of ./go-txtfile-viewer:
  -d string
    	The dir where to serve (default ".")
  -p int
    	The listen port (default 8080)
```

具体运行效果如下：

通过 `./go-txtfile-viewer/main.go -d t` 运行。

目录页：

![dir](https://github.com/yuweizzz/go-txtfile-viewer/blob/main/view_dir.png)

txt 文件页：

![txt](https://github.com/yuweizzz/go-txtfile-viewer/blob/main/view_txt.png)

markdown 文件页：

![markdown](https://github.com/yuweizzz/go-txtfile-viewer/blob/main/view_markdown.png)

## License

[MIT license](https://github.com/yuweizzz/go-txtfile-viewer/blob/main/LICENSE)
