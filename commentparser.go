package pbmeta

import (
	"github.com/davyxu/golexer"
)

// 自定义的token id
const (
	Token_EOF = iota
	Token_LeftBrace
	Token_RightBrace
	Token_WhiteSpace
	Token_LineEnd
	Token_UnixStyleComment
	Token_Identifier
	Token_Unknown
)

type CommentParser struct {
	*golexer.Parser
}

func parseTaggedComment(src string, commentArray []*TaggedComment) []*TaggedComment {

	return NewCommentParser(src).Run(src, commentArray)
}

func (self *CommentParser) Run(src string, commentArray []*TaggedComment) (ret []*TaggedComment) {

	defer golexer.ErrorCatcher(func(err error) {

		log.Errorln(err, self.TokenID(), self.TokenValue())

		ret = commentArray

	})

	self.Lexer().Start(src)

	for {

		self.NextToken()

		//log.Debugln("#", self.TokenID(), self.TokenValue())

		if self.TokenID() == Token_WhiteSpace {
			continue
		}

		// 读取标头
		if self.TokenID() == Token_LeftBrace {

			self.NextToken()

			var tc TaggedComment

			tagNameToken := self.Expect(Token_Identifier)

			self.Expect(Token_RightBrace)

			tc.Name = tagNameToken.Value()

			for {

				//log.Debugln("#", self.TokenID(), self.TokenValue())

				tc.Data += self.TokenValue()

				self.NextToken()

				if self.TokenID() == Token_LineEnd || self.TokenID() == Token_EOF {
					break
				}
			}
			commentArray = append(commentArray, &tc)

		}

		if self.TokenID() == Token_EOF {
			break
		}

	}

	return commentArray
}

func NewCommentParser(srcName string) *CommentParser {

	l := golexer.NewLexer()

	// 匹配顺序从高到低

	l.AddMatcher(golexer.NewSignMatcher(Token_LeftBrace, "["))
	l.AddMatcher(golexer.NewSignMatcher(Token_RightBrace, "]"))

	l.AddMatcher(golexer.NewWhiteSpaceMatcher(Token_WhiteSpace))
	l.AddMatcher(golexer.NewLineEndMatcher(Token_LineEnd))
	l.AddIgnoreMatcher(golexer.NewUnixStyleCommentMatcher(Token_UnixStyleComment))

	l.AddMatcher(golexer.NewIdentifierMatcher(Token_Identifier))

	l.AddMatcher(golexer.NewUnknownMatcher(Token_Unknown))

	return &CommentParser{
		Parser: golexer.NewParser(l, srcName),
	}
}
