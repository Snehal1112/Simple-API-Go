package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testOkAkuns(t *testing.T) {
	t.Parallel()

	query := OkAkuns(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOkAkunsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkAkun.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkAkunsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkAkuns(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOkAkunsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkAkunSlice{OkAkun}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOkAkunsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OkAkunExists(tx, OkAkun.IDAkun)
	if err != nil {
		t.Errorf("Unable to check if OkAkun exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OkAkunExistsG to return true, but got false.")
	}
}
func testOkAkunsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	OkAkunFound, err := FindOkAkun(tx, OkAkun.IDAkun)
	if err != nil {
		t.Error(err)
	}

	if OkAkunFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOkAkunsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkAkuns(tx).Bind(OkAkun); err != nil {
		t.Error(err)
	}
}

func testOkAkunsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := OkAkuns(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOkAkunsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkunOne := &OkAkun{}
	OkAkunTwo := &OkAkun{}
	if err = randomize.Struct(seed, OkAkunOne, OkAkunDBTypes, false, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}
	if err = randomize.Struct(seed, OkAkunTwo, OkAkunDBTypes, false, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkunOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkAkunTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkAkuns(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOkAkunsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	OkAkunOne := &OkAkun{}
	OkAkunTwo := &OkAkun{}
	if err = randomize.Struct(seed, OkAkunOne, OkAkunDBTypes, false, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}
	if err = randomize.Struct(seed, OkAkunTwo, OkAkunDBTypes, false, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkunOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = OkAkunTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func OkAkunBeforeInsertHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunAfterInsertHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunAfterSelectHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunBeforeUpdateHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunAfterUpdateHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunBeforeDeleteHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunAfterDeleteHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunBeforeUpsertHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func OkAkunAfterUpsertHook(e boil.Executor, o *OkAkun) error {
	*o = OkAkun{}
	return nil
}

func testOkAkunsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &OkAkun{}
	o := &OkAkun{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, OkAkunDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OkAkun object: %s", err)
	}

	AddOkAkunHook(boil.BeforeInsertHook, OkAkunBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	OkAkunBeforeInsertHooks = []OkAkunHook{}

	AddOkAkunHook(boil.AfterInsertHook, OkAkunAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	OkAkunAfterInsertHooks = []OkAkunHook{}

	AddOkAkunHook(boil.AfterSelectHook, OkAkunAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	OkAkunAfterSelectHooks = []OkAkunHook{}

	AddOkAkunHook(boil.BeforeUpdateHook, OkAkunBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	OkAkunBeforeUpdateHooks = []OkAkunHook{}

	AddOkAkunHook(boil.AfterUpdateHook, OkAkunAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	OkAkunAfterUpdateHooks = []OkAkunHook{}

	AddOkAkunHook(boil.BeforeDeleteHook, OkAkunBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	OkAkunBeforeDeleteHooks = []OkAkunHook{}

	AddOkAkunHook(boil.AfterDeleteHook, OkAkunAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	OkAkunAfterDeleteHooks = []OkAkunHook{}

	AddOkAkunHook(boil.BeforeUpsertHook, OkAkunBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	OkAkunBeforeUpsertHooks = []OkAkunHook{}

	AddOkAkunHook(boil.AfterUpsertHook, OkAkunAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	OkAkunAfterUpsertHooks = []OkAkunHook{}
}
func testOkAkunsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkAkunsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx, OkAkunColumns...); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOkAkunToManyAkunOkProfiles(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkAkun
	var b, c OkProfile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, OkProfileDBTypes, false, OkProfileColumnsWithDefault...)
	randomize.Struct(seed, &c, OkProfileDBTypes, false, OkProfileColumnsWithDefault...)

	b.AkunID = a.IDAkun
	c.AkunID = a.IDAkun
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	OkProfile, err := a.AkunOkProfiles(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range OkProfile {
		if v.AkunID == b.AkunID {
			bFound = true
		}
		if v.AkunID == c.AkunID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OkAkunSlice{&a}
	if err = a.L.LoadAkunOkProfiles(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AkunOkProfiles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.AkunOkProfiles = nil
	if err = a.L.LoadAkunOkProfiles(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.AkunOkProfiles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", OkProfile)
	}
}

func testOkAkunToManyAddOpAkunOkProfiles(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a OkAkun
	var b, c, d, e OkProfile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, OkAkunDBTypes, false, strmangle.SetComplement(OkAkunPrimaryKeyColumns, OkAkunColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OkProfile{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, OkProfileDBTypes, false, strmangle.SetComplement(OkProfilePrimaryKeyColumns, OkProfileColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*OkProfile{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAkunOkProfiles(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.IDAkun != first.AkunID {
			t.Error("foreign key was wrong value", a.IDAkun, first.AkunID)
		}
		if a.IDAkun != second.AkunID {
			t.Error("foreign key was wrong value", a.IDAkun, second.AkunID)
		}

		if first.R.Akun != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Akun != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.AkunOkProfiles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.AkunOkProfiles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.AkunOkProfiles(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOkAkunsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = OkAkun.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOkAkunsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OkAkunSlice{OkAkun}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOkAkunsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := OkAkuns(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	OkAkunDBTypes = map[string]string{`CreatedAt`: `datetime`, `IDAkun`: `int`, `Password`: `varchar`, `Role`: `varchar`, `Token`: `text`, `UpdatedAt`: `datetime`, `Username`: `varchar`}
	_             = bytes.MinRead
)

func testOkAkunsUpdate(t *testing.T) {
	t.Parallel()

	if len(OkAkunColumns) == len(OkAkunPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	if err = OkAkun.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOkAkunsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(OkAkunColumns) == len(OkAkunPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	OkAkun := &OkAkun{}
	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, OkAkun, OkAkunDBTypes, true, OkAkunPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(OkAkunColumns, OkAkunPrimaryKeyColumns) {
		fields = OkAkunColumns
	} else {
		fields = strmangle.SetComplement(
			OkAkunColumns,
			OkAkunPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(OkAkun))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OkAkunSlice{OkAkun}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOkAkunsUpsert(t *testing.T) {
	t.Parallel()

	if len(OkAkunColumns) == len(OkAkunPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	OkAkun := OkAkun{}
	if err = randomize.Struct(seed, &OkAkun, OkAkunDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = OkAkun.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkAkun: %s", err)
	}

	count, err := OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &OkAkun, OkAkunDBTypes, false, OkAkunPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OkAkun struct: %s", err)
	}

	if err = OkAkun.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert OkAkun: %s", err)
	}

	count, err = OkAkuns(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
