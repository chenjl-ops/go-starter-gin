package apollo

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
	"os"
)

const (
	runEnvKey     = "RUNTIME_ENV"
	runClusterKey = "RUNTIME_CLUSTER"
	appNameKey    = "RUNTIME_APP_NAME"
)

var Config *Specification

//获取当前运行环境
func GetEnv() string {
	env := os.Getenv(runEnvKey)
	return env
}

//获取当前运行应用名称
func GetAppName() string {
	appName := os.Getenv(appNameKey)
	return appName
}

//获取当前运行cluster name
func GetRunCluster() string {
	clusterName := os.Getenv(runClusterKey)
	return clusterName
}

func GetApolloUrl(env string) string {
	url := fmt.Sprintf("http://configserver-%s.chj.cloud", env)
	// URL Apollo 地址
	// url, _ := fmt.Printf("http://configserver-%s.xxx.xxx", env)
	return url
}

func NewApollo() (apollo *Apollo, err error) {
	appName := GetAppName()
	if appName == "" {
		return nil, errors.Errorf("Env appName Has Empty")
	}
	env := GetEnv()
	if env == "" {
		return nil, errors.Errorf("Env env Has Empty")
	}
	cluster := GetRunCluster()
	if cluster == "" {
		return nil, errors.Errorf("Env cluster Has Empty")
	}
	apolloServerUrl := GetApolloUrl(env)
	// apolloServerUrl := apolloConfigServers[env]
	// fmt.Println("URL:", apolloServerUrl)
	apollo = &Apollo{
		AppID:           appName,
		Cluster:         cluster,
		NameSpaceName:   "application",
		ApolloServerUrl: apolloServerUrl,
	}

	return apollo, nil
}

func ReadRemoteConfig() error {
	return ReadRemoteConfigCustom(nil)
}

func ReadRemoteConfigCustom(input *Apollo) error {
	config, err := NewApollo()
	if err != nil {
		return err
	}
	// 以输入参数 覆盖New对象
	if input != nil {
		if input.Cluster != "" {
			config.Cluster = input.Cluster
		}
		if input.NameSpaceName != "" {
			config.NameSpaceName = input.NameSpaceName
		}
		if input.ApolloServerUrl != "" {
			config.ApolloServerUrl = input.ApolloServerUrl
		}
		if input.AppID != "" {
			config.AppID = input.AppID
		}
	}
	remote.SetAppID(config.AppID)
	v := viper.New()
	v.SetConfigType("prop") // prop为固定写法
	err = v.AddRemoteProvider("apollo", config.ApolloServerUrl, config.NameSpaceName)
	if err != nil {
		return err
	}
	err = v.ReadRemoteConfig()
	if err != nil {
		return err
	}
	err = v.Unmarshal(&Config)
	if err != nil {
		return err
	}
	// 调试apollo配置信息使用
	// fmt.Printf("%+v\n", config)
	fmt.Printf("apollo data: %+v", Config)
	//fmt.Println("redis_sentinels:", v.GetString("redis_sentinels"))

	//监听配置变化
	//v.WatchRemoteConfigOnChannel()
	//for {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("app.AllSettings:", v.AllSettings())
	//}
	return nil
}
