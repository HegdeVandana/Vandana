// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (//"fmt"
         //"database/sql"
        //"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"crypto/tls"
	"net/http"
	//"sync"
	"sync/atomic"
       
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/models"
	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi/operations"
	"github.com/go-swagger/go-swagger/examples/tutorials/todo-list/server-complete/restapi/operations/todos"
)




//go:generate swagger generate server --target .. --name TodoList --spec ../swagger.yml
 //var  items = make(map[int64]*models.Item)
 var lastID int64

 
//var itemsLock = &sync.Mutex{}
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&models.Profile{}) {
		db.CreateTable(&models.Profile{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Profile{})
              
	}
return db        
}


        

func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}
        





func getProfile(since int64, limit int32) (result[] *models.Profile){
         
        // Connection to the database
	db := InitDb()
	// Close connection database
	 
       defer db.Close()
       var profiles []*models.Profile
        
      db.Find(&profiles)
    
     
        
       
        
	
	return profiles 
}        

func getidProfile(id int64) (profile *models.Profile){        
    
       
        // Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()
       var profiles []*models.Profile
       db.First(&profiles,id)

    for _, profile:= range profiles{
    if profile.ID==id{
       
          return profile

}
}
return
}
        
       


func configureFlags(api *operations.TodoListAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TodoListAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	
	api.TodosFindTodosHandler = todos.FindTodosHandlerFunc(func(params todos.FindTodosParams) middleware.Responder {
		mergedParams := todos.NewFindTodosParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return todos.NewFindTodosOK().WithPayload(getProfile(*mergedParams.Since, *mergedParams.Limit))
	})
        api.TodosOnetodosHandler = todos.OnetodosHandlerFunc(func(params todos.OnetodosParams) middleware.Responder {

                if err:=getidProfile(params.ID);err == nil {
			return todos.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500})
		}
                 

		return todos.NewOnetodosOK().WithPayload(getidProfile(params.ID))


	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}





