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
	err := db.AutoMigrate(&model.Remark{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Remark{}) fail: %s", err)
	}
}

func Test_remarkQuery(t *testing.T) {
	remark := newRemark(db)
	remark = *remark.As(remark.TableName())
	do := remark.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(remark.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <remarks> fail:", err)
		return
	}

	_, ok := remark.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from remark success")
	}

	err = do.Create(&model.Remark{})
	if err != nil {
		t.Error("create item in table <remarks> fail:", err)
	}

	err = do.Save(&model.Remark{})
	if err != nil {
		t.Error("create item in table <remarks> fail:", err)
	}

	err = do.CreateInBatches([]*model.Remark{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <remarks> fail:", err)
	}

	_, err = do.Select(remark.ALL).Take()
	if err != nil {
		t.Error("Take() on table <remarks> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <remarks> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <remarks> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <remarks> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Remark{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <remarks> fail:", err)
	}

	_, err = do.Select(remark.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <remarks> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <remarks> fail:", err)
	}

	_, err = do.Select(remark.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <remarks> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <remarks> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <remarks> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <remarks> fail:", err)
	}

	_, err = do.ScanByPage(&model.Remark{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <remarks> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <remarks> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <remarks> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <remarks> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <remarks> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <remarks> fail:", err)
	}
}