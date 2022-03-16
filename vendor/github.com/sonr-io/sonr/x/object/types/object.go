package types

func (o *ObjectDoc) Validate(b *ObjectDoc) bool {
	if o.GetLabel() != b.GetLabel() {
		return false
	}

	for k, v := range o.GetFields() {
		if b.GetFields()[k] != v {
			return false
		}
	}
	return true
}
