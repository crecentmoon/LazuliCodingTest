package infra

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/crecentmoon/lazuli-coding-test/internal/app/repository"
	"github.com/crecentmoon/lazuli-coding-test/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlHandler struct {
	db *gorm.DB
}

var txKey = struct{}{}

func NewMySQLHandler(cfg config.Config) (repository.SqlHandler, error) {
	dbURL := cfg.Database.Host
	port := cfg.Database.Port
	dbName := cfg.Database.DBName
	user := cfg.Database.Username
	password := cfg.Database.Password

	dsn := user + ":" + password + "@tcp(" + dbURL + ":" + strconv.Itoa(port) + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Println(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db

	return sqlHandler, err
}

func (handler *SqlHandler) Execute(ctx context.Context, query string, params ...interface{}) (uint, error) {
	tx, ok := ctx.Value(&txKey).(*sql.Tx)
	if ok {
		res, err := tx.Exec(query, params...)
		if err != nil {
			return 0, err
		}

		rows, err := res.RowsAffected()
		if rows < 0 || err != nil {
			return 0, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}

		return uint(id), nil
	}

	db, err := handler.db.DB()
	if err != nil {
		return 0, err
	}

	res, err := db.Exec(query, params...)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()
	if rows < 0 || err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func (handler *SqlHandler) Query(obj interface{}, sql string, params ...interface{}) error {
	if err := handler.db.Raw(sql, params...).Scan(obj).Error; err != nil {
		return err
	}

	return nil
}

func (handler *SqlHandler) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	db, err := handler.db.DB()
	if err != nil {
		return nil, err
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, &txKey, tx)

	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return v, nil
}
