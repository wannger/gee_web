package gee

import (
	"gee_web/common"
	"log"
	"net/http"
)

// 路由内部逻辑操作

// 字典树的节点信息
type node struct {
	part     string  // 当前节点的路由部分，包含":","*"
	children []*node // 儿子节点
}

func (n *node) insert(parts []string, pos int) {
	if n.part == "*" || len(parts) == pos {
		return
	}
	isHas := false
	for i := 0; i < len(n.children); i++ {
		if n.children[i].part[0] == '*' || n.children[i].part[0] == ':' || n.children[i].part == parts[pos] {
			isHas = true
			n.children[i].insert(parts, pos+1)
			break
		}
	}
	if !isHas {
		n.children = append(n.children, &node{
			part:     parts[pos],
			children: make([]*node, 0),
		})
		n.children[len(n.children)-1].insert(parts, pos+1)
	}
}

func (n *node) search(parts []string, pos int) (bool, string) {
	if len(parts) == pos {
		return true, ""
	}
	for i := 0; i < len(n.children); i++ {
		if n.children[i].part[0] == ':' || n.children[i].part == parts[pos] {
			isMapping, treePath := n.children[i].search(parts, pos+1)
			if !isMapping {
				return false, ""
			} else {
				return true, "/" + n.children[i].part + treePath
			}
		} else if n.children[i].part[0] == '*' {
			return true, "/" + n.children[i].part
		}
	}
	// 匹配不成功
	return false, ""
}

type router struct {
	roots    map[string]*node // 不同类型方法的字典树根结点
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	// pattern 分解成数组
	parts := common.ParsePattern(pattern)
	// 插入路由字典树中
	key := method + "-" + pattern
	// 当前方法类型没有，新建一个空结点
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{
			part:     "",
			children: make([]*node, 0),
		}
	}
	r.roots[method].insert(parts, 0)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	// c.Path 分解成数组
	paths := common.ParsePattern(c.Path)
	// 匹配路由字典树，调用业务逻辑方法
	isMapping, treePath := r.roots[c.Method].search(paths, 0)
	log.Printf("treePath: %+v\n", treePath)
	if isMapping && treePath != "" {
		key := c.Method + "-" + treePath
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404: %s", c.Path)
	}
}
