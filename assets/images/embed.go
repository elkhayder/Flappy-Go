package image_assets

import _ "embed"

var (
	//go:embed icon.png
	GameIcon_png []byte

	/**
	 * Digits
	 */

	//go:embed digits/0.png
	Digit0_png []byte

	//go:embed digits/1.png
	Digit1_png []byte

	//go:embed digits/2.png
	Digit2_png []byte

	//go:embed digits/3.png
	Digit3_png []byte

	//go:embed digits/4.png
	Digit4_png []byte

	//go:embed digits/5.png
	Digit5_png []byte

	//go:embed digits/6.png
	Digit6_png []byte

	//go:embed digits/7.png
	Digit7_png []byte

	//go:embed digits/8.png
	Digit8_png []byte

	//go:embed digits/9.png
	Digit9_png []byte

	/**
	 * UI
	 */

	//go:embed ui/Checkbox_Checked.png
	Checkbox_Checked []byte
	//go:embed ui/Checkbox_Unchecked.png
	Checkbox_Unchecked []byte
)
