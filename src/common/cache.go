package common

import (
	"bytes"
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	blog "github.com/beego/bee/logger"
	"github.com/pkg/errors"
	"time"
)

var cc cache.Cache

func InitCache() {

	cacheConfig := beego.AppConfig.String("cache.cache")
	cc = nil
	if "redis" == cacheConfig {
		initRedis()
	}

}

func initRedis() {

	blog.Log.Info("redis start ...")

	var err error

	defer func() {
		if r := recover(); r != nil {
			cc = nil
		}
	}()

	host := beego.AppConfig.String("cache.redis.host")
	password := beego.AppConfig.String("cache.redis.password")
	blog.Log.Infof("info", "connect redis param :"+host)

	cc, err = cache.NewCache("redis", `{"conn":"`+host+`","dbNum":"0","password":"`+password+`"}`)

	if err != nil {
		blog.Log.Errorf("connect redis failed ", err.Error())
	}

}

// SetCache 设置缓存
func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cache.cache is null ")
	}

	defer func() {
		if r := recover(); r != nil {
			blog.Log.Errorf("error", r)
			cc = nil
		}
	}()

	timeouts := time.Duration(timeout) * time.Second
	err = cc.Put(key, data, timeouts)

	if err != nil {
		blog.Log.Errorf("info", "SetCache failed key:"+key)
		return err
	}
	return err
}

// GetCache 获得缓存信息
func GetCache(key string, to interface{}) error {

	if cc == nil {
		return errors.New("cache.cache is null")
	}
	defer func() {
		if r := recover(); r != nil {
			blog.Log.Errorf("error", r)
			cc = nil
		}
	}()

	data := cc.Get(key)
	if data == nil {
		blog.Log.Infof("no key exists %s ", key)
		return nil
	}
	err := Decode(data.([]byte), to)
	if err != nil {
		blog.Log.Errorf("error", err)
		blog.Log.Errorf("error", "GetCache failed key:"+key)
	}
	return nil
}

func DeleteCache(key string) (err error) {
	if cc == nil {
		return errors.New("cache.cache is null")
	}
	defer func() {
		if r := recover(); r != nil {
			blog.Log.Errorf("error", r)
			cc = nil
		}
	}()

	err = cc.Delete(key)
	if err != nil {
		return errors.New("Cache delete failed key " + key)
	}
	return err
}

// Encode 用gob进行数据编码
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode 用gob进行数据解码
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
