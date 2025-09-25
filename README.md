## 📄 README.md 構成案 (Prototypus AI)

### \# Prototypus AI: The Data Transformation Tool

> **[ここにプロジェクトのバッジやロゴを挿入]** 例: 

### 🚀 概要 (About)

`Prototypus AI` は、**Gemini API** を活用したコマンドラインツールです。様々な形式の入力ファイル（Markdown、CSV、JSONなど）を読み込み、ユーザーの指定したプロンプトと出力形式に基づいて、AIによるデータ変換や整形を実行することを目的としています。

このプロジェクトは、PythonとGoという異なる言語で**同一のCLI構造を再現し、両者の開発手法とエコシステムを比較学習する**ために作成されました。

| 言語 | CLI コマンド | 特徴 |
| :--- | :--- | :--- |
| **Python** | `ptai` | `python-fire` による簡単なCLI定義。仮想環境での管理。 |
| **Go** | `./ptai` | `cobra` による構造化されたCLI定義。単一バイナリでの配布。 |

-----

### 🛠️ 技術スタック (Tech Stack)

このプロジェクトを構築するために使用した主要な技術とライブラリです。

  * **メイン言語:** Python 3.12+ / Go 1.22+
  * **CLI フレームワーク:** `python-fire` (Python) / `spf13/cobra` (Go)
  * **AI SDK:** `google-generativeai` (Python) / `github.com/google/generative-ai-go/genai` (Go)
  * **依存管理:** `pip` & `pyproject.toml` (Python) / `go mod` (Go)

-----

### 📦 インストールとセットアップ (Installation & Setup)

#### 1\. 前提条件

  * Git
  * Python 3.12 以上
  * Go 1.22 以上

#### 2\. 環境変数の設定

Gemini APIを利用するには、APIキーが必要です。プロジェクトルートに **`.env`** ファイルを作成し、キーを設定してください。

```bash
# .env ファイルの内容
GEMINI_API_KEY="YOUR_API_KEY_HERE"
```

#### 3\. Python 版のセットアップ (ptai)

```bash
# 仮想環境を作成・有効化
python -m venv .venv
source .venv/bin/activate

# 開発モードでインストール
pip install -e .

# コマンド確認
ptai hello
```

#### 4\. Go 版のセットアップ (./ptai)

```bash
# プロジェクトルートで実行
go build -o ptai ./cmd/ptai

# コマンド確認
./ptai convert --help
```

-----

### 🚀 使用方法 (Usage)

#### 1\. データ変換 (`convert`)

Markdownファイルを読み込み、指定されたプロンプトでJSON形式に変換します。

```bash
# Go 版の実行例
./ptai convert \
    --input-path notes.md \
    --output-format json \
    --prompt "Convert the markdown into a list of JSON objects, summarizing each section."
```

#### 2\. その他のコマンド

| コマンド | 説明 |
| :--- | :--- |
| `ptai hello` | シンプルな挨拶を表示し、CLIが正しく動作しているか確認します。 |
| `ptai byebye --message "See ya"` | メッセージを表示してプログラムを終了します。 |

-----

### 📚 学習ポイント (Learning Notes)

**(※ このセクションは、特に勉強目的のプロジェクトで非常に有用です)**

  * **CLI構造の設計**: Python (クラスベース) と Go (Cobraツリー構造) で、CLIのサブコマンドとフラグ（引数）をどのように定義するかを比較しました。
  * **モジュール分離**: `cli.py` / `main.go` から、コアロジックを `core/gemini_client` へ委譲する設計パターンを習得しました。
  * **エラーハンドリング**: ファイルの存在確認や、コアロジックからのエラー（`PerformConversion` の `return "", err`）を CLI のエントリーポイントで捕捉し、ユーザーフレンドリーなメッセージとして出力する練習をしました。

-----

### 📜 ライセンス (License)

このプロジェクトは **MIT License** の下で公開されています。詳細については [`LICENSE`](https://www.google.com/search?q=LICENSE) ファイルを参照してください。
