package formats

type VersionFormat struct {
	Version string
	AutoDeployment string
	Stage string
	IsAutoBuild string
	IsAutoDeploy string
	RepoURL string
	lastBuildResult string
	DisplayRoles []string
	CurrentBuildStatus string
	PromoteStatus string
	DeployedBuid string
	IsInitialStage bool
}
