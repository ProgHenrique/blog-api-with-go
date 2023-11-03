package app

import (
	"database/sql"
	"net/http"

	createpost "github.com/ProgHenrique/api-blog/app/modules/Posts/useCases/createPost"
	deletepost "github.com/ProgHenrique/api-blog/app/modules/Posts/useCases/deletePost"
	getpost "github.com/ProgHenrique/api-blog/app/modules/Posts/useCases/getPost"
	getposts "github.com/ProgHenrique/api-blog/app/modules/Posts/useCases/getPosts"
	updatepost "github.com/ProgHenrique/api-blog/app/modules/Posts/useCases/updatePost"
	"github.com/ProgHenrique/api-blog/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// App has Router and Database instances
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// RequestHandlerFunction handles the requests
type RequestHandlerFunction func(db *sql.DB, w http.ResponseWriter, r *http.Request)

func (app *App) Initialize() {
	db, err := sql.Open(config.Provider, config.DATABASE_URL)
	if err != nil {
		panic(err)
	}

	app.DB = db
	app.Router = mux.NewRouter()
	app.setRouters()
}

func (app *App) setRouters() {
	app.Post("/posts/create", app.handleRequest(createpost.CreatePostController))
	app.GET("/posts/get-all", app.handleRequest(getposts.GetPostsController))
	app.GET("/posts/{_id}", app.handleRequest(getpost.GetPostController))
	app.DELETE("/posts/{_id}/delete", app.handleRequest(deletepost.DeletePostController))
	app.PUT("/posts/{_id}/update", app.handleRequest(updatepost.UpdatePostController))
}

func (app *App) Post(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) GET(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) DELETE(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("DELETE")
}

func (app *App) PUT(path string, f func(res http.ResponseWriter, req *http.Request)) {
	app.Router.HandleFunc(path, f).Methods("PUT")
}

func (app *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		handler(app.DB, res, req)
	}
}

func (app *App) Run(port string) {
	http.ListenAndServe(port, app.Router)
}
