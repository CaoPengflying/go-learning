package chain

import "strings"

type SensitiveWordFilter interface {
	DoFilter(content string) bool
}

type SexWordFilter struct {
}

func (s *SexWordFilter) DoFilter(content string) bool {
	if strings.Contains(content, "黄色") {
		return false
	}
	return true
}

type PoliticalWordFilter struct {
}

func (p *PoliticalWordFilter) DoFilter(content string) bool {
	if strings.Contains(content, "反动") {
		return false
	}
	return true
}

type FilterChain struct {
	filters []SensitiveWordFilter
}

func (chain *FilterChain) AddFilter(filter SensitiveWordFilter) {
	chain.filters = append(chain.filters, filter)
}

func (chain *FilterChain) Filter(content string) bool {
	for _, filter := range chain.filters {
		if !filter.DoFilter(content) {
			return false
		}
	}
	return true
}
