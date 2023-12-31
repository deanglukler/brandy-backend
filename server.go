package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Product struct {
    ID    int    `json:"id"`
    BrandName  string `json:"brand_name"`
    ProductName string `json:"product_name"`
    Category string `json:"category"`
    Location string `json:"location"`
    ImageURL string `json:"image_url"`
}

func main() {

    port := os.Getenv("PORT")
 
	if port == "" {
		port = "8080"
	}
    
    e := echo.New()

    e.GET("/products", getAllProducts)
    e.GET("/products/search", searchProducts)

    e.Logger.Fatal(e.Start(":" + port))
}

func getAllProducts(c echo.Context) error {
    log.Println("Getting all products...")
    
    connStr := os.Getenv("CONNSTR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

    log.Println("Connected to database!")

    rows, err := db.Query("SELECT id, brand_name, product_name, category, location, image_url FROM products")
    if err != nil {
        log.Fatal(err)
        return err
    }
    defer rows.Close()

    var products []Product

    for rows.Next() {
        var product Product
        if err := rows.Scan(&product.ID, &product.BrandName, &product.ProductName, &product.Category, &product.Location, &product.ImageURL); err != nil {
            return err
        }
        products = append(products, product)
    }

    return c.JSON(http.StatusOK, products)
}

func searchProducts(c echo.Context) error {
    query := c.QueryParam("q")

    connStr := os.Getenv("CONNSTR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

    log.Println("Searching for products with query:", query)

    words := strings.Split(query, " ")
    var whereClauses []string
    for _, word := range words {
        whereClauses = append(whereClauses, fmt.Sprintf(`
            product_name ILIKE '%%%s%%' 
            OR brand_name ILIKE '%%%s%%' 
            OR category ILIKE '%%%s%%' 
            OR location ILIKE '%%%s%%'
        `, word, word, word, word))
    }
    whereClause := strings.Join(whereClauses, " OR ")
    sqlQuery := fmt.Sprintf(`
        SELECT 
            id, 
            brand_name, 
            product_name, 
            category, 
            location,
            image_url
        FROM 
            products 
        WHERE 
            %s
    `, whereClause)

    rows, err := db.Query(sqlQuery)

    if err != nil {
        log.Fatal(err)
        return err
    }
    defer rows.Close()

    var products []Product = []Product{}

    for rows.Next() {
        var product Product
        if err := rows.Scan(&product.ID, &product.BrandName, &product.ProductName, &product.Category, &product.Location, &product.ImageURL); err != nil {
            return err
        }
        products = append(products, product)
    }

    return c.JSON(http.StatusOK, products)
}