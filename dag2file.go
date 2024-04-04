package merkledag

import (
	"fmt"
	"io/ioutil"
)

// Hash to file
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) ([]byte, error) {
	// 根据 hash 和 path，返回对应的文件，其中 hash 对应的类型是 tree

	// 检查存储中是否存在给定的哈希值
	exists, err := store.Has(hash)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("Hash not found in store")
	}

	// 从存储中获取与哈希值对应的数据
	data, err := store.Get(hash)
	if err != nil {
		return nil, err
	}

	// 将数据写入文件
	err = ioutil.WriteFile(path, data, 0644)
	if err != nil {
		return nil, err
	}

	// 读取文件内容
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
