package main

const (
	length  = 10
	sLength = length - 1

	nucleotideA uint32 = 0
	nucleotideC uint32 = 1 << 18
	nucleotideG uint32 = 2 << 18
	nucleotideT uint32 = 3 << 18
)

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= length {
		return nil
	}

	var (
		result    = make([]string, 0, 128)
		sequences = make(map[uint32]int)
		sequence  = initSequence(s)
		count     int
	)

	for i := sLength; i < len(s); i++ {
		sequence = insertNucleotide(sequence, s[i])
		if count = sequences[sequence]; count == 1 {
			result = append(result, s[i-sLength:i+1])
		}

		sequences[sequence] = count + 1
	}

	return result
}

func initSequence(s string) (result uint32) {
	for i := 0; i < sLength; i++ {
		result = insertNucleotide(result, s[i])
	}

	return result
}

func insertNucleotide(sequence uint32, nucleotide byte) uint32 {
	if nucleotide == 'A' {
		return (sequence >> 2) | nucleotideA
	}

	if nucleotide == 'C' {
		return (sequence >> 2) | nucleotideC
	}

	if nucleotide == 'G' {
		return (sequence >> 2) | nucleotideG
	}

	return (sequence >> 2) | nucleotideT
}
