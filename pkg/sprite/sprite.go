package sprite

type Sprites struct{}

func New() Sprites {
	return Sprites{}
}

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

func RJ() string {
	return `                                                ○
                                                ║
██████████████▙                ╔═█        █████████████
█ ║          ║ █               ║ █    ()═[███▄█████▄███]═()
█ ║          ║ █               ║ █        ███▄▄▄▄▄▄▄███
██████████████▛    █═╗         ║ █      ▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅▅
█ ║          ╚═█▙   ▜███████████▛       ███ ▒ ▒ ░░░░░████`
}
