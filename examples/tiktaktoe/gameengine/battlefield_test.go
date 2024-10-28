/*
 *
 *
 * MIT NON-AI License
 *
 * Copyright (c) 2024-2024 Aleksei Kotelnikov(gudron2s@gmail.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of the software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions.
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 *
 * In addition, the following restrictions apply:
 *
 * 1. The Software and any modifications made to it may not be used for the purpose of training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining. This condition applies to any derivatives,
 * modifications, or updates based on the Software code. Any usage of the Software in an AI-training dataset is considered a breach of this License.
 *
 * 2. The Software may not be included in any dataset used for training or improving machine learning algorithms,
 * including but not limited to artificial intelligence, natural language processing, or data mining.
 *
 * 3. Any person or organization found to be in violation of these restrictions will be subject to legal action and may be held liable
 * for any damages resulting from such use.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 */

package gameengine

import (
	"fmt"
	"testing"
)

func TestSLC1(t *testing.T) {
	originSlice := make([]int, 0, 8)
	originSlice = append(originSlice, []int{1, 2, 3, 4, 5}...)

	clonedSlice := append(originSlice[1:3], []int{255, 256}...)
	clonedSlice2 := append(originSlice[:3], []int{355, 356, 357, 358, 359, 360}...)
	clonedSlice2[0] = 450
	clonedSlice2[1] = 452

	fmt.Printf("newSlice: %v, len: %d, cap: %d\n", originSlice, len(originSlice), cap(originSlice))
	fmt.Printf("clonedSlice: %v, len: %d, cap: %d\n", clonedSlice, len(clonedSlice), cap(clonedSlice))
	fmt.Printf("clonedSlice2: %v, len: %d, cap: %d\n", clonedSlice2, len(clonedSlice2), cap(clonedSlice2))
}

func TestSLC3(t *testing.T) {
	originSlice := []int{1, 2, 3, 4, 5}

	// Новый срез с len = 0 и cap = 4
	newSlice := originSlice[0:0:4]

	fmt.Printf("newSlice: %v, len: %d, cap: %d\n", newSlice, len(newSlice), cap(newSlice))

	// Добавляем элементы в newSlice
	newSlice = append(newSlice, 10, 20, 30)
	fmt.Printf("newSlice после append: %v, len: %d, cap: %d\n", newSlice, len(newSlice), cap(newSlice))

	// Проверяем исходный срез
	fmt.Printf("originSlice: %v\n", originSlice)
}

func TestSLC2(t *testing.T) {
	var test = [6]string{"adc", "def", "ghi", "ooz", "voooz", "booz"}

	originSlice := test[:]
	secondSlice := test[1:3]

	fmt.Printf("originSlice: %v, len: %d, cap: %d\n", originSlice, len(originSlice), cap(originSlice))

	fmt.Printf("secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))

	secondSlice = append(secondSlice, "ZZZ", "YYY")

	fmt.Printf("new secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("originSlice: %v, len: %d, cap: %d\n", originSlice, len(originSlice), cap(originSlice))

	secondSlice = append(secondSlice, "HAHA", "HAHA2", "HAHA3")

	fmt.Printf("new new secondSlice: %v, len: %d, cap: %d\n", secondSlice, len(secondSlice), cap(secondSlice))
	fmt.Printf("originSlice: %v, len: %d, cap: %d\n", originSlice, len(originSlice), cap(originSlice))
}

func TestArray(t *testing.T) {
	var test = [3]string{"adc", "def", "ghi"}

	clone := test

	t.Logf("test: %p", &test)
	t.Logf("clone: %p", &clone)

	clone[0] = "xyz"

	t.Logf("test: %v", test)
	t.Logf("clone: %v", clone)

	test[2] = "XXX"

	t.Logf("test: %v", test)
	t.Logf("clone: %v", clone)
}

func TestSlice(t *testing.T) {
	origin := make([]uint8, 5, 15)

	origin[0] = 1
	origin[1] = 2
	origin[2] = 3
	origin[3] = 4
	origin[4] = 5

	t.Logf("test fields: %p, len: %d, cap: %d", origin, len(origin), cap(origin))
	clone := append(origin[:0:0], origin[:]...)
	testPassSlice(origin)
	t.Log(origin[4])
	t.Logf("clone fileds: %p, len: %d, cap: %d", clone, len(clone), cap(clone))

	clone[0] = 255
	//clone[6] = 255

	t.Logf("test fields values: %v, len: %d, cap: %d", origin, len(origin), cap(origin))
	t.Logf("clone fields values: %v, len: %d, cap: %d", clone, len(clone), cap(clone))
}

func TestSlice2(t *testing.T) {
	origin := make([]uint8, 5)

	origin[0] = 1
	origin[1] = 2
	origin[2] = 3
	origin[3] = 4
	origin[4] = 5

	t.Logf("test fields: %p, len: %d, cap: %d", origin, len(origin), cap(origin))
	clone := origin[:]
	t.Logf("clone fileds: %p, len: %d, cap: %d", clone, len(clone), cap(clone))

	clone[0] = 255

	t.Logf("test fields values: %v, len: %d, cap: %d", origin, len(origin), cap(origin))
	t.Logf("clone fields values: %v, len: %d, cap: %d", clone, len(clone), cap(clone))
}

func testPassSlice(slc []uint8) {
	slc = append(slc, []uint8{222, 223, 224, 225, 226, 227}...)

	slc[0] = 254
	slc[4] = 253

	fmt.Println(slc)
}

type BattleField struct {
	Players [2]string
	Fields  []uint8
	Size    uint8
}

func (bf *BattleField) Clone() *BattleField {
	data := *bf

	//data.Fields = append([]uint8(nil), bf.Fields...)
	//data.Fields = make([]uint8, len(bf.Fields))
	//copy(data.Fields, bf.Fields)
	data.Fields = append(bf.Fields[:0:0], bf.Fields...)

	return &data
}

func TestStruct(t *testing.T) {
	test := &BattleField{
		Players: [2]string{"abc", "def"},
		Fields:  []uint8{10, 15, 29},
		Size:    3,
	}

	clone := test.Clone()

	t.Logf("test struct: %p", &test)
	t.Logf("clone struct: %p", &clone)

	t.Logf("test players: %p", &test.Players)
	t.Logf("clone players: %p", &clone.Players)

	t.Logf("test fields: %p", test.Fields)
	t.Logf("clone fileds: %p", clone.Fields)

	test.Fields[0] = 255

	t.Logf("test fields values: %v", test.Fields)
	t.Logf("clone fields values: %v", clone.Fields)

	test.Players[0] = "XXX"

	t.Logf("test players values: %v", test.Players)
	t.Logf("clone players values: %v", clone.Players)
}
