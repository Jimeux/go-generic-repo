package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"

	"github.com/Jimeux/go-generic-repo/data"
)

var (
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")
)

func main() {
	db, err := xorm.NewEngine("mysql", fmt.Sprintf("root:@(:%s)/%s", dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Job
	jobRepo := data.NewJobRepository(db)
	job := &data.Job{
		Name: "Job 1",
	}
	if err := jobRepo.Create(ctx, job); err != nil {
		log.Fatal(err)
	}
	j, err := jobRepo.GetByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", j)
	count, err := jobRepo.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count %d\n", count)

	// Person
	personRepo := data.NewPersonRepository(db)
	person := &data.Person{
		Name: "Job 1",
	}
	if err := personRepo.Create(ctx, person); err != nil {
		log.Fatal(err)
	}
	p, err := personRepo.GetByID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
	pcount, err := personRepo.Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count %d\n", pcount)
}

func tx(err error, ctx context.Context, db *xorm.Engine, personRepo *data.PersonRepository, jobRepo *data.JobRepository) {
	person, err := data.Tx(ctx, db, func(ctx context.Context) (*data.Job, error) {
		if err := personRepo.Create(ctx, &data.Person{Name: "Person"}); err != nil {
			return nil, err
		}
		jb := &data.Job{Name: "Job"}
		if err := jobRepo.Create(ctx, jb); err != nil {
			return nil, err
		}
		return jb, nil
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = person
}
