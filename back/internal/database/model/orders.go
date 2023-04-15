package model

/*
DB Table Details
-------------------------------------


Table: orders
[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
[ 1] id_product                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 2] client                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
[ 3] date                                           DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": 48,    "id_product": 36,    "client": 46,    "date": "2143-03-16T15:54:06.682797346-04:00"}



*/

// Orders struct is a row record of the orders table in the store database
type Orders struct {
	//[ 0] id                                             INT4                 null: false  primary: true   isArray: false  auto: false  col: INT4            len: -1      default: []
	ID int `json:"id" db:"id"`
	//[ 1] id_product                                     INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	IdProduct int `json:"id_product" db:"id_product"`
	//[ 2] client                                         INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: []
	Client int `json:"client" db:"client"`
	//[ 3] date                                           DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	Date string `json:"date" db:"date"`
}

var ordersTableInfo = &TableInfo{
	Name: "orders",
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
			Name:               "id_product",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "IDProduct",
			GoFieldType:        "int",
			JSONFieldName:      "id_product",
			ProtobufFieldName:  "id_product",
			ProtobufType:       "int",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "client",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Client",
			GoFieldType:        "int",
			JSONFieldName:      "client",
			ProtobufFieldName:  "client",
			ProtobufType:       "int",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "DATE",
			DatabaseTypePretty: "DATE",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATE",
			ColumnLength:       -1,
			GoFieldName:        "Date",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *Orders) TableName() string {
	return "orders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Orders) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Orders) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Orders) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *Orders) TableInfo() *TableInfo {
	return ordersTableInfo
}
