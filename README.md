# LOGGING 
SINGLETON

in fact we not need make object logger manually with singleton

recommended to make object logger manually so not use logger singleton. bettter use logrus.New()

logger singleton by default use TextFormatter & JSONFormatter