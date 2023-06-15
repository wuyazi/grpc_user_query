package svc

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/wuyazi/grpc_user_query/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

type ServiceContext struct {
	Config config.Config
	ReadDB sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := openDB(c)
	if err != nil {
		logx.Errorf("query openDB error: %v", err)
		panic(err)
	}
	readDB := sqlx.NewSqlConnFromDB(db)
	prometheus.MustRegister(collectors.NewDBStatsCollector(db, "query-ReadDB"))

	logx.DisableStat()
	return &ServiceContext{
		Config: c,
		ReadDB: readDB,
	}
}

func openDB(c config.Config) (db *sql.DB, err error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local&interpolateParams=true",
		c.DataBase.User, c.DataBase.Password, c.DataBase.Host, c.DataBase.Port, c.DataBase.Db)
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		return nil, fmt.Errorf("openDB error: %w", err)
	}
	db.SetConnMaxLifetime(time.Second * 300)
	return db, nil
}
