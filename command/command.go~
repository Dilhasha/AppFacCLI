package command

type Command interface {
	Metadata() CommandMetadata
	Configs(reqs CommandRequirements) CommandConfigs
	Requirements() CommandRequirements
	Run(c CommandConfigs)
}



