package malleable

type Profile struct {
	SampleName           string             `parser:"( \"set\" \"sample_name\" @String \";\""`
	SleepTime            int                `parser:"| \"set\" \"sleeptime\" @Number \";\""`
	Jitter               int                `parser:"| \"set\" \"jitter\" @Number \";\""`
	UserAgent            string             `parser:"| \"set\" \"useragent\" @String \";\""`
	DataJitter           int                `parser:"| \"set\" \"data_jitter\" @Number \";\""`
	HostStage            Boolean            `parser:"| \"set\" \"host_stage\" @Boolean \";\""`
	Pipename             string             `parser:"| \"set\" \"pipename\" @String \";\""`
	PipenameStager       string             `parser:"| \"set\" \"pipename_stager\" @String \";\""`
	SMBFrameHeader       string             `parser:"| \"set\" \"smb_frame_header\" @String \";\""`
	TCPPort              int                `parser:"| \"set\" \"tcp_port\" @Number \";\""`
	TCPFrameHeader       string             `parser:"| \"set\" \"tcp_frame_header\" @String \";\""`
	SSHBanner            string             `parser:"| \"set\" \"ssh_banner\" @String \";\""`
	SSHPipename          string             `parser:"| \"set\" \"ssh_pipename\" @String \";\""`
	StealTokenAccessMask string             `parser:"| \"set\" \"steal_token_access_mask\" @String \";\""`
	TasksMaxSize         string             `parser:"| \"set\" \"tasks_max_size\" @String \";\""`
	TasksProxyMaxSize    string             `parser:"| \"set\" \"tasks_proxy_max_size\" @String \";\""`
	TasksDNSProxyMaxSize string             `parser:"| \"set\" \"tasks_dns_proxy_max_size\" @String \";\""`
	HeadersRemove        CommaSeparatedList `parser:"| \"set\" \"headers_remove\" @String \";\""`

	DNSBeacon        []DNSBeacon        `parser:"| \"dns-beacon\" @@"`
	HTTPSCertificate []HTTPSCertificate `parser:"| \"https-certificate\" @@"`
	CodeSigner       CodeSigner         `parser:"| \"code-signer\" \"{\" @@ \"}\""`
	HTTPConfig       HTTPConfig         `parser:"| \"http-config\" \"{\" @@ \"}\""`
	HTTPGet          []HTTPGet          `parser:"| \"http-get\" @@"`
	HTTPPost         []HTTPPost         `parser:"| \"http-post\" @@"`
	HTTPStager       []HTTPStager       `parser:"| \"http-stager\" @@"`
	Stage            Stage              `parser:"| \"stage\" \"{\" @@ \"}\""`
	ProcessInject    ProcessInject      `parser:"| \"process-inject\" \"{\" @@ \"}\""`
	PostEx           PostEx             `parser:"| \"post-ex\" \"{\" @@ \"}\" )*"`
}

func (d Profile) String() string {
	return printStruct(0, d)
}

type DNSBeacon struct {
	Name string `parser:"@String? \"{\""`

	DNSIdle          string `parser:"( \"set\" \"dns_idle\" @String \";\""`
	DNSMaxTXT        int    `parser:"| \"set\" \"dns_max_txt\" @Number \";\""`
	DNSSleep         int    `parser:"| \"set\" \"dns_sleep\" @Number \";\""`
	DNSTTL           int    `parser:"| \"set\" \"dns_ttl\" @Number \";\""`
	MaxDNS           int    `parser:"| \"set\" \"maxdns\" @Number \";\""`
	DNSStagerPrepend string `parser:"| \"set\" \"dns_stager_prepend\" @String \";\""`
	DNSStagerSubhost string `parser:"| \"set\" \"dns_stager_subhost\" @String \";\""`
	Beacon           string `parser:"| \"set\" \"beacon\" @String \";\""`
	GetA             string `parser:"| \"set\" \"get_A\" @String \";\""`
	GetAAAA          string `parser:"| \"set\" \"get_AAAA\" @String \";\""`
	GetTXT           string `parser:"| \"set\" \"get_TXT\" @String \";\""`
	PutMetadata      string `parser:"| \"set\" \"put_metadata\" @String \";\""`
	PutOutput        string `parser:"| \"set\" \"put_output\" @String \";\""`
	NSResponse       string `parser:"| \"set\" \"ns_response\" @String \";\")* \"}\""`
}

func (b DNSBeacon) String() string {
	return printNamed(0, "dns-beacon", b.Name, b)
}

type HTTPSCertificate struct {
	Name string `parser:"@String? \"{\""`

	Keystore string `parser:"( \"set\" \"keystore\" @String \";\""`
	Password string `parser:"| \"set\" \"password\" @String \";\""`

	C        string `parser:"| \"set\" \"C\" @String \";\""`
	CN       string `parser:"| \"set\" \"CN\" @String \";\""`
	L        string `parser:"| \"set\" \"L\" @String \";\""`
	O        string `parser:"| \"set\" \"O\" @String \";\""`
	OU       string `parser:"| \"set\" \"OU\" @String \";\""`
	ST       string `parser:"| \"set\" \"ST\" @String \";\""`
	Validity int    `parser:"| \"set\" \"validity\" @Number \";\")* \"}\""`
}

func (b HTTPSCertificate) String() string {
	return printNamed(0, "https-certificate", b.Name, b)
}

type CodeSigner struct {
	Keystore        string  `parser:"( \"set\" \"keystore\" @String \";\""`
	Password        string  `parser:"| \"set\" \"password\" @String \";\""`
	Alias           string  `parser:"| \"set\" \"alias\" @String \";\""`
	DigestAlgorithm string  `parser:"| \"set\" \"digest_algorithm\" @String \";\""`
	Timestamp       Boolean `parser:"| \"set\" \"timestamp\" @Boolean \";\""`
	TimestampURL    string  `parser:"| \"set\" \"timestamp_url\" @String \";\" )*"`
}

func (b CodeSigner) String() string {
	return printUnnamed(0, "code-signer", b)
}

type HTTPConfig struct {
	HeadersOrder       CommaSeparatedList `parser:"( \"set\" \"headers\" @String \";\""`
	Headers            []Header           `parser:"| \"header\" @@ \";\""`
	TrustXForwardedFor Boolean            `parser:"| \"set\" \"trust_x_forwarded_for\" @Boolean \";\""`
	BlockUserAgents    CommaSeparatedList `parser:"| \"set\" \"block_useragents\" @String \";\")*"`
}

func (b HTTPConfig) String() string {
	return printUnnamed(0, "http-config", b)
}

type HTTPGet struct {
	Name string `parser:"@String? \"{\""`

	Verb   string        `parser:"( \"set\" \"verb\" @String \";\""`
	URI    URIs          `parser:"| \"set\" \"uri\" @String \";\""`
	Client HTTPGetClient `parser:"| \"client\" \"{\" @@ \"}\""`
	Server HTTPServer    `parser:"| \"server\" \"{\" @@ \"}\" )* \"}\""`
}

func (b HTTPGet) String() string {
	return printNamed(0, "http-get", b.Name, b)
}

type HTTPGetClient struct {
	Headers    []Header    `parser:"( \"header\" @@ \";\""`
	Parameters []Parameter `parser:"| \"parameter\" @@ \";\""`
	Metadata   []Function  `parser:"| \"metadata\" \"{\" @@* \"}\" )*"`
}

func (b HTTPGetClient) String() string {
	return printUnnamed(1, "client", b)
}

type HTTPPost struct {
	Name string `parser:"@String? \"{\""`

	Verb   string         `parser:"( \"set\" \"verb\" @String \";\""`
	URI    URIs           `parser:"| \"set\" \"uri\" @String \";\""`
	Client HTTPPostClient `parser:"| \"client\" \"{\" @@ \"}\""`
	Server HTTPServer     `parser:"| \"server\" \"{\" @@ \"}\" )* \"}\""`
}

func (b HTTPPost) String() string {
	return printNamed(0, "http-post", b.Name, b)
}

type HTTPPostClient struct {
	Headers    []Header    `parser:"( \"header\" @@ \";\""`
	Parameters []Parameter `parser:"| \"parameter\" @@ \";\""`
	Output     []Function  `parser:"| \"output\" \"{\" @@* \"}\""`
	ID         []Function  `parser:"| \"id\" \"{\" @@* \"}\" )*"`
}

func (b HTTPPostClient) String() string {
	return printUnnamed(1, "client", b)
}

type HTTPStager struct {
	Name string `parser:"@String? \"{\""`

	URIx86 URIs             `parser:"( \"set\" \"uri_x86\" @String \";\""`
	URIx64 URIs             `parser:"| \"set\" \"uri_x64\" @String \";\""`
	Client HTTPStagerClient `parser:"| \"client\" \"{\" @@ \"}\""`
	Server HTTPServer       `parser:"| \"server\" \"{\" @@ \"}\" )* \"}\""`
}

func (b HTTPStager) String() string {
	return printNamed(0, "http-stager", b.Name, b)
}

type HTTPStagerClient struct {
	Headers    []Header    `parser:"( \"header\" @@ \";\""`
	Parameters []Parameter `parser:"| \"parameter\" @@ \";\" )*"`
}

func (b HTTPStagerClient) String() string {
	return printUnnamed(1, "client", b)
}

type HTTPServer struct {
	Headers []Header   `parser:"( \"header\" @@ \";\""`
	Output  []Function `parser:"| \"output\" \"{\" @@* \"}\" )*"`
}

func (b HTTPServer) String() string {
	return printUnnamed(1, "server", b)
}

type Stage struct {
	Checksum      int     `parser:"( \"set\" \"checksum\" @Number \";\""`
	CompileTime   string  `parser:"| \"set\" \"compile_time\" @String \";\""`
	EntryPoint    int     `parser:"| \"set\" \"entry_point\" @Number \";\""`
	ImageSizeX86  int     `parser:"| \"set\" \"image_size_x86\" @Number \";\""`
	ImageSizeX64  int     `parser:"| \"set\" \"image_size_x64\" @Number \";\""`
	Name          string  `parser:"| \"set\" \"name\" @String \";\""`
	RichHeader    string  `parser:"| \"set\" \"rich_header\" @String \";\""`
	UseRWX        Boolean `parser:"| \"set\" \"userwx\" @Boolean \";\""`
	Cleanup       Boolean `parser:"| \"set\" \"cleanup\" @Boolean \";\""`
	SleepMask     Boolean `parser:"| \"set\" \"sleep_mask\" @Boolean \";\""`
	StompPE       Boolean `parser:"| \"set\" \"stomppe\" @Boolean \";\""`
	Obfuscate     Boolean `parser:"| \"set\" \"obfuscate\" @Boolean \";\""`
	Allocator     string  `parser:"| \"set\" \"allocator\" @String \";\""`
	MagicMZX86    string  `parser:"| \"set\" \"magic_mz_x86\" @String \";\""`
	MagicMZX64    string  `parser:"| \"set\" \"magic_mz_x64\" @String \";\""`
	MagicPE       string  `parser:"| \"set\" \"magic_pe\" @String \";\""`
	SmartInject   Boolean `parser:"| \"set\" \"smartinject\" @Boolean \";\""`
	ModuleX86     string  `parser:"| \"set\" \"module_x86\" @String \";\""`
	ModuleX64     string  `parser:"| \"set\" \"module_x64\" @String \";\""`
	SyscallMethod string  `parser:"| \"set\" \"syscall_method\" @String \";\""`

	TransformX86 []Function `parser:"| \"transform-x86\" \"{\" @@* \"}\""`
	TransformX64 []Function `parser:"| \"transform-x64\" \"{\" @@* \"}\""`

	Data      []Data    `parser:"| \"data\" @String \";\""`
	Strings   []String  `parser:"| \"string\" @String \";\""`
	SwtringsW []StringW `parser:"| \"stringw\" @String \";\" )*"`
}

func (b Stage) String() string {
	return printUnnamed(0, "stage", b)
}

type ProcessInject struct {
	Allocator      string  `parser:"( \"set\" \"allocator\" @String \";\""`
	BOFAllocator   string  `parser:"| \"set\" \"bof_allocator\" @String \";\""`
	BOFReuseMemory Boolean `parser:"| \"set\" \"bof_reuse_memory\" @Boolean \";\""`
	MinAlloc       int     `parser:"| \"set\" \"min_alloc\" @Number \";\""`
	UseRWX         Boolean `parser:"| \"set\" \"userwx\" @Boolean \";\""`
	StartRWX       Boolean `parser:"| \"set\" \"startrwx\" @Boolean \";\""`

	TransformX86 []Function `parser:"| \"transform-x86\" \"{\" @@* \"}\""`
	TransformX64 []Function `parser:"| \"transform-x64\" \"{\" @@* \"}\""`

	Execute []Function `parser:"| \"execute\" \"{\" @@* \"}\" )*"`
}

func (b ProcessInject) String() string {
	return printUnnamed(0, "process-inject", b)
}

type PostEx struct {
	SpawnToX86  string  `parser:"( \"set\" \"spawnto_x86\" @String \";\""`
	SpawnToX64  string  `parser:"| \"set\" \"spawnto_x64\" @String \";\""`
	Obfuscate   Boolean `parser:"| \"set\" \"obfuscate\" @Boolean \";\""`
	SmartInject Boolean `parser:"| \"set\" \"smartinject\" @Boolean \";\""`
	AmsiDisable Boolean `parser:"| \"set\" \"amsi_disable\" @Boolean \";\""`
	ThreadHint  string  `parser:"| \"set\" \"thread_hint\" @String \";\""`
	PipeName    string  `parser:"| \"set\" \"pipename\" @String \";\""`
	Keylogger   string  `parser:"| \"set\" \"keylogger\" @String \";\" )*"`
}

func (b PostEx) String() string {
	return printUnnamed(0, "post-ex", b)
}
