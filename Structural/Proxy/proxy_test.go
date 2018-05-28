package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	var  searcher Searcher
	searcher = &ProxySearcher{&RealSearch{},&AccessValidator{},&Logger{}}
	searcher.DoSearch("杨过", "玉女心经")

}