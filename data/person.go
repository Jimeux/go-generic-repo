package data

import "xorm.io/xorm"

type Person struct {
	ID         int64  `xorm:"id"`
	GivenName  string `xorm:"given_name"`
	FamilyName string `xorm:"family_name"`
}

type PersonRepository struct {
	*Repository[Person, int64]
}

func NewPersonRepository(db *xorm.Engine) *PersonRepository {
	return &PersonRepository{
		Repository: NewRepository[Person, int64](db),
	}
}
