package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)



func main() {
	connStr := "postgresql://dean:pass4dean@database-1.cwj7xe0iqba4.us-east-1.rds.amazonaws.com/brandy-db"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
        CREATE TABLE IF NOT EXISTS products (
            id SERIAL PRIMARY KEY,
            brand_name VARCHAR(255) NOT NULL,
            product_name VARCHAR(255) NOT NULL,
            category VARCHAR(255) NOT NULL,
            location VARCHAR(255) NOT NULL
        );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	brandProductPairs := generateBrandProductPairs(20)
	for _, pair := range brandProductPairs {
		insertSQL := "INSERT INTO products (brand_name, product_name, category, location) VALUES ($1, $2, $3, $4)"
		_, err := db.Exec(insertSQL, pair.BrandName, pair.ProductName, pair.Category, randomLocation())
		if err != nil {
			log.Println("Error inserting record:", err)
		}
	}

	fmt.Println("Database seeded successfully!")
}


func randomLocation() string {
	locations := []string{"Los Angeles", "New York", "Seattle", "Dallas", "Miami"}
	return locations[randomInt(0, len(locations))]
}

func randomInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

var coffeeProductPairs = []struct {
	BrandName   string
	ProductName string
	Category    string 
}{
	{"Star Coffee", "Coffee Beans", "Coffee"},
	{"BeanBuzz", "Coffee Grinder", "Coffee Accessories"},
	{"CaffeineCo", "Coffee Maker", "Coffee Appliances"},
	{"JavaJoy", "Coffee Mug", "Drinkware"},
	{"FilterFresh", "Coffee Filter", "Coffee Accessories"},
	{"RoastMaster", "Coffee Scoop", "Coffee Accessories"},
	{"BrewBliss", "Coffee Thermos", "Drinkware"},
	{"PressPerfect", "Coffee Press", "Barista Tools"},
	{"SyrupSipper", "Coffee Syrup", "Coffee Accessories"},
	{"RoastCraft", "Coffee Roaster", "Barista Tools"},
	{"ChillBrew", "Iced Coffee Maker", "Barista Tools"},
	{"FrothFusion", "Coffee Frother", "Coffee Accessories"},
	{"SipSense", "Coffee Tumbler", "Drinkware"},
	{"PerkPal", "Coffee Pot", "Coffee Appliances"},
	{"AromaBrew", "Coffee Percolator", "Coffee Appliances"},
	{"KettleKraft", "Coffee Kettle", "Coffee Appliances"},
	{"PodPioneer", "Coffee Pods", "Coffee Accessories"},
	{"CreamyDelight", "Coffee Creamer", "Coffee Accessories"},
	{"StirMaster", "Coffee Stirrer", "Coffee Accessories"},
	{"BeanVault", "Coffee Storage Container", "Coffee Accessories"},
}

func generateBrandProductPairs(count int) []struct {
	BrandName   string
	ProductName string
	Category    string
} {
	var pairs []struct {
		BrandName   string
		ProductName string
		Category    string
	}

	for i := 0; i < count; i++ {
		randomPair := coffeeProductPairs[i]

		pair := struct {
			BrandName   string
			ProductName string
			Category    string
		}{
			BrandName:   randomPair.BrandName,
			ProductName: randomPair.ProductName,
			Category:    randomPair.Category,
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

