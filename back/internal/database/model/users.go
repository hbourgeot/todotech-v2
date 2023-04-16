package model

import "time"

/*
DB Table Details
-------------------------------------


Table: users
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] created                                        TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
[ 2] email                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] hashed_password                                TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 19,    "created": "2027-07-02T05:48:30.41202247-04:00",    "email": "HbDPDRwmLRSNRxTMomWfspJRa",    "hashed_password": "JkaXQgrgsevHOFrOonYSorwat"}



*/

// Users struct is a row record of the users table in the store database
type Users struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int `json:"id"`
	//[ 1] created                                        TIMESTAMPTZ          null: false  primary: false  isArray: false  auto: false  col: TIMESTAMPTZ     len: -1      default: []
	Created time.Time `json:"created"`
	//[ 2] email                                          TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	Email string `json:"email"`
	//[ 3] hashed_password                                TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
	HashedPassword []byte `json:"hashed_password" db:"hashed_password"`
}

var usersTableInfo = &TableInfo{
	Name: "users",
	Columns: []*ColumnInfo{
		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "created",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TIMESTAMPTZ",
			DatabaseTypePretty: "TIMESTAMPTZ",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMPTZ",
			ColumnLength:       -1,
			GoFieldName:        "Created",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created",
			ProtobufFieldName:  "created",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "email",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "hashed_password",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "HashedPassword",
			GoFieldType:        "string",
			JSONFieldName:      "hashed_password",
			ProtobufFieldName:  "hashed_password",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *Users) TableName() string {
	return "users"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *Users) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *Users) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *Users) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *Users) TableInfo() *TableInfo {
	return usersTableInfo
}
