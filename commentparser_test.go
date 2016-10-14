package pbmeta

import (
	"testing"
)

func TestCommentParser(t *testing.T) {

	var tag []*TaggedComment

	tag = parseTaggedComment("[agent] client -> battle", tag)
	tag = parseTaggedComment("[ab] awef\n\n[cd] 23\n", tag)
	for _, r := range tag {
		t.Log(r)
	}
}
