// Copyright 2013 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

package ui

import (
	"fmt"
)

type Key int

const (
	KeyA Key = iota
	KeyB
	KeyC
	KeyD
	KeyE
	KeyF
	KeyG
	KeyH
	KeyI
	KeyJ
	KeyK
	KeyL
	KeyM
	KeyN
	KeyO
	KeyP
	KeyQ
	KeyR
	KeyS
	KeyT
	KeyU
	KeyV
	KeyW
	KeyX
	KeyY
	KeyZ
	KeyAltLeft
	KeyAltRight
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyArrowUp
	KeyBackquote
	KeyBackslash
	KeyBackspace
	KeyBracketLeft
	KeyBracketRight
	KeyCapsLock
	KeyComma
	KeyContextMenu
	KeyControlLeft
	KeyControlRight
	KeyDelete
	KeyDigit0
	KeyDigit1
	KeyDigit2
	KeyDigit3
	KeyDigit4
	KeyDigit5
	KeyDigit6
	KeyDigit7
	KeyDigit8
	KeyDigit9
	KeyEnd
	KeyEnter
	KeyEqual
	KeyEscape
	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyHome
	KeyInsert
	KeyIntlBackslash
	KeyMetaLeft
	KeyMetaRight
	KeyMinus
	KeyNumLock
	KeyNumpad0
	KeyNumpad1
	KeyNumpad2
	KeyNumpad3
	KeyNumpad4
	KeyNumpad5
	KeyNumpad6
	KeyNumpad7
	KeyNumpad8
	KeyNumpad9
	KeyNumpadAdd
	KeyNumpadDecimal
	KeyNumpadDivide
	KeyNumpadEnter
	KeyNumpadEqual
	KeyNumpadMultiply
	KeyNumpadSubtract
	KeyPageDown
	KeyPageUp
	KeyPause
	KeyPeriod
	KeyPrintScreen
	KeyQuote
	KeyScrollLock
	KeySemicolon
	KeyShiftLeft
	KeyShiftRight
	KeySlash
	KeySpace
	KeyTab
	KeyReserved0
	KeyReserved1
	KeyReserved2
	KeyReserved3
	KeyMax = KeyReserved3
)

func (k Key) String() string {
	switch k {
	case KeyA:
		return "KeyA"
	case KeyB:
		return "KeyB"
	case KeyC:
		return "KeyC"
	case KeyD:
		return "KeyD"
	case KeyE:
		return "KeyE"
	case KeyF:
		return "KeyF"
	case KeyG:
		return "KeyG"
	case KeyH:
		return "KeyH"
	case KeyI:
		return "KeyI"
	case KeyJ:
		return "KeyJ"
	case KeyK:
		return "KeyK"
	case KeyL:
		return "KeyL"
	case KeyM:
		return "KeyM"
	case KeyN:
		return "KeyN"
	case KeyO:
		return "KeyO"
	case KeyP:
		return "KeyP"
	case KeyQ:
		return "KeyQ"
	case KeyR:
		return "KeyR"
	case KeyS:
		return "KeyS"
	case KeyT:
		return "KeyT"
	case KeyU:
		return "KeyU"
	case KeyV:
		return "KeyV"
	case KeyW:
		return "KeyW"
	case KeyX:
		return "KeyX"
	case KeyY:
		return "KeyY"
	case KeyZ:
		return "KeyZ"
	case KeyAltLeft:
		return "KeyAltLeft"
	case KeyAltRight:
		return "KeyAltRight"
	case KeyArrowDown:
		return "KeyArrowDown"
	case KeyArrowLeft:
		return "KeyArrowLeft"
	case KeyArrowRight:
		return "KeyArrowRight"
	case KeyArrowUp:
		return "KeyArrowUp"
	case KeyBackquote:
		return "KeyBackquote"
	case KeyBackslash:
		return "KeyBackslash"
	case KeyBackspace:
		return "KeyBackspace"
	case KeyBracketLeft:
		return "KeyBracketLeft"
	case KeyBracketRight:
		return "KeyBracketRight"
	case KeyCapsLock:
		return "KeyCapsLock"
	case KeyComma:
		return "KeyComma"
	case KeyContextMenu:
		return "KeyContextMenu"
	case KeyControlLeft:
		return "KeyControlLeft"
	case KeyControlRight:
		return "KeyControlRight"
	case KeyDelete:
		return "KeyDelete"
	case KeyDigit0:
		return "KeyDigit0"
	case KeyDigit1:
		return "KeyDigit1"
	case KeyDigit2:
		return "KeyDigit2"
	case KeyDigit3:
		return "KeyDigit3"
	case KeyDigit4:
		return "KeyDigit4"
	case KeyDigit5:
		return "KeyDigit5"
	case KeyDigit6:
		return "KeyDigit6"
	case KeyDigit7:
		return "KeyDigit7"
	case KeyDigit8:
		return "KeyDigit8"
	case KeyDigit9:
		return "KeyDigit9"
	case KeyEnd:
		return "KeyEnd"
	case KeyEnter:
		return "KeyEnter"
	case KeyEqual:
		return "KeyEqual"
	case KeyEscape:
		return "KeyEscape"
	case KeyF1:
		return "KeyF1"
	case KeyF2:
		return "KeyF2"
	case KeyF3:
		return "KeyF3"
	case KeyF4:
		return "KeyF4"
	case KeyF5:
		return "KeyF5"
	case KeyF6:
		return "KeyF6"
	case KeyF7:
		return "KeyF7"
	case KeyF8:
		return "KeyF8"
	case KeyF9:
		return "KeyF9"
	case KeyF10:
		return "KeyF10"
	case KeyF11:
		return "KeyF11"
	case KeyF12:
		return "KeyF12"
	case KeyF13:
		return "KeyF13"
	case KeyF14:
		return "KeyF14"
	case KeyF15:
		return "KeyF15"
	case KeyF16:
		return "KeyF16"
	case KeyF17:
		return "KeyF17"
	case KeyF18:
		return "KeyF18"
	case KeyF19:
		return "KeyF19"
	case KeyF20:
		return "KeyF20"
	case KeyF21:
		return "KeyF21"
	case KeyF22:
		return "KeyF22"
	case KeyF23:
		return "KeyF23"
	case KeyF24:
		return "KeyF24"
	case KeyHome:
		return "KeyHome"
	case KeyInsert:
		return "KeyInsert"
	case KeyIntlBackslash:
		return "KeyIntlBackslash"
	case KeyMetaLeft:
		return "KeyMetaLeft"
	case KeyMetaRight:
		return "KeyMetaRight"
	case KeyMinus:
		return "KeyMinus"
	case KeyNumLock:
		return "KeyNumLock"
	case KeyNumpad0:
		return "KeyNumpad0"
	case KeyNumpad1:
		return "KeyNumpad1"
	case KeyNumpad2:
		return "KeyNumpad2"
	case KeyNumpad3:
		return "KeyNumpad3"
	case KeyNumpad4:
		return "KeyNumpad4"
	case KeyNumpad5:
		return "KeyNumpad5"
	case KeyNumpad6:
		return "KeyNumpad6"
	case KeyNumpad7:
		return "KeyNumpad7"
	case KeyNumpad8:
		return "KeyNumpad8"
	case KeyNumpad9:
		return "KeyNumpad9"
	case KeyNumpadAdd:
		return "KeyNumpadAdd"
	case KeyNumpadDecimal:
		return "KeyNumpadDecimal"
	case KeyNumpadDivide:
		return "KeyNumpadDivide"
	case KeyNumpadEnter:
		return "KeyNumpadEnter"
	case KeyNumpadEqual:
		return "KeyNumpadEqual"
	case KeyNumpadMultiply:
		return "KeyNumpadMultiply"
	case KeyNumpadSubtract:
		return "KeyNumpadSubtract"
	case KeyPageDown:
		return "KeyPageDown"
	case KeyPageUp:
		return "KeyPageUp"
	case KeyPause:
		return "KeyPause"
	case KeyPeriod:
		return "KeyPeriod"
	case KeyPrintScreen:
		return "KeyPrintScreen"
	case KeyQuote:
		return "KeyQuote"
	case KeyScrollLock:
		return "KeyScrollLock"
	case KeySemicolon:
		return "KeySemicolon"
	case KeyShiftLeft:
		return "KeyShiftLeft"
	case KeyShiftRight:
		return "KeyShiftRight"
	case KeySlash:
		return "KeySlash"
	case KeySpace:
		return "KeySpace"
	case KeyTab:
		return "KeyTab"
	}
	return fmt.Sprintf("Key(%d)", k)
}
