package olimpiad

var counter = -1

type Olimp struct {
	name string
	subject string
	level int
	olimpNumber int
}

func NewOlimpiad(name string, subject string, level int) Olimp {
	counter++
	return Olimp{name: name, subject: subject, level: level, olimpNumber: counter}
}

func (olimp Olimp) GetName() string {
	return olimp.name
}

func (olimp Olimp) GetSubject() string {
	return olimp.subject
}

func (olimp Olimp) GetLevel() int {
	return olimp.level
}

func (olimp Olimp) GetOlimpNumber() int {
	return olimp.olimpNumber
}