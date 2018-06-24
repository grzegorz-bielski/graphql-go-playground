package schema

//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
import "bytes"

// String parses schema from generated bindata.go (AssetNames(), MustAsset())
func String() string {
	buffer := bytes.Buffer{}
	for _, name := range AssetNames() {
		bytes := MustAsset(name)
		buffer.Write(bytes)

		if len(bytes) > 0 && bytes[len(bytes)-1] != '\n' {
			buffer.WriteByte('\n')
		}
	}
	return buffer.String()
}
