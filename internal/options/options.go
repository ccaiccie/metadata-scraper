package options

import (
	"github.com/spf13/pflag"
)

var metadataScraperOptions = &MetadataScraperOptions{}

func init() {
	metadataScraperOptions.addFlags(pflag.CommandLine)
	pflag.Parse()
}

// GetMetadataScraperOptions returns the pointer of MetadataScraperOptions
func GetMetadataScraperOptions() *MetadataScraperOptions {
	return metadataScraperOptions
}

// MetadataScraperOptions contains frequent command line and application options.
type MetadataScraperOptions struct {
	// Domain is the target domain to search on search engines
	Domain string
	// FileType is the file type to search for on the specified domain
	FileType string
}

func (mso *MetadataScraperOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&mso.Domain, "domain", "nytimes.com", "target domain to search on search engines")
	fs.StringVar(&mso.FileType, "fileType", "docx", "file type to search for on the specified domain")
}
