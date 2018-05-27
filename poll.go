package main

import (
	"database/sql"
	"gopoll/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
        CREATE TABLE IF NOT EXISTS polls(
                id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                name VARCHAR NOT NULL,
                topic VARCHAR NOT NULL,
                src VARCHAR NOT NULL,
                upvotes INTEGER NOT NULL,
                downvotes INTEGER NOT NULL,
                UNIQUE(name)
        );

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Angular','Awesome Angular', 'https://cdn.colorlib.com/wp/wp-content/uploads/sites/2/angular-logo.png', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Vue', 'Voguish Vue','https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('React','Remarkable React','https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Ember','Excellent Ember','https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png', 1, 0);

        INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Knockout','Knightly Knockout','https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png', 1, 0);
   `
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the database
	db := initDB("storage.db")
	migrate(db)

	// Define the HTTP routes
	e.File("/", "public/index.html")
	e.GET("/polls", handlers.GetPolls(db))
	e.PUT("/update/:index", handlers.UpdatePoll(db))

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
