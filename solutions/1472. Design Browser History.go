type BrowserHistory struct {
	history  []string
	position int
}

func Constructor(homepage string) BrowserHistory {
	b := BrowserHistory{}
	b.history, b.position = append(b.history, homepage), 1
	return b
}

func (this *BrowserHistory) Visit(url string) {
	if len(this.history) > this.position {
		this.history = this.history[:this.position+1]
	}
	this.history = append(this.history, url)
	this.position = len(this.history) - 1
}

func (this *BrowserHistory) Back(steps int) string {
	this.position -= min(steps, this.position)
	return this.history[this.position]
}

func (this *BrowserHistory) Forward(steps int) string {
	maxForward := len(this.history) - this.position - 1

	this.position += min(steps, maxForward)

	return this.history[this.position]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}