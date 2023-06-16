package configs

type (
	Config struct {
		Mysql  Mysql  `mapstructure:"mysql"`
		Server Server `mapstructure:"server"`
	}

	Mysql struct {
		Host             string `mapstructure:"host"`
		Port             int    `mapstructure:"port"`
		DBName           string `mapstructure:"db_name"`
		Username         string `mapstructure:"username"`
		Password         string `mapstructure:"password"`
		MaxOpenConns     int    `mapstructure:"max_open_conns"`
		MaxIdleConns     int    `mapstructure:"max_idle_conns"`
		MaxLifetime      int    `mapstructure:"max_life_time"`
		TablePrefix      string `mapstructure:"table_prefix"`
		SingularTable    bool   `mapstructure:"singular_table"`
		MaxExecutionTime int    `mapstructure:"max_execution_time"`
	}

	Server struct {
		Port           string `mapstructure:"port"`
		Env            string `mapstructure:"env"`
		LogLevel       string `mapstructure:"log_level"`
		PrometheusPort string `mapstructure:"prometheus_port"`
	}
)
