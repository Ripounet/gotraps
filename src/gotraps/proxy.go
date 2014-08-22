package gotraps

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"html"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

//
// Just forward the response from another site, to the user.
// Used to bypass the "Access-Control-Allow-Origin" for Go snippet compilation+execution.
//
func compile(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	remoteUrl := "http://play.golang.org/compile"
	trapcodeEscaped := r.FormValue("body")
	//c.Infof("%v", trapcodeEscaped) 
	trapcode := html.UnescapeString(trapcodeEscaped)
	values := url.Values{
		"version": []string{ "2" },
		"body": []string{ trapcode },
	}
	//c.Infof("%v", values)

	resp, err := urlfetch.Client(c).PostForm(remoteUrl, values)
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

func sendJsonError(w http.ResponseWriter, err error){
	w.WriteHeader( 500 )
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