package models

type ServiceConfig struct {
	// node app
	WorkingDirectory string `json:"working_directory"`
	EntryFile string `json:"entry_file"`
	LogDirectory string `json:"log_directory"`

	// windows service
	Name string `json:"name"`
	ServiceType uint32 `json:"service_type"`
	StartType uint32 `json:"start_type"`
	ErrorControl uint32 `json:"error_control"`
	LoadOrderGroup string `json:"load_order_group"`
	TagId uint32 `json:"tag_id"`
	Dependencies []string `json:"dependencies"`
	ServiceStartName string `json:"service_start_name"`
	DisplayName string `json:"display_name"`
	Password string `json:"password"`
	Description string `json:"description"`
	SidType uint32 `json:"sid_type"`
	DelayedAutoStart bool `json:"delayed_auto_start"`
}