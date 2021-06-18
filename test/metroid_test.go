package metroid_test

import (
	metroid2 "game-password-codecs/metroid"
	"testing"
)

var (
	metroid = metroid2.Metroid{}
)

func TestBufferRotateLeft(t *testing.T) {
	var rotationBits uint8 = 3
	data := []uint8{0b00100000, 0, 0, 1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, rotationBits, 0}

	b, _ := metroid2.NewBuffer(data)
	b.RotateLeft()

	actual := b.GetBitsMapString()
	expected := `00000000 00000000 00000000 00001000 00000000 00000000 00000000 00000000 00010000 00000000 00000000 00000000 00000000 00000000 00000000 00000001 00000011 00000000
                                                              GAME STATE                                                                          SHIFT  CHECKSUM
[_____________________________________________________________________________________________________________________________________________] [______] [______]`
	if actual != expected {
		t.Errorf("invalid left rotation:\n%s\nvs.\n%s", actual, expected)
	}
}

func TestBufferRotateRight(t *testing.T) {
	var rotationBits uint8 = 3
	data := []uint8{0, 0, 0, 1, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0b00000010, rotationBits, 0}

	b, _ := metroid2.NewBuffer(data)
	b.RotateRight()

	actual := b.GetBitsMapString()
	expected := `01000000 00000000 00000000 00000000 00100000 00000000 00000000 00000000 00000000 01000000 00000000 00000000 00000000 00000000 00000000 00000000 00000011 00000000
                                                              GAME STATE                                                                          SHIFT  CHECKSUM
[_____________________________________________________________________________________________________________________________________________] [______] [______]`
	if actual != expected {
		t.Errorf("invalid left rotation:\n%s\nvs.\n%s", actual, expected)
	}
}

func TestBufferGetBytesString(t *testing.T) {
	data := []uint8{0xFF, 0xAB, 0x71, 0x99, 0x20, 0x33, 0x00, 0x8, 0x58, 0xCD, 0x9B, 0x17, 0x61, 0x94, 0x4, 0x1, 0x3, 0x99}
	b, _ := metroid2.NewBuffer(data)

	actual := b.GetBytesString()
	expected := `[HEX]	[DEC]	[BIN]    
ff	255	11111111    
ab	171	10101011    
71	113	01110001    
99	153	10011001    
20	032	00100000    
33	051	00110011    
00	000	00000000    
08	008	00001000    
58	088	01011000    
cd	205	11001101    
9b	155	10011011    
17	023	00010111    
61	097	01100001    
94	148	10010100    
04	004	00000100    
01	001	00000001    
03	003	00000011    
99	153	10011001`
	if actual != expected {
		t.Errorf("invalid bytes string:\n%s\nvs.\n%s", actual, expected)
	}
}

func TestGetTranslatedPasswordBytes(t *testing.T) {
	s1 := "aAbBcCdDeEfFgGhHiIjJkKlL"
	actual1, err1 := metroid.GetTranslatedPasswordBytes(s1)
	if err1 != nil {
		t.Errorf("invalid password bytes: %v", err1)
	}

	expected1 := []uint8{
		36, 10, 37, 11, 38, 12,
		39, 13, 40, 14, 41, 15,
		42, 16, 43, 17, 44, 18,
		45, 19, 46, 20, 47, 21,
	}
	for i := range actual1 {
		if actual1[i] != expected1[i] {
			t.Errorf("invalid bytes string:\n%d\nvs.\n%d", actual1[i], expected1[i])
		}
	}

	s2 := "                        "
	actual2, err2 := metroid.GetTranslatedPasswordBytes(s2)
	if err2 != nil {
		t.Errorf("invalid password bytes: %v", err2)
	}

	expected2 := []uint8{
		255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255,
		255, 255, 255, 255, 255, 255,
	}
	for i := range actual2 {
		if actual2[i] != expected2[i] {
			t.Errorf("invalid bytes string:\n%d\nvs.\n%d", actual2[i], expected2[i])
		}
	}

	s3 := "!                       "
	_, err3 := metroid.GetTranslatedPasswordBytes(s3)
	if err3 == nil {
		t.Error("invalid password bytes; expected error")
	}
}

func TestJustinBailey(t *testing.T) {
	p := "JUSTINBAILEY------------"
	state, decodeErr := metroid.Decode(p)
	if decodeErr != nil {
		t.Errorf("invalid decoding: %v\n", decodeErr)
	}

	swimsuit1 := state.Armor.Swimsuit
	if !swimsuit1 {
		t.Errorf("invalid Armor: swimsuit=%t\n", swimsuit1)
	}

	password, encodeErr := metroid.Encode(*state)
	if encodeErr != nil || password != p {
		t.Errorf("invalid encoding: password=%s, err=%v\n", password, encodeErr)
	}
}

func TestNarpassword(t *testing.T) {
	p := "NARPASSWORD0000000000000"
	state, decodeErr := metroid.Decode(p)
	if decodeErr != nil {
		t.Errorf("invalid decoding: %v\n", decodeErr)
	}

	password, encodeErr := metroid.Encode(*state)
	if encodeErr != nil || password != p {
		t.Errorf("invalid encoding: password=%s, err=%v\n", password, encodeErr)
	}
}

func TestEngageRidley(t *testing.T) {
	p := "ENGAGERIDLEYMUFFINKICKER"
	state, decodeErr := metroid.Decode(p)
	if decodeErr != nil {
		t.Errorf("invalid decoding: %v\n", decodeErr)
	}

	if !state.Reset {
		t.Errorf("invalid: reset: %t\n", state.Reset)
	}

	password, encodeErr := metroid.Encode(*state)
	if encodeErr != nil || password != p {
		t.Errorf("invalid encoding: password=%s, err=%v\n", password, encodeErr)
	}
}
