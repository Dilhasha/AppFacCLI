package command


type Command interface {
	Metadata() CommandMetadata
	Run(c CommandConfigs) (bool,string)
}




