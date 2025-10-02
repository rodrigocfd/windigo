//go:build windows

package win

import (
	"fmt"
	"strings"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/wstr"
)

type (
	// High-level abstraction for .ini file contents, loaded with [IniLoad].
	Ini struct {
		// Path to the .ini file that has been loaded. If you want to save the
		// .ini somewhere else, you may change this value.
		Path string
		// Sections of the .ini file. You may modify or rearrange these before
		// saving the file.
		Sections []IniSection
	}

	// High-level abstraction for .ini file sections.
	IniSection struct {
		// Name of the section. You may change this value before saving the
		// file.
		Name string
		// Entries of the section. You may modify or rearrange these before
		// saving the file.
		Entries []IniEntry
	}

	// High-level abstracion for .ini entries.
	IniEntry struct {
		// Name of the key. You may change this value before saving the file.
		Key string
		// Actual value. You may change this value before saving the file.
		Value string
	}
)

// Creates an [Ini] object by reading an .ini file, parsing its contents.
//
// Example:
//
//	ini, _ := win.IniLoad("C:\\Temp\\foo.ini")
func IniLoad(iniPath string) (*Ini, error) {
	lines, err := iniLoadLines(iniPath)
	if err != nil {
		return nil, fmt.Errorf("IniLoad: %w", err)
	}

	me := &Ini{}
	var curSection IniSection

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue // skip blank lines
		}

		if line[0] == '[' && line[len(line)-1] == ']' { // [section] ?
			if curSection.Name != "" {
				me.Sections = append(me.Sections, curSection)
			}
			curSection = IniSection{ // create a new section with the given name
				Name:    strings.TrimSpace(line[1 : len(line)-1]),
				Entries: make([]IniEntry, 0),
			}
		} else if curSection.Name != "" {
			keyVal := strings.SplitN(line, "=", 2)
			curSection.Entries = append(curSection.Entries, IniEntry{
				Key:   strings.TrimSpace(keyVal[0]),
				Value: strings.TrimSpace(keyVal[1]),
			})
		}
	}

	if curSection.Name != "" { // for the last section
		me.Sections = append(me.Sections, curSection)
	}

	me.Path = iniPath // keep
	return me, nil
}

func iniLoadLines(iniPath string) ([]string, error) {
	fin, err := FileMapOpen(iniPath, co.FOPEN_READ_EXISTING)
	if err != nil {
		return nil, err
	}
	defer fin.Close()
	return wstr.SplitLines(string(fin.HotSlice())), nil
}

// Serializes the contents.
func (me *Ini) Serialize() string {
	var serialized strings.Builder

	for idxSection := range me.Sections {
		section := &me.Sections[idxSection]
		serialized.WriteString("[" + section.Name + "]\r\n")
		for idxEntry := range section.Entries {
			value := &section.Entries[idxEntry]
			serialized.WriteString(value.Key + "=" + value.Value + "\r\n")
		}

		isLast := idxSection == len(me.Sections)-1
		if !isLast {
			serialized.WriteString("\r\n")
		}
	}

	return serialized.String()
}

// Serializes and saves the contents into the .ini file.
func (me *Ini) SaveToFile(filePath string) error {
	if err := FileWriteAll(filePath, []byte(me.Serialize())); err != nil {
		return fmt.Errorf("Ini.SaveToFile: %w", err)
	}
	return nil
}

// Returns a pointer to the section with the given name, or nil if not existing.
func (me *Ini) GetSection(name string) *IniSection {
	for idx := range me.Sections {
		if me.Sections[idx].Name == name {
			return &me.Sections[idx]
		}
	}
	return nil
}

// Returns the given value, if existing.
func (me *Ini) Get(section, key string) (string, bool) {
	if pSection := me.GetSection(section); pSection != nil {
		return pSection.Get(key)
	}
	return "", false
}

// Sets the given value. If the section/key pair doesn't exist, create it.
func (me *Ini) Set(section, key, value string) {
	if pSection := me.GetSection(section); pSection != nil {
		pSection.Set(key, value) // section exists
	} else {
		me.Sections = append(me.Sections, IniSection{ // section doesn't exist
			Name:    section,
			Entries: []IniEntry{{key, value}},
		})
	}
}

// Returns the given value, if existing.
func (me *IniSection) Get(key string) (string, bool) {
	for idx := range me.Entries {
		if me.Entries[idx].Key == key {
			return me.Entries[idx].Value, true
		}
	}
	return "", false
}

// Sets the given value. If the key doesn't exist, creates it.
func (me *IniSection) Set(key, value string) {
	for idx := range me.Entries {
		if me.Entries[idx].Key == key {
			me.Entries[idx].Value = value // key exists, set and return
			return
		}
	}
	me.Entries = append(me.Entries, IniEntry{key, value}) // key doesn't exist
}
