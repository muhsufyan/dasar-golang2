# LOGGING 
FORMATTER

logrus by default have 2 formatter, TextFormatter (default), JSONFormatter (format message log in JSON format, this option 2). but we can make new formatter if we want (just implement Formatter interface)

for change formatter (ex from text to json) use SetFormatter()

previous (logging-3-output) file application.log use TextFormatter

