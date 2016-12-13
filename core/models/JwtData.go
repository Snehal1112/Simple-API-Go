package models



type JwtData struct {
	UserName    string    `boil:"username" json:"username" toml:"username" yaml:"username"`
	PhoneNumber string    `boil:"phone_number" json:"phone_number" toml:"phone_number" yaml:"phone_number"`
	Email       string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	Name    	string 	   `boil:"name" json:"name" toml:"name" yaml:"name"`
	Role         string    `boil:"role" json:"role" toml:"role" yaml:"role"`
	AkunId		int			`boil:"akun_id" json:"akun_id" toml:"akun_id" yaml:"akun_id"`
	ProfileId	int 		`boil:"profile_id" json:"profile_id" toml:"profile_id" yaml:"profile_id"`
}
 