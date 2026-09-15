package components

import "charm.land/lipgloss/v2"

const LOGO = `                         
 ▄▄▄▄▄                 ▄▄                        
██▀▀▀▀█▄                ██                       
▀██▄  ▄▀ ▄              ██   ▄                   
  ▀██▄▄  ███▄███▄ ▄███▄ ██   ███▄███▄ ▄▀▀█▄ ████▄
▄   ▀██▄ ██ ██ ██ ██ ██ ██   ██ ██ ██ ▄█▀██ ██ ██
▀██████▀▄██ ██ ▀█▄▀███▀▄██  ▄██ ██ ▀█▄▀█▄██▄████▀
                                            ██   
                                            ▀    
   Scuffed way to search .tiny (v1) mappings																						
`

var LogoCenterAligned = lipgloss.NewStyle().
	AlignHorizontal(0.4).
	Width(MAX_PAGE_WIDTH)
