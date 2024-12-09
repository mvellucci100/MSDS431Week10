package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
    _ "modernc.org/sqlite" // Import the sqlite driver to register it
)

func main() {
	// Open SQLite database
	db, err := sql.Open("sqlite", "./movie.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create tables
	createTables(db)

	// Populate tables with data from CSV files
	populateMoviesTable(db)
	populateGenresTable(db)
	populateMoviesGenresTable(db)

	// Query example: Retrieve movies by genre
	queryMoviesByGenre(db)
}

func createTables(db *sql.DB) {
	// Create Movies table
	createMoviesTable := `
	CREATE TABLE IF NOT EXISTS movies (
		movie_id INTEGER PRIMARY KEY,
		title TEXT,
		release_year INTEGER,
		rating FLOAT,
		votes INTEGER
	);
	`
	_, err := db.Exec(createMoviesTable)
	if err != nil {
		panic(err)
	}

	// Create Genres table
	createGenresTable := `
	CREATE TABLE IF NOT EXISTS genres (
		genre_id INTEGER PRIMARY KEY,
		genre TEXT
	);
	`
	_, err = db.Exec(createGenresTable)
	if err != nil {
		panic(err)
	}

	// Create MoviesGenres table
	createMoviesGenresTable := `
	CREATE TABLE IF NOT EXISTS movies_genres (
		movie_id INTEGER,
		genre_id INTEGER,
		FOREIGN KEY (movie_id) REFERENCES movies(movie_id),
		FOREIGN KEY (genre_id) REFERENCES genres(genre_id)
	);
	`
	_, err = db.Exec(createMoviesGenresTable)
	if err != nil {
		panic(err)
	}
}

func populateMoviesTable(db *sql.DB) {
	// Open the CSV file
	file, err := os.Open("IMDB-movies.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, ",")
		if len(columns) < 5 {
			continue
		}

		// Extract movie data
		title := columns[0]
		releaseYear := columns[1]
		rating := columns[2]
		votes := columns[3]

		// Insert into movies table
		_, err := db.Exec("INSERT INTO movies (title, release_year, rating, votes) VALUES (?, ?, ?, ?)", title, releaseYear, rating, votes)
		if err != nil {
			fmt.Println("Error inserting movie data:", err)
		}
	}
}

func populateGenresTable(db *sql.DB) {
	// Open the genres CSV file
	file, err := os.Open("IMDB-movies_genres.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Split(line, ",")
		if len(columns) < 1 {
			continue
		}

		// Extract genre
		genre := columns[0]

		// Insert into genres table
		_, err := db.Exec("INSERT INTO genres (genre) VALUES (?)", genre)
		if err != nil {
			fmt.Println("Error inserting genre data:", err)
		}
	}
}

func populateMoviesGenresTable(db *sql.DB) {
	// Insert data into the movies_genres table
	// Here, you'd relate each movie to its genres
	// You would need additional logic to match movies to genres based on your data structure
}

func queryMoviesByGenre(db *sql.DB) {
	// Query to join movies and genres and get movies with their genres
	rows, err := db.Query(`
		SELECT movies.title, genres.genre
		FROM movies
		INNER JOIN movies_genres ON movies.movie_id = movies_genres.movie_id
		INNER JOIN genres ON movies_genres.genre_id = genres.genre_id
		WHERE genres.genre = 'Action'
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var genre string
		if err := rows.Scan(&title, &genre); err != nil {
			panic(err)
		}
		fmt.Println(title, ":", genre)
	}
}
