package main

// Config - configuration of the application
type Config struct {
	DbType     string `default:"sqlite" split_words:"true"` // sqlite or mysql
	DbHost     string `required:"false" split_words:"true"`
	DbName     string `required:"true" split_words:"true"`
	DbPort     string `default:"3306" split_words:"true"`
	DbUser     string `required:"false" split_words:"true"`
	DbPassword string `required:"false" split_words:"true"`
}
