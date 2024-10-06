package data

var StabilityTestAddrList []string = []string{
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
}

var StabilityTestIPList []string = []string{
	"8.8.8.8",
	"9.9.9.9",
	"8.8.4.4",
}

var StabilityGrade []string = []string{"A", "B", "C", "D", "F", "G", "H", "I", "J", "K"}

var StabilityGradeDescription map[string]string = map[string]string{
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