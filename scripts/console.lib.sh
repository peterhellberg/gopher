################################################################################
# An overly simple set of functions for displaying color and style in a console
# By Thomas.Cherry@gmail.com
# Simply copy into any project to add color.
################################################################################

# Wrap a color number in a terminal escape sequence
function colorwrap() {
	color=$1
	printf "\033[0;${color}m"
}

# Wrap a style number in a terminal escape sequence
function stylewrap() {
	style=$1
	printf "\033[${style}m"
}

################################################################################
# With 8 colors you can support the color models:
# RGB (light), RYG (traffic lights), CMYK (printing)

# FOREGROUND 					; background
# ----------------------------- ; -----------------------------
BLACK=30 						; black=40
RED=31 							; red=41
GREEN=32 						; green=42
YELLOW=33						; yellow=43
BLUE=34							; blue=44
MAGENTA=35 						; magenta=45
CYAN=36 						; cyan=46
WHITE=37 						; white=47

# Escaped colors, F for foreground, B for background
FBLACK=$(colorwrap $BLACK)		; BBLACK=$(colorwrap $black)
FRED=$(colorwrap $RED)			; BRED=$(colorwrap $red)
FGREEN=$(colorwrap $GREEN)		; BGREEN=$(colorwrap $green)
FYELLOW=$(colorwrap $YELLOW)	; BYELLOW=$(colorwrap $yellow)
FBLUE=$(colorwrap $BLUE) 		; BBLUE=$(colorwrap $blue)
FMAGENTA=$(colorwrap $MAGENTA)	; BMAGENTA=$(colorwrap $magenta)
FCYAN=$(colorwrap $CYAN) 		; BCYAN=$(colorwrap $cyan)
FWHITE=$(colorwrap $WHITE) 		; BWHITE=$(colorwrap $white)

# Style codes and their escaped values
nc=0 		; NC=$(stylewrap ${nc})
bold=1 		; BOLD=$(stylewrap ${bold})
faint=2		; FAINT=$(stylewrap ${faint})
italix=3	; ITALIX=$(stylewrap ${italix})		# not well supported
under=4 	; UNDER=$(stylewrap ${under})
blink=5 	; BLINK=$(stylewrap ${blink})
fast=6 		; FAST=$(stylewrap ${fast})			# not well supported
inverse=7 	; INVERSE=$(stylewrap ${inverse})

################################################################################
# Functions for use externally

# Color print with style using just color numbers
function cprints() {
	color=$1	# Example: $RED
	bg=$2		# Example: $green
	style=$3	# Example: $bold
	msg="$4"
	printf "\033[${color};${bg};${style}m%s\033[0m" "$msg"
}

# Color print with style followed by a new line
function cprintsln() {
	cprints "$1" "$2"
	echo
}

# Color print, can mix styles, but can not mix foreground and background
function cprint() {
	color="$1"
	msg="$2"
	if tty -s ; then
		printf "${color}%s${NC}" "$msg"
	else
		printf "%s" "$msg"
	fi
}

# Color print followed by a new line
function cprintln() {
	cprint "$1" "$2"
	echo
}
