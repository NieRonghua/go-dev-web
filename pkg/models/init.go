package models

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbMysqlMap map[string]*MysqlDBWrap

func init() {
	dbMysqlMap = make(map[string]*MysqlDBWrap)
}

type MysqlDBWrap struct {
	Host     string
	Username string
	Password string
	Schema   string
	Charset  string
	Loc      string

	MaxConnections  int
	MaxConnLifetime int64

	ConnectionTimeout int

	d     *sql.DB
	valid bool
}

func (db *MysqlDBWrap) genMysqlConnString() string {
	//check if contains 3306
	hostArr := strings.Split(db.Host, ":")
	if len(hostArr) == 1 {
		db.Host = fmt.Sprintf("%s:3306", db.Host)
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&loc=%s&parseTime=true&timeout=%ds&time_zone=%s", db.Username, db.Password, db.Host, db.Schema, db.Charset, db.Loc, db.ConnectionTimeout, "%27%2b00:00%27")
}

func (db *MysqlDBWrap) init() error {
	var err error
	if db.Charset == "" {
		db.Charset = "utf8mb4"
	}

	if db.Loc == "" {
		db.Loc = "UTC"
	}

	if db.ConnectionTimeout == 0 {
		db.ConnectionTimeout = 10
	}

	connString := db.genMysqlConnString()
	db.d, err = sql.Open("mysql", connString)
	if err == nil {
		err = db.d.Ping()
		if err == nil {
			db.valid = true
			return nil
		}

	} else {
		return err
	}

	fmt.Println("------------db info--------------")
	fmt.Printf("%+v", db)
	if db.MaxConnLifetime != 0 {
		db.d.SetConnMaxLifetime(time.Duration(db.MaxConnLifetime) * time.Second)
	} else { // default one hout
		db.d.SetConnMaxLifetime(1 * time.Hour)
	}

	if db.MaxConnections != 0 {
		db.d.SetMaxIdleConns(db.MaxConnections)
		db.d.SetMaxOpenConns(db.MaxConnections)
	}

	go func() {
		for {
			time.Sleep(10 * time.Second)
			err := db.d.Ping()
			if err != nil {
				_ = fmt.Errorf("sql ping db fail:%s", err)
			}
		}
	}()

	return err
}

func (db *MysqlDBWrap) SetLoc(loc string) {
	db.Loc = url.QueryEscape(loc)
}

func (db *MysqlDBWrap) Valid() bool {
	return db.valid
}

func (db *MysqlDBWrap) DB() *sql.DB {
	if db.valid {
		return db.d
	}
	return nil
}

func M(key string) *sql.DB {
	db, ok := dbMysqlMap[key]
	if !ok {
		return nil
	}
	if !db.Valid() {
		return nil
	}
	return db.d
}

func defaultDB() *sql.DB {
	return dbMysqlMap["default"].d
}

func InitMysqlDB(key string, wrap *MysqlDBWrap) error {
	err := wrap.init()
	if err != nil {
		return err
	}

	dbMysqlMap[key] = wrap
	return nil
}
