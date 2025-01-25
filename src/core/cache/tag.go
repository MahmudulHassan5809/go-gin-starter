package cache

import "fmt"

type CacheTag string

const (
    UserData CacheTag = "USER_DATA_%v"
    UserAccessToken CacheTag = "USER_ACCESS_TOKEN_%v"
    UserRefreshToken CacheTag = "USER_REFRESH_TOKEN_%v"
)


func (c CacheTag) Format(args ...interface{}) string {
    return fmt.Sprintf(string(c), args...)
}