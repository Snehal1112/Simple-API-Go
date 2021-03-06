// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

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

// Okproject is an object representing the database table.
type Okproject struct {
	IDProject   int         `boil:"id_project" json:"id_project" toml:"id_project" yaml:"id_project"`
	NamaProject string      `boil:"nama_project" json:"nama_project" toml:"nama_project" yaml:"nama_project"`
	Desc        null.String `boil:"desc" json:"desc,omitempty" toml:"desc" yaml:"desc,omitempty"`
	GitURL      string      `boil:"git_url" json:"git_url" toml:"git_url" yaml:"git_url"`

	R *okprojectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L okprojectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// okprojectR is where relationships are stored.
type okprojectR struct {
	ProjectOkprojecttasks OkprojecttaskSlice
}

// okprojectL is where Load methods for each relationship are stored.
type okprojectL struct{}

var (
	okprojectColumns               = []string{"id_project", "nama_project", "desc", "git_url"}
	okprojectColumnsWithoutDefault = []string{"nama_project", "desc", "git_url"}
	okprojectColumnsWithDefault    = []string{"id_project"}
	okprojectPrimaryKeyColumns     = []string{"id_project"}
)

type (
	// OkprojectSlice is an alias for a slice of pointers to Okproject.
	// This should generally be used opposed to []Okproject.
	OkprojectSlice []*Okproject
	// OkprojectHook is the signature for custom Okproject hook methods
	OkprojectHook func(boil.Executor, *Okproject) error

	okprojectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	okprojectType                 = reflect.TypeOf(&Okproject{})
	okprojectMapping              = queries.MakeStructMapping(okprojectType)
	okprojectPrimaryKeyMapping, _ = queries.BindMapping(okprojectType, okprojectMapping, okprojectPrimaryKeyColumns)
	okprojectInsertCacheMut       sync.RWMutex
	okprojectInsertCache          = make(map[string]insertCache)
	okprojectUpdateCacheMut       sync.RWMutex
	okprojectUpdateCache          = make(map[string]updateCache)
	okprojectUpsertCacheMut       sync.RWMutex
	okprojectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var okprojectBeforeInsertHooks []OkprojectHook
var okprojectBeforeUpdateHooks []OkprojectHook
var okprojectBeforeDeleteHooks []OkprojectHook
var okprojectBeforeUpsertHooks []OkprojectHook

var okprojectAfterInsertHooks []OkprojectHook
var okprojectAfterSelectHooks []OkprojectHook
var okprojectAfterUpdateHooks []OkprojectHook
var okprojectAfterDeleteHooks []OkprojectHook
var okprojectAfterUpsertHooks []OkprojectHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Okproject) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Okproject) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Okproject) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Okproject) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Okproject) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Okproject) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Okproject) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Okproject) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Okproject) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range okprojectAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOkprojectHook registers your hook function for all future operations.
func AddOkprojectHook(hookPoint boil.HookPoint, okprojectHook OkprojectHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		okprojectBeforeInsertHooks = append(okprojectBeforeInsertHooks, okprojectHook)
	case boil.BeforeUpdateHook:
		okprojectBeforeUpdateHooks = append(okprojectBeforeUpdateHooks, okprojectHook)
	case boil.BeforeDeleteHook:
		okprojectBeforeDeleteHooks = append(okprojectBeforeDeleteHooks, okprojectHook)
	case boil.BeforeUpsertHook:
		okprojectBeforeUpsertHooks = append(okprojectBeforeUpsertHooks, okprojectHook)
	case boil.AfterInsertHook:
		okprojectAfterInsertHooks = append(okprojectAfterInsertHooks, okprojectHook)
	case boil.AfterSelectHook:
		okprojectAfterSelectHooks = append(okprojectAfterSelectHooks, okprojectHook)
	case boil.AfterUpdateHook:
		okprojectAfterUpdateHooks = append(okprojectAfterUpdateHooks, okprojectHook)
	case boil.AfterDeleteHook:
		okprojectAfterDeleteHooks = append(okprojectAfterDeleteHooks, okprojectHook)
	case boil.AfterUpsertHook:
		okprojectAfterUpsertHooks = append(okprojectAfterUpsertHooks, okprojectHook)
	}
}

// OneP returns a single okproject record from the query, and panics on error.
func (q okprojectQuery) OneP() *Okproject {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single okproject record from the query.
func (q okprojectQuery) One() (*Okproject, error) {
	o := &Okproject{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for okproject")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Okproject records from the query, and panics on error.
func (q okprojectQuery) AllP() OkprojectSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Okproject records from the query.
func (q okprojectQuery) All() (OkprojectSlice, error) {
	var o OkprojectSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Okproject slice")
	}

	if len(okprojectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Okproject records in the query, and panics on error.
func (q okprojectQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Okproject records in the query.
func (q okprojectQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count okproject rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q okprojectQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q okprojectQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if okproject exists")
	}

	return count > 0, nil
}

// ProjectOkprojecttasksG retrieves all the okprojecttask's okprojecttask via project_id column.
func (o *Okproject) ProjectOkprojecttasksG(mods ...qm.QueryMod) okprojecttaskQuery {
	return o.ProjectOkprojecttasks(boil.GetDB(), mods...)
}

// ProjectOkprojecttasks retrieves all the okprojecttask's okprojecttask with an executor via project_id column.
func (o *Okproject) ProjectOkprojecttasks(exec boil.Executor, mods ...qm.QueryMod) okprojecttaskQuery {
	queryMods := []qm.QueryMod{
		qm.Select("`a`.*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`a`.`project_id`=?", o.IDProject),
	)

	query := Okprojecttasks(exec, queryMods...)
	queries.SetFrom(query.Query, "`okprojecttask` as `a`")
	return query
}

// LoadProjectOkprojecttasks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (okprojectL) LoadProjectOkprojecttasks(e boil.Executor, singular bool, maybeOkproject interface{}) error {
	var slice []*Okproject
	var object *Okproject

	count := 1
	if singular {
		object = maybeOkproject.(*Okproject)
	} else {
		slice = *maybeOkproject.(*OkprojectSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &okprojectR{}
		}
		args[0] = object.IDProject
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &okprojectR{}
			}
			args[i] = obj.IDProject
		}
	}

	query := fmt.Sprintf(
		"select * from `okprojecttask` where `project_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load okprojecttask")
	}
	defer results.Close()

	var resultSlice []*Okprojecttask
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice okprojecttask")
	}

	if len(okprojecttaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ProjectOkprojecttasks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IDProject == foreign.ProjectID {
				local.R.ProjectOkprojecttasks = append(local.R.ProjectOkprojecttasks, foreign)
				break
			}
		}
	}

	return nil
}

// AddProjectOkprojecttasksG adds the given related objects to the existing relationships
// of the okproject, optionally inserting them as new records.
// Appends related to o.R.ProjectOkprojecttasks.
// Sets related.R.Project appropriately.
// Uses the global database handle.
func (o *Okproject) AddProjectOkprojecttasksG(insert bool, related ...*Okprojecttask) error {
	return o.AddProjectOkprojecttasks(boil.GetDB(), insert, related...)
}

// AddProjectOkprojecttasksP adds the given related objects to the existing relationships
// of the okproject, optionally inserting them as new records.
// Appends related to o.R.ProjectOkprojecttasks.
// Sets related.R.Project appropriately.
// Panics on error.
func (o *Okproject) AddProjectOkprojecttasksP(exec boil.Executor, insert bool, related ...*Okprojecttask) {
	if err := o.AddProjectOkprojecttasks(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddProjectOkprojecttasksGP adds the given related objects to the existing relationships
// of the okproject, optionally inserting them as new records.
// Appends related to o.R.ProjectOkprojecttasks.
// Sets related.R.Project appropriately.
// Uses the global database handle and panics on error.
func (o *Okproject) AddProjectOkprojecttasksGP(insert bool, related ...*Okprojecttask) {
	if err := o.AddProjectOkprojecttasks(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddProjectOkprojecttasks adds the given related objects to the existing relationships
// of the okproject, optionally inserting them as new records.
// Appends related to o.R.ProjectOkprojecttasks.
// Sets related.R.Project appropriately.
func (o *Okproject) AddProjectOkprojecttasks(exec boil.Executor, insert bool, related ...*Okprojecttask) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProjectID = o.IDProject
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `okprojecttask` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"project_id"}),
				strmangle.WhereClause("`", "`", 0, okprojecttaskPrimaryKeyColumns),
			)
			values := []interface{}{o.IDProject, rel.IDPTask}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProjectID = o.IDProject
		}
	}

	if o.R == nil {
		o.R = &okprojectR{
			ProjectOkprojecttasks: related,
		}
	} else {
		o.R.ProjectOkprojecttasks = append(o.R.ProjectOkprojecttasks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &okprojecttaskR{
				Project: o,
			}
		} else {
			rel.R.Project = o
		}
	}
	return nil
}

// OkprojectsG retrieves all records.
func OkprojectsG(mods ...qm.QueryMod) okprojectQuery {
	return Okprojects(boil.GetDB(), mods...)
}

// Okprojects retrieves all the records using an executor.
func Okprojects(exec boil.Executor, mods ...qm.QueryMod) okprojectQuery {
	mods = append(mods, qm.From("`okproject`"))
	return okprojectQuery{NewQuery(exec, mods...)}
}

// FindOkprojectG retrieves a single record by ID.
func FindOkprojectG(idProject int, selectCols ...string) (*Okproject, error) {
	return FindOkproject(boil.GetDB(), idProject, selectCols...)
}

// FindOkprojectGP retrieves a single record by ID, and panics on error.
func FindOkprojectGP(idProject int, selectCols ...string) *Okproject {
	retobj, err := FindOkproject(boil.GetDB(), idProject, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOkproject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOkproject(exec boil.Executor, idProject int, selectCols ...string) (*Okproject, error) {
	okprojectObj := &Okproject{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `okproject` where `id_project`=?", sel,
	)

	q := queries.Raw(exec, query, idProject)

	err := q.Bind(okprojectObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from okproject")
	}

	return okprojectObj, nil
}

// FindOkprojectP retrieves a single record by ID with an executor, and panics on error.
func FindOkprojectP(exec boil.Executor, idProject int, selectCols ...string) *Okproject {
	retobj, err := FindOkproject(exec, idProject, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Okproject) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Okproject) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Okproject) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Okproject) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no okproject provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(okprojectColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	okprojectInsertCacheMut.RLock()
	cache, cached := okprojectInsertCache[key]
	okprojectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			okprojectColumns,
			okprojectColumnsWithDefault,
			okprojectColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(okprojectType, okprojectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(okprojectType, okprojectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `okproject` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `okproject` () VALUES ()"
		}

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `okproject` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, okprojectPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into okproject")
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

	o.IDProject = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == okprojectMapping["IDProject"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDProject,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for okproject")
	}

CacheNoHooks:
	if !cached {
		okprojectInsertCacheMut.Lock()
		okprojectInsertCache[key] = cache
		okprojectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Okproject record. See Update for
// whitelist behavior description.
func (o *Okproject) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Okproject record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Okproject) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Okproject, and panics on error.
// See Update for whitelist behavior description.
func (o *Okproject) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Okproject.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Okproject) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	okprojectUpdateCacheMut.RLock()
	cache, cached := okprojectUpdateCache[key]
	okprojectUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(okprojectColumns, okprojectPrimaryKeyColumns, whitelist)
		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update okproject, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `okproject` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, okprojectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(okprojectType, okprojectMapping, append(wl, okprojectPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update okproject row")
	}

	if !cached {
		okprojectUpdateCacheMut.Lock()
		okprojectUpdateCache[key] = cache
		okprojectUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q okprojectQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q okprojectQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for okproject")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OkprojectSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OkprojectSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OkprojectSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OkprojectSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), okprojectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `okproject` SET %s WHERE (`id_project`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(okprojectPrimaryKeyColumns), len(colNames)+1, len(okprojectPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in okproject slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Okproject) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Okproject) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Okproject) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Okproject) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no okproject provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(okprojectColumnsWithDefault, o)

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

	okprojectUpsertCacheMut.RLock()
	cache, cached := okprojectUpsertCache[key]
	okprojectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			okprojectColumns,
			okprojectColumnsWithDefault,
			okprojectColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			okprojectColumns,
			okprojectPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert okproject, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "okproject", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `okproject` WHERE `id_project`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(okprojectType, okprojectMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(okprojectType, okprojectMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for okproject")
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

	o.IDProject = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == okprojectMapping["IDProject"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDProject,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for okproject")
	}

CacheNoHooks:
	if !cached {
		okprojectUpsertCacheMut.Lock()
		okprojectUpsertCache[key] = cache
		okprojectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Okproject record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Okproject) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Okproject record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Okproject) DeleteG() error {
	if o == nil {
		return errors.New("models: no Okproject provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Okproject record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Okproject) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Okproject record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Okproject) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Okproject provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), okprojectPrimaryKeyMapping)
	sql := "DELETE FROM `okproject` WHERE `id_project`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from okproject")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q okprojectQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q okprojectQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no okprojectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from okproject")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OkprojectSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OkprojectSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Okproject slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OkprojectSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OkprojectSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Okproject slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(okprojectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), okprojectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `okproject` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, okprojectPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(okprojectPrimaryKeyColumns), 1, len(okprojectPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from okproject slice")
	}

	if len(okprojectAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Okproject) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Okproject) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Okproject) ReloadG() error {
	if o == nil {
		return errors.New("models: no Okproject provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Okproject) Reload(exec boil.Executor) error {
	ret, err := FindOkproject(exec, o.IDProject)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkprojectSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkprojectSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkprojectSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OkprojectSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkprojectSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	okprojects := OkprojectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), okprojectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `okproject`.* FROM `okproject` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, okprojectPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(okprojectPrimaryKeyColumns), 1, len(okprojectPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&okprojects)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OkprojectSlice")
	}

	*o = okprojects

	return nil
}

// OkprojectExists checks if the Okproject row exists.
func OkprojectExists(exec boil.Executor, idProject int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `okproject` where `id_project`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, idProject)
	}

	row := exec.QueryRow(sql, idProject)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if okproject exists")
	}

	return exists, nil
}

// OkprojectExistsG checks if the Okproject row exists.
func OkprojectExistsG(idProject int) (bool, error) {
	return OkprojectExists(boil.GetDB(), idProject)
}

// OkprojectExistsGP checks if the Okproject row exists. Panics on error.
func OkprojectExistsGP(idProject int) bool {
	e, err := OkprojectExists(boil.GetDB(), idProject)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OkprojectExistsP checks if the Okproject row exists. Panics on error.
func OkprojectExistsP(exec boil.Executor, idProject int) bool {
	e, err := OkprojectExists(exec, idProject)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
