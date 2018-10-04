package main // import "server"

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"dbbooks"
)

var table = "novel"

// ResultResponse : Create, Read 결과 반환용
type ResultResponse struct{ Message string }

func dbInit(c echo.Context) error {
	dbbooks.CreateTable(table)
	return c.JSON(http.StatusOK, &ResultResponse{Message: "Table creation done"})
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, dbbooks.SelectData(0, table))
}

func create(c echo.Context) error {
	book := new(dbbooks.Book)
	book.Title = c.FormValue("Title")
	book.Author = c.FormValue("Author")

	dbbooks.InsertData(book, table)
	return c.JSON(http.StatusOK, &ResultResponse{Message: "New item creation done"})
}

func read(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("ID"))
	return c.JSON(http.StatusOK, dbbooks.SelectData(_id, table))
}

func update(c echo.Context) error {
	book := new(dbbooks.Book)
	if e := c.Bind(book); e != nil {
		panic(e.Error())
	}

	dbbooks.UpdateData(book, table)

	return c.JSON(http.StatusOK, &ResultResponse{Message: "Update done"})
}

func delete(c echo.Context) error {
	// Error is occured so, auth is waived.
	// auth := echo.Map{}
	// if e := c.Bind(&auth); e != nil {
	// 	panic(e.Error())
	// }

	_id, _ := strconv.Atoi(c.Param("ID"))

	fmt.Println(_id)

	dbbooks.DeleteData(_id, table)
	return c.JSON(http.StatusOK, &ResultResponse{Message: "Delete done"})
}

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResult := &ResultResponse{
			Message: "Contents not found",
		}
		return c.JSON(http.StatusNotFound, errorResult)
	}

	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/dbinit", dbInit)

	e.GET("/books", index)

	e.POST("/books", create)
	e.GET("/books/:id", read)
	e.PUT("/books", update)
	e.DELETE("/books/:id", delete)

	// e.Logger.Fatal(e.Start("127.0.0.1:1323"))
	e.Logger.Fatal(e.Start(":1323"))
}
