package urlmanager

import (
	"regexp"
)

type IUrlManager interface {
	ParseRequest(path string) (moduleId string, controllerId string, actionId string)
}

type _DefaultUrlManager struct {
	rules *regexp.Regexp
}

func NewDefault() IUrlManager {
	exp, err := regexp.Compile(`/(\w+)/(\w+)/(\w+)$`)
	if err != nil {
		return nil
	}
	return &_DefaultUrlManager{
		exp,
	}
}

func (urlmanager *_DefaultUrlManager) ParseRequest(path string) (moduleId string, controllerId string, actionId string) {
	match := urlmanager.rules.FindStringSubmatch(path)
	if len(match) >= 3 {
		return match[0], match[1], match[2]
	}
	return "", "", ""
}
