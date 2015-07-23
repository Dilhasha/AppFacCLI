package command


type Command interface {
	Metadata() CommandMetadata
	Configs(reqs CommandRequirements) CommandConfigs
	Requirements(args []string) CommandRequirements
	Run(c CommandConfigs)
}




