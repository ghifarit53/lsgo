package main

import "fmt"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[97m"
)

type FileStyle struct {
	Icon  string
	Color string
}

var fileStyle = map[string]FileStyle{
	// default icon
	"default": {"󰈔", White},

	// directories
	"dir":               {"󰉋", Blue},
	"dir_symlink":       {"󰉋", Cyan},
	"dir_empty":         {"󰉖", Blue},
	"dir_empty_symlink": {"󰉖", Cyan},

	// default for files
	"file_empty":      {"󰈤", White},
	"file_symlink":    {"󰌹", Cyan},
	"file_executable": {"", Red},

	// developer icons
	"c":     {"", Blue},
	"cpp":   {"", Blue},
	"dart":  {"", Cyan},
	"go":    {"", Blue},
	"html":  {"", Green},
	"java":  {"", Red},
	"js":    {"", Yellow},
	"jsx":   {"", Yellow},
	"kt":    {"", Magenta},
	"perl":  {"", Blue},
	"php":   {"", Magenta},
	"py":    {"", Yellow},
	"rb":    {"", Red},
	"rs":    {"", Blue},
	"sh":    {"", Magenta},
	"swift": {"", Red},
	"ts":    {"", Blue},
	"tsx":   {"", Blue},

	// media files
	"aac":  {"󰝚", Red},
	"flac": {"󰝚", Red},
	"mp3":  {"󰝚", Red},
	"ogg":  {"󰝚", Red},
	"wav":  {"󰝚", Red},

	"avi":  {"󰎁", Red},
	"mkv":  {"󰎁", Red},
	"mov":  {"󰎁", Red},
	"mp4":  {"󰎁", Red},
	"webm": {"󰎁", Red},

	"bmp":  {"󰋩", Blue},
	"gif":  {"󰋩", Blue},
	"heic": {"󰋩", Blue},
	"icns": {"󰋩", Blue},
	"ico":  {"󰋩", Blue},
	"jpeg": {"󰋩", Blue},
	"jpg":  {"󰋩", Blue},
	"png":  {"󰋩", Blue},
	"svg":  {"󰋩", Blue},
	"webp": {"󰋩", Blue},

	// documents
	"doc":  {"󱎒", Blue},
	"docx": {"󱎒", Blue},
	"epub": {"󰂺", Green},
	"mobi": {"󰂺", Green},
	"pdf":  {"", Red},
	"ppt":  {"󱎐", Red},
	"pptx": {"󱎐", Red},
	"xls":  {"󱎏", Green},
	"xlsx": {"󱎏", Green},

	// archive
	"7z":  {"", Green},
	"bz2": {"", Green},
	"gz":  {"", Green},
	"rar": {"", Green},
	"tar": {"", Green},
	"xz":  {"", Green},
	"zip": {"", Green},

	// disc image
	"dmg": {"", White},
	"iso": {"", White},

	// font files
	"otf":   {"", White},
	"ttf":   {"", White},
	"woff":  {"", White},
	"woff2": {"", White},

	// misc
	"txt":   {"󰈙", White},
	"csv":   {"󰸫", White},
	"env":   {"󰒓", Yellow},
	"ipynb": {"", Yellow},
	"json":  {"", Yellow},
	"lock":  {"󰌾", Yellow},
	"md":    {"", White},
	"toml":  {"󰒓", Magenta},
	"xml":   {"󰗀", Magenta},
	"yaml":  {"󰒓", Magenta},
	"yml":   {"󰒓", Magenta},
}

func getStyle(info FileInfo) FileStyle {
	key := "default"

	if info.isDirectory {
		if info.isSymlink && info.isEmpty {
			key = "dir_empty_symlink"
		} else if info.isSymlink {
			key = "dir_symlink"
		} else if info.isEmpty {
			key = "dir_empty"
		} else {
			key = "dir"
		}
	} else {
		if info.isSymlink {
			key = "file_symlink"
		} else if info.isExecutable {
			key = "file_executable"
		} else if info.isEmpty {
			key = "file_empty"
		} else if style, ok := fileStyle[info.extension]; ok {
			return style
		}
	}

	return fileStyle[key]
}

func printIconFilename(info FileInfo) {
	style := getStyle(info)
	fmt.Printf("%s%s %s%s\n", style.Color, style.Icon, Reset, info.name)
}
