package server

func RunServer() {
	app := NewApp()

	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
