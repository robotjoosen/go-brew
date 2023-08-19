package brew_program

type Sprites struct{}

func (s Sprites) Logo() string {
	return `╭─────╮      ╭────╮              
│  ╭──│───╮  │ ── ╯╭───╮───╮─╮─╮─╮
│  │  │ . │  │ ─── │  ─╯ ──╯ │ │ │ ╭─╮
╰─────╰───╯  ╰─────╯─╯ ╯───╯─────╯ ╰─╯`
}

func (s Sprites) CoffeeAnimated(frame int) string {
	frames := []string{
		` 
       (
   (    }
  {  } (
`, `
       (
   {    }   
  (  ) (
`, `
       {
   (    )   
  {  ) (
`, `
       (
   (    )   
  (  } {
`,
	}

	return frames[frame%len(frames)] + `
├~~~~~~~~┤
│        │─╮
│ GoBrew │─╯
╰────────╯`
}
