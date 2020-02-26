// string service microservice using go kit
package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Transport expose the service to network. In this first example we utilize JSON over http
func main() {

	logger := log.NewLogfmtLogger(os.Stdout)

	var svc StringService
	svc = stringService{}

	svc = loggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	http.ListenAndServe(":9000", nil)
	//log.Fatal(http.ListenAndServe(":9000",nil))

}
