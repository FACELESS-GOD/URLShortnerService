package Model

import (
	"errors"

	"github.com/FACELESS-GOD/URLShortnerService.git/Package/Utility"
	"github.com/redis/go-redis/v9"
)

func GetURL(Id string) (string, error) {

	old_URL, err := Utility.RedisInstance.Get(Utility.Context, Id).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}

	if err == redis.Nil {
		err := errors.New("Invalid URL")
		return "", err
	} else {
		return old_URL, nil
	}

}
