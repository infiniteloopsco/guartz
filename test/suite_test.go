package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"

	"github.com/infiniteloopsco/guartz/endpoint"
	"github.com/infiniteloopsco/guartz/models"
	"github.com/infiniteloopsco/guartz/utils"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ts     *httptest.Server
	client utils.Client
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	fmt.Println("Suite found")
	RunSpecs(t, "Api Suite")
}

var _ = BeforeSuite(func() {
	initEnv()
	utils.InitLogTest()
	models.InitDB()
	models.InitCron()
	cleanDB()
	ts = httptest.NewServer(endpoint.GetMainEngine())
	client = utils.Client{
		&http.Client{},
		ts.URL,
		"application/json",
	}
})

var _ = AfterSuite(func() {
	models.Gdb.Close()
	ts.Close()
})

var _ = BeforeEach(func() {
	cleanDB()
})

func cleanDB() {
	utils.Log.Info("***Cleaning***")
	models.Gdb.Delete(models.Execution{})
	models.Gdb.Delete(models.Task{})
}

func initEnv() {
	path := ".env_test"
	for i := 1; ; i++ {
		if err := godotenv.Load(path); err != nil {
			if i > 3 {
				panic("Error loading .env_test file")
			} else {
				path = "../" + path
			}
		} else {
			break
		}
	}
}
