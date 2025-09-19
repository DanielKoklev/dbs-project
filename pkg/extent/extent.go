package extent

const DefaultExtentSize = 4 << 20 // 4 MiB

type Extent struct {
	ID string
	Data []byte
}

