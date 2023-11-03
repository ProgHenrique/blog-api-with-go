package main

import (
	"fmt"

	"github.com/ProgHenrique/api-blog/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	var port string = "3000"
	fmt.Println(`Server running @` + port)
	app.Run(":" + port)
}
