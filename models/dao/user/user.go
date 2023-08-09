package user

import (
	"github.com/gocraft/dbr"
	"time"
	"goCommand/models/definitions"
)

type AccountUser struct {
	AccountID      int       `db:"account_id"`
	AccountName    string    `db:"account_name"`
	MailAddress    string    `db:"mail_address"`
	CreateDatetime time.Time `db:"create_datetime"`
	UpdateDateTime time.Time `db:"update_datetime"`
}

func (p AccountUser) InsertAccountUser(tx *dbr.Tx, userInfo *AccountUser) (err error) {
	//setting insert data
	p.AccountName = userInfo.AccountName
	p.MailAddress = userInfo.MailAddress
	p.CreateDatetime = time.Now()
	p.UpdateDateTime = time.Now()

	//insert
	_, err = tx.InsertInto(definitions.AccountUser).
		Columns(
			"account_name",
			"mail_address",
			"create_datetime",
			"update_datetime",
		).
		Values(
			p.AccountName,
			p.MailAddress,
			p.CreateDatetime,
			p.UpdateDateTime,
		).
		Exec()
	return
}

func (p AccountUser) UpdateAccountUser(tx *dbr.Tx, userInfo *AccountUser) (err error) {
	//setting insert data
	p.AccountName = userInfo.AccountName
	p.MailAddress = userInfo.MailAddress
	p.UpdateDateTime = time.Now()

	//update
	_, err = tx.Update(definitions.AccountUser).
		Set("account_name", p.AccountName).
		Set("mail_address", p.MailAddress).
		Set("update_datetime", p.UpdateDateTime).
		Where("account_id = ?", p.AccountID).
		Exec()
	return
}

func DeleteAccountUser(tx *dbr.Tx, accountID int) (err error) {
	//delete
	_, err = tx.DeleteFrom(definitions.AccountUser).
		Where("account_id = ?", accountID).
		Exec()
	return
}
