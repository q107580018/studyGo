package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"text/template"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	for _, fileInfo := range fileInfoArr {
		templateName := fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath := path.Join(TEMPLATE_DIR, templateName)
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t

	}
}
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e)
				// logging
				log.Printf("WARN: panic in %v - %v\n", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	err = templates[tmpl].Execute(w, locals)
	return
}
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := renderHtml(w, "upload.html", nil)
		check(err)
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(path.Join(UPLOAD_DIR, filename))
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := path.Join(UPLOAD_DIR, imageId)
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func listhandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	err = renderHtml(w, "list.html", locals)
	check(err)
}
func main() {

	http.HandleFunc("/", safeHandler(listhandler))
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
