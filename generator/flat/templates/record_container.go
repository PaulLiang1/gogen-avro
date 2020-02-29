package templates

const RecordContainerTemplate = `
import (
	"io"

	"github.com/actgardner/gogen-avro/container"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

func {{ .NewWriterMethod }}(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	t := {{ .ConstructorMethod }}
	return container.NewWriter(writer, codec, recordsPerBlock, t.Schema())
}

// container reader
type {{ .RecordReaderTypeName }} struct {
	r io.Reader
	p *vm.Program
}

func New{{ .RecordReaderTypeName }}(r io.Reader) (*{{ .RecordReaderTypeName }}, error){
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := {{ .ConstructorMethod }}
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &{{ .RecordReaderTypeName }} {
		r: containerReader,
		p: deser,
	}, nil
}

func (r {{ .RecordReaderTypeName }}) Read() (t {{ .GoType }}, err error) {
	err = vm.Eval(r.r, r.p, &t)
	return
}
`
