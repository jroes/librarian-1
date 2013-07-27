package librarian

type Table struct {
  Name        string         // Name of the table the relation connects to.
  PrimaryKey  string         // Database column that is the models primary key.
  Model       ModelInterface // Model calling Table's New method generates.
  Reflections []ReflectionInterface
}

func (self Table) New() ModelInterface {
  return nil
}

func (self Table) DestroyAll() error {
  return nil
}

func (self Table) Select(columns ...interface{}) *Relation {
  return nil
}

func (self Table) Where(conditions ...interface{}) *Relation {
  return nil

}

func (self Table) Distinct() *Relation {
  return nil
}

func (self Table) Unique() *Relation {
  return nil
}

func (self Table) Order() *Relation {
  return nil
}

func (self Table) Group() *Relation {
  return nil
}

func (self Table) Having() *Relation {
  return nil
}

func (self Table) Limit() *Relation {
  return nil
}

func (self Table) Offset() *Relation {
  return nil
}

func (self Table) Lock() *Relation {
  return nil
}

func (self Table) First() ([]interface{}, error) {
  return nil, nil
}

func (self Table) Last() ([]interface{}, error) {
  return nil, nil
}

func (self Table) All() ([]interface{}, error) {
  return nil, nil
}
