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

// OkProject is an object representing the database table.
type OkProject struct {
	IDProject   int         `boil:"id_project" json:"id_project" toml:"id_project" yaml:"id_project"`
	NamaProject string      `boil:"nama_project" json:"nama_project" toml:"nama_project" yaml:"nama_project"`
	Desc        null.String `boil:"desc" json:"desc,omitempty" toml:"desc" yaml:"desc,omitempty"`
	GitURL      string      `boil:"git_url" json:"git_url" toml:"git_url" yaml:"git_url"`

	R *OkProjectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L OkProjectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// OkProjectR is where relationships are stored.
type OkProjectR struct {
	ProjectOkProjectTasks OkProjectTaskSlice
}

// OkProjectL is where Load methods for each relationship are stored.
type OkProjectL struct{}

var (
	OkProjectColumns               = []string{"id_project", "nama_project", "desc", "git_url"}
	OkProjectColumnsWithoutDefault = []string{"nama_project", "desc", "git_url"}
	OkProjectColumnsWithDefault    = []string{"id_project"}
	OkProjectPrimaryKeyColumns     = []string{"id_project"}
)

type (
	// OkProjectSlice is an alias for a slice of pointers to OkProject.
	// This should generally be used opposed to []OkProject.
	OkProjectSlice []*OkProject
	// OkProjectHook is the signature for custom OkProject hook methods
	OkProjectHook func(boil.Executor, *OkProject) error

	OkProjectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	OkProjectType                 = reflect.TypeOf(&OkProject{})
	OkProjectMapping              = queries.MakeStructMapping(OkProjectType)
	OkProjectPrimaryKeyMapping, _ = queries.BindMapping(OkProjectType, OkProjectMapping, OkProjectPrimaryKeyColumns)
	OkProjectInsertCacheMut       sync.RWMutex
	OkProjectInsertCache          = make(map[string]insertCache)
	OkProjectUpdateCacheMut       sync.RWMutex
	OkProjectUpdateCache          = make(map[string]updateCache)
	OkProjectUpsertCacheMut       sync.RWMutex
	OkProjectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var OkProjectBeforeInsertHooks []OkProjectHook
var OkProjectBeforeUpdateHooks []OkProjectHook
var OkProjectBeforeDeleteHooks []OkProjectHook
var OkProjectBeforeUpsertHooks []OkProjectHook

var OkProjectAfterInsertHooks []OkProjectHook
var OkProjectAfterSelectHooks []OkProjectHook
var OkProjectAfterUpdateHooks []OkProjectHook
var OkProjectAfterDeleteHooks []OkProjectHook
var OkProjectAfterUpsertHooks []OkProjectHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OkProject) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OkProject) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OkProject) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OkProject) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OkProject) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OkProject) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OkProject) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OkProject) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OkProject) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOkProjectHook registers your hook function for all future operations.
func AddOkProjectHook(hookPoint boil.HookPoint, OkProjectHook OkProjectHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		OkProjectBeforeInsertHooks = append(OkProjectBeforeInsertHooks, OkProjectHook)
	case boil.BeforeUpdateHook:
		OkProjectBeforeUpdateHooks = append(OkProjectBeforeUpdateHooks, OkProjectHook)
	case boil.BeforeDeleteHook:
		OkProjectBeforeDeleteHooks = append(OkProjectBeforeDeleteHooks, OkProjectHook)
	case boil.BeforeUpsertHook:
		OkProjectBeforeUpsertHooks = append(OkProjectBeforeUpsertHooks, OkProjectHook)
	case boil.AfterInsertHook:
		OkProjectAfterInsertHooks = append(OkProjectAfterInsertHooks, OkProjectHook)
	case boil.AfterSelectHook:
		OkProjectAfterSelectHooks = append(OkProjectAfterSelectHooks, OkProjectHook)
	case boil.AfterUpdateHook:
		OkProjectAfterUpdateHooks = append(OkProjectAfterUpdateHooks, OkProjectHook)
	case boil.AfterDeleteHook:
		OkProjectAfterDeleteHooks = append(OkProjectAfterDeleteHooks, OkProjectHook)
	case boil.AfterUpsertHook:
		OkProjectAfterUpsertHooks = append(OkProjectAfterUpsertHooks, OkProjectHook)
	}
}

// OneP returns a single OkProject record from the query, and panics on error.
func (q OkProjectQuery) OneP() *OkProject {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single OkProject record from the query.
func (q OkProjectQuery) One() (*OkProject, error) {
	o := &OkProject{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OkProject")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all OkProject records from the query, and panics on error.
func (q OkProjectQuery) AllP() OkProjectSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OkProject records from the query.
func (q OkProjectQuery) All() (OkProjectSlice, error) {
	var o OkProjectSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OkProject slice")
	}

	if len(OkProjectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all OkProject records in the query, and panics on error.
func (q OkProjectQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OkProject records in the query.
func (q OkProjectQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OkProject rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q OkProjectQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q OkProjectQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OkProject exists")
	}

	return count > 0, nil
}

// ProjectOkProjectTasksG retrieves all the OkProjectTask's OkProjectTask via project_id column.
func (o *OkProject) ProjectOkProjectTasksG(mods ...qm.QueryMod) OkProjectTaskQuery {
	return o.ProjectOkProjectTasks(boil.GetDB(), mods...)
}

// ProjectOkProjectTasks retrieves all the OkProjectTask's OkProjectTask with an executor via project_id column.
func (o *OkProject) ProjectOkProjectTasks(exec boil.Executor, mods ...qm.QueryMod) OkProjectTaskQuery {
	queryMods := []qm.QueryMod{
		qm.Select("`a`.*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`a`.`project_id`=?", o.IDProject),
	)

	query := OkProjectTasks(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProjectTask` as `a`")
	return query
}

// LoadProjectOkProjectTasks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProjectL) LoadProjectOkProjectTasks(e boil.Executor, singular bool, maybeOkProject interface{}) error {
	var slice []*OkProject
	var object *OkProject

	count := 1
	if singular {
		object = maybeOkProject.(*OkProject)
	} else {
		slice = *maybeOkProject.(*OkProjectSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProjectR{}
		}
		args[0] = object.IDProject
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProjectR{}
			}
			args[i] = obj.IDProject
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProjectTask` where `project_id` in (%s)",
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
		object.R.ProjectOkProjectTasks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IDProject == foreign.ProjectID {
				local.R.ProjectOkProjectTasks = append(local.R.ProjectOkProjectTasks, foreign)
				break
			}
		}
	}

	return nil
}

// AddProjectOkProjectTasks adds the given related objects to the existing relationships
// of the OkProject, optionally inserting them as new records.
// Appends related to o.R.ProjectOkProjectTasks.
// Sets related.R.Project appropriately.
func (o *OkProject) AddProjectOkProjectTasks(exec boil.Executor, insert bool, related ...*OkProjectTask) error {
	var err error
	for _, rel := range related {
		rel.ProjectID = o.IDProject
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "project_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &OkProjectR{
			ProjectOkProjectTasks: related,
		}
	} else {
		o.R.ProjectOkProjectTasks = append(o.R.ProjectOkProjectTasks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &OkProjectTaskR{
				Project: o,
			}
		} else {
			rel.R.Project = o
		}
	}
	return nil
}

// OkProjectsG retrieves all records.
func OkProjectsG(mods ...qm.QueryMod) OkProjectQuery {
	return OkProjects(boil.GetDB(), mods...)
}

// OkProjects retrieves all the records using an executor.
func OkProjects(exec boil.Executor, mods ...qm.QueryMod) OkProjectQuery {
	mods = append(mods, qm.From("`OkProject`"))
	return OkProjectQuery{NewQuery(exec, mods...)}
}

// FindOkProjectG retrieves a single record by ID.
func FindOkProjectG(idProject int, selectCols ...string) (*OkProject, error) {
	return FindOkProject(boil.GetDB(), idProject, selectCols...)
}

// FindOkProjectGP retrieves a single record by ID, and panics on error.
func FindOkProjectGP(idProject int, selectCols ...string) *OkProject {
	retobj, err := FindOkProject(boil.GetDB(), idProject, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOkProject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOkProject(exec boil.Executor, idProject int, selectCols ...string) (*OkProject, error) {
	OkProjectObj := &OkProject{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OkProject` where `id_project`=?", sel,
	)

	q := queries.Raw(exec, query, idProject)

	err := q.Bind(OkProjectObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OkProject")
	}

	return OkProjectObj, nil
}

// FindOkProjectP retrieves a single record by ID with an executor, and panics on error.
func FindOkProjectP(exec boil.Executor, idProject int, selectCols ...string) *OkProject {
	retobj, err := FindOkProject(exec, idProject, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OkProject) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OkProject) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OkProject) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OkProject) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkProject provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(OkProjectColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	OkProjectInsertCacheMut.RLock()
	cache, cached := OkProjectInsertCache[key]
	OkProjectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			OkProjectColumns,
			OkProjectColumnsWithDefault,
			OkProjectColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(OkProjectType, OkProjectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(OkProjectType, OkProjectMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `OkProject` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OkProject` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, OkProjectPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into OkProject")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkProjectMapping["IDProject"] {
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
		return errors.Wrap(err, "models: unable to populate default values for OkProject")
	}

CacheNoHooks:
	if !cached {
		OkProjectInsertCacheMut.Lock()
		OkProjectInsertCache[key] = cache
		OkProjectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single OkProject record. See Update for
// whitelist behavior description.
func (o *OkProject) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OkProject record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OkProject) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OkProject, and panics on error.
// See Update for whitelist behavior description.
func (o *OkProject) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OkProject.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OkProject) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	OkProjectUpdateCacheMut.RLock()
	cache, cached := OkProjectUpdateCache[key]
	OkProjectUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(OkProjectColumns, OkProjectPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update OkProject, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OkProject` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, OkProjectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(OkProjectType, OkProjectMapping, append(wl, OkProjectPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update OkProject row")
	}

	if !cached {
		OkProjectUpdateCacheMut.Lock()
		OkProjectUpdateCache[key] = cache
		OkProjectUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q OkProjectQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q OkProjectQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for OkProject")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OkProjectSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OkProjectSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OkProjectSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OkProjectSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `OkProject` SET %s WHERE (`id_project`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkProjectPrimaryKeyColumns), len(colNames)+1, len(OkProjectPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in OkProject slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OkProject) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OkProject) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OkProject) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OkProject) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkProject provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(OkProjectColumnsWithDefault, o)

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

	OkProjectUpsertCacheMut.RLock()
	cache, cached := OkProjectUpsertCache[key]
	OkProjectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			OkProjectColumns,
			OkProjectColumnsWithDefault,
			OkProjectColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			OkProjectColumns,
			OkProjectPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert OkProject, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "OkProject", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OkProject` WHERE `id_project`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(OkProjectType, OkProjectMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(OkProjectType, OkProjectMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for OkProject")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkProjectMapping["IDProject"] {
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
		return errors.Wrap(err, "models: unable to populate default values for OkProject")
	}

CacheNoHooks:
	if !cached {
		OkProjectUpsertCacheMut.Lock()
		OkProjectUpsertCache[key] = cache
		OkProjectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single OkProject record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkProject) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OkProject record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OkProject) DeleteG() error {
	if o == nil {
		return errors.New("models: no OkProject provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OkProject record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkProject) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OkProject record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OkProject) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkProject provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), OkProjectPrimaryKeyMapping)
	sql := "DELETE FROM `OkProject` WHERE `id_project`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from OkProject")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q OkProjectQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q OkProjectQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no OkProjectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkProject")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OkProjectSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OkProjectSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no OkProject slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OkProjectSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OkProjectSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkProject slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(OkProjectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `OkProject` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkProjectPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkProjectPrimaryKeyColumns), 1, len(OkProjectPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkProject slice")
	}

	if len(OkProjectAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OkProject) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OkProject) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OkProject) ReloadG() error {
	if o == nil {
		return errors.New("models: no OkProject provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OkProject) Reload(exec boil.Executor) error {
	ret, err := FindOkProject(exec, o.IDProject)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkProjectSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkProjectSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkProjectSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OkProjectSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkProjectSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	OkProjects := OkProjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `OkProject`.* FROM `OkProject` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkProjectPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(OkProjectPrimaryKeyColumns), 1, len(OkProjectPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&OkProjects)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OkProjectSlice")
	}

	*o = OkProjects

	return nil
}

// OkProjectExists checks if the OkProject row exists.
func OkProjectExists(exec boil.Executor, idProject int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `OkProject` where `id_project`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, idProject)
	}

	row := exec.QueryRow(sql, idProject)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OkProject exists")
	}

	return exists, nil
}

// OkProjectExistsG checks if the OkProject row exists.
func OkProjectExistsG(idProject int) (bool, error) {
	return OkProjectExists(boil.GetDB(), idProject)
}

// OkProjectExistsGP checks if the OkProject row exists. Panics on error.
func OkProjectExistsGP(idProject int) bool {
	e, err := OkProjectExists(boil.GetDB(), idProject)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OkProjectExistsP checks if the OkProject row exists. Panics on error.
func OkProjectExistsP(exec boil.Executor, idProject int) bool {
	e, err := OkProjectExists(exec, idProject)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
