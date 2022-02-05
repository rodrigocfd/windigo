package win

// Enumerates all modules.
func (hProcSnap HPROCSNAPSHOT) EnumModules(
	enumFunc func(me32 *MODULEENTRY32)) error {

	me32 := MODULEENTRY32{}
	me32.SetDwSize()

	found, err := hProcSnap.Module32First(&me32)
	for {
		if err != nil {
			return err
		} else if !found {
			break
		}
		enumFunc(&me32)
		found, err = hProcSnap.Module32Next(&me32)
	}
	return nil
}

// Enumerates all processes.
func (hProcSnap HPROCSNAPSHOT) EnumProcesses(
	enumFunc func(me32 *PROCESSENTRY32)) error {

	pe32 := PROCESSENTRY32{}
	pe32.SetDwSize()

	found, err := hProcSnap.Process32First(&pe32)
	for {
		if err != nil {
			return err
		} else if !found {
			break
		}
		enumFunc(&pe32)
		found, err = hProcSnap.Process32Next(&pe32)
	}
	return nil
}

// Enumerates all threads.
func (hProcSnap HPROCSNAPSHOT) EnumThreads(
	enumFunc func(me32 *THREADENTRY32)) error {

	te32 := THREADENTRY32{}
	te32.SetDwSize()

	found, err := hProcSnap.Thread32First(&te32)
	for {
		if err != nil {
			return err
		} else if !found {
			break
		}
		enumFunc(&te32)
		found, err = hProcSnap.Thread32Next(&te32)
	}
	return nil
}
