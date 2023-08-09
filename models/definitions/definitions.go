package definitions

const (
	Insert    = "insert"
	Delete    = "delete"
	Update    = "update"
	InsertAll = "insertall"
	DeleteAll = "deleteall"
)

const (
	BookDB     = "bookdb"
)

// table
const (
	AccountUser = "account_user"
	Book        = "book"
)

// status
const (
	Start   = 1
	Reading = 2
	Finish  = 3
)


// common message
const (
	//insert
	InsertSuccess       = "Successfully insert data in %v"
	InsertErr           = "Insert %v error: %v"

	// update
	UpdateErr     = "Update %v error: %v"
	UpdateSuccess = "Successfully update data in %v"

	// update
	DeleteErr     = "Update %v error: %v"
	DeleteSuccess = "Successfully delete data in %v"

	//deleteAll
	DeleteAllSuccess = "Successfully delete all data in %v"

	// trannsaction
	TransactionErrorFormat = "Transaction %v commit error : %v"
)

