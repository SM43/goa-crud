package blogapi

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	// blank
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	blog "github.com/sm43/goa-crud/gen/blog"
	"github.com/stretchr/testify/assert"
)

var newLock sync.Mutex

// NewGormDB ...
func NewGormDB() (*gorm.DB, func() error) {
	newLock.Lock()
	defer newLock.Unlock()
	defaultDb, err := sql.Open("postgres", "user=postgres password=postgres dbname=goa_crud sslmode=disable")
	if err != nil {
		fmt.Println("Failed")
	}
	testDbName := fmt.Sprintf("%s_test_%d", "goa_crud", time.Now().UnixNano())

	_, err = defaultDb.Exec(fmt.Sprintf("CREATE DATABASE %s;", testDbName))
	if err != nil {
		fmt.Println("Failed to create db", err)
	}

	testDbString := fmt.Sprintf("user=postgres password=postgres dbname=%s sslmode=disable", testDbName)
	db, err := gorm.Open("postgres", testDbString)
	if err != nil {
		fmt.Println("Failed to connect")
	}

	db.AutoMigrate(Blog{}, Comment{})

	closeFn := func() error {
		_ = db.Close()
		_, err = defaultDb.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDbName))
		if err != nil {
			fmt.Println("Error ", err)
			return err
		}
		return defaultDb.Close()
	}
	return db, closeFn
}

// LoadFixtureDirs ...
func LoadFixtureDirs(db *gorm.DB, fixtureDir string) error {

	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB()),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		fmt.Println("Failed to create new Fixture", err)
	}
	if err := fixtures.Load(); err != nil {
		fmt.Println("Can't Load Data", err)
	}
	return nil
}

func TestBlog_Create(t *testing.T) {
	t.Run("create blog", func(t *testing.T) {
		t.Parallel()
		db, closeDb := NewGormDB()
		defer closeDb()
		logger := log.New(os.Stderr, "[blogapi] ", log.Ltime)
		blogSvc := NewBlog(db, logger)
		b := &blog.Blog{Name: "shivam"}
		err := blogSvc.Create(context.Background(), b)
		assert.NoError(t, err)
	})
}

func TestBlog_List(t *testing.T) {
	t.Run("read blogs", func(t *testing.T) {
		t.Parallel()
		db, closeDb := NewGormDB()
		LoadFixtureDirs(db, "fixtures")
		defer closeDb()
		logger := log.New(os.Stderr, "[blogapi] ", log.Ltime)
		blogSvc := NewBlog(db, logger)
		blogList, err := blogSvc.List(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, len(blogList), 3)
	})
}
