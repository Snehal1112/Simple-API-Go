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

// OkAkun is an object representing the database table.
type OkAkun struct {
	IDAkun    int       `boil:"id_akun" json:"id_akun" toml:"id_akun" yaml:"id_akun"`
	Username  string    `boil:"username" json:"username" toml:"username" yaml:"username"`
	Password  string    `boil:"password" json:"password" toml:"password" yaml:"password"`
	Token     string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Role      string    `boil:"role" json:"role" toml:"role" yaml:"role"`

	R *OkAkunR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L OkAkunL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// OkAkunR is where relationships are stored.
type OkAkunR struct {
	AkunOkProfiles OkProfileSlice
}

// OkAkunL is where Load methods for each relationship are stored.
type OkAkunL struct{}

var (
	OkAkunColumns               = []string{"id_akun", "username", "password", "token", "created_at", "updated_at", "role"}
	OkAkunColumnsWithoutDefault = []string{"username", "password", "token", "created_at", "updated_at", "role"}
	OkAkunColumnsWithDefault    = []string{"id_akun"}
	OkAkunPrimaryKeyColumns     = []string{"id_akun"}
)

type (
	// OkAkunSlice is an alias for a slice of pointers to OkAkun.
	// This should generally be used opposed to []OkAkun.
	OkAkunSlice []*OkAkun
	// OkAkunHook is the signature for custom OkAkun hook methods
	OkAkunHook func(boil.Executor, *OkAkun) error

	OkAkunQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	OkAkunType                 = reflect.TypeOf(&OkAkun{})
	OkAkunMapping              = queries.MakeStructMapping(OkAkunType)
	OkAkunPrimaryKeyMapping, _ = queries.BindMapping(OkAkunType, OkAkunMapping, OkAkunPrimaryKeyColumns)
	OkAkunInsertCacheMut       sync.RWMutex
	OkAkunInsertCache          = make(map[string]insertCache)
	OkAkunUpdateCacheMut       sync.RWMutex
	OkAkunUpdateCache          = make(map[string]updateCache)
	OkAkunUpsertCacheMut       sync.RWMutex
	OkAkunUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var OkAkunBeforeInsertHooks []OkAkunHook
var OkAkunBeforeUpdateHooks []OkAkunHook
var OkAkunBeforeDeleteHooks []OkAkunHook
var OkAkunBeforeUpsertHooks []OkAkunHook

var OkAkunAfterInsertHooks []OkAkunHook
var OkAkunAfterSelectHooks []OkAkunHook
var OkAkunAfterUpdateHooks []OkAkunHook
var OkAkunAfterDeleteHooks []OkAkunHook
var OkAkunAfterUpsertHooks []OkAkunHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OkAkun) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OkAkun) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OkAkun) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OkAkun) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OkAkun) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OkAkun) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OkAkun) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OkAkun) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OkAkun) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range OkAkunAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOkAkunHook registers your hook function for all future operations.
func AddOkAkunHook(hookPoint boil.HookPoint, OkAkunHook OkAkunHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		OkAkunBeforeInsertHooks = append(OkAkunBeforeInsertHooks, OkAkunHook)
	case boil.BeforeUpdateHook:
		OkAkunBeforeUpdateHooks = append(OkAkunBeforeUpdateHooks, OkAkunHook)
	case boil.BeforeDeleteHook:
		OkAkunBeforeDeleteHooks = append(OkAkunBeforeDeleteHooks, OkAkunHook)
	case boil.BeforeUpsertHook:
		OkAkunBeforeUpsertHooks = append(OkAkunBeforeUpsertHooks, OkAkunHook)
	case boil.AfterInsertHook:
		OkAkunAfterInsertHooks = append(OkAkunAfterInsertHooks, OkAkunHook)
	case boil.AfterSelectHook:
		OkAkunAfterSelectHooks = append(OkAkunAfterSelectHooks, OkAkunHook)
	case boil.AfterUpdateHook:
		OkAkunAfterUpdateHooks = append(OkAkunAfterUpdateHooks, OkAkunHook)
	case boil.AfterDeleteHook:
		OkAkunAfterDeleteHooks = append(OkAkunAfterDeleteHooks, OkAkunHook)
	case boil.AfterUpsertHook:
		OkAkunAfterUpsertHooks = append(OkAkunAfterUpsertHooks, OkAkunHook)
	}
}

// OneP returns a single OkAkun record from the query, and panics on error.
func (q OkAkunQuery) OneP() *OkAkun {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single OkAkun record from the query.
func (q OkAkunQuery) One() (*OkAkun, error) {
	o := &OkAkun{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OkAkun")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all OkAkun records from the query, and panics on error.
func (q OkAkunQuery) AllP() OkAkunSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all OkAkun records from the query.
func (q OkAkunQuery) All() (OkAkunSlice, error) {
	var o OkAkunSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OkAkun slice")
	}

	if len(OkAkunAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all OkAkun records in the query, and panics on error.
func (q OkAkunQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all OkAkun records in the query.
func (q OkAkunQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OkAkun rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q OkAkunQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q OkAkunQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OkAkun exists")
	}

	return count > 0, nil
}

// AkunOkProfilesG retrieves all the OkProfile's OkProfile via akun_id column.
func (o *OkAkun) AkunOkProfilesG(mods ...qm.QueryMod) OkProfileQuery {
	return o.AkunOkProfiles(boil.GetDB(), mods...)
}

// AkunOkProfiles retrieves all the OkProfile's OkProfile with an executor via akun_id column.
func (o *OkAkun) AkunOkProfiles(exec boil.Executor, mods ...qm.QueryMod) OkProfileQuery {
	queryMods := []qm.QueryMod{
		qm.Select("`a`.*"),
	}

	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`a`.`akun_id`=?", o.IDAkun),
	)

	query := OkProfiles(exec, queryMods...)
	queries.SetFrom(query.Query, "`OkProfile` as `a`")
	return query
}

// LoadAkunOkProfiles allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (OkAkunL) LoadAkunOkProfiles(e boil.Executor, singular bool, maybeOkAkun interface{}) error {
	var slice []*OkAkun
	var object *OkAkun

	count := 1
	if singular {
		object = maybeOkAkun.(*OkAkun)
	} else {
		slice = *maybeOkAkun.(*OkAkunSlice)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &OkAkunR{}
		}
		args[0] = object.IDAkun
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &OkAkunR{}
			}
			args[i] = obj.IDAkun
		}
	}

	query := fmt.Sprintf(
		"select * from `OkProfile` where `akun_id` in (%s)",
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

	if len(OkProfileAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.AkunOkProfiles = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.IDAkun == foreign.AkunID {
				local.R.AkunOkProfiles = append(local.R.AkunOkProfiles, foreign)
				break
			}
		}
	}

	return nil
}

// AddAkunOkProfiles adds the given related objects to the existing relationships
// of the OkAkun, optionally inserting them as new records.
// Appends related to o.R.AkunOkProfiles.
// Sets related.R.Akun appropriately.
func (o *OkAkun) AddAkunOkProfiles(exec boil.Executor, insert bool, related ...*OkProfile) error {
	var err error
	for _, rel := range related {
		rel.AkunID = o.IDAkun
		if insert {
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			if err = rel.Update(exec, "akun_id"); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}
		}
	}

	if o.R == nil {
		o.R = &OkAkunR{
			AkunOkProfiles: related,
		}
	} else {
		o.R.AkunOkProfiles = append(o.R.AkunOkProfiles, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &OkProfileR{
				Akun: o,
			}
		} else {
			rel.R.Akun = o
		}
	}
	return nil
}

// OkAkunsG retrieves all records.
func OkAkunsG(mods ...qm.QueryMod) OkAkunQuery {
	return OkAkuns(boil.GetDB(), mods...)
}

// OkAkuns retrieves all the records using an executor.
func OkAkuns(exec boil.Executor, mods ...qm.QueryMod) OkAkunQuery {
	mods = append(mods, qm.From("`OkAkun`"))
	return OkAkunQuery{NewQuery(exec, mods...)}
}

// FindOkAkunG retrieves a single record by ID.
func FindOkAkunG(idAkun int, selectCols ...string) (*OkAkun, error) {
	return FindOkAkun(boil.GetDB(), idAkun, selectCols...)
}

// FindOkAkunGP retrieves a single record by ID, and panics on error.
func FindOkAkunGP(idAkun int, selectCols ...string) *OkAkun {
	retobj, err := FindOkAkun(boil.GetDB(), idAkun, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindOkAkun retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOkAkun(exec boil.Executor, idAkun int, selectCols ...string) (*OkAkun, error) {
	OkAkunObj := &OkAkun{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OkAkun` where `id_akun`=?", sel,
	)

	q := queries.Raw(exec, query, idAkun)

	err := q.Bind(OkAkunObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OkAkun")
	}

	return OkAkunObj, nil
}

// FindOkAkunP retrieves a single record by ID with an executor, and panics on error.
func FindOkAkunP(exec boil.Executor, idAkun int, selectCols ...string) *OkAkun {
	retobj, err := FindOkAkun(exec, idAkun, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OkAkun) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *OkAkun) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *OkAkun) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *OkAkun) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkAkun provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(OkAkunColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	OkAkunInsertCacheMut.RLock()
	cache, cached := OkAkunInsertCache[key]
	OkAkunInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			OkAkunColumns,
			OkAkunColumnsWithDefault,
			OkAkunColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(OkAkunType, OkAkunMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(OkAkunType, OkAkunMapping, returnColumns)
		if err != nil {
			return err
		}
		cache.query = fmt.Sprintf("INSERT INTO `OkAkun` (`%s`) VALUES (%s)", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OkAkun` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, OkAkunPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into OkAkun")
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

	o.IDAkun = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkAkunMapping["IDAkun"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDAkun,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkAkun")
	}

CacheNoHooks:
	if !cached {
		OkAkunInsertCacheMut.Lock()
		OkAkunInsertCache[key] = cache
		OkAkunInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single OkAkun record. See Update for
// whitelist behavior description.
func (o *OkAkun) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single OkAkun record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *OkAkun) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the OkAkun, and panics on error.
// See Update for whitelist behavior description.
func (o *OkAkun) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the OkAkun.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *OkAkun) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	OkAkunUpdateCacheMut.RLock()
	cache, cached := OkAkunUpdateCache[key]
	OkAkunUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(OkAkunColumns, OkAkunPrimaryKeyColumns, whitelist)
		if len(wl) == 0 {
			return errors.New("models: unable to update OkAkun, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OkAkun` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, OkAkunPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(OkAkunType, OkAkunMapping, append(wl, OkAkunPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update OkAkun row")
	}

	if !cached {
		OkAkunUpdateCacheMut.Lock()
		OkAkunUpdateCache[key] = cache
		OkAkunUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q OkAkunQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q OkAkunQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for OkAkun")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OkAkunSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o OkAkunSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o OkAkunSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OkAkunSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkAkunPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"UPDATE `OkAkun` SET %s WHERE (`id_akun`) IN (%s)",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkAkunPrimaryKeyColumns), len(colNames)+1, len(OkAkunPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in OkAkun slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *OkAkun) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *OkAkun) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *OkAkun) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *OkAkun) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no OkAkun provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(OkAkunColumnsWithDefault, o)

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

	OkAkunUpsertCacheMut.RLock()
	cache, cached := OkAkunUpsertCache[key]
	OkAkunUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		var ret []string
		whitelist, ret = strmangle.InsertColumnSet(
			OkAkunColumns,
			OkAkunColumnsWithDefault,
			OkAkunColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)
		update := strmangle.UpdateColumnSet(
			OkAkunColumns,
			OkAkunPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert OkAkun, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "OkAkun", update, whitelist)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OkAkun` WHERE `id_akun`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(OkAkunType, OkAkunMapping, whitelist)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(OkAkunType, OkAkunMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for OkAkun")
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

	o.IDAkun = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == OkAkunMapping["IDAkun"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.IDAkun,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OkAkun")
	}

CacheNoHooks:
	if !cached {
		OkAkunUpsertCacheMut.Lock()
		OkAkunUpsertCache[key] = cache
		OkAkunUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single OkAkun record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkAkun) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single OkAkun record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OkAkun) DeleteG() error {
	if o == nil {
		return errors.New("models: no OkAkun provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single OkAkun record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *OkAkun) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single OkAkun record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OkAkun) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkAkun provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), OkAkunPrimaryKeyMapping)
	sql := "DELETE FROM `OkAkun` WHERE `id_akun`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from OkAkun")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q OkAkunQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q OkAkunQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no OkAkunQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkAkun")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o OkAkunSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o OkAkunSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no OkAkun slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o OkAkunSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OkAkunSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no OkAkun slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(OkAkunBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkAkunPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"DELETE FROM `OkAkun` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkAkunPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(o)*len(OkAkunPrimaryKeyColumns), 1, len(OkAkunPrimaryKeyColumns)),
	)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from OkAkun slice")
	}

	if len(OkAkunAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *OkAkun) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *OkAkun) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OkAkun) ReloadG() error {
	if o == nil {
		return errors.New("models: no OkAkun provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OkAkun) Reload(exec boil.Executor) error {
	ret, err := FindOkAkun(exec, o.IDAkun)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkAkunSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *OkAkunSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkAkunSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty OkAkunSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OkAkunSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	OkAkuns := OkAkunSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), OkAkunPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf(
		"SELECT `OkAkun`.* FROM `OkAkun` WHERE (%s) IN (%s)",
		strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, OkAkunPrimaryKeyColumns), ","),
		strmangle.Placeholders(dialect.IndexPlaceholders, len(*o)*len(OkAkunPrimaryKeyColumns), 1, len(OkAkunPrimaryKeyColumns)),
	)

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&OkAkuns)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OkAkunSlice")
	}

	*o = OkAkuns

	return nil
}

// OkAkunExists checks if the OkAkun row exists.
func OkAkunExists(exec boil.Executor, idAkun int) (bool, error) {
	var exists bool

	sql := "select exists(select 1 from `OkAkun` where `id_akun`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, idAkun)
	}

	row := exec.QueryRow(sql, idAkun)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OkAkun exists")
	}

	return exists, nil
}

// OkAkunExistsG checks if the OkAkun row exists.
func OkAkunExistsG(idAkun int) (bool, error) {
	return OkAkunExists(boil.GetDB(), idAkun)
}

// OkAkunExistsGP checks if the OkAkun row exists. Panics on error.
func OkAkunExistsGP(idAkun int) bool {
	e, err := OkAkunExists(boil.GetDB(), idAkun)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// OkAkunExistsP checks if the OkAkun row exists. Panics on error.
func OkAkunExistsP(exec boil.Executor, idAkun int) bool {
	e, err := OkAkunExists(exec, idAkun)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
