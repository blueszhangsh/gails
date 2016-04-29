package gails

//Migration migration
type Migration interface {
	Migrate()
	Seed()
}
