package utility

import(
	"golang.org/x/crypto/bcrypt"
)


func Encode(str string,cost int)(hashedString string , err error){

	res,err := bcrypt.GenerateFromPassword([]byte(str),cost)
	if err != nil {
		return string(res),err
	}
	return string(res),nil
}

