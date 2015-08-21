package main

import (
  "sort"
)

type lessFunc func(p1, p2 *Configuration) bool

type multiSorter struct {
	configurations []Configuration
	less           []lessFunc
}

func (ms *multiSorter) Sort(configurations []Configuration) {
	ms.configurations = configurations
	sort.Sort(ms)
}

func OrderBy(lessFuncs [][]lessFunc) *multiSorter {
	if len(lessFuncs) > 0 {
		for _,lessFunc := range lessFuncs {
			return &multiSorter {
				less: lessFunc,
			}
		}
	} 
	return &multiSorter {}
}

func (ms *multiSorter) Len() int {
	return len(ms.configurations)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.configurations[i], ms.configurations[j] = ms.configurations[j], ms.configurations[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.configurations[i], &ms.configurations[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p,q):
			return true
		case less(q,p):
			return false
		}
	}
	return ms.less[k](p,q)
}

var SortByNameAsc = func(c1, c2 *Configuration) bool {
	return c1.Name < c2.Name
}
var SortByNameDesc = func(c1, c2 *Configuration) bool {
	return c1.Name > c2.Name
}
var SortByHostnameAsc = func(c1, c2 *Configuration) bool {
	return c1.Hostname < c2.Hostname
}
var SortByHostnameDesc = func(c1, c2 *Configuration) bool {
	return c1.Hostname < c2.Hostname
}
var SortByPortAsc = func(c1, c2 *Configuration) bool {
	return c1.Port < c2.Port
}
var SortByPortDesc = func(c1, c2 *Configuration) bool {
	return c1.Port > c2.Port
}
var SortByUsernameAsc = func(c1, c2 *Configuration) bool {
	return c1.Port > c2.Port
}
var SortByUsernameDesc = func(c1, c2 *Configuration) bool {
	return c1.Port > c2.Port
}

var sort_keys = map[string]func(c1, c2 *Configuration) bool {
	"name"          : SortByNameAsc,
	"name:asc"      : SortByNameAsc,
	"name:desc"     : SortByNameDesc,
	"hostname"      : SortByHostnameAsc,
	"hostname:asc"  : SortByHostnameAsc,
	"hostname:desc" : SortByHostnameDesc,
	"port"          : SortByPortAsc,
	"port:asc"      : SortByPortAsc,
	"port:desc"     : SortByPortDesc,
	"username"      : SortByUsernameAsc,
	"username:asc"  : SortByUsernameAsc,
	"username:desc" : SortByUsernameDesc,
}

func SortConfigs(fields []string, c Configurations) {
	fieldLessFuncs := make([][]lessFunc,len(fields))
	for i,el := range fields  {
		fieldLessFuncs[i] = []lessFunc{sort_keys[el]}
	}
	OrderBy(fieldLessFuncs).Sort(c.List)
}


