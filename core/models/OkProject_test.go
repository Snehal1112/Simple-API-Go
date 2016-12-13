package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOkProjects(t *testing.T) {
	t.Parallel()

	query := OkProjects(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOkProjectsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProject.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkProjectsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProjects(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkProjectsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkProjectSlice{OkProject}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOkProjectsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OkProjectExists(tx, OkProject.IDProject)
	if err != nil {
		t.Errorf("Unable to check if OkProject exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OkProjectExistsG to return true, but got false.")
	}
}
func testOkProjectsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	OkProjectFound, err := FindOkProject(tx, OkProject.IDProject)
	if err != nil {
		t.Error(err)
	}

	if OkProjectFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOkProjectsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProjects(tx).Bind(OkProject); err != nil {
		t.Error(err)
	}
}

func testOkProjectsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := OkProjects(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOkProjectsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProjectOne := &OkProject{}
	OkProjectTwo := &OkProject{}
	if err = randomize.Struct(seed, OkProjectOne, OkProjectDBTypes, false, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}
	if err = randomize.Struct(seed, OkProjectTwo, OkProjectDBTypes, false, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkProjectTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkProjects(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOkProjectsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	OkProjectOne := &OkProject{}
	OkProjectTwo := &OkProject{}
	if err = randomize.Struct(seed, OkProjectOne, OkProjectDBTypes, false, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}
	if err = randomize.Struct(seed, OkProjectTwo, OkProjectDBTypes, false, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProjectOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkProjectTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func OkProjectBeforeInsertHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectAfterInsertHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectAfterSelectHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectBeforeUpdateHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectAfterUpdateHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectBeforeDeleteHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectAfterDeleteHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectBeforeUpsertHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func OkProjectAfterUpsertHook(e boil.Executor, o *OkProject) error {
	*o = OkProject{}
	return nil
}

func testOkProjectsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &OkProject{}
	o := &OkProject{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, OkProjectDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OkProject object: %s", err)
	}

	AddOkProjectHook(boil.BeforeInsertHook, OkProjectBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	OkProjectBeforeInsertHooks = []OkProjectHook{}

	AddOkProjectHook(boil.AfterInsertHook, OkProjectAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	OkProjectAfterInsertHooks = []OkProjectHook{}

	AddOkProjectHook(boil.AfterSelectHook, OkProjectAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	OkProjectAfterSelectHooks = []OkProjectHook{}

	AddOkProjectHook(boil.BeforeUpdateHook, OkProjectBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	OkProjectBeforeUpdateHooks = []OkProjectHook{}

	AddOkProjectHook(boil.AfterUpdateHook, OkProjectAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	OkProjectAfterUpdateHooks = []OkProjectHook{}

	AddOkProjectHook(boil.BeforeDeleteHook, OkProjectBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	OkProjectBeforeDeleteHooks = []OkProjectHook{}

	AddOkProjectHook(boil.AfterDeleteHook, OkProjectAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	OkProjectAfterDeleteHooks = []OkProjectHook{}

	AddOkProjectHook(boil.BeforeUpsertHook, OkProjectBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	OkProjectBeforeUpsertHooks = []OkProjectHook{}

	AddOkProjectHook(boil.AfterUpsertHook, OkProjectAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	OkProjectAfterUpsertHooks = []OkProjectHook{}
}
func testOkProjectsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkProjectsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx, OkProjectColumns...); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkProjectToManyProjectOkProjectTasks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProject
	var b, c OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)
	randomize.Struct(seed, &c, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)

	b.ProjectID = a.IDProject
	c.ProjectID = a.IDProject
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	OkProjectTask, err := a.ProjectOkProjectTasks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range OkProjectTask {
		if v.ProjectID == b.ProjectID {
			bFound = true
		}
		if v.ProjectID == c.ProjectID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OkProjectSlice{&a}
	if err = a.L.LoadProjectOkProjectTasks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ProjectOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ProjectOkProjectTasks = nil
	if err = a.L.LoadProjectOkProjectTasks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ProjectOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", OkProjectTask)
	}
}

func testOkProjectToManyAddOpProjectOkProjectTasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProject
	var b, c, d, e OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProjectDBTypes, false, strmangle.SetComplement(OkProjectPrimaryKeyColumns, OkProjectColumnsWithoutDefault)...); err != nil {
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
		err = a.AddProjectOkProjectTasks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.IDProject != first.ProjectID {
			t.Error("foreign key was wrong value", a.IDProject, first.ProjectID)
		}
		if a.IDProject != second.ProjectID {
			t.Error("foreign key was wrong value", a.IDProject, second.ProjectID)
		}

		if first.R.Project != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Project != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ProjectOkProjectTasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ProjectOkProjectTasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ProjectOkProjectTasks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOkProjectsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProject.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOkProjectsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkProjectSlice{OkProject}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOkProjectsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkProjects(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	OkProjectDBTypes = map[string]string{`Desc`: `text`, `GitURL`: `text`, `IDProject`: `int`, `NamaProject`: `varchar`}
	_                = bytes.MinRead
)

func testOkProjectsUpdate(t *testing.T) {
	t.Parallel()

	if len(OkProjectColumns) == len(OkProjectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	if err = OkProject.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOkProjectsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(OkProjectColumns) == len(OkProjectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkProject := &OkProject{}
	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkProject, OkProjectDBTypes, true, OkProjectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(OkProjectColumns, OkProjectPrimaryKeyColumns) {
		fields = OkProjectColumns
	} else {
		fields = strmangle.SetComplement(
			OkProjectColumns,
			OkProjectPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(OkProject))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OkProjectSlice{OkProject}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOkProjectsUpsert(t *testing.T) {
	t.Parallel()

	if len(OkProjectColumns) == len(OkProjectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	OkProject := OkProject{}
	if err = randomize.Struct(seed, &OkProject, OkProjectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProject.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkProject: %s", err)
	}

	count, err := OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &OkProject, OkProjectDBTypes, false, OkProjectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkProject struct: %s", err)
	}

	if err = OkProject.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkProject: %s", err)
	}

	count, err = OkProjects(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
