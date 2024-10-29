# これはなに
パスワードジェネレーターです。Go言語とFyneライブラリを使用して、カスタマイズ可能なパスワードを簡単に生成することができます。
生成するパスワードは数字や記号などの組み合わせを選択することができます。

# 前提条件（Prerequisites）
- Go: 1.18以上の安定板で動作確認済みです
- Fine: インストール必須です。インストールしておいてください。

# 準備（Preparation）
## ライブラリの準備
以下のGoライブラリをインストールする必要があるので行ってください
- Fyne:`go get fyne.io/fyne/v2`でダウンロードできます。
- その他の依存関係：`go mod tidy`で入手可能です
## モジュールパスの変更
- モジュールパス`module github.com/YOUR_NAME/password_generator`を適切な形に変更してください

# 使い方（How to Use）
1. **準備（Preparation）**
   - **パスワードの長さの指定**: アプリの入力フィールドにパスワードの長さ（例：12）を入力してください。
   - **文字種の選択**: 以下の4つのオプションから文字種を選択してください。
     - アルファベットのみ
     - アルファベットと数字
     - 数字のみ
     - アルファベット、数字、記号

2. **パスワードの生成（Generate Password）**
   - **パスワードの生成**: 「パスワードを生成」ボタンをクリックして、ランダ