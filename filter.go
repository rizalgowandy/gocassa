package cmagic

type filter struct {
	t table
	rs []Relation
}

func (f filter) Replace(i interface{}) error {
	return nil
}

func (f filter) Update(m map[string]interface{}) error {
	//m, ok := toMap(i)
	//if !ok {
	//	return errors.New("Update: value not understood")
	//}
	//id, ok := m[c.info.primaryKey]
	//if !ok {
	//	return errors.New("Update: primary key not found")
	//}
	//fields, values := keyValues(m)
	//for k, v := range m {
	//	if k == c.info.primaryKey {
	//		continue
	//	}
	//	fields = append(fields, k)
	//	values = append(values, v)
	//}
	//stmt := g.UpdateById(c.keySpace.name, c.info.primaryKey, fields)
	//sess := c.keySpace.session
	//return sess.Query(stmt, append(values, id)...).Exec()
	return nil
}

func (f filter) Delete() error {
	// return c.keySpace.session.Query(g.DeleteById(c.keySpace.name, c.info.primaryKey), id).Exec()
	return nil
}

func (f filter) Query() Query {
	return &query{
		f: f,
	}
}