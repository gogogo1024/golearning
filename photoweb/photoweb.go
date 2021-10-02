package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)
const (
	UPLOAD_DIR = "./uploads"
	TEMPLATE_DIR = "./views" 
	LIST_DIR = 0x0001
)
var templates = make(map[string] *template.Template)

// 上传文件
func uploadHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "GET"{
	 if err := rendHtml(w,"upload",nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
	}
	if r.Method == "POST"{
		f,h,err1 := r.FormFile("image")
		check(err1)
		filename :=h.Filename
		defer f.Close()
		t, err2 :=os.Create(UPLOAD_DIR+"/"+filename)
		check(err2)
		defer t.Close()
		_, err3 :=io.Copy(t,f) 
		check(err3)
		http.Error(w,err3.Error(),http.StatusInternalServerError)
		http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
	}
}

// 浏览文件
func viewHandler(w http.ResponseWriter,r *http.Request)  {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR+"/"+imageId
	if exists :=isExists(imagePath); exists {
		http.NotFound(w,r)
		return
	}
	w.Header().Set("Content-type","image")
	http.ServeFile(w,r,imagePath)
}

//判定文件存在
func isExists(path string)bool  {
	_, err:=os.Stat(path)
	if err != nil {
		return true
	}
	return os.IsExist(err)
}

func  listHandler(w http.ResponseWriter, r *http.Request)  {
	fileInfoArr, err:=ioutil.ReadDir(UPLOAD_DIR)
	check(err)
	locals := make(map[string]interface{})
	images :=[] string {}

	for _,fileInfo := range fileInfoArr{
		images =append(images,fileInfo.Name())
	}
	locals["images"] = images
	if err =rendHtml(w,"list",locals); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func rendHtml(w http.ResponseWriter,tmpl string, locals map[string]interface{})  error {
	err := templates[tmpl].Execute(w,locals)
	return err
} 
func check(err error)  {
	if err != nil {
		panic(err)
	}
} 
func safeHandler(fn http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter,r *http.Request){
		defer func(){
			if err,ok:=recover().(error); ok {
				http.Error(w,err.Error(),http.StatusInternalServerError)
				log.Println("WARN: panic in %v. - %v",fn,err)
				log.Println(string(debug.Stack()))
		}
	}()
	fn(w,r)
	}
}

func staticDirHandler(mux *http.ServeMux,prefix string, staticDir string, flags int)  {
	mux.HandleFunc(prefix,func(w http.ResponseWriter, r *http.Request){
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & LIST_DIR) ==0 {
			if exists :=isExists(file); !exists {
				http.NotFound(w,r)
				return
			}
		}
		http.ServeFile(w,r,file)

	})
}
func init()  {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName,templatePath string

	for _,fileInfo :=range fileInfoArr{
		templateName = fileInfo.Name()
		ext := path.Ext(templateName)
		if ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR+"/"+templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName[0:len(templateName)-len(ext)]] = t
	}
}
func main() {
	mux :=http.NewServeMux()
	staticDirHandler(mux,"/assets","./public",0)
	
	// mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/", listHandler)

	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ",err.Error())
	}
}  