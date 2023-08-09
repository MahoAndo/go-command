package book

import (
	"github.com/gocraft/dbr"
	"github.com/thoas/go-funk"
	"time"
	"goCommand/models/definitions"
)

type Book struct {
	AccountID      int       `db:"account_id"`
	Title          string    `db:"title"`
	Author         string    `db:"author"`
	Status         string    `db:"status"`
	Note           string    `db:"note"`
	CreateDatetime time.Time `db:"create_datetime"`
	UpdateDateTime time.Time `db:"update_datetime"`
}

func (p Book) selectAllBookRecord(tx *dbr.Tx, book Book) (err error) {
	var result []Book
	_, err = tx.Select("*").
		From(definitions.Book).
		Where("account_id = ? ", book.AccountID).
		Load(&result)
	return
}

func selectMaxBookRecordNo(tx *dbr.Tx, accountID int) (err error, maxNo int) {
	var result []int
	_, err = tx.Select("book_no").
		From(definitions.Book).
		Where("account_id = ? ", accountID).
		Load(&result)

	//getting max no from BookNO array
	maxNo = funk.MaxInt(result)
	return
}

func (p Book) InsertBookRecord(tx *dbr.Tx, book *Book) (err error) {
	//setting insert data
	p.AccountID = book.AccountID
	p.Title = book.Title
	p.Author = book.Author
	p.Status = book.Status
	p.Note = book.Note
	p.CreateDatetime = time.Now()
	p.UpdateDateTime = time.Now()

	//get max no of book record
	e, maxNo := selectMaxBookRecordNo(tx, book.AccountID)
	if e != nil {
		return
	} else {
		p.BookNO = maxNo + 1
	}

	//insert
	_, err = tx.InsertInto(definitions.Book).
		Columns(
			"account_id",
			"title",
			"author",
			"status",
			"note",
			"create_datetime",
			"update_datetime",
		).
		Values(
			p.AccountID,
			p.Title,
			p.Author,
			p.Status,
			p.Note,
			p.CreateDatetime,
			p.UpdateDateTime,
		).Exec()
	return
}

func (p Book) UpdateBookRecord(tx *dbr.Tx, book *Book) (err error) {
	//setting insert data
	p.Title = book.Title
	p.Author = book.Author
	p.Status = book.Status
	p.Note = book.Note
	p.UpdateDateTime = time.Now()

	//update
	_, err = tx.Update(definitions.Book).
		Set("title", p.Title).
		Set("author", p.Author).
		Set("status", p.Status).
		Set("note", p.CreateDatetime).
		Set("update_datetime", p.UpdateDateTime).
		Where("account_id = ? and title = ?", book.AccountID, book.Title).
		Exec()
	return
}

func DeleteBookRecord(tx *dbr.Tx, book *Book) (err error) {
	//delete
	_, err = tx.DeleteFrom(definitions.Book).
	Where("account_id = ? and title = ?", book.AccountID, book.Title).
		Exec()
	return
}
