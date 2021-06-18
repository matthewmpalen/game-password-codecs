package metroid

import (
	"encoding/json"
)

type (
	Bosses struct {
		RidleyKilled      bool `json:"ridley_killed"`       // 124
		KraidKilled       bool `json:"kraid_killed"`        // 126
		Zebetite1Killed   bool `json:"zebetite1_killed"`    // 53
		Zebetite2Killed   bool `json:"zebetite2_killed"`    // 54
		Zebetite3Killed   bool `json:"zebetite3_killed"`    // 55
		Zebetite4Killed   bool `json:"zebetite4_killed"`    // 56
		Zebetite5Killed   bool `json:"zebetite5_killed"`    // 57
		MotherBrainKilled bool `json:"mother_brain_killed"` // 58
	}

	Statues struct {
		RidleyStatueRaised bool `json:"ridley_statue_raised"` // 125
		KraidStatueRaised  bool `json:"kraid_statue_raised"`  // 127
	}

	Doors struct {
		LongBeam        bool `json:"long_beam"`         // 2
		TourianBridge   bool `json:"tourian_bridge"`    // 3
		Bombs           bool `json:"bombs"`             // 5
		BrinstarIceBeam bool `json:"brinstar_ice_beam"` // 7
		Varia           bool `json:"varia"`             // 10
		NorfairIceBeam  bool `json:"norfair_ice_beam"`  // 15
		HighJumpBoots   bool `json:"high_jump_boots"`   // 23
		ScrewAttack     bool `json:"screw_attack"`      // 25
		WaveBeam        bool `json:"wave_beam"`         // 29
		KraidsLair1     bool `json:"kraids_lair1"`      // 32
		KraidsLair2     bool `json:"kraids_lair2"`      // 35
		KraidsLair3     bool `json:"kraids_lair3"`      // 37
		KraidsLair4     bool `json:"kraids_lair4"`      // 38
		KraidsRoom      bool `json:"kraids_room"`       // 41
		RidleysLair     bool `json:"ridleys_lair"`      // 44
		RidleysRoom     bool `json:"ridleys_room"`      // 47
		Tourian1        bool `json:"tourian1"`          // 50
		Tourian2        bool `json:"tourian2"`          // 51
		Tourian3        bool `json:"tourian3"`          // 52
	}

	MissileContainers struct {
		Brinstar1    bool `json:"brinstar1"`     //1
		Brinstar2    bool `json:"brinstar2"`     //8
		Norfair1     bool `json:"norfair1"`      //13
		Norfair2     bool `json:"norfair2"`      //14
		Norfair3     bool `json:"norfair3"`      //16
		Norfair4     bool `json:"norfair4"`      //17
		Norfair5     bool `json:"norfair5"`      //18
		Norfair6     bool `json:"norfair6"`      //19
		Norfair7     bool `json:"norfair7"`      //20
		Norfair8     bool `json:"norfair8"`      //21
		Norfair9     bool `json:"norfair9"`      //22
		Norfair10    bool `json:"norfair10"`     //27
		Norfair11    bool `json:"norfair11"`     //28
		Norfair12    bool `json:"norfair12"`     //31
		KraidsLair1  bool `json:"kraids_lair1"`  //33
		KraidsLair2  bool `json:"kraids_lair2"`  //34
		KraidsLair3  bool `json:"kraids_lair3"`  //39
		KraidsLair4  bool `json:"kraids_lair4"`  //40
		RidleysLair1 bool `json:"ridleys_lair1"` //43
		RidleysLair2 bool `json:"ridleys_lair2"` //46
		RidleysLair3 bool `json:"ridleys_lair3"` //49
	}

	EnergyTanks struct {
		Brinstar1    bool `json:"brinstar1"`     //4
		Brinstar2    bool `json:"brinstar2"`     //9
		Brinstar3    bool `json:"brinstar3"`     //12
		Norfair      bool `json:"norfair"`       //30
		KraidsLair   bool `json:"kraids_lair"`   //36
		KraidsRoom   bool `json:"kraids_room"`   //42
		RidleysLair1 bool `json:"ridleys_lair1"` //45
		RidleysLair2 bool `json:"ridleys_lair2"` //48
	}

	Clock struct {
		Hours   uint32 `json:"hours"`
		Minutes uint32 `json:"minutes"`
		Seconds uint32 `json:"seconds"`
	}

	GameAge struct {
		// 88 - 119 (4 bytes)
		Ticks uint32 `json:"ticks"`
		NTSC  Clock  `json:"ntsc"`
		PAL   Clock  `json:"pal"`
	}

	StartLocation struct {
		// 64 - 66
		Name string `json:"name"`
	}

	Powerups struct {
		Bombs         bool `json:"bombs"`           //72
		HighJumpBoots bool `json:"high_jump_boots"` //73
		LongBeam      bool `json:"long_beam"`       //74
		ScrewAttack   bool `json:"screw_attack"`    //75
		MorphBall     bool `json:"morph_ball"`      //76
		Varia         bool `json:"varia"`           //77
		WaveBeam      bool `json:"wave_beam"`       //78
		IceBeam       bool `json:"ice_beam"`        //79

		// Removed from map
		MorphBallTaken     bool `json:"morph_ball_taken"`      //0
		BombsTaken         bool `json:"bombs_taken"`           //6
		VariaTaken         bool `json:"varia_taken"`           //11
		HighJumpBootsTaken bool `json:"high_jump_boots_taken"` //24
		ScrewAttackTaken   bool `json:"screw_attack_taken"`    //26
	}

	Missiles struct {
		// 80-87
		Count uint8 `json:"count"`
	}

	Armor struct {
		Swimsuit bool `json:"swimsuit"` //71
	}

	GameState struct {
		buffer            Buffer
		Bosses            Bosses            `json:"bosses"`
		Statues           Statues           `json:"statues"`
		Doors             Doors             `json:"doors"`
		MissileContainers MissileContainers `json:"missile_containers"`
		EnergyTanks       EnergyTanks       `json:"energy_tanks"`
		GameAge           GameAge           `json:"game_age"`
		StartLocation     StartLocation     `json:"start_location"`
		Powerups          Powerups          `json:"powerups"`
		Missiles          Missiles          `json:"missiles"`
		Armor             Armor             `json:"armor"`
		Reset             bool              `json:"reset"`
		DebugEnabled      bool              `json:"debug_enabled"`
	}
)

func NewBosses(buf Buffer) Bosses {
	return Bosses{
		RidleyKilled:      buf.IsFlagSet(BIT_124_RIDLEY_KILLED),
		KraidKilled:       buf.IsFlagSet(BIT_126_KRAID_KILLED),
		Zebetite1Killed:   buf.IsFlagSet(BIT_53_ZEBETITE_1_KILLED),
		Zebetite2Killed:   buf.IsFlagSet(BIT_54_ZEBETITE_2_KILLED),
		Zebetite3Killed:   buf.IsFlagSet(BIT_55_ZEBETITE_3_KILLED),
		Zebetite4Killed:   buf.IsFlagSet(BIT_56_ZEBETITE_4_KILLED),
		Zebetite5Killed:   buf.IsFlagSet(BIT_57_ZEBETITE_5_KILLED),
		MotherBrainKilled: buf.IsFlagSet(BIT_58_MOTHER_BRAIN_KILLED),
	}
}

func NewStatues(buf Buffer) Statues {
	return Statues{
		RidleyStatueRaised: buf.IsFlagSet(BIT_125_RIDLEY_STATUE_RAISED),
		KraidStatueRaised:  buf.IsFlagSet(BIT_127_KRAID_STATUE_RAISED),
	}
}

func NewDoors(buf Buffer) Doors {
	return Doors{
		LongBeam:        buf.IsFlagSet(BIT_2_RED_DOOR_LONG_BEAM),
		TourianBridge:   buf.IsFlagSet(BIT_3_RED_DOOR_TOURIAN_BRIDGE),
		Bombs:           buf.IsFlagSet(BIT_5_RED_DOOR_BOMBS),
		BrinstarIceBeam: buf.IsFlagSet(BIT_7_RED_DOOR_ICE_BEAM_BRINSTAR),
		Varia:           buf.IsFlagSet(BIT_10_RED_DOOR_VARIA),
		NorfairIceBeam:  buf.IsFlagSet(BIT_15_RED_DOOR_ICE_BEAM_NORFAIR),
		HighJumpBoots:   buf.IsFlagSet(BIT_23_RED_DOOR_HIGH_JUMP_BOOTS),
		ScrewAttack:     buf.IsFlagSet(BIT_25_RED_DOOR_SCREW_ATTACK),
		WaveBeam:        buf.IsFlagSet(BIT_29_RED_DOOR_WAVE_BEAM),
		KraidsLair1:     buf.IsFlagSet(BIT_32_RED_DOOR_KRAIDS_LAIR),
		KraidsLair2:     buf.IsFlagSet(BIT_35_RED_DOOR_KRAIDS_LAIR),
		KraidsLair3:     buf.IsFlagSet(BIT_37_RED_DOOR_KRAIDS_LAIR),
		KraidsLair4:     buf.IsFlagSet(BIT_38_RED_DOOR_KRAIDS_LAIR),
		KraidsRoom:      buf.IsFlagSet(BIT_41_RED_DOOR_KRAIDS_ROOM),
		RidleysLair:     buf.IsFlagSet(BIT_44_RED_DOOR_RIDLEYS_LAIR),
		RidleysRoom:     buf.IsFlagSet(BIT_47_YELLOW_DOOR_RIDLEYS_ROOM),
		Tourian1:        buf.IsFlagSet(BIT_50_YELLOW_DOOR_TOURIAN),
		Tourian2:        buf.IsFlagSet(BIT_51_RED_DOOR_TOURIAN),
		Tourian3:        buf.IsFlagSet(BIT_52_RED_DOOR_TOURIAN),
	}
}

func NewMissileContainers(buf Buffer) MissileContainers {
	return MissileContainers{
		Brinstar1:    buf.IsFlagSet(BIT_1_MISSILE_CONTAINER_BRINSTAR),
		Brinstar2:    buf.IsFlagSet(BIT_8_MISSILE_CONTAINER_BRINSTAR),
		Norfair1:     buf.IsFlagSet(BIT_13_MISSILE_CONTAINER_NORFAIR),
		Norfair2:     buf.IsFlagSet(BIT_14_MISSILE_CONTAINER_NORFAIR),
		Norfair3:     buf.IsFlagSet(BIT_16_MISSILE_CONTAINER_NORFAIR),
		Norfair4:     buf.IsFlagSet(BIT_17_MISSILE_CONTAINER_NORFAIR),
		Norfair5:     buf.IsFlagSet(BIT_18_MISSILE_CONTAINER_NORFAIR),
		Norfair6:     buf.IsFlagSet(BIT_19_MISSILE_CONTAINER_NORFAIR),
		Norfair7:     buf.IsFlagSet(BIT_20_MISSILE_CONTAINER_NORFAIR),
		Norfair8:     buf.IsFlagSet(BIT_21_MISSILE_CONTAINER_NORFAIR),
		Norfair9:     buf.IsFlagSet(BIT_22_MISSILE_CONTAINER_NORFAIR),
		Norfair10:    buf.IsFlagSet(BIT_27_MISSILE_CONTAINER_NORFAIR),
		Norfair11:    buf.IsFlagSet(BIT_28_MISSILE_CONTAINER_NORFAIR),
		Norfair12:    buf.IsFlagSet(BIT_31_MISSILE_CONTAINER_NORFAIR),
		KraidsLair1:  buf.IsFlagSet(BIT_33_MISSILE_CONTAINER_KRAIDS_LAIR),
		KraidsLair2:  buf.IsFlagSet(BIT_34_MISSILE_CONTAINER_KRAIDS_LAIR),
		KraidsLair3:  buf.IsFlagSet(BIT_39_MISSILE_CONTAINER_KRAIDS_LAIR),
		KraidsLair4:  buf.IsFlagSet(BIT_40_MISSILE_CONTAINER_KRAIDS_LAIR),
		RidleysLair1: buf.IsFlagSet(BIT_43_MISSILE_CONTAINER_RIDLEYS_LAIR),
		RidleysLair2: buf.IsFlagSet(BIT_46_MISSILE_CONTAINER_RIDLEYS_LAIR),
		RidleysLair3: buf.IsFlagSet(BIT_49_MISSILE_CONTAINER_RIDLEYS_LAIR),
	}
}

func NewEnergyTanks(buf Buffer) EnergyTanks {
	return EnergyTanks{
		Brinstar1:    buf.IsFlagSet(BIT_4_ENERGY_TANK_BRINSTAR),
		Brinstar2:    buf.IsFlagSet(BIT_9_ENERGY_TANK_BRINSTAR),
		Brinstar3:    buf.IsFlagSet(BIT_12_ENERGY_TANK_BRINSTAR),
		Norfair:      buf.IsFlagSet(BIT_30_ENERGY_TANK_NORFAIR),
		KraidsLair:   buf.IsFlagSet(BIT_36_ENERGY_TANK_KRAIDS_LAIR),
		KraidsRoom:   buf.IsFlagSet(BIT_42_ENERGY_TANK_KRAIDS_ROOM),
		RidleysLair1: buf.IsFlagSet(BIT_45_ENERGY_TANK_RIDLEYS_LAIR),
		RidleysLair2: buf.IsFlagSet(BIT_48_ENERGY_TANK_ROOM_BEHIND_RIDLEY),
	}
}

func NewClock(ticks uint32, hertz uint32) Clock {
	totalSeconds := float64(ticks) * (256 / float64(hertz))
	intSeconds := int(totalSeconds)

	hours := intSeconds / 3600
	mins := (intSeconds - (hours * 3600)) / 60
	secs := intSeconds - (hours * 3600) - (mins * 60)

	return Clock{
		Hours:   uint32(hours),
		Minutes: uint32(mins),
		Seconds: uint32(secs),
	}
}

func NewGameAge(buf Buffer) GameAge {
	bytes := buf.GetBytes(gameAgeByte, gameAgeByte+4)
	var ticks uint32 = 0
	for i, b1 := range bytes {
		b2 := uint32(b1) << (i * 8)
		ticks += b2
	}

	return GameAge{
		Ticks: ticks,
		NTSC:  NewClock(ticks, NTSC_HERTZ),
		PAL:   NewClock(ticks, PAL_HERTZ),
	}
}

func NewStartLocation(buf Buffer) StartLocation {
	norfairBit := buf.IsFlagSet(BIT_64_START_IN_NORFAIR)
	kraidsLairBit := buf.IsFlagSet(BIT_65_START_IN_KRAIDS_LAIR)
	ridleysLairBit := buf.IsFlagSet(BIT_66_START_IN_RIDLEYS_LAIR)

	var name string
	if !norfairBit && !kraidsLairBit && !ridleysLairBit {
		name = "Brinstar"
	} else if norfairBit && !kraidsLairBit && !ridleysLairBit {
		name = "Norfair"
	} else if !norfairBit && kraidsLairBit && !ridleysLairBit {
		name = "Kraid's Lair"
	} else if !norfairBit && !kraidsLairBit && ridleysLairBit {
		name = "Ridley's Lair"
	} else if norfairBit && kraidsLairBit {
		name = "Tourian"
	} else {
		name = "invalid"
	}

	return StartLocation{
		Name: name,
	}
}

func NewPowerups(buf Buffer) Powerups {
	return Powerups{
		Bombs:              buf.IsFlagSet(BIT_72_HAS_BOMBS),
		HighJumpBoots:      buf.IsFlagSet(BIT_73_HAS_HIGH_JUMP_BOOTS),
		LongBeam:           buf.IsFlagSet(BIT_74_HAS_LONG_BEAM),
		ScrewAttack:        buf.IsFlagSet(BIT_75_HAS_SCREW_ATTACK),
		MorphBall:          buf.IsFlagSet(BIT_76_HAS_MORPH_BALL),
		Varia:              buf.IsFlagSet(BIT_77_HAS_VARIA),
		WaveBeam:           buf.IsFlagSet(BIT_78_HAS_WAVE_BEAM),
		IceBeam:            buf.IsFlagSet(BIT_79_HAS_ICE_BEAM),
		MorphBallTaken:     buf.IsFlagSet(BIT_0_MORPH_BALL_TAKEN),
		BombsTaken:         buf.IsFlagSet(BIT_6_BOMBS_TAKEN),
		VariaTaken:         buf.IsFlagSet(BIT_11_VARIA_TAKEN),
		HighJumpBootsTaken: buf.IsFlagSet(BIT_24_HIGH_JUMP_BOOTS_TAKEN),
		ScrewAttackTaken:   buf.IsFlagSet(BIT_26_SCREW_ATTACK_TAKEN),
	}
}

func NewMissiles(buf Buffer) Missiles {
	return Missiles{
		Count: buf.Byte(missilesByte),
	}
}

func NewArmor(buf Buffer) Armor {
	return Armor{
		Swimsuit: buf.IsFlagSet(BIT_71_SWIMSUIT),
	}
}

func NewGameState(buf Buffer, debugEnabled bool) *GameState {
	return &GameState{
		buffer:            buf,
		Bosses:            NewBosses(buf),
		Statues:           NewStatues(buf),
		Doors:             NewDoors(buf),
		MissileContainers: NewMissileContainers(buf),
		EnergyTanks:       NewEnergyTanks(buf),
		GameAge:           NewGameAge(buf),
		StartLocation:     NewStartLocation(buf),
		Powerups:          NewPowerups(buf),
		Missiles:          NewMissiles(buf),
		Armor:             NewArmor(buf),
		Reset:             buf.IsFlagSet(BIT_67_RESET),
		DebugEnabled:      debugEnabled,
	}
}

func NewGameStateFromJSON(jsonString string) (*GameState, error) {
	var state GameState
	err := json.Unmarshal([]byte(jsonString), &state)
	if err != nil {
		return nil, err
	}

	state.BuildBuffer()

	return &state, nil
}

func (g GameState) Buffer() Buffer {
	return g.buffer
}

func (g *GameState) BuildBuffer() {
	var data []uint8
	for i := 0; i < 18; i++ {
		data = append(data, 0)
	}

	buf, _ := NewBuffer(data)
	buf.SetBossesFromState(g.Bosses)
	buf.SetStatuesFromState(g.Statues)
	buf.SetDoorsFromState(g.Doors)
	buf.SetMissileContainersFromState(g.MissileContainers)
	buf.SetEnergyTanksFromState(g.EnergyTanks)

	buf.SetGameAgeFromState(g.GameAge)
	buf.SetStartLocationFromState(g.StartLocation)
	buf.SetPowerupsFromState(g.Powerups)
	buf.SetMissilesFromState(g.Missiles)
	buf.SetArmorFromState(g.Armor)
	buf.SetResetFromState(g.Reset)
	buf.BuildChecksum()

	g.buffer = *buf
}

func (g GameState) ToJSON(pretty bool) string {
	var bytes []byte
	if pretty {
		bytes, _ = json.MarshalIndent(g, "", "    ")
	} else {
		bytes, _ = json.Marshal(g)
	}

	return string(bytes)
}
