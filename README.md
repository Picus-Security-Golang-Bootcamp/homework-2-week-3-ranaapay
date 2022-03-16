## Homework | Week 3

The book list is taken from bookData.json and assigned to the bookList of the Library struct. After the operations are done, the final version of bookList is written to the file before the program closes.

### list command
It prints books that not deleted in library.
```
go run main.go -command list
```

### search command
Print all books where the value in the entered name parameter is in the name of the book.
```
go run main.go -command search -name <bookName>
go run main.go -command search -name Lord of the Ring: The Return of the King
go run main.go -command search -name Lor
```
* Upper and lower case does not matter.

### get command
It prints book by id, even if book has been deleted.
```
go run main.go -command get -id <bookID>
go run main.go -command get -id 5
```

### delete command
It changes the isDeleted field in book.
```
go run main.go -command delete -id <bookID>
go run main.go -command delete -id 5
```

### buy command
It buys the book according to the given id and quantity value. It updates the stock number field in the book. It returns the final situation of the book and how much bought the book.
```
go run main.go -command buy -id <bookID> -quantity <quantity>
go run main.go -command buy -id 5 -quantity 2
```

###default
Returns how the flag should be used.
