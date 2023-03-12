package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var energySum, experienceSum, result int
	energySum, experienceSum = initialEnergy, initialExperience
	for i := 0; i < len(energy); i++ {
		if energySum <= energy[i] {
			result += energy[i] - energySum + 1
			energySum = energy[i] + 1
		}
		if experienceSum <= experience[i] {
			result += experience[i] - experienceSum + 1
			experienceSum = experience[i] + 1
		}
		energySum -= energy[i]
		experienceSum += experience[i]
	}
	return result
}
