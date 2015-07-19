package command


type Command interface {
	Metadata() CommandMetadata
	//Configs(reqs CommandRequirements) CommandConfigs
	Configs() CommandConfigs
	Requirements() CommandRequirements
	Run(c CommandConfigs)
}




