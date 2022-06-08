package globals

import "github.com/spf13/viper"

var Secret = []byte(viper.GetString("SECRET"))

const Userkey = "user"
