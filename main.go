package main

import (
	"flag"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tamnook/test_issue/queue"
)

var N int

func main() {

	flagN := flag.String("N", "1", "")
	flag.Parse()
	var err error
	N, err = strconv.Atoi(*flagN)
	if err != nil {
		panic(err)
	}

	queue := queue.InitQueue()

	server := echo.New()
	server.POST("/addIssue", func(c echo.Context) error {

		return c.String(http.StatusOK, "")
	})
	server.Logger.Fatal(server.Start(":8080"))
}

func ArithmeticProgression(n int, d, n1, I float64) float64 {
	ans := n1
	for i := 0; i < n-1; i++ {
		time.Sleep(time.Duration(int(I*1000)) * time.Millisecond)
		ans += d
	}
	return ans
}
