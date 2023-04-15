package model

/*
DB Table Details
-------------------------------------


Table: customers
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] name                                           VARCHAR(200)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 200     default: []
[ 2] document_type                                  USER_DEFINED         null: false  primary: false  isArray: false  auto: false  col: USER_DEFINED    len: -1      default: []
[ 3] document                                       VARCHAR(15)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
[ 4] total_expense                                  MONEY                null: false  primary: false  isArray: false  auto: false  col: MONEY           len: -1      default: [0]


JSON Sample
-------------------------------------
{    "id": 76,    "name": "JwGCIOvYiXPJkkWnfQOfPXthn",    "document_type": "PdSLgEFdTVyHpgLBcLJEKXWiB",    "document": "UyhwEkJBvAkSgWbpcpAKGNxSE",    "total_expense": 0.8231917813202464}



*/

// Customers struct is a row record of the customers table in the store database
type Customers struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int `json:"id"`
	//[ 1] name                                           VARCHAR(200)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 200     default: []
	Name string `json:"name"`
	//[ 3] document                                       VARCHAR(15)          null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 15      default: []
	Document string `json:"document"`
	//[ 4] total_expense                                  MONEY                null: false  primary: false  isArray: false  auto: false  col: MONEY           len: -1      default: [0]
	TotalExpense string `json:"total_expense" db:"total_expense"`
}

var customersTableInfo = &TableInfo{
	Name: "customers",
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
			DatabaseTypePretty: "VARCHAR(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       200,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "document_type",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "USER_DEFINED",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "USER_DEFINED",
			ColumnLength:       -1,
			GoFieldName:        "DocumentType",
			GoFieldType:        "string",
			JSONFieldName:      "document_type",
			ProtobufFieldName:  "document_type",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "document",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(15)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       15,
			GoFieldName:        "Document",
			GoFieldType:        "string",
			JSONFieldName:      "document",
			ProtobufFieldName:  "document",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "total_expense",
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
			GoFieldName:        "TotalExpense",
			GoFieldType:        "float64",
			JSONFieldName:      "total_expense",
			ProtobufFieldName:  "total_expense",
			ProtobufType:       "float",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Customers) TableName() string {
	return "customers"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Customers) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Customers) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Customers) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Customers) TableInfo() *TableInfo {
	return customersTableInfo
}
