package util

var (
  /// encoded file header and actual body separator
	HeaderBodySeparator   string = "\n~<=>-\n"
  /// encoded file header (k-v) pair separator
	KeyValuePairSeparator string = "%->"
  /// string between a key and a value 
	KeyValueSeparator     string = "-<>~"
)
