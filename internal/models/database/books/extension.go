package books

// you can extend extension field to enrich model
type extension struct {
	LastUpdated int64
}

func (model *Books) GetLastUpdated() int64 {
	if model.ext == nil {
		model.ext = new(extension)
		model.ext.LastUpdated = model.UpdateAt
	}
	return model.ext.LastUpdated
}

func (model *Books) AsBriefInfo() map[string]interface{} {
	return map[string]interface{}{
		"id":   model.Id,
		"name": model.Name,
	}
}
