package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/YOUR_NAME/password_generator/generator"
)

// InitUI はGUIを初期化し、ウィンドウに設定します。
func InitUI(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("パスワードジェネレーター")

	// パスワードの長さ入力
	lengthEntry := widget.NewEntry()
	lengthEntry.SetPlaceHolder("例: 12")

	// 文字種の選択肢
	charsetOptions := []string{
		"1. アルファベットのみ",
		"2. アルファベットと数字",
		"3. 数字のみ",
		"4. アルファベット、数字、記号",
	}
	charsetRadio := widget.NewRadioGroup(charsetOptions, nil)
	charsetRadio.Horizontal = false

	// 生成されたパスワードを表示するラベル
	passwordLabel := widget.NewLabel("生成されたパスワード: ")

	// 生成ボタン
	generateButton := widget.NewButton("パスワードを生成", func() {
		// パスワードの長さを取得
		lengthStr := lengthEntry.Text
		length, err := strconv.Atoi(lengthStr)
		if err != nil || length <= 0 {
			dialog.ShowError(fmt.Errorf("パスワードの文字数を選択してください"), myWindow)
			return
		}

		// 選択された文字種を取得
		if charsetRadio.Selected == "" {
			dialog.ShowError(fmt.Errorf("作成するパスワードのパターンを選択してください"), myWindow)
			return
		}

		var choice int
		switch charsetRadio.Selected {
		case charsetOptions[0]:
			choice = generator.AlphaOnly
		case charsetOptions[1]:
			choice = generator.AlphaNum
		case charsetOptions[2]:
			choice = generator.NumOnly
		case charsetOptions[3]:
			choice = generator.AlphaNumSymbol
		default:
			choice = 0
		}

		if choice == 0 {
			dialog.ShowError(fmt.Errorf("エラーです。作成するパスワードのパターンを再度選択してください。"), myWindow)
			return
		}

		// 文字セットを決定
		charset, err := generator.GetCharset(choice)
		if err != nil {
			dialog.ShowError(fmt.Errorf("文字セットの取得中にエラーが発生しました: %v", err), myWindow)
			return
		}

		// パスワードを生成
		password, err := generator.GeneratePassword(length, charset)
		if err != nil {
			dialog.ShowError(fmt.Errorf("パスワードの生成中にエラーが発生しました: %v", err), myWindow)
			return
		}

		// 生成されたパスワードを表示
		passwordLabel.SetText(fmt.Sprintf("生成されたパスワード: %s", password))
	})

	// レイアウトの作成
	form := container.NewVBox(
		widget.NewLabel("パスワードの長さを入力してください:"),
		lengthEntry,
		widget.NewLabel("文字種の選択肢:"),
		charsetRadio,
		generateButton,
		passwordLabel,
	)

	myWindow.SetContent(form)
	myWindow.Resize(fyne.NewSize(400, 300))
	return myWindow
}
