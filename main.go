package main

// Import resty into your code and refer it as `resty`.
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pennz/antlr_lifestyle/lifestyle"
	"github.com/pennz/antlr_lifestyle/model"

	_ "github.com/heroku/x/hmetrics/onload"
)

func testResty() {
	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("Error      :", err)
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Status     :", resp.Status())
	fmt.Println("Time       :", resp.Time())
	fmt.Println("Received At:", resp.ReceivedAt())
	fmt.Println("Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("DNSLookup    :", ti.DNSLookup)
	fmt.Println("ConnTime     :", ti.ConnTime)
	fmt.Println("TLSHandshake :", ti.TLSHandshake)
	fmt.Println("ServerTime   :", ti.ServerTime)
	fmt.Println("ResponseTime :", ti.ResponseTime)
	fmt.Println("TotalTime    :", ti.TotalTime)
	fmt.Println("IsConnReused :", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime :", ti.ConnIdleTime)

	/* Output
			Response Info:
			Error      : <nil>
			Status Code: 200
			Status     : 200 OK
			Time       : 465.301137ms
			Received At: 2019-06-16 01:52:33.772456 -0800 PST m=+0.466672260
			Body       :
			 {
			   "args": {},
		     "headers": {
			     "Accept-Encoding": "gzip",
			     "Host": "httpbin.org",
			     "User-Agent": "go-resty/2.0.0 (https://github.com/go-resty/resty)"
			   },
		     "origin": "0.0.0.0",
		   "url": "https://httpbin.org/get"
	   }

	   Request Trace Info:
	   DNSLookup    : 2.21124ms
	   ConnTime     : 393.875795ms
	   TLSHandshake : 319.313546ms
	   ServerTime   : 71.109256ms
	   ResponseTime : 94.466µs
	   TotalTime    : 465.301137ms
	   IsConnReused : false
	   IsConnWasIdle: false
	   ConnIdleTime : 0s
	*/
}

func main() {
	testResty()
	env, err := model.GetFakeDB()
	if err != nil {
		log.Fatal(err)
	}
	ts, err := env.DS.AllThings()
	for _, t := range ts {
		fmt.Println(t)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	log.Println("Check before")
	log.Println("Preparation Done")
	log.Println(lifestyle.Actions)
	router.Run(":" + port)
	log.Fatal("Check, router is down")
}
