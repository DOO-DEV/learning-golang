package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Context is an object that can be safely passed to multiple
// goroutines to provide a way to implement a timeout or cancellation to a function.
// 3rd party libraries that make HTTP requests or database requests
// typically have support for providing your own context so you cancel those operations.

// func main() {
// 	// if time limit that we define exceeded wil cancel the request
// 	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	// create HTTP request
// 	req, err := http.NewRequestWithContext(timeoutContext,
// 		http.MethodGet, "https://dev-to-uploads.s3.amazonaws.com/uploads/logos/resized_logo_UQww2soKuUsjaOGNB38o.png", nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// preforme HTTP request
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer res.Body.Close()

// 	// get data from HTTP request
// 	imageData, err := ioutil.ReadAll(res.Body)

// 	fmt.Printf("Downloaded image of size %d\n", len(imageData))

// 	// context with gin server
// 	Server()
// }

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		// one of the things that ctx.reques.context does is, if client cancel request it cancle
		// the operation in our program to for preventing run remain of our logic
		
		// also we can add our timeout context to it
		timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second *5)
		defer cancel()

		req, err := http.NewRequestWithContext(timeoutContext,http.MethodGet,"https://google.com", nil)
		if err != nil {
			panic(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		ctx.Data(200, "text/html", data)

	})
	r.Run(":3000")
}