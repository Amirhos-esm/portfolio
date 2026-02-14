package models

func (d Experience) GetStartDate() string {
	return d.StartDate.Format("Jan 2006")
}
func (d Experience) GetEndDate() string {
	return d.EndDate.Format("Jan 2006")
}

func (d Education) GetStartDate() string {
	return d.StartDate.Format("Jan 2006")
}
func (d Education) GetEndDate() string {
	return d.EndDate.Format("Jan 2006")
}

func (p *Project) AddTag(tag string) {
	if tag == "" {
		return
	}
	// if p.Tags == nil {
	// 	p.Tags = make([]string, 0)
	// }

	for _, t := range p.Tags {
		if t == tag {
			return // already exists
		}
	}

	p.Tags = append(p.Tags, tag)
}

func (p *Project) RemoveTag(tag string) {
	for i, t := range p.Tags {
		if t == tag {
			p.Tags = append(p.Tags[:i], p.Tags[i+1:]...)
			return
		}
	}
}

func (p *Project) AddGalleryImage(id string) {

	for _, img := range p.GalleryImages {
		if img == id {
			return // already exists
		}
	}
	p.GalleryImages = append(p.GalleryImages, id)
}

func (p *Project) RemoveGalleryImage(id string) bool {
	for i, img := range p.GalleryImages {
		if img == id {
			p.GalleryImages = append(p.GalleryImages[:i], p.GalleryImages[i+1:]...)
			return true
		}
	}
	return false
}

func (p *Project) AddFeature(feature *ProjectFeature) {
	if feature == nil {
		return
	}
	var max  uint = 0
	for _, f := range p.Features {
		if f.ID > uint(max) {
			max = f.ID
		}
	}

	feature.ID = max + 1
	p.Features = append(p.Features, feature)
}

func (p *Project) RemoveFeature(featureID uint) bool {
	for i, f := range p.Features {
		if f.ID == featureID {
			p.Features = append(p.Features[:i], p.Features[i+1:]...)
			return true
		}
	}
	return false
}

func (p *Project) AddTech(category, tech string) bool {
	if category == "" || tech == "" {
		return false
	}
	var MyMap []string
	value := p.TechStack[category]
	if value != nil {
		// M, IsOk := value.([]string)
		// if IsOk == false {
		// 	panic(fmt.Errorf(" value.([]string) "))
		// }
		// MyMap = M
		MyMap = value
	} else {
		MyMap = make([]string, 0)
	}

	for _, t := range MyMap {
		if t == tech {
			return false
		}
	}
	MyMap = append(MyMap, tech)
	p.TechStack[category] = MyMap

	return true
}

func deleteByValue(slice []string, value string) []string {
	for i, v := range slice {
		if v == value {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (p *Project) RemoveTech(category, tech string) bool {
	if category == "" || tech == "" {
		return false
	}

	value := p.TechStack[category]

	value = deleteByValue(value, tech)
	if len(value) != 0 {
		p.TechStack[category] = value
	} else {
		delete(p.TechStack, category)
	}

	return true
}

func (s *Skills) AddSoftSkill(key string) bool {
	if key == "" {
		return false
	}
	for _, t := range s.Soft {
		if t == key {
			return false
		}
	}
	s.Soft = append(s.Soft, key)
	return true

}
func (s *Skills) AddTecknicalSkill(key string) bool {
	if key == "" {
		return false
	}
	for _, t := range s.Soft {
		if t == key {
			return false
		}
	}
	s.Soft = append(s.Soft, key)
	return true
}

func (s *Skills) DeleteSoftSkill(key string) bool {
	if key == "" {
		return false
	}
	for _, t := range s.Soft {
		if t == key {
			return false
		}
	}
	s.Soft = deleteByValue(s.Soft, key)
	return true
}
func (s *Skills) DeleteTecknicalSkill(key string) bool {

	if key == "" {
		return false
	}
	for _, t := range s.Technical {
		if t == key {
			return false
		}
	}
	s.Technical = deleteByValue(s.Technical, key)
	return true
}
