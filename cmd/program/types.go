package program

// A Project contains the data for the project folder
// being created, and methods that help with that process
type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
	ProjectType  string
	DBDriver     string
	FrameworkMap map[string]Framework
	DBDriverMap  map[string]Driver
}

// A Framework contains the name and templater for a
// given Framework
type Framework struct {
	packageName []string
	templater   Templater
}

type Driver struct {
	packageName []string
	templater   DBDriverTemplater
}

// A Templater has the methods that help build the files
// in the Project folder, and is specific to a Framework
type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
	RoutesWithDB() []byte
	ServerWithDB() []byte
	TestHandler() []byte
}

type DBDriverTemplater interface {
	Service() []byte
	Env() []byte
}

var (
	chiPackage     = []string{"github.com/go-chi/chi/v5"}
	gorillaPackage = []string{"github.com/gorilla/mux"}
	routerPackage  = []string{"github.com/julienschmidt/httprouter"}
	ginPackage     = []string{"github.com/gin-gonic/gin"}
	fiberPackage   = []string{"github.com/gofiber/fiber/v2"}
	echoPackage    = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}

	mysqlDriver    = []string{"github.com/go-sql-driver/mysql"}
	postgresDriver = []string{"github.com/lib/pq"}
	sqliteDriver   = []string{"github.com/mattn/go-sqlite3"}
	mongoDriver    = []string{"go.mongodb.org/mongo-driver"}

	godotenvPackage = []string{"github.com/joho/godotenv"}
)

const (
	root                 = "/"
	cmdApiPath           = "cmd/api"
	internalServerPath   = "internal/server"
	internalDatabasePath = "internal/database"
	testHandlerPath      = "tests"
	standardLib          = "standard = library"
)
