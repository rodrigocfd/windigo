//go:build windows

package win

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Reads all the contents of the file at once, immediately. Calls:
//   - [CreateFile]
//   - [HFILE.GetFileSizeEx]
//   - [HFILE.ReadFile]
//   - [HFILE.CloseHandle]
//
// Example:
//
//	contents, _ := win.FileReadAll("C:\\Temp\\foo.txt")
func FileReadAll(filePath string) ([]byte, error) {
	fin, err := FileOpen(filePath, co.FOPEN_READ_EXISTING)
	if err != nil {
		return nil, fmt.Errorf("FileReadAll: %w", err)
	}
	defer fin.Close()

	sz, err := fin.Size()
	if err != nil {
		return nil, fmt.Errorf("FileReadAll: %w", err)
	}

	ret := make([]byte, sz)
	if _, err := fin.Read(ret); err != nil {
		return nil, fmt.Errorf("FileReadAll: %w", err)
	}

	return ret, nil
}

// Truncates the file, then writes all the contents at once, immediately. Calls:
//   - [CreateFile]
//   - [HFILE.SetEndOfFile]
//   - [HFILE.WriteFile]
//   - [HFILE.CloseHandle]
//
// Example:
//
//	contents := []byte("my text")
//	_ = win.FileWriteAll("C:\\Temp\\foo.txt", contents)
func FileWriteAll(filePath string, contents []byte) error {
	fout, err := FileOpen(filePath, co.FOPEN_RW_OPEN_OR_CREATE)
	if err != nil {
		return fmt.Errorf("FileWriteAll: %w", err)
	}
	defer fout.Close()

	if err := fout.Truncate(); err != nil {
		return fmt.Errorf("FileWriteAll: %w", err)
	}

	if _, err := fout.Write(contents); err != nil {
		return fmt.Errorf("FileWriteAll: %w", err)
	}

	return nil
}

// High-level abstraction to [HFILE], providing several operations.
//
// If you simply need to read or write the contents, consider using the
// [FileReadAll] and [FileWriteAll] functions.
//
// Implements the following standard io interfaces:
//   - [io.ByteReader]
//   - [io.ByteWriter]
//   - [io.Closer]
//   - [io.Reader]
//   - [io.Seeker]
//   - [io.StringWriter]
//   - [io.Writer]
//
// Created with [FileOpen].
//
// Example:
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_RW_OPEN_OR_CREATE)
//	defer f.Close()
//
//	fmt.Fprintf(f, "foo")
type File struct {
	hFile HFILE
}

// Constructs a new [File] by calling [CreateFile].
//
// ⚠️ You must defer [File.Close].
//
// Example:
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_RW_OPEN_OR_CREATE)
//	defer f.Close()
func FileOpen(filePath string, desiredAccess co.FOPEN) (*File, error) {
	var access co.GENERIC
	var share co.FILE_SHARE
	var disposition co.DISPOSITION

	switch desiredAccess {
	case co.FOPEN_READ_EXISTING:
		access = co.GENERIC_READ
		share = co.FILE_SHARE_READ
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.FOPEN_RW_EXISTING:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_EXISTING
	case co.FOPEN_RW_OPEN_OR_CREATE:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_OPEN_ALWAYS
	case co.FOPEN_RW_CREATE:
		access = co.GENERIC_READ | co.GENERIC_WRITE
		share = co.FILE_SHARE_NONE
		disposition = co.DISPOSITION_CREATE_NEW
	}

	hFile, err := CreateFile(filePath, access, share, nil,
		disposition, co.FILE_ATTRIBUTE_NORMAL, co.FILE_FLAG_NONE,
		co.SECURITY_NONE, 0)
	if err != nil {
		return nil, fmt.Errorf("FileOpen CreateFile: %w", err)
	}

	return &File{hFile}, nil
}

// Implements [io.Closer].
//
// Calls [HFILE.CloseHandle].
func (me *File) Close() error {
	if me.hFile != 0 {
		err := me.hFile.CloseHandle()
		me.hFile = 0
		if err != nil {
			return fmt.Errorf("File.Close HFILE.CloseHandle: %w", err)
		}
	}
	return nil
}

// Returns the underlying handle.
func (me *File) Hfile() HFILE {
	return me.hFile
}

// Implements [io.Reader].
//
// Calls [HFILE.ReadFile] to read the file contents from its current internal
// pointer up to the buffer size.
func (me *File) Read(p []byte) (int, error) {
	numRead, err := me.hFile.ReadFile(p, nil)
	if err != nil {
		return 0, fmt.Errorf("File.Read HFILE.ReadFile: %w", err)
	}

	if numRead < len(p) { // buffer not completely filled, surely there's no more to read
		return int(numRead), io.EOF
	} else if numRead == 0 { // EOF found
		return 0, io.EOF
	} else { // still more to read
		return int(numRead), nil
	}
}

// Rewinds the internal file pointer and reads all contents at once, then
// rewinds the pointer again. Returns a []byte with the contents.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.ReadFile].
func (me *File) ReadAllAsSlice() ([]byte, error) {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	fileSize, err := me.Size()
	if err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	buf := make([]byte, fileSize)
	if _, err := me.Read(buf); err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("File.ReadAllAsSlice: %w", err)
	}

	return buf, nil
}

// Rewinds the internal file pointer and reads all contents at once, then
// rewinds the pointer again. Returns a [Vec] with the contents.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.ReadFile].
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// Example:
//
//	f, _ := win.FileOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data, _ := f.ReadAllAsVec()
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
//
// [SetFilePointerEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
// [ReadFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-readfile
func (me *File) ReadAllAsVec() (Vec[byte], error) {
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	fileSize, err := me.Size()
	if err != nil {
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	heapBuf := NewVecSized[byte](fileSize, 0x00)
	if _, err := me.Read(heapBuf.HotSlice()); err != nil {
		heapBuf.Free()
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	if _, err := me.Seek(0, io.SeekStart); err != nil {
		heapBuf.Free()
		return Vec[byte]{}, fmt.Errorf("File.ReadAllAsVec: %w", err)
	}

	return heapBuf, nil
}

// Implements [io.ByteReader].
//
// Calls [HFILE.ReadFile].
func (me *File) ReadByte() (byte, error) {
	var buf [1]byte
	if _, err := me.Read(buf[:]); err != nil {
		return 0, fmt.Errorf("File.ReadByte: %w", err)
	}
	return buf[0], nil
}

// Truncates or expands the file, according to the new size. Zero will empty the
// file. The internal file pointer will rewind.
//
// Calls [HFILE.SetFilePointerEx] and [HFILE.SetEndOfFile].
//
// For some reason, sometimes the resized file ends up with more space than
// requested. So be careful with this function.
//
// Panics if numBytes is negative.
func (me *File) Resize(numBytes int) error {
	utl.PanicNeg(numBytes)

	// Simply go beyond file limits if needed.
	if _, err := me.Seek(int64(numBytes), io.SeekStart); err != nil {
		return fmt.Errorf("File.Resize: %w", err)
	}

	if err := me.hFile.SetEndOfFile(); err != nil {
		return fmt.Errorf("File.Resize HFILE.SetEndOfFile: %w", err)
	}

	// Rewind pointer.
	if _, err := me.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("File.Resize: %w", err)
	}

	return nil
}

// Implements [io.Seeker].
//
// Moves the internal pointer with [HFILE.SetFilePointerEx].
func (me *File) Seek(offset int64, whence int) (int64, error) {
	var moveMethod co.FILE_FROM
	switch whence {
	case io.SeekCurrent:
		moveMethod = co.FILE_FROM_CURRENT
	case io.SeekStart:
		moveMethod = co.FILE_FROM_BEGIN
	case io.SeekEnd:
		moveMethod = co.FILE_FROM_END
	}

	newOff, err := me.hFile.SetFilePointerEx(int(offset), moveMethod)
	if err != nil {
		return 0, fmt.Errorf("File.Seek HFILE.SetFilePointerEx: %w", err)
	}
	return int64(newOff), nil
}

// Retrieves the file size with [HFILE.GetFileSizeEx]. This value is not cached.
func (me *File) Size() (int, error) {
	sz, err := me.hFile.GetFileSizeEx()
	if err != nil {
		return 0, fmt.Errorf("File.Size HFILE.GetFileSizeEx: %w", err)
	}
	return sz, nil
}

// Sets the file size to zero, deleting all its contents.
func (me *File) Truncate() error {
	if err := me.Resize(0); err != nil {
		return fmt.Errorf("File.Truncate: %w", err)
	}
	return nil
}

// Implements [io.Writer].
//
// Calls [HFILE.WriteFile] to write a slice at current internal pointer
// position.
func (me *File) Write(p []byte) (int, error) {
	written, err := me.hFile.WriteFile(p, nil)
	if err != nil {
		return 0, fmt.Errorf("File.Write HFILE.WriteFile: %w", err)
	}
	return int(written), nil
}

// Implements [io.ByteWriter].
//
// Calls [HFILE.WriteFile] to write a byte at current internal pointer position.
func (me *File) WriteByte(c byte) error {
	if _, err := me.Write([]byte{c}); err != nil {
		return fmt.Errorf("File.WriteByte: %w", err)
	}
	return nil
}

// Implements [io.StringWriter].
//
// Calls [HFILE.WriteFile] to write a string at current internal pointer
// position.
func (me *File) WriteString(s string) (int, error) {
	serialized := []byte(s)
	written, err := me.Write(serialized)
	if err != nil {
		return 0, fmt.Errorf("File.WriteString: %w", err)
	}
	return written, nil
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// High-level abstraction to [HFILEMAP], providing several operations.
//
// Note that memory-mapped files may present issues in x86 architectures; if so,
// just use the ordinary [File].
//
// If you simply need to read or write the contents at once, consider using the
// simpler [FileReadAll] and [FileWriteAll] functions.
//
// Created with [FileMapOpen].
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
type FileMap struct {
	file *File
	hMap HFILEMAP
	pMem HFILEMAPVIEW
	sz   int
}

// Constructs a new [FileMap] by opening the file and mapping it into memory
// with [HFILE.CreateFileMapping].
//
// Note that memory-mapped files may present issues in x86 architectures; if so,
// just call the ordinary [FileOpen] to work with a non-memory-mapped [File].
//
// ⚠️ You must defer [FileMap.Close].
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
func FileMapOpen(filePath string, desiredAccess co.FOPEN) (*FileMap, error) {
	file, err := FileOpen(filePath, desiredAccess)
	if err != nil {
		return nil, fmt.Errorf("FileMapOpen: %w", err)
	}

	// Map into memory.
	pageFlags := co.PAGE_READONLY
	if desiredAccess != co.FOPEN_READ_EXISTING {
		pageFlags = co.PAGE_READWRITE
	}

	hMap, err := file.Hfile().CreateFileMapping(nil, pageFlags, co.SEC_NONE, 0, "")
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("FileMapOpen HFILE.CreateFileMapping: %w", err)
	}

	// Get pointer to data block.
	mapFlags := co.FILE_MAP_READ
	if desiredAccess != co.FOPEN_READ_EXISTING {
		mapFlags = co.FILE_MAP_WRITE
	}

	pMem, err := hMap.MapViewOfFile(mapFlags, 0, 0)
	if err != nil {
		hMap.CloseHandle()
		file.Close()
		return nil, fmt.Errorf("FileMapOpen HFILEMAP.MapViewOfFile: %w", err)
	}

	// Cache file size.
	sz, err := file.Size()
	if err != nil {
		pMem.UnmapViewOfFile()
		hMap.CloseHandle()
		file.Close()
		return nil, fmt.Errorf("FileMapOpen: %w", err)
	}

	return &FileMap{file, hMap, pMem, sz}, nil
}

// Unmaps and releases the file resource.
func (me *FileMap) Close() error {
	var errRet error
	if me.pMem != 0 {
		err := me.pMem.UnmapViewOfFile()
		me.pMem = 0
		if err != nil {
			errRet = fmt.Errorf("FileMap.Close HFILEMAPVIEW.UnmapViewOfFile: %w", err)
		}
	}

	if me.hMap != 0 {
		err := me.hMap.CloseHandle()
		me.hMap = 0
		if err != nil && errRet == nil { // only report if pMem.UnmapViewOfFile() succeeded
			errRet = fmt.Errorf("FileMap.Close HFILEMAP.CloseHandle: %w", err)
		}
	}

	err := me.file.Close()
	me.sz = 0
	if err != nil && errRet == nil { // only report if hMap.CloseHandle() succeeded
		errRet = fmt.Errorf("FileMap.Close: %w", err)
	}

	return errRet
}

// Returns a slice to the memory-mapped bytes.
//
// The [FileMap] object must remain open while the slice is being used.
func (me *FileMap) HotSlice() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(me.pMem)), me.sz)
}

// Returns a new []byte with a copy of all data in the file.
func (me *FileMap) ReadAllAsSlice() []byte {
	return me.ReadChunkAsSlice(0, me.sz)
}

// Returns a new [Vec] with a copy of all data in the file.
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data := f.ReadAllAsVec()
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
func (me *FileMap) ReadAllAsVec() Vec[byte] {
	return me.ReadChunkAsVec(0, me.sz)
}

// Returns a new []byte with a copy of data, start with offset, and with the
// given length.
//
// Panics if offset or length is negative.
func (me *FileMap) ReadChunkAsSlice(offset, length int) []byte {
	utl.PanicNeg(offset, length)
	hotSlice := me.HotSlice()
	buf := make([]byte, length)
	copy(buf, hotSlice[offset:offset+length])
	return buf
}

// Returns a [Vec] with a copy of the data, start with offset, and with the
// given length.
//
// Panics if offset or length is negative.
//
// ⚠️ You must defer [Vec.Free] on the returned Vec.
//
// Example:
//
//	f, _ := win.FileMapOpen("C:\\Temp\\foo.txt", co.FOPEN_READ_EXISTING)
//	defer f.Close()
//
//	data := f.ReadChunkAsVec(0, 30)
//	defer data.Free()
//
//	txt := string(data.HotSlice())
//	println(txt)
func (me *FileMap) ReadChunkAsVec(offset, length int) Vec[byte] {
	utl.PanicNeg(offset, length)
	hotSlice := me.HotSlice()
	heapBuf := NewVec[byte]()
	heapBuf.Append(hotSlice[offset : offset+length]...)
	return heapBuf
}

// Retrieves the file size. This value is cached.
func (me *FileMap) Size() int {
	return me.sz
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

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

// Constructs a new [Ini] object by reading an .ini file, parsing its contents.
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

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// Returns a new []string with all files and folders within searchPath.
//
// If fileExtension isn't empty, brings only the files and folders with this
// extension.
//
// Does not search recursively. For a recursive search, use [PathEnumDeep].
//
// Calls:
//   - [FindFirstFile]
//   - [HFIND.FindNextFile]
//   - [HFIND.FindClose]
//
// Example:
//
//	paths := win.PathEnum("C:\\Temp", "")
//	mp3s := win.PathEnum("C:\\Temp", "mp3")
func PathEnum(searchPath, fileExtension string) ([]string, error) {
	if strings.Contains(searchPath, "*") {
		return nil, fmt.Errorf("invalid path: %s", searchPath)
	} else if strings.Contains(fileExtension, "*") {
		return nil, fmt.Errorf("invalid file extension: %s", fileExtension)
	}
	searchPath = PathTrimBackslash(searchPath)
	basePath := searchPath
	searchPath += "\\*"
	fileExtension = strings.TrimSpace(fileExtension)
	if fileExtension != "" {
		searchPath += "." + fileExtension
	}

	var wfd WIN32_FIND_DATA
	hFind, found, err := FindFirstFile(searchPath, &wfd)
	if err != nil {
		return nil, fmt.Errorf("PathEnum FindFirstFile: %w", err)
	} else if !found {
		return []string{}, nil // empty, not an error
	}
	defer hFind.FindClose()

	files := make([]string, 0, 20) // arbitrary
	for found {
		fileNameFound := wfd.CFileName()
		if fileNameFound != ".." && fileNameFound != "." {
			files = append(files, basePath+"\\"+fileNameFound)
		}

		if found, err = hFind.FindNextFile(&wfd); err != nil {
			return nil, fmt.Errorf("PathEnum HFIND.FindNextFile: %w", err)
		}
	}
	PathSort(files)
	return files, nil
}

// Returns a new []string with all files within searchPath.
//
// If fileExtension isn't empty, brings only the files with this extension.
//
// Searches recursively. For a non-recursive search, use [PathEnum].
//
// Calls:
//   - [FindFirstFile]
//   - [HFIND.FindNextFile]
//   - [HFIND.FindClose]
//
// Example:
//
//	paths := win.PathEnumDeep("C:\\Temp", "")
//	mp3s := win.PathEnumDeep("C:\\Temp", "mp3")
func PathEnumDeep(searchPath, fileExtension string) ([]string, error) {
	foundFiles, err := PathEnum(searchPath, "") // if we pass extension, subfolders will be skipped
	if err != nil {
		return nil, fmt.Errorf("PathEnumDeep: %w", err)
	}
	if len(foundFiles) == 0 {
		return []string{}, nil
	}

	files := make([]string, 0, len(foundFiles)+20) // arbitrary
	for _, f := range foundFiles {
		if !PathIsFolder(f) {
			if fileExtension == "" || PathHasExtension(f, fileExtension) { // manual extension filter
				files = append(files, f)
			}
		} else {
			nestedFiles, err := PathEnumDeep(f, fileExtension) // recursively
			if err != nil {
				return nil, err // don't wrap to avoid recursion repetition
			}
			files = append(files, nestedFiles...)
		}
	}
	return files, nil
}

// Returns true if the given file or folder exists. Calls [GetFileAttributes].
//
// Panics on error.
func PathExists(path string) bool {
	attr, err := GetFileAttributes(path)
	if err != nil {
		panic(err)
	}
	return attr != co.FILE_ATTRIBUTE_INVALID
}

// Retrieves the file name of the path.
func PathGetFileName(path string) string {
	if slashIdx := strings.LastIndex(path, "\\"); slashIdx == -1 {
		return path // path contains just the file name
	} else {
		return path[slashIdx+1:]
	}
}

// Retrieves the path without the file name itself, and without trailing
// backslash.
func PathGetPath(path string) string {
	if slashIdx := strings.LastIndex(path, "\\"); slashIdx == -1 {
		return "" // path contains just the file name
	} else {
		return path[0:slashIdx]
	}
}

// Returns whether the path ends with at least one of the given extensions.
//
// Example:
//
//	docPath := "C:\\Temp\\foo.txt"
//	isDocument := win.PathHasExtension(docPath, "txt", "doc")
func PathHasExtension(path string, extensions ...string) bool {
	pathUpper := strings.ToUpper(path)
	for _, extension := range extensions {
		if strings.HasSuffix(pathUpper, strings.ToUpper(extension)) {
			return true
		}
	}
	return false
}

// Returns true if the given path is a folder, and not a file. Calls
// [GetFileAttributes].
//
// Panics on error.
func PathIsFolder(path string) bool {
	attr, err := GetFileAttributes(path)
	if err != nil {
		panic(err)
	}
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_DIRECTORY) != 0
}

// Returns true if the given file or folder is hidden. Calls
// [GetFileAttributes].
//
// Panics on error.
func PathIsHidden(path string) bool {
	attr, err := GetFileAttributes(path)
	if err != nil {
		panic(err)
	}
	return attr != co.FILE_ATTRIBUTE_INVALID &&
		(attr&co.FILE_ATTRIBUTE_HIDDEN) != 0
}

// Sorts the paths alphabetically, case insensitive, in-place.
func PathSort(paths []string) {
	sort.Slice(paths, func(a, b int) bool {
		return strings.ToUpper(paths[a]) < strings.ToUpper(paths[b])
	})
}

// Replaces the current extension by the new one.
//
// Panics if the path doesn't have a file name.
func PathSwapExtension(path, newExtension string) string {
	if !strings.HasPrefix(newExtension, ".") {
		newExtension = "." + newExtension // must start with a dot
	}

	if strings.HasSuffix(path, "\\") {
		panic(fmt.Sprintf("Path doesn't have a file name: %s", path))
	}

	if idxDot := strings.LastIndex(path, "."); idxDot == -1 {
		return path + newExtension
	} else {
		return path[:idxDot] + newExtension
	}
}

// Returns a new string removing any trailing backslash.
func PathTrimBackslash(path string) string {
	path = strings.TrimSpace(path)
	lastIdx := len(path) - 1
	for path[lastIdx] == '\\' {
		lastIdx--
	}
	return path[:lastIdx+1]
}

// * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *

// Dynamic, manually OS [heap-allocated] memory block array.
//
// Created with:
//   - [NewVec]
//   - [NewVecReserved]
//   - [NewVecSized]
//
// Do not store Go pointers, slices or strings in a Vec – this will make the GC
// believe they are no more in use, thus collecting them.
//
// [heap-allocated]: https://learn.microsoft.com/en-us/windows/win32/Memory/heap-functions
type Vec[T any] struct {
	data  []T // Slice to the heap-allocated memory, its length is the capacity.
	inUse int // Number of elements effectively being used.
}

// Constructs a new, unallocated [Vec].
//
// Do not store Go pointers, slices or strings in a Vec – this will make the GC
// believe they are no more in use, thus collecting them.
//
// ⚠️ You must defer [Vec.Free].
//
// Example:
//
//	pts := win.NewVec[win.POINT]()
//	defer pts.Free()
func NewVec[T any]() Vec[T] {
	return Vec[T]{
		data:  nil,
		inUse: 0,
	}
}

// Constructs a new [Vec] with preallocated memory, but zero elements.
//
// Do not store Go pointers, slices or strings in a Vec – this will make the GC
// believe they are no more in use, thus collecting them.
//
// Panics if numElems is negative.
//
// ⚠️ You must defer [Vec.Free].
//
// Example:
//
//	pts := win.NewVecReserved[win.POINT](30)
//	defer pts.Free()
func NewVecReserved[T any](numElems int) Vec[T] {
	var me Vec[T]
	me.Reserve(numElems)
	return me
}

// Constructs a new [Vec] with numElems copies of elem.
//
// Do not store Go pointers, slices or strings in a Vec – this will make the GC
// believe they are no more in use, thus collecting them.
//
// Panics if numElems is negative.
//
// ⚠️ You must defer [Vec.Free].
//
// Example:
//
//	pts := win.NewVecSized(30, win.POINT{})
//	defer pts.Free()
func NewVecSized[T any](numElems int, elem T) Vec[T] {
	var me Vec[T]
	me.AppendN(numElems, elem)
	return me
}

// Appends new elements, increasing the buffer size if needed.
//
// Example:
//
//	bigNums := win.NewVec[uint64]()
//	defer bigNums.Free()
//
//	bigNums.Append(200)
//
//	others := []uint64{10, 20, 30}
//	bigNums.Append(others...)
func (me *Vec[T]) Append(elems ...T) {
	me.Reserve(me.inUse + len(elems))
	for _, elem := range elems {
		me.data[me.inUse] = elem
		me.inUse++
	}
}

// Appends numElems copies of elem, increasing the buffer size if needed.
//
// Panics if numElems is negative.
func (me *Vec[T]) AppendN(numElems int, elem T) {
	me.Reserve(me.inUse + numElems)
	for i := 0; i < numElems; i++ {
		me.data[me.inUse] = elem
		me.inUse++
	}
}

// Removes all elements, keeping the reserved size.
func (me *Vec[T]) Clear() {
	var dummy T
	for i := 0; i < me.inUse; i++ {
		me.data[i] = dummy
	}
	me.inUse = 0
}

// Releases the allocated heap memory, if any.
func (me *Vec[T]) Free() {
	if me.data != nil {
		hHeap, _ := GetProcessHeap()
		hHeap.HeapFree(co.HEAP_NS_NONE, unsafe.Pointer(unsafe.SliceData(me.data)))
		me.data = nil
		me.inUse = 0
	}
}

// Returns a pointer the element at the given position.
//
// If the buffer is changed for whathever reason – like by adding an element or
// reserving more space –, this pointer will be no longer valid.
//
// Does not perform bounds check. Panics if index is negative.
func (me *Vec[T]) Get(index int) *T {
	return &me.data[index]
}

// Returns a slice over the current elements.
//
// If the data is changed for whathever reason – like by adding an element or
// reserving more space –, the slice will be no longer valid.
func (me *Vec[T]) HotSlice() []T {
	if me.inUse == 0 {
		return []T{}
	} else {
		return me.data[:me.inUse]
	}
}

// Returns true if there are no elements.
func (me *Vec[T]) IsEmpty() bool {
	return me.inUse == 0
}

// Returns the number of elements currently stored, not counting the reserved
// space.
func (me *Vec[T]) Len() int {
	return me.inUse
}

// Returns a pointer to allocated memory block.
//
// If the buffer is changed for whathever reason – like by adding an element or
// reserving more space –, this pointer will be no longer valid.
func (me *Vec[T]) Ptr() unsafe.Pointer {
	if me.IsEmpty() {
		return nil
	} else {
		return unsafe.Pointer(unsafe.SliceData(me.data))
	}
}

// Allocates memory for the given number of elements, reserving the space,
// without adding elements.
//
// This method is intended for optimization purposes. If you want to create a
// buffer to receive data, use [Vec.Resize] instead.
//
// If amount is smaller than the current buffer size, does nothing; that is,
// this function only grows the buffer.
//
// Panics if numElems is negative.
func (me *Vec[T]) Reserve(numElems int) {
	utl.PanicNeg(numElems)
	if numElems > len(me.data) {
		newSizeBytes := numElems * me.szElem()
		hHeap, _ := GetProcessHeap()
		if me.data == nil { // first allocation
			ptr, _ := hHeap.HeapAlloc(co.HEAP_ALLOC_ZERO_MEMORY, newSizeBytes)
			me.data = unsafe.Slice((*T)(ptr), numElems) // store slice
		} else {
			curPtr := unsafe.Pointer(unsafe.SliceData(me.data))
			newPtr, _ := hHeap.HeapReAlloc(co.HEAP_REALLOC_ZERO_MEMORY, curPtr, newSizeBytes)
			me.data = unsafe.Slice((*T)(newPtr), numElems) // store new slice
		}
	}
}

// Returns the actual number of allocated elements in the buffer.
func (me *Vec[T]) Reserved() int {
	return len(me.data)
}

// Resizes the internal buffer to the given number of elements. If increased,
// the given element is used to fill the new positions.
//
// Panics if numElems is negative.
func (me *Vec[T]) Resize(numElems int, elemToFill T) {
	utl.PanicNeg(numElems)
	if numElems > me.inUse { // enlarge
		me.AppendN(numElems-me.inUse, elemToFill)
	} else if me.inUse > numElems { // shrink
		var dummy T
		for i := numElems; i < me.inUse; i++ {
			me.data[i] = dummy // fill the unused memory
		}
		me.inUse = numElems
	}
}

// Size of a single element, in bytes.
func (me *Vec[T]) szElem() int {
	var dummy T
	return int(unsafe.Sizeof(dummy))
}
