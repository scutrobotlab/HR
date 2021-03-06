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

func newScore(db *gorm.DB) score {
	_score := score{}

	_score.scoreDo.UseDB(db)
	_score.scoreDo.UseModel(&model.Score{})

	tableName := _score.scoreDo.TableName()
	_score.ALL = field.NewField(tableName, "*")
	_score.AdminID = field.NewUint(tableName, "admin_id")
	_score.ApplicantID = field.NewUint(tableName, "applicant_id")
	_score.Group = field.NewString(tableName, "group")
	_score.Score = field.NewFloat64(tableName, "score")
	_score.StandardID = field.NewUint(tableName, "standard_id")
	_score.EvaluationDetails = field.NewString(tableName, "evaluation_details")
	_score.UpdatedAt = field.NewTime(tableName, "updated_at")
	_score.Admin = scoreAdmin{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Admin", "model.Admin"),
		Standard: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Admin.Standard", "model.Standard"),
		},
	}

	_score.Applicant = scoreApplicant{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Applicant", "model.Applicant"),
		Intents: struct {
			field.RelationField
			Applicant struct {
				field.RelationField
			}
			OptionalTime struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Applicant.Intents", "model.Intent"),
			Applicant: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Applicant.Intents.Applicant", "model.Applicant"),
			},
			OptionalTime: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Applicant.Intents.OptionalTime", "model.OptionalTime"),
			},
		},
		Answers: struct {
			field.RelationField
			Applicant struct {
				field.RelationField
			}
			Question struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Applicant.Answers", "model.Answer"),
			Applicant: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Applicant.Answers.Applicant", "model.Applicant"),
			},
			Question: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Applicant.Answers.Question", "model.Question"),
			},
		},
	}

	_score.Standard = scoreStandard{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Standard", "model.Standard"),
	}

	_score.fillFieldMap()

	return _score
}

type score struct {
	scoreDo scoreDo

	ALL               field.Field
	AdminID           field.Uint
	ApplicantID       field.Uint
	Group             field.String
	Score             field.Float64
	StandardID        field.Uint
	EvaluationDetails field.String
	UpdatedAt         field.Time
	Admin             scoreAdmin

	Applicant scoreApplicant

	Standard scoreStandard

	fieldMap map[string]field.Expr
}

func (s score) Table(newTableName string) *score {
	s.scoreDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s score) As(alias string) *score {
	s.scoreDo.DO = *(s.scoreDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *score) updateTableName(table string) *score {
	s.ALL = field.NewField(table, "*")
	s.AdminID = field.NewUint(table, "admin_id")
	s.ApplicantID = field.NewUint(table, "applicant_id")
	s.Group = field.NewString(table, "group")
	s.Score = field.NewFloat64(table, "score")
	s.StandardID = field.NewUint(table, "standard_id")
	s.EvaluationDetails = field.NewString(table, "evaluation_details")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *score) WithContext(ctx context.Context) *scoreDo { return s.scoreDo.WithContext(ctx) }

func (s score) TableName() string { return s.scoreDo.TableName() }

func (s score) Alias() string { return s.scoreDo.Alias() }

func (s *score) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *score) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 10)
	s.fieldMap["admin_id"] = s.AdminID
	s.fieldMap["applicant_id"] = s.ApplicantID
	s.fieldMap["group"] = s.Group
	s.fieldMap["score"] = s.Score
	s.fieldMap["standard_id"] = s.StandardID
	s.fieldMap["evaluation_details"] = s.EvaluationDetails
	s.fieldMap["updated_at"] = s.UpdatedAt

}

func (s score) clone(db *gorm.DB) score {
	s.scoreDo.ReplaceDB(db)
	return s
}

type scoreAdmin struct {
	db *gorm.DB

	field.RelationField

	Standard struct {
		field.RelationField
	}
}

func (a scoreAdmin) Where(conds ...field.Expr) *scoreAdmin {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a scoreAdmin) WithContext(ctx context.Context) *scoreAdmin {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a scoreAdmin) Model(m *model.Score) *scoreAdminTx {
	return &scoreAdminTx{a.db.Model(m).Association(a.Name())}
}

type scoreAdminTx struct{ tx *gorm.Association }

func (a scoreAdminTx) Find() (result *model.Admin, err error) {
	return result, a.tx.Find(&result)
}

func (a scoreAdminTx) Append(values ...*model.Admin) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a scoreAdminTx) Replace(values ...*model.Admin) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a scoreAdminTx) Delete(values ...*model.Admin) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a scoreAdminTx) Clear() error {
	return a.tx.Clear()
}

func (a scoreAdminTx) Count() int64 {
	return a.tx.Count()
}

type scoreApplicant struct {
	db *gorm.DB

	field.RelationField

	Intents struct {
		field.RelationField
		Applicant struct {
			field.RelationField
		}
		OptionalTime struct {
			field.RelationField
		}
	}
	Answers struct {
		field.RelationField
		Applicant struct {
			field.RelationField
		}
		Question struct {
			field.RelationField
		}
	}
}

func (a scoreApplicant) Where(conds ...field.Expr) *scoreApplicant {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a scoreApplicant) WithContext(ctx context.Context) *scoreApplicant {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a scoreApplicant) Model(m *model.Score) *scoreApplicantTx {
	return &scoreApplicantTx{a.db.Model(m).Association(a.Name())}
}

type scoreApplicantTx struct{ tx *gorm.Association }

func (a scoreApplicantTx) Find() (result *model.Applicant, err error) {
	return result, a.tx.Find(&result)
}

func (a scoreApplicantTx) Append(values ...*model.Applicant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a scoreApplicantTx) Replace(values ...*model.Applicant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a scoreApplicantTx) Delete(values ...*model.Applicant) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a scoreApplicantTx) Clear() error {
	return a.tx.Clear()
}

func (a scoreApplicantTx) Count() int64 {
	return a.tx.Count()
}

type scoreStandard struct {
	db *gorm.DB

	field.RelationField
}

func (a scoreStandard) Where(conds ...field.Expr) *scoreStandard {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a scoreStandard) WithContext(ctx context.Context) *scoreStandard {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a scoreStandard) Model(m *model.Score) *scoreStandardTx {
	return &scoreStandardTx{a.db.Model(m).Association(a.Name())}
}

type scoreStandardTx struct{ tx *gorm.Association }

func (a scoreStandardTx) Find() (result *model.Standard, err error) {
	return result, a.tx.Find(&result)
}

func (a scoreStandardTx) Append(values ...*model.Standard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a scoreStandardTx) Replace(values ...*model.Standard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a scoreStandardTx) Delete(values ...*model.Standard) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a scoreStandardTx) Clear() error {
	return a.tx.Clear()
}

func (a scoreStandardTx) Count() int64 {
	return a.tx.Count()
}

type scoreDo struct{ gen.DO }

func (s scoreDo) Debug() *scoreDo {
	return s.withDO(s.DO.Debug())
}

func (s scoreDo) WithContext(ctx context.Context) *scoreDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s scoreDo) Clauses(conds ...clause.Expression) *scoreDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s scoreDo) Returning(value interface{}, columns ...string) *scoreDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s scoreDo) Not(conds ...gen.Condition) *scoreDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s scoreDo) Or(conds ...gen.Condition) *scoreDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s scoreDo) Select(conds ...field.Expr) *scoreDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s scoreDo) Where(conds ...gen.Condition) *scoreDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s scoreDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *scoreDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s scoreDo) Order(conds ...field.Expr) *scoreDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s scoreDo) Distinct(cols ...field.Expr) *scoreDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s scoreDo) Omit(cols ...field.Expr) *scoreDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s scoreDo) Join(table schema.Tabler, on ...field.Expr) *scoreDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s scoreDo) LeftJoin(table schema.Tabler, on ...field.Expr) *scoreDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s scoreDo) RightJoin(table schema.Tabler, on ...field.Expr) *scoreDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s scoreDo) Group(cols ...field.Expr) *scoreDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s scoreDo) Having(conds ...gen.Condition) *scoreDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s scoreDo) Limit(limit int) *scoreDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s scoreDo) Offset(offset int) *scoreDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s scoreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *scoreDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s scoreDo) Unscoped() *scoreDo {
	return s.withDO(s.DO.Unscoped())
}

func (s scoreDo) Create(values ...*model.Score) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s scoreDo) CreateInBatches(values []*model.Score, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s scoreDo) Save(values ...*model.Score) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s scoreDo) First() (*model.Score, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Score), nil
	}
}

func (s scoreDo) Take() (*model.Score, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Score), nil
	}
}

func (s scoreDo) Last() (*model.Score, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Score), nil
	}
}

func (s scoreDo) Find() ([]*model.Score, error) {
	result, err := s.DO.Find()
	return result.([]*model.Score), err
}

func (s scoreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Score, err error) {
	buf := make([]*model.Score, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s scoreDo) FindInBatches(result *[]*model.Score, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s scoreDo) Attrs(attrs ...field.AssignExpr) *scoreDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s scoreDo) Assign(attrs ...field.AssignExpr) *scoreDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s scoreDo) Joins(fields ...field.RelationField) *scoreDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s scoreDo) Preload(fields ...field.RelationField) *scoreDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s scoreDo) FirstOrInit() (*model.Score, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Score), nil
	}
}

func (s scoreDo) FirstOrCreate() (*model.Score, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Score), nil
	}
}

func (s scoreDo) FindByPage(offset int, limit int) (result []*model.Score, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s scoreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s *scoreDo) withDO(do gen.Dao) *scoreDo {
	s.DO = *do.(*gen.DO)
	return s
}
