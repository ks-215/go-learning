# Go言語学習記録 - More types

## 📝 やったこと

- A Tour of Go（More types）を学習
- ポインタ、構造体、配列、スライス、マップ、関数値について学習

---

## ❌ エラー集

### runtime.main_main·f: function main is undeclared in the main package
```
runtime.main_main·f: function main is undeclared in the main package
```

**原因**: packageの関数が存在していない。packageを`main`にしているので`func main`が必要。

---

### invalid argument: index out of bounds
```
invalid argument: index 2 out of bounds [0:2]
```

**原因**: 配列に要素が2個しかないのに3つ目の値を指定している。

**例**:
```go
arr := [2]int{1, 2}
fmt.Println(arr[2])  // エラー！インデックスは0と1のみ
```

---

### invalid append: argument must be a slice
```
invalid append: argument must be a slice; have arr1 (variable of type [3]int)
```

**原因**: arr1はサイズ3の配列にもかかわらず、`append`で4つ目を追加しようとしている。配列なので後から要素を追加/削除はできない。

**重要**: `append`はスライスに対してのみ使用可能。配列には使えない。

```go
arr := [3]int{1, 2, 3}
arr = append(arr, 4)  // エラー！配列には使えない

slice := []int{1, 2, 3}
slice = append(slice, 4)  // OK！スライスなら可能
```

---

### panic: runtime error: slice bounds out of range
```
panic: runtime error: slice bounds out of range [3:2]
```

**原因**: 配列の容量を超えたスライスでのアクセスによるエラー。値が無いから取れない。

**例**:
```go
arr := [2]int{1, 2}
s := arr[3:2]  // エラー！開始位置が終了位置より大きい
```

---

### tautological condition: nil == nil
```
tautological condition: nil == nil
```

**意味**: 「恒真条件（トートロジー）」という。if文などで`nil`であるかを判定するコードに発生。

**原因**: 今の書き方だと常に`nil`だからこのif文は意味がないというお知らせ。実行に支障はない。

```go
if nil == nil {  // 警告！これは常にtrue
    // 処理
}
```

---

### redundant type from array, slice, or map composite
```
redundant type from array, slice, or map composite
```

**意味**: 冗長な型指定になっているという警告。

**例**:
```go
// 冗長な書き方
m := map[string]Vertex{
    "Bell Labs": Vertex{40.68433, -74.39967},  // Vertexが冗長
}

// 推奨される書き方
m := map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},  // 型を省略できる
}
```

---

## 📚 学習メモ

### パッケージmainについて

**実行可能なプログラムの場合**:
- `package main`は必須
- `func main()`が必要
- `go run`や`go build`で実行するもの

**ライブラリの場合**:
- `package main`は不要
- 任意のパッケージ名を使用

---

### 変数宣言と代入の違いを再整理

| 演算子 | 意味 | 使用例 |
|--------|------|--------|
| `:=` | 宣言 + 代入 | `x := 10` |
| `=` | 代入のみ | `x = 20` |

**注意**: `:=`は関数内でのみ使用可能。パッケージレベルでは`var`を使用。

---

### 配列

Pythonのリストみたいなやつ。**固定長**の複数要素をまとめたリテラル。

**重要な特徴**:
- サイズも指定する必要がある
- **サイズも型の一部となる**

```go
var arr1 [3]int  // サイズ3の整数型配列
var arr2 [4]int  // サイズ4の整数型配列
```

`[3]int`と`[4]int`は**別の型**！
- サイズ3の数値型
- サイズ4の数値型

**宣言例**:
```go
// 方法1: サイズを明示
arr1 := [3]int{1, 2, 3}

// 方法2: 要素から自動でサイズを決定
arr2 := [...]int{1, 2, 3, 4, 5}

// 方法3: 一部のみ初期化（残りはゼロ値）
arr3 := [5]string{"a", "b"}  // {"a", "b", "", "", ""}
```

---

### スライス

**定義**: リストへの参照のようなもの。元の配列への「窓口」。

**重要な特徴**:
- 要素を変更すると、その元となる配列の対応する要素が変更される
- 最初の要素は含むが**最後の要素は除いた区間**

**例**:
```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]  // インデックス1, 2, 3を含むスライス（4は含まない）
// s = [2, 3, 4]
```

#### スライスの長さと容量

スライスは**長さ**と**容量**の両方を持っている：

| 概念 | 説明 |
|------|------|
| 長さ (`len`) | 要素の数 |
| 容量 (`cap`) | 現在の開始位置から元の配列の終わりまでの個数 |

**視覚的な説明**:
```
元の配列: [1, 2, 3, 4, 5]
           0  1  2  3  4  <- インデックス

s := arr[1:4]

スライスs:    [2, 3, 4]
               ↑      ↑
             start   end

len(s) = 3  (要素数)
cap(s) = 4  (インデックス1から配列の最後まで = 5-1 = 4)
```

#### PythonとGoのスライスの違い

| 言語 | 動作 |
|------|------|
| Python | スライスは**新しいリスト**を作る（コピー） |
| Go | 元の配列への**窓口**（参照） |

**⚠️ 重要**: スライスを変更すると元の配列も変わるので、`len`/`cap`の感覚をしっかりつけておく必要がある。

**例**:
```go
names := [4]string{"John", "Paul", "George", "Ringo"}

a := names[0:2]  // ["John", "Paul"]
b := names[1:3]  // ["Paul", "George"]

b[0] = "XXX"     // bを変更

fmt.Println(a)     // ["John", "XXX"]  ← aも変わる！
fmt.Println(names) // ["John", "XXX", "George", "Ringo"] ← 元の配列も変わる！
```

---

### 配列とスライスの違い

| 項目 | 配列 | スライス |
|------|------|----------|
| 長さ | 固定長 | 可変長 |
| 宣言 | `[3]int` | `[]int` |
| 追加/削除 | 不可 | 可能（`append`など） |
| 型 | サイズも型の一部 | 長さは型の一部ではない |

---

## 💻 コード例とサンプル

### ポインタ (sample1)

```go
func sample1() {
    i, j := 42, 2701

    p := &i         // i へのポインタを p に入れる
    fmt.Println(*p) // i のポインタを通して値を表示（42が表示される）

    *p = 21         // i のポインタを通して値を更新（i の値を直接更新）
    fmt.Println(i)  // 21が表示される

    p = &j          // j へのポインタを p に入れる
    *p = *p / 37    // ポインタを通して j の値を取得して37で割った値に更新
    fmt.Println(j)  // 73が表示される
}
```

**ポイント**:
- `&変数`: アドレス演算子（ポインタを取得）
- `*ポインタ`: デリファレンス演算子（ポインタの指す値を取得/変更）

---

### 構造体 (sample2, sample3)

```go
type Vertex struct {
    X int
    Y int
}

func sample2() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)  // 4
}

func sample3() {
    v := Vertex{1, 2}
    p := &v      // 構造体へのポインタ
    p.X = 1e9    // (*p).X の省略形
    fmt.Println(v)
}
```

**構造体リテラルの種類**:
```go
var (
    v1 = Vertex{1, 2}      // X:1, Y:2
    v2 = Vertex{X: 1}      // X:1, Y:0（Yはゼロ値）
    v3 = Vertex{}          // X:0, Y:0
    p  = &Vertex{1, 2}     // ポインタを返す
)
```

---

### スライスの操作 (sample8)

```go
func sample8() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

    s = s[:0]     // 長さ0にする
    printSlice(s) // len=0 cap=6 []

    s = s[:4]     // 長さを4に拡張
    printSlice(s) // len=4 cap=6 [2 3 5 7]

    s = s[2:]     // 最初の2要素をスキップ
    printSlice(s) // len=2 cap=4 [5 7]

    s = s[1:]     // さらに最初の要素をスキップ
    printSlice(s) // len=1 cap=3 [7]
}
```

**容量の変化に注目**:
- 開始位置を変えると容量も変わる
- `s[2:]`とすると、容量は元の配列の残り部分のみになる

---

### make関数でスライスを作成 (sample10)

```go
func sample10() {
    a := make([]int, 5)      // len=5 cap=5
    b := make([]int, 0, 5)   // len=0 cap=5
    c := b[:2]               // len=2 cap=5
    d := c[2:5]              // len=3 cap=3
}
```

**make関数の構文**:
```go
make([]T, length, capacity)
```

---

### append関数 (sample12)

```go
func sample12() {
    var s []int
    printSlice3(s)  // len=0 cap=0 []

    s = append(s, 0)        // len=1 cap=1 [0]
    s = append(s, 1)        // len=2 cap=2 [0 1]
    s = append(s, 2, 3, 4)  // len=5 cap=6 [0 1 2 3 4]
}
```

**容量の自動拡張**:
- スライスの容量が足りなくなると自動的に拡張される
- 通常、容量が2倍になる

---

### range (sample13, sample14)

```go
// インデックスと値の両方を取得
func sample13() {
    pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}

// インデックスのみ使用（値は不要）
func sample14() {
    pow := make([]int, 11)
    for i := range pow {
        pow[i] = 1 << uint(i)  // ビットシフト演算
    }
    for _, value := range pow {  // インデックスは不要
        fmt.Printf("%d\n", value)
    }
}
```

---

### マップ (sample16)

```go
type Vertex2 struct {
    Lat, Long float64
}

var m = map[string]Vertex2{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
```

**マップの操作**:
```go
// 作成
m := make(map[string]int)

// 追加・更新
m["key"] = 42

// 取得
value := m["key"]

// 削除
delete(m, "key")

// 存在確認
value, ok := m["key"]
if ok {
    // キーが存在する
}
```

---

### 関数値とクロージャ (sample17, fibona)

```go
// 関数を変数に代入
func sample17() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))  // 13

    // 関数を引数として渡す
    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}

// クロージャ（状態を保持する関数）
func fibona() func() int {
    a, b := 0, 1
    return func() int {
        result := a
        a, b = b, a+b
        return result
    }
}

func main() {
    f := fibona()
    for i := 0; i < 10; i++ {
        fmt.Println(f())  // 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
    }
}
```

**クロージャの特徴**:
- 外側の関数のローカル変数（`a`, `b`）を保持し続ける
- 呼び出すたびに状態が更新される

---

## 🎯 重要ポイントまとめ

### スライスで覚えておくべきこと

1. **スライスは参照**: 元の配列を変更してしまう
2. **len vs cap**: 
   - `len`: 現在の要素数
   - `cap`: 拡張可能な最大容量
3. **スライス範囲**: `a[low:high]` は`low`を含み`high`を含まない
4. **配列には使えない**: `append`はスライス専用

### 型の理解

- 配列: `[3]int` ← サイズも型の一部
- スライス: `[]int` ← サイズは型の一部ではない
- マップ: `map[string]int`

### nil スライス

```go
var s []int  // s == nil, len(s) == 0, cap(s) == 0
```

---

## 🔗 参考コード

完全なコードは以下の関数で確認可能:
- `sample1` - ポインタの基礎
- `sample2`, `sample3` - 構造体
- `sample4` - 配列
- `sample5` - スライスの参照特性
- `sample8` - スライスの長さと容量
- `sample10` - make関数
- `sample12` - append関数
- `sample17` - 関数値
- `fibona` - クロージャ
