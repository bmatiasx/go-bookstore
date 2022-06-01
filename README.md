# Go Bookstore

This app is a web service with the basic APIs to handle
the creation, update, retrieval and delete of books in 
the context of a book store.

## Build

Go to the directory `cmd/main` and there enter the following
command in the terminal:
```bash
go build
```

Once the build is done now the project can be run with the command:
```bash
go run main.go
```

## Example requests
### Create Book

```bash
curl --location --request POST 'localhost:9090/book/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Mastering Go",
    "author": "Mihalis Tsoukalos",
    "publication": "Packt"
}'
```

### Get all Books
```bash
curl --location --request GET 'localhost:9090/book'
```

### Get Book by ID
```bash
curl --location --request GET 'localhost:9090/book/1'
```

### Update Book
```bash
curl --location --request PUT 'localhost:9090/book/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Java Concurrency in Practice",
    "publication": "Pearson Education"
}'
```

### Delete Book
```bash
curl --location --request DELETE 'localhost:9090/book/2'
```