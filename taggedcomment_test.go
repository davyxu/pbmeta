package pbmeta

import (
	"testing"
)

func TestComment(t *testing.T) {

	s := "[tabtoy] a:b [table] \" \\\"c:d \\\"[ \" "
	t.Logf("|%s|", s)
	cmtarr := parseTaggedComment(s)

	if len(cmtarr) != 2 {
		t.Fail()
		t.Fatal("num failed", len(cmtarr))
	}

	if cmtarr[0].Name != "tabtoy" {
		t.Fail()
		t.Fatal("1 name failed", cmtarr[0].Name)
	}

	if cmtarr[0].Data != " a:b " {
		t.Fail()
		t.Fatal("1 meta failed", cmtarr[0].Data)
	}

	if cmtarr[1].Name != "table" {
		t.Fail()
		t.Fatal("2 name failed", cmtarr[1].Name)
	}

	if cmtarr[1].Data != " \" \"c:d \"[ \" " {
		t.Fail()
		t.Fatal("2 meta failed", cmtarr[1].Data)
	}

}
