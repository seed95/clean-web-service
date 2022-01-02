package user

type (
	User struct {
		ID          uint
		Username    string
		Password    string
		Firstname   string
		Lastname    string
		Email       string
		PhoneNumber string
		Gender      Gender
		Role        Role
	}
)
