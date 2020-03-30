package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	//"image"
	//"image/color"
	//"io"
	"github.com/disintegration/imaging"
	"github.com/jiandahao/servemux"
	"io/ioutil"
	"log"
	//"net/http"
	"os"
)

func main(){
	router := servemux.NewRouter()
	router.Post("/",handleImage)
	compressImage()
	http.ListenAndServe(":12345",http.HandlerFunc(router.ServeHTTP))
}

func handleImage(w http.ResponseWriter, r *http.Request){
	if err := r.ParseMultipartForm(1<<20);err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	image, _,_ := r.FormFile("image")
	//fileBytes,_ := ioutil.ReadAll(image)

	//img,_ := imaging.Decode(image)
	//img = imaging.Resize(img,200,200,imaging.Lanczos)
	//fileBytes := &bytes.Buffer{}
	//imaging.Encode(fileBytes,img,imaging.JPEG)
	Bytes,_ := compress(image,1000,1000)
	ioutil.WriteFile("http_out_image",Bytes,os.ModePerm)
}
func compress(r io.Reader, width int, height int)([]byte,error){
	img,_ := imaging.Decode(r)
	img = imaging.Resize(img,width,height,imaging.Lanczos)
	imageBytes := &bytes.Buffer{}
	imaging.Encode(imageBytes,img, imaging.JPEG)
	return imageBytes.Bytes(),nil
}
func compressImage(){
	img, err := imaging.Open("./cloth.jpeg")
	fmt.Println(img.Bounds().Min)
	fmt.Println(img.Bounds().Max)
	if err != nil{
		fmt.Println(err)
		return
	}
	time1 := time.Now()
	img = imaging.Resize(img,200,200,imaging.Lanczos)
	time2 := time.Now()
	fmt.Println(time2.Sub(time1))
	buf := &bytes.Buffer{}
	imaging.Encode(buf,img,imaging.JPEG)
	if err := ioutil.WriteFile("./compress_out.jpeg",buf.Bytes(),os.ModePerm);err != nil{
		log.Println(err)
		return
	}

	//dst := imaging.New(100, 100, color.NRGBA{0, 0, 0, 0})
	//dst = imaging.Paste(dst, img, image.Pt(0, 0))
	//err = imaging.Save(dst, "cloth_compress.jpg")
	//if err != nil {
	//	log.Fatalf("failed to save image: %v", err)
	//}
}