package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   int
	mux sync.Mutex
} // Inc increments the counter.

var counterMux = SafeCounter{v: 0}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	server := &http.Server{Addr: ":3000", Handler: mux}
	fmt.Println("Started Server")
	log.Fatal(server.ListenAndServe())
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Hello"))
	r.Body.Close()
	counterMux.Inc()
	fmt.Println(counterMux.Value())

}

func check(e error) { //Simple Support function so we dont have to have so much boilerplate. This is not the best practice tbh
	if e != nil {
		panic(e)
	}
}

func (c *SafeCounter) Inc() {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the value c.v.
	c.v++
	c.mux.Unlock()
} // Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the value c.v.
	defer c.mux.Unlock()
	return c.v
}
