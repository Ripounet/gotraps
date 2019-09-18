package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/compile", compile)

	// Handlers for static files: necessary for local dev.
	// Might be bypassed/ignored by app.yaml in prod.
	// http.Handle("/static/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static/workaround/assets"))))
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir("content"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}

//
// Just forward the response from another site, to the user.
// Used to bypass the "Access-Control-Allow-Origin" for Go snippet compilation+execution.
//
func compile(w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	trapcodeEscaped := r.FormValue("body")
	trapname := r.FormValue("trapname")
	log.Printf("Compile [%v]", trapname)

	trapcode := html.UnescapeString(trapcodeEscaped)
	values := url.Values{
		"version": []string{"2"},
		"body":    []string{trapcode},
	}
	//c.Infof("%v", values)

	resp, err := post(c, values)
	//c.Infof("%v", resp)
	if err != nil {
		log.Printf("%v", err)
		sendJsonError(w, err)
		return
	}
	defer resp.Body.Close()
	x, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v", err)
		sendJsonError(w, err)
		return
	}
	_, err = w.Write(x)
	if err != nil {
		log.Printf("%v", err)
		sendJsonError(w, err)
		return
	}
}

const REMOTE_PLAYGROUND_COMPILE_URL = "https://play.golang.org/compile"
const GOTRAPS_UNIQUE_USER_AGENT = "go-traps.appspot.com"

// This works but doesn't explicitly add User-Agent header
//
// However according to https://developers.google.com/appengine/docs/go/urlfetch/#headers_identifying_request_source :
// "User-Agent. This header can be modified but App Engine will append an
// identifier string to allow servers to identify App Engine requests.
// The appended string has the format "AppEngine-Google; (+http://code.google.com/appengine; appid: APPID)",
// where APPID is your app's identifier."
//
// After further investigation: the header value is "AppEngine-Google; (+http://code.google.com/appengine; appid: s~go-traps)"
func post(c context.Context, values url.Values) (*http.Response, error) {
	return http.PostForm(REMOTE_PLAYGROUND_COMPILE_URL, values)
}

// This does not work: http.DefaultTransport and http.DefaultClient are not available in App Engine.
func postWithUserAgent(c context.Context, values url.Values) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", REMOTE_PLAYGROUND_COMPILE_URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", GOTRAPS_UNIQUE_USER_AGENT)

	return client.Do(req)
}

func sendJsonError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	fmt.Fprint(w, Response{"Errors": err.Error(), "Events": nil})
}

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}
