package framework

// IGroup 代表前綴分組
type IGroup interface {
	// 实现HttpMethod方法
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	// 实现嵌套group
	Group(string) IGroup
}

// Group struct 实现了IGroup
type Group struct {
	core   *Core  // 指向core结构
	parent *Group //指向上一个Group，如果有的话
	prefix string // 这个group的通用前缀
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}
func (g *Group) Get(uri string, handler ControllerHandler) {
	g.core.Get(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Post(uri string, handler ControllerHandler) {
	g.core.Post(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Put(uri string, handler ControllerHandler) {
	g.core.Put(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Delete(uri string, handler ControllerHandler) {
	g.core.Delete(g.getAbsolutePrefix()+uri, handler)
}

func (g *Group) Group(uri string) IGroup {
	newGroup := NewGroup(g.core, uri)
	newGroup.parent = g
	return newGroup
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}
