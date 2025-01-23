package cache

import "fmt"

type CacheTag string

const (
    UserData CacheTag = "USER_DATA_%v"
)


func (c CacheTag) Format(args ...interface{}) string {
    return fmt.Sprintf(string(c), args...)
}