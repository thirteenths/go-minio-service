package models

type ImageUnit struct {
	BucketName string
	ObjectName string
	ObjectPath string
	ObjectSize int64
}

type ObjectUnit struct {
	BucketName  string
	ObjectName  string
	ObjectBytes []byte
	ObjectSize  int64
	ContentType string
}

type Object struct {
}
