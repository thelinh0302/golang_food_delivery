package usermodel

import (
	"Tranning_food/common"
	"Tranning_food/component/digitProvider"
	"Tranning_food/component/tokenprovider"
	"errors"
	"strings"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"-" gorm:"column:role;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (res *UserOTP) ValidateOTP() error {
	res.Phone = strings.TrimSpace(res.Phone)
	if len(res.Phone) == 0 {
		return common.NewCustomError(
			errors.New("phone has already required"),
			"phone has already required",
			"errPhoneRequired",
		)
	}
	return nil
}

func (res *UserCreate) Validate() error {

	res.Email = strings.TrimSpace(res.Email)
	res.Password = strings.TrimSpace(res.Password)
	res.Phone = strings.TrimSpace(res.Phone)
	if len(res.Email) == 0 {
		return common.NewCustomError(
			errors.New("email has already required"),
			"email has already required",
			"ErrEmailRequired",
		)
	} else if len(res.Password) == 0 {
		return common.NewCustomError(
			errors.New("password has already required"),
			"password has already required",
			"ErrPasswordRequired",
		)
	} else if len(res.Phone) == 0 {
		return common.NewCustomError(
			errors.New("phone has already required"),
			"phone has already required",
			"ErrPhoneRequired",
		)
	}
	return nil
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

type UserOTP struct {
	Phone string `json:"phone" form: "phone" gorm:"column:phone;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

type ResOTP struct {
	AccessToken *tokenprovider.Token `json:"access_token"`
	Otp         *digitProvider.Otp   `json:"otp"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrPhoneOrPasswordInvalid = common.NewCustomError(
		errors.New("phone or password invalid"),
		"phone or password invalid",
		"ErrphoneOrPasswordInvalid",
	)

	ErrPhoneInvalid = common.NewCustomError(
		errors.New("phone does not exist"),
		"phone does not exist",
		"ErrPhoneInvalid",
	)

	ErrEmailorPhoneExisted = common.NewCustomError(
		errors.New("email/phone has already existed"),
		"email/phone has already existed",
		"ErrEmailorPhoneExisted",
	)
	ErrPhoneExisted = common.NewCustomError(
		errors.New("phone has already existed"),
		"phone has already existed",
		"ErrPhoneExisted",
	)

	ErrRequiredEmailorPassword = common.NewCustomError(
		errors.New("Please enter your email or password"),
		"Please enter your email or password",
		"ErrRequiredEmailorPassword",
	)
)
