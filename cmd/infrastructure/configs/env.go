package configs

type Config struct {
	DBHost     string `mapstructure:"mariadb_host"`
	DBUserName string `mapstructure:"mariadb_user"`
	DBPassword string `mapstructure:"mariadb_password"`
	DBName     string `mapstructure:"mariadb_db"`
	DBPort     string `mapstructure:"mariadb_port"`

	Environment string `toml:"ENVIRONMENT"`

	AWSAccessKeyID     string `mapstructure:"aws_access_key_id"`
	AWSSecretAccessKey string `mapstructure:"aws_secret_access_key"`
	AWSRegion          string `mapstructure:"aws_region"`
	S3BucketName       string `mapstructure:"s3_bucket_name"`
}
