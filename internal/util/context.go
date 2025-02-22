package util

type ContextKey string

const (
	// context key for the file data.
	KeyFileData ContextKey = "filedata"

	// example for cli usage
	CliExample = `  trex                   # Run with default sample text
  trex myfile.txt        # Process specific file
  trex /path/to/file.md  # Process file with absolute path`

	// HACK: Not working for longer default text

	// default text content
	DefaultText = `Lorem ipsum odor amet, consectetuer adipiscing elit. Semper tristique curabitur netus facilisi commodo pellentesque. Dignissim habitant metus massa fermentum aliquam leo praesent vestibulum. Duis et enim ex non conubia leo. Aptent gravida hendrerit odio ultricies cras dolor vulputate placerat? Posuere lacus interdum; ac curae nibh sit vestibulum. Fusce elementum nec sed purus sollicitudin, class ullamcorper purus!`
)
