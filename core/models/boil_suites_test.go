package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("OkAkuns", testOkAkuns)
	t.Run("OkProfiles", testOkProfiles)
	t.Run("OkProjects", testOkProjects)
	t.Run("OkProjectTasks", testOkProjectTasks)
	t.Run("OkTasks", testOkTasks)
}

func TestDelete(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsDelete)
	t.Run("OkProfiles", testOkProfilesDelete)
	t.Run("OkProjects", testOkProjectsDelete)
	t.Run("OkProjectTasks", testOkProjectTasksDelete)
	t.Run("OkTasks", testOkTasksDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsQueryDeleteAll)
	t.Run("OkProfiles", testOkProfilesQueryDeleteAll)
	t.Run("OkProjects", testOkProjectsQueryDeleteAll)
	t.Run("OkProjectTasks", testOkProjectTasksQueryDeleteAll)
	t.Run("OkTasks", testOkTasksQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsSliceDeleteAll)
	t.Run("OkProfiles", testOkProfilesSliceDeleteAll)
	t.Run("OkProjects", testOkProjectsSliceDeleteAll)
	t.Run("OkProjectTasks", testOkProjectTasksSliceDeleteAll)
	t.Run("OkTasks", testOkTasksSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsExists)
	t.Run("OkProfiles", testOkProfilesExists)
	t.Run("OkProjects", testOkProjectsExists)
	t.Run("OkProjectTasks", testOkProjectTasksExists)
	t.Run("OkTasks", testOkTasksExists)
}

func TestFind(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsFind)
	t.Run("OkProfiles", testOkProfilesFind)
	t.Run("OkProjects", testOkProjectsFind)
	t.Run("OkProjectTasks", testOkProjectTasksFind)
	t.Run("OkTasks", testOkTasksFind)
}

func TestBind(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsBind)
	t.Run("OkProfiles", testOkProfilesBind)
	t.Run("OkProjects", testOkProjectsBind)
	t.Run("OkProjectTasks", testOkProjectTasksBind)
	t.Run("OkTasks", testOkTasksBind)
}

func TestOne(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsOne)
	t.Run("OkProfiles", testOkProfilesOne)
	t.Run("OkProjects", testOkProjectsOne)
	t.Run("OkProjectTasks", testOkProjectTasksOne)
	t.Run("OkTasks", testOkTasksOne)
}

func TestAll(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsAll)
	t.Run("OkProfiles", testOkProfilesAll)
	t.Run("OkProjects", testOkProjectsAll)
	t.Run("OkProjectTasks", testOkProjectTasksAll)
	t.Run("OkTasks", testOkTasksAll)
}

func TestCount(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsCount)
	t.Run("OkProfiles", testOkProfilesCount)
	t.Run("OkProjects", testOkProjectsCount)
	t.Run("OkProjectTasks", testOkProjectTasksCount)
	t.Run("OkTasks", testOkTasksCount)
}

func TestHooks(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsHooks)
	t.Run("OkProfiles", testOkProfilesHooks)
	t.Run("OkProjects", testOkProjectsHooks)
	t.Run("OkProjectTasks", testOkProjectTasksHooks)
	t.Run("OkTasks", testOkTasksHooks)
}

func TestInsert(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsInsert)
	t.Run("OkAkuns", testOkAkunsInsertWhitelist)
	t.Run("OkProfiles", testOkProfilesInsert)
	t.Run("OkProfiles", testOkProfilesInsertWhitelist)
	t.Run("OkProjects", testOkProjectsInsert)
	t.Run("OkProjects", testOkProjectsInsertWhitelist)
	t.Run("OkProjectTasks", testOkProjectTasksInsert)
	t.Run("OkProjectTasks", testOkProjectTasksInsertWhitelist)
	t.Run("OkTasks", testOkTasksInsert)
	t.Run("OkTasks", testOkTasksInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("OkProfileToOkAkunUsingAkun", testOkProfileToOneOkAkunUsingAkun)
	t.Run("OkProjectTaskToOkProfileUsingAssignedto", testOkProjectTaskToOneOkProfileUsingAssignedto)
	t.Run("OkProjectTaskToOkProfileUsingAssigner", testOkProjectTaskToOneOkProfileUsingAssigner)
	t.Run("OkProjectTaskToOkProjectUsingProject", testOkProjectTaskToOneOkProjectUsingProject)
	t.Run("OkProjectTaskToOkTaskUsingTask", testOkProjectTaskToOneOkTaskUsingTask)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("OkAkunToAkunOkProfiles", testOkAkunToManyAkunOkProfiles)
	t.Run("OkProfileToAssignedtoOkProjectTasks", testOkProfileToManyAssignedtoOkProjectTasks)
	t.Run("OkProfileToAssignerOkProjectTasks", testOkProfileToManyAssignerOkProjectTasks)
	t.Run("OkProjectToProjectOkProjectTasks", testOkProjectToManyProjectOkProjectTasks)
	t.Run("OkTaskToTaskOkProjectTasks", testOkTaskToManyTaskOkProjectTasks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("OkProfileToOkAkunUsingAkun", testOkProfileToOneSetOpOkAkunUsingAkun)
	t.Run("OkProjectTaskToOkProfileUsingAssignedto", testOkProjectTaskToOneSetOpOkProfileUsingAssignedto)
	t.Run("OkProjectTaskToOkProfileUsingAssigner", testOkProjectTaskToOneSetOpOkProfileUsingAssigner)
	t.Run("OkProjectTaskToOkProjectUsingProject", testOkProjectTaskToOneSetOpOkProjectUsingProject)
	t.Run("OkProjectTaskToOkTaskUsingTask", testOkProjectTaskToOneSetOpOkTaskUsingTask)
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
	t.Run("OkAkunToAkunOkProfiles", testOkAkunToManyAddOpAkunOkProfiles)
	t.Run("OkProfileToAssignedtoOkProjectTasks", testOkProfileToManyAddOpAssignedtoOkProjectTasks)
	t.Run("OkProfileToAssignerOkProjectTasks", testOkProfileToManyAddOpAssignerOkProjectTasks)
	t.Run("OkProjectToProjectOkProjectTasks", testOkProjectToManyAddOpProjectOkProjectTasks)
	t.Run("OkTaskToTaskOkProjectTasks", testOkTaskToManyAddOpTaskOkProjectTasks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsReload)
	t.Run("OkProfiles", testOkProfilesReload)
	t.Run("OkProjects", testOkProjectsReload)
	t.Run("OkProjectTasks", testOkProjectTasksReload)
	t.Run("OkTasks", testOkTasksReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsReloadAll)
	t.Run("OkProfiles", testOkProfilesReloadAll)
	t.Run("OkProjects", testOkProjectsReloadAll)
	t.Run("OkProjectTasks", testOkProjectTasksReloadAll)
	t.Run("OkTasks", testOkTasksReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsSelect)
	t.Run("OkProfiles", testOkProfilesSelect)
	t.Run("OkProjects", testOkProjectsSelect)
	t.Run("OkProjectTasks", testOkProjectTasksSelect)
	t.Run("OkTasks", testOkTasksSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsUpdate)
	t.Run("OkProfiles", testOkProfilesUpdate)
	t.Run("OkProjects", testOkProjectsUpdate)
	t.Run("OkProjectTasks", testOkProjectTasksUpdate)
	t.Run("OkTasks", testOkTasksUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsSliceUpdateAll)
	t.Run("OkProfiles", testOkProfilesSliceUpdateAll)
	t.Run("OkProjects", testOkProjectsSliceUpdateAll)
	t.Run("OkProjectTasks", testOkProjectTasksSliceUpdateAll)
	t.Run("OkTasks", testOkTasksSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("OkAkuns", testOkAkunsUpsert)
	t.Run("OkProfiles", testOkProfilesUpsert)
	t.Run("OkProjects", testOkProjectsUpsert)
	t.Run("OkProjectTasks", testOkProjectTasksUpsert)
	t.Run("OkTasks", testOkTasksUpsert)
}
