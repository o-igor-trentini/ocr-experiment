package appconst

type Env string

const (
	GIN_HOST Env = "GIN_HOST"
	GIN_PORT Env = "GIN_PORT"

	GIN_MODE   Env = "MODE"
	PG_HOST    Env = "PG_HOST"
	PG_USER    Env = "PG_USER"
	PG_PASS    Env = "PG_PASS"
	PG_NAME    Env = "PG_NAME"
	PG_PORT    Env = "PG_PORT"
	PG_SCH     Env = "PG_SCH"
	PG_SCH_LOG Env = "PG_SCH_LOG"
)
