package redui

import (
	"github.com/abdfnx/redui/redis"
	"github.com/abdfnx/redui/config"
	"github.com/abdfnx/redui/core"

	"github.com/rivo/tview"
)

type primitiveKey struct {
	Primitive tview.Primitive
	Key       string
}

// RedUI is a redis console user interface
type RedUI struct {
	metaPanel           *tview.TextView
	mainPanel           *tview.Flex
	outputPanel         *tview.List
	keyItemsPanel       *tview.List
	summaryPanel        *tview.TextView
	findPanel           *tview.InputField
	helpPanel           *tview.Flex
	helpMessagePanel    *tview.TextView
	helpServerInfoPanel *tview.TextView

	commandPanel        *tview.Flex
	commandInputField   *tview.InputField
	commandResultPanel  *tview.TextView
	commandMode         bool

	leftPanel  			*tview.Flex
	rightPanel 			*tview.Flex

	layout 				*tview.Flex
	pages  				*tview.Pages
	app    				*tview.Application

	redisClient      	redis.RedisClient
	outputChan       	chan core.OutputMessage
	uiViewUpdateChan 	chan func()

	itemSelectedHandler func(index int, key string) func()

	maxKeyLimit       	int
	maxCharacterLimit 	int

	version   			string
	buildDate 		    string

	focusPrimitives     []primitiveKey
	currentFocusIndex   int

	config      		config.Config
	keyBindings         core.KeyBindings

	findKeyHistories    []string
	commandKeyHistories []string
}
