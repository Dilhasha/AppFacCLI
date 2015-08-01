package formats

type AppVersionFormat struct {
	Key string
	Owner string
	Type string
	Language string
	MappedSubDomain string
	CustomUrl string
	CustomUrlVerificationCode string
	Description string
	RepositoryType string
	RepositoryTypeDisplayName string
	branchCount int64 `json:",string"`
	TypeDisplayName string
	IsUploadable bool
	Name string
	RepoAccessability string
	InProduction bool
	ApplicationCreationStatus string
	Users []UserFormat
	Versions []VersionFormat
}
