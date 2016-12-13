package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOkProjectTasks(t *testing.T) {
	t.Parallel()

	query := OkProjectTasks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOkProjectTasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProjectTask.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkProjectTasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProjectTasks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkProjectTasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkProjectTaskSlice{OkProjectTask}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOkProjectTasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OkProjectTaskExists(tx, OkProjectTask.IDPTask)
	if err != nil {
		t.Errorf("Unable to check if OkProjectTask exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OkProjectTaskExistsG to return true, but got false.")
	}
}
func testOkProjectTasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	OkProjectTaskFound, err := FindOkProjectTask(tx, OkProjectTask.IDPTask)
	if err != nil {
		t.Error(err)
	}

	if OkProjectTaskFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOkProjectTasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProjectTasks(tx).Bind(OkProjectTask); err != nil {
		t.Error(err)
	}
}

func testOkProjectTasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := OkProjectTasks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOkProjectTasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTaskOne := &OkProjectTask{}
	OkProjectTaskTwo := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTaskOne, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}
	if err = randomize.Struct(seed, OkProjectTaskTwo, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTaskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkProjectTaskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkProjectTasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOkProjectTasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	OkProjectTaskOne := &OkProjectTask{}
	OkProjectTaskTwo := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTaskOne, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}
	if err = randomize.Struct(seed, OkProjectTaskTwo, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTaskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkProjectTaskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func OkProjectTaskBeforeInsertHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskAfterInsertHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskAfterSelectHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskBeforeUpdateHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskAfterUpdateHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskBeforeDeleteHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskAfterDeleteHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskBeforeUpsertHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func OkProjectTaskAfterUpsertHook(e boil.Executor, o *OkProjectTask) error {
	*o = OkProjectTask{}
	return nil
}

func testOkProjectTasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &OkProjectTask{}
	o := &OkProjectTask{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, OkProjectTaskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OkProjectTask object: %s", err)
	}

	AddOkProjectTaskHook(boil.BeforeInsertHook, OkProjectTaskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskBeforeInsertHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.AfterInsertHook, OkProjectTaskAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskAfterInsertHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.AfterSelectHook, OkProjectTaskAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskAfterSelectHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.BeforeUpdateHook, OkProjectTaskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskBeforeUpdateHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.AfterUpdateHook, OkProjectTaskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskAfterUpdateHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.BeforeDeleteHook, OkProjectTaskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskBeforeDeleteHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.AfterDeleteHook, OkProjectTaskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskAfterDeleteHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.BeforeUpsertHook, OkProjectTaskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskBeforeUpsertHooks = []OkProjectTaskHook{}

	AddOkProjectTaskHook(boil.AfterUpsertHook, OkProjectTaskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	OkProjectTaskAfterUpsertHooks = []OkProjectTaskHook{}
}
func testOkProjectTasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkProjectTasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx, OkProjectTaskColumns...); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkProjectTaskToOneOkProfileUsingAssignedto(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OkProjectTask
	var foreign OkProfile

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AssignedtoID = foreign.IDProfile
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Assignedto(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.IDProfile != foreign.IDProfile {
		t.Errorf("want: %v, got %v", foreign.IDProfile, check.IDProfile)
	}

	slice := OkProjectTaskSlice{&local}
	if err = local.L.LoadAssignedto(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Assignedto == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Assignedto = nil
	if err = local.L.LoadAssignedto(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Assignedto == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOkProjectTaskToOneOkProfileUsingAssigner(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OkProjectTask
	var foreign OkProfile

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AssignerID = foreign.IDProfile
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Assigner(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.IDProfile != foreign.IDProfile {
		t.Errorf("want: %v, got %v", foreign.IDProfile, check.IDProfile)
	}

	slice := OkProjectTaskSlice{&local}
	if err = local.L.LoadAssigner(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Assigner == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Assigner = nil
	if err = local.L.LoadAssigner(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Assigner == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOkProjectTaskToOneOkProjectUsingProject(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OkProjectTask
	var foreign OkProject

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ProjectID = foreign.IDProject
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Project(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.IDProject != foreign.IDProject {
		t.Errorf("want: %v, got %v", foreign.IDProject, check.IDProject)
	}

	slice := OkProjectTaskSlice{&local}
	if err = local.L.LoadProject(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Project == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Project = nil
	if err = local.L.LoadProject(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Project == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOkProjectTaskToOneOkTaskUsingTask(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OkProjectTask
	var foreign OkTask

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, OkTaskDBTypes, true, OkTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkTask struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TaskID = foreign.IDTask
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Task(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.IDTask != foreign.IDTask {
		t.Errorf("want: %v, got %v", foreign.IDTask, check.IDTask)
	}

	slice := OkProjectTaskSlice{&local}
	if err = local.L.LoadTask(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Task == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Task = nil
	if err = local.L.LoadTask(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Task == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOkProjectTaskToOneSetOpOkProfileUsingAssignedto(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProjectTask
	var b, c OkProfile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProjectTaskDBTypes, false, strmangle.SetComplement(OkProjectTaskPrimaryKeyColumns, OkProjectTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OkProfile{&b, &c} {
		err = a.SetAssignedto(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Assignedto != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AssignedtoOkProjectTasks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AssignedtoID != x.IDProfile {
			t.Error("foreign key was wrong value", a.AssignedtoID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AssignedtoID))
		reflect.Indirect(reflect.ValueOf(&a.AssignedtoID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AssignedtoID != x.IDProfile {
			t.Error("foreign key was wrong value", a.AssignedtoID, x.IDProfile)
		}
	}
}
func testOkProjectTaskToOneSetOpOkProfileUsingAssigner(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProjectTask
	var b, c OkProfile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProjectTaskDBTypes, false, strmangle.SetComplement(OkProjectTaskPrimaryKeyColumns, OkProjectTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OkProfile{&b, &c} {
		err = a.SetAssigner(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Assigner != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AssignerOkProjectTasks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AssignerID != x.IDProfile {
			t.Error("foreign key was wrong value", a.AssignerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AssignerID))
		reflect.Indirect(reflect.ValueOf(&a.AssignerID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AssignerID != x.IDProfile {
			t.Error("foreign key was wrong value", a.AssignerID, x.IDProfile)
		}
	}
}
func testOkProjectTaskToOneSetOpOkProjectUsingProject(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProjectTask
	var b, c OkProject

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProjectTaskDBTypes, false, strmangle.SetComplement(OkProjectTaskPrimaryKeyColumns, OkProjectTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, OkProjectDBTypes, false, strmangle.SetComplement(OkProjectPrimaryKeyColumns, OkProjectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, OkProjectDBTypes, false, strmangle.SetComplement(OkProjectPrimaryKeyColumns, OkProjectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OkProject{&b, &c} {
		err = a.SetProject(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Project != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ProjectOkProjectTasks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProjectID != x.IDProject {
			t.Error("foreign key was wrong value", a.ProjectID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProjectID))
		reflect.Indirect(reflect.ValueOf(&a.ProjectID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ProjectID != x.IDProject {
			t.Error("foreign key was wrong value", a.ProjectID, x.IDProject)
		}
	}
}
func testOkProjectTaskToOneSetOpOkTaskUsingTask(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProjectTask
	var b, c OkTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProjectTaskDBTypes, false, strmangle.SetComplement(OkProjectTaskPrimaryKeyColumns, OkProjectTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, OkTaskDBTypes, false, strmangle.SetComplement(OkTaskPrimaryKeyColumns, OkTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, OkTaskDBTypes, false, strmangle.SetComplement(OkTaskPrimaryKeyColumns, OkTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OkTask{&b, &c} {
		err = a.SetTask(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Task != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TaskOkProjectTasks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TaskID != x.IDTask {
			t.Error("foreign key was wrong value", a.TaskID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TaskID))
		reflect.Indirect(reflect.ValueOf(&a.TaskID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TaskID != x.IDTask {
			t.Error("foreign key was wrong value", a.TaskID, x.IDTask)
		}
	}
}
func testOkProjectTasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProjectTask.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOkProjectTasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkProjectTaskSlice{OkProjectTask}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOkProjectTasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkProjectTasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	OkProjectTaskDBTypes = map[string]string{`AssignedtoID`: `int`, `AssignerID`: `int`, `CreatedAt`: `datetime`, `DueDate`: `datetime`, `IDPTask`: `int`, `ProjectID`: `int`, `Status`: `varchar`, `TaskID`: `int`, `UpdatedAt`: `datetime`}
	_                    = bytes.MinRead
)

func testOkProjectTasksUpdate(t *testing.T) {
	t.Parallel()

	if len(OkProjectTaskColumns) == len(OkProjectTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	if err = OkProjectTask.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOkProjectTasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(OkProjectTaskColumns) == len(OkProjectTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkProjectTask := &OkProjectTask{}
	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkProjectTask, OkProjectTaskDBTypes, true, OkProjectTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(OkProjectTaskColumns, OkProjectTaskPrimaryKeyColumns) {
		fields = OkProjectTaskColumns
	} else {
		fields = strmangle.SetComplement(
			OkProjectTaskColumns,
			OkProjectTaskPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(OkProjectTask))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OkProjectTaskSlice{OkProjectTask}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOkProjectTasksUpsert(t *testing.T) {
	t.Parallel()

	if len(OkProjectTaskColumns) == len(OkProjectTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	OkProjectTask := OkProjectTask{}
	if err = randomize.Struct(seed, &OkProjectTask, OkProjectTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectTask.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkProjectTask: %s", err)
	}

	count, err := OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &OkProjectTask, OkProjectTaskDBTypes, false, OkProjectTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkProjectTask struct: %s", err)
	}

	if err = OkProjectTask.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkProjectTask: %s", err)
	}

	count, err = OkProjectTasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
