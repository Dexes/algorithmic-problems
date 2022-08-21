package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) (result int) {
	for i := 0; i < len(energy); i++ {
		if need := energy[i] + 1; initialEnergy < need {
			result += need - initialEnergy
			initialEnergy = 1
		} else {
			initialEnergy -= energy[i]
		}

		if need := experience[i] + 1; initialExperience < need {
			result += need - initialExperience
			initialExperience = need
		}

		initialExperience += experience[i]
	}

	return result
}
