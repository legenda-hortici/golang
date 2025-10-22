// Реализация хеш-таблицы

package main

import "fmt"

const ArraySize = 7

// HashTable структура для хеш-таблицы
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket структура, представляющая список для разрешения коллизий
type bucket struct {
	head *bucketNode
}

// bucketNode структура узла в списке
type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

// Insert вставляет элемент в хеш-таблицу
func (h *HashTable) Insert(key string, value int) {
	index := hash(key)
	h.array[index].insert(key, value)
}

// Search возвращает значение по ключу в хеш-таблице
func (h *HashTable) Search(key string) (int, bool) {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete удаляет элемент по ключу из хеш-таблицы
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// hash простая хеш-функция
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// insert добавляет элемент в список
func (b *bucket) insert(key string, value int) {
	if _, found := b.search(key); !found {
		newNode := &bucketNode{key: key, value: value}
		newNode.next = b.head
		b.head = newNode
	} else {
		currNode := b.head
		for currNode != nil {
			if currNode.key == key {
				currNode.value = value
				return
			}
			currNode = currNode.next
		}
	}
}

// search ищет ключ в списке
func (b *bucket) search(key string) (int, bool) {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

// delete удаляет ключ из списка
func (b *bucket) delete(key string) {
	if b.head == nil {
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prevNode := b.head
	currNode := b.head.next

	for currNode != nil {
		if currNode.key == key {
			prevNode.next = currNode.next
			return
		}
		prevNode = currNode
		currNode = currNode.next
	}
}

func main() {
	hashTable := &HashTable{}
	for i := range hashTable.array {
		hashTable.array[i] = &bucket{}
	}

	fmt.Println("Hash Table Created")
	hashTable.Insert("name", 123)
	hashTable.Insert("age", 456)

	name, exists := hashTable.Search("name")
	age, exists := hashTable.Search("age")

	fmt.Println("name:", name, exists)
	fmt.Println("age:", age, exists)

	hashTable.Delete("name")

	name, exists = hashTable.Search("name")

	fmt.Println("name:", name, exists)
}
