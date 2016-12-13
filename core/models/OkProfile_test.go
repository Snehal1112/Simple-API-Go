package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOkProfiles(t *testing.T) {
	t.Parallel()

	query := OkProfiles(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOkProfilesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProfile.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkProfilesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProfiles(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkProfilesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkProfileSlice{OkProfile}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOkProfilesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OkProfileExists(tx, OkProfile.IDProfile)
	if err != nil {
		t.Errorf("Unable to check if OkProfile exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OkProfileExistsG to return true, but got false.")
	}
}
func testOkProfilesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	OkProfileFound, err := FindOkProfile(tx, OkProfile.IDProfile)
	if err != nil {
		t.Error(err)
	}

	if OkProfileFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOkProfilesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProfiles(tx).Bind(OkProfile); err != nil {
		t.Error(err)
	}
}

func testOkProfilesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := OkProfiles(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOkProfilesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfileOne := &OkProfile{}
	OkProfileTwo := &OkProfile{}
	if err = randomize.Struct(seed, OkProfileOne, OkProfileDBTypes, false, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}
	if err = randomize.Struct(seed, OkProfileTwo, OkProfileDBTypes, false, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfileOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkProfileTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkProfiles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOkProfilesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	OkProfileOne := &OkProfile{}
	OkProfileTwo := &OkProfile{}
	if err = randomize.Struct(seed, OkProfileOne, OkProfileDBTypes, false, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}
	if err = randomize.Struct(seed, OkProfileTwo, OkProfileDBTypes, false, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfileOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkProfileTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func OkProfileBeforeInsertHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileAfterInsertHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileAfterSelectHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileBeforeUpdateHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileAfterUpdateHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileBeforeDeleteHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileAfterDeleteHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileBeforeUpsertHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func OkProfileAfterUpsertHook(e boil.Executor, o *OkProfile) error {
	*o = OkProfile{}
	return nil
}

func testOkProfilesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &OkProfile{}
	o := &OkProfile{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, OkProfileDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OkProfile object: %s", err)
	}

	AddOkProfileHook(boil.BeforeInsertHook, OkProfileBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	OkProfileBeforeInsertHooks = []OkProfileHook{}

	AddOkProfileHook(boil.AfterInsertHook, OkProfileAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	OkProfileAfterInsertHooks = []OkProfileHook{}

	AddOkProfileHook(boil.AfterSelectHook, OkProfileAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	OkProfileAfterSelectHooks = []OkProfileHook{}

	AddOkProfileHook(boil.BeforeUpdateHook, OkProfileBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	OkProfileBeforeUpdateHooks = []OkProfileHook{}

	AddOkProfileHook(boil.AfterUpdateHook, OkProfileAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	OkProfileAfterUpdateHooks = []OkProfileHook{}

	AddOkProfileHook(boil.BeforeDeleteHook, OkProfileBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	OkProfileBeforeDeleteHooks = []OkProfileHook{}

	AddOkProfileHook(boil.AfterDeleteHook, OkProfileAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	OkProfileAfterDeleteHooks = []OkProfileHook{}

	AddOkProfileHook(boil.BeforeUpsertHook, OkProfileBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	OkProfileBeforeUpsertHooks = []OkProfileHook{}

	AddOkProfileHook(boil.AfterUpsertHook, OkProfileAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	OkProfileAfterUpsertHooks = []OkProfileHook{}
}
func testOkProfilesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkProfilesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx, OkProfileColumns...); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkProfileToManyAssignedtoOkProjectTasks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProfile
	var b, c OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)
	randomize.Struct(seed, &c, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)

	b.AssignedtoID = a.IDProfile
	c.AssignedtoID = a.IDProfile
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	OkProjectTask, err := a.AssignedtoOkProjectTasks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range OkProjectTask {
		if v.AssignedtoID == b.AssignedtoID {
			bFound = true
		}
		if v.AssignedtoID == c.AssignedtoID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OkProfileSlice{&a}
	if err = a.L.LoadAssignedtoOkProjectTasks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssignedtoOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AssignedtoOkProjectTasks = nil
	if err = a.L.LoadAssignedtoOkProjectTasks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssignedtoOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", OkProjectTask)
	}
}

func testOkProfileToManyAssignerOkProjectTasks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProfile
	var b, c OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)
	randomize.Struct(seed, &c, OkProjectTaskDBTypes, false, OkProjectTaskColumnsWithDefault...)

	b.AssignerID = a.IDProfile
	c.AssignerID = a.IDProfile
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	OkProjectTask, err := a.AssignerOkProjectTasks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range OkProjectTask {
		if v.AssignerID == b.AssignerID {
			bFound = true
		}
		if v.AssignerID == c.AssignerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OkProfileSlice{&a}
	if err = a.L.LoadAssignerOkProjectTasks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssignerOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AssignerOkProjectTasks = nil
	if err = a.L.LoadAssignerOkProjectTasks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AssignerOkProjectTasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", OkProjectTask)
	}
}

func testOkProfileToManyAddOpAssignedtoOkProjectTasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProfile
	var b, c, d, e OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
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
		err = a.AddAssignedtoOkProjectTasks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.IDProfile != first.AssignedtoID {
			t.Error("foreign key was wrong value", a.IDProfile, first.AssignedtoID)
		}
		if a.IDProfile != second.AssignedtoID {
			t.Error("foreign key was wrong value", a.IDProfile, second.AssignedtoID)
		}

		if first.R.Assignedto != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Assignedto != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AssignedtoOkProjectTasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AssignedtoOkProjectTasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AssignedtoOkProjectTasks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOkProfileToManyAddOpAssignerOkProjectTasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProfile
	var b, c, d, e OkProjectTask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
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
		err = a.AddAssignerOkProjectTasks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.IDProfile != first.AssignerID {
			t.Error("foreign key was wrong value", a.IDProfile, first.AssignerID)
		}
		if a.IDProfile != second.AssignerID {
			t.Error("foreign key was wrong value", a.IDProfile, second.AssignerID)
		}

		if first.R.Assigner != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Assigner != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AssignerOkProjectTasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AssignerOkProjectTasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AssignerOkProjectTasks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOkProfileToOneOkAkunUsingAkun(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local OkProfile
	var foreign OkAkun

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AkunID = foreign.IDAkun
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Akun(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.IDAkun != foreign.IDAkun {
		t.Errorf("want: %v, got %v", foreign.IDAkun, check.IDAkun)
	}

	slice := OkProfileSlice{&local}
	if err = local.L.LoadAkun(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Akun == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Akun = nil
	if err = local.L.LoadAkun(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Akun == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOkProfileToOneSetOpOkAkunUsingAkun(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkProfile
	var b, c OkAkun

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, OkAkunDBTypes, false, strmangle.SetComplement(OkAkunPrimaryKeyColumns, OkAkunColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, OkAkunDBTypes, false, strmangle.SetComplement(OkAkunPrimaryKeyColumns, OkAkunColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OkAkun{&b, &c} {
		err = a.SetAkun(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Akun != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.AkunOkProfiles[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AkunID != x.IDAkun {
			t.Error("foreign key was wrong value", a.AkunID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AkunID))
		reflect.Indirect(reflect.ValueOf(&a.AkunID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AkunID != x.IDAkun {
			t.Error("foreign key was wrong value", a.AkunID, x.IDAkun)
		}
	}
}
func testOkProfilesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkProfile.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOkProfilesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkProfileSlice{OkProfile}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOkProfilesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkProfiles(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	OkProfileDBTypes = map[string]string{`AkunID`: `int`, `CreatedAt`: `datetime`, `CurrentStatus`: `varchar`, `FirstName`: `varchar`, `IDProfile`: `int`, `LastName`: `varchar`, `Photo`: `varchar`, `UpdatedAt`: `datetime`}
	_                = bytes.MinRead
)

func testOkProfilesUpdate(t *testing.T) {
	t.Parallel()

	if len(OkProfileColumns) == len(OkProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfileColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	if err = OkProfile.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOkProfilesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(OkProfileColumns) == len(OkProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkProfile := &OkProfile{}
	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkProfile, OkProfileDBTypes, true, OkProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(OkProfileColumns, OkProfilePrimaryKeyColumns) {
		fields = OkProfileColumns
	} else {
		fields = strmangle.SetComplement(
			OkProfileColumns,
			OkProfilePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(OkProfile))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OkProfileSlice{OkProfile}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOkProfilesUpsert(t *testing.T) {
	t.Parallel()

	if len(OkProfileColumns) == len(OkProfilePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	OkProfile := OkProfile{}
	if err = randomize.Struct(seed, &OkProfile, OkProfileDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkProfile.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkProfile: %s", err)
	}

	count, err := OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &OkProfile, OkProfileDBTypes, false, OkProfilePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkProfile struct: %s", err)
	}

	if err = OkProfile.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkProfile: %s", err)
	}

	count, err = OkProfiles(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
