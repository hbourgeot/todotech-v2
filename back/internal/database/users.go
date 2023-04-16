package database

import (
	"context"

	"todotech.henrry.online/internal/database/model"
)

/*


DB Table Details
-------------------------------------
Table: users
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] created                                        TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 2] email                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] hashed_password                                TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []



PrimaryKeyNamesList    : [id]
PrimaryKeysJoined      : id
NonPrimaryKeyNamesList : [created email hashed_password]
NonPrimaryKeysJoined   : created,email,hashed_password
delSql                 : DELETE FROM `users` where id = ?
updateSql              : UPDATE `users` set created = ?, email = ?, hashed_password = ? WHERE id = ?
insertSql              : INSERT INTO `users` ( id,  created,  email,  hashed_password) values ( ?, ?, ?, ? )
selectOneSql           : SELECT * FROM `users` WHERE id = ?
selectMultiSql         : SELECT * FROM `users`


*/

// GetUsers is a function to get a single record from the users table in the store database
// error - ErrNotFound, db Find error
func GetUsers(ctx context.Context, email string) (record *model.Users, err error) {
	sql := "SELECT * FROM users WHERE email = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	record = &model.Users{}
	err = DB.GetContext(ctx, record, sql, email)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// AddUsers is a function to add a single record to users table in the store database
// error - ErrInsertFailed, db save call failed
func AddUsers(ctx context.Context, record *model.Users) (result *model.Users, RowsAffected int64, err error) {
	sql := "INSERT INTO users ( id,  created,  email,  hashed_password) values ( ?, ?, ?, ? )"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult := DB.QueryRowContext(ctx, sql, record.ID, record.Created, record.Email, record.HashedPassword)
	err = dbResult.Scan(record.ID, record.Created, record.Email, record.HashedPassword)

	return record, 1, err
}

// UpdateUsers is a function to update a single record from users table in the store database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateUsers(ctx context.Context, argID int, updated *model.Users) (result *model.Users, RowsAffected int64, err error) {
	sql := "UPDATE `users` set created = ?, email = ?, hashed_password = ? WHERE id = ?"
	sql = DB.Rebind(sql)

	if Logger != nil {
		Logger(ctx, sql)
	}

	dbResult, err := DB.ExecContext(ctx, sql, updated.Created, updated.Email, updated.HashedPassword, argID)
	if err != nil {
		return nil, 0, err
	}

	rows, err := dbResult.RowsAffected()
	updated.ID = argID

	return updated, rows, err
}
