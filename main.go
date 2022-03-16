package main

import (
	"PicusHomework2/src/service"
	"PicusHomework2/src/storage"
	"flag"
	"fmt"
	"os"
)

var (
	command     = flag.String("command", "", "chooses which command operation it wants to perform")
	requestId   = flag.Int("id", 0, "id of the book")
	requestName = flag.String("name", "", "name of the book")
	quantity    = flag.Int("quantity", 0, "specifies the number of books to be purchased")
)

var usage = `Usage: ./PicusHomework2 [options...]
Options:
	-command   Chooses which command operation it wants to perform.
				list, search, get, delete, and buy; commands that valid.
	-id        Entered command requests a value to determine the books.
	-name      Entered command requests a name to determine the books.
	-quantity  After buy command specifies the number of books to be purchased.
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()

	library := storage.NewLibrary()
	service := service.NewService(library)

	switch *command {
	case "list":
		service.WriteBookList()
	case "search":
		service.SearchBooks(*requestName)
	case "get":
		service.WriteBookById(*requestId)
	case "delete":
		service.DeleteBookById(*requestId)
	case "buy":
		service.BuyBookByItsId(*requestId, *quantity)
	default:
		flag.Usage()
	}

	defer service.SaveChanges()
}
