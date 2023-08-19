package brew_program

import (
	"sync"

	"github.com/rs/zerolog/log"
)

const (
	LayoutSplashScreen  = "splash-screen"
	LayoutSchema        = "schema"
	LayoutHelper        = "helper"
	LayoutConfiguration = "configuration"
)

var (
	allowedLayouts = map[string]string{
		LayoutSplashScreen:  LayoutSplashScreen,
		LayoutSchema:        LayoutSchema,
		LayoutHelper:        LayoutHelper,
		LayoutConfiguration: LayoutConfiguration,
	}
)

type Layout struct {
	mux     *sync.RWMutex
	current string
}

func NewLayout() *Layout {
	return &Layout{
		mux: new(sync.RWMutex),
	}
}

func (l *Layout) Set(layout string) *Layout {
	l.mux.RLock()
	defer l.mux.RUnlock()

	if _, ok := allowedLayouts[layout]; !ok {
		log.Error().Str("layout", layout).Msg("layout is not allowed")

		return l
	}

	l.current = layout

	return l
}

func (l *Layout) Get() string {
	l.mux.RLock()
	defer l.mux.RUnlock()

	return l.current
}

func (l *Layout) Render() string {
	l.mux.RLock()
	defer l.mux.RUnlock()

	var layoutString string

	switch l.current {
	case LayoutSplashScreen:
		layoutString = l.splashScreen()
	case LayoutHelper:
		layoutString = l.helper()
	case LayoutSchema:
		layoutString = l.schema()
	case LayoutConfiguration:
		layoutString = l.configuration()
	}

	return layoutString
}

func (l *Layout) splashScreen() string {
	return `
╭─────────────────────────────────────╮
│                                     │
│                                     │
│                 (                   │
│                ) ) )                │
│               ( ( (                 │
│             │        │─╮            │
│             │ GoBrew │─╯            │
│             ╰────────╯              │
│                                     │
│                                     │
│                                     │
╰─────────────────────────────────────╯
`
}

func (l *Layout) schema() string {
	return `
╭────────────────┬──────────┬─────────╮  
│  steps         │  coffee  │ time    │   
│────────────────┼──────────┼─────────│
│  1st pour      │  60g     │  45s    │  
│  2nd pour      │  60g     │  45s    │  
│  3rd pour      │  60g     │  45s    │  
│  4th pour      │  60g     │  45s    │  
│  5th pour      │  60g     │  45s    │  
│                │          │         │  
│                │          │         │  
│                │          │         │  
│                │          │         │
╰────────────────┴──────────┴─────────╯
`
}

func (l *Layout) helper() string {
	return `
╭────────────────┬────────────────────╮  
│  steps         │                    │   
│  ■ 1st pour    │        ○  ●        │          
│  ■ 2nd pour    │     ○        *     │      
│  □ 3rd pour    │    ○   wait   *    │     
│                │    ○   pour   *    │     
│                │     ○        *     │      
│                │        ○  ○        │  
│                │                    │
│                │    120g => 180g    │  
│                │                    │
│                │       3m30s        │
╰────────────────┴────────────────────╯
* pour | ○ wait`
}

func (l *Layout) configuration() string {
	return `
╭─────────────────────────────────────╮
│  configuration                      │
│──────────────────┬──────────────────│
│  brew recipe     │   tetsu          │
│  ratio           │   1:16           │
│  coffee          │   60g            │
│  flavor          │   bright         │
│  concentration   │   strong         │
│                  │                  │
│                  │                  │
│                  │                  │
│                  │                  │
╰──────────────────┴──────────────────╯
`
}
