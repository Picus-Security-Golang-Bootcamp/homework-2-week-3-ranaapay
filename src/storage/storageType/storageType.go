package storageType

import "fmt"

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      Author `json:"author"`
	Price       string `json:"price"`
	PublishTime string `json:"publishTime"`
	ISBN        string `json:"isbn"`
	PageNumber  int    `json:"pageNumber"`
	StockCode   string `json:"stockCode"`
	StockNumber int    `json:"stockNumber"`
	IsDeleted   bool   `json:"deleted"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (b *Book) GetBookId() int {
	return b.Id
}
func (b *Book) GetBookName() string {
	return b.Name
}
func (b *Book) GetAuthorName() string {
	return b.Author.Name
}
func (b *Book) GetPublishTime() string {
	return b.PublishTime
}
func (b *Book) GetPrice() string {
	return b.Price
}
func (b *Book) GetPageNumber() int {
	return b.PageNumber
}
func (b *Book) GetISBN() string {
	return b.ISBN
}
func (b *Book) GetStockCode() string {
	return b.StockCode
}
func (b *Book) GetIsDeleted() bool {
	return b.IsDeleted
}
func (b *Book) GetStockNumber() int {
	return b.StockNumber
}
func (b *Book) SetStockNumber(num int) {
	b.StockNumber = num
}
func (b *Book) SetIsDeleted(isDeleted bool) {
	b.IsDeleted = isDeleted
}

func (b *Book) PrintBook() {
	fmt.Printf("Book : %d ---------------\n"+
		"Name : %s\n"+
		"Author Name : %s\n"+
		"Publish Year : %s\n"+
		"Price : %s\n"+
		"Page Number : %d\n"+
		"ISBN : %s\n"+
		"Stock Code : %s\n"+
		"StockNumber : %d\n"+
		"Deleted : %t\n",
		b.GetBookId(), b.GetBookName(), b.GetAuthorName(), b.GetPublishTime(), b.GetPrice(),
		b.GetPageNumber(), b.GetISBN(), b.GetStockCode(), b.GetStockNumber(), b.GetIsDeleted())
}
