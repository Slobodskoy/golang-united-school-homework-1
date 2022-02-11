package main

import (
	"fmt"
	"strings"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return fmt.Sprintf("%s!",
		strings.TrimRight(emoji.Sprint("Hello :world_map:"), " "))
}
