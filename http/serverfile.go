package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func SaveFile(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Write([]byte("the method must be POST"))
		return
	}
	var (
		status int
		err    error
	)
	defer func() {
		if nil != err {
			http.Error(w, err.Error(), status)
		}
	}()

	// parse request
	const _24K = (1 << 20) * 24
	if err = req.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		fmt.Println("Parse MultipartForm failed")
		return
	}

	if req.MultipartForm == nil {
		fmt.Println("MultipartForm is null")
		status = http.StatusInternalServerError
		return
	} else if req.MultipartForm.File == nil {
		fmt.Println("MultipartForm.File is null")
		status = http.StatusInternalServerError
		return
	}

	for name, fileHeaders := range req.MultipartForm.File {
		fmt.Println(name)
		for _, fileHeader := range fileHeaders {
			var infile multipart.File
			if infile, err = fileHeader.Open(); err != nil {
				status = http.StatusInternalServerError
				return
			}

			//创建文件用来存储
			var outfile *os.File
			if outfile, err = os.Create(fileHeader.Filename); err != nil {
				status = http.StatusInternalServerError
				return
			}
			defer outfile.Close()

			var written int64
			if written, err = io.Copy(outfile, infile); err != nil {
				status = http.StatusInternalServerError
				return
			}

			w.Write([]byte("上传文件：" + fileHeader.Filename + ";长度:" + strconv.Itoa(int(written))))
		}
	}
	filename := req.URL.Query()["filename"][0]
	w.Write([]byte("hello " + filename))
}

func main() {
	http.HandleFunc("/file", SaveFile)
	http.ListenAndServe(":8080", nil)
}
