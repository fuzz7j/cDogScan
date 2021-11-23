package plugins

import "cDogScan/config"

type Choose func(Info *config.Info) (result bool, err error)

var (
	ScanFuncMap map[string]Choose
)

func init()  {
	ScanFuncMap = make(map[string]Choose)
	ScanFuncMap["21"] = FtpScan
	ScanFuncMap["22"] = SshScan
	ScanFuncMap["161"] = SnmpScan
	ScanFuncMap["445"] = SmbScan
	ScanFuncMap["1433"] = MssqlScan
	ScanFuncMap["3306"] = MysqlScan
	ScanFuncMap["5432"] = PostgresScan
	ScanFuncMap["6379"] = RedisScan
	ScanFuncMap["27017"] = MongodbScan
	ScanFuncMap["9200"] = ElasticScan
	ScanFuncMap["2181"] = ZookeeperScan
	ScanFuncMap["11211"] = MemcachedScan
	ScanFuncMap["5984"] = CouchDBScan
}