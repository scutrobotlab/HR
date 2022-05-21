package biz

import (
	"context"
	"fmt"

	"github.com/scutrobotlab/HR/dal/query"
)

var q = query.Q

func Query(ctx context.Context) {
	admin := q.Admin
	do := admin.WithContext(ctx)

	data, err := do.Take()
	catchError("Take", err)
	fmt.Printf("got %+v\n", data)

	dataArray, err := do.Find()
	catchError("Find", err)
	fmt.Printf("got %+v\n", dataArray)

	data, err = do.Where(admin.ID.Eq(1)).Take()
	catchError("Take", err)
	fmt.Printf("got %+v\n", data)

	// dataArray, err = do.Where(admin.Name.Gt(18)).Order(admin.Username).Find()
	// catchError("Find", err)
	// fmt.Printf("got %+v\n", dataArray)

	dataArray, err = do.Select(admin.ID, admin.Name).Order(admin.ID.Desc()).Find()
	catchError("Find", err)
	fmt.Printf("got %+v\n", dataArray)

	// info, err := do.Where(admin.ID.Eq(1)).UpdateSimple(admin.Age.Add(1))
	// catchError("Update", err)
	// fmt.Printf("got %+v\n", info)
}

func catchError(detail string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", detail, err)
	}
}
