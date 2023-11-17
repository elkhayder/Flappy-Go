package image_assets

import _ "embed"

var (
	//go:embed favicon.ico
	Favicon_ico []byte

	/**
	 * UI
	 */

	//go:embed ui/Checkbox_Checked.png
	Checkbox_Checked []byte
	//go:embed ui/Checkbox_Unchecked.png
	Checkbox_Unchecked []byte
)
