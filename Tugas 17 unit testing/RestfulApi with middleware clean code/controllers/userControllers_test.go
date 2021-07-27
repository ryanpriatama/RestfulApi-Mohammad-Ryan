package controllers

import (
	"encoding/json"
	//"myapp/project/config"
	"myapp/project/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	//"github.com/stretchr/testify/assert"
)

func testGetBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []models.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 1 {
		t.Errorf("expected one book, got: %#v", books)
		return
	}
	if books[0].Name != "Harry Potter" {
		t.Errorf("expected Harry Potter, got: %#v", books[0].Name)
	}
}

func testPostBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []models.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 1 {
		t.Errorf("expected one book, got: %#v", books)
		return
	}
	if books[0].Name != "Harry Potter" {
		t.Errorf("expected Harry Potter, got: %#v", books[0].Name)
	}
}

func testDeleteBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []models.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 0 {
		t.Errorf("expected None book, got: %#v", books)
		return
	}
}

func testUpdateBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []models.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 1 {
		t.Errorf("expected one book, got: %#v", books)
		return
	}
	if books[0].Name != "LOTR" {
		t.Errorf("expected Harry Potter, got: %#v", books[0].Name)
	}
}

func TestMockGetBookController(t *testing.T) {
	m := models.NewMockBookModel()
	bookController := CreateGetBookController(m)
	// insert data baru
	m.Insert(models.Book{Name: "Harry Potter"})
	// test controller
	testGetBookController(t, bookController)
}

func TestMockPostBookController(t *testing.T) {
	m := models.NewMockBookModel()
	bookController := CreatePostBookController(m)
	// insert data baru
	m.Insert(models.Book{Name: "Harry Potter"})
	// test controller
	testPostBookController(t, bookController)
}

func TestMockDeleteBookController(t *testing.T) {
	m := models.NewMockBookModel()
	bookController := CreateDeleteBookController(m)
	// insert data baru
	m.Insert(models.Book{Name: "Harry Potter"})
	// test controller
	testDeleteBookController(t, bookController)
}

func TestMockUpdateBookController(t *testing.T) {
	m := models.NewMockBookModel()
	bookController := CreateUpdateBookController(m)
	// insert data baru
	m.Insert(models.Book{Name: "Harry Potter"})
	m.Insert(models.Book{Name: "LOTR"})
	// test controller
	testUpdateBookController(t, bookController)
}

/*
func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	users := models.User{
		Name:     "ryan",
		Password: "111",
		Email:    "ryan@gmail.com",
	}

	var err error
	if err = config.DB.Save(&users).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUserController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUserControllers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// open file
			// convert struct
			var users models.UserResponse
			err := json.Unmarshal([]byte(body), &users)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(users.Data))
		}
	}
}
*/
