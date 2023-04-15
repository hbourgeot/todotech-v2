package model

import "github.com/lib/pq"

/*
DB Table Details
-------------------------------------


Table: products
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] name                                           VARCHAR(300)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 300     default: []
[ 2] description                                    TEXT                 null: false  primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 3] images                                         _VARCHAR             null: false  primary: false  isArray: false  auto: false  col: _VARCHAR        len: -1      default: []
[ 4] price                                          MONEY                null: false  primary: false  isArray: false  auto: false  col: MONEY           len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 82,    "name": "ejlADyvggyKJxECYqjLFUHXMb",    "description": "nipAEcxofWOFGWIBPsMMUSylB",    "images": 0.7609915134081692}



*/

// Products struct is a row record of the products table in the store database
type Products struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
	Price       string         `json:"price"`
}

var productsTableInfo = &TableInfo{
	Name: "products",
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
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       300,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "description",
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
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              4,
			Name:               "price",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "MONEY",
			DatabaseTypePretty: "MONEY",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "MONEY",
			ColumnLength:       -1,
			GoFieldName:        "Price",
			GoFieldType:        "float64",
			JSONFieldName:      "price",
			ProtobufFieldName:  "price",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *Products) TableName() string {
	return "products"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Products) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Products) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Products) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *Products) TableInfo() *TableInfo {
	return productsTableInfo
}
