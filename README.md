# DRGDoubleXP
Program that prints out upcoming Double XP missions and the potential Experience gained from them

# How to Use
With Go and VS Code installed, clone the repo and open the folder in VS Code. Then run the application and it'll print a list of missions for today that are:
- Double XP
- Have a total XP ( (Mission Base XP * Hazard Bonus) + Predicted Resource XP (4000) ) over a threshold of 10000
- Mission List is ordered by upcoming date, but can be changed to sort by Experience from Highest to Lowest.

  Example Output:
  
  ![image](https://github.com/user-attachments/assets/58550b3b-d530-46af-afdf-315fb4ec57c9)

# Adjustable Parameters
- Forecasting Period set by `searchDays` (Default is 1, ideally don't forecast for too long)
- Mission Hazard Level set by `baseHazardBonus` (Default is Hazard 3, has consts for Hazard 1-5)
- Predicted Resource XP set by `predicted` (Default is 4000, this value represents how much XP you might get from resource gathering and other non Objective related XP)
- Missions List Order set by `sortKeys()` (Default is `ORDER_BY_DATE`, can be changed to `ORDER_BY_XP`)

# How does it work?
The Deep Rock Galactic Live Mission Tracker at doublexp.net has an API for mission forecasting.

The program simply makes an API call for the forecasted day, unmarshals it into structs, then parses the data using data maps defined in `values.go`

I was considering putting in Event chance rolls but the data from the drg wiki.gg is missing data for the Data Cell event and its also too much work to calculate potential
experience gains that aren't guaranteed. I instead opted for a "Theoretical Experience" breakdown that assumes a random roll of 2 from the 1000 XP resource pool, and 1 from the 1500 XP resource pool.

This theoretical value is extremely unlikely though, and the total XP you'll get will fall somewhere between the Mission Experience and the Mission Experience + Predicted + 1 Event Roll.

Note that any Mission with Lithophage Outbreak will not roll the Meteor Impact/Corruptor Event, and any Mission with Industrial Sabotage will NOT roll a Prospector Drone/Data Deposit/Rival Communications Event.
