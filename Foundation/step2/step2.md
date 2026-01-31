# Go言語学習記録 - フロー制御構文

## 📝 やったこと

- A Tour of Go（Flow control statements）を学習

---

## 📚 学習メモ

### for文

```go
func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

**動作説明**:
- `i := 0` でカウンタを0で初期化
- `i++` でiが5になるまで実行して処理を終わりにする
- つまり条件式が`false`になった時に終了する

**書式**:
```go
for 初期値; 条件; 後処理 {
    // 処理内容
}
```

**各部分の説明**:
- **初期値**: forループのスコープ内のみ有効、記述は任意
- **条件**: `true`の間ループが実行される。条件を付けない場合は無限ループになるので注意
- **後処理**: イテレーション（各ループ時の処理）の最後に実行される、記述は任意

**補足**:
```go
// 条件のみのfor文（whileのように使える）
for i < 10 {
    fmt.Println(i)
    i++
}

// 無限ループ
for {
    // 処理
}
```

---

### if文

**書式**:
```go
if [変数名] := [変数の値]; [条件式] {
    // 処理内容
}
```

**特徴**:
- スコープ内で有効な変数を設定できる
- この変数はif文のスコープ内（および対応するelse節）でのみ使用可能

**例**:
```go
if v := math.Pow(x, n); v < lim {
    return v
} else {
    // vはここでも使える
    fmt.Printf("%g >= %g\n", v, lim)
}
// vはここでは使えない
```

---

### Switch文

**基本的な特徴**:
- if文と似たような構文
- 条件に指定した値に対してテストをしたい値を指定することで、if文のような条件分岐の処理を書くことができる
- テストの値を複数指定することもできる
- 条件を省略すると条件を`true`としたのと同じになる（`switch true`と同じ）

**Goのswitchの重要な特徴**:
- **自動break**: caseに合致したらそこでswitchスコープは終了して次に進む
- 他の言語（C、Javaなど）だと抜けるために`break`文が必要になるが、Goでは不要

**例**:
```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    fmt.Printf("%s.\n", os)
}
```

#### fallthroughキーワード

**使い方**:
- `fallthrough`キーワードを使用することで、caseに合致しても次のcaseに進むことができる
- 他の言語のデフォルト動作（breakがない状態）を再現できる

**重要な注意点**:
- `fallthrough`で次の処理に進むときは、**次のcaseの条件に合致するかに関わらず処理が実行される**
- `fallthrough`が記載された直後のcase文だけが実行される
- その次のcaseについては条件判定が行われ、当てはまらない場合は自動breakされる

**例**:
```go
switch num {
case 1:
    fmt.Println("1")
    fallthrough  // 次のcaseに強制的に進む
case 2:
    fmt.Println("2")  // num=1でもこれが実行される
case 3:
    fmt.Println("3")  // ここは条件判定される（num=1なので実行されない）
}
```

---

### defer文

**基本的な動作**:
- `defer`へ渡した関数の実行を、**呼び出し元の関数が終了するまで遅延**させる
- つまり、関数の最後に実行される

**重要な特徴 - スタック（LIFO: Last In First Out）**:
- 複数のdefer文を書いた場合、**後に宣言されたものから順に実行される**
- 最初に書かれたdeferが最も遅延されるため、最後に実行される
- 「何があっても最後に実行したい処理」にdeferを使う

**例**:
```go
func main() {
    defer fmt.Println("1番目に書いた defer")
    defer fmt.Println("2番目に書いた defer")
    defer fmt.Println("3番目に書いた defer")
    fmt.Println("通常の処理")
}

// 出力:
// 通常の処理
// 3番目に書いた defer
// 2番目に書いた defer
// 1番目に書いた defer
```

**主な用途**:
- ファイルのクローズ処理
- データベース接続の解放
- ロックの解放
など、リソースの後始末に使われる

---

## 🔍 参考リンク

- [フォーマット指定子についての参考記事](https://qiita.com/atsutama/items/81a7e1a72775910bcb8c)

---

## ❓ 疑問・TODO

### fmt.Printlnとprintlnの違い
**TODO**: 時間見つけて調べる

**簡易メモ（調査後に追記予定）**:
- `fmt.Println`: fmtパッケージの公式関数、本番コードで使用
- `println`: 組み込み関数、主にデバッグ用途、標準エラー出力に出力される場合がある
