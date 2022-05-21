// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/scutrobotlab/HR/dal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Group{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Group{}) fail: %s", err)
	}
}

func Test_groupQuery(t *testing.T) {
	group := newGroup(db)
	group = *group.As(group.TableName())
	do := group.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(group.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <groups> fail:", err)
		return
	}

	_, ok := group.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from group success")
	}

	err = do.Create(&model.Group{})
	if err != nil {
		t.Error("create item in table <groups> fail:", err)
	}

	err = do.Save(&model.Group{})
	if err != nil {
		t.Error("create item in table <groups> fail:", err)
	}

	err = do.CreateInBatches([]*model.Group{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <groups> fail:", err)
	}

	_, err = do.Select(group.ALL).Take()
	if err != nil {
		t.Error("Take() on table <groups> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <groups> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <groups> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <groups> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Group{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <groups> fail:", err)
	}

	_, err = do.Select(group.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <groups> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <groups> fail:", err)
	}

	_, err = do.Select(group.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <groups> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <groups> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <groups> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <groups> fail:", err)
	}

	_, err = do.ScanByPage(&model.Group{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <groups> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <groups> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <groups> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <groups> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <groups> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <groups> fail:", err)
	}
}