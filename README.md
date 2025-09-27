# Prototypus AI (Go): High-Performance Data Transformation CLI

[![Status](https://img.shields.io/badge/Status-In%20Development-orange)](https://github.com/snknsk/prototypus-ai-go)
[![Language](https://img.shields.io/badge/Language-Go-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

### 🚀 概要 (About)

`Prototypus AI` は、**Go言語** と **`cobra`** ライブラリを使用して構築された高性能なコマンドラインツールです。主に、**Gemini API** を活用してファイルの読み込み、AIによるデータ変換や整形を高速に実行するために設計されています。

このプロジェクトは、Go言語の**標準的なCLI設計（Cobra）**と**パッケージ構造（`internal/`）**を習得し、配布しやすい単一バイナリを作成するために作成されました。

### 🛠️ 技術スタック (Tech Stack)

  * **メイン言語:** **Go 1.22+**
  * **CLI フレームワーク:** **`spf13/cobra`** (構造化されたCLIコマンドの構築)
  * **API クライアント:** **`github.com/google/generative-ai-go/genai`** (Go用 Gemini SDK)
  * **パッケージ管理:** **`go mod`**

-----

### 📦 インストールとセットアップ (Installation & Setup)

#### 1\. 前提条件

  * Git
  * Go 1.22 以上

#### 2\. 環境変数の設定

Gemini APIを利用するには、APIキーが必要です。プロジェクトの実行場所（通常はプロジェクトルート）に **`.env`** ファイルを作成し、キーを設定してください。

```bash
# .env ファイルの内容
GEMINI_API_KEY="YOUR_API_KEY_HERE"
```

#### 3\. プロジェクトのクローンとビルド

Go版は単一の実行ファイルをビルドすることで動作します。

```bash
# リポジトリをクローン
git clone https://github.com/snknsk/prototypus-ai-go
cd prototypus-ai-go

# 依存関係をダウンロード
go mod tidy

# 実行可能ファイルをビルド（プロジェクトルートに 'ptai' が作成される）
go build -o ptai ./cmd/ptai

# コマンド確認
./ptai hello
```

-----

### 🚀 使用方法 (Usage)

ビルドした実行ファイル **`./ptai`** でアプリケーションを実行できます。

#### 1\. データ変換 (`convert`)

このコマンドは、指定されたファイルを読み込み、AIに変換を依頼します。

```bash
echo '{"message": "Hello from the shell!", "status": "success", "code": 200}' > ./documents/data.md
./ptai convert \
    --input-path ./documents/data.md \
    --output-format json \
    --prompt "Extract all headings and their subsequent content into a single JSON array."
```

#### 2\. コマンドの詳細

| コマンド | 説明 | 主なフラグ |
| :--- | :--- | :--- |
| `./ptai convert` | ファイルをAI変換します。 | `--input-path`, `--output-format`, `--prompt` |
| `./ptai hello` | CLIが動作していることを確認します。 | なし |
| `./ptai --help` | **`ptai`** コマンド全体のヘルプを表示します。 | なし |

-----

### 📚 学習ポイント (Learning Notes)

  * **Go Standard Project Layout**: `cmd/ptai/main.go` をエントリーポイントとし、コアロジックを `internal/core` に分離する Goの標準的な構成を実践しました。
  * **Cobra CLI Design**: `cobra` を使用して、Pythonの `fire` と同等の機能を持つ構造化されたサブコマンド（`convert`）とフラグを定義する方法を習得しました。
  * **単一バイナリ配布**: `go build` により、依存関係を含んだ実行ファイル一つで動作するアプリケーションを作成する能力を習得しました。

-----

### 📜 ライセンス (License)

このプロジェクトは **MIT License** の下で公開されています。詳細については [`LICENSE`](https://www.google.com/search?q=LICENSE) ファイルを参照してください。
