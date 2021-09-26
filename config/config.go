package config

import "sync"

var PortList = map[string]int{
	"ftp":           21,
	"ssh":           22,
	"smb":           445,
	"mongodb":       27017,
	"mysql":         3306,
	"mssql":         1433,
	"redis":         6379,
	"postgres":      5432,
	"couchdb":       5984,
	"elasticsearch": 9200,
	"zookeeper":     2181,
	"memcached":     11211,
}

var UserList = map[string][]string{
	"ftp":      {"ftp", "admin", "www", "web", "root", "db", "wwwroot", "data", "anymous"},
	"mysql":    {"root", "mysql"},
	"mssql":    {"sa", "sql"},
	"smb":      {"administrator", "admin", "guest"},
	"postgres": {"postgres", "admin"},
	"ssh":      {"root", "admin"},
	"mongodb":  {"root", "admin"},
}

var Passwords = []string{"", "{user}1234", "{user}123456", "{user}12345", "{user}@123456", "{user}@12345", "{user}#123456", "{user}#12345", "{user}_123456", "{user}_12345", "{user}123!@#", "{user}!@#$", "{user}!@#", "{user}~!@", "{user}!@#123", "qweasdzxc", "{user}2017", "{user}2016", "{user}2015", "{user}@2017", "{user}@2016", "{user}@2015", "admin888", "administrator", "administrator123", "ftp", "ftppass", "12345", "1234", "qwerty", "1q2w3e4r", "qazwsx", "123qwe", "123qaz", "oracle", "1234567", "123456qwerty", "password123", "1q2w3e", "q1w2e3r4", "user", "mysql", "web{user}", "88888888", "q1w2e3r4{user}", "root123", "web", "mongodb", "mongodb123", "mongodbpass", "test12345", "test123456", "pass", "7654321", "888888", "987654321", "147258369", "123asd", "qwer123", "root3306", "Q1W2E3b3{user}", "sys", "dbadmin", "oracle{user}", "q1w2e3r4admin", "postgres", "admin#123", "redhat", "apache", "sa123", "sa", "sqlpass", "sql123", "sqlserver", "qwer1234admin", "wwwroot", "data", "testadmin", "webadmin", "adminsys", "orcl", "testpostgres", "webadministrator", "manager", "guest", "db2admin", "testsaroot", "backup", "upload", "linux", "temp", "nagios", "user1", "www", "test1", "portaladmin", "guest", "anymous", "123456", "admin", "admin123", "root", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "{user}@123#4", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "0", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!"}
var DefaultPorts = "21,22,445,1433,2181,3306,5432,5984,6379,9200,27017"

type Info struct {
	Host      string
	Port      string
	Service   string
	Username  string
	Password  string
	Usernames []string
	Passwords []string
	Timeout   int64
	Scantype  string
	Domain    string
	Scantypes []string
}

var Num int64
var End int64
var SuccessHash sync.Map
var NoOutput bool
var NoScan bool

var (
	Hostfile  string
	Userfile  string
	Passfile  string
	AliveAddr []Info
	Thread    int
)
