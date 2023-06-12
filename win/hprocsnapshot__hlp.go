//go:build windows

package win

// This helper method enumerates all modules.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, return false.
func (hProcSnap HPROCSNAPSHOT) EnumModules(
	callback func(me32 *MODULEENTRY32) bool) error {

	me32 := MODULEENTRY32{}
	me32.SetDwSize()

	found, err := hProcSnap.Module32First(&me32)
	for {
		if err != nil {
			return err
		} else if !found {
			break
		}
		if !callback(&me32) {
			break
		}
		found, err = hProcSnap.Module32Next(&me32)
	}
	return nil
}

// This helper method enumerates all processes.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, return false.
func (hProcSnap HPROCSNAPSHOT) EnumProcesses(
	callback func(me32 *PROCESSENTRY32) bool) error {

	pe32 := PROCESSENTRY32{}
	pe32.SetDwSize()

	found, err := hProcSnap.Process32First(&pe32)
	for {
		if err != nil {
			return err
		} else if !found {
			break
		}
		if !callback(&pe32) {
			break
		}
		found, err = hProcSnap.Process32Next(&pe32)
	}
	return nil
}

// This helper method enumerates all threads.
//
// To continue enumeration, the callback function must return true; to stop
// enumeration, return false.
func (hProcSnap HPROCSNAPSHOT) EnumThreads(
	callback func(me32 *THREADENTRY32) bool) error {

	te32 := THREADENTRY32{}
	te32.SetDwSize()

	found, err := hProcSnap.Thread32First(&te32)
	for {
		if err != nil {
			return err
		} else if !found {
			break
		}
		if !callback(&te32) {
			break
		}
		found, err = hProcSnap.Thread32Next(&te32)
	}
	return nil
}
