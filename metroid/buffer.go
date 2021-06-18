package metroid

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type (
	Buffer struct {
		data []uint8
	}
)

func NewBuffer(data []uint8) (*Buffer, error) {
	if len(data) != 18 {
		return nil, errors.New("invalid data")
	}

	return &Buffer{
		data: data,
	}, nil
}

func (b Buffer) Data() []uint8 {
	return b.data
}

func (b *Buffer) GetBitsString() string {
	var s1 []string
	for _, b := range b.data {
		s1 = append(s1, fmt.Sprintf("%08b", b))
	}

	return strings.Join(s1, " ")
}

func (b *Buffer) GetBitsMapString() string {
	s1 := b.GetBitsString()
	var s2 = `
                                                              GAME STATE                                                                          SHIFT  CHECKSUM
[_____________________________________________________________________________________________________________________________________________] [______] [______]`

	return s1 + s2
}

func (b Buffer) GetBytesString() string {
	tokens := []string{"[HEX]\t[DEC]\t[BIN]"}
	for _, b := range b.data {
		s := fmt.Sprintf("\n%02x\t%03d\t%08b", b, b, b)
		tokens = append(tokens, s)
	}

	return strings.Join(tokens, "    ")
}

func (b Buffer) ShiftByte() uint8 {
	return b.data[shiftByte]
}

func (b Buffer) GetBit(bitIndex uint8) uint8 {
	byteIndex := bitIndex / 8
	bitPos := uint8(1) << (bitIndex % 8)
	return b.data[byteIndex] & bitPos
}

func (b *Buffer) SetFlag(flag Flag, enabled bool) {
	var val uint8 = 0
	if enabled {
		val = 1
	}

	bitIndex := uint8(flag)
	byteIndex := bitIndex / 8
	bitPos := val << (bitIndex % 8)
	b.data[byteIndex] |= bitPos
}

func (b Buffer) IsFlagSet(flag Flag) bool {
	return b.GetBit(uint8(flag)) > 0
}

func (b Buffer) Byte(i uint8) uint8 {
	return b.data[i]
}

func (b *Buffer) SetByte(i, val uint8) {
	b.data[i] = val
}

func (b Buffer) GetBytes(start, end uint8) []uint8 {
	return b.data[start:end]
}

func (b *Buffer) RotateLeft() {
	var carry uint8 = 1
	var carryTemp uint8 = 0
	rotateAmount := b.data[shiftByte]

	var i uint8 = 0
	for i = 0; i < rotateAmount; i++ {
		temp := b.data[15]

		for j := 15; j >= 0; j-- {
			carryTemp = (b.data[j] & 0x80) >> 7
			b.data[j] = ((b.data[j] << 1) & 0xff) | (carry & 0x1)
			carry = carryTemp
		}

		carryTemp = (temp & 0x80) >> 7
		temp = ((temp << 1) & 0xff) | (carry & 0x1)
		carry = carryTemp

		b.data[15] = temp
	}
}

func (b *Buffer) RotateRight() {
	var carry uint8 = 1
	var carryTemp uint8 = 0
	rotateAmount := b.data[shiftByte]

	var i uint8 = 0
	for i = 0; i < rotateAmount; i++ {
		temp := b.data[0]

		for j := 0; j < 16; j++ {
			carryTemp = b.data[j] & 0x1
			b.data[j] = (b.data[j] >> 1) | ((carry & 0x1) << 7)
			carry = carryTemp
		}

		carryTemp = temp & 0x1
		temp = (temp >> 1) | ((carry & 0x1) << 7)
		carry = carryTemp
		b.data[0] = temp
	}
}

func (b Buffer) CalculateChecksum() uint8 {
	var actual uint8 = 0
	for i := 0; i < 17; i++ {
		actual += b.data[i]
	}

	return actual
}

func (b *Buffer) BuildChecksum() {
	b.data[checksumByte] = b.CalculateChecksum()
}

func (b Buffer) ValidateChecksum() bool {
	return b.CalculateChecksum() == b.data[checksumByte]
}

func (b *Buffer) SetBossesFromState(bos Bosses) {
	b.SetFlag(BIT_124_RIDLEY_KILLED, bos.RidleyKilled)
	b.SetFlag(BIT_126_KRAID_KILLED, bos.KraidKilled)
	b.SetFlag(BIT_53_ZEBETITE_1_KILLED, bos.Zebetite1Killed)
	b.SetFlag(BIT_54_ZEBETITE_2_KILLED, bos.Zebetite2Killed)
	b.SetFlag(BIT_55_ZEBETITE_3_KILLED, bos.Zebetite3Killed)
	b.SetFlag(BIT_56_ZEBETITE_4_KILLED, bos.Zebetite4Killed)
	b.SetFlag(BIT_57_ZEBETITE_5_KILLED, bos.Zebetite5Killed)
	b.SetFlag(BIT_58_MOTHER_BRAIN_KILLED, bos.MotherBrainKilled)
}

func (b *Buffer) SetStatuesFromState(s Statues) {
	b.SetFlag(BIT_125_RIDLEY_STATUE_RAISED, s.RidleyStatueRaised)
	b.SetFlag(BIT_127_KRAID_STATUE_RAISED, s.KraidStatueRaised)
}

func (b *Buffer) SetDoorsFromState(d Doors) {
	b.SetFlag(BIT_2_RED_DOOR_LONG_BEAM, d.LongBeam)
	b.SetFlag(BIT_3_RED_DOOR_TOURIAN_BRIDGE, d.TourianBridge)
	b.SetFlag(BIT_58_MOTHER_BRAIN_KILLED, d.Bombs)
	b.SetFlag(BIT_7_RED_DOOR_ICE_BEAM_BRINSTAR, d.BrinstarIceBeam)
	b.SetFlag(BIT_10_RED_DOOR_VARIA, d.Varia)
	b.SetFlag(BIT_15_RED_DOOR_ICE_BEAM_NORFAIR, d.NorfairIceBeam)
	b.SetFlag(BIT_23_RED_DOOR_HIGH_JUMP_BOOTS, d.HighJumpBoots)
	b.SetFlag(BIT_25_RED_DOOR_SCREW_ATTACK, d.ScrewAttack)
	b.SetFlag(BIT_29_RED_DOOR_WAVE_BEAM, d.WaveBeam)
	b.SetFlag(BIT_32_RED_DOOR_KRAIDS_LAIR, d.KraidsLair1)
	b.SetFlag(BIT_35_RED_DOOR_KRAIDS_LAIR, d.KraidsLair2)
	b.SetFlag(BIT_37_RED_DOOR_KRAIDS_LAIR, d.KraidsLair3)
	b.SetFlag(BIT_38_RED_DOOR_KRAIDS_LAIR, d.KraidsLair4)
	b.SetFlag(BIT_41_RED_DOOR_KRAIDS_ROOM, d.KraidsRoom)
	b.SetFlag(BIT_44_RED_DOOR_RIDLEYS_LAIR, d.RidleysLair)
	b.SetFlag(BIT_47_YELLOW_DOOR_RIDLEYS_ROOM, d.RidleysRoom)
	b.SetFlag(BIT_50_YELLOW_DOOR_TOURIAN, d.Tourian1)
	b.SetFlag(BIT_51_RED_DOOR_TOURIAN, d.Tourian2)
	b.SetFlag(BIT_52_RED_DOOR_TOURIAN, d.Tourian3)
}

func (b *Buffer) SetMissileContainersFromState(c MissileContainers) {
	b.SetFlag(BIT_1_MISSILE_CONTAINER_BRINSTAR, c.Brinstar1)
	b.SetFlag(BIT_8_MISSILE_CONTAINER_BRINSTAR, c.Brinstar2)
	b.SetFlag(BIT_13_MISSILE_CONTAINER_NORFAIR, c.Norfair1)
	b.SetFlag(BIT_14_MISSILE_CONTAINER_NORFAIR, c.Norfair2)
	b.SetFlag(BIT_16_MISSILE_CONTAINER_NORFAIR, c.Norfair3)
	b.SetFlag(BIT_17_MISSILE_CONTAINER_NORFAIR, c.Norfair4)
	b.SetFlag(BIT_18_MISSILE_CONTAINER_NORFAIR, c.Norfair5)
	b.SetFlag(BIT_19_MISSILE_CONTAINER_NORFAIR, c.Norfair6)
	b.SetFlag(BIT_20_MISSILE_CONTAINER_NORFAIR, c.Norfair7)
	b.SetFlag(BIT_21_MISSILE_CONTAINER_NORFAIR, c.Norfair8)
	b.SetFlag(BIT_22_MISSILE_CONTAINER_NORFAIR, c.Norfair9)
	b.SetFlag(BIT_27_MISSILE_CONTAINER_NORFAIR, c.Norfair10)
	b.SetFlag(BIT_28_MISSILE_CONTAINER_NORFAIR, c.Norfair11)
	b.SetFlag(BIT_31_MISSILE_CONTAINER_NORFAIR, c.Norfair12)
	b.SetFlag(BIT_33_MISSILE_CONTAINER_KRAIDS_LAIR, c.KraidsLair1)
	b.SetFlag(BIT_34_MISSILE_CONTAINER_KRAIDS_LAIR, c.KraidsLair2)
	b.SetFlag(BIT_39_MISSILE_CONTAINER_KRAIDS_LAIR, c.KraidsLair3)
	b.SetFlag(BIT_40_MISSILE_CONTAINER_KRAIDS_LAIR, c.KraidsLair4)
	b.SetFlag(BIT_43_MISSILE_CONTAINER_RIDLEYS_LAIR, c.RidleysLair1)
	b.SetFlag(BIT_46_MISSILE_CONTAINER_RIDLEYS_LAIR, c.RidleysLair2)
	b.SetFlag(BIT_49_MISSILE_CONTAINER_RIDLEYS_LAIR, c.RidleysLair3)
}

func (b *Buffer) SetEnergyTanksFromState(t EnergyTanks) {
	b.SetFlag(BIT_4_ENERGY_TANK_BRINSTAR, t.Brinstar1)
	b.SetFlag(BIT_9_ENERGY_TANK_BRINSTAR, t.Brinstar2)
	b.SetFlag(BIT_12_ENERGY_TANK_BRINSTAR, t.Brinstar3)
	b.SetFlag(BIT_30_ENERGY_TANK_NORFAIR, t.Norfair)
	b.SetFlag(BIT_36_ENERGY_TANK_KRAIDS_LAIR, t.KraidsLair)
	b.SetFlag(BIT_42_ENERGY_TANK_KRAIDS_ROOM, t.KraidsRoom)
	b.SetFlag(BIT_45_ENERGY_TANK_RIDLEYS_LAIR, t.RidleysLair1)
	b.SetFlag(BIT_48_ENERGY_TANK_ROOM_BEHIND_RIDLEY, t.RidleysLair2)
}

func (b *Buffer) SetTicks(ticks uint32) {
	for i := 0; i < 4; i++ {
		byte := uint8(ticks & 0xff)
		b.data[gameAgeByte+i] = byte
		ticks >>= 8
	}
}

func (b *Buffer) SetGameAgeFromState(a GameAge) {
	if a.Ticks > 0 {
		b.SetTicks(a.Ticks)
	} else if a.NTSC.Hours > 0 || a.NTSC.Minutes > 0 || a.NTSC.Seconds > 0 {
		totalSeconds := (a.NTSC.Hours * 3600) + (a.NTSC.Minutes * 60) + a.NTSC.Seconds
		fTicks := float64(totalSeconds) * (float64(NTSC_HERTZ) / 256)
		ticks := uint32(math.Round(fTicks))
		b.SetTicks(ticks)
	} else if a.PAL.Hours > 0 || a.PAL.Minutes > 0 || a.PAL.Seconds > 0 {
		totalSeconds := (a.PAL.Hours * 3600) + (a.PAL.Minutes * 60) + a.PAL.Seconds
		fTicks := float64(totalSeconds) * (float64(PAL_HERTZ) / 256)
		ticks := uint32(math.Round(fTicks))
		b.SetTicks(ticks)
	}
}

func (b *Buffer) SetStartLocationFromState(l StartLocation) {
	switch l.Name {
	case "Brinstar":
		b.SetFlag(BIT_64_START_IN_NORFAIR, false)
		b.SetFlag(BIT_65_START_IN_KRAIDS_LAIR, false)
		b.SetFlag(BIT_66_START_IN_RIDLEYS_LAIR, false)
	case "Norfair":
		b.SetFlag(BIT_64_START_IN_NORFAIR, true)
		b.SetFlag(BIT_65_START_IN_KRAIDS_LAIR, false)
		b.SetFlag(BIT_66_START_IN_RIDLEYS_LAIR, false)
	case "Kraid's Lair":
		b.SetFlag(BIT_64_START_IN_NORFAIR, false)
		b.SetFlag(BIT_65_START_IN_KRAIDS_LAIR, true)
		b.SetFlag(BIT_66_START_IN_RIDLEYS_LAIR, false)
	case "Ridley's Lair":
		b.SetFlag(BIT_64_START_IN_NORFAIR, false)
		b.SetFlag(BIT_65_START_IN_KRAIDS_LAIR, false)
		b.SetFlag(BIT_66_START_IN_RIDLEYS_LAIR, true)
	case "Tourian":
		b.SetFlag(BIT_64_START_IN_NORFAIR, true)
		b.SetFlag(BIT_65_START_IN_KRAIDS_LAIR, true)
		b.SetFlag(BIT_66_START_IN_RIDLEYS_LAIR, false)
	default:
		b.SetFlag(BIT_64_START_IN_NORFAIR, true)
		b.SetFlag(BIT_65_START_IN_KRAIDS_LAIR, true)
		b.SetFlag(BIT_66_START_IN_RIDLEYS_LAIR, true)
	}
}

func (b *Buffer) SetPowerupsFromState(p Powerups) {
	b.SetFlag(BIT_72_HAS_BOMBS, p.Bombs)
	b.SetFlag(BIT_73_HAS_HIGH_JUMP_BOOTS, p.HighJumpBoots)
	b.SetFlag(BIT_74_HAS_LONG_BEAM, p.LongBeam)
	b.SetFlag(BIT_75_HAS_SCREW_ATTACK, p.ScrewAttack)
	b.SetFlag(BIT_76_HAS_MORPH_BALL, p.MorphBall)
	b.SetFlag(BIT_77_HAS_VARIA, p.Varia)
	b.SetFlag(BIT_78_HAS_WAVE_BEAM, p.WaveBeam)
	b.SetFlag(BIT_79_HAS_ICE_BEAM, p.IceBeam)

	b.SetFlag(BIT_0_MORPH_BALL_TAKEN, p.MorphBallTaken)
	b.SetFlag(BIT_6_BOMBS_TAKEN, p.BombsTaken)
	b.SetFlag(BIT_11_VARIA_TAKEN, p.VariaTaken)
	b.SetFlag(BIT_24_HIGH_JUMP_BOOTS_TAKEN, p.HighJumpBootsTaken)
	b.SetFlag(BIT_26_SCREW_ATTACK_TAKEN, p.ScrewAttackTaken)
}

func (b *Buffer) SetMissilesFromState(m Missiles) {
	b.SetByte(missilesByte, m.Count)
}

func (b *Buffer) SetArmorFromState(armor Armor) {
	b.SetFlag(BIT_71_SWIMSUIT, armor.Swimsuit)
}

func (b *Buffer) SetResetFromState(reset bool) {
	b.SetFlag(BIT_67_RESET, reset)
}
