package merkledag

import "hash"

type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

func Add(store KVStore, node Node, h hash.Hash) []byte {
	switch node.Type() {
	case FILE:
		file := node.(File)
		store.Save(file.Name(), file.Bytes())
		return calculateHash(file.Bytes(), h)
	case DIR:
		dir := node.(Dir)
		dirIterator := dir.It()

		var concatenatedHashes string
		for dirIterator.Next() {
			childNode := dirIterator.Node()
			childHash := Add(store, childNode, h)
			concatenatedHashes += hex.EncodeToString(childHash)
		}
		return calculateHash([]byte(concatenatedHashes), h)
	default:
		return nil
	}
}

func calculateHash(data []byte, h hash.Hash) []byte {
	h.Write(data)
	hashValue := h.Sum(nil)
	h.Reset()
	return hashValue
}
