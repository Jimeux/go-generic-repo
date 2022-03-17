package data

import "xorm.io/xorm"

type Job struct {
	ID   int64  `xorm:"id"`
	Name string `xorm:"name"`
}

type JobRepository struct {
	*Repository[Job, int64]
}

func NewJobRepository(db *xorm.Engine) *JobRepository {
	return &JobRepository{
		Repository: NewRepository[Job, int64](db),
	}
}
