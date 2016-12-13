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

// OkProjectTask is an object representing the database table.
type OkProjectTask struct {
	IDPTask      int       `boil:"id_p_task" json:"id_p_task" toml:"id_p_task" yaml:"id_p_task"`
	ProjectID    int       `boil:"project_id" json:"project_id" toml:"project_id" yaml:"project_id"`
	TaskID       int       `boil:"task_id" json:"task_id" toml:"task_id" yaml:"task_id"`
	AssignedtoID int       `boil:"assignedto_id" json:"assignedto_id" toml:"assignedto_id" yaml:"assignedto_id"`
	AssignerID   int       `boil:"assigner_id" json:"assigner_id" toml:"assigner_id" yaml:"assigner_id"`
	CreatedAt    null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt    null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Status       string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	DueDate      null.Time `boil:"due_date" json:"due_date,omitempty" toml:"due_date" yaml:"due_date,omitempty"`

	R *OkProjectTaskR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L OkProjectTaskL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// OkProjectTaskR is where relationships are stored.
type OkProjectTaskR struct {
	Assignedto *OkProfile
	Assigner   *OkProfile
	Project    *OkProject
	Task       *OkTask
}

// OkProjectTaskL is where Load methods for each relationship are stored.
type OkProjectTaskL struct{}

var (
	OkProjectTaskColumns               = []string{"id_p_task", "project_id", "task_id", "assignedto_id", "assigner_id", "created_at", "updated_at", "status", "due_date"}
	OkProjectTaskColumnsWithoutDefault = []string{"project_id", "task_id", "assignedto_id", "assigner_id", "created_at", "updated_at", "status", "due_date"}
	OkProjectTaskColumnsWithDefault    = []string{"id_p_task"}
	OkProjectTaskPrimaryKeyColumns     = []string{"id_p_task"}
)

type (
	// OkProjectTaskSlice is an alias for a slice of pointers to OkProjectTask.
	// This should generally be used opposed to []OkProjectTask.
	OkProjectTaskSlice []*OkProjectTask
	// OkProjectTaskHook is the signature for custom OkProjectTask hook methods
	OkProjectTaskHook func(boil.Executor, *OkProjectTask) error

	OkProjectTaskQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	OkProjectTaskType                 = reflect.TypeOf(&OkProjectTask{})
	OkProjectTaskMapping              = queries.MakeStructMapping(OkProjectTaskType)
	OkProjectTaskPrimaryKeyMapping, _ = queries.BindMapping(OkProjectTaskType, OkProjectTaskMapping, OkProjectTaskPrimaryKeyColumns)
	OkProjectTaskInsertCacheMut       sync.RWMutex
	OkProjectTaskInsertCache          = make(map[string]insertCache)
	OkProjectTaskUpdateCacheMut       sync.RWMutex
	OkProjectTaskUpdateCache          = make(map[string]updateCache)
	OkProjectTaskUpsertCacheMut       sync.RWMutex
	OkProjectTaskUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var OkProjectTaskBeforeInsertHooks []OkProjectTaskHook
var OkProjectTaskBeforeUpdateHooks []OkProjectTaskHook
var OkProjectTaskBeforeDeleteHooks []OkProjectTaskHook
var OkProjectTaskBeforeUpsertHooks []OkProjectTaskHook

var OkProjectTaskAfterInsertHooks []OkProjectTaskHook
var OkProjectTaskAfterSelectHooks []OkProjectTaskHook
var OkProjectTaskAfterUpdateHooks []OkProjectTaskHook
var OkProjectTaskAfterDeleteHooks []OkProjectTaskHook
var OkProjectTaskAfterUpsertHooks []OkProjectTaskHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OkProjectTask) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OkProjectTask) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OkProjectTask) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OkProjectTask) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OkProjectTask) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OkProjectTask) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OkProjectTask) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OkProjectTask) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OkProjectTask) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkProjectTaskAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOkProjectTaskHook registers your hook function for all future operations.
func AddOkProjectTaskHook(hookPoint boil.HookPoint, OkProjectTaskHook OkProjectTaskHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		OkProjectTaskBeforeInsertHooks = append(OkProjectTaskBeforeInsertHooks, OkProjectTaskHook)
	case boil.BeforeUpdateHook:
		OkProjectTaskBeforeUpdateHooks = append(OkProjectTaskBeforeUpdateHooks, OkProjectTaskHook)
	case boil.BeforeDeleteHook:
		OkProjectTaskBeforeDeleteHooks = append(OkProjectTaskBeforeDeleteHooks, OkProjectTaskHook)
	case boil.BeforeUpsertHook:
		OkProjectTaskBeforeUpsertHooks = append(OkProjectTaskBeforeUpsertHooks, OkProjectTaskHook)
	case boil.AfterInsertHook:
		OkProjectTaskAfterInsertHooks = append(OkProjectTaskAfterInsertHooks, OkProjectTaskHook)
	case boil.AfterSelectHook:
		OkProjectTaskAfterSelectHooks = append(OkProjectTaskAfterSelectHooks, OkProjectTaskHook)
	case boil.AfterUpdateHook:
		OkProjectTaskAfterUpdateHooks = append(OkProjectTaskAfterUpdateHooks, OkProjectTaskHook)
	case boil.AfterDeleteHook:
		OkProjectTaskAfterDeleteHooks = append(OkProjectTaskAfterDeleteHooks, OkProjectTaskHook)
	case boil.AfterUpsertHook:
		OkProjectTaskAfterUpsertHooks = append(OkProjectTaskAfterUpsertHooks, OkProjectTaskHook)
	}
}

// OneP returns a single OkProjectTask record from the query, and panics on error.
func (q OkProjectTaskQuery) OneP() *OkProjectTask {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single OkProjectTask record from the query.
func (q OkProjectTaskQuery) One() (*OkProjectTask, error) {
	o := &OkProjectTask{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OkProjectTask")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all OkProjectTask records from the query, and panics on error.
func (q OkProjectTaskQuery) AllP() OkProjectTaskSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OkProjectTask records from the query.
func (q OkProjectTaskQuery) All() (OkProjectTaskSlice, error) {
	var o OkProjectTaskSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OkProjectTask slice")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all OkProjectTask records in the query, and panics on error.
func (q OkProjectTaskQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OkProjectTask records in the query.
func (q OkProjectTaskQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OkProjectTask rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q OkProjectTaskQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q OkProjectTaskQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OkProjectTask exists")
	}

	return count > 0, nil
}

// AssignedtoG pointed to by the foreign key.
func (o *OkProjectTask) AssignedtoG(mods ...qm.QueryMod) OkProfileQuery {
	return o.Assignedto(boil.GetDB(), mods...)
}

// Assignedto pointed to by the foreign key.
func (o *OkProjectTask) Assignedto(exec boil.Executor, mods ...qm.QueryMod) OkProfileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id_profile=?", o.AssignedtoID),
	}

	queryMods = append(queryMods, mods...)

	query := OkProfiles(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProfile`")

	return query
}

// AssignerG pointed to by the foreign key.
func (o *OkProjectTask) AssignerG(mods ...qm.QueryMod) OkProfileQuery {
	return o.Assigner(boil.GetDB(), mods...)
}

// Assigner pointed to by the foreign key.
func (o *OkProjectTask) Assigner(exec boil.Executor, mods ...qm.QueryMod) OkProfileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id_profile=?", o.AssignerID),
	}

	queryMods = append(queryMods, mods...)

	query := OkProfiles(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProfile`")

	return query
}

// ProjectG pointed to by the foreign key.
func (o *OkProjectTask) ProjectG(mods ...qm.QueryMod) OkProjectQuery {
	return o.Project(boil.GetDB(), mods...)
}

// Project pointed to by the foreign key.
func (o *OkProjectTask) Project(exec boil.Executor, mods ...qm.QueryMod) OkProjectQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id_project=?", o.ProjectID),
	}

	queryMods = append(queryMods, mods...)

	query := OkProjects(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProject`")

	return query
}

// TaskG pointed to by the foreign key.
func (o *OkProjectTask) TaskG(mods ...qm.QueryMod) OkTaskQuery {
	return o.Task(boil.GetDB(), mods...)
}

// Task pointed to by the foreign key.
func (o *OkProjectTask) Task(exec boil.Executor, mods ...qm.QueryMod) OkTaskQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id_task=?", o.TaskID),
	}

	queryMods = append(queryMods, mods...)

	query := OkTasks(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkTask`")

	return query
}

// LoadAssignedto allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProjectTaskL) LoadAssignedto(e boil.Executor, singular bool, maybeOkProjectTask interface{}) error {
	var slice []*OkProjectTask
	var object *OkProjectTask

	count := 1
	if singular {
		object = maybeOkProjectTask.(*OkProjectTask)
	} else {
		slice = *maybeOkProjectTask.(*OkProjectTaskSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProjectTaskR{}
		}
		args[0] = object.AssignedtoID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProjectTaskR{}
			}
			args[i] = obj.AssignedtoID
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProfile` where `id_profile` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkProfile")
	}
	defer results.Close()

	var resultSlice []*OkProfile
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkProfile")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Assignedto = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AssignedtoID == foreign.IDProfile {
				local.R.Assignedto = foreign
				break
			}
		}
	}

	return nil
}

// LoadAssigner allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProjectTaskL) LoadAssigner(e boil.Executor, singular bool, maybeOkProjectTask interface{}) error {
	var slice []*OkProjectTask
	var object *OkProjectTask

	count := 1
	if singular {
		object = maybeOkProjectTask.(*OkProjectTask)
	} else {
		slice = *maybeOkProjectTask.(*OkProjectTaskSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProjectTaskR{}
		}
		args[0] = object.AssignerID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProjectTaskR{}
			}
			args[i] = obj.AssignerID
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProfile` where `id_profile` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkProfile")
	}
	defer results.Close()

	var resultSlice []*OkProfile
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkProfile")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Assigner = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.AssignerID == foreign.IDProfile {
				local.R.Assigner = foreign
				break
			}
		}
	}

	return nil
}

// LoadProject allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProjectTaskL) LoadProject(e boil.Executor, singular bool, maybeOkProjectTask interface{}) error {
	var slice []*OkProjectTask
	var object *OkProjectTask

	count := 1
	if singular {
		object = maybeOkProjectTask.(*OkProjectTask)
	} else {
		slice = *maybeOkProjectTask.(*OkProjectTaskSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProjectTaskR{}
		}
		args[0] = object.ProjectID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProjectTaskR{}
			}
			args[i] = obj.ProjectID
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProject` where `id_project` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkProject")
	}
	defer results.Close()

	var resultSlice []*OkProject
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkProject")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Project = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ProjectID == foreign.IDProject {
				local.R.Project = foreign
				break
			}
		}
	}

	return nil
}

// LoadTask allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkProjectTaskL) LoadTask(e boil.Executor, singular bool, maybeOkProjectTask interface{}) error {
	var slice []*OkProjectTask
	var object *OkProjectTask

	count := 1
	if singular {
		object = maybeOkProjectTask.(*OkProjectTask)
	} else {
		slice = *maybeOkProjectTask.(*OkProjectTaskSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkProjectTaskR{}
		}
		args[0] = object.TaskID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkProjectTaskR{}
			}
			args[i] = obj.TaskID
		}
	}

	query := fmt.Sprintf(
		"select * from `OkTask` where `id_task` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)

	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load OkTask")
	}
	defer results.Close()

	var resultSlice []*OkTask
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice OkTask")
	}

	if len(OkProjectTaskAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if singular && len(resultSlice) != 0 {
		object.R.Task = resultSlice[0]
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TaskID == foreign.IDTask {
				local.R.Task = foreign
				break
			}
		}
	}

	return nil
}

// SetAssignedto of the OkProjectTask to the related item.
// Sets o.R.Assignedto to related.
// Adds o to related.R.AssignedtoOkProjectTasks.
func (o *OkProjectTask) SetAssignedto(exec boil.Executor, insert bool, related *OkProfile) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `OkProjectTask` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"assignedto_id"}),
		strmangle.WhereClause("`", "`", 0, OkProjectTaskPrimaryKeyColumns),
	)
	values := []interface{}{related.IDProfile, o.IDPTask}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AssignedtoID = related.IDProfile

	if o.R == nil {
		o.R = &OkProjectTaskR{
			Assignedto: related,
		}
	} else {
		o.R.Assignedto = related
	}

	if related.R == nil {
		related.R = &OkProfileR{
			AssignedtoOkProjectTasks: OkProjectTaskSlice{o},
		}
	} else {
		related.R.AssignedtoOkProjectTasks = append(related.R.AssignedtoOkProjectTasks, o)
	}

	return nil
}

// SetAssigner of the OkProjectTask to the related item.
// Sets o.R.Assigner to related.
// Adds o to related.R.AssignerOkProjectTasks.
func (o *OkProjectTask) SetAssigner(exec boil.Executor, insert bool, related *OkProfile) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `OkProjectTask` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"assigner_id"}),
		strmangle.WhereClause("`", "`", 0, OkProjectTaskPrimaryKeyColumns),
	)
	values := []interface{}{related.IDProfile, o.IDPTask}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AssignerID = related.IDProfile

	if o.R == nil {
		o.R = &OkProjectTaskR{
			Assigner: related,
		}
	} else {
		o.R.Assigner = related
	}

	if related.R == nil {
		related.R = &OkProfileR{
			AssignerOkProjectTasks: OkProjectTaskSlice{o},
		}
	} else {
		related.R.AssignerOkProjectTasks = append(related.R.AssignerOkProjectTasks, o)
	}

	return nil
}

// SetProject of the OkProjectTask to the related item.
// Sets o.R.Project to related.
// Adds o to related.R.ProjectOkProjectTasks.
func (o *OkProjectTask) SetProject(exec boil.Executor, insert bool, related *OkProject) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `OkProjectTask` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"project_id"}),
		strmangle.WhereClause("`", "`", 0, OkProjectTaskPrimaryKeyColumns),
	)
	values := []interface{}{related.IDProject, o.IDPTask}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProjectID = related.IDProject

	if o.R == nil {
		o.R = &OkProjectTaskR{
			Project: related,
		}
	} else {
		o.R.Project = related
	}

	if related.R == nil {
		related.R = &OkProjectR{
			ProjectOkProjectTasks: OkProjectTaskSlice{o},
		}
	} else {
		related.R.ProjectOkProjectTasks = append(related.R.ProjectOkProjectTasks, o)
	}

	return nil
}

// SetTask of the OkProjectTask to the related item.
// Sets o.R.Task to related.
// Adds o to related.R.TaskOkProjectTasks.
func (o *OkProjectTask) SetTask(exec boil.Executor, insert bool, related *OkTask) error {
	var err error
	if insert {
		if err = related.Insert(exec); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `OkProjectTask` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"task_id"}),
		strmangle.WhereClause("`", "`", 0, OkProjectTaskPrimaryKeyColumns),
	)
	values := []interface{}{related.IDTask, o.IDPTask}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TaskID = related.IDTask

	if o.R == nil {
		o.R = &OkProjectTaskR{
			Task: related,
		}
	} else {
		o.R.Task = related
	}

	if related.R == nil {
		related.R = &OkTaskR{
			TaskOkProjectTasks: OkProjectTaskSlice{o},
		}
	} else {
		related.R.TaskOkProjectTasks = append(related.R.TaskOkProjectTasks, o)
	}

	return nil
}

// OkProjectTasksG retrieves all records.
func OkProjectTasksG(mods ...qm.QueryMod) OkProjectTaskQuery {
	return OkProjectTasks(boil.GetDB(), mods...)
}

// OkProjectTasks retrieves all the records using an executor.
func OkProjectTasks(exec boil.Executor, mods ...qm.QueryMod) OkProjectTaskQuery {
	mods = append(mods, qm.From("`OkProjectTask`"))
	return OkProjectTaskQuery{NewQuery(exec, mods...)}
}

// FindOkProjectTaskG retrieves a single record by ID.
func FindOkProjectTaskG(idPTask int, selectCols ...string) (*OkProjectTask, error) {
	return FindOkProjectTask(boil.GetDB(), idPTask, selectCols...)
}

// FindOkProjectTaskGP retrieves a single record by ID, and panics on error.
func FindOkProjectTaskGP(idPTask int, selectCols ...string) *OkProjectTask {
	retobj, err := FindOkProjectTask(boil.GetDB(), idPTask, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOkProjectTask retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOkProjectTask(exec boil.Executor, idPTask int, selectCols ...string) (*OkProjectTask, error) {
	OkProjectTaskObj := &OkProjectTask{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OkProjectTask` where `id_p_task`=?", sel,
	)

	q := queries.Raw(exec, query, idPTask)

	err := q.Bind(OkProjectTaskObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OkProjectTask")
	}

	return OkProjectTaskObj, nil
}

// FindOkProjectTaskP retrieves a single record by ID with an executor, and panics on error.
func FindOkProjectTaskP(exec boil.Executor, idPTask int, selectCols ...string) *OkProjectTask {
	retobj, err := FindOkProjectTask(exec, idPTask, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OkProjectTask) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OkProjectTask) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OkProjectTask) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OkProjectTask) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkProjectTask provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(OkProjectTaskColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	OkProjectTaskInsertCacheMut.RLock()
	cache, cached := OkProjectTaskInsertCache[key]
	OkProjectTaskInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			OkProjectTaskColumns,
			OkProjectTaskColumnsWithDefault,
			OkProjectTaskColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(OkProjectTaskType, OkProjectTaskMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(OkProjectTaskType, OkProjectTaskMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `OkProjectTask` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OkProjectTask` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, OkProjectTaskPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into OkProjectTask")
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

	o.IDPTask = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkProjectTaskMapping["IDPTask"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDPTask,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkProjectTask")
	}

CacheNoHooks:
	if !cached {
		OkProjectTaskInsertCacheMut.Lock()
		OkProjectTaskInsertCache[key] = cache
		OkProjectTaskInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single OkProjectTask record. See Update for
// whitelist behavior description.
func (o *OkProjectTask) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OkProjectTask record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OkProjectTask) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OkProjectTask, and panics on error.
// See Update for whitelist behavior description.
func (o *OkProjectTask) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OkProjectTask.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OkProjectTask) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	OkProjectTaskUpdateCacheMut.RLock()
	cache, cached := OkProjectTaskUpdateCache[key]
	OkProjectTaskUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(OkProjectTaskColumns, OkProjectTaskPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update OkProjectTask, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OkProjectTask` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, OkProjectTaskPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(OkProjectTaskType, OkProjectTaskMapping, append(wl, OkProjectTaskPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update OkProjectTask row")
	}

	if !cached {
		OkProjectTaskUpdateCacheMut.Lock()
		OkProjectTaskUpdateCache[key] = cache
		OkProjectTaskUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q OkProjectTaskQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q OkProjectTaskQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for OkProjectTask")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OkProjectTaskSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OkProjectTaskSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OkProjectTaskSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OkProjectTaskSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProjectTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `OkProjectTask` SET %s WHERE (`id_p_task`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkProjectTaskPrimaryKeyColumns), len(colNames)+1, len(OkProjectTaskPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in OkProjectTask slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OkProjectTask) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OkProjectTask) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OkProjectTask) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OkProjectTask) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkProjectTask provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(OkProjectTaskColumnsWithDefault, o)

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

	OkProjectTaskUpsertCacheMut.RLock()
	cache, cached := OkProjectTaskUpsertCache[key]
	OkProjectTaskUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			OkProjectTaskColumns,
			OkProjectTaskColumnsWithDefault,
			OkProjectTaskColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			OkProjectTaskColumns,
			OkProjectTaskPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert OkProjectTask, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "OkProjectTask", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OkProjectTask` WHERE `id_p_task`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(OkProjectTaskType, OkProjectTaskMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(OkProjectTaskType, OkProjectTaskMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for OkProjectTask")
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

	o.IDPTask = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkProjectTaskMapping["IDPTask"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDPTask,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkProjectTask")
	}

CacheNoHooks:
	if !cached {
		OkProjectTaskUpsertCacheMut.Lock()
		OkProjectTaskUpsertCache[key] = cache
		OkProjectTaskUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single OkProjectTask record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkProjectTask) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OkProjectTask record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OkProjectTask) DeleteG() error {
	if o == nil {
		return errors.New("models: no OkProjectTask provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OkProjectTask record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkProjectTask) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OkProjectTask record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OkProjectTask) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkProjectTask provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), OkProjectTaskPrimaryKeyMapping)
	sql := "DELETE FROM `OkProjectTask` WHERE `id_p_task`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from OkProjectTask")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q OkProjectTaskQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q OkProjectTaskQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no OkProjectTaskQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkProjectTask")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OkProjectTaskSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OkProjectTaskSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no OkProjectTask slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OkProjectTaskSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OkProjectTaskSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkProjectTask slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(OkProjectTaskBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProjectTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `OkProjectTask` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkProjectTaskPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkProjectTaskPrimaryKeyColumns), 1, len(OkProjectTaskPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkProjectTask slice")
	}

	if len(OkProjectTaskAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OkProjectTask) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OkProjectTask) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OkProjectTask) ReloadG() error {
	if o == nil {
		return errors.New("models: no OkProjectTask provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OkProjectTask) Reload(exec boil.Executor) error {
	ret, err := FindOkProjectTask(exec, o.IDPTask)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkProjectTaskSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkProjectTaskSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkProjectTaskSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OkProjectTaskSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkProjectTaskSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	OkProjectTasks := OkProjectTaskSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkProjectTaskPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `OkProjectTask`.* FROM `OkProjectTask` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkProjectTaskPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(OkProjectTaskPrimaryKeyColumns), 1, len(OkProjectTaskPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&OkProjectTasks)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OkProjectTaskSlice")
	}

	*o = OkProjectTasks

	return nil
}

// OkProjectTaskExists checks if the OkProjectTask row exists.
func OkProjectTaskExists(exec boil.Executor, idPTask int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `OkProjectTask` where `id_p_task`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, idPTask)
	}

	row := exec.QueryRow(sql, idPTask)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OkProjectTask exists")
	}

	return exists, nil
}

// OkProjectTaskExistsG checks if the OkProjectTask row exists.
func OkProjectTaskExistsG(idPTask int) (bool, error) {
	return OkProjectTaskExists(boil.GetDB(), idPTask)
}

// OkProjectTaskExistsGP checks if the OkProjectTask row exists. Panics on error.
func OkProjectTaskExistsGP(idPTask int) bool {
	e, err := OkProjectTaskExists(boil.GetDB(), idPTask)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OkProjectTaskExistsP checks if the OkProjectTask row exists. Panics on error.
func OkProjectTaskExistsP(exec boil.Executor, idPTask int) bool {
	e, err := OkProjectTaskExists(exec, idPTask)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
