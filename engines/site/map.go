package site

import (
	"crypto/aes"
	"fmt"
	"time"

	"github.com/facebookgo/inject"
	"github.com/garyburd/redigo/redis"
	"github.com/itpkg/gails"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

func Logger() *logging.Logger {
	return logging.MustGetLogger("gails")
}

func Database() (*gorm.DB, error) {
	//postgresql: "user=%s password=%s host=%s port=%d dbname=%s sslmode=%s"
	args := ""
	for k, v := range viper.GetStringMap("database.extras") {
		args += fmt.Sprintf(" %s=%s ", k, v)
	}
	db, err := gorm.Open(viper.GetString("database.adapter"), args)
	if err != nil {
		return nil, err
	}
	if !gails.IsProduction() {
		db.LogMode(true)
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(viper.GetInt("database.pool.max_idle"))
	db.DB().SetMaxOpenConns(viper.GetInt("database.pool.max_open"))
	return db, nil
}

func Redis() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial(
				"tcp",
				fmt.Sprintf(
					"%s:%d",
					viper.GetString("redis.host"),
					viper.GetInt("redis.port"),
				),
			)
			if e != nil {
				return nil, e
			}
			if _, e = c.Do("SELECT", viper.GetInt("redis.db")); e != nil {
				c.Close()
				return nil, e
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Aes() (*gails.Aes, error) {
	cip, err := aes.NewCipher([]byte(viper.GetString("secrets")[80:112]))
	if err != nil {
		return nil, err
	}
	return &gails.Aes{Cip: cip}, nil
}

func (p *Engine) Map(inj *inject.Graph) error {
	db, err := Database()
	if err != nil {
		return err
	}
	_aes, err := Aes()
	if err != nil {
		return err
	}

	inj.Provide(
		&inject.Object{
			Value: Redis(),
		},
		&inject.Object{
			Value: db,
		},
		&inject.Object{
			Value: _aes,
		},
		&inject.Object{
			Value: Logger(),
		},
	)
	return nil
}
