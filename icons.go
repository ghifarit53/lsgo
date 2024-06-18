package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// first, check for directory
// then check whether it's a symlink or not
// if not a directory, check whether it's a symlinked file
// else, check whether it's an executable file
// else, check whether it has extension, then do pattern matchin
// if it doesn't have extension, check whether it's a text file
// else, it's likely a binary file

func printIconAndFilenames(fileInfoArray []FileInfo) {
	f, err := os.ReadFile(os.Getenv("HOME") + "/.config/lsi_config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var config Configuration

	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatal(err)
	}

	icon, color := config.General.File.Icon, colors[config.General.File.Color]

	for _, fileInfo := range fileInfoArray {
		if fileInfo.isDirectory {
			// Handle directory
			if fileInfo.isSymlink {
				color = colors[config.General.FolderSymlink.Color]
			} else {
				color = colors[config.General.Folder.Color]
			}
			if fileInfo.isEmpty {
				icon = config.General.FolderEmpty.Icon
			} else {
				icon = config.General.Folder.Icon
			}
		} else {
			// Handle file
			if fileInfo.isSymlink {
				// Symlinked file
				icon = config.General.FileSymlink.Icon
				color = colors[config.General.FileSymlink.Color]
			} else if fileInfo.isExecutable {
				// // Executable file
				icon = config.General.FileExecutable.Icon
				color = colors[config.General.FileExecutable.Color]
			} else if fileInfo.isEmpty {
				// Empty file with any extension
				icon = config.General.FileEmpty.Icon
				color = colors[config.General.FileEmpty.Color]
			} else if fileInfo.isTextFile {
				// Text file
				if fileInfo.extension != "" {
					switch fileInfo.extension {
					case "astro":
						icon = config.Developer.Astro.Icon
						color = colors[config.Developer.Astro.Color]

					case "c":
						icon = config.Developer.C.Icon
						color = colors[config.Developer.C.Color]

					case "cpp":
						icon = config.Developer.Cpp.Icon
						color = colors[config.Developer.Cpp.Color]

					case "csv":
						icon = config.Developer.CSV.Icon
						color = colors[config.Developer.CSV.Color]

					case "css":
						icon = config.Developer.CSS.Icon
						color = colors[config.Developer.CSS.Color]

					case "dart":
						icon = config.Developer.Dart.Icon
						color = colors[config.Developer.Dart.Color]

					case "env":
						icon = config.Developer.Env.Icon
						color = colors[config.Developer.Env.Color]

					case "fnl":
						icon = config.Developer.Fennel.Icon
						color = colors[config.Developer.Fennel.Color]

					case "go":
						icon = config.Developer.Go.Icon
						color = colors[config.Developer.Go.Color]

					case "h":
						icon = config.Developer.HeaderC.Icon
						color = colors[config.Developer.HeaderC.Color]

					case "hpp":
						icon = config.Developer.HeaderCpp.Icon
						color = colors[config.Developer.HeaderCpp.Color]

					case "html":
						icon = config.Developer.HTML.Icon
						color = colors[config.Developer.HTML.Color]

					case "json":
						icon = config.Developer.JSON.Icon
						color = colors[config.Developer.JSON.Color]

					case "java":
						icon = config.Developer.Java.Icon
						color = colors[config.Developer.Java.Color]

					case "js":
						icon = config.Developer.JavaScript.Icon
						color = colors[config.Developer.JavaScript.Color]

					case "jsx":
						icon = config.Developer.JavaScriptReact.Icon
						color = colors[config.Developer.JavaScriptReact.Color]

					case "ipynb":
						icon = config.Developer.JupyterNotebook.Icon
						color = colors[config.Developer.JupyterNotebook.Color]

					case "kt":
						icon = config.Developer.Kotlin.Icon
						color = colors[config.Developer.Kotlin.Color]

					case "lock":
						icon = config.Developer.Lock.Icon
						color = colors[config.Developer.Lock.Color]

					case "lua":
						icon = config.Developer.Lua.Icon
						color = colors[config.Developer.Lua.Color]

					case "md":
						icon = config.Developer.Markdown.Icon
						color = colors[config.Developer.Markdown.Color]

					case "php":
						icon = config.Developer.PHP.Icon
						color = colors[config.Developer.PHP.Color]

					case "perl":
						icon = config.Developer.Perl.Icon
						color = colors[config.Developer.Perl.Color]

					case "py":
						icon = config.Developer.Python.Icon
						color = colors[config.Developer.Python.Color]

					case "r":
						icon = config.Developer.R.Icon
						color = colors[config.Developer.R.Color]

					case "rb":
						icon = config.Developer.R.Icon
						color = colors[config.Developer.Ruby.Color]

					case "rs":
						icon = config.Developer.Rust.Icon
						color = colors[config.Developer.Rust.Color]

					case "sql":
						icon = config.Developer.SQL.Icon
						color = colors[config.Developer.SQL.Color]

					case "sh":
						icon = config.Developer.Shell.Icon
						color = colors[config.Developer.Shell.Color]

					case "swift":
						icon = config.Developer.Swift.Icon
						color = colors[config.Developer.Swift.Color]

					case "toml":
						icon = config.Developer.TOML.Icon
						color = colors[config.Developer.TOML.Color]

					case "ts":
						icon = config.Developer.TypeScript.Icon
						color = colors[config.Developer.TypeScript.Color]

					case "typ":
						icon = config.Developer.Typst.Icon
						color = colors[config.Developer.Typst.Color]

					case "vim":
						icon = config.Developer.Vim.Icon
						color = colors[config.Developer.Vim.Color]

					case "vue":
						icon = config.Developer.Vue.Icon
						color = colors[config.Developer.Vue.Color]

					case "xml":
						icon = config.Developer.XML.Icon
						color = colors[config.Developer.XML.Color]

					case "yaml", "yml":
						icon = config.Developer.YAML.Icon
						color = colors[config.Developer.YAML.Color]

					case "zig":
						icon = config.Developer.Zig.Icon
						color = colors[config.Developer.Zig.Color]

					default:
						icon = config.General.FileText.Icon
						color = colors[config.General.FileText.Color]
					}
				} else {
					icon = config.General.FileText.Icon
					color = colors[config.General.FileText.Color]
				}
			} else {
				if fileInfo.extension != "" {
					switch fileInfo.extension {
					case "mp3", "aac", "flac", "wav", "ogg":
						icon = config.Filetype.Audio.Icon
						color = colors[config.Filetype.Audio.Color]

					case "mp4", "mkv", "mov", "avi", "webm":
						icon = config.Filetype.Video.Icon
						color = colors[config.Filetype.Video.Color]

					case "jpg", "jpeg", "png", "icns", "svg", "webp", "heic":
						icon = config.Filetype.Picture.Icon
						color = colors[config.Filetype.Picture.Color]

					case "epub", "mobi":
						icon = config.Filetype.EBook.Icon
						color = colors[config.Filetype.EBook.Color]

					case "pdf":
						icon = config.Filetype.PDF.Icon
						color = colors[config.Filetype.PDF.Color]

					case "doc", "docx":
						icon = config.Filetype.Word.Icon
						color = colors[config.Filetype.Word.Color]

					case "xls", "xlsx":
						icon = config.Filetype.Excel.Icon
						color = colors[config.Filetype.Excel.Color]

					case "ppt", "pptx":
						icon = config.Filetype.Powerpoint.Icon
						color = colors[config.Filetype.Powerpoint.Color]

					case "dmg", "iso":
						icon = config.Filetype.DiskImage.Icon
						color = colors[config.Filetype.DiskImage.Color]

					case "ttf", "otf", "woff", "woff2":
						icon = config.Filetype.Font.Icon
						color = colors[config.Filetype.Font.Color]

					case "zip", "gz", "rar", "7z", "tar", "bz2", "xz":
						icon = config.Filetype.Archive.Icon
						color = colors[config.Filetype.Archive.Color]
					}
				}
			}
		}

		if fileInfo.size > 0 {
			fmt.Printf("%v%v %v%v (%v)\n", color, icon, colors["reset"], fileInfo.name, formatFileSize(fileInfo.size))
		} else {
			fmt.Printf("%v%v %v%v\n", color, icon, colors["reset"], fileInfo.name)
		}
	}
}
