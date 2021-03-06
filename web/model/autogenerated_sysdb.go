// Code generated by go-queryset. DO NOT EDIT.
package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set SysDBQuerySet

// SysDBQuerySet is an queryset type for SysDB
type SysDBQuerySet struct {
	db *gorm.DB
}

// NewSysDBQuerySet constructs new SysDBQuerySet
func NewSysDBQuerySet(db *gorm.DB) SysDBQuerySet {
	return SysDBQuerySet{
		db: db.Model(&SysDB{}),
	}
}

func (qs SysDBQuerySet) w(db *gorm.DB) SysDBQuerySet {
	return NewSysDBQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) All(ret *[]SysDB) error {
	return qs.db.Find(ret).Error
}

// Count is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *SysDB) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) CreatedAtEq(createdAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) CreatedAtGt(createdAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) CreatedAtGte(createdAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) CreatedAtLt(createdAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) CreatedAtLte(createdAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) CreatedAtNe(createdAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// DBNameEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) DBNameEq(dBName string) SysDBQuerySet {
	return qs.w(qs.db.Where("db_name = ?", dBName))
}

// DBNameIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) DBNameIn(dBName string, dBNameRest ...string) SysDBQuerySet {
	iArgs := []interface{}{dBName}
	for _, arg := range dBNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("db_name IN (?)", iArgs))
}

// DBNameNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) DBNameNe(dBName string) SysDBQuerySet {
	return qs.w(qs.db.Where("db_name != ?", dBName))
}

// DBNameNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) DBNameNotIn(dBName string, dBNameRest ...string) SysDBQuerySet {
	iArgs := []interface{}{dBName}
	for _, arg := range dBNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("db_name NOT IN (?)", iArgs))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *SysDB) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) Delete() error {
	return qs.db.Delete(SysDB{}).Error
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) GetUpdater() SysDBUpdater {
	return NewSysDBUpdater(qs.db)
}

// HostEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) HostEq(host string) SysDBQuerySet {
	return qs.w(qs.db.Where("host = ?", host))
}

// HostIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) HostIn(host string, hostRest ...string) SysDBQuerySet {
	iArgs := []interface{}{host}
	for _, arg := range hostRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("host IN (?)", iArgs))
}

// HostNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) HostNe(host string) SysDBQuerySet {
	return qs.w(qs.db.Where("host != ?", host))
}

// HostNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) HostNotIn(host string, hostRest ...string) SysDBQuerySet {
	iArgs := []interface{}{host}
	for _, arg := range hostRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("host NOT IN (?)", iArgs))
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDEq(ID uint64) SysDBQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDGt(ID uint64) SysDBQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDGte(ID uint64) SysDBQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDIn(ID uint64, IDRest ...uint64) SysDBQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDLt(ID uint64) SysDBQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDLte(ID uint64) SysDBQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDNe(ID uint64) SysDBQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) IDNotIn(ID uint64, IDRest ...uint64) SysDBQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) Limit(limit int) SysDBQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs SysDBQuerySet) One(ret *SysDB) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderAscByCreatedAt() SysDBQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderAscByID() SysDBQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByPort is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderAscByPort() SysDBQuerySet {
	return qs.w(qs.db.Order("port ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderAscByUpdatedAt() SysDBQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderDescByCreatedAt() SysDBQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderDescByID() SysDBQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByPort is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderDescByPort() SysDBQuerySet {
	return qs.w(qs.db.Order("port DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) OrderDescByUpdatedAt() SysDBQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// PasswordEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PasswordEq(password string) SysDBQuerySet {
	return qs.w(qs.db.Where("password = ?", password))
}

// PasswordIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PasswordIn(password string, passwordRest ...string) SysDBQuerySet {
	iArgs := []interface{}{password}
	for _, arg := range passwordRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("password IN (?)", iArgs))
}

// PasswordNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PasswordNe(password string) SysDBQuerySet {
	return qs.w(qs.db.Where("password != ?", password))
}

// PasswordNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PasswordNotIn(password string, passwordRest ...string) SysDBQuerySet {
	iArgs := []interface{}{password}
	for _, arg := range passwordRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("password NOT IN (?)", iArgs))
}

// PortEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortEq(port int) SysDBQuerySet {
	return qs.w(qs.db.Where("port = ?", port))
}

// PortGt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortGt(port int) SysDBQuerySet {
	return qs.w(qs.db.Where("port > ?", port))
}

// PortGte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortGte(port int) SysDBQuerySet {
	return qs.w(qs.db.Where("port >= ?", port))
}

// PortIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortIn(port int, portRest ...int) SysDBQuerySet {
	iArgs := []interface{}{port}
	for _, arg := range portRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("port IN (?)", iArgs))
}

// PortLt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortLt(port int) SysDBQuerySet {
	return qs.w(qs.db.Where("port < ?", port))
}

// PortLte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortLte(port int) SysDBQuerySet {
	return qs.w(qs.db.Where("port <= ?", port))
}

// PortNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortNe(port int) SysDBQuerySet {
	return qs.w(qs.db.Where("port != ?", port))
}

// PortNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) PortNotIn(port int, portRest ...int) SysDBQuerySet {
	iArgs := []interface{}{port}
	for _, arg := range portRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("port NOT IN (?)", iArgs))
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetCreatedAt(createdAt time.Time) SysDBUpdater {
	u.fields[string(SysDBDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDBName is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetDBName(dBName string) SysDBUpdater {
	u.fields[string(SysDBDBSchema.DBName)] = dBName
	return u
}

// SetHost is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetHost(host string) SysDBUpdater {
	u.fields[string(SysDBDBSchema.Host)] = host
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetID(ID uint64) SysDBUpdater {
	u.fields[string(SysDBDBSchema.ID)] = ID
	return u
}

// SetPassword is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetPassword(password string) SysDBUpdater {
	u.fields[string(SysDBDBSchema.Password)] = password
	return u
}

// SetPort is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetPort(port int) SysDBUpdater {
	u.fields[string(SysDBDBSchema.Port)] = port
	return u
}

// SetShowName is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetShowName(showName string) SysDBUpdater {
	u.fields[string(SysDBDBSchema.ShowName)] = showName
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetUpdatedAt(updatedAt time.Time) SysDBUpdater {
	u.fields[string(SysDBDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetUser is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) SetUser(user string) SysDBUpdater {
	u.fields[string(SysDBDBSchema.User)] = user
	return u
}

// ShowNameEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) ShowNameEq(showName string) SysDBQuerySet {
	return qs.w(qs.db.Where("show_name = ?", showName))
}

// ShowNameIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) ShowNameIn(showName string, showNameRest ...string) SysDBQuerySet {
	iArgs := []interface{}{showName}
	for _, arg := range showNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("show_name IN (?)", iArgs))
}

// ShowNameNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) ShowNameNe(showName string) SysDBQuerySet {
	return qs.w(qs.db.Where("show_name != ?", showName))
}

// ShowNameNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) ShowNameNotIn(showName string, showNameRest ...string) SysDBQuerySet {
	iArgs := []interface{}{showName}
	for _, arg := range showNameRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("show_name NOT IN (?)", iArgs))
}

// Update is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u SysDBUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UpdatedAtEq(updatedAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UpdatedAtGt(updatedAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UpdatedAtGte(updatedAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UpdatedAtLt(updatedAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UpdatedAtLte(updatedAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UpdatedAtNe(updatedAt time.Time) SysDBQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// UserEq is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UserEq(user string) SysDBQuerySet {
	return qs.w(qs.db.Where("user = ?", user))
}

// UserIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UserIn(user string, userRest ...string) SysDBQuerySet {
	iArgs := []interface{}{user}
	for _, arg := range userRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("user IN (?)", iArgs))
}

// UserNe is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UserNe(user string) SysDBQuerySet {
	return qs.w(qs.db.Where("user != ?", user))
}

// UserNotIn is an autogenerated method
// nolint: dupl
func (qs SysDBQuerySet) UserNotIn(user string, userRest ...string) SysDBQuerySet {
	iArgs := []interface{}{user}
	for _, arg := range userRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("user NOT IN (?)", iArgs))
}

// ===== END of query set SysDBQuerySet

// ===== BEGIN of SysDB modifiers

// SysDBDBSchemaField describes database schema field. It requires for method 'Update'
type SysDBDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f SysDBDBSchemaField) String() string {
	return string(f)
}

// SysDBDBSchema stores db field names of SysDB
var SysDBDBSchema = struct {
	ID        SysDBDBSchemaField
	ShowName  SysDBDBSchemaField
	Host      SysDBDBSchemaField
	Port      SysDBDBSchemaField
	User      SysDBDBSchemaField
	Password  SysDBDBSchemaField
	DBName    SysDBDBSchemaField
	CreatedAt SysDBDBSchemaField
	UpdatedAt SysDBDBSchemaField
}{

	ID:        SysDBDBSchemaField("id"),
	ShowName:  SysDBDBSchemaField("show_name"),
	Host:      SysDBDBSchemaField("host"),
	Port:      SysDBDBSchemaField("port"),
	User:      SysDBDBSchemaField("user"),
	Password:  SysDBDBSchemaField("password"),
	DBName:    SysDBDBSchemaField("db_name"),
	CreatedAt: SysDBDBSchemaField("created_at"),
	UpdatedAt: SysDBDBSchemaField("updated_at"),
}

// Update updates SysDB fields by primary key
// nolint: dupl
func (o *SysDB) Update(db *gorm.DB, fields ...SysDBDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":         o.ID,
		"show_name":  o.ShowName,
		"host":       o.Host,
		"port":       o.Port,
		"user":       o.User,
		"password":   o.Password,
		"db_name":    o.DBName,
		"created_at": o.CreatedAt,
		"updated_at": o.UpdatedAt,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update SysDB %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// SysDBUpdater is an SysDB updates manager
type SysDBUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewSysDBUpdater creates new SysDB updater
// nolint: dupl
func NewSysDBUpdater(db *gorm.DB) SysDBUpdater {
	return SysDBUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&SysDB{}),
	}
}

// ===== END of SysDB modifiers

// ===== END of all query sets
