// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/scutrobotlab/HR/dal/model"
)

func newApplicant(db *gorm.DB) applicant {
	_applicant := applicant{}

	_applicant.applicantDo.UseDB(db)
	_applicant.applicantDo.UseModel(&model.Applicant{})

	tableName := _applicant.applicantDo.TableName()
	_applicant.ALL = field.NewField(tableName, "*")
	_applicant.ID = field.NewUint(tableName, "id")
	_applicant.CreatedAt = field.NewTime(tableName, "created_at")
	_applicant.UpdatedAt = field.NewTime(tableName, "updated_at")
	_applicant.DeletedAt = field.NewField(tableName, "deleted_at")
	_applicant.Name = field.NewString(tableName, "name")
	_applicant.Gender = field.NewBool(tableName, "gender")
	_applicant.Phone = field.NewString(tableName, "phone")
	_applicant.Avatar = field.NewString(tableName, "avatar")
	_applicant.Form = field.NewString(tableName, "form")

	_applicant.fillFieldMap()

	return _applicant
}

type applicant struct {
	applicantDo applicantDo

	ALL       field.Field
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Gender    field.Bool
	Phone     field.String
	Avatar    field.String
	Form      field.String

	fieldMap map[string]field.Expr
}

func (a applicant) Table(newTableName string) *applicant {
	a.applicantDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a applicant) As(alias string) *applicant {
	a.applicantDo.DO = *(a.applicantDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *applicant) updateTableName(table string) *applicant {
	a.ALL = field.NewField(table, "*")
	a.ID = field.NewUint(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Name = field.NewString(table, "name")
	a.Gender = field.NewBool(table, "gender")
	a.Phone = field.NewString(table, "phone")
	a.Avatar = field.NewString(table, "avatar")
	a.Form = field.NewString(table, "form")

	a.fillFieldMap()

	return a
}

func (a *applicant) WithContext(ctx context.Context) *applicantDo {
	return a.applicantDo.WithContext(ctx)
}

func (a applicant) TableName() string { return a.applicantDo.TableName() }

func (a applicant) Alias() string { return a.applicantDo.Alias() }

func (a *applicant) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *applicant) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["name"] = a.Name
	a.fieldMap["gender"] = a.Gender
	a.fieldMap["phone"] = a.Phone
	a.fieldMap["avatar"] = a.Avatar
	a.fieldMap["form"] = a.Form
}

func (a applicant) clone(db *gorm.DB) applicant {
	a.applicantDo.ReplaceDB(db)
	return a
}

type applicantDo struct{ gen.DO }

func (a applicantDo) Debug() *applicantDo {
	return a.withDO(a.DO.Debug())
}

func (a applicantDo) WithContext(ctx context.Context) *applicantDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a applicantDo) Clauses(conds ...clause.Expression) *applicantDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a applicantDo) Returning(value interface{}, columns ...string) *applicantDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a applicantDo) Not(conds ...gen.Condition) *applicantDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a applicantDo) Or(conds ...gen.Condition) *applicantDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a applicantDo) Select(conds ...field.Expr) *applicantDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a applicantDo) Where(conds ...gen.Condition) *applicantDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a applicantDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *applicantDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a applicantDo) Order(conds ...field.Expr) *applicantDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a applicantDo) Distinct(cols ...field.Expr) *applicantDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a applicantDo) Omit(cols ...field.Expr) *applicantDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a applicantDo) Join(table schema.Tabler, on ...field.Expr) *applicantDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a applicantDo) LeftJoin(table schema.Tabler, on ...field.Expr) *applicantDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a applicantDo) RightJoin(table schema.Tabler, on ...field.Expr) *applicantDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a applicantDo) Group(cols ...field.Expr) *applicantDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a applicantDo) Having(conds ...gen.Condition) *applicantDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a applicantDo) Limit(limit int) *applicantDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a applicantDo) Offset(offset int) *applicantDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a applicantDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *applicantDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a applicantDo) Unscoped() *applicantDo {
	return a.withDO(a.DO.Unscoped())
}

func (a applicantDo) Create(values ...*model.Applicant) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a applicantDo) CreateInBatches(values []*model.Applicant, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a applicantDo) Save(values ...*model.Applicant) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a applicantDo) First() (*model.Applicant, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Applicant), nil
	}
}

func (a applicantDo) Take() (*model.Applicant, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Applicant), nil
	}
}

func (a applicantDo) Last() (*model.Applicant, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Applicant), nil
	}
}

func (a applicantDo) Find() ([]*model.Applicant, error) {
	result, err := a.DO.Find()
	return result.([]*model.Applicant), err
}

func (a applicantDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Applicant, err error) {
	buf := make([]*model.Applicant, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a applicantDo) FindInBatches(result *[]*model.Applicant, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a applicantDo) Attrs(attrs ...field.AssignExpr) *applicantDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a applicantDo) Assign(attrs ...field.AssignExpr) *applicantDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a applicantDo) Joins(fields ...field.RelationField) *applicantDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a applicantDo) Preload(fields ...field.RelationField) *applicantDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a applicantDo) FirstOrInit() (*model.Applicant, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Applicant), nil
	}
}

func (a applicantDo) FirstOrCreate() (*model.Applicant, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Applicant), nil
	}
}

func (a applicantDo) FindByPage(offset int, limit int) (result []*model.Applicant, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a applicantDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a *applicantDo) withDO(do gen.Dao) *applicantDo {
	a.DO = *do.(*gen.DO)
	return a
}