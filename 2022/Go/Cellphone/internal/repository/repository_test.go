package repository_test

import (
	"bytes"
	"cellphone/internal/api"
	"cellphone/internal/app_config"
	"cellphone/internal/dbconn"
	"cellphone/internal/entity"
	"cellphone/internal/entry"
	"cellphone/internal/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func insertCellphone(port int, cellphone *entity.Cellphone) (*http.Response, error) {
	jsonValue, err := json.Marshal(cellphone)

	if err != nil {
		return nil, err
	}

	return http.Post(fmt.Sprintf("http://localhost:%d/Cellphone", port), "application/json", bytes.NewBuffer(jsonValue))
}

func serveSingleCellphone(port int, providerId int) (*http.Response, error) {
	return http.Get(fmt.Sprintf("http://localhost:%d/Cellphone/%d", port, providerId))
}

func TestCellphoneSingle(t *testing.T) {
	config := app_config.Main{
		DbType:   dbconn.DB_MOCK,
		RepoType: repository.REPO_MOCK,
		ApiType:  api.API_NAT,
	}
	_, _, apiEngine, err := entry.InitializeBackend(config)

	go func() {
		switch engine := apiEngine.(type) {
		case *gin.Engine:
			err = engine.Run()
		case *api.NativeApiService:
			err = http.ListenAndServe(":8080", nil)
		}
	}()

	time.Sleep(2_000_000_000)

	if err != nil {
		t.Fatalf("Error is not nil: %s", err.Error())
	}

	var cellphone entity.Cellphone
	for i := 1; i <= 100; i++ {
		cellphone.Number = strconv.Itoa(i)
		cellphone.ProviderId = i % 10
		resp, err := insertCellphone(8080, &cellphone)

		if err != nil {
			t.Fatalf("insertCellphone Failed: %s", err.Error())
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("insertCellphone status returned: %s", http.StatusText(resp.StatusCode))
		}

		resp.Body.Close()
	}

	for i := 1; i <= 100; i++ {
		resp, err := serveSingleCellphone(8080, i%10)

		if err != nil {
			t.Fatalf("insertCellphone Failed: %s", err.Error())
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("insertCellphone status returned: %s", http.StatusText(resp.StatusCode))
		}

		bytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			t.Fatalf("Failed to read body: %s", err.Error())
		}

		err = json.Unmarshal(bytes, &cellphone)

		if err != nil {
			t.Fatalf("Failed to unmarshall: %s", err.Error())
		}

		i, err := strconv.Atoi(cellphone.Number)

		if err != nil {
			t.Fatalf("Failed to convert number %s", cellphone.Number)
		}

		if i < 0 || i > 100 {
			t.Fail()
		}

		resp.Body.Close()
	}

	resp, err := serveSingleCellphone(8080, 1)

	if resp.StatusCode != http.StatusNotFound {
		t.Fail()
	}
}
