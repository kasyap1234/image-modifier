package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()
	
	e.Start(":8080")

}

// go will ingest images into a queue another goroutine will pick images concurrenlty and process them ? 
