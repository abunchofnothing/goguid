/*
Generate a random, probably unique 32-digit GUID (with dashes)

By CJ Thompson
Nov 07, 2014
*/
package guid

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func uniqid() string {
	rand.Seed(time.Now().UnixNano())
	hash := md5.New()
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	for i := 0; i < 128; i++ {
		io.WriteString(hash, string(chars[rand.Intn(len(chars))]))
	}

	return string(hash.Sum(nil))
}

func Generate() string {
	str := uniqid()

	return fmt.Sprintf("%x-%x-%x-%x-%x", str[0:4], str[4:6], str[6:8], str[8:10], str[10:16])
}
