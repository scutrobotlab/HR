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
	err := db.AutoMigrate(&model.Admin{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Admin{}) fail: %s", err)
	}
}

func Test_adminQuery(t *testing.T) {
	admin := newAdmin(db)
	admin = *admin.As(admin.TableName())
	do := admin.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(admin.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <admin> fail:", err)
		return
	}

	_, ok := admin.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from admin success")
	}

	err = do.Create(&model.Admin{})
	if err != nil {
		t.Error("create item in table <admin> fail:", err)
	}

	err = do.Save(&model.Admin{})
	if err != nil {
		t.Error("create item in table <admin> fail:", err)
	}

	err = do.CreateInBatches([]*model.Admin{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <admin> fail:", err)
	}

	_, err = do.Select(admin.ALL).Take()
	if err != nil {
		t.Error("Take() on table <admin> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <admin> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <admin> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <admin> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Admin{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <admin> fail:", err)
	}

	_, err = do.Select(admin.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <admin> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <admin> fail:", err)
	}

	_, err = do.Select(admin.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <admin> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <admin> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <admin> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <admin> fail:", err)
	}

	_, err = do.ScanByPage(&model.Admin{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <admin> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <admin> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <admin> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <admin> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <admin> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <admin> fail:", err)
	}
}