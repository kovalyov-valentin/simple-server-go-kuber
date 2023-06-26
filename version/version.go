package version

var (
	// BuildTime - это временная метка момента, когда был собран двоичный файл
	BuildTime = "unset"
	// Commit - это хэш последней фиксации на момент создания двоичного файла
	Commit    = "unset"
	// Release is a semantic version of current build
	Release   = "unset"
)
