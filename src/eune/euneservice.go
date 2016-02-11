package main

import (
	"compress/gzip"
	"net/http"
	"fmt"
	"github.com/golang/glog"
	"dp/dpds"
	"strings"
)
	
// euneHandler handles all route requests for eune requests
// If there is a dot at the end of the route, that dot will be processed and returned.
func euneHandler(responseWriter http.ResponseWriter, req *http.Request) {
    responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    responseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	if req.Method == "OPTIONS" {
		responseWriter.Header().Set("Allow", "GET,OPTIONS")
		responseWriter.WriteHeader(http.StatusOK)
		return
	}
	
	// Route used.
	route := req.URL.Path
	paramMap := req.URL.Query()
    
    routeChunks := strings.Split(route, "/")
    if len(routeChunks) <= 1 {
		responseWriter.WriteHeader(http.StatusNotFound)
    	return
    }
	firstRoute := "/" + routeChunks[1]
	dottree := dpds.DtFactory.GetInstance()

    rd := new(dpds.RequestDot)
    rd.Init(paramMap, firstRoute, route, "")
    firstDot := dottree.GetDot(firstRoute)
    firstDot.Process(rd)

	responseWriter.Header().Add("Content-Type", "text/plain")
	responseWriter.Header().Add("Content-Encoding", "gzip")

    var jsonResults = rd.GetResult()

    if len(jsonResults) == 0 {
		responseWriter.WriteHeader(http.StatusNotFound)
    	return
    } else {
        //
        // Ready to generate output back to client.
	    //
	    gzipWriter := gzip.NewWriter(responseWriter)
        fmt.Fprintf(gzipWriter, jsonResults)
        gzipWriter.Flush()
    }
}

// On execution of the web service, set up all available routes to DOT's in the DOT tree.
func main() {
	glog.Error("Starting up.")
	dottree := dpds.DtFactory.GetInstance()
	routeMap := dottree.GenerateRoutes()
	
    for route, _ := range routeMap {
    	http.HandleFunc(route, euneHandler)
    }
	
	http.ListenAndServe(":8080", nil)
}
