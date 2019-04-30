package dal

import (
	"github.com/yeejlan/maru"
	"github.com/jmoiron/sqlx"
)

//add helper function to sqlx.DB
type DBHelper struct {
	DB *sqlx.DB
}

//new DBHelper
func NewDBHelper(db *sqlx.DB) *DBHelper {
	return &DBHelper{
		DB: db,
	}
}

//select one record
func (this *DBHelper) SelectOne(out interface{}, sql string, params map[string]interface{}) error {
	stmt, err := this.DB.PrepareNamed(sql)
	if err!= nil {
		return maru.WrapError(err, 3)
	}

	err = stmt.Get(out, params)
	if err!= nil && err.Error() == "sql: no rows in result set" { //ignore not found error
		return nil
	}
	if err!= nil {
		return maru.WrapError(err, 3)
	}
	return nil
}

//select
func (this *DBHelper) Select(out interface{}, sql string, params interface{}) error {
	stmt, err := this.DB.PrepareNamed(sql)
	if err!= nil {
		return maru.WrapError(err, 3)
	}

	err = stmt.Select(out, params)
	if err!= nil {
		return maru.WrapError(err, 3)
	}
	return nil
}

//insert
func (this *DBHelper) Insert(sql string, params interface{}) (int, error) {
	stmt, err := this.DB.PrepareNamed(sql)
	if err!= nil {
		return 0, maru.WrapError(err, 3)
	}

	result, err := stmt.Exec(params)
	if err!= nil {
		return 0, maru.WrapError(err, 3)
	}
	insertId, _ := result.LastInsertId()
	return int(insertId), nil
}

//update or delete
func (this *DBHelper) Update(sql string, params interface{}) (int, error) {
	stmt, err := this.DB.PrepareNamed(sql)
	if err!= nil {
		return 0, maru.WrapError(err, 3)
	}

	result, err := stmt.Exec(params)
	if err!= nil {
		return 0, maru.WrapError(err, 3)
	}
	rowsAffected, _ := result.RowsAffected()
	return int(rowsAffected), nil
}