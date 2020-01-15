package dbproject

//Config ...
type Config struct {
	BindAddr    string `required:"true" split_words:"true"`
	LogLevel    string `required:"true" split_words:"true"`
	DatabaseURL string `required:"true" split_words:"true"`
	SessionKey  string `required:"true" split_words:"true"`
	SessionPath string `required:"true" split_words:"true"`
}
