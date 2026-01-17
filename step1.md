# Go言語学習記録 - エラーとつまずきポイント

## 📝 やったこと

- GitHubアカウント作成
- 最初のリポジトリ作成（[go-learning](https://github.com/ks-215/go-learning.git)）
- A Tour of Go を始める（Basicsのみ完了）

---

## ❌ エラー集

### main redeclared in this block
```
main redeclared in this blockcompilerDuplicateDecl hello.go(5, 6): other declaration of main
```

**内容**: mainが再定義されているとのエラー。実行はできた。

**原因**: 同じディレクトリ内で別のファイルにmain関数があるとエラーがでる。当時の状況で言えば、`packages.go`にmain関数を記載したが、同ディレクトリの`hello.go`にもmain関数があるため。

---

### variables() (no value) used as value
```
variables() (no value) used as value
```

**内容**: variables変数が値を返していないエラー

---

### undefined: {パッケージ名}
```
例）undefined: math.pi
```

**原因**:
- mathがインポートされていないか、piがエクスポートされていない
- mathがインポートされている場合、piの頭文字を大文字に変える必要がある（`math.Pi`）

**注意点**: インポートしたパッケージを使わないとコンパイルエラーになるため、VScodeは自動的に使用していないパッケージを削除する設定になっている。

---

### 型の不一致エラー
```
prog.go:10:28: cannot use x * x + y * y (value of type int) as float64 value in argument to math.Sqrt
```

**内容**: math.Sqrtはfloat型が必要なのにintが与えられている

```
prog.go:11:15: cannot use f (variable of type float64) as uint value in variable declaration
```

**内容**: fはfloat64型であり、uint型ではないよのエラー。型宣言ではuint型を指定しているのに初期値がfloat64型になっている。

---

### constant overflows int
```
constant overflows int
```

**内容**: int型に収まらないbit数によるエラー

---

## 📚 その他の学習メモ

### factored import statement
括弧でまとめられたインポート文のこと

```go
import (
    "fmt"
    "math"
)
```

---

### naked return
関数の戻り値となる変数に名前を付けて、名前を付けた変数をスコープ内に使うとreturnに何も書かなくて良くなる。

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // x, y を返す
}
```

**注意**: 長い関数では可読性が損なわれるため使わないように注意。

---

### 暗黙的な型宣言（Short variable declarations）
var宣言の代わりに「`:=`」の代入文を使うことで、型宣言ができる。

```go
// 明示的な型宣言
var i int = 42

// 暗黙的な型宣言
i := 42  // 型が推論される
```

⇒ 代入された値の型を宣言するってこと？明示的な型を指定しない場合は、代入する値から型推論されるよね

---

### ゼロ値
変数に初期値を与えずに宣言した場合の値のこと

| 型 | ゼロ値 |
|---|---|
| 数値型 | `0` |
| bool型 | `false` |
| string型 | `""` (空文字) |

---

### 定数（Constants）
`const`キーワードを使って変数と同じように宣言できる。プログラム実行中に値が変更されない固定値のこと。

**使用可能な型**:
- 文字（charset）
- 文字列（string）
- boolean
- 数値（numeric）

**注意**: 暗黙的な宣言である「`:=`」は使えない。

**特性**: 数値であり型の無い定数は、状況によって必要な型を取ることになる。

---

### ビットシフト演算
例をもとに解説：

```go
Big = 1 << 100   // 先頭から右に0を100付けた101ビットの二進数
Small = Big >> 99 // 末尾から左に0を99個外した2ビットの二進数(10=2)
```

**注意点**:
- Bigはint64型の場合は64bitに収まらないので`overflows`エラーが出る
- 型をfloat64にした場合は、指数表記ができるのでエラーにならない

---

### 指数表記（Floating-point representation）
64bitを使い分けて広範囲の数を表現する方法のこと。以下に分けられる：

#### 構造
1. **符号（1bit）**:
   - 正の数の場合は0、負の数の場合は1を付ける

2. **指数部（11bit）**:
   - 指数に1023（負の指数も表現できるための調整値）を足した数
   - その数を二進数に変換した箇所であり指数がいくつなのかを表す

3. **仮数部（52bit）**:
   - 指数がどれくらい大きいものなのかを表す
   - 例）2の100乗であれば、仮数部は100となる

**ポイント**: 数値そのものを保存するのではなく、数値の構造を保存する方法になっている。

**表記例**: 最終的に「有効数字+(e+n)」という形で表現される
```
1.2676506002282295e+29
```
「e+29」の部分は何桁あるかってこと
