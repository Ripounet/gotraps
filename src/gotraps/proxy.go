package gotraps

import (
	"appengine"
	"appengine/urlfetch"
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
	
	remoteUrl := "http://play.golang.org/compile"
	trapcode := r.FormValue("body") 
	values := url.Values{
		"version": []string{ "2" },
		"body": []string{ trapcode },
	}
	//c.Infof("%v", values)

	resp, err := urlfetch.Client(c).PostForm(remoteUrl, values)
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	defer resp.Body.Close()
	x, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	_, err = w.Write(x)
	if err != nil {
		c.Errorf("%v", err)
		return
	}

}
