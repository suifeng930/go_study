package model

// new struct
type student struct {
	Name  string
	score float64
}

// 使用工厂模式 因为student结构体首字母小写，因此只能在model包中使用
// 采用工厂模式解决包的引入问题

func NewStudent(name string, score float64) *student {

	return &student{
		Name:  name,
		score: score,
	}
}

// 如果score 字段首字符小写，则，该字段在其它包中将不可以直接使用，因此提供一个方法可解决这个问题
func (s *student) GetScore() float64 {
	return s.score
}
