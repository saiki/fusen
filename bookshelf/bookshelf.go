package bookshelf

type Fusen struct {
	left  int
	top   int
	color string
	body  string
}

func NewFusen(left, top int, color, body string) *Fusen {
	return &Fusen{left: left, top: top, color: color, body:body}
}

type Page struct {
	fusen map[int]*Fusen
}

func (self *Page) Init() *Page {
	self.fusen = make(map[int]*Fusen)
	return self
}

func (self *Page) Add(fusen *Fusen) int {
	next := len(self.fusen) + 1
	self.fusen[next] = fusen
	return next
}

type Note struct {
	pages map[int]*Page
}

func (self *Note) Init() *Note {
	self.pages = make(map[int]*Page)
	return self
}

func (self *Note) AddPage(page *Page) int {
	next := len(self.pages) + 1
	self.pages[next] = page
	return next
}

func (self *Note) AddFusen(fusen *Fusen, page int) int {
	index := self.pages[page].Add(fusen)
	return index
}
