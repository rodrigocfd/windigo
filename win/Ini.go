package win

import (
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

type _IniT struct{}

// Functions to load/save an INI file into a map[string]map[string]string.
var Ini _IniT

// Lods an INI file from a file.
func (_IniT) FromFile(filePath string) (map[string]map[string]string, error) {
	fin, err := OpenFileMapped(filePath, co.OPEN_FILEMAP_MODE_READ)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	return Ini.FromStr(string(fin.ReadAll()))
}

// Loads an INI file from a string blob.
func (_IniT) FromStr(textBlob string) (map[string]map[string]string, error) {
	iniMap := make(map[string]map[string]string)
	curSectionName := ""

	for _, line := range strings.Split(textBlob, "\n") {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue // skip blank lines
		}

		if line[0] == '[' && line[len(line)-1] == ']' { // [section] ?
			curSectionName = line[1 : len(line)-1]
		} else {
			keyVal := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(keyVal[0])
			val := strings.TrimSpace(keyVal[1])

			if curSection, hasSection := iniMap[curSectionName]; hasSection {
				curSection[key] = val
			} else {
				iniMap[curSectionName] = make(map[string]string)
				iniMap[curSectionName][key] = val
			}
		}
	}

	return iniMap, nil
}

// Serializes the map into a string.
func (_IniT) ToStr(iniMap map[string]map[string]string) string {
	out := ""

	if namelessSection, hasSection := iniMap[""]; hasSection { // blank section?
		for k, v := range namelessSection {
			out += k + "=" + v + "\r\n"
		}
		out += "\r\n"
	}

	for sectionName, section := range iniMap {
		if sectionName == "" { // skip blank section, already written if any
			continue
		}

		out += "[" + sectionName + "]\n"
		for k, v := range section {
			out += k + "=" + v + "\r\n"
		}

		out += "\r\n"
	}

	return out
}

// Serializes the map into a file.
func (_IniT) ToFile(iniMap map[string]map[string]string, filePath string) error {
	serialized := Ini.ToStr(iniMap)

	fout, err := OpenFile(filePath, co.OPEN_FILE_RW_OPEN_OR_CREATE)
	if err != nil {
		return err
	}
	defer fout.Close()

	return fout.EraseAndWrite([]byte(serialized))
}
