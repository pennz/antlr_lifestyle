package main

// Import resty into your code and refer it as `resty`.
import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/pennz/antlr_lifestyle/lifestyle"
	"github.com/pennz/antlr_lifestyle/logic"
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
	env, err := model.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	go logic.FeedbackLoop("test-logs_topic", []string{"anonymous.info"})

	r := setupRouter()
	addModelFunc2Router(r, env)
	r.Run(":" + port)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

// parse_do call python? to parse
func parse_do(deeds []string) error {
	log.Println(deeds)
	return nil
}

func get_data(c *gin.Context) (map[string][]string, error) {
	env, err := model.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	q := c.Request.URL.Query()
	v, ok := q["do"]
	if ok {
		parse_do(v)
		err = env.DS.AddAction(&lifestyle.Action{ActionType: "TestType"})
	}
	if err != nil {
		log.Println(err)
	} else {
		log.Println("AddAction is fine")
	}

	return q, err
}

func addModelFunc2Router(r *gin.Engine, env model.Env) {
	r.GET("/actions", func(c *gin.Context) {
		ts, err := env.DS.AllActions()
		if err != nil {
			c.String(403, "403 Error")
		}
		s := ""
		for _, t := range ts {
			s += fmt.Sprintf("%v\n", t)
		}
		c.String(200, s)
	})

	r.GET("/api", func(c *gin.Context) {
		ts, err := get_data(c)

		if err != nil {
			c.String(403, "403 Error")
		}
		s := ""
		for _, t := range ts {
			s += fmt.Sprintf("%v\n", t)
		}

		c.String(200, s)
	})

	r.GET("/relations", func(c *gin.Context) {
		ts, err := env.DS.AllRelations()
		if err != nil {
			c.String(403, "403 Error")
		}
		s := ""
		for _, t := range ts {
			s += fmt.Sprintf("%v\n", t)
		}
		c.String(200, s)
	})
	r.GET("/thing", func(c *gin.Context) {
		t, err := env.DS.Thing("good")
		if err != nil {
			c.String(403, "403 Error")
		}
		s := ""
		s += fmt.Sprintf("%v\n", t)
		c.String(200, s)
	})
	r.GET("/things", func(c *gin.Context) {
		ts, err := env.DS.AllThings()
		if err != nil {
			c.String(403, "403 Error")
		}
		s := ""
		for _, t := range ts {
			s += fmt.Sprintf("%v\n", t)
		}
		c.String(200, s)
	})
}
