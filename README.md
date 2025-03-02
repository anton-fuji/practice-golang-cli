# Go CLI Development Practice with Cobra

---

## 📌 インストール
```sh
go install github.com/spf13/cobra-cli@latest
```

---

## 🛠️ プロジェクトの作成
```sh
cobra-cli init mycli
cd mycli
go mod tidy
```

---

## 📌 コマンドの追加
```sh
cobra-cli add hello
```
`cmd/hello.go` を編集：
```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "Say hello",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello, Cobra!")
    },
}

func init() {
    rootCmd.AddCommand(helloCmd)
}
```

実行：
```sh
go run main.go hello
```
出力：
```
Hello, Cobra!
```

---

## 📂 コマンドの階層構造
サブコマンドを作成する：
```sh
cobra-cli add user
cobra-cli add create --parent user
```

---

##  フラグの利用
フラグを追加する場合：
```go
var name string

var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "Say hello",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Hello, %s!\n", name)
    },
}

func init() {
    rootCmd.AddCommand(helloCmd)
    helloCmd.Flags().StringVarP(&name, "name", "n", "World", "Your name")
}
```
実行：
```sh
go run main.go hello --name Alice
```
出力：
```
Hello, Alice!
```

---

## ⚙️ 設定ファイルの利用（Viper）
`config.yaml` を用意：
```yaml
name: Alice
```
`cmd/hello.go` を修正：
```go
import (
    "github.com/spf13/viper"
)

func init() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.ReadInConfig()
    rootCmd.AddCommand(helloCmd)
}

var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "Say hello from config",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Hello, %s!\n", viper.GetString("name"))
    },
}
```
実行：
```sh
go run main.go hello
```
出力：
```
Hello, Alice!
```

---

## 🔧 シェル補完の追加
```sh
cobra-cli add completion
```
`cmd/completion.go` を編集：
```go
var completionCmd = &cobra.Command{
    Use:   "completion",
    Short: "Generate shell completion script",
    Run: func(cmd *cobra.Command, args []string) {
        rootCmd.GenBashCompletion(os.Stdout)
    },
}
```
Bash 用の補完を有効化：
```sh
source <(go run main.go completion)
```

---

## 🚀 例：GitHub リポジトリ情報取得 CLI
プロジェクトの作成：
```sh
cobra-cli init ghcli
cd ghcli
cobra-cli add repo
```
`cmd/repo.go` を編集：
```go
package cmd

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
    Use:   "repo [owner] [repo]",
    Short: "Get repository info from GitHub",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        owner, repo := args[0], args[1]
        url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

        resp, err := http.Get(url)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        defer resp.Body.Close()

        var result map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&result)

        fmt.Println("Repository Name:", result["full_name"])
        fmt.Println("Description:", result["description"])
        fmt.Println("Stars:", result["stargazers_count"])
    },
}

func init() {
    rootCmd.AddCommand(repoCmd)
}
```
実行：
```sh
go run main.go repo torvalds linux
```
出力：
```
Repository Name: torvalds/linux
Description: Linux kernel source tree
Stars: 150000
```

---

## 📌 まとめ
1. `cobra-cli` をインストールし、CLI プロジェクトを作成
2. コマンドとサブコマンドを作成
3. フラグを利用してコマンドの動作をカスタマイズ
4. `viper` を統合して設定ファイルを活用
5. シェル補完を追加
6. 実践的な CLI ツールを開発

Cobra を活用して、Go で強力な CLI アプリを開発しましょう！ 🚀

