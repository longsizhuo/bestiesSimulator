package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

//go:embed Source/HarmonyOS_Sans_SC_Regular.ttf
var customFont []byte

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "newFont",
		StaticContent: customFont,
	}
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("闺闺模拟器")
	myWindow.Resize(fyne.NewSize(400, 400))
	outputLabel := widget.NewLabel("\n")

	nameLabel := widget.NewLabel("闺闺名字")
	entry := widget.NewEntry()
	entry.SetPlaceHolder("输入闺闺名字")

	content := container.NewVBox(entry)
	name := entry.Text
	confirmCallback := func(response bool) {
		if response {
			name = entry.Text
			nameLabel.SetText("闺闺名字: " + name)
		}
	}

	dialog.ShowCustomConfirm("输入闺闺名字", "确认", "取消", content, confirmCallback, myWindow)

	submitButton := widget.NewButton("闺闺伤我心了", func() {
		beBlack(outputLabel, name)
	})
	submitButton2 := widget.NewButton("闺闺回来了", func() {
		beWhite(outputLabel, name)
	})

	resetButton := widget.NewButton("重新开始", func() {
		entry.SetText("")         // 清空输入框
		nameLabel.SetText("闺闺名字") // 清空输出标签
	})

	content = container.NewVBox(
		nameLabel,
		submitButton,
		submitButton2,
		outputLabel,
		resetButton,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}

func updateLoveMessage(output *widget.Label, start int, end int, step int, finalMessage string, name string) {
	var text string
	for i := start; i != end; i -= step {
		text += fmt.Sprintf("爱%s%%%d\n\n", name, i*10)
		output.Text += text
	}
	text += finalMessage
	output.SetText(text)
}

func beBlack(output *widget.Label, name string) {
	updateLoveMessage(output, 10, -1, 1, "能量不够，马上关机", name)
}

func beWhite(output *widget.Label, name string) {
	updateLoveMessage(output, 0, 11, -1, "你的闺闺回来了", name)
}
