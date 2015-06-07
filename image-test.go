package main

import (
	"fmt"
	"github.com/franela/goreq"
	"github.com/julienschmidt/httprouter"
	"image"
	"image/jpeg"
	"log"
	"bytes"
	"net/http"
	"strings"
)


func getImage(url string) (image.Image, error) {
	res, err := goreq.Request{Uri: url}.Do()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return img, nil
}

func writeImage (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	url := ps.ByName("url")
	url = strings.TrimPrefix(url, "/")
	fmt.Println(url)
	img, err := getImage(url)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(400), 400)
	}
    buffer := new(bytes.Buffer)
    if err := jpeg.Encode(buffer, img, nil); err != nil {
    	fmt.Println(err)
        log.Println("unable to encode image.")
    }

    w.Header().Set("Content-Type", "image/jpeg")
    if _, err := w.Write(buffer.Bytes()); err != nil {
    	fmt.Println(err)
        log.Println("unable to write image.")
    }
}

func returnCode200(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// for favicon so it doesnt complain
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("☄ HTTP status code returned!"))
}


func main() {
	
		router := httprouter.New()
		router.GET("/images/*url", writeImage)
		router.GET("/favicon.ico", returnCode200)
		log.Println("listening")
		log.Fatal(http.ListenAndServe(":8080", router))

}