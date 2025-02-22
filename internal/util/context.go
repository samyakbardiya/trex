package util

type ContextKey string

const (
	// context key for the file data.
	KeyFileData ContextKey = "filedata"

	// example for cli usage
	CliExample = `  trex                   # Run with default sample text
  trex myfile.txt        # Process specific file
  trex /path/to/file.md  # Process file with absolute path`

	// default text content
	DefaultText = `
Lorem ipsum odor amet, consectetuer adipiscing elit. Semper tristique curabitur netus facilisi commodo pellentesque. Dignissim habitant metus massa fermentum aliquam leo praesent vestibulum. Duis et enim ex non conubia leo. Aptent gravida hendrerit odio ultricies cras dolor vulputate placerat? Posuere lacus interdum; ac curae nibh sit vestibulum. Fusce elementum nec sed purus sollicitudin, class ullamcorper purus!

Etiam finibus purus dolor semper eu posuere mi lectus. Conubia lacus augue dolor porttitor leo quisque blandit. Potenti risus maecenas a potenti class velit fringilla mauris. Montes in faucibus gravida luctus aptent iaculis. Ex condimentum curabitur tempor ad at. Elementum fringilla fusce mauris primis porta. Ut adipiscing dis cursus id nec hendrerit efficitur. Sit montes taciti neque; cursus ante venenatis. Sagittis risus ex eget habitant non. Volutpat varius orci aptent; facilisis blandit rhoncus est.

Congue fringilla parturient donec aptent mattis nam. Ac conubia eu efficitur ac aenean non fusce. Penatibus rhoncus cras diam justo primis lobortis. Ad quam ullamcorper vestibulum vulputate senectus. Placerat tristique sollicitudin nisl varius penatibus consequat vivamus. Primis habitant nam libero cubilia, tortor nulla malesuada. Erat elit conubia fusce consectetur, litora blandit dui suscipit. Ultricies habitant magna magnis habitasse et dapibus malesuada. Vestibulum gravida consequat risus cursus, sociosqu dis. Mus primis augue bibendum penatibus ac euismod.

Massa hac conubia cursus elementum tempor laoreet posuere dictum. Molestie non sem pretium vitae orci nec. Inceptos ad imperdiet dis curae pellentesque conubia eget. Non purus etiam senectus consequat vehicula ullamcorper habitasse netus condimentum. Consectetur volutpat vivamus est; integer fames tincidunt mus. Orci tristique ornare odio, potenti sociosqu class ligula consequat. Dui feugiat adipiscing ultrices imperdiet turpis pellentesque magna risus. Cubilia montes litora nibh praesent habitasse. Sollicitudin cras nullam interdum lorem vivamus ex sociosqu primis.

Venenatis pellentesque ultricies hac condimentum; vel enim. Ligula mauris tristique auctor nam elit fames et? Diam fringilla habitant orci nisi convallis nibh velit. Malesuada arcu taciti nisi bibendum ultrices lacus porttitor. Nibh sem praesent rhoncus ultricies tempor commodo orci. Tristique ante lacinia ipsum orci eu nisi. Risus maximus cursus tincidunt cras lorem orci velit dolor hac. Turpis a taciti natoque; sit ut arcu suspendisse lacinia. Ornare convallis mus volutpat etiam pulvinar euismod.
`
)
