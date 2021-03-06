package version

import "github.com/esimov/diagram/color"

// Diagram CLI logo
const Logo = `
  ██████╗ ██╗ █████╗  ██████╗ ██████╗  █████╗ ███╗   ███╗
  ██╔══██╗██║██╔══██╗██╔════╝ ██╔══██╗██╔══██╗████╗ ████║
  ██║  ██║██║███████║██║  ███╗██████╔╝███████║██╔████╔██║
  ██║  ██║██║██╔══██║██║   ██║██╔══██╗██╔══██║██║╚██╔╝██║
  ██████╔╝██║██║  ██║╚██████╔╝██║  ██║██║  ██║██║ ╚═╝ ██║
  ╚═════╝ ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝`

// Name of application
const Name = "diagram"

// Description of application
const Description = "Transform ASCII texts to hand drawing diagrams"

// Version number
const Version = "v.1.0.1"

//ColorLogo with color
func DrawLogo() string {
	var logo string

	logo += "\n\n"
	logo += color.StringRandom("  ██████╗ ██╗ █████╗  ██████╗ ██████╗  █████╗ ███╗   ███╗\n")
	logo += color.StringRandom("  ██╔══██╗██║██╔══██╗██╔════╝ ██╔══██╗██╔══██╗████╗ ████║\n")
	logo += color.StringRandom("  ██║  ██║██║███████║██║  ███╗██████╔╝███████║██╔████╔██║\n")
	logo += color.StringRandom("  ██║  ██║██║██╔══██║██║   ██║██╔══██╗██╔══██║██║╚██╔╝██║\n")
	logo += color.StringRandom("  ██████╔╝██║██║  ██║╚██████╔╝██║  ██║██║  ██║██║ ╚═╝ ██║\n")
	logo += color.StringRandom("  ╚═════╝ ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝ " + Version)

	return logo
}