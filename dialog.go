package main

import "github.com/mattn/go-gtk/gtk"

type dialogResult int

const (
	yes dialogResult = iota
	no
)

func handleDialog(conf *config) (dialogResult, string) {
	gtk.Init(nil)
	gtk.MainIteration()
	defer gtk.MainQuit()

	var msgType gtk.MessageType
	var buttons gtk.ButtonsType

	switch conf.msgType {
	case confirm:
		msgType = gtk.MESSAGE_QUESTION
		buttons = gtk.BUTTONS_YES_NO
	case info:
		msgType = gtk.MESSAGE_INFO
		buttons = gtk.BUTTONS_OK
	case warning:
		msgType = gtk.MESSAGE_WARNING
		buttons = gtk.BUTTONS_OK
	case password:
		msgType = gtk.MESSAGE_QUESTION
		buttons = gtk.BUTTONS_OK_CANCEL
	}

	dialog := gtk.NewMessageDialog(nil,
		gtk.DIALOG_MODAL|gtk.DIALOG_DESTROY_WITH_PARENT,
		msgType,
		buttons,
		conf.message)
	defer dialog.Destroy()

	dialog.SetTitle(conf.title)

	var input *gtk.Entry
	if conf.msgType == password {
		table := gtk.NewTable(1, 2, false)

		label := gtk.NewLabel("Password:")
		table.AttachDefaults(label, 0, 1, 0, 1)

		input = gtk.NewEntry()
		input.SetVisibility(false)
		input.SetActivatesDefault(true)
		table.AttachDefaults(input, 1, 2, 0, 1)

		dialog.GetVBox().PackEnd(table, true, true, 2)
		dialog.ShowAll()
	}

	response := dialog.Run()

	var result dialogResult
	if response == gtk.RESPONSE_YES || response == gtk.RESPONSE_OK {
		result = yes
	} else {
		result = no
	}

	if result == yes && input != nil {
		return result, input.GetText()
	}
	return result, ""
}
