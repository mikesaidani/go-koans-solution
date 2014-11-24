package go_koans

import "bytes"

func aboutCommonInterfaces() {
  {
    in := new(bytes.Buffer)
    in.WriteString("hello world")

    out := new(bytes.Buffer)

    out.ReadFrom(in)

    assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
  }

  {
    in := new(bytes.Buffer)
    in.WriteString("hello world")

    out := new(bytes.Buffer)

    out.Write(in.Bytes()[:5])

    assert(out.String() == "hello") // duplicate only a portion of the io.Reader
  }
}
