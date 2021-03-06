package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// node は、節点を表す。
type node struct {
	key    int
	parent *node // 親
	left   *node // 左の子
	right  *node // 右の子
}

// root は、二分探索木のrootを表す。
var root *node

// insert は、treeにnodeを挿入する。
func insert(z *node) {
	var y *node // xの親
	x := root

	for x != nil { // xがnilじゃない、つまり要素が存在する
		y = x // 親を設定
		if z.key < x.key {
			x = x.left // 左の子に進む
		} else {
			x = x.right // 右の子に進む
		}
	}
	// z(今回の挿入するnode)の親をyにする
	z.parent = y

	if y == nil { // yがnilということは、二分探索木は空だということ
		root = z
	} else if z.key < y.key {
		y.left = z // yの左の子にzを設定する
	} else {
		y.right = z // yの右の子にzを設定する
	}
}

// isExist は、引数で与えられた値が二分探索木の中に存在するかどうかを返す。
func isExist(u *node, key int) bool {
	return find(u, key) != nil
}

// find は、引数で与えられた値が二分探索木の中に存在する場合、そのnodeを返す。
func find(u *node, key int) *node {
	for u != nil && key != u.key { // u.key == key にならない場合は、uを葉まで進める
		if key < u.key {
			u = u.left // 左部分木
		} else {
			u = u.right // 右部分木
		}
	}

	return u
}

// delete は、指定されたkeyを持つnodeを二分探索木から削除する。
func delete(z *node) {
	// y = 削除する対象を接点
	y := &node{}

	if z.left == nil || z.right == nil {
		y = z
	} else { // zが2つの子nodeを持つ場合には
		// zの次のnode
		y = getSuccessor(z)
	}

	// x = yの子node
	x := &node{}
	if y.left != nil {
		// yの左の子nodeが存在する場合は、xはyの左の子nodeとする
		x = y.left
	} else {
		// yの右の子nodeが存在する場合は、xはyの右の子nodeとする
		x = y.right
	}

	if x != nil {
		// xが存在する場合には、xの親nodeは、yの親nodeにする
		x.parent = y.parent
	}

	if y.parent == nil {
		// yがrootの場合、xをrootにする
		root = x
	} else if y == y.parent.left {
		// yがその親nodeの左の子nodeの場合、yの親nodeの左の子nodeをxにする
		y.parent.left = x
	} else {
		// yがその親nodeの右の子nodeの場合、yの親nodeの右の子nodeをxにする
		y.parent.right = x
	}

	if !reflect.DeepEqual(y, z) {
		// zの次のnodeが削除された場合、yのkeyをzのkeyにする
		z.key = y.key
	}
}

// getSuccessor は、引数で与えられたnodeの次nodeを返す。
func getSuccessor(x *node) *node {
	// 右の子nodeが存在する場合は、右nodeの中の最小のnodeが次のnodeとなる
	if x.right != nil {
		return getMinimum(x.right)
	}

	y := x.parent
	// 親のnodeが存在し、親nodeの右の子nodeが対象nodeの場合
	for y != nil && x == y.right {
		// xの親nodeをxにし、yの親nodeをyにする
		x = y
		y = y.parent
	}
	return y
}

// getMinimum は、引数で与えられたnodeをrootとする部分木の中で最小のkeyを持つnodeを返す。
func getMinimum(x *node) *node {
	for x.left != nil {
		x = x.left
	}
	return x
}

var buf = bufio.NewWriter(os.Stdout)

func preOrder(u *node) {
	if u == nil {
		return
	}

	buf.WriteString(fmt.Sprintf(" %d", u.key))
	preOrder(u.left)
	preOrder(u.right)
}
func inOrder(u *node) {
	if u == nil {
		return
	}

	inOrder(u.left)
	buf.WriteString(fmt.Sprintf(" %d", u.key))
	inOrder(u.right)
}

func execute(method string, num int) error {
	target := &node{
		key: num,
	}

	switch method {
	case "insert":
		insert(target)
		return nil
	case "find":
		if isExist(root, num) {
			fmt.Println("yes")
			return nil
		}
		fmt.Println("no")
		return nil
	case "print":
		inOrder(root)
		buf.WriteString("\n")
		preOrder(root)
		buf.WriteString("\n")
		buf.Flush()
		return nil
	case "delete":
		delete(find(root, num))
		return nil
	default:
		return errors.New("unexpected method")
	}
}

const (
	methodIndex = iota
	numIndex
)

func getMethodAndNumFromInput(input string) (method string, num int, err error) {
	s := strings.Split(input, " ")

	method = s[methodIndex]
	if len(s) == 2 {
		num, err = strconv.Atoi(s[numIndex])
		if err != nil {
			return "", -1, err
		}
	}

	return
}

var sc = bufio.NewScanner(os.Stdin)

func scanToInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	n := scanToInt()
	for i := 0; i < n; i++ {
		sc.Scan()
		input := sc.Text()
		method, num, err := getMethodAndNumFromInput(input)
		if err != nil {
			panic(err)
		}
		if err := execute(method, num); err != nil {
			panic(err)
		}
	}
}
