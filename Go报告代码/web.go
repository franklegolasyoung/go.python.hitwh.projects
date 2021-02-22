package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)
func HelloServer(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method {
		file,handler,err:=r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return 
		}
		fmt.Println(handler.Header)
		defer file.Close()
		f,err:=os.OpenFile("./"+handler.Filename, os.O_WRONLY | os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return 
		}
		defer f.Close()
		size, err:= io.Copy(f, file)
		if err != nil {
			fmt.Println(err)
			return 
		}
		fmt.Fprintf(w, "The size of the file is %d", size)
		return 
	}
	w.Header().Add("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(200)
	html := `
	<form enctype="multipart/form-data" action="/" method="POST">
	Choose a file:<input name="file" tyoe="file" /><br/>
	</form>`
	io.WriteString(w, html)
}

func main(){
	http.HandleFunc("/", HelloServer)
    err := http.ListenAndServe(":8200", nil)
    if err != nil {
    	fmt.Println(err)
    }
}