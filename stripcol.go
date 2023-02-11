package stripcol

func StripColor(text string) string {
	strippedText := make([]byte, 0, len(text))

	// Iterate through the text character by character
	for i := 0; i < len(text); i++ {
		if text[i] != 27 {
			strippedText = append(strippedText, text[i])
		} else {
			// Check if the current character is an escape character
			if i+1 < len(text) && text[i+1] == '[' {
				i += 2
				// Check if the following characters are digits
				for i < len(text) && text[i] >= '0' && text[i] <= '9' {
					i++
				}
				// Check if the next character is ';'
				for i < len(text) && text[i] == ';' {
					i++
					// Check if the following characters are digits
					for i < len(text) && text[i] >= '0' && text[i] <= '9' {
						i++
					}
				}
				// Check if the next character is 'm'
				if i < len(text) && text[i] == 'm' {
					continue
				}
				// If the next character is not 'm', add the escape character to the strippedText
				strippedText = append(strippedText, text[i-1])
			} else {
				strippedText = append(strippedText, text[i])
			}
		}
	}

	return string(strippedText)
}
