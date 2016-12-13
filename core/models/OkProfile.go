package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// OkProfile is an object representing the database table.
type OkProfile struct {
	IDProfile     int       `boil:"id_profile" json:"id_profile" toml:"id_profile" yaml:"id_profile"`
	AkunID        int       `boil:"akun_id" json:"akun_id" toml:"akun_id" yaml:"akun_id"`
	FirstName     string    `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName      string    `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Photo         string    `boil:"photo" json:"photo" toml:"photo" yaml:"photo"`
	CurrentStatus string    `boil:"current_status" json:"current_status" toml:"current_status" yaml:"current_status"`
	CreatedAt     null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *OkProfileR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L OkProfileL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// OkProfileR is where relationships are stored.
type OkProfileR struct {
	Akun                     *OkAkun
	AssignedtoOkProjectTasks OkProjectTaskSlice
	AssignerOkProjectTasks   OkProjectTaskSlice
}

// OkProfileL is where Load methods for each relationship are stored.
type OkProfileL struct{}

var (
	OkProfileColumns               = []string{"id_profile", "akun_id", "first_name", "last_name", "photo", "current_status", "created_at", "updated_at"}
	OkProfileColumnsWithoutDefault = []string{"akun_id", "first_name", "last_name", "photo", "current_status", "created_at", "updated_at"}
	OkProfileColumnsWithDefault    = []string{"id_profile"}
	OkProfilePrimaryKeyColumns     = []string{"id_profile"}
)

type (
	// OkProfileSlice is an alias for a slice of pointers to OkProfile.
	// This should generally be used opposed to []OkProfile.
	OkProfileSlice []*OkProfile
	// OkProfileHook is the signature for custom OkProfile hook methods
	OkProfileHook func(boil.Executor, *OkProfile) error

	OkProfileQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	OkProfileType                 = reflect.TypeOf(&OkProfile{})
	OkProfileMapping              = queries.MakeStructMapping(OkProfileType)
	OkProfilePrimaryKeyMapping, _ = queries.BindMapping(OkProfileType, OkProfileMapping, OkProfilePrimaryKeyColumns)
	OkProfileInsertCacheMut       sync.RWMutex
	OkProfileInsertCache          = make(map[string]insertCache)
	OkProfileUpdateCacheMut       sync.RWMutex
	OkProfileUpdateCache          = make(map[string]updateCache)
	OkProfileUpsertCacheMut       sync.RWMutex
	OkProfileUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var OkProfileBeforeInsertHooks []OkProfileHook
var OkProfileBeforeUpdateHooks []OkProfileHook
var OkProfileBeforeDeleteHooks []OkProfileHook
var OkProfileBeforeUpsertHooks []OkProfileHook

var OkProfileAfterInsertHooks []OkProfileHook
var OkProfileAfterSelectHooks []OkProfileHook
var OkProfileAfterUpdateHooks []OkProfileHook
var OkProfileAfterDeleteHooks []OkProfileHook
var OkProfileAfterUpsertHooks []OkProfileHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OkProfile) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OkProfile) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OkProfile) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OkProfile) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OkProfile) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OkProfile) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OkProfile) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OkProfile) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OkProfile) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProfileAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOkProfileHook registers your hook function for all future operations.
func AddOkProfileHook(hookPoint boil.HookPoint, OkProfileHook OkProfileHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		OkProfileBeforeInsertHooks = append(OkProfileBeforeInsertHooks, OkProfileHook)
	case boil.BeforeUpdateHook:
		OkProfileBeforeUpdateHooks = append(OkProfileBeforeUpdateHooks, OkProfileHook)
	case boil.BeforeDeleteHook:
		OkProfileBeforeDeleteHooks = append(OkProfileBeforeDeleteHooks, OkProfileHook)
	case boil.BeforeUpsertHook:
		OkProfileBeforeUpsertHooks = append(OkProfileBeforeUpsertHooks, OkProfileHook)
	case boil.AfterInsertHook:
		OkProfileAfterInsertHooks = append(OkProfileAfterInsertHooks, OkProfileHook)
	case boil.AfterSelectHook:
		OkProfileAfterSelectHooks = append(OkProfileAfterSelectHooks, OkProfileHook)
	case boil.AfterUpdateHook:
		OkProfileAfterUpdateHooks = append(OkProfileAfterUpdateHooks, OkProfileHook)
	case boil.AfterDeleteHook:
		OkProfileAfterDeleteHooks = append(OkProfileAfterDeleteHooks, OkProfileHook)
	case boil.AfterUpsertHook:
		OkProfileAfterUpsertHooks = append(OkProfileAfterUpsertHooks, OkProfileHook)
	}
}

// OneP returns a single OkProfile record from the query, and panics on error.
func (q OkProfileQuery) OneP() *OkProfile {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single OkProfile record from the query.
func (q OkProfileQuery) One() (*OkProfile, error) {
	o := &OkProfile{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OkProfile")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all OkProfile records from the query, and panics on error.
func (q OkProfileQuery) AllP() OkProfileSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OkProfile records from the query.
func (q OkProfileQuery) All() (OkProfileSlice, error) {
	var o OkProfileSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OkProfile slice")
	}

	if len(OkProfileAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all OkProfile records in the query, and panics on error.
func (q OkProfileQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OkProfile records in the query.
func (q OkProfileQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OkProfile rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q OkProfileQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q OkProfileQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OkProfile exists")
	}

	return count > 0, nil
}

// AkunG pointed to by the foreign key.
func (o *OkProfile) AkunG(mods ...qm.QueryMod) OkAkunQuery {
	return o.Akun(boil.GetDB(), mods...)
}

// Akun pointed to by the foreign key.
func (o *OkProfile) Akun(exec boil.Executor, mods ...qm.QueryMod) OkAkunQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id_akun=?", o.AkunID),
	}

	queryMods = append(queryMods, mods...)

	query := OkAkuns(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkAkun`")

	return query
}

// AssignedtoOkProjectTasksG retrieves all the OkProjectTask's OkProjectTask via assignedto_id column.
func (o *OkProfile) AssignedtoOkProjectTasksG(mods ...qm.QueryMod) OkProjectTaskQuery {
	return o.AssignedtoOkProjectTasks(boil.GetDB(), mods...)
}

// AssignedtoOkProjectTasks retrieves all the OkProjectTask's OkProjectTask with an executor via assignedto_id column.
func (o *OkProfile) AssignedtoOkProjectTasks(exec boil.Executor, mods ...qm.QueryMod) OkProjectTaskQuery {
	queryMods := []qm.QueryMod{
		qm.Select("`a`.*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`a`.`assignedto_id`=?", o.IDProfile),
	)

	query := OkProjectTasks(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProjectTask` as `a`")
	return query
}

// AssignerOkProjectTasksG retrieves all the OkProjectTask's OkProjectTask via assigner_id column.
func (o *OkProfile) AssignerOkProjectTasksG(mods ...qm.QueryMod) OkProjectTaskQuery {
	return o.AssignerOkProjectTasks(boil.GetDB(), mods...)
}

// AssignerOkProjectTasks retrieves all the OkProjectTask's OkProjectTask with an executor via assigner_id column.
func (o *OkProfile) AssignerOkProjectTasks(exec boil.Executor, mods ...qm.QueryMod) OkProjectTaskQuery {
	queryMods := []qm.QueryMod{
		qm.Select("`a`.*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`a`.`assigner_id`=?", o.IDProfile),
	)

	query := OkProjectTasks(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProjectTask` as `a`")
	return query
}

// LoadAkun allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProfileL) LoadAkun(e boil.Executor, singular bool, maybeOkProfile interface{}) error {
	var slice []*OkProfile
	var object *OkProfile

	count := 1
	if singular {
		object = maybeOkProfile.(*OkProfile)
	} else {
		slice = *maybeOkProfile.(*OkProfileSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProfileR{}
		}
		args[0] = object.AkunID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProfileR{}
			}
			args[i] = obj.AkunID
		}
	}

	query := fmt.Sprintf(
		"select * from `OkAkun` where `id_akun` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkAkun")
	}
	defer results.Close()

	var resultSlice []*OkAkun
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkAkun")
	}

	if len(OkProfileAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Akun = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AkunID == foreign.IDAkun {
				local.R.Akun = foreign
				break
			}
		}
	}

	return nil
}

// LoadAssignedtoOkProjectTasks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProfileL) LoadAssignedtoOkProjectTasks(e boil.Executor, singular bool, maybeOkProfile interface{}) error {
	var slice []*OkProfile
	var object *OkProfile

	count := 1
	if singular {
		object = maybeOkProfile.(*OkProfile)
	} else {
		slice = *maybeOkProfile.(*OkProfileSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProfileR{}
		}
		args[0] = object.IDProfile
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProfileR{}
			}
			args[i] = obj.IDProfile
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProjectTask` where `assignedto_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkProjectTask")
	}
	defer results.Close()

	var resultSlice []*OkProjectTask
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkProjectTask")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AssignedtoOkProjectTasks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IDProfile == foreign.AssignedtoID {
				local.R.AssignedtoOkProjectTasks = append(local.R.AssignedtoOkProjectTasks, foreign)
				break
			}
		}
	}

	return nil
}

// LoadAssignerOkProjectTasks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProfileL) LoadAssignerOkProjectTasks(e boil.Executor, singular bool, maybeOkProfile interface{}) error {
	var slice []*OkProfile
	var object *OkProfile

	count := 1
	if singular {
		object = maybeOkProfile.(*OkProfile)
	} else {
		slice = *maybeOkProfile.(*OkProfileSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProfileR{}
		}
		args[0] = object.IDProfile
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProfileR{}
			}
			args[i] = obj.IDProfile
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProjectTask` where `assigner_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkProjectTask")
	}
	defer results.Close()

	var resultSlice []*OkProjectTask
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkProjectTask")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AssignerOkProjectTasks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IDProfile == foreign.AssignerID {
				local.R.AssignerOkProjectTasks = append(local.R.AssignerOkProjectTasks, foreign)
				break
			}
		}
	}

	return nil
}

// SetAkun of the OkProfile to the related item.
// Sets o.R.Akun to related.
// Adds o to related.R.AkunOkProfiles.
func (o *OkProfile) SetAkun(exec boil.Executor, insert bool, related *OkAkun) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `OkProfile` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"akun_id"}),
		strmangle.WhereClause("`", "`", 0, OkProfilePrimaryKeyColumns),
	)
	values := []interface{}{related.IDAkun, o.IDProfile}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AkunID = related.IDAkun

	if o.R == nil {
		o.R = &OkProfileR{
			Akun: related,
		}
	} else {
		o.R.Akun = related
	}

	if related.R == nil {
		related.R = &OkAkunR{
			AkunOkProfiles: OkProfileSlice{o},
		}
	} else {
		related.R.AkunOkProfiles = append(related.R.AkunOkProfiles, o)
	}

	return nil
}

// AddAssignedtoOkProjectTasks adds the given related objects to the existing relationships
// of the OkProfile, optionally inserting them as new records.
// Appends related to o.R.AssignedtoOkProjectTasks.
// Sets related.R.Assignedto appropriately.
func (o *OkProfile) AddAssignedtoOkProjectTasks(exec boil.Executor, insert bool, related ...*OkProjectTask) error {
	var err error
	for _, rel := range related {
		rel.AssignedtoID = o.IDProfile
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "assignedto_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &OkProfileR{
			AssignedtoOkProjectTasks: related,
		}
	} else {
		o.R.AssignedtoOkProjectTasks = append(o.R.AssignedtoOkProjectTasks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &OkProjectTaskR{
				Assignedto: o,
			}
		} else {
			rel.R.Assignedto = o
		}
	}
	return nil
}

// AddAssignerOkProjectTasks adds the given related objects to the existing relationships
// of the OkProfile, optionally inserting them as new records.
// Appends related to o.R.AssignerOkProjectTasks.
// Sets related.R.Assigner appropriately.
func (o *OkProfile) AddAssignerOkProjectTasks(exec boil.Executor, insert bool, related ...*OkProjectTask) error {
	var err error
	for _, rel := range related {
		rel.AssignerID = o.IDProfile
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "assigner_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &OkProfileR{
			AssignerOkProjectTasks: related,
		}
	} else {
		o.R.AssignerOkProjectTasks = append(o.R.AssignerOkProjectTasks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &OkProjectTaskR{
				Assigner: o,
			}
		} else {
			rel.R.Assigner = o
		}
	}
	return nil
}

// OkProfilesG retrieves all records.
func OkProfilesG(mods ...qm.QueryMod) OkProfileQuery {
	return OkProfiles(boil.GetDB(), mods...)
}

// OkProfiles retrieves all the records using an executor.
func OkProfiles(exec boil.Executor, mods ...qm.QueryMod) OkProfileQuery {
	mods = append(mods, qm.From("`OkProfile`"))
	return OkProfileQuery{NewQuery(exec, mods...)}
}

// FindOkProfileG retrieves a single record by ID.
func FindOkProfileG(idProfile int, selectCols ...string) (*OkProfile, error) {
	return FindOkProfile(boil.GetDB(), idProfile, selectCols...)
}

// FindOkProfileGP retrieves a single record by ID, and panics on error.
func FindOkProfileGP(idProfile int, selectCols ...string) *OkProfile {
	retobj, err := FindOkProfile(boil.GetDB(), idProfile, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOkProfile retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOkProfile(exec boil.Executor, idProfile int, selectCols ...string) (*OkProfile, error) {
	OkProfileObj := &OkProfile{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OkProfile` where `id_profile`=?", sel,
	)

	q := queries.Raw(exec, query, idProfile)

	err := q.Bind(OkProfileObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OkProfile")
	}

	return OkProfileObj, nil
}

// FindOkProfileP retrieves a single record by ID with an executor, and panics on error.
func FindOkProfileP(exec boil.Executor, idProfile int, selectCols ...string) *OkProfile {
	retobj, err := FindOkProfile(exec, idProfile, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OkProfile) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OkProfile) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OkProfile) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OkProfile) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkProfile provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(OkProfileColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	OkProfileInsertCacheMut.RLock()
	cache, cached := OkProfileInsertCache[key]
	OkProfileInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			OkProfileColumns,
			OkProfileColumnsWithDefault,
			OkProfileColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(OkProfileType, OkProfileMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(OkProfileType, OkProfileMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `OkProfile` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OkProfile` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, OkProfilePrimaryKeyColumns))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into OkProfile")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.IDProfile = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkProfileMapping["IDProfile"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDProfile,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkProfile")
	}

CacheNoHooks:
	if !cached {
		OkProfileInsertCacheMut.Lock()
		OkProfileInsertCache[key] = cache
		OkProfileInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single OkProfile record. See Update for
// whitelist behavior description.
func (o *OkProfile) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OkProfile record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OkProfile) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OkProfile, and panics on error.
// See Update for whitelist behavior description.
func (o *OkProfile) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OkProfile.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OkProfile) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	OkProfileUpdateCacheMut.RLock()
	cache, cached := OkProfileUpdateCache[key]
	OkProfileUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(OkProfileColumns, OkProfilePrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update OkProfile, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OkProfile` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, OkProfilePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(OkProfileType, OkProfileMapping, append(wl, OkProfilePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update OkProfile row")
	}

	if !cached {
		OkProfileUpdateCacheMut.Lock()
		OkProfileUpdateCache[key] = cache
		OkProfileUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q OkProfileQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q OkProfileQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for OkProfile")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OkProfileSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OkProfileSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OkProfileSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OkProfileSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `OkProfile` SET %s WHERE (`id_profile`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkProfilePrimaryKeyColumns), len(colNames)+1, len(OkProfilePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in OkProfile slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OkProfile) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OkProfile) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OkProfile) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OkProfile) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkProfile provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(OkProfileColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	OkProfileUpsertCacheMut.RLock()
	cache, cached := OkProfileUpsertCache[key]
	OkProfileUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			OkProfileColumns,
			OkProfileColumnsWithDefault,
			OkProfileColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			OkProfileColumns,
			OkProfilePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert OkProfile, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "OkProfile", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OkProfile` WHERE `id_profile`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(OkProfileType, OkProfileMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(OkProfileType, OkProfileMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for OkProfile")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.IDProfile = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkProfileMapping["IDProfile"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDProfile,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkProfile")
	}

CacheNoHooks:
	if !cached {
		OkProfileUpsertCacheMut.Lock()
		OkProfileUpsertCache[key] = cache
		OkProfileUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single OkProfile record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkProfile) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OkProfile record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OkProfile) DeleteG() error {
	if o == nil {
		return errors.New("models: no OkProfile provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OkProfile record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkProfile) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OkProfile record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OkProfile) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkProfile provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), OkProfilePrimaryKeyMapping)
	sql := "DELETE FROM `OkProfile` WHERE `id_profile`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from OkProfile")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q OkProfileQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q OkProfileQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no OkProfileQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkProfile")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OkProfileSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OkProfileSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no OkProfile slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OkProfileSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OkProfileSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkProfile slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(OkProfileBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `OkProfile` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkProfilePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkProfilePrimaryKeyColumns), 1, len(OkProfilePrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkProfile slice")
	}

	if len(OkProfileAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OkProfile) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OkProfile) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OkProfile) ReloadG() error {
	if o == nil {
		return errors.New("models: no OkProfile provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OkProfile) Reload(exec boil.Executor) error {
	ret, err := FindOkProfile(exec, o.IDProfile)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkProfileSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkProfileSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkProfileSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OkProfileSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkProfileSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	OkProfiles := OkProfileSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProfilePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `OkProfile`.* FROM `OkProfile` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkProfilePrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(OkProfilePrimaryKeyColumns), 1, len(OkProfilePrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&OkProfiles)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OkProfileSlice")
	}

	*o = OkProfiles

	return nil
}

// OkProfileExists checks if the OkProfile row exists.
func OkProfileExists(exec boil.Executor, idProfile int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `OkProfile` where `id_profile`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, idProfile)
	}

	row := exec.QueryRow(sql, idProfile)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OkProfile exists")
	}

	return exists, nil
}

// OkProfileExistsG checks if the OkProfile row exists.
func OkProfileExistsG(idProfile int) (bool, error) {
	return OkProfileExists(boil.GetDB(), idProfile)
}

// OkProfileExistsGP checks if the OkProfile row exists. Panics on error.
func OkProfileExistsGP(idProfile int) bool {
	e, err := OkProfileExists(boil.GetDB(), idProfile)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OkProfileExistsP checks if the OkProfile row exists. Panics on error.
func OkProfileExistsP(exec boil.Executor, idProfile int) bool {
	e, err := OkProfileExists(exec, idProfile)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
