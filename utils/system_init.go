package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app inited")
}
func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	fmt.Println("mysql inited")
	// user := models.UserBasic{}
	// DB.Find(&user)
	// fmt.Println(user)
}
func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	// pong, err := Red.Ping(context.Background()).Result()
	// if err != nil {
	// 	fmt.Println("init redis 。。。。", err)
	// } else {
	// 	fmt.Println(" Redis inited 。。。。", pong)
	// }
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish...", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe...", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe...", msg.Payload)
	return msg.Payload, err
}
