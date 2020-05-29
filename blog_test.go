package blogapi

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	// blank
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	blog "github.com/sm43/goa-crud/gen/blog"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	db      *gorm.DB
	closeDb func() error
	blogSvc blog.Service
}

var test TestConfig

func TestMain(m *testing.M) {

	var err error
	if test, err = ConnectDb(); err != nil {
		fmt.Println("Failed to Connect to Db. Details - ", err)
		return
	}
	if err = LoadFixture(test.db, "fixtures"); err != nil {
		fmt.Println("Failed to load Data. Details - ", err)
		return
	}
	logger := log.New(os.Stderr, "[blogapi] ", log.Ltime)
	test.blogSvc = NewBlog(test.db, logger)

	defer os.Exit(m.Run())
	defer test.closeDb()

}

// ConnectDb ...
func ConnectDb() (TestConfig, error) {

	// Connect to Application Db
	appDBName := "goa_crud"
	appDbString := fmt.Sprintf("user=postgres password=postgres dbname=%s sslmode=disable", appDBName)
	defaultDb, err := sql.Open("postgres", appDbString)
	if err != nil {
		return TestConfig{}, err
	}

	// Create a test db - appDbname_test
	testDbName := fmt.Sprintf("%s_test", appDBName)
	_, err = defaultDb.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDbName))
	if err != nil {
		return TestConfig{}, err
	}
	_, err = defaultDb.Exec(fmt.Sprintf("CREATE DATABASE %s;", testDbName))
	if err != nil {
		return TestConfig{}, err
	}

	// Connect to test db
	testDbString := fmt.Sprintf("user=postgres password=postgres dbname=%s sslmode=disable", testDbName)
	db, err := gorm.Open("postgres", testDbString)
	if err != nil {
		return TestConfig{}, err
	}

	db.AutoMigrate(Blog{}, Comment{})

	closeDb := func() error {
		_ = db.Close()
		_, err = defaultDb.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", testDbName))
		if err != nil {
			fmt.Println("Error ", err)
			return err
		}
		return defaultDb.Close()
	}
	return TestConfig{db: db, closeDb: closeDb}, nil
}

// LoadFixture ...
func LoadFixture(db *gorm.DB, fixtureDir string) error {

	fixtures, err := testfixtures.New(
		testfixtures.Database(db.DB()),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixtureDir),
	)
	if err != nil {
		return err
	}
	if err := fixtures.Load(); err != nil {
		return err
	}
	return nil
}

// Create a blog
func Test_Create(t *testing.T) {
	LoadFixture(test.db, "fixtures")
	c := []*blog.Comment{
		{Comment: "2019"},
		{Comment: "Movie"},
	}
	b := &blog.Blog{Name: "Karwaan", Comments: c}
	payload := &blog.CreatePayload{Blog: b, Auth: ""}
	err := test.blogSvc.Create(context.Background(), payload)
	assert.NoError(t, err)

	// If the Blog already exist, expect error
	err = test.blogSvc.Create(context.Background(), payload)
	assert.Error(t, err)

}

// Fetch the list of the blog
func Test_List(t *testing.T) {
	LoadFixture(test.db, "fixtures")
	blogList, err := test.blogSvc.List(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, len(blogList), 3)
}

// Fetch a blog using ID
func Test_Show(t *testing.T) {
	LoadFixture(test.db, "fixtures")
	payload := &blog.ShowPayload{ID: 1}
	blog, err := test.blogSvc.Show(context.Background(), payload)
	assert.NoError(t, err)
	assert.Equal(t, blog.Name, "Why A is not B")
	assert.Equal(t, blog.Comments[0].Comment, "Bcoz A is A and B is B")

	// If the blog does not exist, expect error
	payload.ID = 11
	blog, err = test.blogSvc.Show(context.Background(), payload)
	assert.Error(t, err)
}

// Remove a blog using ID / If a blog doesn't exit it doesn't return error
// TODO: Check for record in code
func Test_Remove(t *testing.T) {
	LoadFixture(test.db, "fixtures")
	removePayload := &blog.RemovePayload{ID: 1}
	err := test.blogSvc.Remove(context.Background(), removePayload)
	assert.NoError(t, err)
}

// Add a comment to a blog using ID
func Test_Add(t *testing.T) {
	LoadFixture(test.db, "fixtures")
	c := blog.Comment{Comment: "2019"}
	addPayload := &blog.AddPayload{ID: 1, Comments: &c}
	err := test.blogSvc.Add(context.Background(), addPayload)
	assert.NoError(t, err)

	// If the blog does not exist, expect error
	addPayload.ID = 11
	err = test.blogSvc.Add(context.Background(), addPayload)
	assert.Error(t, err)
}
