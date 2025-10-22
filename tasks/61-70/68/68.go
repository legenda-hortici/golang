package main

type Data struct {
	ID      int
	Payload map[string]interface{}
}

type Reader interface {
	Read() []*Data
}

type Processor interface {
	Process(Data) ([]*Data, error)
}

type Writer interface {
	Write([]*Data)
}

type Manager interface {
	Manage() // blocking
}

// Необходимо имплементировать интерфейс Manager так, чтобы Manager мог принимать данные из одного Reader
// обрабатывать полученные данные на каждом из списка Processor и результирующие данные передавать в Writer.
// При возникновении ошибки при обработке, прочитанные из Reader данные необходимо пропустить.

type manager struct {
	reader    Reader
	writer    Writer
	processor []Processor
}

func (m *manager) Manage() {
	for {
		data := m.reader.Read()
		if len(data) == 0 {
			continue
		}

		for _, d := range data {
			cur := []*Data{d}

			var err error
			for _, p := range m.processor {
				var out []*Data
				out, err = p.Process(*d)
				if err != nil {
					cur = nil
					break
				}
				cur = out
			}

			if cur != nil {
				m.writer.Write(cur)
			}
		}
	}
}
