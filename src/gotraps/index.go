// Go Traps server-side processing (if ever needed)
package gotraps

import "net/http"

func init() {
	http.HandleFunc("/compile", compile)
}
