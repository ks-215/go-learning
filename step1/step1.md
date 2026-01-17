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

**内容**: math.Sqrtはfloat64型が必要なのにint型が与えられている

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
- 文字（rune/byte）
- 文字列（string）
- ブール値（bool）
- 数値（int, float, complex など）

**注意**: 暗黙的な宣言である「`:=`」は使えない。

**特性**: 数値であり型の無い定数は、状況によって必要な型を取ることになる。

---

### ビットシフト演算
例をもとに解説：

```go
Big = 1 << 100   // 1を左に100ビットシフト（= 2の100乗）
Small = Big >> 99 // Bigを右に99ビットシフト（= 2の100乗 ÷ 2の99乗 = 2）
```

**ビットシフトの仕組み**:
- `<<` (左シフト): ビットを左に移動させる。1つシフトするごとに2倍になる
- `>>` (右シフト): ビットを右に移動させる。1つシフトするごとに2で割る
- `1 << 100` は実質的に 2^100 を意味する

**注意点**:
- Bigはint64型の場合は64bitに収まらないので`overflows`エラーが出る
- 型をfloat64にした場合は、指数表記ができるのでエラーにならない

---

### 指数表記（Floating-point representation）
64bit浮動小数点数（float64）を使って広範囲の数を表現する方法のこと。

#### IEEE 754形式の構造（float64の場合）
1. **符号ビット（1bit）**:
   - 正の数の場合は0、負の数の場合は1

2. **指数部（11bit）**:
   - 実際の指数に1023を足した値（バイアス表現）
   - 負の指数も正の数として表現できる仕組み
   - 例：指数が100なら、1023+100=1123を2進数で格納

3. **仮数部（52bit）**:
   - 有効数字の小数部分を格納
   - 暗黙の1があるため、実際は53bit精度（1.xxxxx...の形）

**ポイント**: 数値そのものを保存するのではなく、「符号 × 仮数 × 2^指数」という形式で数値の構造を保存する方法。

**表記例**: 
```
1.2676506002282295e+29
```
これは `1.2676506002282295 × 10^29` を意味する。「e+29」の部分は「10の29乗倍」という意味。

**2^100の場合**:
- 指数表記では約 `1.2676506002282294e+30` と表示される
- これは 2^100 ≈ 1.267... × 10^30 という意味
