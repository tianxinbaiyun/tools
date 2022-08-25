package utils

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"os"
	"strings"
	"github.com/flopp/go-findfont"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		switch  {
		case strings.Contains(path, "Songti.ttc"):
			fmt.Println(path)
			os.Setenv("FYNE_FONT", path)
			return
		case strings.Contains(path, "simkai.ttf"):
			fmt.Println(path)
			os.Setenv("FYNE_FONT", path)
			return
		}
	}
}

// NewApp 实例化
func NewApp()  {
	//新建一个app
	a := app.New()
	//新建一个窗口
	w := a.NewWindow("IP查询")
	//主界面框架布局
	MainShow(w)
	//尺寸
	w.Resize(fyne.Size{Width: 400, Height: 300})
	//w居中显示
	w.CenterOnScreen()
	//循环运行
	w.ShowAndRun()
	err := os.Unsetenv("FYNE_FONT")
	if err != nil {
		return
	}
}

var (
	ip = widget.NewLabel("IP ：")
	location = widget.NewLabel("LOCATION ：")
)

// MainShow 主界面函数
func MainShow(w fyne.Window) {
	w.SetContent(bindData())
}

func bindData() fyne.CanvasObject {
	formStruct := struct {
		IP string
	}{}

	formData := binding.BindStruct(&formStruct)
	form := newFormWithData(formData)
	form.OnSubmit = func() {
		value,_ := formData.GetValue("IP")
		info:= GetIpInfo(value.(string))
		if info.IP!=""{
			ip.SetText("IP ："+info.IP)
			location.SetText("LOCATION ："+info.Location)
		}
	}

	return container.NewBorder(container.NewVBox(form), nil, nil, nil, container.NewVBox(ip, location))
}

// newFormWithData newFormWithData
func newFormWithData(data binding.DataMap) *widget.Form {
	keys := data.Keys()
	items := make([]*widget.FormItem, len(keys))
	for i, k := range keys {
		data, err := data.GetItem(k)
		if err != nil {
			items[i] = widget.NewFormItem(k, widget.NewLabel(err.Error()))
		}
		items[i] = widget.NewFormItem(k, createBoundItem(data))
	}

	return widget.NewForm(items...)
}

// createBoundItem createBoundItem
func createBoundItem(v binding.DataItem) fyne.CanvasObject {
	switch val := v.(type) {
	case binding.Bool:
		return widget.NewCheckWithData("", val)
	case binding.Float:
		s := widget.NewSliderWithData(0, 1, val)
		s.Step = 0.01
		return s
	case binding.Int:
		return widget.NewEntryWithData(binding.IntToString(val))
	case binding.String:
		return widget.NewEntryWithData(val)
	default:
		return widget.NewLabel("")
	}
}
