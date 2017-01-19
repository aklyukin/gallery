package main

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type userTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *userTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("users").
func (v *userTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *userTableType) Columns() []string {
	return []string{"id", "email", "password", "reg_date"}
}

// NewStruct makes a new struct for that view or table.
func (v *userTableType) NewStruct() reform.Struct {
	return new(User)
}

// NewRecord makes a new record for that table.
func (v *userTableType) NewRecord() reform.Record {
	return new(User)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *userTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// UserTable represents users view or table in SQL database.
var UserTable = &userTableType{
	s: parse.StructInfo{Type: "User", SQLSchema: "", SQLName: "users", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "Email", PKType: "", Column: "email"}, {Name: "Password", PKType: "", Column: "password"}, {Name: "RegDate", PKType: "", Column: "reg_date"}}, PKFieldIndex: 0},
	z: new(User).Values(),
}

// String returns a string representation of this struct or record.
func (s User) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Email: " + reform.Inspect(s.Email, true)
	res[2] = "Password: " + reform.Inspect(s.Password, true)
	res[3] = "RegDate: " + reform.Inspect(s.RegDate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *User) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Email,
		s.Password,
		s.RegDate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *User) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Email,
		&s.Password,
		&s.RegDate,
	}
}

// View returns View object for that struct.
func (s *User) View() reform.View {
	return UserTable
}

// Table returns Table object for that record.
func (s *User) Table() reform.Table {
	return UserTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *User) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *User) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *User) HasPK() bool {
	return s.ID != UserTable.z[UserTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *User) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = UserTable
	_ reform.Struct = new(User)
	_ reform.Table  = UserTable
	_ reform.Record = new(User)
	_ fmt.Stringer  = new(User)
)

type photoTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *photoTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("photos").
func (v *photoTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *photoTableType) Columns() []string {
	return []string{"id", "photo_name", "filename", "create_date"}
}

// NewStruct makes a new struct for that view or table.
func (v *photoTableType) NewStruct() reform.Struct {
	return new(Photo)
}

// NewRecord makes a new record for that table.
func (v *photoTableType) NewRecord() reform.Record {
	return new(Photo)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *photoTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// PhotoTable represents photos view or table in SQL database.
var PhotoTable = &photoTableType{
	s: parse.StructInfo{Type: "Photo", SQLSchema: "", SQLName: "photos", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "PhotoName", PKType: "", Column: "photo_name"}, {Name: "FileName", PKType: "", Column: "filename"}, {Name: "CreateDate", PKType: "", Column: "create_date"}}, PKFieldIndex: 0},
	z: new(Photo).Values(),
}

// String returns a string representation of this struct or record.
func (s Photo) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "PhotoName: " + reform.Inspect(s.PhotoName, true)
	res[2] = "FileName: " + reform.Inspect(s.FileName, true)
	res[3] = "CreateDate: " + reform.Inspect(s.CreateDate, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Photo) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.PhotoName,
		s.FileName,
		s.CreateDate,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Photo) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.PhotoName,
		&s.FileName,
		&s.CreateDate,
	}
}

// View returns View object for that struct.
func (s *Photo) View() reform.View {
	return PhotoTable
}

// Table returns Table object for that record.
func (s *Photo) Table() reform.Table {
	return PhotoTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Photo) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Photo) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Photo) HasPK() bool {
	return s.ID != PhotoTable.z[PhotoTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Photo) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = PhotoTable
	_ reform.Struct = new(Photo)
	_ reform.Table  = PhotoTable
	_ reform.Record = new(Photo)
	_ fmt.Stringer  = new(Photo)
)

func init() {
	parse.AssertUpToDate(&UserTable.s, new(User))
	parse.AssertUpToDate(&PhotoTable.s, new(Photo))
}
