package apollo

//apollo自身结构配置
type Apollo struct {
	AppID           string
	Cluster         string
	NameSpaceName   string
	ApolloServerUrl string
}

//apollo内容配置
type Specification struct {
	RedisSentinels string `envconfig:"REDIS_SENTINELS" mapstructure:"redis_sentinels"`
	RedisCluster   string `envconfig:"REDIS_CLUSTER" mapstructure:"redis_cluster"`
}

