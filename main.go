package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"sort"
	"time"

	"github.com/dariubs/percent"
)

var (
	comparators = ComparatorsMap{
		ORDER_BY_DATE: DateComparator,
		ORDER_BY_XP:   ExperienceComparator,
	}
	baseHazardBonus     = HAZARD_3
	doubleMissions      = UniqueMission{}
	experienceThreshold = 5000
	offsetTime          = time.Now().Add(-time.Minute * 30)
	currentTime         = time.Now()
	// How much resource XP obtained from a mission, 2 xp per non-Objective
	predicted = 1000
)

const (
	ORDER_BY_DATE = 1
	ORDER_BY_XP   = 2
)

func main() {
	searchDays := 2
	for i := 0; i < searchDays; i++ {
		missions := getMissionsFor(currentTime.AddDate(0, 0, i))
		for _, data := range missions {
			for _, biome := range data.Biomes {
				if offsetTime.Before(data.TimeStamp) {
					searchMissionData(data.TimeStamp, biome)
				}
			}
		}
	}
	printMissionData()
}

func getMissionsFor(requestedDate time.Time) DataBlock {
	res, err := http.Get(fmt.Sprintf("https://doublexp.net/static/json/bulkmissions/%s.json", requestedDate.Format(time.DateOnly)))
	if err != nil {
		fmt.Println("Its broken")
	}
	defer res.Body.Close()
	jsonBody, _ := io.ReadAll(res.Body)
	var missions DataBlock
	json.Unmarshal(jsonBody, &missions)
	return missions
}

func printMissionData() {
	// Create a new map that is a sorted version of doubleMissions
	keys := sortKeys(doubleMissions, ORDER_BY_XP)

	for _, key := range keys {
		mission := doubleMissions[key]
		fmt.Printf("%s (starts in %s)\n", mission.Timestamp.Format(time.DateTime), mission.Timestamp.Sub(currentTime).Truncate(time.Second))
		fmt.Println(mission.Mission.toString())
		fmt.Println("-------")
	}
}

func sortKeys(missions UniqueMission, compare int) []int {
	keys := make([]int, 0, len(missions))
	for key := range missions {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return comparators[compare](missions, keys, i, j)
	})
	return keys
}

func searchMissionData(timestamp time.Time, missions []DRGMission) {
	for i := 0; i < len(missions); i++ {
		mission := missions[i]
		if mission.hasDoubleXpMutator() && (mission.getExperience()+predicted)*2 >= experienceThreshold {
			applicableMission := TimeStampMission{timestamp, mission}
			doubleMissions[mission.Seed] = applicableMission
		}
	}
}

type ComparatorsMap map[int]func(UniqueMission, []int, int, int) bool

func ExperienceComparator(missions UniqueMission, keys []int, i int, j int) bool {
	return missions[keys[i]].Mission.getExperience() > missions[keys[j]].Mission.getExperience()
}

func DateComparator(missions UniqueMission, keys []int, i int, j int) bool {
	return missions[keys[i]].Timestamp.Before(missions[keys[j]].Timestamp)
}

type UniqueMission map[int]TimeStampMission

type TimeStampMission struct {
	Timestamp time.Time
	Mission   DRGMission
}

type DataBlock map[time.Time]RolloverData

type RolloverData struct {
	TimeStamp time.Time   `json:"timestamp"`
	Biomes    ActiveBiome `json:"Biomes"`
}

type ActiveBiome map[string][]DRGMission

type DRGMission struct {
	Seed               int      `json:"Seed"`
	PrimaryObjective   string   `json:"PrimaryObjective"`
	SecondaryObjective string   `json:"SecondaryObjective"`
	MissionWarnings    []string `json:"MissionWarnings"`
	MissionMutator     string   `json:"MissionMutator"`
	Complexity         string   `json:"Complexity"`
	Length             string   `json:"Length"`
	CodeName           string   `json:"CodeName"`
	IncludedIn         []string `json:"included_in"`
	ID                 int      `json:"id"`
}

func (mission DRGMission) toString() string {
	text := fmt.Sprintf("Objectives: %s, %s\nCave: Length %s Complexity %s", mission.PrimaryObjective, mission.SecondaryObjective, mission.Length, mission.Complexity)
	if len(mission.MissionWarnings) > 0 {
		text += fmt.Sprintf("\nWarnings: %v", mission.MissionWarnings)
	}
	hazard := baseHazardBonus + mission.getHazardBonus()
	xp := mission.getExperience()
	text += fmt.Sprintf("\nHazard Bonus: %d%%", hazard)
	text += fmt.Sprintf("\nMission Experience: %d (%d)", xp*2, xp)
	text += fmt.Sprintf("\nMission Experience + Predicted Resources: %d (%d)", (xp+predicted)*2, xp+predicted)
	text += mission.getTheoreticalExperience()

	return text
}

func (mission DRGMission) getHazardBonus() int {
	hazard := caveHazardBonus[[2]string{mission.Length, mission.Complexity}]
	for i := 0; i < len(mission.MissionWarnings); i++ {
		hazard += warningBonus[mission.MissionWarnings[i]]
	}
	return hazard
}

func (mission DRGMission) getSecondaryType() int {
	if slices.Contains(typeOneSecondary, mission.SecondaryObjective) {
		return 1
	} else {
		return 2
	}
}

func (mission DRGMission) getTheoreticalExperience() string {
	experience := mission.getExperience() + predicted
	hazard := mission.getHazardBonus()

	bonusOne := fmt.Sprintf("Any 2 of: Machine Event, Core Stone, Meteor Impact, Data Cell: %d", int(percent.Percent(hazard+100, 2000)))
	bonusTwo := fmt.Sprintf("Tyrant Shard: %d", int(percent.Percent(hazard+100, 1500)))

	xp := experience + int(percent.Percent(hazard+100, 3500))
	theoreticalTotal := fmt.Sprintf("Total Theoretical Experience: %d (%d)", xp*2, xp)
	return fmt.Sprintf("\n\n%s\n%s\n\n%s", bonusOne, bonusTwo, theoreticalTotal)
}

func (mission DRGMission) getExperience() int {
	hazard := baseHazardBonus + mission.getHazardBonus()
	secondary := mission.getSecondaryType()
	experience := experienceTable[MissionParameters{mission.Length, mission.Complexity, mission.PrimaryObjective, secondary}]
	return int(percent.Percent(hazard+100, experience))
}

func (mission DRGMission) hasDoubleXpMutator() bool {
	return mission.MissionMutator == "Double XP"
}
