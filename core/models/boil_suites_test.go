// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Okakuns", testOkakuns)
	t.Run("Okprofiles", testOkprofiles)
	t.Run("Okprojects", testOkprojects)
	t.Run("Okprojecttasks", testOkprojecttasks)
	t.Run("Oktasks", testOktasks)
}

func TestDelete(t *testing.T) {
	t.Run("Okakuns", testOkakunsDelete)
	t.Run("Okprofiles", testOkprofilesDelete)
	t.Run("Okprojects", testOkprojectsDelete)
	t.Run("Okprojecttasks", testOkprojecttasksDelete)
	t.Run("Oktasks", testOktasksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Okakuns", testOkakunsQueryDeleteAll)
	t.Run("Okprofiles", testOkprofilesQueryDeleteAll)
	t.Run("Okprojects", testOkprojectsQueryDeleteAll)
	t.Run("Okprojecttasks", testOkprojecttasksQueryDeleteAll)
	t.Run("Oktasks", testOktasksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Okakuns", testOkakunsSliceDeleteAll)
	t.Run("Okprofiles", testOkprofilesSliceDeleteAll)
	t.Run("Okprojects", testOkprojectsSliceDeleteAll)
	t.Run("Okprojecttasks", testOkprojecttasksSliceDeleteAll)
	t.Run("Oktasks", testOktasksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Okakuns", testOkakunsExists)
	t.Run("Okprofiles", testOkprofilesExists)
	t.Run("Okprojects", testOkprojectsExists)
	t.Run("Okprojecttasks", testOkprojecttasksExists)
	t.Run("Oktasks", testOktasksExists)
}

func TestFind(t *testing.T) {
	t.Run("Okakuns", testOkakunsFind)
	t.Run("Okprofiles", testOkprofilesFind)
	t.Run("Okprojects", testOkprojectsFind)
	t.Run("Okprojecttasks", testOkprojecttasksFind)
	t.Run("Oktasks", testOktasksFind)
}

func TestBind(t *testing.T) {
	t.Run("Okakuns", testOkakunsBind)
	t.Run("Okprofiles", testOkprofilesBind)
	t.Run("Okprojects", testOkprojectsBind)
	t.Run("Okprojecttasks", testOkprojecttasksBind)
	t.Run("Oktasks", testOktasksBind)
}

func TestOne(t *testing.T) {
	t.Run("Okakuns", testOkakunsOne)
	t.Run("Okprofiles", testOkprofilesOne)
	t.Run("Okprojects", testOkprojectsOne)
	t.Run("Okprojecttasks", testOkprojecttasksOne)
	t.Run("Oktasks", testOktasksOne)
}

func TestAll(t *testing.T) {
	t.Run("Okakuns", testOkakunsAll)
	t.Run("Okprofiles", testOkprofilesAll)
	t.Run("Okprojects", testOkprojectsAll)
	t.Run("Okprojecttasks", testOkprojecttasksAll)
	t.Run("Oktasks", testOktasksAll)
}

func TestCount(t *testing.T) {
	t.Run("Okakuns", testOkakunsCount)
	t.Run("Okprofiles", testOkprofilesCount)
	t.Run("Okprojects", testOkprojectsCount)
	t.Run("Okprojecttasks", testOkprojecttasksCount)
	t.Run("Oktasks", testOktasksCount)
}

func TestHooks(t *testing.T) {
	t.Run("Okakuns", testOkakunsHooks)
	t.Run("Okprofiles", testOkprofilesHooks)
	t.Run("Okprojects", testOkprojectsHooks)
	t.Run("Okprojecttasks", testOkprojecttasksHooks)
	t.Run("Oktasks", testOktasksHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Okakuns", testOkakunsInsert)
	t.Run("Okakuns", testOkakunsInsertWhitelist)
	t.Run("Okprofiles", testOkprofilesInsert)
	t.Run("Okprofiles", testOkprofilesInsertWhitelist)
	t.Run("Okprojects", testOkprojectsInsert)
	t.Run("Okprojects", testOkprojectsInsertWhitelist)
	t.Run("Okprojecttasks", testOkprojecttasksInsert)
	t.Run("Okprojecttasks", testOkprojecttasksInsertWhitelist)
	t.Run("Oktasks", testOktasksInsert)
	t.Run("Oktasks", testOktasksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("OkprofileToOkakunUsingAkun", testOkprofileToOneOkakunUsingAkun)
	t.Run("OkprojecttaskToOkprofileUsingAssignedto", testOkprojecttaskToOneOkprofileUsingAssignedto)
	t.Run("OkprojecttaskToOkprofileUsingAssigner", testOkprojecttaskToOneOkprofileUsingAssigner)
	t.Run("OkprojecttaskToOkprojectUsingProject", testOkprojecttaskToOneOkprojectUsingProject)
	t.Run("OkprojecttaskToOktaskUsingTask", testOkprojecttaskToOneOktaskUsingTask)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("OkakunToAkunOkprofiles", testOkakunToManyAkunOkprofiles)
	t.Run("OkprofileToAssignedtoOkprojecttasks", testOkprofileToManyAssignedtoOkprojecttasks)
	t.Run("OkprofileToAssignerOkprojecttasks", testOkprofileToManyAssignerOkprojecttasks)
	t.Run("OkprojectToProjectOkprojecttasks", testOkprojectToManyProjectOkprojecttasks)
	t.Run("OktaskToTaskOkprojecttasks", testOktaskToManyTaskOkprojecttasks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("OkprofileToOkakunUsingAkun", testOkprofileToOneSetOpOkakunUsingAkun)
	t.Run("OkprojecttaskToOkprofileUsingAssignedto", testOkprojecttaskToOneSetOpOkprofileUsingAssignedto)
	t.Run("OkprojecttaskToOkprofileUsingAssigner", testOkprojecttaskToOneSetOpOkprofileUsingAssigner)
	t.Run("OkprojecttaskToOkprojectUsingProject", testOkprojecttaskToOneSetOpOkprojectUsingProject)
	t.Run("OkprojecttaskToOktaskUsingTask", testOkprojecttaskToOneSetOpOktaskUsingTask)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("OkakunToAkunOkprofiles", testOkakunToManyAddOpAkunOkprofiles)
	t.Run("OkprofileToAssignedtoOkprojecttasks", testOkprofileToManyAddOpAssignedtoOkprojecttasks)
	t.Run("OkprofileToAssignerOkprojecttasks", testOkprofileToManyAddOpAssignerOkprojecttasks)
	t.Run("OkprojectToProjectOkprojecttasks", testOkprojectToManyAddOpProjectOkprojecttasks)
	t.Run("OktaskToTaskOkprojecttasks", testOktaskToManyAddOpTaskOkprojecttasks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Okakuns", testOkakunsReload)
	t.Run("Okprofiles", testOkprofilesReload)
	t.Run("Okprojects", testOkprojectsReload)
	t.Run("Okprojecttasks", testOkprojecttasksReload)
	t.Run("Oktasks", testOktasksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Okakuns", testOkakunsReloadAll)
	t.Run("Okprofiles", testOkprofilesReloadAll)
	t.Run("Okprojects", testOkprojectsReloadAll)
	t.Run("Okprojecttasks", testOkprojecttasksReloadAll)
	t.Run("Oktasks", testOktasksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Okakuns", testOkakunsSelect)
	t.Run("Okprofiles", testOkprofilesSelect)
	t.Run("Okprojects", testOkprojectsSelect)
	t.Run("Okprojecttasks", testOkprojecttasksSelect)
	t.Run("Oktasks", testOktasksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Okakuns", testOkakunsUpdate)
	t.Run("Okprofiles", testOkprofilesUpdate)
	t.Run("Okprojects", testOkprojectsUpdate)
	t.Run("Okprojecttasks", testOkprojecttasksUpdate)
	t.Run("Oktasks", testOktasksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Okakuns", testOkakunsSliceUpdateAll)
	t.Run("Okprofiles", testOkprofilesSliceUpdateAll)
	t.Run("Okprojects", testOkprojectsSliceUpdateAll)
	t.Run("Okprojecttasks", testOkprojecttasksSliceUpdateAll)
	t.Run("Oktasks", testOktasksSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Okakuns", testOkakunsUpsert)
	t.Run("Okprofiles", testOkprofilesUpsert)
	t.Run("Okprojects", testOkprojectsUpsert)
	t.Run("Okprojecttasks", testOkprojecttasksUpsert)
	t.Run("Oktasks", testOktasksUpsert)
}
