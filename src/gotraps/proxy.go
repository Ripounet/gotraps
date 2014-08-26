package gotraps

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"net/url"
)

//
// Just forward the response from another site, to the user.
// Used to bypass the "Access-Control-Allow-Origin" for Go snippet compilation+execution.
//
func compile(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	trapcodeEscaped := r.FormValue("body")
	//c.Infof("%v", trapcodeEscaped)
	trapcode := html.UnescapeString(trapcodeEscaped)
	values := url.Values{
		"version": []string{"2"},
		"body":    []string{trapcode},
	}
	//c.Infof("%v", values)

	resp, err := post(c, values)
	c.Infof("%v", resp)
	if err != nil {
		c.Errorf("%v", err)
		sendJsonError(w, err)
		return
	}
	defer resp.Body.Close()
	x, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("%v", err)
		sendJsonError(w, err)
		return
	}
	_, err = w.Write(x)
	if err != nil {
		c.Errorf("%v", err)
		sendJsonError(w, err)
		return
	}

}

const REMOTE_PLAYGROUND_COMPILE_URL = "http://play.golang.org/compile"
const GOTRAPS_UNIQUE_USER_AGENT = "go-traps.appspot.com"

// This works but doesn't explicitly add User-Agent header
//
// However according to https://developers.google.com/appengine/docs/go/urlfetch/#headers_identifying_request_source :
// "User-Agent. This header can be modified but App Engine will append an 
// identifier string to allow servers to identify App Engine requests. 
// The appended string has the format "AppEngine-Google; (+http://code.google.com/appengine; appid: APPID)", 
// where APPID is your app's identifier."
func post(c appengine.Context, values url.Values) (*http.Response, error) {
	return urlfetch.Client(c).PostForm(REMOTE_PLAYGROUND_COMPILE_URL, values)
}

// This does not work: http.DefaultTransport and http.DefaultClient are not available in App Engine.
func postWithUserAgent(c appengine.Context, values url.Values) (*http.Response, error) {
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
