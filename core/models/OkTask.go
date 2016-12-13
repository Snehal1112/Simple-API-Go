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

// OkTask is an object representing the database table.
type OkTask struct {
	IDTask     int         `boil:"id_task" json:"id_task" toml:"id_task" yaml:"id_task"`
	TitleTask  string      `boil:"title_task" json:"title_task" toml:"title_task" yaml:"title_task"`
	Urgency    string      `boil:"urgency" json:"urgency" toml:"urgency" yaml:"urgency"`
	Difficulty string      `boil:"difficulty" json:"difficulty" toml:"difficulty" yaml:"difficulty"`
	DescTask   null.String `boil:"desc_task" json:"desc_task,omitempty" toml:"desc_task" yaml:"desc_task,omitempty"`

	R *OkTaskR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L OkTaskL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// OkTaskR is where relationships are stored.
type OkTaskR struct {
	TaskOkProjectTasks OkProjectTaskSlice
}

// OkTaskL is where Load methods for each relationship are stored.
type OkTaskL struct{}

var (
	OkTaskColumns               = []string{"id_task", "title_task", "urgency", "difficulty", "desc_task"}
	OkTaskColumnsWithoutDefault = []string{"title_task", "urgency", "difficulty", "desc_task"}
	OkTaskColumnsWithDefault    = []string{"id_task"}
	OkTaskPrimaryKeyColumns     = []string{"id_task"}
)

type (
	// OkTaskSlice is an alias for a slice of pointers to OkTask.
	// This should generally be used opposed to []OkTask.
	OkTaskSlice []*OkTask
	// OkTaskHook is the signature for custom OkTask hook methods
	OkTaskHook func(boil.Executor, *OkTask) error

	OkTaskQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	OkTaskType                 = reflect.TypeOf(&OkTask{})
	OkTaskMapping              = queries.MakeStructMapping(OkTaskType)
	OkTaskPrimaryKeyMapping, _ = queries.BindMapping(OkTaskType, OkTaskMapping, OkTaskPrimaryKeyColumns)
	OkTaskInsertCacheMut       sync.RWMutex
	OkTaskInsertCache          = make(map[string]insertCache)
	OkTaskUpdateCacheMut       sync.RWMutex
	OkTaskUpdateCache          = make(map[string]updateCache)
	OkTaskUpsertCacheMut       sync.RWMutex
	OkTaskUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var OkTaskBeforeInsertHooks []OkTaskHook
var OkTaskBeforeUpdateHooks []OkTaskHook
var OkTaskBeforeDeleteHooks []OkTaskHook
var OkTaskBeforeUpsertHooks []OkTaskHook

var OkTaskAfterInsertHooks []OkTaskHook
var OkTaskAfterSelectHooks []OkTaskHook
var OkTaskAfterUpdateHooks []OkTaskHook
var OkTaskAfterDeleteHooks []OkTaskHook
var OkTaskAfterUpsertHooks []OkTaskHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OkTask) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OkTask) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OkTask) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OkTask) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OkTask) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OkTask) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OkTask) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OkTask) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OkTask) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkTaskAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOkTaskHook registers your hook function for all future operations.
func AddOkTaskHook(hookPoint boil.HookPoint, OkTaskHook OkTaskHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		OkTaskBeforeInsertHooks = append(OkTaskBeforeInsertHooks, OkTaskHook)
	case boil.BeforeUpdateHook:
		OkTaskBeforeUpdateHooks = append(OkTaskBeforeUpdateHooks, OkTaskHook)
	case boil.BeforeDeleteHook:
		OkTaskBeforeDeleteHooks = append(OkTaskBeforeDeleteHooks, OkTaskHook)
	case boil.BeforeUpsertHook:
		OkTaskBeforeUpsertHooks = append(OkTaskBeforeUpsertHooks, OkTaskHook)
	case boil.AfterInsertHook:
		OkTaskAfterInsertHooks = append(OkTaskAfterInsertHooks, OkTaskHook)
	case boil.AfterSelectHook:
		OkTaskAfterSelectHooks = append(OkTaskAfterSelectHooks, OkTaskHook)
	case boil.AfterUpdateHook:
		OkTaskAfterUpdateHooks = append(OkTaskAfterUpdateHooks, OkTaskHook)
	case boil.AfterDeleteHook:
		OkTaskAfterDeleteHooks = append(OkTaskAfterDeleteHooks, OkTaskHook)
	case boil.AfterUpsertHook:
		OkTaskAfterUpsertHooks = append(OkTaskAfterUpsertHooks, OkTaskHook)
	}
}

// OneP returns a single OkTask record from the query, and panics on error.
func (q OkTaskQuery) OneP() *OkTask {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single OkTask record from the query.
func (q OkTaskQuery) One() (*OkTask, error) {
	o := &OkTask{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OkTask")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all OkTask records from the query, and panics on error.
func (q OkTaskQuery) AllP() OkTaskSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OkTask records from the query.
func (q OkTaskQuery) All() (OkTaskSlice, error) {
	var o OkTaskSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OkTask slice")
	}

	if len(OkTaskAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all OkTask records in the query, and panics on error.
func (q OkTaskQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OkTask records in the query.
func (q OkTaskQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OkTask rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q OkTaskQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q OkTaskQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OkTask exists")
	}

	return count > 0, nil
}

// TaskOkProjectTasksG retrieves all the OkProjectTask's OkProjectTask via task_id column.
func (o *OkTask) TaskOkProjectTasksG(mods ...qm.QueryMod) OkProjectTaskQuery {
	return o.TaskOkProjectTasks(boil.GetDB(), mods...)
}

// TaskOkProjectTasks retrieves all the OkProjectTask's OkProjectTask with an executor via task_id column.
func (o *OkTask) TaskOkProjectTasks(exec boil.Executor, mods ...qm.QueryMod) OkProjectTaskQuery {
	queryMods := []qm.QueryMod{
		qm.Select("`a`.*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`a`.`task_id`=?", o.IDTask),
	)

	query := OkProjectTasks(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProjectTask` as `a`")
	return query
}

// LoadTaskOkProjectTasks allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkTaskL) LoadTaskOkProjectTasks(e boil.Executor, singular bool, maybeOkTask interface{}) error {
	var slice []*OkTask
	var object *OkTask

	count := 1
	if singular {
		object = maybeOkTask.(*OkTask)
	} else {
		slice = *maybeOkTask.(*OkTaskSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkTaskR{}
		}
		args[0] = object.IDTask
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkTaskR{}
			}
			args[i] = obj.IDTask
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProjectTask` where `task_id` in (%s)",
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
		object.R.TaskOkProjectTasks = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IDTask == foreign.TaskID {
				local.R.TaskOkProjectTasks = append(local.R.TaskOkProjectTasks, foreign)
				break
			}
		}
	}

	return nil
}

// AddTaskOkProjectTasks adds the given related objects to the existing relationships
// of the OkTask, optionally inserting them as new records.
// Appends related to o.R.TaskOkProjectTasks.
// Sets related.R.Task appropriately.
func (o *OkTask) AddTaskOkProjectTasks(exec boil.Executor, insert bool, related ...*OkProjectTask) error {
	var err error
	for _, rel := range related {
		rel.TaskID = o.IDTask
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "task_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &OkTaskR{
			TaskOkProjectTasks: related,
		}
	} else {
		o.R.TaskOkProjectTasks = append(o.R.TaskOkProjectTasks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &OkProjectTaskR{
				Task: o,
			}
		} else {
			rel.R.Task = o
		}
	}
	return nil
}

// OkTasksG retrieves all records.
func OkTasksG(mods ...qm.QueryMod) OkTaskQuery {
	return OkTasks(boil.GetDB(), mods...)
}

// OkTasks retrieves all the records using an executor.
func OkTasks(exec boil.Executor, mods ...qm.QueryMod) OkTaskQuery {
	mods = append(mods, qm.From("`OkTask`"))
	return OkTaskQuery{NewQuery(exec, mods...)}
}

// FindOkTaskG retrieves a single record by ID.
func FindOkTaskG(idTask int, selectCols ...string) (*OkTask, error) {
	return FindOkTask(boil.GetDB(), idTask, selectCols...)
}

// FindOkTaskGP retrieves a single record by ID, and panics on error.
func FindOkTaskGP(idTask int, selectCols ...string) *OkTask {
	retobj, err := FindOkTask(boil.GetDB(), idTask, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOkTask retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOkTask(exec boil.Executor, idTask int, selectCols ...string) (*OkTask, error) {
	OkTaskObj := &OkTask{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OkTask` where `id_task`=?", sel,
	)

	q := queries.Raw(exec, query, idTask)

	err := q.Bind(OkTaskObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OkTask")
	}

	return OkTaskObj, nil
}

// FindOkTaskP retrieves a single record by ID with an executor, and panics on error.
func FindOkTaskP(exec boil.Executor, idTask int, selectCols ...string) *OkTask {
	retobj, err := FindOkTask(exec, idTask, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OkTask) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OkTask) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OkTask) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OkTask) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkTask provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(OkTaskColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	OkTaskInsertCacheMut.RLock()
	cache, cached := OkTaskInsertCache[key]
	OkTaskInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			OkTaskColumns,
			OkTaskColumnsWithDefault,
			OkTaskColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(OkTaskType, OkTaskMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(OkTaskType, OkTaskMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `OkTask` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OkTask` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, OkTaskPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into OkTask")
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

	o.IDTask = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkTaskMapping["IDTask"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDTask,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkTask")
	}

CacheNoHooks:
	if !cached {
		OkTaskInsertCacheMut.Lock()
		OkTaskInsertCache[key] = cache
		OkTaskInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single OkTask record. See Update for
// whitelist behavior description.
func (o *OkTask) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OkTask record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OkTask) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OkTask, and panics on error.
// See Update for whitelist behavior description.
func (o *OkTask) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OkTask.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OkTask) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	OkTaskUpdateCacheMut.RLock()
	cache, cached := OkTaskUpdateCache[key]
	OkTaskUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(OkTaskColumns, OkTaskPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update OkTask, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OkTask` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, OkTaskPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(OkTaskType, OkTaskMapping, append(wl, OkTaskPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update OkTask row")
	}

	if !cached {
		OkTaskUpdateCacheMut.Lock()
		OkTaskUpdateCache[key] = cache
		OkTaskUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q OkTaskQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q OkTaskQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for OkTask")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OkTaskSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OkTaskSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OkTaskSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OkTaskSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `OkTask` SET %s WHERE (`id_task`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkTaskPrimaryKeyColumns), len(colNames)+1, len(OkTaskPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in OkTask slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OkTask) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OkTask) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OkTask) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OkTask) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkTask provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(OkTaskColumnsWithDefault, o)

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

	OkTaskUpsertCacheMut.RLock()
	cache, cached := OkTaskUpsertCache[key]
	OkTaskUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			OkTaskColumns,
			OkTaskColumnsWithDefault,
			OkTaskColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			OkTaskColumns,
			OkTaskPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert OkTask, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "OkTask", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OkTask` WHERE `id_task`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(OkTaskType, OkTaskMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(OkTaskType, OkTaskMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for OkTask")
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

	o.IDTask = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkTaskMapping["IDTask"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDTask,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkTask")
	}

CacheNoHooks:
	if !cached {
		OkTaskUpsertCacheMut.Lock()
		OkTaskUpsertCache[key] = cache
		OkTaskUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single OkTask record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkTask) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OkTask record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OkTask) DeleteG() error {
	if o == nil {
		return errors.New("models: no OkTask provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OkTask record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkTask) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OkTask record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OkTask) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkTask provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), OkTaskPrimaryKeyMapping)
	sql := "DELETE FROM `OkTask` WHERE `id_task`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from OkTask")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q OkTaskQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q OkTaskQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no OkTaskQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkTask")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OkTaskSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OkTaskSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no OkTask slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OkTaskSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OkTaskSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkTask slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(OkTaskBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `OkTask` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkTaskPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkTaskPrimaryKeyColumns), 1, len(OkTaskPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkTask slice")
	}

	if len(OkTaskAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OkTask) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OkTask) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OkTask) ReloadG() error {
	if o == nil {
		return errors.New("models: no OkTask provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OkTask) Reload(exec boil.Executor) error {
	ret, err := FindOkTask(exec, o.IDTask)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkTaskSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkTaskSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkTaskSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OkTaskSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkTaskSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	OkTasks := OkTaskSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `OkTask`.* FROM `OkTask` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkTaskPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(OkTaskPrimaryKeyColumns), 1, len(OkTaskPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&OkTasks)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OkTaskSlice")
	}

	*o = OkTasks

	return nil
}

// OkTaskExists checks if the OkTask row exists.
func OkTaskExists(exec boil.Executor, idTask int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `OkTask` where `id_task`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, idTask)
	}

	row := exec.QueryRow(sql, idTask)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OkTask exists")
	}

	return exists, nil
}

// OkTaskExistsG checks if the OkTask row exists.
func OkTaskExistsG(idTask int) (bool, error) {
	return OkTaskExists(boil.GetDB(), idTask)
}

// OkTaskExistsGP checks if the OkTask row exists. Panics on error.
func OkTaskExistsGP(idTask int) bool {
	e, err := OkTaskExists(boil.GetDB(), idTask)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OkTaskExistsP checks if the OkTask row exists. Panics on error.
func OkTaskExistsP(exec boil.Executor, idTask int) bool {
	e, err := OkTaskExists(exec, idTask)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
