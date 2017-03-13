// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOktasks(t *testing.T) {
	t.Parallel()

	query := Oktasks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOktasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = oktask.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOktasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Oktasks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOktasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OktaskSlice{oktask}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOktasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OktaskExists(tx, oktask.IDTask)
	if err != nil {
		t.Errorf("Unable to check if Oktask exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OktaskExistsG to return true, but got false.")
	}
}
func testOktasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	oktaskFound, err := FindOktask(tx, oktask.IDTask)
	if err != nil {
		t.Error(err)
	}

	if oktaskFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOktasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Oktasks(tx).Bind(oktask); err != nil {
		t.Error(err)
	}
}

func testOktasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Oktasks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOktasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktaskOne := &Oktask{}
	oktaskTwo := &Oktask{}
	if err = randomize.Struct(seed, oktaskOne, oktaskDBTypes, false, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}
	if err = randomize.Struct(seed, oktaskTwo, oktaskDBTypes, false, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktaskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = oktaskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Oktasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOktasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	oktaskOne := &Oktask{}
	oktaskTwo := &Oktask{}
	if err = randomize.Struct(seed, oktaskOne, oktaskDBTypes, false, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}
	if err = randomize.Struct(seed, oktaskTwo, oktaskDBTypes, false, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktaskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = oktaskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func oktaskBeforeInsertHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskAfterInsertHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskAfterSelectHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskBeforeUpdateHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskAfterUpdateHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskBeforeDeleteHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskAfterDeleteHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskBeforeUpsertHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func oktaskAfterUpsertHook(e boil.Executor, o *Oktask) error {
	*o = Oktask{}
	return nil
}

func testOktasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Oktask{}
	o := &Oktask{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, oktaskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Oktask object: %s", err)
	}

	AddOktaskHook(boil.BeforeInsertHook, oktaskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	oktaskBeforeInsertHooks = []OktaskHook{}

	AddOktaskHook(boil.AfterInsertHook, oktaskAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	oktaskAfterInsertHooks = []OktaskHook{}

	AddOktaskHook(boil.AfterSelectHook, oktaskAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	oktaskAfterSelectHooks = []OktaskHook{}

	AddOktaskHook(boil.BeforeUpdateHook, oktaskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	oktaskBeforeUpdateHooks = []OktaskHook{}

	AddOktaskHook(boil.AfterUpdateHook, oktaskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	oktaskAfterUpdateHooks = []OktaskHook{}

	AddOktaskHook(boil.BeforeDeleteHook, oktaskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	oktaskBeforeDeleteHooks = []OktaskHook{}

	AddOktaskHook(boil.AfterDeleteHook, oktaskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	oktaskAfterDeleteHooks = []OktaskHook{}

	AddOktaskHook(boil.BeforeUpsertHook, oktaskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	oktaskBeforeUpsertHooks = []OktaskHook{}

	AddOktaskHook(boil.AfterUpsertHook, oktaskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	oktaskAfterUpsertHooks = []OktaskHook{}
}
func testOktasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOktasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx, oktaskColumns...); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOktaskToManyTaskOkprojecttasks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Oktask
	var b, c Okprojecttask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, okprojecttaskDBTypes, false, okprojecttaskColumnsWithDefault...)
	randomize.Struct(seed, &c, okprojecttaskDBTypes, false, okprojecttaskColumnsWithDefault...)

	b.TaskID = a.IDTask
	c.TaskID = a.IDTask
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	okprojecttask, err := a.TaskOkprojecttasks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range okprojecttask {
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

	slice := OktaskSlice{&a}
	if err = a.L.LoadTaskOkprojecttasks(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TaskOkprojecttasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TaskOkprojecttasks = nil
	if err = a.L.LoadTaskOkprojecttasks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TaskOkprojecttasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", okprojecttask)
	}
}

func testOktaskToManyAddOpTaskOkprojecttasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Oktask
	var b, c, d, e Okprojecttask

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, oktaskDBTypes, false, strmangle.SetComplement(oktaskPrimaryKeyColumns, oktaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Okprojecttask{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, okprojecttaskDBTypes, false, strmangle.SetComplement(okprojecttaskPrimaryKeyColumns, okprojecttaskColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Okprojecttask{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTaskOkprojecttasks(tx, i != 0, x...)
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

		if a.R.TaskOkprojecttasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TaskOkprojecttasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TaskOkprojecttasks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOktasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = oktask.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOktasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OktaskSlice{oktask}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOktasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Oktasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	oktaskDBTypes = map[string]string{`DescTask`: `longtext`, `Difficulty`: `varchar`, `IDTask`: `int`, `TitleTask`: `varchar`, `Urgency`: `varchar`}
	_             = bytes.MinRead
)

func testOktasksUpdate(t *testing.T) {
	t.Parallel()

	if len(oktaskColumns) == len(oktaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	if err = oktask.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOktasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(oktaskColumns) == len(oktaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	oktask := &Oktask{}
	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, oktask, oktaskDBTypes, true, oktaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(oktaskColumns, oktaskPrimaryKeyColumns) {
		fields = oktaskColumns
	} else {
		fields = strmangle.SetComplement(
			oktaskColumns,
			oktaskPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(oktask))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OktaskSlice{oktask}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOktasksUpsert(t *testing.T) {
	t.Parallel()

	if len(oktaskColumns) == len(oktaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	oktask := Oktask{}
	if err = randomize.Struct(seed, &oktask, oktaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = oktask.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Oktask: %s", err)
	}

	count, err := Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &oktask, oktaskDBTypes, false, oktaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Oktask struct: %s", err)
	}

	if err = oktask.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Oktask: %s", err)
	}

	count, err = Oktasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
