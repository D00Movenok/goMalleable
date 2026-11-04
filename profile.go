package malleable

type Profile struct {
	SampleName           string             `parser:"( \"set\" \"sample_name\" @String \";\""`
	SleepTime            int                `parser:"| \"set\" \"sleeptime\" @String \";\""`
	Jitter               int                `parser:"| \"set\" \"jitter\" @String \";\""`
	Sleep                string             `parser:"| \"set\" \"sleep\" @String \";\""`
	UserAgent            string             `parser:"| \"set\" \"useragent\" @String \";\""`
	DataJitter           int                `parser:"| \"set\" \"data_jitter\" @String \";\""`
	CheckinDelay         int                `parser:"| \"set\" \"checkin_delay\" @String \";\""`
	HostStage            Boolean            `parser:"| \"set\" \"host_stage\" @String \";\""`
	Pipename             string             `parser:"| \"set\" \"pipename\" @String \";\""`
	PipenameStager       string             `parser:"| \"set\" \"pipename_stager\" @String \";\""`
	SMBFrameHeader       string             `parser:"| \"set\" \"smb_frame_header\" @String \";\""`
	TCPPort              int                `parser:"| \"set\" \"tcp_port\" @String \";\""`
	TCPFrameHeader       string             `parser:"| \"set\" \"tcp_frame_header\" @String \";\""`
	SSHBanner            string             `parser:"| \"set\" \"ssh_banner\" @String \";\""`
	SSHPipename          string             `parser:"| \"set\" \"ssh_pipename\" @String \";\""`
	StealTokenAccessMask int                `parser:"| \"set\" \"steal_token_access_mask\" @String \";\""`
	TasksMaxSize         int                `parser:"| \"set\" \"tasks_max_size\" @String \";\""`
	TasksProxyMaxSize    int                `parser:"| \"set\" \"tasks_proxy_max_size\" @String \";\""`
	TasksDNSProxyMaxSize int                `parser:"| \"set\" \"tasks_dns_proxy_max_size\" @String \";\""`
	HeadersRemove        CommaSeparatedList `parser:"| \"set\" \"headers_remove\" @String \";\""`

	DNSBeacon        []DNSBeacon        `parser:"| \"dns-beacon\" @@"`
	HTTPBeacon       []HTTPBeacon       `parser:"| \"http-beacon\" @@"`
	HTTPSCertificate []HTTPSCertificate `parser:"| \"https-certificate\" @@"`
	CodeSigner       CodeSigner         `parser:"| \"code-signer\" \"{\" @@ \"}\""`
	HTTPConfig       HTTPConfig         `parser:"| \"http-config\" \"{\" @@ \"}\""`
	HTTPGet          []HTTPGet          `parser:"| \"http-get\" @@"`
	HTTPPost         []HTTPPost         `parser:"| \"http-post\" @@"`
	HTTPStager       []HTTPStager       `parser:"| \"http-stager\" @@"`
	Stage            Stage              `parser:"| \"stage\" \"{\" @@ \"}\""`
	ProcessInject    ProcessInject      `parser:"| \"process-inject\" \"{\" @@ \"}\""`
	PostEx           PostEx             `parser:"| \"post-ex\" \"{\" @@ \"}\""`
	HostProfiles     HTTPHostProfiles   `parser:"| \"http-host-profiles\" \"{\" @@ \"}\" )*"`
}

func (d Profile) String() string {
	return printStruct(0, d)
}

type DNSBeacon struct {
	Name string `parser:"@String? \"{\""`

	DNSIdle          string `parser:"( \"set\" \"dns_idle\" @String \";\""`
	DNSMaxTXT        int    `parser:"| \"set\" \"dns_max_txt\" @String \";\""`
	DNSSleep         int    `parser:"| \"set\" \"dns_sleep\" @String \";\""`
	DNSTTL           int    `parser:"| \"set\" \"dns_ttl\" @String \";\""`
	MaxDNS           int    `parser:"| \"set\" \"maxdns\" @String \";\""`
	DNSStagerPrepend string `parser:"| \"set\" \"dns_stager_prepend\" @String \";\""`
	DNSStagerSubhost string `parser:"| \"set\" \"dns_stager_subhost\" @String \";\""`
	Beacon           string `parser:"| \"set\" \"beacon\" @String \";\""`
	GetA             string `parser:"| \"set\" \"get_A\" @String \";\""`
	GetAAAA          string `parser:"| \"set\" \"get_AAAA\" @String \";\""`
	GetTXT           string `parser:"| \"set\" \"get_TXT\" @String \";\""`
	PutMetadata      string `parser:"| \"set\" \"put_metadata\" @String \";\""`
	PutOutput        string `parser:"| \"set\" \"put_output\" @String \";\""`
	NSResponse       string `parser:"| \"set\" \"ns_response\" @String \";\""`
	CommMode         string `parser:"| \"set\" \"comm_mode\" @String \";\""`

	DNSOverHTTPS []DNSOverHTTPS `parser:"| \"dns-over-https\" @@ )* \"}\""`
}

func (b DNSBeacon) String() string {
	return printNamed(0, "dns-beacon", b.Name, b)
}

type DNSOverHTTPS struct {
	Name string `parser:"@Ident? \"{\""`

	DohVerb        string             `parser:"( \"set\" \"doh_verb\" @String \";\""`
	DohUserAgent   string             `parser:"| \"set\" \"doh_useragent\" @String \";\""`
	DohProxyServer string             `parser:"| \"set\" \"doh_proxy_server\" @String \";\""`
	DohServer      CommaSeparatedList `parser:"| \"set\" \"doh_server\" @String \";\""`
	DohAccept      string             `parser:"| \"set\" \"doh_accept\" @String \";\""`
	Headers        []Header           `parser:"| \"header\" @@ \";\" )* \"}\""`
}

func (b DNSOverHTTPS) String() string {
	return printUnnamed(1, "dns-over-https", b)
}

type HTTPBeacon struct {
	Name string `parser:"@String? \"{\""`

	Library            string  `parser:"( \"set\" \"library\" @String \";\""`
	DataRequired       Boolean `parser:"| \"set\" \"data_required\" @String \";\""`
	DataRequiredLength string  `parser:"| \"set\" \"data_required_length\" @String \";\" )* \"}\""`
}

func (b HTTPBeacon) String() string {
	return printNamed(0, "http-beacon", b.Name, b)
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
	Validity int    `parser:"| \"set\" \"validity\" @String \";\")* \"}\""`
}

func (b HTTPSCertificate) String() string {
	return printNamed(0, "https-certificate", b.Name, b)
}

type CodeSigner struct {
	Keystore        string  `parser:"( \"set\" \"keystore\" @String \";\""`
	Password        string  `parser:"| \"set\" \"password\" @String \";\""`
	Alias           string  `parser:"| \"set\" \"alias\" @String \";\""`
	DigestAlgorithm string  `parser:"| \"set\" \"digest_algorithm\" @String \";\""`
	Timestamp       Boolean `parser:"| \"set\" \"timestamp\" @String \";\""`
	TimestampURL    string  `parser:"| \"set\" \"timestamp_url\" @String \";\" )*"`
}

func (b CodeSigner) String() string {
	return printUnnamed(0, "code-signer", b)
}

type HTTPConfig struct {
	HeadersOrder       CommaSeparatedList `parser:"( \"set\" \"headers\" @String \";\""`
	Headers            []Header           `parser:"| \"header\" @@ \";\""`
	TrustXForwardedFor Boolean            `parser:"| \"set\" \"trust_x_forwarded_for\" @String \";\""`
	BlockUserAgents    CommaSeparatedList `parser:"| \"set\" \"block_useragents\" @String \";\""`
	AllowUserAgents    CommaSeparatedList `parser:"| \"set\" \"allow_useragents\" @String \";\")*"`
}

func (b HTTPConfig) String() string {
	return printUnnamed(0, "http-config", b)
}

type HTTPGet struct {
	Name string `parser:"@String? \"{\""`

	Verb   string             `parser:"( \"set\" \"verb\" @String \";\""`
	URI    SpaceSeparatedList `parser:"| \"set\" \"uri\" @String \";\""`
	Client HTTPGetClient      `parser:"| \"client\" \"{\" @@ \"}\""`
	Server HTTPServer         `parser:"| \"server\" \"{\" @@ \"}\" )* \"}\""`
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

	Verb                   string             `parser:"( \"set\" \"verb\" @String \";\""`
	URI                    SpaceSeparatedList `parser:"| \"set\" \"uri\" @String \";\""`
	ClientMaxPostPostSize  int                `parser:"| \"set\" \"client_max_post_post_size\" @String \";\""`
	ClientMaxPostGetSize   int                `parser:"| \"set\" \"client_max_post_get_size\" @String \";\""`
	ClientMaxPostGetPacket int                `parser:"| \"set\" \"client_max_post_get_packet\" @String \";\""`
	Client                 HTTPPostClient     `parser:"| \"client\" \"{\" @@ \"}\""`
	Server                 HTTPServer         `parser:"| \"server\" \"{\" @@ \"}\" )* \"}\""`
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

	URIx86 SpaceSeparatedList `parser:"( \"set\" \"uri_x86\" @String \";\""`
	URIx64 SpaceSeparatedList `parser:"| \"set\" \"uri_x64\" @String \";\""`
	Client HTTPStagerClient   `parser:"| \"client\" \"{\" @@ \"}\""`
	Server HTTPServer         `parser:"| \"server\" \"{\" @@ \"}\" )* \"}\""`
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
	Checksum      int     `parser:"( \"set\" \"checksum\" @String \";\""`
	CompileTime   string  `parser:"| \"set\" \"compile_time\" @String \";\""`
	EntryPoint    int     `parser:"| \"set\" \"entry_point\" @String \";\""`
	ImageSizeX86  int     `parser:"| \"set\" \"image_size_x86\" @String \";\""`
	ImageSizeX64  int     `parser:"| \"set\" \"image_size_x64\" @String \";\""`
	Name          string  `parser:"| \"set\" \"name\" @String \";\""`
	RichHeader    string  `parser:"| \"set\" \"rich_header\" @String \";\""`
	UseRWX        Boolean `parser:"| \"set\" \"userwx\" @String \";\""`
	Cleanup       Boolean `parser:"| \"set\" \"cleanup\" @String \";\""`
	SleepMask     Boolean `parser:"| \"set\" \"sleep_mask\" @String \";\""`
	StompPE       Boolean `parser:"| \"set\" \"stomppe\" @String \";\""`
	Obfuscate     Boolean `parser:"| \"set\" \"obfuscate\" @String \";\""`
	Allocator     string  `parser:"| \"set\" \"allocator\" @String \";\""`
	MagicMZX86    string  `parser:"| \"set\" \"magic_mz_x86\" @String \";\""`
	MagicMZX64    string  `parser:"| \"set\" \"magic_mz_x64\" @String \";\""`
	MagicPE       string  `parser:"| \"set\" \"magic_pe\" @String \";\""`
	SmartInject   Boolean `parser:"| \"set\" \"smartinject\" @String \";\""`
	ModuleX86     string  `parser:"| \"set\" \"module_x86\" @String \";\""`
	ModuleX64     string  `parser:"| \"set\" \"module_x64\" @String \";\""`
	SyscallMethod string  `parser:"| \"set\" \"syscall_method\" @String \";\""`
	DataStoreSize int     `parser:"| \"set\" \"data_store_size\" @String \";\""`
	EafBypass     Boolean `parser:"| \"set\" \"eaf_bypass\" @String \";\""`

	RdllLoader         string  `parser:"| \"set\" \"rdll_loader\" @String \";\""`
	RdllUseSyscalls    Boolean `parser:"| \"set\" \"rdll_use_syscalls\" @String \";\""`
	CopyPEHeader       Boolean `parser:"| \"set\" \"copy_pe_header\" @String \";\""`
	RdllUseDriploading Boolean `parser:"| \"set\" \"rdll_use_driploading\" @String \";\""`
	RdllDriploadDelay  int     `parser:"| \"set\" \"rdll_dripload_delay\" @String \";\""`

	BeaconGate []Function `parser:"| \"beacon_gate\" \"{\" @@* \"}\""`

	TransformX86 []Function `parser:"| \"transform-x86\" \"{\" @@* \"}\""`
	TransformX64 []Function `parser:"| \"transform-x64\" \"{\" @@* \"}\""`

	TransformObfuscate []Function `parser:"| \"transform-obfuscate\" \"{\" @@* \"}\""`

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
	BOFReuseMemory Boolean `parser:"| \"set\" \"bof_reuse_memory\" @String \";\""`
	MinAlloc       int     `parser:"| \"set\" \"min_alloc\" @String \";\""`
	UseRWX         Boolean `parser:"| \"set\" \"userwx\" @String \";\""`
	StartRWX       Boolean `parser:"| \"set\" \"startrwx\" @String \";\""`
	UseDriploading Boolean `parser:"| \"set\" \"use_driploading\" @String \";\""`
	DriploadDelay  int     `parser:"| \"set\" \"dripload_delay\" @String \";\""`

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
	Obfuscate   Boolean `parser:"| \"set\" \"obfuscate\" @String \";\""`
	SmartInject Boolean `parser:"| \"set\" \"smartinject\" @String \";\""`
	AmsiDisable Boolean `parser:"| \"set\" \"amsi_disable\" @String \";\""`
	Cleanup     Boolean `parser:"| \"set\" \"cleanup\" @String \";\""`
	ThreadHint  string  `parser:"| \"set\" \"thread_hint\" @String \";\""`
	PipeName    string  `parser:"| \"set\" \"pipename\" @String \";\""`
	Keylogger   string  `parser:"| \"set\" \"keylogger\" @String \";\""`

	TransformX86 []Function `parser:"| \"transform-x86\" \"{\" @@* \"}\""`
	TransformX64 []Function `parser:"| \"transform-x64\" \"{\" @@* \"}\" )*"`
}

func (b PostEx) String() string {
	return printUnnamed(0, "post-ex", b)
}

type HTTPHostProfiles struct {
	Profiles []HostProfile `parser:"( @@ )*"`
}

func (b HTTPHostProfiles) String() string {
	return printUnnamed(0, "http-host-profiles", b)
}

type HostProfile struct {
	Name string `parser:"@Ident? \"{\""`

	HostName string              `parser:"( \"set\" \"host-name\" @String \";\""`
	HTTPGet  HostProfileHTTPGet  `parser:"| \"http-get\" \"{\" @@ \"}\""`
	HTTPPost HostProfileHTTPPost `parser:"| \"http-post\" \"{\" @@ \"}\" )* \"}\""`
}

func (b HostProfile) String() string {
	return printUnnamed(1, "profile", b)
}

type HostProfileHTTPGet struct {
	URI        string      `parser:"( \"set\" \"uri\" @String \";\""`
	Headers    []Header    `parser:"| \"header\" @@ \";\""`
	Parameters []Parameter `parser:"| \"parameter\" @@ \";\" )*"`
}

func (b HostProfileHTTPGet) String() string {
	return printUnnamed(2, "http-get", b) //nolint: mnd // not needed by design
}

type HostProfileHTTPPost struct {
	URI        string      `parser:"( \"set\" \"uri\" @String \";\""`
	Headers    []Header    `parser:"| \"header\" @@ \";\""`
	Parameters []Parameter `parser:"| \"parameter\" @@ \";\" )*"`
}

func (b HostProfileHTTPPost) String() string {
	return printUnnamed(2, "http-post", b) //nolint: mnd // not needed by design
}
