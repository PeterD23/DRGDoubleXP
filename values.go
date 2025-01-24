package main

type MissionParameters struct {
	Length     string
	Complexity string
	Primary    string
	Secondary  int
}

type CaveHazardBonus map[[2]string]int
type WarningBonus map[string]int
type Experience map[MissionParameters]int

const (
	HAZARD_1 = 25
	HAZARD_2 = 50
	HAZARD_3 = 75
	HAZARD_4 = 100
	HAZARD_5 = 133
	mine     = "Mining Expedition"
	egg      = "Egg Hunt"
	salvage  = "Salvage Operation"
	point    = "Point Extraction"
	elim     = "Elimination"
	escort   = "Escort Duty"
	refine   = "On-Site Refining"
	sabotage = "Industrial Sabotage"
	scan     = "Deep Scan"
)

var (
	typeOneSecondary = []string{"ApocaBlooms", "Boolo Caps", "Bha Barnacles", "Fossils", "Fester Fleas", "Glyphid Eggs", "Ebonuts", "Gunk Seeds"}
	experienceTable  = Experience{
		{"1", "1", mine, 1}: 2400, {"1", "1", mine, 2}: 2385,
		{"2", "1", mine, 1}: 2645, {"2", "1", mine, 2}: 2625,
		{"2", "2", mine, 1}: 2890, {"2", "2", mine, 2}: 2855,
		{"3", "2", mine, 1}: 3845, {"3", "2", mine, 2}: 3790,
		{"3", "3", mine, 1}: 4800, {"3", "3", mine, 2}: 4735,

		{"1", "1", egg, 1}: 1650, {"1", "1", egg, 2}: 1635,
		{"2", "2", egg, 1}: 2475, {"2", "2", egg, 2}: 2435,
		{"3", "2", egg, 1}: 3300, {"3", "2", egg, 2}: 3235,

		{"2", "2", salvage, 1}: 2800, {"2", "2", salvage, 2}: 2765,
		{"3", "3", salvage, 1}: 3600, {"3", "3", salvage, 2}: 3545,

		{"2", "3", point, 1}: 1715, {"2", "3", point, 2}: 1715,
		{"3", "3", point, 1}: 2450, {"3", "3", point, 2}: 2500,

		{"2", "2", elim, 1}: 2940, {"2", "2", elim, 2}: 2915,
		{"3", "3", elim, 1}: 4165, {"3", "3", elim, 2}: 4115,

		{"2", "2", escort, 1}: 3899, {"2", "2", escort, 2}: 3969,
		{"2", "3", escort, 1}: 4200, {"2", "3", escort, 2}: 4165,
		{"3", "2", escort, 1}: 4800, {"3", "2", escort, 2}: 4755,
		{"3", "3", escort, 1}: 5400, {"3", "3", escort, 2}: 5345,

		{"2", "2", refine, 1}: 2680, {"2", "2", refine, 2}: 2640,
		{"2", "3", refine, 1}: 3520, {"2", "3", refine, 2}: 3475,

		{"2", "1", sabotage, 1}: 3500, {"2", "1", sabotage, 2}: 3485,
		{"2", "2", sabotage, 1}: 5250, {"2", "2", sabotage, 2}: 5210,

		{"1", "2", scan, 1}: 2900, {"1", "2", scan, 2}: 2885,
		{"2", "3", scan, 1}: 4843, {"2", "3", scan, 2}: 4817,
	}
	caveHazardBonus = CaveHazardBonus{
		{"1", "1"}: 0,
		{"2", "1"}: 10,
		{"1", "2"}: 10,
		{"2", "2"}: 20,
		{"3", "2"}: 30,
		{"2", "3"}: 30,
		{"3", "3"}: 40,
	}
	warningBonus = WarningBonus{
		"Cave Leech Cluster":   15,
		"Parasites":            15,
		"Regenerative Bugs":    15,
		"Exploder Infestation": 20,
		"Low Oxygen":           20,
		"Mactera Plague":       20,
		"Swarmageddon":         20,
		"Ebonite Outbreak":     20,
		"Lethal Enemies":       25,
		"Elite Threat":         30,
		"Haunted Cave":         30,
		"Rival Presence":       30,
		"Shield Disruption":    30,
		"Duck and Cover":       30,
		"Lithophage Outbreak":  50,
	}
)
