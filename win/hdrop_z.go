//go:build windows

package win

// This helper method retrieves all file names with DragQueryFile() and calls
// DragFinish().
func (hDrop HDROP) ListFilesAndFinish() []string {
	var pathBuf [_MAX_PATH + 1]uint16 // buffer to receive all paths
	count := hDrop.DragQueryFile(0xffff_ffff, nil, 0)
	paths := make([]string, 0, count) // paths to be returned

	for i := uint32(0); i < count; i++ {
		hDrop.DragQueryFile(i, &pathBuf[0], uint32(len(pathBuf)))
		paths = append(paths, Str.FromNativeSlice(pathBuf[:]))
	}
	hDrop.DragFinish()

	Path.Sort(paths)
	return paths
}
