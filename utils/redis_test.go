package utils_test

import (
	"context"
	"fmt"
	"peony/utils"
	"testing"
)

func TestRedis(t *testing.T) {
	var ctx context.Context
	utils.InitRedis()
	utils.RDB.HSet(ctx, "Msg", "email", "123456@qq.com", "err", "")
	defer utils.RDB.Del(ctx, "Msg")

	c := utils.RDB.HGet(ctx, "Msg", "email").String()
	fmt.Println(c)
}
