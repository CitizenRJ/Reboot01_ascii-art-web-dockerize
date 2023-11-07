package asciiartfs

import "strings"

// Print the full outcome
func PrintBanners(banners, arr []string) string {
	var lines []string

	for i, ch := range banners {
		if i != 0 && ch != "" && banners[i-1] != "" {
			lines = append(lines, "")
		}
		if ch == "" {
			continue
		}

		for j := 0; j < 8; j++ {
			var line strings.Builder
			for _, char := range ch {
				n := (int(char)-32)*9 + j
				if n >= 0 && n < len(arr) {
					line.WriteString(arr[n])
				}
			}
			lines = append(lines, line.String())
		}
	}
	return strings.Join(lines, "\n")
}
