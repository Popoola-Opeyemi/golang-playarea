package interfaces

import (
	"bytes"
	"fmt"
)

type Writer interface {
	write([]byte) (int, error)
}

type Closer interface {
	close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type bufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *bufferedWriterCloser) write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)

	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 0 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))

		if err != nil {
			return 0, err
		}
	}
	return n, nil

}

func (bwc *bufferedWriterCloser) close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)

		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *bufferedWriterCloser {
	return &bufferedWriterCloser{
		bytes.NewBuffer([]byte{}),
	}
}

func TestInterface() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.write([]byte("The quick fox jumped over the fences with lots of power"))
	wc.close()

	bww := wc.(*bufferedWriterCloser)
	bww.write([]byte("meaning of work"))

	// creating an using normally
	bwc := &bufferedWriterCloser{buffer: bytes.NewBuffer([]byte{})}
	bwc.write([]byte("the meaning of life"))
}
