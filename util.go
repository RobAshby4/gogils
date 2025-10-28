package main

func BPrint(outputBuffer *string, output string) {
	*outputBuffer = *outputBuffer + output
}

func BPrintln(outputBuffer *string, output string) {
	*outputBuffer = *outputBuffer + output + "\n"
}

func SliceContainsPtr(slice []*Item, target *Item) bool {
	for _, p := range slice {
		if p == target {
			return true
		}
	}
	return false
}
