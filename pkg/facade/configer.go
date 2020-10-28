package facade

type Configer interface {
	Env(envName string, defaultValue ...interface{}) interface{}
	Add(name string, configuration map[string]interface{})
	Get(path string, defaultValue ...interface{}) interface{}
	GetInterface(path string, defaultValue ...interface{}) (value interface{})
	GetString(path string, defaultValue ...interface{}) string
	GetInt(path string, defaultValue ...interface{}) int
	GetInt64(path string, defaultValue ...interface{}) int64
	GetUint(path string, defaultValue ...interface{}) uint
	GetBool(path string, defaultValue ...interface{}) bool
}
