package main

type FileInfo struct {
	name         string
	extension    string
	size         int64
	isDirectory  bool
	isEmpty      bool
	isSymlink    bool
	isTextFile   bool
	isExecutable bool
}

type ByDirectoryAndName []FileInfo

func (a ByDirectoryAndName) Len() int      { return len(a) }
func (a ByDirectoryAndName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDirectoryAndName) Less(i, j int) bool {
	if a[i].isDirectory && !a[j].isDirectory {
		return true
	} else if !a[i].isDirectory && a[j].isDirectory {
		return false
	} else {
		return a[i].name < a[j].name
	}
}

type IconColor struct {
	Icon  string
	Color string
}

type Configuration struct {
	General   GeneralIcons   `yaml:"general"`
	Developer DeveloperIcons `yaml:"developer"`
	Filetype  FiletypeIcons
}

type GeneralIcons struct {
	File               IconColor `yaml:"file"`
	FileEmpty          IconColor `yaml:"file_empty"`
	FileExecutable     IconColor `yaml:"file_executable"`
	FileSymlink        IconColor `yaml:"file_symlink"`
	FileText           IconColor `yaml:"file_text"`
	Folder             IconColor `yaml:"folder"`
	FolderEmpty        IconColor `yaml:"folder_empty"`
	FolderSymlink      IconColor `yaml:"folder_symlink"`
	FolderSymlinkEmpty IconColor `yaml:"folder_symlink_empty"`
}

type DeveloperIcons struct {
	Astro           IconColor `yaml:"astro"`
	C               IconColor `yaml:"c"`
	CSS             IconColor `yaml:"css"`
	CSV             IconColor `yaml:"csv"`
	Cpp             IconColor `yaml:"cpp"`
	Dart            IconColor `yaml:"dart"`
	Env             IconColor `yaml:"env"`
	Fennel          IconColor `yaml:"fennel"`
	Go              IconColor `yaml:"go"`
	HeaderC         IconColor `yaml:"h"`
	HeaderCpp       IconColor `yaml:"hpp"`
	HTML            IconColor `yaml:"html"`
	JSON            IconColor `yaml:"json"`
	Java            IconColor `yaml:"java"`
	JavaScript      IconColor `yaml:"javascript"`
	JavaScriptReact IconColor `yaml:"javascript_react"`
	JupyterNotebook IconColor `yaml:"jupyter_notebook"`
	Kotlin          IconColor `yaml:"kotlin"`
	Lock            IconColor `yaml:"lock"`
	Lua             IconColor `yaml:"lua"`
	Markdown        IconColor `yaml:"markdown"`
	PHP             IconColor `yaml:"php"`
	Perl            IconColor `yaml:"perl"`
	Python          IconColor `yaml:"python"`
	R               IconColor `yaml:"r"`
	Ruby            IconColor `yaml:"ruby"`
	Rust            IconColor `yaml:"rust"`
	SQL             IconColor `yaml:"sql"`
	Shell           IconColor `yaml:"sh"`
	Swift           IconColor `yaml:"swift"`
	TOML            IconColor `yaml:"toml"`
	TypeScript      IconColor `yaml:"typescript"`
	Typst           IconColor `yaml:"typst"`
	Vim             IconColor `yaml:"vim"`
	Vue             IconColor `yaml:"vue"`
	XML             IconColor `yaml:"xml"`
	YAML            IconColor `yaml:"yaml"`
	Zig             IconColor `yaml:"zig"`
}

type FiletypeIcons struct {
	Picture    IconColor `yaml:"picture"`
	Video      IconColor `yaml:"video"`
	Audio      IconColor `yaml:"audio"`
	EBook      IconColor `yaml:"ebook"`
	PDF        IconColor `yaml:"pdf"`
	Word       IconColor `yaml:"word"`
	Excel      IconColor `yaml:"excel"`
	Powerpoint IconColor `yaml:"powerpoint"`
	Archive    IconColor `yaml:"archive"`
	Font       IconColor `yaml:"font"`
	DiskImage  IconColor `yaml:"disk_image"`
}
