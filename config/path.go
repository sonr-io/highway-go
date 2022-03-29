package config

type filePathOptType int

// Option Type Enum
const (
	filePathOptionTypeSuffix    filePathOptType = iota // Suffix
	filePathOptionTypePrefix                           // Prefix
	filePathOptionTypeReplace                          // Replace
	filePathOptionTypeSeparator                        // Separator
)

// FilePathOption is a function option for FilePath.
type FilePathOption interface {
	Apply() *filePathOptions
}

// filePathOpt is a struct that implements FilePathOption.
type filePathOpt struct {
	FilePathOption
	value string
	filePathOptType
}

// Apply returns the value set for the specified option.
func (fio *filePathOpt) Apply() *filePathOptions {
	if fio.filePathOptType == filePathOptionTypeSuffix {
		return &filePathOptions{Suffix: fio.value}
	} else if fio.filePathOptType == filePathOptionTypePrefix {
		return &filePathOptions{Prefix: fio.value}
	} else if fio.filePathOptType == filePathOptionTypeReplace {
		return &filePathOptions{Replace: fio.value}
	} else if fio.filePathOptType == filePathOptionTypeSeparator {
		return &filePathOptions{Separator: fio.value}
	}
	return nil
}

// FilePathOptFunc is a function that returns a FilePathOption.
type FilePathOptFunc func(path string) FilePathOption

// WithSuffix sets the suffix for the file path.
var WithSuffix FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypeSuffix,
		value:           path,
	}
}

// WithPrefix sets the prefix for the file path.
var WithPrefix FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypePrefix,
		value:           path,
	}
}

// WithReplace sets the replace string for the file path.
var WithReplace FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypeReplace,
		value:           path,
	}
}

// WithSeparator sets the separator for the file path.
var WithSeparator FilePathOptFunc = func(path string) FilePathOption {
	return &filePathOpt{
		filePathOptType: filePathOptionTypeSeparator,
		value:           path,
	}
}

// filePathOptions is a struct for holding file path options.
type filePathOptions struct {
	// Options
	Suffix    string // Add Suffix to file name
	Prefix    string // Add Prefix to file name
	Replace   string // Replace filename with this string
	Separator string // Default is "-"

	// Properties
	fileName  string
	baseName  string
	extension string

	// Internal
	suffixSet    bool
	prefixSet    bool
	replaceSet   bool
	separatorSet bool
}
