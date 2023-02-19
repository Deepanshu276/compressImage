package ui

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Deepanshu276/imgcompress/cmd"
	"github.com/rivo/tview"
)

func ShowUI() {

	app := tview.NewApplication()

	inputField := tview.NewInputField().
		SetLabel("Input Image/URL: ").
		SetPlaceholder("Enter the input imag`e path or URL")

	outputField := tview.NewInputField().
		SetLabel("Output Image Path: ").
		SetPlaceholder("Enter the output image path")

	resizeField := tview.NewInputField().
		SetLabel("Resize Percentage: ").
		SetPlaceholder("Enter the resize percentage")

	form := tview.NewForm().
		AddFormItem(inputField).
		AddFormItem(outputField).
		AddFormItem(resizeField)

	form.AddButton("Compress", func() {
		input := inputField.GetText()
		output := outputField.GetText()
		resizeStr := resizeField.GetText()

		if input == "" {
			showError(app, "Input path cannot be empty")
			return
		}

		if resizeStr == "" {
			showError(app, "Resize percentage cannot be empty")
			return
		}

		resize, err := strconv.Atoi(resizeStr)
		if err != nil {
			showError(app, "Invalid resize percentage")
			return
		}
		if output != "" && output[0] == '~' {
			output, _ = filepath.Abs(filepath.Join(os.Getenv("HOME"), output[1:]))
		}

		if resize < 30 {
			modal := tview.NewModal().
				SetText(fmt.Sprintf("Resize percentage is greater than 70%%. Do you want to continue?[yellow]\n(Click 'yes' to continue or 'no' to go back)[-]")).
				AddButtons([]string{"Yes", "No"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonIndex == 0 {
						if isURL(input) {
							if output != "" {
								err = cmd.CompressImage(input, output, resize)
							} else {
								err = cmd.CompressImage(input, "", resize)
							}
							//err = cmd.CompressImage(input, output, resize)
						} else {
							absInput, err := filepath.Abs(filepath.Join("..", input))
							if err != nil {
								showError(app, err.Error())
								return
							}
							if output != "" {
								err = cmd.CompressImage(filepath.Join(absInput), output, resize)
							} else {
								err = cmd.CompressImage(filepath.Join(absInput), "", resize)
							}

							//err = cmd.CompressImage(filepath.Join(absInput), output, resize)

							if err != nil {
								showError(app, err.Error())
								return

							}
						}
						showMessage(app, "Image compressed successfully")
					} else {
						app.Stop()
					}
				})

			err = app.SetRoot(modal, true).SetFocus(modal).Run()
			if err != nil {
				panic(err)
			}
		} else {
			if isURL(input) {
				if output != "" {
					err = cmd.CompressImage(input, output, resize)
				} else {
					err = cmd.CompressImage(input, "", resize)
				}
				//err = cmd.CompressImage(input, output, resize)
			} else {
				absInput, err := filepath.Abs(filepath.Join("..", input))
				if err != nil {
					showError(app, err.Error())
					return
				}

				err = cmd.CompressImage(absInput, output, resize)

				if err != nil {
					showError(app, err.Error())
					return

				}
				if output != "" {
					err = cmd.CompressImage(absInput, output, resize)
				} else {
					err = cmd.CompressImage(absInput, "", resize)
				}

				if err != nil {
					showError(app, err.Error())
					return
				}
			}
			showMessage(app, "Image compressed successfully")
		}

	})

	/*if resize < 30 {
			modal := tview.NewModal().
				SetText(fmt.Sprintf("Resize percentage is greater than 70%%. Do you want to continue?[yellow]\n(Click 'yes' to continue or 'no' to go back)[-]")).
				AddButtons([]string{"Yes", "No"}).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonIndex == 0 {
						if isURL(input) {
							err = cmd.CompressImage(input, output, resize)
						} else {
							absInput, err := filepath.Abs(filepath.Join("..", input))
							if err != nil {
								showError(app, err.Error())
								return
							}
							err = cmd.CompressImage(filepath.Join(absInput, filepath.Base(input)), output, resize)
							if err != nil {
								showError(app, err.Error())
								return
							}
						}
						showMessage(app, "Image compressed successfully")
					} else {
						app.Stop()
					}
				})

			err = app.SetRoot(modal, true).SetFocus(modal).Run()
			if err != nil {
				panic(err)
			}
		} else {
			if isURL(input) {
				err = cmd.CompressImage(input, output, resize)
			} else {
				absInput, err := filepath.Abs(filepath.Join("..", input))
				if err != nil {
					showError(app, err.Error())
					return
				}
				err = cmd.CompressImage(absInput, output, resize)

				if err != nil {
					showError(app, err.Error())
					return
				}
			}
			showMessage(app, "Image compressed successfully")
		}

	})*/

	form.AddButton("Quit", func() {
		app.Stop()
	})

	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
func isURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func showError(app *tview.Application, msg string) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("[red]%s[-]", msg)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.Stop()
		})

	//app := tview.NewApplication()
	err := app.SetRoot(modal, true).SetFocus(modal).Run()
	if err != nil {
		panic(err)
	}
}

func showMessage(app *tview.Application, msg string) {
	modal := tview.NewModal().
		SetText(fmt.Sprintf("[green]%s[-]", msg)).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			app.Stop()
		})

	//app := tview.NewApplication()
	err := app.SetRoot(modal, true).SetFocus(modal).Run()
	if err != nil {
		panic(err)
	}
}
