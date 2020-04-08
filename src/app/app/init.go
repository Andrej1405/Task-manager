package app

import (
	"github.com/revel/revel"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type project struct {
	id int
	projectname string
	date string
}

// type employee struct {
// 	id int
// 	surname string
// 	name string
// 	position string
// }
	
// type status struct {
// 	id int
// 	namestatus string
// }

var Db *sql.DB

func InitDB() (err error) {
	connect := "user=postgres password=1111 dbname=db_taskManager sslmode=disable"

	Db, err := sql.Open("postgres", connect)
	if err != nil {
		panic(err)
	}
	defer Db.Close()
	fmt.Println("DB Connected", Db)

rows, err := Db.Query("select * from projects")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

	projects := []project{}

	for rows.Next() {
		p := project{}
		
        err := rows.Scan(&p.id, &p.projectname, &p.date)
        if err != nil{
            panic(err)
		}
		
        projects = append(projects, p)
    }
	return
}

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	InitDB()

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
