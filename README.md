# MSDS431Week100

### Introduction
This project involves developing a local movie database application using Go and SQLite. The goal is to create a simple relational database that stores movie data from the IMDb dataset, performs SQL queries, and allows for interactions beyond just querying the database. Additionally, we'll explore the possibilities of adding more features to the database and extending the application to create a robust movie tracking system

### Process
Download and Extract IMDb Data
The first step is to download the necessary data files from the Northwestern University IMDb dataset archive. These files are in CSV format and contain information about movies, genres, ratings, and more.

You can download the archive from the Northwestern Library Archive.

Once you have the archive, extract the six CSV files. The files you will work with are:

IMDB-movies.csv
IMDB-movies_genres.csv
(Additional files, depending on what information you want to include)
Step 2: Define the SQLite Schema
Movies Table
To store movie data, we'll create a table called movies to hold the relevant information for each movie, such as:

movie_id (integer, primary key)
title (text)
release_year (integer)
rating (float)
votes (integer)
Genres Table
To store movie genres, we'll create a genres table with the following columns:

genre_id (integer, primary key)
genre (text)
MoviesGenres Table (Many-to-Many Relationship)
Since a movie can have multiple genres, we need a many-to-many relationship table:

movie_id (integer, foreign key referencing movies)
genre_id (integer, foreign key referencing genres)
Step 3: Create and Populate the SQLite Database with Go
We'll use a pure Go SQLite library to avoid cgo dependencies, such as go-sqlite3, which requires cgo. The alternative is using cznic/sqlite, a cgo-free SQLite implementation.

The Go program will:

Set up an SQLite database.
Define the schema (tables) for the movies and genres.
Load the movie and genre data from the CSV files into the database.

### Purpose
The personal movie database serves several purposes beyond what IMDb offers:
Organizational Tool
* Keep track of your own collection
* Quickly identify where a movie is stored or where to watch it (Netflix, Hulu, Peacock)
* Saves storage

Personalized Ratings
* You can track your movies based on your own ratings instead of public ones

Movie Exploration
* Organize your collection by genres, ratings, and you can create custom filters for searching and sorting movies
* Track which movies you want to watch next based off of certain criteria like genre or rating

### User Interactions Beyond SQL Queries
A Go application could provide a user interface to:

* Track movies in your collection.
* Display personalized movie recommendations based on user ratings.
* Allow users to search for movies by genre, rating, or release year.
* Allow users to add new movies and genres to the database.

