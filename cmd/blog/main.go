package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
	// Blank for package side effect: loads postgres drivers
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	blogapi "github.com/sm43/goa-crud"
	blog "github.com/sm43/goa-crud/gen/blog"
	oauth "github.com/sm43/goa-crud/gen/oauth"
	user "github.com/sm43/goa-crud/gen/user"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[blogapi] ", log.Ltime)
	}

	{
		// loads values from .env into the system
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
			return
		}
	}
	// Database Connection
	var (
		db *gorm.DB
	)
	{
		var err error
		db, err = gorm.Open("postgres", "user=postgres password=postgres dbname=goa_crud sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successful Db Connection..!!")
		defer db.Close()
		db.LogMode(true)
		db.AutoMigrate(blogapi.Blog{}, blogapi.Comment{}, blogapi.User{})
	}

	// Initialize the services.
	var (
		oauthSvc oauth.Service
		blogSvc  blog.Service
		userSvc  user.Service
	)
	{
		oauthSvc = blogapi.NewOauth(db, logger)
		blogSvc = blogapi.NewBlog(db, logger)
		userSvc = blogapi.NewUser(db, logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		oauthEndpoints *oauth.Endpoints
		blogEndpoints  *blog.Endpoints
		userEndpoints  *user.Endpoints
	)
	{
		oauthEndpoints = oauth.NewEndpoints(oauthSvc)
		blogEndpoints = blog.NewEndpoints(blogSvc)
		userEndpoints = user.NewEndpoints(userSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8000"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, oauthEndpoints, blogEndpoints, userEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
