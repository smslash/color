package logic

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Color string

const (
	ColorBlack      = "\u001b[30m"
	ColorRed        = "\u001b[31m"
	ColorGreen      = "\u001b[32m"
	ColorYellow     = "\u001b[33m"
	ColorBlue       = "\u001b[34m"
	ColorPurple     = "\u001b[35m"
	ColorAqua       = "\u001b[36m"
	ColorCyan       = "\033[36m"
	ColorOrange     = "\033[38;5;208m"
	ColorViolet     = "\033[38;5;129m"
	ColorLightGreen = "\033[38;5;46m"
	ColorGray       = "\033[37m"
	ColorBold       = "\033[1m"
	ColorReset      = "\u001b[0m"
)

var colors = map[string]Color{
	"black":       ColorBlack,
	"red":         ColorRed,
	"green":       ColorGreen,
	"yellow":      ColorYellow,
	"blue":        ColorBlue,
	"purple":      ColorPurple,
	"aqua":        ColorAqua,
	"cyan":        ColorCyan,
	"orange":      ColorOrange,
	"violet":      ColorViolet,
	"light-green": ColorLightGreen,
	"gray":        ColorGray,
	"bold":        ColorBold,
	"white":       ColorReset,
}

func colorize(color Color, message string) string {
	return fmt.Sprintf("%s%s%s", string(color), message, string(ColorReset))
}

func container(r rune, str string) bool {
	for _, v := range str {
		if r == v {
			return true
		}
	}
	return false
}

func Process(str string, data []byte) string {
	if len(os.Args) != 3 && len(os.Args) != 4 && len(os.Args) != 5 {
		fmt.Println("Worn't arguments")
		return ""
	}
	var colorFlag string
	flag.StringVar(&colorFlag, "color", "", "color for output")
	flag.Parse()
	if len(str) == 0 {
		return ""
	}

	s := ""
	for i := 0; i < len(str); i++ {
		if str[i] == 10 {
			s += "\\" + "n"
		} else {
			s += string(str[i])
		}
	}
	s += " "

	replaced := strings.ReplaceAll(string(data), "\n\n", "\n")
	splited := strings.Split(replaced, "\n")

	ascii := make(map[byte]int)
	var q byte
	for q = 32; q <= 126; q++ {
		ascii[q] = (int(q)-32)*8 + 1
	}

	var newLine []int
	counter := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == 92 && s[i+1] == 'n' {
			newLine = append(newLine, i)
			counter++
			i++
		}
	}
	var result []string
	if len(os.Args) == 4 {
		letters := os.Args[2]
		if len(letters) == 0 {
			letters = " "
		}
		if counter == 0 {
			for i := 0; i < 8; i++ {
				l := 0
				for j := 0; j < len(s)-1; j++ {
					if l+len(letters) < len(s) && s[l:l+len(letters)] == letters {
						for _, v := range letters {
							if color, ok := colors[colorFlag]; ok {
								result = append(result, colorize(color, string(splited[(ascii[byte(v)])+i])))
							} else {
								result = append(result, splited[ascii[byte(v)]+i])
							}
						}
						l += len(letters) - 1
					} else if l < len(s)-1 {
						result = append(result, splited[ascii[s[l]]+i])
					}
					l++
				}
				result = append(result, "\n")
			}
		} else {
			for i := 0; i < 8; i++ {
				l := 0
				for j := 0; j < len(s)-1; j++ {
					if l+1 < (len(s)-1) && s[l] == 92 && s[l+1] == 'n' {
						l++
						break
					}
					if l+len(letters) < len(s) && s[l:l+len(letters)] == letters {
						for _, v := range letters {
							if color, ok := colors[colorFlag]; ok {
								result = append(result, colorize(color, string(splited[(ascii[byte(v)])+i])))
							} else {
								result = append(result, splited[ascii[byte(v)]+i])
							}
						}
						l += len(letters) - 1
					} else {
						if l < len(s) {
							result = append(result, splited[ascii[s[l]]+i])
						}
					}
					l++
				}
				result = append(result, "\n")
			}
			for k := 0; k < len(newLine); k++ {
				for i := 0; i < 8; i++ {
					l := 0
					if l < len(s)-1 {
						l = newLine[k] + 2
					}
					for j := newLine[k] + 2; j < len(s)-1; j++ {
						if l+1 < (len(s)-1) && s[l] == 92 && s[l+1] == 'n' {
							l++
							break
						}
						if l+len(letters) < len(s) && s[l:l+len(letters)] == letters {
							for _, v := range letters {
								if color, ok := colors[colorFlag]; ok {
									result = append(result, colorize(color, string(splited[(ascii[byte(v)])+i])))
								} else {
									result = append(result, splited[ascii[byte(v)]+i])
								}
							}
							l += len(letters) - 1
						} else {
							if l < len(s) {
								result = append(result, splited[ascii[s[l]]+i])
							}
						}
						l++
					}
					result = append(result, "\n")
				}
			}
		}
	} else if len(os.Args) == 5 {
		letters := os.Args[2]
		if len(letters) == 0 {
			letters = " "
		}
		if counter == 0 {
			for i := 0; i < 8; i++ {
				l := 0
				for j := 0; j < len(s)-1; j++ {
					if l+len(letters) < len(s) && s[l:l+len(letters)] == letters {
						for _, v := range letters {
							if color, ok := colors[colorFlag]; ok {
								result = append(result, colorize(color, string(splited[(ascii[byte(v)])+i])))
							} else {
								result = append(result, splited[ascii[byte(v)]+i])
							}
						}
						l += len(letters) - 1
					} else if l < len(s)-1 {
						result = append(result, splited[ascii[s[l]]+i])
					}
					l++
				}
				result = append(result, "\n")
			}
		} else {
			for i := 0; i < 8; i++ {
				l := 0
				for j := 0; j < len(s)-1; j++ {
					if l+1 < (len(s)-1) && s[l] == 92 && s[l+1] == 'n' {
						l++
						break
					}
					if l+len(letters) < len(s) && s[l:l+len(letters)] == letters {
						for _, v := range letters {
							if color, ok := colors[colorFlag]; ok {
								result = append(result, colorize(color, string(splited[(ascii[byte(v)])+i])))
							} else {
								result = append(result, splited[ascii[byte(v)]+i])
							}
						}
						l += len(letters) - 1
					} else {
						if l < len(s) {
							result = append(result, splited[ascii[s[l]]+i])
						}
					}
					l++
				}
				result = append(result, "\n")
			}
			for k := 0; k < len(newLine); k++ {
				for i := 0; i < 8; i++ {
					l := 0
					if l < len(s)-1 {
						l = newLine[k] + 2
					}
					for j := newLine[k] + 2; j < len(s)-1; j++ {
						if l+1 < (len(s)-1) && s[l] == 92 && s[l+1] == 'n' {
							l++
							break
						}
						if l+len(letters) < len(s) && s[l:l+len(letters)] == letters {
							for _, v := range letters {
								if color, ok := colors[colorFlag]; ok {
									result = append(result, colorize(color, string(splited[(ascii[byte(v)])+i])))
								} else {
									result = append(result, splited[ascii[byte(v)]+i])
								}
							}
							l += len(letters) - 1
						} else {
							if l < len(s) {
								result = append(result, splited[ascii[s[l]]+i])
							}
						}
						l++
					}
					result = append(result, "\n")
				}
			}
		}
	} else if len(os.Args) == 3 {
		if counter == 0 {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(s)-1; j++ {
					if color, ok := colors[colorFlag]; ok {
						result = append(result, colorize(color, string(splited[(ascii[s[j]])+i])))
					} else {
						result = append(result, splited[ascii[s[j]]+i])
					}
				}
				result = append(result, "\n")
			}
		} else {
			for i := 0; i < 8; i++ {
				for j := 0; j < len(s)-1; j++ {
					if j+1 < (len(s)-1) && s[j] == 92 && s[j+1] == 'n' {
						break
					}
					if color, ok := colors[colorFlag]; ok {
						result = append(result, colorize(color, string(splited[(ascii[s[j]])+i])))
					} else {
						result = append(result, splited[ascii[s[j]]+i])
					}
				}
				result = append(result, "\n")
			}
			for k := 0; k < len(newLine); k++ {
				for i := 0; i < 8; i++ {
					for j := newLine[k] + 2; j < len(s)-1; j++ {
						if j+1 < (len(s)-1) && s[j] == 92 && s[j+1] == 'n' {
							break
						}
						if color, ok := colors[colorFlag]; ok {
							result = append(result, colorize(color, string(splited[(ascii[s[j]])+i])))
						} else {
							result = append(result, splited[ascii[s[j]]+i])
						}
					}
					result = append(result, "\n")
				}
			}
		}
	}

	res := ""
	for i := 0; i < len(result); i++ {
		res += result[i]
	}

	text := strings.ReplaceAll(res, "\n\n\n\n\n\n\n\n", "\n")

	new_line_str := ""
	if counter == 0 {
		return text
	} else if counter == (len(s)-1)/2 {
		for i := 0; i < counter; i++ {
			new_line_str += "\n"
		}
		return new_line_str
	}
	return text
}
