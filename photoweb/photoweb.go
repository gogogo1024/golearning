package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)
const (UPLOAD_DIR = "./uploads")
const (TEMPLATE_DIR = "./views")

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
		f,h,err := r.FormFile("image")
		check(err)
		filename :=h.Filename
		defer f.Close()
		t, err :=os.Create(UPLOAD_DIR+"/"+filename)
		check(err)
		defer t.Close()
		if _,err :=io.Copy(t,f); err != nil {
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
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
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ",err.Error())
	}
}  