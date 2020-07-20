package saml

//Config mapping to the config file for SAML
type Config struct {
	MetadataFile        string `yaml:"metadatafile"`
	JksFile     		string `yaml:"jksfile"`
	JksPassword         string `yaml:"jkspassword"`
	KeystoreAlias      	string `yaml:"keystorealias"`
	AudienceUri        	string `yaml:"audienceuri"`
}
