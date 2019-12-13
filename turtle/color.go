package turtle

// from: https://github.com/lucasb-eyer/go-colorful/blob/master/doc/gradientgen/gradientgen.go

import "github.com/lucasb-eyer/go-colorful"

func mustParseHex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}

// GradientTable contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,256]
type GradientTable []struct {
	Col colorful.Color
	Pos float64
}

// GetInterpolatedColorFor returns a HCL-blend between the two colors around `t`.
// Note: It relies heavily on the fact that the gradient keypoints are sorted.
func (gt GradientTable) GetInterpolatedColorFor(t float64) colorful.Color {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, t).Clamped()
		}
	}

	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return gt[len(gt)-1].Col
}

// NewPalette returns the gradient color palette
func NewPalette() GradientTable {
	// The "keypoints" of the gradient.
	keypoints := GradientTable{
		{mustParseHex("#cc9900"), 0.0},
		{mustParseHex("#97ff01"), 32},
		{mustParseHex("#67ffff"), 64},
		{mustParseHex("#ed70f8"), 96},
		{mustParseHex("#ed70f8"), 128},
		{mustParseHex("#f6ce75"), 160},
		{mustParseHex("#f6ce75"), 192},
		{mustParseHex("#f6ce75"), 224},
		{mustParseHex("#cc9900"), 256},
	}
	return keypoints
}
