// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/next/model"
)

func newOrganization(db *gorm.DB, opts ...gen.DOOption) organization {
	_organization := organization{}

	_organization.organizationDo.UseDB(db, opts...)
	_organization.organizationDo.UseModel(&model.Organization{})

	tableName := _organization.organizationDo.TableName()
	_organization.ALL = field.NewAsterisk(tableName)
	_organization.CreatedAt = field.NewTime(tableName, "created_at")
	_organization.ID = field.NewString(tableName, "id")
	_organization.Title = field.NewString(tableName, "title")
	_organization.Slug = field.NewString(tableName, "slug")
	_organization.PublicKey = field.NewString(tableName, "public_key")

	_organization.fillFieldMap()

	return _organization
}

type organization struct {
	organizationDo

	ALL       field.Asterisk
	CreatedAt field.Time
	ID        field.String
	Title     field.String
	Slug      field.String
	PublicKey field.String

	fieldMap map[string]field.Expr
}

func (o organization) Table(newTableName string) *organization {
	o.organizationDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o organization) As(alias string) *organization {
	o.organizationDo.DO = *(o.organizationDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *organization) updateTableName(table string) *organization {
	o.ALL = field.NewAsterisk(table)
	o.CreatedAt = field.NewTime(table, "created_at")
	o.ID = field.NewString(table, "id")
	o.Title = field.NewString(table, "title")
	o.Slug = field.NewString(table, "slug")
	o.PublicKey = field.NewString(table, "public_key")

	o.fillFieldMap()

	return o
}

func (o *organization) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *organization) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 5)
	o.fieldMap["created_at"] = o.CreatedAt
	o.fieldMap["id"] = o.ID
	o.fieldMap["title"] = o.Title
	o.fieldMap["slug"] = o.Slug
	o.fieldMap["public_key"] = o.PublicKey
}

func (o organization) clone(db *gorm.DB) organization {
	o.organizationDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o organization) replaceDB(db *gorm.DB) organization {
	o.organizationDo.ReplaceDB(db)
	return o
}

type organizationDo struct{ gen.DO }

type IOrganizationDo interface {
	gen.SubQuery
	Debug() IOrganizationDo
	WithContext(ctx context.Context) IOrganizationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IOrganizationDo
	WriteDB() IOrganizationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IOrganizationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IOrganizationDo
	Not(conds ...gen.Condition) IOrganizationDo
	Or(conds ...gen.Condition) IOrganizationDo
	Select(conds ...field.Expr) IOrganizationDo
	Where(conds ...gen.Condition) IOrganizationDo
	Order(conds ...field.Expr) IOrganizationDo
	Distinct(cols ...field.Expr) IOrganizationDo
	Omit(cols ...field.Expr) IOrganizationDo
	Join(table schema.Tabler, on ...field.Expr) IOrganizationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IOrganizationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IOrganizationDo
	Group(cols ...field.Expr) IOrganizationDo
	Having(conds ...gen.Condition) IOrganizationDo
	Limit(limit int) IOrganizationDo
	Offset(offset int) IOrganizationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IOrganizationDo
	Unscoped() IOrganizationDo
	Create(values ...*model.Organization) error
	CreateInBatches(values []*model.Organization, batchSize int) error
	Save(values ...*model.Organization) error
	First() (*model.Organization, error)
	Take() (*model.Organization, error)
	Last() (*model.Organization, error)
	Find() ([]*model.Organization, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Organization, err error)
	FindInBatches(result *[]*model.Organization, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Organization) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IOrganizationDo
	Assign(attrs ...field.AssignExpr) IOrganizationDo
	Joins(fields ...field.RelationField) IOrganizationDo
	Preload(fields ...field.RelationField) IOrganizationDo
	FirstOrInit() (*model.Organization, error)
	FirstOrCreate() (*model.Organization, error)
	FindByPage(offset int, limit int) (result []*model.Organization, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IOrganizationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o organizationDo) Debug() IOrganizationDo {
	return o.withDO(o.DO.Debug())
}

func (o organizationDo) WithContext(ctx context.Context) IOrganizationDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o organizationDo) ReadDB() IOrganizationDo {
	return o.Clauses(dbresolver.Read)
}

func (o organizationDo) WriteDB() IOrganizationDo {
	return o.Clauses(dbresolver.Write)
}

func (o organizationDo) Session(config *gorm.Session) IOrganizationDo {
	return o.withDO(o.DO.Session(config))
}

func (o organizationDo) Clauses(conds ...clause.Expression) IOrganizationDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o organizationDo) Returning(value interface{}, columns ...string) IOrganizationDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o organizationDo) Not(conds ...gen.Condition) IOrganizationDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o organizationDo) Or(conds ...gen.Condition) IOrganizationDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o organizationDo) Select(conds ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o organizationDo) Where(conds ...gen.Condition) IOrganizationDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o organizationDo) Order(conds ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o organizationDo) Distinct(cols ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o organizationDo) Omit(cols ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o organizationDo) Join(table schema.Tabler, on ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o organizationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o organizationDo) RightJoin(table schema.Tabler, on ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o organizationDo) Group(cols ...field.Expr) IOrganizationDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o organizationDo) Having(conds ...gen.Condition) IOrganizationDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o organizationDo) Limit(limit int) IOrganizationDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o organizationDo) Offset(offset int) IOrganizationDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o organizationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IOrganizationDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o organizationDo) Unscoped() IOrganizationDo {
	return o.withDO(o.DO.Unscoped())
}

func (o organizationDo) Create(values ...*model.Organization) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o organizationDo) CreateInBatches(values []*model.Organization, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o organizationDo) Save(values ...*model.Organization) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o organizationDo) First() (*model.Organization, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organization), nil
	}
}

func (o organizationDo) Take() (*model.Organization, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organization), nil
	}
}

func (o organizationDo) Last() (*model.Organization, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organization), nil
	}
}

func (o organizationDo) Find() ([]*model.Organization, error) {
	result, err := o.DO.Find()
	return result.([]*model.Organization), err
}

func (o organizationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Organization, err error) {
	buf := make([]*model.Organization, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o organizationDo) FindInBatches(result *[]*model.Organization, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o organizationDo) Attrs(attrs ...field.AssignExpr) IOrganizationDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o organizationDo) Assign(attrs ...field.AssignExpr) IOrganizationDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o organizationDo) Joins(fields ...field.RelationField) IOrganizationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o organizationDo) Preload(fields ...field.RelationField) IOrganizationDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o organizationDo) FirstOrInit() (*model.Organization, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organization), nil
	}
}

func (o organizationDo) FirstOrCreate() (*model.Organization, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Organization), nil
	}
}

func (o organizationDo) FindByPage(offset int, limit int) (result []*model.Organization, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o organizationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o organizationDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o organizationDo) Delete(models ...*model.Organization) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *organizationDo) withDO(do gen.Dao) *organizationDo {
	o.DO = *do.(*gen.DO)
	return o
}