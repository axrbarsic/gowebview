package basic

import (
	"fmt"
	"net/http"
	"time"

	"github.com/srinathh/mobilehtml5app/server"
	"golang.org/x/net/context"
)

const greetString = "greetstring"

var srv *server.Server

// Start is called by the native portion of the webapp to start the web server.
// It returns the server root URL (without the trailing slash) and any errors.
func Start() (string, error) {
	srv := server.NewServer()
	srv.HandleFunc(server.GET, "/", index)
	srv.HandleFunc(server.GET, "/:name", hello)
	return srv.Start("127.0.0.1:0", map[string]string{greetString: "Namaste"})
}

// Stop is called by the native portion of the webapp to stop the web server.
func Stop() {
	srv.Stop(time.Millisecond * 100)
}

// These two are autogenerated sample handlers for your webapp to get you started.

func index(_ context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<html><body><div><a href='/Alice'>Alice</a></div><div><a href='/Bob'>Bob</a></div></body></html>"))
}

func hello(c context.Context, w http.ResponseWriter, r *http.Request) {
	name := c.Value("name")
	greetstring := c.Value(greetString)
	fmt.Fprintf(w, "<html><body><div>%s %s!</div><div><a href='/'>Back</a></div></body></html>", greetstring, name)
}