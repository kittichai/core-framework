package cassandra

import (
	"github.com/gocql/gocql"
)

type CassandraDB struct {
	*gocql.Session
}

func NewCassandraDB(hosts []string, keyspace string) (*CassandraDB, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &CassandraDB{session}, nil
}

func (db *CassandraDB) Close() {
	db.Session.Close()
}
