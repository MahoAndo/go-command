package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"goCommand/conf"
	"goCommand/models/dao"
	"goCommand/models/dao/book"
	"goCommand/models/definitions"
)

type BookData struct {
	Config *conf.Config
}

func NewBookData(c *conf.Config) (*BookData, error) {
	a := &BookData{
		Config: c,
	}
	return a, nil
}

func (d BookData) Insert(b *book.Book) {
	var t book.Book
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//insert data
	err = t.InsertBookRecord(tx, b)
	if err != nil {
		color.Red(definitions.InsertErr, definitions.Book, err)
		return
	}
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, definitions.Book, err)
		return
	}

	if err == nil {
		color.Green(definitions.InsertSuccess, definitions.Book)
	}
}

func (d BookData) Update(b *book.Book) {
	var t book.Book
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//update data
	err = t.UpdateBookRecord(tx, b)
	if err != nil {
		color.Red(definitions.UpdateErr, definitions.Book, err)
		return
	}
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, definitions.Book, err)
		return
	}

	if err == nil {
		color.Green(definitions.UpdateSuccess, definitions.Book)
	}
}

func (d BookData) Delete(b *book.Book) {
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//delete data
	err = book.DeleteBookRecord(tx, b)
	if err != nil {
		color.Red(definitions.DeleteErr, definitions.Book, err)
		return
	}
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, definitions.Book, err)
		return
	}

	if err == nil {
		color.Green(definitions.DeleteSuccess, definitions.Book)
	}
}

func (d BookData) DeleteAll(b *book.Book) {
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//delete
	_, err = tx.DeleteFrom(definitions.Book).
		Where("account_id = ?", b.AccountID).
		Exec()
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, "delete all in book", err)
		return
	}
	color.Green("Successfully delete all data in DB")
}

func (d BookData) Build() *cobra.Command {
	p := &book.Book{}
	var mode string
	var title, author, note string
	var accountID, status int
	
	cc := &cobra.Command{
		Use:   "bookCmd",
		Short: "Setting data for the book record",
		Example: `1. INSERT:
		bookCmd book --mode insert --table book --accountid 10 --title Harry --author JKR --status 1 --note so much fun!!
			2. UPDATE:
			bookCmd book --mode update --table book --accountid 10 --title HarryPotter --author JKR --status 2 --note so much fun!!
			3. DELETE:
			bookCmd book --mode delete --table book --accountid 10 --title HarryPotter
			4. DELETE ALL:
			bookCmd book --mode deleteall --table book`,
		Run: func(cmd *cobra.Command, args []string) {
			switch mode {
			case definitions.Insert:
				color.Blue("insert mode")
				p.AccountID = accountID
				p.Title = title
				p.Author = author
				p.Status = status
				p.Note = note
				d.Insert(p)
			case definitions.Update:
				color.Blue("update mode")
				p.AccountID = accountID
				p.Title = title
				p.Author = author
				p.Status = status
				p.Note = note
				d.update(p)
			case definitions.Delete:
				color.Blue("delete mode")
				p.AccountID = accountID
				p.Title = title
				d.Delete(p)
			case definitions.DeleteAll:
				color.Blue("delete all mode")
				color.Yellow("table: all in book ")
				p.AccountID = accountID
				d.DeleteAll()
			default:
				color.Blue("default")
			}
		},
	}
	return cc

}

