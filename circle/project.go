package circle

type Project struct {
	Name string
	URL  string
}

type Projects []Project

func (p Projects) Len() int {
	return len(p)
}

func (p Projects) Index(i int) string {
	return p[i].Name
}

func (p *Projects) Remove(i int) {
	*p = append([]Project(*p)[:i], []Project(*p)[i+1:]...)
}

func (p *Projects) Add(project Project) {
	*p = append(*p, project)
}
