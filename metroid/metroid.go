package metroid

import (
	"errors"
	"fmt"
	"strings"
)

type (
	Flag uint
)

const (
	// BYTE 0
	BIT_0_MORPH_BALL_TAKEN Flag = iota
	BIT_1_MISSILE_CONTAINER_BRINSTAR
	BIT_2_RED_DOOR_LONG_BEAM
	BIT_3_RED_DOOR_TOURIAN_BRIDGE
	BIT_4_ENERGY_TANK_BRINSTAR
	BIT_5_RED_DOOR_BOMBS
	BIT_6_BOMBS_TAKEN
	BIT_7_RED_DOOR_ICE_BEAM_BRINSTAR

	// BYTE 1
	BIT_8_MISSILE_CONTAINER_BRINSTAR
	BIT_9_ENERGY_TANK_BRINSTAR
	BIT_10_RED_DOOR_VARIA
	BIT_11_VARIA_TAKEN
	BIT_12_ENERGY_TANK_BRINSTAR
	BIT_13_MISSILE_CONTAINER_NORFAIR
	BIT_14_MISSILE_CONTAINER_NORFAIR
	BIT_15_RED_DOOR_ICE_BEAM_NORFAIR

	// BYTE 2
	BIT_16_MISSILE_CONTAINER_NORFAIR
	BIT_17_MISSILE_CONTAINER_NORFAIR
	BIT_18_MISSILE_CONTAINER_NORFAIR
	BIT_19_MISSILE_CONTAINER_NORFAIR
	BIT_20_MISSILE_CONTAINER_NORFAIR
	BIT_21_MISSILE_CONTAINER_NORFAIR
	BIT_22_MISSILE_CONTAINER_NORFAIR
	BIT_23_RED_DOOR_HIGH_JUMP_BOOTS

	// BYTE 3
	BIT_24_HIGH_JUMP_BOOTS_TAKEN
	BIT_25_RED_DOOR_SCREW_ATTACK
	BIT_26_SCREW_ATTACK_TAKEN
	BIT_27_MISSILE_CONTAINER_NORFAIR
	BIT_28_MISSILE_CONTAINER_NORFAIR
	BIT_29_RED_DOOR_WAVE_BEAM
	BIT_30_ENERGY_TANK_NORFAIR
	BIT_31_MISSILE_CONTAINER_NORFAIR

	// BYTE 4
	BIT_32_RED_DOOR_KRAIDS_LAIR
	BIT_33_MISSILE_CONTAINER_KRAIDS_LAIR
	BIT_34_MISSILE_CONTAINER_KRAIDS_LAIR
	BIT_35_RED_DOOR_KRAIDS_LAIR
	BIT_36_ENERGY_TANK_KRAIDS_LAIR
	BIT_37_RED_DOOR_KRAIDS_LAIR
	BIT_38_RED_DOOR_KRAIDS_LAIR
	BIT_39_MISSILE_CONTAINER_KRAIDS_LAIR

	// BYTE 5
	BIT_40_MISSILE_CONTAINER_KRAIDS_LAIR
	BIT_41_RED_DOOR_KRAIDS_ROOM
	BIT_42_ENERGY_TANK_KRAIDS_ROOM
	BIT_43_MISSILE_CONTAINER_RIDLEYS_LAIR
	BIT_44_RED_DOOR_RIDLEYS_LAIR
	BIT_45_ENERGY_TANK_RIDLEYS_LAIR
	BIT_46_MISSILE_CONTAINER_RIDLEYS_LAIR
	BIT_47_YELLOW_DOOR_RIDLEYS_ROOM

	// BYTE 6
	BIT_48_ENERGY_TANK_ROOM_BEHIND_RIDLEY
	BIT_49_MISSILE_CONTAINER_RIDLEYS_LAIR
	BIT_50_YELLOW_DOOR_TOURIAN
	BIT_51_RED_DOOR_TOURIAN
	BIT_52_RED_DOOR_TOURIAN
	BIT_53_ZEBETITE_1_KILLED
	BIT_54_ZEBETITE_2_KILLED
	BIT_55_ZEBETITE_3_KILLED

	// BYTE 7
	BIT_56_ZEBETITE_4_KILLED
	BIT_57_ZEBETITE_5_KILLED
	BIT_58_MOTHER_BRAIN_KILLED
	BIT_59_UNKNOWN
	BIT_60_UNKNOWN
	BIT_61_UNKNOWN
	BIT_62_UNKNOWN
	BIT_63_UNKNOWN

	// BYTE 8
	BIT_64_START_IN_NORFAIR
	BIT_65_START_IN_KRAIDS_LAIR
	BIT_66_START_IN_RIDLEYS_LAIR
	BIT_67_RESET
	BIT_68_UNKNOWN
	BIT_69_UNKNOWN
	BIT_70_UNKNOWN
	BIT_71_SWIMSUIT

	// BYTE 9
	BIT_72_HAS_BOMBS
	BIT_73_HAS_HIGH_JUMP_BOOTS
	BIT_74_HAS_LONG_BEAM
	BIT_75_HAS_SCREW_ATTACK
	BIT_76_HAS_MORPH_BALL
	BIT_77_HAS_VARIA
	BIT_78_HAS_WAVE_BEAM
	BIT_79_HAS_ICE_BEAM

	// BYTE 10
	BIT_80_MISSILES_COUNT
	BIT_81_MISSILES_COUNT
	BIT_82_MISSILES_COUNT
	BIT_83_MISSILES_COUNT
	BIT_84_MISSILES_COUNT
	BIT_85_MISSILES_COUNT
	BIT_86_MISSILES_COUNT
	BIT_87_MISSILES_COUNT

	// BYTE 11
	BIT_88_GAME_AGE
	BIT_89_GAME_AGE
	BIT_90_GAME_AGE
	BIT_91_GAME_AGE
	BIT_92_GAME_AGE
	BIT_93_GAME_AGE
	BIT_94_GAME_AGE
	BIT_95_GAME_AGE

	// BYTE 12
	BIT_96_GAME_AGE
	BIT_97_GAME_AGE
	BIT_98_GAME_AGE
	BIT_99_GAME_AGE
	BIT_100_GAME_AGE
	BIT_101_GAME_AGE
	BIT_102_GAME_AGE
	BIT_103_GAME_AGE

	// BYTE 13
	BIT_104_GAME_AGE
	BIT_105_GAME_AGE
	BIT_106_GAME_AGE
	BIT_107_GAME_AGE
	BIT_108_GAME_AGE
	BIT_109_GAME_AGE
	BIT_110_GAME_AGE
	BIT_111_GAME_AGE

	// BYTE 14
	BIT_112_GAME_AGE
	BIT_113_GAME_AGE
	BIT_114_GAME_AGE
	BIT_115_GAME_AGE
	BIT_116_GAME_AGE
	BIT_117_GAME_AGE
	BIT_118_GAME_AGE
	BIT_119_GAME_AGE

	// BYTE 15
	BIT_120_UNKNOWN
	BIT_121_UNKNOWN
	BIT_122_UNKNOWN
	BIT_123_UNKNOWN
	BIT_124_RIDLEY_KILLED
	BIT_125_RIDLEY_STATUE_RAISED
	BIT_126_KRAID_KILLED
	BIT_127_KRAID_STATUE_RAISED
)

const (
	missilesByte = 10
	gameAgeByte  = 11
	shiftByte    = 16
	checksumByte = 17

	VALID_CHARACTERS  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz?-"
	PASSWORD_OVERRIDE = "NARPASSWORD00000"

	NTSC_HERTZ uint32 = 60
	PAL_HERTZ  uint32 = 50
)

type (
	Metroid struct {
	}
)

func (m Metroid) GetTranslatedPasswordBytes(password string) ([]uint8, error) {
	if len(password) != 24 {
		return nil, errors.New("invalid password length")
	}

	var passwordBytes []uint8
	for _, rune := range password {
		if rune == ' ' {
			passwordBytes = append(passwordBytes, 0xff)
		} else {
			passwordByte := strings.Index(VALID_CHARACTERS, string(rune))
			if passwordByte == -1 {
				return nil, fmt.Errorf("invalid character: %v", rune)
			}

			passwordBytes = append(passwordBytes, uint8(passwordByte))
		}
	}

	return passwordBytes, nil
}

func (m Metroid) Get6BitPasswordBytes(bytes []uint8) []uint8 {
	var newBytes []uint8

	for i := 0; i < 6; i++ {
		startIndex := i * 4
		newVal1 := (bytes[startIndex] << 2) | (bytes[startIndex+1] >> 4)
		newVal2 := (bytes[startIndex+1] << 4) | (bytes[startIndex+2] >> 2)
		newVal3 := (bytes[startIndex+2] << 6) | (bytes[startIndex+3])

		newBytes = append(newBytes, newVal1)
		newBytes = append(newBytes, newVal2)
		newBytes = append(newBytes, newVal3)
	}

	return newBytes
}

func (m Metroid) Get8BitPasswordBytes(bytes []uint8) []uint8 {
	var newBytes []uint8

	for i := 0; i < 6; i++ {
		startIndex := i * 3
		newVal1 := bytes[startIndex] >> 2
		newVal2 := ((bytes[startIndex] & 0x3) << 4) | (bytes[startIndex+1] >> 4)
		newVal3 := ((bytes[startIndex+1] & 0xf) << 2) | (bytes[startIndex+2] >> 6)
		newVal4 := bytes[startIndex+2] & 0x3f

		newBytes = append(newBytes, newVal1)
		newBytes = append(newBytes, newVal2)
		newBytes = append(newBytes, newVal3)
		newBytes = append(newBytes, newVal4)
	}

	return newBytes
}

func (m Metroid) GetOriginalPassword(bytes []uint8) (string, error) {
	var chars []string

	for i, byte := range bytes {
		if byte == 0xff {
			chars = append(chars, " ")
		} else if byte > 63 {
			return "", fmt.Errorf("byte %d invalid: %d not Metroid alphabet", i, byte)
		} else {
			chars = append(chars, string(VALID_CHARACTERS[byte]))
		}
	}

	return strings.Join(chars, ""), nil
}

func (m Metroid) Decode(password string) (*GameState, error) {
	bytes1, err := m.GetTranslatedPasswordBytes(password)
	if err != nil {
		return nil, err
	}

	bytes2 := m.Get6BitPasswordBytes(bytes1)
	buf, bufErr := NewBuffer(bytes2)
	if bufErr != nil {
		return nil, bufErr
	}

	buf.RotateLeft()

	if strings.HasPrefix(password, PASSWORD_OVERRIDE) {
		return NewGameState(*buf, true), nil
	}

	if !buf.ValidateChecksum() {
		return nil, fmt.Errorf("decode failed: invalid checksum:\n%s\n%s", buf.GetBitsMapString(), buf.GetBytesString())
	}

	return NewGameState(*buf, false), nil
}

func (m Metroid) Encode(state GameState) (string, error) {
	buf := state.buffer
	if !state.DebugEnabled && !buf.ValidateChecksum() {
		return "", errors.New("encode failed: invalid checksum")
	}

	buf.RotateRight()
	rotatedBytes := buf.Data()
	bytes := m.Get8BitPasswordBytes(rotatedBytes)
	return m.GetOriginalPassword(bytes)
}

func PrettyPassword(password string) string {
	return fmt.Sprintf("%s %s %s %s", password[0:6], password[6:12], password[12:18], password[18:24])
}
