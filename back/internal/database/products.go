package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"todotech.henrry.online/internal/database/model"
)

/*


DB Table Details
-------------------------------------
Table: products
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] name                                           VARCHAR(300)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 300     default: []
[ 2] description                                    TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] images                                         _VARCHAR             null: false  primary: false  isArray: false  auto: false  col: _VARCHAR        len: -1      default: []
[ 4] price                                          MONEY                null: false  primary: false  isArray: false  auto: false  col: MONEY           len: -1      default: []



PrimaryKeyNamesList    : [id]
PrimaryKeysJoined      : id
NonPrimaryKeyNamesList : [name description images price]
NonPrimaryKeysJoined   : name,description,images,price
delSql                 : DELETE FROM `products` where id = ?
updateSql              : UPDATE `products` set name = ?, description = ?, images = ?, price = ? WHERE id = ?
insertSql              : INSERT INTO `products` ( id,  name,  description,  images,  price) values ( ?, ?, ?, ?, ? )
selectOneSql           : SELECT * FROM `products` WHERE id = ?
selectMultiSql         : SELECT * FROM `products`


*/

// GetAllProducts is a function to get a slice of record(s) from products table in the store database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProducts(ctx context.Context, order string) (results []*model.Products, totalRows int, err error) {
	if order != "" {
		if strings.ContainsAny(order, "'\"") {
			order = ""
		}
	}

	if order == "" {
		order = "id"
	}

	sql := fmt.Sprintf("SELECT * FROM products ORDER BY %s", order)
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	err = DB.SelectContext(ctx, &results, sql)
	if err != nil {
		return nil, -1, err
	}

	cnt, err := GetRowCount(ctx, "products")
	if err != nil {
		return results, -2, err
	}

	return results, cnt, err
}

// GetProducts is a function to get a single record from the products table in the store database
// error - ErrNotFound, db Find error
func GetProducts(ctx context.Context, argID int) (record *model.Products, err error) {
	sql := "SELECT * FROM products WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	record = &model.Products{}
	err = DB.GetContext(ctx, record, sql, argID)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddProducts is a function to add a single record to products table in the store database
// error - ErrInsertFailed, db save call failed
func AddProducts(ctx context.Context, record *model.Products) (result *model.Products, RowsAffected int64, err error) {
	sql := "INSERT INTO products ( id,  name,  description,  price, images) values ( ?, ?, ?, ?, ? )"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult := DB.QueryRowContext(ctx, sql, record.ID, record.Name, record.Description, record.Price, pq.Array(record.Images))
	err = dbResult.Scan(record.ID, record.Name, record.Description, record.Price, record.Images)
	fmt.Println(record, dbResult.Err())
	return record, 1, err
}

// UpdateProducts is a function to update a single record from products table in the store database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProducts(ctx context.Context, argID int, updated *model.Products) (result *model.Products, RowsAffected int64, err error) {
	sql := "UPDATE products set name = ?, description = ?, images = ?, price = ? WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult, err := DB.ExecContext(ctx, sql, updated.Name, updated.Description, updated.Price, argID)
	if err != nil {
		return nil, 0, err
	}

	rows, err := dbResult.RowsAffected()
	updated.ID = argID

	return updated, rows, err
}

// DeleteProducts is a function to delete a single record from products table in the store database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProducts(ctx context.Context, argID int) (rowsAffected int64, err error) {
	sql := "DELETE FROM products where id = ?"
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
