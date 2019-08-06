package main

type Site struct {
	name       string
	profileURL string
	checkerFn  checkerFunc
	regexCheck string
	probeURL   string
	noPeriod   string
}

var sites = []Site{
	{
		name: "500px",
		// errorMsg:   "Oops! This page doesn\u2019t exist.",
		// checkBy:    "message",
		// rank:       3144,
		profileURL: "https://500px.com/%s",
		// mainUrl:    "https://500px.com/",
		checkerFn: bodyChecker("Oops! This page doesn\u2019t exist."),
	},
	{
		name: "9GAG",
		// checkBy:    "status_code",
		// rank:       308,
		profileURL: "https://9gag.com/u/%s",
		// mainUrl:    "https://9gag.com/",
		checkerFn: statusChecker,
	},
	{
		name: "About.me",
		// checkBy:    "status_code",
		// rank:       14151,
		profileURL: "https://about.me/%s",
		// mainUrl:    "https://about.me/",
		checkerFn: statusChecker,
	},
	{
		name: "Academia.edu",
		// checkBy:    "status_code",
		// rank:       184,
		profileURL: "https://independent.academia.edu/%s",
		// mainUrl:    "https://www.academia.edu/",
		checkerFn: statusChecker,
	},
	{
		name: "AngelList",
		// checkBy:    "status_code",
		// rank:       3377,
		profileURL: "https://angel.co/%s",
		// mainUrl:    "https://angel.co/",
		checkerFn: statusChecker,
	},
	{
		name: "Aptoide",
		// checkBy:    "status_code",
		// rank:       4933,
		profileURL: "https://%s.en.aptoide.com/",
		// mainUrl:    "https://en.aptoide.com/",
		checkerFn: statusChecker,
	},
	{
		name: "AskFM",
		// checkBy:    "status_code",
		// rank:       1162,
		profileURL: "https://ask.fm/%s",
		// mainUrl:    "https://ask.fm/",
		checkerFn: statusChecker,
	},
	{
		name: "BLIP.fm",
		// checkBy:    "status_code",
		// rank:       366949,
		profileURL: "https://blip.fm/%s",
		// mainUrl:    "https://blip.fm/",
		checkerFn: statusChecker,
	},
	{
		name: "Badoo",
		// checkBy:    "status_code",
		// rank:       1268,
		profileURL: "https://badoo.com/profile/%s",
		// mainUrl:    "https://badoo.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Bandcamp",
		// checkBy:    "status_code",
		// rank:       607,
		profileURL: "https://www.bandcamp.com/%s",
		// mainUrl:    "https://www.bandcamp.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Basecamp",
		//errorMsg:   "The account you were looking for doesn't exist",
		// checkBy:    "message",
		// rank:       1754,
		profileURL: "https://%s.basecamphq.com",
		// mainUrl:    "https://basecamp.com/",
		checkerFn: bodyChecker("The account you were looking for doesn't exist"),
	},
	{
		name: "Behance",
		// checkBy:    "status_code",
		// rank:       384,
		profileURL: "https://www.behance.net/%s",
		// mainUrl:    "https://www.behance.net/",
		checkerFn: statusChecker,
	},
	{
		name: "BitBucket",
		// checkBy:    "status_code",
		// rank:       989,
		profileURL: "https://bitbucket.org/%s/",
		// mainUrl:    "https://bitbucket.org/",
		checkerFn: statusChecker,
	},
	{
		name: "BitCoinForum",
		// errorMsg:   "The user whose profile you are trying to view does not exist.",
		// checkBy:    "message",
		// rank:       507091,
		profileURL: "https://bitcoinforum.com/profile/%s",
		// mainUrl:    "https://bitcoinforum.com",
		checkerFn: bodyChecker("The user whose profile you are trying to view does not exist."),
	},
	{
		name: "Blogger",
		// checkBy:    "status_code",
		// rank:       211,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.blogspot.com",
		// mainUrl:    "https://www.blogger.com/",
		checkerFn: statusChecker,
	},
	{
		name: "BuzzFeed",
		// checkBy:    "status_code",
		// rank:       283,
		profileURL: "https://buzzfeed.com/%s",
		// mainUrl:    "https://buzzfeed.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Canva",
		// checkBy: "status_code",
		// errorUrl:   "https://www.canva.com/%s",
		// rank:       172,
		profileURL: "https://www.canva.com/%s",
		// mainUrl:    "https://www.canva.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Carbonmade",
		// checkBy:    "response_url",
		// errorUrl:   "https://carbonmade.com/fourohfour?domain=%s.carbonmade.com",
		// rank:       31822,
		profileURL: "https://%s.carbonmade.com",
		// mainUrl:    "https://carbonmade.com/",
		checkerFn: redirectChecker("https://carbonmade.com/fourohfour?domain=%s.carbonmade.com"),
	},
	{
		name: "CashMe",
		// checkBy:    "status_code",
		// rank:       12351551,
		profileURL: "https://cash.me/%s",
		// mainUrl:    "https://cash.me/",
		checkerFn: statusChecker,
	},
	{
		name: "Cloob",
		// checkBy:    "status_code",
		// rank:       9254,
		profileURL: "https://www.cloob.com/name/%s",
		// mainUrl:    "https://www.cloob.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Codecademy",
		// checkBy:    "status_code",
		// rank:       2943,
		profileURL: "https://www.codecademy.com/%s",
		// mainUrl:    "https://www.codecademy.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Codechef",
		// checkBy:    "status_code",
		// rank:       9098,
		profileURL: "https://www.codechef.com/users/%s",
		// mainUrl:    "https://www.codechef.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Codementor",
		// checkBy:    "status_code",
		// rank:       9600,
		profileURL: "https://www.codementor.io/%s",
		// mainUrl:    "https://www.codementor.io/",
		checkerFn: statusChecker,
	},
	{
		name: "Codepen",
		// checkBy:    "status_code",
		// rank:       831,
		profileURL: "https://codepen.io/%s",
		// mainUrl:    "https://codepen.io/",
		checkerFn: statusChecker,
	},
	{
		name: "Coderwall",
		//errorMsg:   "404! Our feels when that url is used",
		//checkBy:    "message",
		// rank:       10061,
		profileURL: "https://coderwall.com/%s",
		// mainUrl:    "https://coderwall.com/",
		checkerFn: bodyChecker("404! Our feels when that url is used"),
	},
	{
		name: "ColourLovers",
		// checkBy:    "status_code",
		// rank:       32386,
		profileURL: "https://www.colourlovers.com/lover/%s",
		// mainUrl:    "https://www.colourlovers.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Contently",
		// errorMsg:   "We can't find that page!",
		// checkBy:    "message",
		// rank:       58845,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.contently.com/",
		// mainUrl:    "https://contently.com/",
		checkerFn: bodyChecker("We can't find that page!"),
	},
	{
		name: "Coroflot",
		// checkBy:    "status_code",
		// rank:       40375,
		profileURL: "https://www.coroflot.com/%s",
		// mainUrl:    "https://coroflot.com/",
		checkerFn: statusChecker,
	},
	{
		name: "CreativeMarket",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.creativemarket.com/",
		// rank:       2100,
		profileURL: "https://creativemarket.com/%s",
		// mainUrl:    "https://creativemarket.com/",
		checkerFn: redirectChecker("https://www.creativemarket.com/"),
	},
	{
		name: "Crevado",
		// checkBy:    "status_code",
		// rank:       178744,
		profileURL: "https://%s.crevado.com",
		// mainUrl:    "https://crevado.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Crunchyroll",
		// checkBy:    "status_code",
		// rank:       383,
		profileURL: "https://www.crunchyroll.com/user/%s",
		// mainUrl:    "https://www.crunchyroll.com/",
		checkerFn: statusChecker,
	},
	{
		name: "DEV Community",
		// checkBy:    "status_code",
		// rank:       5203,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://dev.to/%s",
		// mainUrl:    "https://dev.to/",
		checkerFn: statusChecker,
	},
	{
		name: "DailyMotion",
		// checkBy:    "status_code",
		// rank:       120,
		profileURL: "https://www.dailymotion.com/%s",
		// mainUrl:    "https://www.dailymotion.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Designspiration",
		// checkBy:    "status_code",
		// rank:       26101,
		profileURL: "https://www.designspiration.net/%s/",
		// mainUrl:    "https://www.designspiration.net/",
		checkerFn: statusChecker,
	},
	{
		name: "DeviantART",
		// checkBy:    "status_code",
		// rank:       222,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.deviantart.com",
		// mainUrl:    "https://deviantart.com",
		checkerFn: statusChecker,
	},
	{
		name: "Discogs",
		// checkBy:    "status_code",
		// rank:       614,
		profileURL: "https://www.discogs.com/user/%s",
		// mainUrl:    "https://www.discogs.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Disqus",
		// checkBy:    "status_code",
		// rank:       1297,
		profileURL: "https://disqus.com/%s",
		// mainUrl:    "https://disqus.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Docker Hub",
		// checkBy:    "status_code",
		// rank:       1993,
		profileURL: "https://hub.docker.com/u/%s/",
		// mainUrl:    "https://hub.docker.com/",
		probeURL:  "https://hub.docker.com/v2/users/%s/",
		checkerFn: statusChecker,
	},
	{
		name: "Dribbble",
		// errorMsg:   "Whoops, that page is gone.",
		// checkBy:    "message",
		// rank:       954,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://dribbble.com/%s",
		// mainUrl:    "https://dribbble.com/",
		checkerFn: bodyChecker("Whoops, that page is gone."),
	},
	{
		name: "EVE Online",
		// checkBy:    "response_url",
		// errorUrl:   "https://eveonline.com",
		// rank:       11619,
		profileURL: "https://evewho.com/pilot/%s/",
		// mainUrl:    "https://eveonline.com",
		checkerFn: redirectChecker("https://eveonline.com"),
	},
	{
		name: "Ebay",
		// errorMsg:   "The User ID you entered was not found",
		// checkBy:    "message",
		// rank:       45,
		profileURL: "https://www.ebay.com/usr/%s",
		// mainUrl:    "https://www.ebay.com/",
		checkerFn: bodyChecker("The User ID you entered was not found"),
	},
	{
		name: "Ello",
		// errorMsg:   "We couldn't find the page you're looking for",
		// checkBy:    "message",
		// rank:       36947,
		profileURL: "https://ello.co/%s",
		// mainUrl:    "https://ello.co/",
		checkerFn: bodyChecker("We couldn't find the page you're looking for"),
	},
	{
		name: "Etsy",
		// checkBy:    "status_code",
		// rank:       166,
		profileURL: "https://www.etsy.com/shop/%s",
		// mainUrl:    "https://www.etsy.com/",
		checkerFn: statusChecker,
	},
	{
		name: "EyeEm",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.eyeem.com/",
		// rank:       32792,
		profileURL: "https://www.eyeem.com/u/%s",
		// mainUrl:    "https://www.eyeem.com/",
		checkerFn: redirectChecker("https://www.eyeem.com/"),
	},
	{
		name: "Facebook",
		// checkBy: "status_code",
		// rank: 3,
		// TODO: Regex
		regexCheck: "^[a-zA-Z0-9]{4,49}$",
		// regexCheck: "^[a-zA-Z0-9]{4,49}(?<!.com|.org|.net)$",
		profileURL: "https://www.facebook.com/%s",
		// mainUrl:    "https://www.facebook.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Fandom",
		// checkBy:    "status_code",
		// rank:       58,
		profileURL: "https://www.fandom.com/u/%s",
		// mainUrl:    "https://www.fandom.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Flickr",
		// checkBy:    "status_code",
		// rank:       454,
		profileURL: "https://www.flickr.com/people/%s",
		// mainUrl:    "https://www.flickr.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Flipboard",
		// checkBy:    "status_code",
		// rank:       3603,
		regexCheck: "^([a-zA-Z0-9_]){1,15}$",
		profileURL: "https://flipboard.com/@%s",
		// mainUrl:    "https://flipboard.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Foursquare",
		// checkBy:    "status_code",
		// rank:       2105,
		profileURL: "https://foursquare.com/%s",
		// mainUrl:    "https://foursquare.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Giphy",
		// checkBy:    "status_code",
		// rank:       472,
		profileURL: "https://giphy.com/%s",
		// mainUrl:    "https://giphy.com/",
		checkerFn: statusChecker,
	},
	{
		name: "GitHub",
		// checkBy: "status_code",
		// rank: 55,
		// TODO: Regex
		regexCheck: "^[a-zA-Z0-9-]{0,38}$",
		// regexCheck: "^[a-zA-Z0-9](?:[a-zA-Z0-9]|-(?=[a-zA-Z0-9])){0,38}$",
		profileURL: "https://www.github.com/%s",
		// mainUrl:    "https://www.github.com/",
		checkerFn: statusChecker,
	},
	{
		name: "GitLab",
		// errorMsg: "[]",
		// checkBy:    "message",
		// rank:       1874,
		profileURL: "https://gitlab.com/%s",
		// mainUrl:    "https://gitlab.com/",
		probeURL:  "https://gitlab.com/api/v4/users?username=%s",
		checkerFn: bodyChecker("Page Not Found"),
	},
	{
		name: "Gitee",
		// checkBy:    "status_code",
		// rank:       2848,
		profileURL: "https://gitee.com/%s",
		// mainUrl:    "https://gitee.com/",
		checkerFn: statusChecker,
	},
	{
		name: "GoodReads",
		// checkBy:    "status_code",
		// rank:       404,
		profileURL: "https://www.goodreads.com/%s",
		// mainUrl:    "https://www.goodreads.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Gravatar",
		// checkBy:    "status_code",
		// rank:       5871,
		profileURL: "http://en.gravatar.com/%s",
		// mainUrl:    "http://en.gravatar.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Gumroad",
		// errorMsg:   "Page not found.",
		// checkBy:    "message",
		// rank:       3148,
		profileURL: "https://www.gumroad.com/%s",
		// mainUrl:    "https://www.gumroad.com/",
		checkerFn: bodyChecker("Page not found."),
	},
	{
		name: "HackerNews",
		// errorMsg:   "No such user.",
		// checkBy:    "message",
		// rank:       2499,
		profileURL: "https://news.ycombinator.com/user?id=%s",
		// mainUrl:    "https://news.ycombinator.com/",
		checkerFn: bodyChecker("No such user."),
	},
	{
		name: "HackerOne",
		// errorMsg:   "Page not found",
		// checkBy:    "message",
		// rank:       41985,
		profileURL: "https://hackerone.com/%s",
		// mainUrl:    "https://hackerone.com/",
		checkerFn: bodyChecker("Page not found"),
	},
	{
		name: "HackerRank",
		// errorMsg:   "Something went wrong",
		// checkBy:    "message",
		// rank:       3715,
		profileURL: "https://hackerrank.com/%s",
		// mainUrl:    "https://hackerrank.com/",
		checkerFn: bodyChecker("Something went wrong"),
	},
	{
		name: "House-Mixes.com",
		// errorMsg:   "Profile Not Found",
		// checkBy:    "message",
		// rank:       375927,
		profileURL: "https://www.house-mixes.com/profile/%s",
		// mainUrl:    "https://www.house-mixes.com/",
		checkerFn: bodyChecker("Profile Not Found"),
	},
	{
		name: "Houzz",
		// errorMsg:   "The page you requested was not found.",
		// checkBy:    "message",
		// rank:       2381,
		profileURL: "https://houzz.com/user/%s",
		// mainUrl:    "https://houzz.com/",
		checkerFn: bodyChecker("The page you requested was not found."),
	},
	{
		name: "HubPages",
		// checkBy:    "status_code",
		// rank:       9531,
		profileURL: "https://hubpages.com/@%s",
		// mainUrl:    "https://hubpages.com/",
		checkerFn: statusChecker,
	},
	{
		name: "IFTTT",
		// errorMsg: "The requested page or file does not exist",
		// checkBy:    "message",
		// rank:       6707,
		profileURL: "https://www.ifttt.com/p/%s",
		// mainUrl:    "https://www.ifttt.com/",
		checkerFn: bodyChecker("The requested page or file does not exist"),
	},
	{
		name: "ImageShack",
		// checkBy:    "response_url",
		// errorUrl:   "https://imageshack.us/",
		// rank:       58288,
		profileURL: "https://imageshack.us/user/%s",
		// mainUrl:    "https://imageshack.us/",
		checkerFn: redirectChecker("https://imageshack.us/"),
	},
	{
		name: "Imgur",
		// checkBy:    "status_code",
		// rank:       54,
		profileURL: "https://imgur.com/user/%s",
		// mainUrl:    "https://imgur.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Instagram",
		// errorMsg:   "The link you followed may be broken",
		// checkBy:    "message",
		// rank:       15,
		profileURL: "https://www.instagram.com/%s",
		// mainUrl:    "https://www.instagram.com/",
		checkerFn: bodyChecker("The link you followed may be broken"),
	},
	{
		name: "Instructables",
		// errorMsg:   "404: We're sorry, things break sometimes",
		// checkBy:    "message",
		// rank:       885,
		profileURL: "https://www.instructables.com/member/%s",
		// mainUrl:    "https://www.instructables.com/",
		checkerFn: bodyChecker("404: We're sorry, things break sometimes"),
	},
	{
		name: "Investing.com",
		// checkBy:    "status_code",
		// rank:       453,
		profileURL: "https://www.investing.com/traders/%s",
		// mainUrl:    "https://www.investing.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Issuu",
		// checkBy:    "status_code",
		// rank:       569,
		profileURL: "https://issuu.com/%s",
		// mainUrl:    "https://issuu.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Itch.io",
		// checkBy:    "status_code",
		// rank:       1999,
		profileURL: "https://%s.itch.io/",
		// mainUrl:    "https://itch.io/",
		checkerFn: statusChecker,
	},
	{
		name: "Jimdo",
		// checkBy:    "status_code",
		noPeriod: "True",
		// rank:       55600,
		profileURL: "https://%s.jimdosite.com",
		// mainUrl:    "https://jimdosite.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Kaggle",
		// checkBy:    "status_code",
		// rank:       2758,
		profileURL: "https://www.kaggle.com/%s",
		// mainUrl:    "https://www.kaggle.com/",
		checkerFn: statusChecker,
	},
	{
		name: "KanoWorld",
		// checkBy:    "status_code",
		// rank:       160033,
		profileURL: "https://api.kano.me/progress/user/%s",
		// mainUrl:    "https://world.kano.me/",
		checkerFn: statusChecker,
	},
	{
		name: "Keybase",
		// checkBy:    "status_code",
		// rank:       78228,
		profileURL: "https://keybase.io/%s",
		// mainUrl:    "https://keybase.io/",
		checkerFn: statusChecker,
	},
	{
		name: "Kik",
		// errorMsg:   "The page you requested was not found",
		// checkBy:    "message",
		// rank:       291593,
		profileURL: "https://ws2.kik.com/user/%s",
		// mainUrl:    "http://kik.me/",
		checkerFn: bodyChecker("The page you requested was not found"),
	},
	{
		name: "Kongregate",
		// errorMsg:   "Sorry, no account with that name was found.",
		// checkBy:    "message",
		// rank:       2235,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://www.kongregate.com/accounts/%s",
		// mainUrl:    "https://www.kongregate.com/",
		checkerFn: bodyChecker("Sorry, no account with that name was found."),
	},
	{
		name: "Launchpad",
		// checkBy:    "status_code",
		// rank:       7887,
		profileURL: "https://launchpad.net/~%s",
		// mainUrl:    "https://launchpad.net/",
		checkerFn: statusChecker,
	},
	{
		name: "LeetCode",
		// checkBy:    "status_code",
		// rank:       2834,
		profileURL: "https://leetcode.com/%s",
		// mainUrl:    "https://leetcode.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Letterboxd",
		// errorMsg:   "Sorry, we can\u2019t find the page you\u2019ve requested.",
		// checkBy:    "message",
		// rank:       3660,
		profileURL: "https://letterboxd.com/%s",
		// mainUrl:    "https://letterboxd.com/",
		checkerFn: bodyChecker("Sorry, we can\u2019t find the page you\u2019ve requested."),
	},
	{
		name: "LiveJournal",
		// checkBy:    "status_code",
		// rank:       237,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.livejournal.com",
		// mainUrl:    "https://www.livejournal.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Mastodon",
		// checkBy:    "status_code",
		// rank:       1389227,
		profileURL: "https://mstdn.io/@%s",
		// mainUrl:    "https://mstdn.io/",
		checkerFn: statusChecker,
	},
	{
		name: "Medium",
		// checkBy:    "status_code",
		// rank:       119,
		profileURL: "https://medium.com/@%s",
		// mainUrl:    "https://medium.com/",
		checkerFn: statusChecker,
	},
	{
		name: "MeetMe",
		//checkBy:    "response_url",
		// errorUrl:   "https://www.meetme.com/",
		// rank:       19198,
		profileURL: "https://www.meetme.com/%s",
		// mainUrl:    "https://www.meetme.com/"
		checkerFn: redirectChecker("https://www.meetme.com/"),
	},
	{
		name: "MixCloud",
		// errorMsg:   "Page Not Found",
		// checkBy:    "message",
		// rank:       2436,
		profileURL: "https://www.mixcloud.com/%s",
		// mainUrl:    "https://www.mixcloud.com/",
		checkerFn: bodyChecker("Page Not Found"),
	},
	{
		name: "MyAnimeList",
		// checkBy:    "status_code",
		// rank:       487,
		profileURL: "https://myanimelist.net/profile/%s",
		// mainUrl:    "https://myanimelist.net/",
		checkerFn: statusChecker,
	},
	{
		name: "Myspace",
		// checkBy:    "status_code",
		// rank:       4638,
		profileURL: "https://myspace.com/%s",
		// mainUrl:    "https://myspace.com/",
		checkerFn: statusChecker,
	},
	{
		name: "NameMC (Minecraft.net skins)",
		// errorMsg:   "Profiles: 0 results",
		// checkBy:    "message",
		// rank:       5520,
		profileURL: "https://namemc.com/profile/%s",
		// mainUrl:    "https://namemc.com/",
		checkerFn: bodyChecker("Profiles: 0 results"),
	},
	{
		name: "Newgrounds",
		// checkBy:    "status_code",
		// rank:       3402,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.newgrounds.com",
		// mainUrl:    "https://newgrounds.com",
		checkerFn: statusChecker,
	},
	{
		name: "OK",
		// checkBy:    "status_code",
		// rank:       63,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_.-]*$",
		profileURL: "https://ok.ru/%s",
		// mainUrl:    "https://ok.ru/",
		checkerFn: statusChecker,
	},
	{
		name: "Pastebin",
		// checkBy: "response_url",
		// errorUrl:   "https://pastebin.com/index",
		// rank:       918,
		profileURL: "https://pastebin.com/u/%s",
		// mainUrl:    "https://pastebin.com/",
		checkerFn: redirectChecker("https://pastebin.com/index"),
	},
	{
		name: "Patreon",
		// checkBy:    "status_code",
		// rank:       243,
		profileURL: "https://www.patreon.com/%s",
		// mainUrl:    "https://www.patreon.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Pexels",
		// errorMsg:   "Ouch, something went wrong!",
		// checkBy:    "message",
		// rank:       638,
		profileURL: "https://www.pexels.com/@%s",
		// mainUrl:    "https://www.pexels.com/",
		checkerFn: bodyChecker("Ouch, something went wrong!"),
	},
	{
		name: "Photobucket",
		// checkBy:    "status_code",
		// rank:       4037,
		profileURL: "https://photobucket.com/user/%s/library",
		// mainUrl:    "https://photobucket.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Pinterest",
		// checkBy:    "status_code",
		// rank:       73,
		profileURL: "https://www.pinterest.com/%s/",
		// mainUrl:    "https://www.pinterest.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Pixabay",
		// checkBy:    "status_code",
		// rank:       392,
		profileURL: "https://pixabay.com/en/users/%s",
		// mainUrl:    "https://pixabay.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Plug.DJ",
		// checkBy:    "status_code",
		// rank:       31273,
		profileURL: "https://plug.dj/@/%s",
		// mainUrl:    "https://plug.dj/",
		checkerFn: statusChecker,
	},
	{
		name: "Pokemon Showdown",
		// checkBy:    "status_code",
		// rank:       3508,
		profileURL: "https://pokemonshowdown.com/users/%s",
		// mainUrl:    "https://pokemonshowdown.com",
		checkerFn: statusChecker,
	},
	{
		name: "ProductHunt",
		// errorMsg:   "Product Hunt is a curation of the best new products",
		// checkBy:    "message",
		// rank:       3303,
		profileURL: "https://www.producthunt.com/@%s",
		// mainUrl:    "https://www.producthunt.com/",
		checkerFn: bodyChecker("Product Hunt is a curation of the best new products"),
	},
	{
		name: "Quora",
		// checkBy: "response_url",
		// errorUrl:   "https://www.quora.com/profile/%s",
		// rank:       69,
		profileURL: "https://www.quora.com/profile/%s",
		// mainUrl:    "https://www.quora.com/",
		checkerFn: redirectChecker("https://www.quora.com/profile/%s"),
	},
	{
		name: "Rajce.net",
		// errorMsg:   "410",
		// checkBy:    "message",
		// rank:       1486,
		profileURL: "https://%s.rajce.idnes.cz/",
		// mainUrl:    "https://www.rajce.idnes.cz/",
		checkerFn: bodyChecker("410"),
	},
	{
		name: "Rate Your Music",
		// checkBy:    "status_code",
		// rank:       4200,
		profileURL: "https://rateyourmusic.com/~%s",
		// mainUrl:    "https://rateyourmusic.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Reddit",
		// checkBy:    "status_code",
		// rank:       20,
		profileURL: "https://www.reddit.com/user/%s",
		// mainUrl:    "https://www.reddit.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Repl.it",
		// errorMsg:   "404",
		// checkBy:    "message",
		// rank:       7412,
		profileURL: "https://repl.it/@%s",
		// mainUrl:    "https://repl.it/",
		checkerFn: bodyChecker("404"),
	},
	{
		name: "ResearchGate",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.researchgate.net/directory/profiles",
		// rank: 114,

		// TODO: Regex
		regexCheck: "^[a-zA-Z0-9-_]+$",
		// regexCheck: "[w+_\\w+",
		profileURL: "https://www.researchgate.net/profile/%s",
		// mainUrl:    "https://www.researchgate.net/",
		checkerFn: redirectChecker("https://www.researchgate.net/directory/profiles"),
	},
	{
		name: "ReverbNation",
		// errorMsg:   "Sorry, we couldn't find that page",
		// checkBy:    "message",
		// rank:       11008,
		profileURL: "https://www.reverbnation.com/%s",
		// mainUrl:    "https://www.reverbnation.com/",
		checkerFn: bodyChecker("Sorry, we couldn't find that page"),
	},
	{
		name: "Roblox",
		// errorMsg:   "Page cannot be found or no longer exists",
		// checkBy:    "message",
		// rank:       108,
		profileURL: "https://www.roblox.com/user.aspx?username=%s",
		// mainUrl:    "https://www.roblox.com/",
		checkerFn: bodyChecker("Page cannot be found or no longer exists"),
	},
	{
		name: "Scribd",
		// errorMsg:   "Page not found",
		// checkBy:    "message",
		// rank:       159,
		profileURL: "https://www.scribd.com/%s",
		// mainUrl:    "https://www.scribd.com/",
		checkerFn: bodyChecker("Page not found"),
	},
	{
		name: "Signal",
		// errorMsg:   "Oops! That page doesn\u2019t exist or is private.",
		// checkBy:    "message",
		// rank:       1539044,
		profileURL: "https://community.signalusers.org/u/%s",
		// mainUrl:    "https://community.signalusers.org",
		checkerFn: bodyChecker("Oops! That page doesn\u2019t exist or is private."),
	},
	{
		name: "Slack",
		// checkBy:    "status_code",
		// rank:       179,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.slack.com",
		// mainUrl:    "https://slack.com",
		checkerFn: statusChecker,
	},
	{
		name: "SlideShare",
		// checkBy:    "status_code",
		// rank:       122,
		profileURL: "https://slideshare.net/%s",
		// mainUrl:    "https://slideshare.net/",
		checkerFn: statusChecker,
	},
	{
		name: "Smashcast",
		// checkBy:    "status_code",
		// rank:       117504,
		profileURL: "https://www.smashcast.tv/api/media/live/%s",
		// mainUrl:    "https://www.smashcast.tv/",
		checkerFn: statusChecker,
	},
	{
		name: "SoundCloud",
		// checkBy:    "status_code",
		// rank:       90,
		profileURL: "https://soundcloud.com/%s",
		// mainUrl:    "https://soundcloud.com/",
		checkerFn: statusChecker,
	},
	{
		name: "SourceForge",
		// checkBy:    "status_code",
		// rank:       363,
		profileURL: "https://sourceforge.net/u/%s",
		// mainUrl:    "https://sourceforge.net/",
		checkerFn: statusChecker,
	},
	{
		name: "Spotify",
		// checkBy:    "status_code",
		// rank:       102,
		profileURL: "https://open.spotify.com/user/%s",
		// mainUrl:    "https://open.spotify.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Star Citizen",
		// checkBy:    "status_code",
		// rank:       8581,
		profileURL: "https://robertsspaceindustries.com/citizens/%s",
		// mainUrl:    "https://robertsspaceindustries.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Steam",
		// errorMsg: "The specified profile could not be found",
		// checkBy:    "message",
		// rank:       126,
		profileURL: "https://steamcommunity.com/id/%s",
		// mainUrl:    "https://steamcommunity.com/",
		checkerFn: bodyChecker("The specified profile could not be found"),
	},
	{
		name: "SteamGroup",
		// errorMsg:   "No group could be retrieved for the given URL",
		// checkBy:    "message",
		// rank:       126,
		profileURL: "https://steamcommunity.com/groups/%s",
		// mainUrl:    "https://steamcommunity.com/",
		checkerFn: bodyChecker("No group could be retrieved for the given URL"),
	},
	{
		name: "Taringa",
		// checkBy:    "status_code",
		// rank:       945,
		profileURL: "https://www.taringa.net/%s",
		// mainUrl:    "https://taringa.net/",
		checkerFn: statusChecker,
	},
	{
		name: "Telegram",
		// errorMsg:   "<meta property=\"twitter:title\" content=\"Telegram: Contact",
		// checkBy:    "message",
		// rank:       593,
		profileURL: "https://t.me/%s",
		// mainUrl:    "https://t.me/",
		checkerFn: bodyChecker("<meta property=\"twitter:title\" content=\"Telegram: Contact"),
	},
	{
		name: "Tinder",
		// errorMsg:   "Looking for Someone?",
		// checkBy:    "message",
		// rank:       1694,
		profileURL: "https://www.gotinder.com/@%s",
		// mainUrl:    "https://tinder.com/",
		checkerFn: bodyChecker("Looking for Someone?"),
	},
	{
		name: "TradingView",
		// checkBy:    "status_code",
		// rank:       515,
		profileURL: "https://www.tradingview.com/u/%s/",
		// mainUrl:    "https://www.tradingview.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Trakt",
		// checkBy:    "status_code",
		// rank:       4647,
		profileURL: "https://www.trakt.tv/users/%s",
		// mainUrl:    "https://www.trakt.tv/",
		checkerFn: statusChecker,
	},
	{
		name: "Trip",
		// errorMsg:   "Page not found",
		// checkBy:    "message",
		// rank:       3085,
		profileURL: "https://www.trip.skyscanner.com/user/%s",
		// mainUrl:    "https://www.trip.skyscanner.com/",
		checkerFn: bodyChecker("Page not found"),
	},
	{
		name: "TripAdvisor",
		// errorMsg:   "This page is on vacation\u2026",
		// checkBy:    "message",
		// rank:       225,
		profileURL: "https://tripadvisor.com/members/%s",
		// mainUrl:    "https://tripadvisor.com/",
		checkerFn: bodyChecker("This page is on vacation\u2026"),
	},
	{
		name: "Twitch",
		// checkBy:    "status_code",
		// rank:       25,
		profileURL: "https://m.twitch.tv/%s",
		// mainUrl:    "https://www.twitch.tv/",
		checkerFn: statusChecker,
	},
	{
		name: "Twitter",
		// errorMsg:   "page doesn\u2019t exist",
		// checkBy:    "message",
		// rank:       10,
		profileURL: "https://www.twitter.com/%s",
		// mainUrl:    "https://www.twitter.com/",
		checkerFn: bodyChecker("page doesn\u2019t exist"),
	},
	{
		name: "Unsplash",
		// checkBy:    "status_code",
		// rank:       519,
		profileURL: "https://unsplash.com/@%s",
		// mainUrl:    "https://unsplash.com/",
		checkerFn: statusChecker,
	},
	{
		name: "VK",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.quora.com/profile/%s",
		// rank:       19,
		profileURL: "https://vk.com/%s",
		// mainUrl:    "https://vk.com/",
		checkerFn: statusChecker,
	},
	{
		name: "VSCO",
		// checkBy:    "status_code",
		// rank:       3366,
		profileURL: "https://vsco.co/%s",
		// mainUrl:    "https://vsco.co/",
		checkerFn: statusChecker,
	},
	{
		name: "Venmo",
		// checkBy:    "status_code",
		// rank:       5299,
		profileURL: "https://venmo.com/%s",
		// mainUrl:    "https://venmo.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Vimeo",
		// checkBy:    "status_code",
		// rank:       127,
		profileURL: "https://vimeo.com/%s",
		// mainUrl:    "https://vimeo.com/",
		checkerFn: statusChecker,
	},
	{
		name: "Virgool",
		// checkBy:    "message",
		// errorMsg:   "۴۰۴",
		// rank:       8175,
		profileURL: "https://virgool.io/@%s",
		// mainUrl:    "https://virgool.io/",
		checkerFn: bodyChecker("۴۰۴"),
	},
	{
		name: "VirusTotal",
		// errorMsg:   "not found",
		// checkBy:    "message",
		// rank:       3934,
		profileURL: "https://www.virustotal.com/ui/users/%s/trusted_users",
		// mainUrl:    "https://www.virustotal.com/",
		checkerFn: bodyChecker("not found"),
	},
	{
		name: "Wattpad",
		// errorMsg:   "userError-404",
		// checkBy:    "message",
		// rank:       495,
		profileURL: "https://www.wattpad.com/user/%s",
		// mainUrl:    "https://www.wattpad.com/",
		checkerFn: bodyChecker("userError-404"),
	},
	{
		name: "We Heart It",
		// errorMsg:   "Oops! You've landed on a moving target!",
		// checkBy:    "message",
		// rank:       3857,
		profileURL: "https://weheartit.com/%s",
		// mainUrl:    "https://weheartit.com/",
		checkerFn: bodyChecker("Oops! You've landed on a moving target!"),
	},
	{
		name: "WebNode",
		// checkBy:    "status_code",
		// rank:       17852,
		profileURL: "https://%s.webnode.cz/",
		// mainUrl:    "https://www.webnode.cz/",
		checkerFn: statusChecker,
	},
	{
		name: "Wikipedia",
		// errorMsg:   "If a page was recently created here, it may not be visible yet because of a delay in updating the database",
		// checkBy:    "message",
		// rank:       5,
		profileURL: "https://www.wikipedia.org/wiki/User:%s",
		// mainUrl:    "https://www.wikipedia.org/",
		checkerFn: bodyChecker("If a page was recently created here, it may not be visible yet because of a delay in updating the database"),
	},
	{
		name: "Wix",
		// checkBy:    "status_code",
		// rank:       402,
		profileURL: "https://%s.wix.com",
		// mainUrl:    "https://wix.com/",
		checkerFn: statusChecker,
	},
	{
		name: "WordPress",
		// checkBy:    "response_url",
		// errorUrl:   "wordpress.com/typo/?subdomain=",
		// rank:       47,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileURL: "https://%s.wordpress.com/",
		// mainUrl:    "https://wordpress.com",
		checkerFn: redirectChecker("wordpress.com/typo/?subdomain="),
	},
	{
		name: "YouNow",
		// errorMsg:   "No users found",
		// checkBy:    "message",
		// rank:       17205,
		profileURL: "https://www.younow.com/%s/",
		// mainUrl:    "https://www.younow.com/",
		probeURL:  "https://api.younow.com/php/api/broadcast/info/user=%s/",
		checkerFn: bodyChecker("No users found"),
	},
	{
		name: "YouPic",
		// checkBy:    "status_code",
		// rank:       43585,
		profileURL: "https://youpic.com/photographer/%s/",
		// mainUrl:    "https://youpic.com/",
		checkerFn: statusChecker,
	},
	{
		name: "YouTube",
		// errorMsg: "Not Found",
		// checkBy:    "message",
		// rank:       2,
		profileURL: "https://www.youtube.com/%s",
		// mainUrl:    "https://www.youtube.com/",
		checkerFn: bodyChecker("Not Found"),
	},
	{
		name: "Zhihu",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.zhihu.com/people/%s",
		// rank:       70,
		profileURL: "https://www.zhihu.com/people/%s",
		// mainUrl:    "https://www.zhihu.com/",
		checkerFn: redirectChecker("https://www.zhihu.com/people/%s"),
	},
	{
		name: "boingboing.net",
		// checkBy:    "status_code",
		// rank:       5301,
		profileURL: "https://bbs.boingboing.net/u/%s",
		// mainUrl:    "https://boingboing.net/",
		checkerFn: statusChecker,
	},
	{
		name: "devRant",
		//checkBy:    "response_url",
		//errorUrl:   "https://devrant.com/",
		// rank:       159249,
		profileURL: "https://devrant.com/users/%s",
		// mainUrl:    "https://devrant.com/",
		checkerFn: redirectChecker("https://devrant.com/"),
	},
	{
		name: "gfycat",
		// checkBy:    "status_code",
		// rank:       230,
		profileURL: "https://gfycat.com/@%s",
		// mainUrl:    "https://gfycat.com/",
		checkerFn: statusChecker,
	},
	{
		name: "iMGSRC.RU",
		// checkBy:    "response_url",
		// errorUrl:   "https://imgsrc.ru/",
		// rank:       8262,
		profileURL: "https://imgsrc.ru/main/user.php?user=%s",
		// mainUrl:    "https://imgsrc.ru/",
		checkerFn: redirectChecker("https://imgsrc.ru/"),
	},
	{
		name: "last.fm",
		// checkBy:    "status_code",
		// rank:       1242,
		profileURL: "https://last.fm/user/%s",
		// mainUrl:    "https://last.fm/",
		checkerFn: statusChecker,
	},
	{
		name: "mixer.com",
		// checkBy:    "status_code",
		// rank:       2145,
		profileURL: "https://mixer.com/%s",
		// mainUrl:    "https://mixer.com/",
		probeURL:  "https://mixer.com/api/v1/channels/%s",
		checkerFn: statusChecker,
	},
	{
		name: "osu!",
		// checkBy:    "status_code",
		// rank:       1813,
		profileURL: "https://osu.ppy.sh/users/%s",
		// mainUrl:    "https://osu.ppy.sh/",
		checkerFn: statusChecker,
	},
	{
		name: "Cent",
		// rank:       1813,
		profileURL: "https://beta.cent.co/@%s",
		checkerFn:  bodyChecker("<title>Cent</title>"),
	},
	{
		name: "Filmogs",
		// rank:       1813,
		profileURL: "https://www.filmo.gs/users/%s",
		checkerFn:  statusChecker,
	},
	{
		name:       "Packagist",
		profileURL: "https://packagist.org/packages/%s/",
		checkerFn:  redirectChecker("/search?q=%s&reason=vendor_not_found"),
	},
	{
		name:       "Trello",
		profileURL: "https://trello.com/%s",
		probeURL:   "https://trello.com/1/Members/%s",
		checkerFn:  statusChecker,
	},
	{
		name:       "Brew",
		profileURL: "https://www.brew.com/%s",
		probeURL:   "https://api.brew.com/api/public/podcast/%s",
		checkerFn:  statusChecker,
	},
	{
		name:       "BuyMeACoffee",
		profileURL: "https://buymeacoff.ee/%s",
		probeURL:   "https://www.buymeacoffee.com/%s",
		checkerFn:  bodyChecker("It returned a 404 error"),
	},

	{
		name:       "TikTok",
		profileURL: "https://www.tiktok.com/@%s",
		checkerFn:  bodyChecker("Couldn&#x27;t find this account"),
	},
	{
		name:       "NPM-Package",
		profileURL: "https://www.npmjs.com/package/%s",
		checkerFn:  statusChecker,
	},
	{
		name:       "NPM",
		profileURL: "https://www.npmjs.com/~%s",
		checkerFn:  statusChecker,
	},
	{
		name:       "Tumblr",
		profileURL: "https://%s.tumblr.com/",
		checkerFn:  statusChecker,
	},
	{
		name:       "OpenCollective",
		profileURL: "https://opencollective.com/%s",
		checkerFn:  bodyChecker("Open Collective - open your finances to your community"),
	},
	{
		name:       "PayPal",
		profileURL: "https://www.paypal.me/%s",
		checkerFn:  bodyChecker("name=\"twitter:title\" content=\"Get your very own PayPal.Me link\""),
	},
	{
		name:       "WordPressOrg",
		profileURL: "https://profiles.wordpress.org/%s/",
		checkerFn:  redirectChecker("https://wordpress.org"),
	},
	{
		name:       "PlayStore",
		profileURL: "https://play.google.com/store/apps/developer?id=%s",
		checkerFn:  statusChecker,
	},
}
