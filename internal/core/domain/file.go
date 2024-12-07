package domain

type File struct {
	Path string
	CID  string
}

func NewFile(path, cid string) (*File, error) {
	if path == "" {
		return nil, ErrInvalidPath
	}
	if cid == "" {
		return nil, ErrInvalidCID
	}
	return &File{
		Path: path,
		CID:  cid,
	}, nil
}
