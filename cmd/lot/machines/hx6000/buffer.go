package hx6000

type HxBuffer struct {
	Buffer      []byte
	ValidLength int
}

func NewHxBuffer() *HxBuffer {
	return &HxBuffer{
		Buffer:      make([]byte, 1024),
		ValidLength: 0,
	}
}

func (b *HxBuffer) Reset() {
	tmp := make([]byte, 1024)
	b.Buffer = tmp
	b.ValidLength = 0
}

func (b *HxBuffer) Clear(deadLength int) {
	// 	if b.ValidLength == deadLength {
	// 		b.Reset()
	// 		return
	// 	}
	// 	// deadLength长度用完抛弃
	// 	b.Buffer = append(b.Buffer[deadLength:], make([]byte, deadLength)...)
	// 	b.ValidLength -= deadLength
	// 	fmt.Println("Clear后长度是", len(b.Buffer))
	tmp := make([]byte, 1024)
	copy(tmp, b.Buffer[deadLength:b.ValidLength])
	b.Buffer = tmp
	b.ValidLength -= deadLength
}
