# runner

指定したディレクトリにある全ファイルに対して、指定したコマンドを実行するツール
（失敗したら5回までリトライする）

## 使い方

1. このパッケージをインポートして、Runner型のインスタンスを作成する
2. Runメソッドで実行する

```
package main

import (
	"github.com/kogai/runner"
)

func main() {
	r := runner.New("/path/to/dir", "your", "command", "--here")
	r.Run()
}
```
