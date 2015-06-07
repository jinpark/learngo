package main

import (
        // "fmt"
        "github.com/gorilla/mux"
        "net/http"
        // "net/url"
        // "io/ioutil"
        "image"
        "log"
        "image/jpeg"
        "bytes"
)

// func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
//         w.Write([]byte("Hello, World!"))
// }

// func Greet(w http.ResponseWriter, r *http.Request) {
//         name := mux.Vars(r)["name"]
//         w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
// }

// func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {

//         // break down the variables for easier assignment
//         vars := mux.Vars(r)
//         name := vars["name"]
//         job := vars["job"]
//         age := vars["age"]
//         w.Write([]byte(fmt.Sprintf("Name is %s ", name)))
//         w.Write([]byte(fmt.Sprintf("Job is %s ", job)))
//         w.Write([]byte(fmt.Sprintf("Age is %s ", age)))
// }

func getImage(url string) (image.Image, error) {  
	var img image.Image
    client := new(http.Client)
    req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if resp, err := client.Do(req); err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()

		if img, _, err = image.Decode(resp.Body); err != nil {
			return nil, err
		}
	}
	
	return img, nil
}

func writeImage(w http.ResponseWriter, r *http.Request) {
	url := mux.Vars(r)["url"]
	img, err := getImage(url)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
	}
    buffer := new(bytes.Buffer)
    if err := jpeg.Encode(buffer, img, nil); err != nil {
        log.Println("unable to encode image.")
    }

    w.Header().Set("Content-Type", "image/jpeg")
    // w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
    if _, err := w.Write(buffer.Bytes()); err != nil {
        log.Println("unable to write image.")
    }
}

// func serveImage(w http.ResponseWriter, r *http.Request, img Executor) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-type", "image/jpeg")
// 	// w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%v, s-maxage=%v", maxage, maxage))
// 	// w.Header().Set("Last-Modified", r.Header.Get("Last-Modified"))

// 	img.Execute(w)
// }


func main() {
        mx := mux.NewRouter()

        mx.HandleFunc("/{url}", writeImage)

        //to handle URL like
        //http://website:8080/person/Boo/CEO/199

        //http://website:8080/person/Boo/CEO/199 <- if age > 199, will cause 404 error

        // mx.HandleFunc("/person/{name}/{job}/{age:[0-199]+}", ProcessPathVariables)

        http.ListenAndServe(":8080", mx)
}