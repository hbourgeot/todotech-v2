package database

import (
	"context"
	"fmt"
	"strings"
	"todotech.henrry.online/internal/database/model"
)

/*


DB Table Details
-------------------------------------
Table: customers
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] name                                           VARCHAR(200)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 200     default: []
[ 2] document_type                                  USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
[ 3] document                                       VARCHAR(15)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
[ 4] total_expense                                  MONEY                null: false  primary: false  isArray: false  auto: false  col: MONEY           len: -1      default: [0]



PrimaryKeyNamesList    : [id]
PrimaryKeysJoined      : id
NonPrimaryKeyNamesList : [name document_type document total_expense]
NonPrimaryKeysJoined   : name,document_type,document,total_expense
delSql                 : DELETE FROM `customers` where id = ?
updateSql              : UPDATE `customers` set name = ?, document_type = ?, document = ?, total_expense = ? WHERE id = ?
insertSql              : INSERT INTO `customers` ( id,  name,  document_type,  document,  total_expense) values ( ?, ?, ?, ?, ? )
selectOneSql           : SELECT * FROM `customers` WHERE id = ?
selectMultiSql         : SELECT * FROM `customers`


*/

// GetAllCustomers is a function to get a slice of record(s) from customers table in the store database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCustomers(ctx context.Context, order string) (results []*model.Customers, totalRows int, err error) {

	if order != "" {
		if strings.ContainsAny(order, "\"") {
			order = ""
		}
	}

	if order == "" {
		order = "id"
	}
	sql := fmt.Sprintf("SELECT * FROM customers ORDER BY %s", order)
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	err = DB.SelectContext(ctx, &results, sql)
	if err != nil {
		fmt.Println("aqui", sql, results)
		return nil, -1, err
	}

	cnt, err := GetRowCount(ctx, "customers")
	if err != nil {
		return results, -2, err
	}

	return results, cnt, err
}

// GetCustomers is a function to get a single record from the customers table in the store database
// error - ErrNotFound, db Find error
func GetCustomers(ctx context.Context, argID int) (record *model.Customers, err error) {
	sql := "SELECT * FROM customers WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	record = &model.Customers{}
	err = DB.GetContext(ctx, record, sql, argID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddCustomers is a function to add a single record to customers table in the store database
// error - ErrInsertFailed, db save call failed
func AddCustomers(ctx context.Context, record *model.Customers) (result *model.Customers, RowsAffected int64, err error) {
	sql := "INSERT INTO customers ( id,  name,  document_type,  document,  total_expense) values ( ?, ?, ?, ?, ? )"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	sql = fmt.Sprintf("%s returning %s", sql, "id")
	dbResult := DB.QueryRowContext(ctx, sql, record.ID, record.Name, record.Document, record.TotalExpense)
	err = dbResult.Scan(record)

	return record, 1, err
}

// UpdateCustomers is a function to update a single record from customers table in the store database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateCustomers(ctx context.Context, argID int, updated *model.Customers) (result *model.Customers, RowsAffected int64, err error) {
	sql := "UPDATE customers set name = ?, document = ?, total_expense = ? WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult, err := DB.ExecContext(ctx, sql, updated.Name, updated.Document, updated.TotalExpense, argID)
	if err != nil {
		return nil, 0, err
	}

	rows, err := dbResult.RowsAffected()
	updated.ID = argID

	return updated, rows, err
}

// DeleteCustomers is a function to delete a single record from customers table in the store database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteCustomers(ctx context.Context, argID int) (rowsAffected int64, err error) {
	sql := "DELETE FROM `customers` where id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	result, err := DB.ExecContext(ctx, sql, argID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
