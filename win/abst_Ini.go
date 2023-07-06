//go:build windows

package win

import (
	"runtime"
	"strings"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to a .ini file.
//
// Contains a slice of sections, which can be freely modified.
//
// Created with IniLoad().
type Ini struct {
	Sections   []IniSection // All sections of this .ini file.
	sourcePath string
}

// Loads the sections and keys of an INI file.
func IniLoad(filePath string) (*Ini, error) {
	me := &Ini{}
	lines, err := me.loadLines(filePath)
	if err != nil {
		return nil, err
	}

	me.Sections = make([]IniSection, 0, 4) // arbitrary
	var curSection IniSection

	for _, line := range lines {
		if len(line) == 0 {
			continue // skip blank lines
		}

		if line[0] == '[' && line[len(line)-1] == ']' { // [section] ?
			if curSection.Name != "" {
				me.Sections = append(me.Sections, curSection)
			}
			curSection = IniSection{ // create a new section with the given name
				Name:   strings.TrimSpace(line[1 : len(line)-1]),
				Values: make([]IniKey, 0, 4), // arbitrary
			}

		} else if curSection.Name != "" {
			keyVal := strings.SplitN(line, "=", 2)
			curSection.Values = append(curSection.Values, IniKey{
				Name:  strings.TrimSpace(keyVal[0]),
				Value: strings.TrimSpace(keyVal[1]),
			})
		}
	}

	if curSection.Name != "" { // for the last section
		me.Sections = append(me.Sections, curSection)
	}

	me.sourcePath = filePath // keep
	return me, nil
}

func (me *Ini) loadLines(filePath string) ([]string, error) {
	if runtime.GOARCH == "386" { // MapViewOfFile may have issues in x86
		fin, err := FileOpen(filePath, co.FILE_OPEN_READ_EXISTING)
		if err != nil {
			return nil, err
		}
		defer fin.Close()
		return fin.ReadLines()
	} else {
		fin, err := FileMappedOpen(filePath, co.FILE_OPEN_READ_EXISTING)
		if err != nil {
			return nil, err
		}
		defer fin.Close()
		return fin.ReadLines(), nil
	}
}

// Returns the IniSection with the given name, if any.
func (me *Ini) Section(name string) (*IniSection, bool) {
	for i := range me.Sections {
		section := &me.Sections[i]
		if section.Name == name {
			return section, true
		}
	}
	return nil, false
}

// Saves the contents to a .ini file.
func (me *Ini) SaveToFile(filePath string) error {
	var serialized strings.Builder

	for s := range me.Sections {
		section := &me.Sections[s]
		serialized.WriteString("[" + section.Name + "]\r\n")
		for v := range section.Values {
			value := &section.Values[v]
			serialized.WriteString(value.Name + "=" + value.Value + "\r\n")
		}

		isLast := s == len(me.Sections)-1
		if !isLast {
			serialized.WriteString("\r\n")
		}
	}

	fout, err := FileOpen(filePath, co.FILE_OPEN_RW_OPEN_OR_CREATE)
	if err != nil {
		return err
	}
	defer fout.Close()

	blob := []byte(serialized.String())
	fout.Resize(len(blob))
	_, err = fout.Write(blob)

	me.sourcePath = filePath // update
	return err
}

// Returns the latest source path of the .ini file.
//
// When IniLoad() is called, this path is stored. When Ini.SaveToFile() is
// called, this new path is stored.
//
// This is useful when you load an .ini file, and you need to update it later.
// When you first open, you pass the file path; on subsequent saves, you just
// call Ini.SourcePath() to retrieve the path, instead of manually saving it
// somewhere.
//
// # Example
//
//	ini, _ := win.IniLoad("C:\\Temp\\foo.ini")
//
//	// modify ini...
//
//	ini.SaveToFile(ini.SourcePath())
func (me *Ini) SourcePath() string {
	return me.sourcePath
}

// Returns the specific value, if existing.
//
// Note that a pointer to the string is returned, so that the value can be
// directly modified.
//
// # Example
//
//	ini, _ := win.IniLoad("C:\\Temp\\foo.ini")
//
//	if val, ok := ini.Value("my section", "my value"); ok {
//		fmt.Printf("Old value: %s\n", *val)
//		*val = "new value"
//	}
func (me *Ini) Value(sectionName, valueName string) (*string, bool) {
	if section, ok := me.Section(sectionName); ok {
		if value, ok := section.Value(valueName); ok {
			return &value.Value, true
		}
	}
	return nil, false
}

//------------------------------------------------------------------------------

// A single section of an Ini.
//
// Contains a slice of keys, which can be freely modified.
type IniSection struct {
	Name   string   // The name of this section.
	Values []IniKey // All values of this section.
}

// Returns the IniKey with the given name, if any.
func (me *IniSection) Value(name string) (*IniKey, bool) {
	for i := range me.Values {
		value := &me.Values[i]
		if value.Name == name {
			return value, true
		}
	}
	return nil, false
}

//------------------------------------------------------------------------------

// A single key of an IniSection.
type IniKey struct {
	Name  string // The name of this key.
	Value string // The value of this key.
}
