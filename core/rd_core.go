package core

import (
	"github.com/gdamore/tcell"
)

type KeyBindings map[string][]tcell.Key

var keyBindings = KeyBindings{
	"search":           {tcell.KeyF2, tcell.KeyCtrlS},
	"keys":             {tcell.KeyF3, tcell.KeyCtrlK},
	"key_list_value":   {tcell.KeyF6, tcell.KeyCtrlY},
	"key_string_value": {tcell.KeyF7, tcell.KeyCtrlA},
	"key_hash":         {tcell.KeyF6, tcell.KeyCtrlY},
	"result":           {tcell.KeyF9, tcell.KeyCtrlO},
	"command":          {tcell.KeyF1, tcell.KeyCtrlN},
	"command_focus":    {tcell.KeyF4, tcell.KeyCtrlF},
	"command_result":   {tcell.KeyF5, tcell.KeyCtrlR},
	"quit":             {tcell.KeyEsc, tcell.KeyCtrlQ},
	"switch_focus":     {tcell.KeyTab},
}

func NewKeyBinding() KeyBindings {
	return keyBindings
}
