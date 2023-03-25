package message

type Message struct {
	Code    int
	Message string
}

type Messages []Message

type Plugin struct {
	File     string
	Messages []Message
}

type Plugins []Plugin

func (p *Plugin) Iterator(doIterator func(Message)) {
	for _, msg := range p.Messages {
		doIterator(msg)
	}
}

func (p *Plugin) Add(m Message) {
	p.Messages = append(p.Messages, m)
}

func (p *Plugins) Add(plugin Plugin) {
	(*p) = append(*p, plugin)
}

func (p *Plugins) IsSet() bool {
	var c int

	for _, plugin := range *p {
		c = c + len(plugin.Messages)
	}

	return c > 0
}
