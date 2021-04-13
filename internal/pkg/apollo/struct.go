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
	// ServerRunPort        string `envconfig:"SERVER_RUN_PORT" mapstructure:"server_run_port"`
	MysqlUserName        string `envconfig:"MYSQL_USERNAME" mapstructure:"mysql_db_user"`
	MysqlPassword        string `envconfig:"MYSQL_PASSWORD" mapstructure:"mysql_db_passwd"`
	MysqlHost            string `envconfig:"MYSQL_HOST" mapstructure:"mysql_db_host"`
	MysqlPort            int    `envconfig:"MYSQL_PORT" mapstructure:"mysql_db_port"`
	MysqlDBName          string `envconfig:"MYSQL_DBNAME" mapstructure:"mysql_db_name"`
	RedisMasterName      string `envconfig:"REDIS_MASTER_NAME" mapstructure:"redis_cluster"`
	RedisSentinelAddress string `envconfig:"REDIS_SENTINEL_ADDRESS" mapstructure:"redis_sentinels"`
	RedisPasswd          string `envconfig:"REDIS_PASSWD" mapstructure:"redis_password"`
}
