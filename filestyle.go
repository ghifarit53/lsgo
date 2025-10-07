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

	// files
	"file_empty":      {"󰈤", White},
	"file_symlink":    {"󰌹", Cyan},
	"file_executable": {"", Red},

	// programming languages
	"c":     {"", Magenta},
	"cpp":   {"", Blue},
	"cs":    {"", Magenta},
	"css":   {"", Green},
	"dart":  {"", Cyan},
	"ex":    {"", Magenta},
	"fish":  {"󰐤", Magenta},
	"gd":    {"", Blue},
	"go":    {"", Blue},
	"h":     {"󰬏", Blue},
	"hpp":   {"󰬏", Blue},
	"hs":    {"", Magenta},
	"html":  {"", Green},
	"java":  {"", Red},
	"js":    {"", Yellow},
	"jsx":   {"", Yellow},
	"kt":    {"", Magenta},
	"kts":   {"", Red},
	"lua":   {"", Blue},
	"m":     {"󰬔", Blue},
	"perl":  {"", Blue},
	"php":   {"", Magenta},
	"py":    {"", Yellow},
	"r":     {"", Blue},
	"rb":    {"", Red},
	"rs":    {"", Blue},
	"sass":  {"", Red},
	"scala": {"", Red},
	"sh":    {"󰐤", Magenta},
	"swift": {"󰛥", Red},
	"ts":    {"", Blue},
	"tsx":   {"", Blue},
	"vue":   {"", Green},
	"zsh":   {"󰐤", Magenta},

	// audio files
	"aac": {"󰝚", Red},
	"m4a": {"󰝚", Red},
	"mp3": {"󰝚", Red},
	"ogg": {"󰝚", Red},

	"flac": {"󱑽", Red},
	"opus": {"󱑽", Red},
	"wav":  {"󱑽", Red},

	// video files
	"3gp":  {"󰎁", Red},
	"flv":  {"󰎁", Red},
	"wmv":  {"󰎁", Red},
	"avi":  {"󰎁", Red},
	"mkv":  {"󰎁", Red},
	"mov":  {"󰎁", Red},
	"mp4":  {"󰎁", Red},
	"webm": {"󰎁", Red},

	// image files
	"bmp":  {"󰋩", Blue},
	"gif":  {"󰵸", Blue},
	"heic": {"󰋩", Blue},
	"jpeg": {"󰋩", Blue},
	"jpg":  {"󰋩", Blue},
	"png":  {"󰋩", Blue},
	"svg":  {"󰋩", Blue},
	"tif":  {"󰋩", Blue},
	"tiff": {"󰋩", Blue},
	"webp": {"󰋩", Blue},

	// icon files
	"icns": {"", Blue},
	"ico":  {"", Blue},

	// documents
	"doc":  {"󰈬", Blue},
	"docx": {"󰈬", Blue},
	"epub": {"󰂺", Green},
	"md":   {"", White},
	"mobi": {"󰂺", Green},
	"odp":  {"󰈧", Red},
	"ods":  {"󰈛", Green},
	"odt":  {"󰈬", Blue},
	"pdf":  {"", Red},
	"ppt":  {"󰈧", Red},
	"pptx": {"󰈧", Red},
	"rtf":  {"󱘍", White},
	"tex":  {"", White},
	"xls":  {"󰈛", Green},
	"xlsx": {"󰈛", Green},

	// archives
	"7z":  {"󰗄", Green},
	"bz2": {"󰗄", Green},
	"gz":  {"󰗄", Green},
	"rar": {"󰗄", Green},
	"tar": {"󰗄", Green},
	"xz":  {"󰗄", Green},
	"zip": {"󰗄", Green},

	// design graphics
	"ai":    {"", Red},
	"blend": {"", Red},
	"fbx":   {"󰆦", Cyan},
	"psd":   {"", Blue},
	"stl":   {"󰆦", Cyan},

	// disc image
	"dmg": {"󰻂", White},
	"iso": {"󰻂", White},

	// font files
	"otf":   {"", White},
	"ttf":   {"", White},
	"woff":  {"", White},
	"woff2": {"", White},

	// console games
	"gb":  {"󰺵", Blue},
	"gba": {"󰺵", Blue},
	"gbc": {"󰺵", Blue},

	// configuration files
	"ini":   {"󰒓", White},
	"plist": {"󰒓", White},
	"toml":  {"󰒓", White},
	"xml":   {"󰗀", White},
	"yaml":  {"󰒓", White},
	"yml":   {"󰒓", White},
	"env":   {"󰇽", Yellow},

	// text data
	"csv":  {"󰸫", White},
	"json": {"", Yellow},
	"lock": {"󰌾", Yellow},

	// misc
	"apk":   {"", Green},
	"exe":   {"", Blue},
	"fit":   {"󰥛", Red},
	"gpx":   {"󰍒", Green},
	"ics":   {"", Red},
	"ipynb": {"󰺂", Blue},
	"jar":   {"", Blue},
	"pkg":   {"󰏓", Yellow},
	"sav":   {"󰆓", Blue},
	"tcx":   {"󰍒", Green},
	"txt":   {"󰈙", White},
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
		} else if isText, err := isTextFile(info.name); err == nil && isText {
			key = "txt"
		}
	}

	return fileStyle[key]
}

func printIconFilename(info FileInfo) {
	style := getStyle(info)
	fmt.Printf("%s%s %s%s\n", style.Color, style.Icon, Reset, info.name)
}
