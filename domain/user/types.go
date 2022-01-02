package user

import (
	"fmt"
	"github.com/seed95/clean-web-service/build/messages"
	"github.com/seed95/clean-web-service/pkg/derrors"
)

type (
	Gender uint
	Role   uint
)

const (
	_ Gender = iota
	Male
	Female
	More
)

const (
	_ Role = iota
	Basic
	Admin
	Seller
	Manager
)

var (
	roleString = map[Role]string{
		Basic:   "basic",
		Admin:   "admin",
		Seller:  "seller",
		Manager: "manager",
	}

	stringRole = map[string]Role{
		roleString[Basic]:   Basic,
		roleString[Admin]:   Admin,
		roleString[Seller]:  Seller,
		roleString[Manager]: Manager,
	}

	genderString = map[Gender]string{
		Male:   "male",
		Female: "female",
		More:   "more",
	}

	stringGender = map[string]Gender{
		genderString[Male]:   Male,
		genderString[Female]: Female,
		genderString[More]:   More,
	}
)

func (r Role) String() string {

	if s, ok := roleString[r]; ok {
		return s
	}

	return fmt.Sprintf("Role(%d)", r)

}

func (r Role) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Role) UnmarshalText(b []byte) error {

	if role, ok := stringRole[string(b)]; ok {
		*r = role
		return nil
	}

	return derrors.New(derrors.Invalid, messages.InvalidRole)

}

func (g Gender) String() string {

	if s, ok := genderString[g]; ok {
		return s
	}

	return fmt.Sprintf("Gender(%d)", g)
}

func (g Gender) MarshalText() ([]byte, error) {
	return []byte(g.String()), nil
}

func (g *Gender) UnmarshalText(b []byte) error {

	if gender, ok := stringGender[string(b)]; ok {
		*g = gender
		return nil
	}

	return derrors.New(derrors.Invalid, messages.InvalidGender)

}
