package core

type Credentials struct {
	ID        int64 `pgx:"column:id;primary_key"`
	Timestamp int64 `pgx:"column:timestamp"`
	Login     int64 `pgx:"column:login"`
	Password  int64 `pgx:"column:password"`
}

func (Credentials) TableName() string {
	return "credentials"
}
