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

func newRemark(db *gorm.DB) remark {
	_remark := remark{}

	_remark.remarkDo.UseDB(db)
	_remark.remarkDo.UseModel(&model.Remark{})

	tableName := _remark.remarkDo.TableName()
	_remark.ALL = field.NewField(tableName, "*")
	_remark.AdminID = field.NewInt32(tableName, "admin_id")
	_remark.ApplicantID = field.NewInt32(tableName, "applicant_id")
	_remark.TheRemark = field.NewString(tableName, "the_remark")

	_remark.fillFieldMap()

	return _remark
}

type remark struct {
	remarkDo remarkDo

	ALL         field.Field
	AdminID     field.Int32
	ApplicantID field.Int32
	TheRemark   field.String

	fieldMap map[string]field.Expr
}

func (r remark) Table(newTableName string) *remark {
	r.remarkDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r remark) As(alias string) *remark {
	r.remarkDo.DO = *(r.remarkDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *remark) updateTableName(table string) *remark {
	r.ALL = field.NewField(table, "*")
	r.AdminID = field.NewInt32(table, "admin_id")
	r.ApplicantID = field.NewInt32(table, "applicant_id")
	r.TheRemark = field.NewString(table, "the_remark")

	r.fillFieldMap()

	return r
}

func (r *remark) WithContext(ctx context.Context) *remarkDo { return r.remarkDo.WithContext(ctx) }

func (r remark) TableName() string { return r.remarkDo.TableName() }

func (r remark) Alias() string { return r.remarkDo.Alias() }

func (r *remark) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *remark) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 3)
	r.fieldMap["admin_id"] = r.AdminID
	r.fieldMap["applicant_id"] = r.ApplicantID
	r.fieldMap["the_remark"] = r.TheRemark
}

func (r remark) clone(db *gorm.DB) remark {
	r.remarkDo.ReplaceDB(db)
	return r
}

type remarkDo struct{ gen.DO }

func (r remarkDo) Debug() *remarkDo {
	return r.withDO(r.DO.Debug())
}

func (r remarkDo) WithContext(ctx context.Context) *remarkDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r remarkDo) Clauses(conds ...clause.Expression) *remarkDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r remarkDo) Returning(value interface{}, columns ...string) *remarkDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r remarkDo) Not(conds ...gen.Condition) *remarkDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r remarkDo) Or(conds ...gen.Condition) *remarkDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r remarkDo) Select(conds ...field.Expr) *remarkDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r remarkDo) Where(conds ...gen.Condition) *remarkDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r remarkDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *remarkDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r remarkDo) Order(conds ...field.Expr) *remarkDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r remarkDo) Distinct(cols ...field.Expr) *remarkDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r remarkDo) Omit(cols ...field.Expr) *remarkDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r remarkDo) Join(table schema.Tabler, on ...field.Expr) *remarkDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r remarkDo) LeftJoin(table schema.Tabler, on ...field.Expr) *remarkDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r remarkDo) RightJoin(table schema.Tabler, on ...field.Expr) *remarkDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r remarkDo) Group(cols ...field.Expr) *remarkDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r remarkDo) Having(conds ...gen.Condition) *remarkDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r remarkDo) Limit(limit int) *remarkDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r remarkDo) Offset(offset int) *remarkDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r remarkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *remarkDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r remarkDo) Unscoped() *remarkDo {
	return r.withDO(r.DO.Unscoped())
}

func (r remarkDo) Create(values ...*model.Remark) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r remarkDo) CreateInBatches(values []*model.Remark, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r remarkDo) Save(values ...*model.Remark) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r remarkDo) First() (*model.Remark, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) Take() (*model.Remark, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) Last() (*model.Remark, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) Find() ([]*model.Remark, error) {
	result, err := r.DO.Find()
	return result.([]*model.Remark), err
}

func (r remarkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Remark, err error) {
	buf := make([]*model.Remark, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r remarkDo) FindInBatches(result *[]*model.Remark, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r remarkDo) Attrs(attrs ...field.AssignExpr) *remarkDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r remarkDo) Assign(attrs ...field.AssignExpr) *remarkDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r remarkDo) Joins(fields ...field.RelationField) *remarkDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r remarkDo) Preload(fields ...field.RelationField) *remarkDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r remarkDo) FirstOrInit() (*model.Remark, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) FirstOrCreate() (*model.Remark, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Remark), nil
	}
}

func (r remarkDo) FindByPage(offset int, limit int) (result []*model.Remark, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r remarkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *remarkDo) withDO(do gen.Dao) *remarkDo {
	r.DO = *do.(*gen.DO)
	return r
}