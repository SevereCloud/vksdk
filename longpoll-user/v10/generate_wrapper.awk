{
	# handler type description
	print "// "$1"Handler handler func for "$1
	# handler type definition
	print "type "$1"Handler func(m "$1")\n"

	# setup function description
	print "// On"$1
	print "//"
	print "// event with code "$7
	print "//"
	print "// TODO: write description"
	# setup function defintion
	print "func (w Wrapper) On"$1"(f "$1"Handler) {"
	print "\tw.longpoll.EventNew("$7", func(i []interface{}) error {"
	print "\t\tevent := "$1"{}"
	print "\t\tif err := event.Parse(i); err != nil {"
	print "\t\t\treturn err"
	print "\t\t}"
	print "\t\tf(event)"
	print ""
	print "\t\treturn nil\n" # note: two newlines
	print "\t})"
	print "}"
}
