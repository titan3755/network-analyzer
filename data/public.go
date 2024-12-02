package data

// this is the list of all hosts that are used for the stability test

var StabilityTestAddrList = []string{
	"google.com",
	"facebook.com",
	"twitter.com",
	"linkedin.com",
	"youtube.com",
	"instagram.com",
	"pinterest.com",
	"reddit.com",
	"tumblr.com",
	"vimeo.com",
	"wordpress.com",
	"blogger.com",
	"medium.com",
	"twitch.tv",
	"soundcloud.com",
	"spotify.com",
	"apple.com",
	"microsoft.com",
	"amazon.com",
	"ebay.com",
} // more to be added later

// this is the list of all IP addresses that are used for the stability test

var StabilityTestIPList = []string{
	"8.8.8.8",
	"9.9.9.9",
	"8.8.4.4",
}

// this is the list of all available stability grades

var StabilityGrade = []string{"A", "B", "C", "D", "F", "G", "H", "I", "J", "K"}

// this is the description of the stability grades

var StabilityGradeDescription = map[string]string{
	"A": "Perfect",
	"B": "Excellent",
	"C": "Very Good",
	"D": "Good",
	"F": "Fair",
	"G": "Poor",
	"H": "Very Poor",
	"I": "Bad",
	"J": "Very Bad",
	"K": "Worst",
}

var SettingsFileName = "settings.prp"

var AvailableSettings = []string{
	"ip_file",
	"output_location",
}

// this is the current version of the app

var CurrentAppVersion = PreviousAppVersionsInclLatestVersion[0]

// this is the list of all previous app versions including the latest version

var PreviousAppVersionsInclLatestVersion = []string{
	"0.7.0",
	"0.6.0",
	"0.5.0",
	"0.4.0",
	"0.3.0",
	"0.2.0",
	"0.1.0",
}

// this is the filename for stability test data

var StabilityTestDataFileName = "stability_test_data.data"

// this is the filename for speed test data

var SpeedTestDataFileName = "speed_test_data.data"
