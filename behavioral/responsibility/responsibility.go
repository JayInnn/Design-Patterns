package responsibility

type Handler interface {
	Process(content string)
	GetNextHandler() Handler
}

// the responsibility for Button handler
type ButtonHandler struct {

}

func (b *ButtonHandler) bNeedProcess(content string) bool {
	return false
}

func (b *ButtonHandler) Process(content string) {
	if b.bNeedProcess(content) {
		// TODO:该handler处理内容
	}
	if b.GetNextHandler() != nil {
		b.GetNextHandler().Process(content)
	}
}

func (b *ButtonHandler) GetNextHandler() Handler {
	return &SubBoardHandler{}
}


// the responsibility for SubBoard handler
type SubBoardHandler struct {

}

func (s *SubBoardHandler) bNeedProcess(content string) bool {
	return true
}

func (s *SubBoardHandler) Process(content string) {
	if s.bNeedProcess(content) {
		// TODO:该handler处理内容
	}
	if s.GetNextHandler() != nil {
		s.GetNextHandler().Process(content)
	}
}

func (s *SubBoardHandler) GetNextHandler() Handler {
	return nil
}


