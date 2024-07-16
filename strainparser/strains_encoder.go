package strainparser

import (
	"math/bits"
)

// precomputed lookup tables for effect index
var effectsLookup = map[string]int{
	"Aroused":   0,
	"Creative":  1,
	"Energetic": 2,
	"Euphoric":  3,
	"Focused":   4,
	"Giggly":    5,
	"Happy":     6,
	"Hungry":    7,
	"Relaxed":   8,
	"Sleepy":    9,
	"Talkative": 10,
	"Tingly":    11,
	"Uplifted":  12,
}
var flavorsLookup = map[string]int{
	"Ammonia":      0,
	"Apple":        1,
	"Appricot":     2,
	"Berry":        3,
	"Blue":         4,
	"Blueberry":    5,
	"Butter":       6,
	"Cheese":       7,
	"Chemical":     8,
	"Chestnut":     9,
	"Citrus":       10,
	"Coffee":       11,
	"Diesel":       12,
	"Earthy":       13,
	"Flowery":      14,
	"Fruit":        15,
	"Grape":        16,
	"Grapefruit":   17,
	"Honey":        18,
	"Lavender":     19,
	"Lemon":        20,
	"Lime":         21,
	"Mango":        22,
	"Menthol":      23,
	"Mint":         24,
	"Minty":        24,
	"Nutty":        25,
	"Orange":       26,
	"Peach":        27,
	"Pear":         28,
	"Pepper":       29,
	"Pine":         30,
	"Pineapple":    31,
	"Plum":         32,
	"Pungent":      33,
	"Rose":         34,
	"Sage":         35,
	"Skunk":        36,
	"Spicy/Herbal": 37,
	"Strawberry":   38,
	"Sweet":        39,
	"Tar":          40,
	"Tea":          41,
	"Tobacco":      42,
	"Tree":         43,
	"Tropical":     44,
	"Vanilla":      45,
	"Violet":       46,
	"Woody":        47,
}

const NUM_EFFECTS = 13
const NUM_FLAVORS = 48

type StrainBitEncoding struct {
	value uint64
}

func (b *StrainBitEncoding) SetBit(position uint) {
	if position < 64 {
		b.value |= (1 << position)
	}
}

func (b *StrainBitEncoding) ClearBit(position uint) {
	if position < 64 {
		b.value &= ^(1 << position)
	}
}

func (b *StrainBitEncoding) ToggleBit(position uint) {
	if position < 64 {
		b.value ^= (1 << position)
	}
}

func (a *StrainBitEncoding) GetOneSimilarity(b *StrainBitEncoding) int {
	return bits.OnesCount(uint(a.value & b.value))
}

func (a *StrainBitEncoding) GetOneSimilarityIgnoreType(b *StrainBitEncoding) int {
	x, y := &StrainBitEncoding{value: a.value}, &StrainBitEncoding{value: b.value}

	var i uint
	for i = 0; i < 3; i += 1 {
		x.ClearBit(i)
		y.ClearBit(i)
	}

	return bits.OnesCount(uint(x.value & y.value))
}

func (a *StrainBitEncoding) GetEffectsOneSimilarity(b *StrainBitEncoding) int {
	x, y := &StrainBitEncoding{value: a.value}, &StrainBitEncoding{value: b.value}
	var i uint
	for i = 0; i < 3; i++ {
		x.ClearBit(i)
		y.ClearBit(i)
	}

	for i = 17; i < 64; i++ {
		x.ClearBit(i)
		y.ClearBit(i)
	}
	return bits.OnesCount(uint(x.value & y.value))
}

func (a *StrainBitEncoding) GetFlavorsOneSimilarity(b *StrainBitEncoding) int {
	x, y := &StrainBitEncoding{value: a.value}, &StrainBitEncoding{value: b.value}
	var i uint
	for i = 0; i < 17; i++ {
		x.ClearBit(i)
		y.ClearBit(i)
	}
	return bits.OnesCount(uint(x.value & y.value))

}

// encode to vector ->
//
//	       0 ..2	  3  ...		 16 17       ...     64
//		[strain type] [    effects    ] [    flavors      ]
func GetStrainEncodings(s StrainEntry) *StrainBitEncoding {
	encoding := &StrainBitEncoding{value: 0}

	//first three bits indicate strain type
	if s.Type == "sativa" {
		encoding.SetBit(2)
	}
	if s.Type == "hybrid" {
		encoding.SetBit(1)
	}
	if s.Type == "indica" {
		encoding.SetBit(0)
	}

	// use lookup table to determine bit positions
	for _, e := range s.Effects {
		if idx, ok := effectsLookup[e]; ok {
			encoding.SetBit(uint(idx + 3))
		}
	}
	for _, f := range s.Flavors {
		if idx, ok := flavorsLookup[f]; ok {
			encoding.SetBit(uint(idx + 17))
		}
	}
	//encoded into a single uint64
	return encoding
}

func NewQueryEncoding(strain_type string, effects, flavors []string) *StrainBitEncoding {
	encoding := &StrainBitEncoding{value: 0}

	//first three bits indicate strain type
	if strain_type == "sativa" {
		encoding.SetBit(2)
	}
	if strain_type == "hybrid" {
		encoding.SetBit(1)
	}
	if strain_type == "indica" {
		encoding.SetBit(0)
	}

	for _, e := range effects {
		if idx, ok := effectsLookup[e]; ok {
			encoding.SetBit(uint(idx + 3))
		}
	}

	for _, f := range flavors {
		if idx, ok := flavorsLookup[f]; ok {
			encoding.SetBit(uint(idx + 17))
		}
	}
	return encoding
}

func CompareStrainEncodings(a, b uint64) int {
	return int(a & b)
}
