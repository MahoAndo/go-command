package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"goCommand/conf"
	"goCommand/models/dao"
	"goCommand/models/dao/user"
	"goCommand/models/definitions"
)

type UserData struct {
	Config *conf.Config
}

func NewUserData(c *conf.Config) (*UserData, error) {
	a := &UserData{
		Config: c,
	}
	return a, nil
}

func (d UserData) Insert(u *user.AccountUser) {
	var t user.AccountUser
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//insert data
	err = t.InsertAccountUser(tx, u)
	if err != nil {
		color.Red(definitions.InsertErr, definitions.AccountUser, err)
		return
	}
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, definitions.AccountUser, err)
		return
	}

	if err == nil {
		color.Green(definitions.InsertSuccess, definitions.AccountUser)
	}
}

func (d UserData) Update(u *user.AccountUser) {
	var t user.AccountUser
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//update data
	err = t.UpdateAccountUser(tx, u)
	if err != nil {
		color.Red(definitions.UpdateErr, definitions.AccountUser, err)
		return
	}
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, definitions.AccountUser, err)
		return
	}

	if err == nil {
		color.Green(definitions.UpdateSuccess, definitions.AccountUser)
	}
}

func (d UserData) Delete(u *user.AccountUser) {
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//delete data
	err = user.DeleteAccountUser(tx, u.AccountID)
	if err != nil {
		color.Red(definitions.DeleteErr, definitions.AccountUser, err)
		return
	}
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, definitions.AccountUser, err)
		return
	}

	if err == nil {
		color.Green(definitions.DeleteSuccess, definitions.AccountUser)
	}
}

func (d UserData) DeleteAll() {
	err := error(nil)

	//make a transaction session
	transaction := dao.GetDBSession(d.Config)

	//start transaction
	tx, _ := transaction.Begin()
	defer tx.Rollback()

	//delete
	_, err = tx.DeleteFrom(definitions.AccountUser).
		Exec()
	if err = tx.Commit(); err != nil {
		color.Red(definitions.TransactionErrorFormat, "delete all in account_user", err)
		return
	}
	color.Green("Successfully delete all data in DB")
}

func (d UserData) Build() *cobra.Command {
	p := &user.AccountUser{}
	var mode string
	var accountname, mailaddress string
	var accountid int
	
	cc := &cobra.Command{
		Use:   "bookCmd",
		Short: "Setting data for the book record",
		Example: `1. INSERT:
		bookCmd book --mode insert --table account_user --accountname Kayle --mailaddress kayle@gmail.com
			2. UPDATE:
			bookCmd book --mode update --table account_user --accountid 10 --accountname Kayle --mailaddress kayle1013@gmail.com
			3. DELETE:
			bookCmd book --mode delete --table account_user --accountid 10
			4. DELETE ALL:
			bookCmd book --mode deleteall --table account_user`,
		Run: func(cmd *cobra.Command, args []string) {
			switch mode {
			case definitions.Insert:
				color.Blue("insert mode")
				color.Yellow("table: account_user")
				p.AccountName = accountName
				p.MailAddress = mailAddress
				d.Insert(p)
			case definitions.Update:
				color.Blue("update mode")
				color.Yellow("table: account_user")
				p.AccountID = accountID
				p.AccountName = accountName
				p.MailAddress = mailAddress
				d.update(p)
			case definitions.Delete:
				color.Blue("delete mode")
				color.Yellow("table: account_user")
				p.AccountID = accountID
				p.Title = title
				d.Delete(p)
			case definitions.DeleteAll:
				color.Blue("delete all mode")
				color.Yellow("table: all in account_user")
				p.AccountID = accountID
				d.DeleteAll()
			default:
				color.Blue("default")
			}
		},
	}
	return cc

}
