package utils

// 列表转化为树形结构
// 1. 传入的列表必须是有序的，即按照父级id进行排序
// 2. 传入的列表必须是所有的数据，不能是分页的数据
// 3. 传入的列表必须是一个结构体的列表
// 4. 传入的列表必须是一个切片
// 5. 传入的列表必须是一个指针类型
// 6. 传入的列表必须是一个公共的结构体，不能是私有的
type Item interface {
	GetId() int
	GetParentId() int
}

type Tree[T Item] struct {
	Value    T
	Children []T
}

func BuildTree[U Item](items []U) []Tree[U] {
	tree := make([]Tree[U], 0)
	for _, item := range items {
		if item.GetParentId() == 0 {
			tree = append(tree, Tree[U]{Value: item})
		} else {
			for i := range tree {
				if tree[i].Value.GetId() == item.GetParentId() {
					tree[i].Children = append(tree[i].Children, item)
				}
			}
		}
	}
	return tree
}

// func BuildTree[T Item](items []T) []Tree[T] {
// 	tree := make([]Tree[T], 0)
// 	for _, item := range items {
// 		if item.ParentId == 0 {
// 			tree = append(tree, Tree[T]{Value: item})
// 		} else {
// 			for i, _ := range tree {
// 				if tree[i].Value.Id == item.ParentId {
// 					tree[i].Children = append(tree[i].Children, item)
// 				}
// 			}
// 		}
// 	}
// 	return tree
// }
