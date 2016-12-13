package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOkTasks(t *testing.T) {
	t.Parallel()

	query := OkTasks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOkTasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkTask.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkTasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkTasks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkTasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkTaskSlice{OkTask}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOkTasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OkTaskExists(tx, OkTask.IDTask)
	if err != nil {
		t.Errorf("Unable to check if OkTask exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OkTaskExistsG to return true, but got false.")
	}
}
func testOkTasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	OkTaskFound, err := FindOkTask(tx, OkTask.IDTask)
	if err != nil {
		t.Error(err)
	}

	if OkTaskFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOkTasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkTasks(tx).Bind(OkTask); err != nil {
		t.Error(err)
	}
}

func testOkTasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := OkTasks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOkTasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTaskOne := &OkTask{}
	OkTaskTwo := &OkTask{}
	if err = randomize.Struct(seed, OkTaskOne, OkTaskDBTypes, false, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}
	if err = randomize.Struct(seed, OkTaskTwo, OkTaskDBTypes, false, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTaskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkTaskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkTasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOkTasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	OkTaskOne := &OkTask{}
	OkTaskTwo := &OkTask{}
	if err = randomize.Struct(seed, OkTaskOne, OkTaskDBTypes, false, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}
	if err = randomize.Struct(seed, OkTaskTwo, OkTaskDBTypes, false, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTaskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkTaskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func OkTaskBeforeInsertHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskAfterInsertHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskAfterSelectHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskBeforeUpdateHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskAfterUpdateHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskBeforeDeleteHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskAfterDeleteHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskBeforeUpsertHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func OkTaskAfterUpsertHook(e boil.Executor, o *OkTask) error {
	*o = OkTask{}
	return nil
}

func testOkTasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &OkTask{}
	o := &OkTask{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, OkTaskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OkTask object: %s", err)
	}

	AddOkTaskHook(boil.BeforeInsertHook, OkTaskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	OkTaskBeforeInsertHooks = []OkTaskHook{}

	AddOkTaskHook(boil.AfterInsertHook, OkTaskAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	OkTaskAfterInsertHooks = []OkTaskHook{}

	AddOkTaskHook(boil.AfterSelectHook, OkTaskAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	OkTaskAfterSelectHooks = []OkTaskHook{}

	AddOkTaskHook(boil.BeforeUpdateHook, OkTaskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	OkTaskBeforeUpdateHooks = []OkTaskHook{}

	AddOkTaskHook(boil.AfterUpdateHook, OkTaskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	OkTaskAfterUpdateHooks = []OkTaskHook{}

	AddOkTaskHook(boil.BeforeDeleteHook, OkTaskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	OkTaskBeforeDeleteHooks = []OkTaskHook{}

	AddOkTaskHook(boil.AfterDeleteHook, OkTaskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	OkTaskAfterDeleteHooks = []OkTaskHook{}

	AddOkTaskHook(boil.BeforeUpsertHook, OkTaskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	OkTaskBeforeUpsertHooks = []OkTaskHook{}

	AddOkTaskHook(boil.AfterUpsertHook, OkTaskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	OkTaskAfterUpsertHooks = []OkTaskHook{}
}
func testOkTasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkTasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx, OkTaskColumns...); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkTaskToManyTaskOkProjectTasks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkTask
	var b, c OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)
	randomize.Struct(seed, &c, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)

	b.TaskID = a.IDTask
	c.TaskID = a.IDTask
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	OkProjectTask, err := a.TaskOkProjectTasks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range OkProjectTask {
		if v.TaskID == b.TaskID {
			bFound = true
		}
		if v.TaskID == c.TaskID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OkTaskSlice{&a}
	if err = a.L.LoadTaskOkProjectTasks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TaskOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TaskOkProjectTasks = nil
	if err = a.L.LoadTaskOkProjectTasks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TaskOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", OkProjectTask)
	}
}

func testOkTaskToManyAddOpTaskOkProjectTasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkTask
	var b, c, d, e OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkTaskDBTypes, false, strmangle.SetComplement(OkTaskPrimaryKeyColumns, OkTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OkProjectTask{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, OkProjectTaskDBTypes, false, strmangle.SetComplement(OkProjectTaskPrimaryKeyColumns, OkProjectTaskColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*OkProjectTask{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTaskOkProjectTasks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.IDTask != first.TaskID {
			t.Error("foreign key was wrong value", a.IDTask, first.TaskID)
		}
		if a.IDTask != second.TaskID {
			t.Error("foreign key was wrong value", a.IDTask, second.TaskID)
		}

		if first.R.Task != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Task != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TaskOkProjectTasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TaskOkProjectTasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TaskOkProjectTasks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOkTasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkTask.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOkTasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkTaskSlice{OkTask}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOkTasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkTasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	OkTaskDBTypes = map[string]string{`DescTask`: `longtext`, `Difficulty`: `varchar`, `IDTask`: `int`, `TitleTask`: `varchar`, `Urgency`: `varchar`}
	_             = bytes.MinRead
)

func testOkTasksUpdate(t *testing.T) {
	t.Parallel()

	if len(OkTaskColumns) == len(OkTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	if err = OkTask.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOkTasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(OkTaskColumns) == len(OkTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkTask := &OkTask{}
	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkTask, OkTaskDBTypes, true, OkTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(OkTaskColumns, OkTaskPrimaryKeyColumns) {
		fields = OkTaskColumns
	} else {
		fields = strmangle.SetComplement(
			OkTaskColumns,
			OkTaskPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(OkTask))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OkTaskSlice{OkTask}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOkTasksUpsert(t *testing.T) {
	t.Parallel()

	if len(OkTaskColumns) == len(OkTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	OkTask := OkTask{}
	if err = randomize.Struct(seed, &OkTask, OkTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkTask.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkTask: %s", err)
	}

	count, err := OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &OkTask, OkTaskDBTypes, false, OkTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	if err = OkTask.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkTask: %s", err)
	}

	count, err = OkTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
