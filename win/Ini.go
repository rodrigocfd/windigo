package win

import (
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to an INI file.
//
// Created with IniLoad().
type Ini interface {
	AddSection(sectionName string) *_IniSection      // Adds a new section, if it doesn't exist yet.
	Get(sectionName, keyName string) (string, bool)  // Retrieves the specific section/key value, if existing.
	SaveToFile(filePath string) error                // Saves the INI to a file.
	Section(sectionName string) (*_IniSection, bool) // Returns the specific section, if existing.
	Sections() []_IniSection                         // Returns all the sections.
}

//------------------------------------------------------------------------------

type (
	_IniSection struct {
		Name string
		Keys []_IniKey
	}
	_IniKey struct {
		Name  string
		Value string
	}
)

type _Ini struct {
	sections []_IniSection
}

// Loads the sections and keys of an INI file.
func IniLoad(filePath string) (Ini, error) {
	me := &_Ini{}
	lines, err := me.loadLines(filePath)
	if err != nil {
		return nil, err
	}

	me.sections = make([]_IniSection, 0, 4) // arbitrary
	curSection := _IniSection{}

	for _, line := range lines {
		if len(line) == 0 {
			continue // skip blank lines
		}

		if line[0] == '[' && line[len(line)-1] == ']' { // [section] ?
			if curSection.Name != "" {
				me.sections = append(me.sections, curSection)
			}
			curSection = _IniSection{ // create a new section with the given name
				Name: strings.TrimSpace(line[1 : len(line)-1]),
				Keys: make([]_IniKey, 0, 4), // arbitrary
			}

		} else if curSection.Name != "" {
			keyVal := strings.SplitN(line, "=", 2)
			curSection.Keys = append(curSection.Keys, _IniKey{
				Name:  strings.TrimSpace(keyVal[0]),
				Value: strings.TrimSpace(keyVal[1]),
			})
		}
	}

	if curSection.Name != "" { // for the last section
		me.sections = append(me.sections, curSection)
	}

	return me, nil
}

func (me *_Ini) loadLines(filePath string) ([]string, error) {
	fin, err := FileMappedOpen(filePath, co.FILE_OPEN_READ_EXISTING)
	if err != nil {
		return nil, err
	}
	defer fin.Close()

	return fin.ReadLines(), nil
}

func (me *_Ini) AddSection(sectionName string) *_IniSection {
	if section, exists := me.Section(sectionName); exists {
		return section
	} else {
		me.sections = append(me.sections, _IniSection{
			Name: sectionName,
			Keys: make([]_IniKey, 0),
		})
		return &me.sections[len(me.sections)-1]
	}
}

func (me *_Ini) Get(sectionName, keyName string) (string, bool) {
	section, exists := me.Section(sectionName)
	if !exists {
		return "", false
	}

	for i := range section.Keys {
		if section.Keys[i].Name == keyName {
			return section.Keys[i].Name, true
		}
	}
	return "", false
}

func (me *_Ini) SaveToFile(filePath string) error {
	serialized := strings.Builder{}

	for i, section := range me.sections {
		serialized.WriteString("[" + section.Name + "]\r\n")
		for _, key := range section.Keys {
			serialized.WriteString(key.Name + "=" + key.Value + "\r\n")
		}

		isLast := i == len(me.sections)-1
		if !isLast {
			serialized.WriteString("\r\n")
		}
	}

	fout, err := FileOpen(filePath, co.FILE_OPEN_RW_OPEN_OR_CREATE)
	if err != nil {
		return err
	}
	defer fout.Close()

	return fout.EraseAndWrite([]byte(serialized.String()))
}

func (me *_Ini) Section(sectionName string) (*_IniSection, bool) {
	for i := range me.sections {
		if me.sections[i].Name == sectionName {
			return &me.sections[i], true
		}
	}
	return nil, false
}

func (me *_Ini) Sections() []_IniSection {
	return me.sections
}

// Adds a new key, if it doesn't exist yet.
func (me *_IniSection) AddKey(keyName, value string) *_IniKey {
	if key, exists := me.Key(keyName); exists {
		return key
	} else {
		me.Keys = append(me.Keys, _IniKey{
			Name:  keyName,
			Value: value,
		})
		return &me.Keys[len(me.Keys)-1]
	}
}

// Returns the specific key, if existing.
func (me *_IniSection) Key(keyName string) (*_IniKey, bool) {
	for i := range me.Keys {
		if me.Keys[i].Name == keyName {
			return &me.Keys[i], true
		}
	}
	return nil, false
}
