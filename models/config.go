package models

//ConfigServer es la configuración de conexión a la base de datos
type ConfigServer struct {
	Protocol string `json:"protocol"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

//ConfigDB es la configuración de conexión a la base de datos
type ConfigDB struct {
	Engine   string `json:"engine"`
	Server   string `json:"server"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
