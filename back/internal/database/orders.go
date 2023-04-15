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
Table: orders
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] id_product                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] client                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] date                                           DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []



PrimaryKeyNamesList    : [id]
PrimaryKeysJoined      : id
NonPrimaryKeyNamesList : [id_product client date]
NonPrimaryKeysJoined   : id_product,client,date
delSql                 : DELETE FROM `orders` where id = ?
updateSql              : UPDATE `orders` set id_product = ?, client = ?, date = ? WHERE id = ?
insertSql              : INSERT INTO `orders` ( id,  id_product,  client,  date) values ( ?, ?, ?, ? )
selectOneSql           : SELECT * FROM `orders` WHERE id = ?
selectMultiSql         : SELECT * FROM `orders`


*/

// GetAllOrders is a function to get a slice of record(s) from orders table in the store database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllOrders(ctx context.Context, order string) (results []*model.Orders, totalRows int, err error) {
	if order != "" {
		if strings.ContainsAny(order, "'\"") {
			order = ""
		}
	}

	if order == "" {
		order = "id"
	}
	sql := fmt.Sprintf("SELECT id, id_product, client, date FROM orders ORDER BY %s", order)
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	err = DB.SelectContext(ctx, &results, sql)
	if err != nil {
		return nil, -1, err
	}

	cnt, err := GetRowCount(ctx, "orders")
	if err != nil {
		return results, -2, err
	}

	return results, cnt, err
}

// GetOrders is a function to get a single record from the orders table in the store database
// error - ErrNotFound, db Find error
func GetOrders(ctx context.Context, argID int) (record *model.Orders, err error) {
	sql := "SELECT * FROM orders WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	record = &model.Orders{}
	err = DB.GetContext(ctx, record, sql, argID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddOrders is a function to add a single record to orders table in the store database
// error - ErrInsertFailed, db save call failed
func AddOrders(ctx context.Context, record *model.Orders) (result *model.Orders, RowsAffected int64, err error) {
	sql := "INSERT INTO orders ( id,  id_product,  client,  date) values ( ?, ?, ?, ? )"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult := DB.QueryRowContext(ctx, sql, record.ID, record.IdProduct, record.Client, record.Date)
	err = dbResult.Scan(record.ID, record.IdProduct, record.Client, record.Date)

	return record, 1, err
}

// UpdateOrders is a function to update a single record from orders table in the store database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateOrders(ctx context.Context, argID int, updated *model.Orders) (result *model.Orders, RowsAffected int64, err error) {
	sql := "UPDATE orders set id_product = ?, client = ?, date = ? WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult, err := DB.ExecContext(ctx, sql, updated.IdProduct, updated.Client, updated.Date, argID)
	if err != nil {
		return nil, 0, err
	}

	rows, err := dbResult.RowsAffected()
	updated.ID = argID

	return updated, rows, err
}

// DeleteOrders is a function to delete a single record from orders table in the store database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteOrders(ctx context.Context, argID int) (rowsAffected int64, err error) {
	sql := "DELETE FROM orders where id = ?"
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
