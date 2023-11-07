package mock_dao

type MockWordDao struct {
}

func (m *MockWordDao) Start(length int, mode string, varargs ...interface{}) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockWordDao) Exists(word string) bool {
	return true
}
