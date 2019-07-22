package main

type Site struct {
	name       string
	profileUrl string
	checkerFn  checkerFunc
	regexCheck string
	rank       int
	probeUrl   string
	noPeriod   string
}

var sites = []Site{
	Site{
		name: "500px",
		// errorMsg:   "Oops! This page doesn\u2019t exist.",
		// checkBy:    "message",
		rank:       3144,
		profileUrl: "https://500px.com/%s",
		// mainUrl:    "https://500px.com/",
		checkerFn: bodyChecker("Oops! This page doesn\u2019t exist."),
	},
	Site{
		name: "9GAG",
		// checkBy:    "status_code",
		rank:       308,
		profileUrl: "https://9gag.com/u/%s",
		// mainUrl:    "https://9gag.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "About.me",
		// checkBy:    "status_code",
		rank:       14151,
		profileUrl: "https://about.me/%s",
		// mainUrl:    "https://about.me/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Academia.edu",
		// checkBy:    "status_code",
		rank:       184,
		profileUrl: "https://independent.academia.edu/%s",
		// mainUrl:    "https://www.academia.edu/",
		checkerFn: statusChecker,
	},
	Site{
		name: "AngelList",
		// checkBy:    "status_code",
		rank:       3377,
		profileUrl: "https://angel.co/%s",
		// mainUrl:    "https://angel.co/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Aptoide",
		// checkBy:    "status_code",
		rank:       4933,
		profileUrl: "https://%s.en.aptoide.com/",
		// mainUrl:    "https://en.aptoide.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "AskFM",
		// checkBy:    "status_code",
		rank:       1162,
		profileUrl: "https://ask.fm/%s",
		// mainUrl:    "https://ask.fm/",
		checkerFn: statusChecker,
	},
	Site{
		name: "BLIP.fm",
		// checkBy:    "status_code",
		rank:       366949,
		profileUrl: "https://blip.fm/%s",
		// mainUrl:    "https://blip.fm/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Badoo",
		// checkBy:    "status_code",
		rank:       1268,
		profileUrl: "https://badoo.com/profile/%s",
		// mainUrl:    "https://badoo.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Bandcamp",
		// checkBy:    "status_code",
		rank:       607,
		profileUrl: "https://www.bandcamp.com/%s",
		// mainUrl:    "https://www.bandcamp.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Basecamp",
		//errorMsg:   "The account you were looking for doesn't exist",
		// checkBy:    "message",
		rank:       1754,
		profileUrl: "https://%s.basecamphq.com",
		// mainUrl:    "https://basecamp.com/",
		checkerFn: bodyChecker("The account you were looking for doesn't exist"),
	},
	Site{
		name: "Behance",
		// checkBy:    "status_code",
		rank:       384,
		profileUrl: "https://www.behance.net/%s",
		// mainUrl:    "https://www.behance.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "BitBucket",
		// checkBy:    "status_code",
		rank:       989,
		profileUrl: "https://bitbucket.org/%s/",
		// mainUrl:    "https://bitbucket.org/",
		checkerFn: statusChecker,
	},
	Site{
		name: "BitCoinForum",
		// errorMsg:   "The user whose profile you are trying to view does not exist.",
		// checkBy:    "message",
		rank:       507091,
		profileUrl: "https://bitcoinforum.com/profile/%s",
		// mainUrl:    "https://bitcoinforum.com",
		checkerFn: bodyChecker("The user whose profile you are trying to view does not exist."),
	},
	Site{
		name: "Blogger",
		// checkBy:    "status_code",
		rank:       211,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.blogspot.com",
		// mainUrl:    "https://www.blogger.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "BuzzFeed",
		// checkBy:    "status_code",
		rank:       283,
		profileUrl: "https://buzzfeed.com/%s",
		// mainUrl:    "https://buzzfeed.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Canva",
		// checkBy: "status_code",
		// errorUrl:   "https://www.canva.com/%s",
		rank:       172,
		profileUrl: "https://www.canva.com/%s",
		// mainUrl:    "https://www.canva.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Carbonmade",
		// checkBy:    "response_url",
		// errorUrl:   "https://carbonmade.com/fourohfour?domain=%s.carbonmade.com",
		rank:       31822,
		profileUrl: "https://%s.carbonmade.com",
		// mainUrl:    "https://carbonmade.com/",
		checkerFn: redirectChecker("https://carbonmade.com/fourohfour?domain=%s.carbonmade.com"),
	},
	Site{
		name: "CashMe",
		// checkBy:    "status_code",
		rank:       12351551,
		profileUrl: "https://cash.me/%s",
		// mainUrl:    "https://cash.me/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Cloob",
		// checkBy:    "status_code",
		rank:       9254,
		profileUrl: "https://www.cloob.com/name/%s",
		// mainUrl:    "https://www.cloob.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Codecademy",
		// checkBy:    "status_code",
		rank:       2943,
		profileUrl: "https://www.codecademy.com/%s",
		// mainUrl:    "https://www.codecademy.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Codechef",
		// checkBy:    "status_code",
		rank:       9098,
		profileUrl: "https://www.codechef.com/users/%s",
		// mainUrl:    "https://www.codechef.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Codementor",
		// checkBy:    "status_code",
		rank:       9600,
		profileUrl: "https://www.codementor.io/%s",
		// mainUrl:    "https://www.codementor.io/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Codepen",
		// checkBy:    "status_code",
		rank:       831,
		profileUrl: "https://codepen.io/%s",
		// mainUrl:    "https://codepen.io/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Coderwall",
		//errorMsg:   "404! Our feels when that url is used",
		//checkBy:    "message",
		rank:       10061,
		profileUrl: "https://coderwall.com/%s",
		// mainUrl:    "https://coderwall.com/",
		checkerFn: bodyChecker("404! Our feels when that url is used"),
	},
	Site{
		name: "ColourLovers",
		// checkBy:    "status_code",
		rank:       32386,
		profileUrl: "https://www.colourlovers.com/lover/%s",
		// mainUrl:    "https://www.colourlovers.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Contently",
		// errorMsg:   "We can't find that page!",
		// checkBy:    "message",
		rank:       58845,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.contently.com/",
		// mainUrl:    "https://contently.com/",
		checkerFn: bodyChecker("We can't find that page!"),
	},
	Site{
		name: "Coroflot",
		// checkBy:    "status_code",
		rank:       40375,
		profileUrl: "https://www.coroflot.com/%s",
		// mainUrl:    "https://coroflot.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "CreativeMarket",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.creativemarket.com/",
		rank:       2100,
		profileUrl: "https://creativemarket.com/%s",
		// mainUrl:    "https://creativemarket.com/",
		checkerFn: redirectChecker("https://www.creativemarket.com/"),
	},
	Site{
		name: "Crevado",
		// checkBy:    "status_code",
		rank:       178744,
		profileUrl: "https://%s.crevado.com",
		// mainUrl:    "https://crevado.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Crunchyroll",
		// checkBy:    "status_code",
		rank:       383,
		profileUrl: "https://www.crunchyroll.com/user/%s",
		// mainUrl:    "https://www.crunchyroll.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "DEV Community",
		// checkBy:    "status_code",
		rank:       5203,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://dev.to/%s",
		// mainUrl:    "https://dev.to/",
		checkerFn: statusChecker,
	},
	Site{
		name: "DailyMotion",
		// checkBy:    "status_code",
		rank:       120,
		profileUrl: "https://www.dailymotion.com/%s",
		// mainUrl:    "https://www.dailymotion.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Designspiration",
		// checkBy:    "status_code",
		rank:       26101,
		profileUrl: "https://www.designspiration.net/%s/",
		// mainUrl:    "https://www.designspiration.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "DeviantART",
		// checkBy:    "status_code",
		rank:       222,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.deviantart.com",
		// mainUrl:    "https://deviantart.com",
		checkerFn: statusChecker,
	},
	Site{
		name: "Discogs",
		// checkBy:    "status_code",
		rank:       614,
		profileUrl: "https://www.discogs.com/user/%s",
		// mainUrl:    "https://www.discogs.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Disqus",
		// checkBy:    "status_code",
		rank:       1297,
		profileUrl: "https://disqus.com/%s",
		// mainUrl:    "https://disqus.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Docker Hub",
		// checkBy:    "status_code",
		rank:       1993,
		profileUrl: "https://hub.docker.com/u/%s/",
		// mainUrl:    "https://hub.docker.com/",
		probeUrl:  "https://hub.docker.com/v2/users/%s/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Dribbble",
		// errorMsg:   "Whoops, that page is gone.",
		// checkBy:    "message",
		rank:       954,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://dribbble.com/%s",
		// mainUrl:    "https://dribbble.com/",
		checkerFn: bodyChecker("Whoops, that page is gone."),
	},
	Site{
		name: "EVE Online",
		// checkBy:    "response_url",
		// errorUrl:   "https://eveonline.com",
		rank:       11619,
		profileUrl: "https://evewho.com/pilot/%s/",
		// mainUrl:    "https://eveonline.com",
		checkerFn: redirectChecker("https://eveonline.com"),
	},
	Site{
		name: "Ebay",
		// errorMsg:   "The User ID you entered was not found",
		// checkBy:    "message",
		rank:       45,
		profileUrl: "https://www.ebay.com/usr/%s",
		// mainUrl:    "https://www.ebay.com/",
		checkerFn: bodyChecker("The User ID you entered was not found"),
	},
	Site{
		name: "Ello",
		// errorMsg:   "We couldn't find the page you're looking for",
		// checkBy:    "message",
		rank:       36947,
		profileUrl: "https://ello.co/%s",
		// mainUrl:    "https://ello.co/",
		checkerFn: bodyChecker("We couldn't find the page you're looking for"),
	},
	Site{
		name: "Etsy",
		// checkBy:    "status_code",
		rank:       166,
		profileUrl: "https://www.etsy.com/shop/%s",
		// mainUrl:    "https://www.etsy.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "EyeEm",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.eyeem.com/",
		rank:       32792,
		profileUrl: "https://www.eyeem.com/u/%s",
		// mainUrl:    "https://www.eyeem.com/",
		checkerFn: redirectChecker("https://www.eyeem.com/"),
	},
	Site{
		name: "Facebook",
		// checkBy: "status_code",
		rank: 3,
		// TODO: Regex
		regexCheck: "^[a-zA-Z0-9]{4,49}$",
		// regexCheck: "^[a-zA-Z0-9]{4,49}(?<!.com|.org|.net)$",
		profileUrl: "https://www.facebook.com/%s",
		// mainUrl:    "https://www.facebook.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Fandom",
		// checkBy:    "status_code",
		rank:       58,
		profileUrl: "https://www.fandom.com/u/%s",
		// mainUrl:    "https://www.fandom.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Flickr",
		// checkBy:    "status_code",
		rank:       454,
		profileUrl: "https://www.flickr.com/people/%s",
		// mainUrl:    "https://www.flickr.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Flipboard",
		// checkBy:    "status_code",
		rank:       3603,
		regexCheck: "^([a-zA-Z0-9_]){1,15}$",
		profileUrl: "https://flipboard.com/@%s",
		// mainUrl:    "https://flipboard.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Foursquare",
		// checkBy:    "status_code",
		rank:       2105,
		profileUrl: "https://foursquare.com/%s",
		// mainUrl:    "https://foursquare.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Giphy",
		// checkBy:    "status_code",
		rank:       472,
		profileUrl: "https://giphy.com/%s",
		// mainUrl:    "https://giphy.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "GitHub",
		// checkBy: "status_code",
		rank: 55,
		// TODO: Regex
		regexCheck: "^[a-zA-Z0-9-]{0,38}$",
		// regexCheck: "^[a-zA-Z0-9](?:[a-zA-Z0-9]|-(?=[a-zA-Z0-9])){0,38}$",
		profileUrl: "https://www.github.com/%s",
		// mainUrl:    "https://www.github.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "GitLab",
		// errorMsg: "[]",
		// checkBy:    "message",
		rank:       1874,
		profileUrl: "https://gitlab.com/%s",
		// mainUrl:    "https://gitlab.com/",
		probeUrl:  "https://gitlab.com/api/v4/users?username=%s",
		checkerFn: bodyChecker("Page Not Found"),
	},
	Site{
		name: "Gitee",
		// checkBy:    "status_code",
		rank:       2848,
		profileUrl: "https://gitee.com/%s",
		// mainUrl:    "https://gitee.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "GoodReads",
		// checkBy:    "status_code",
		rank:       404,
		profileUrl: "https://www.goodreads.com/%s",
		// mainUrl:    "https://www.goodreads.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Gravatar",
		// checkBy:    "status_code",
		rank:       5871,
		profileUrl: "http://en.gravatar.com/%s",
		// mainUrl:    "http://en.gravatar.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Gumroad",
		// errorMsg:   "Page not found.",
		// checkBy:    "message",
		rank:       3148,
		profileUrl: "https://www.gumroad.com/%s",
		// mainUrl:    "https://www.gumroad.com/",
		checkerFn: bodyChecker("Page not found."),
	},
	Site{
		name: "HackerNews",
		// errorMsg:   "No such user.",
		// checkBy:    "message",
		rank:       2499,
		profileUrl: "https://news.ycombinator.com/user?id=%s",
		// mainUrl:    "https://news.ycombinator.com/",
		checkerFn: bodyChecker("No such user."),
	},
	Site{
		name: "HackerOne",
		// errorMsg:   "Page not found",
		// checkBy:    "message",
		rank:       41985,
		profileUrl: "https://hackerone.com/%s",
		// mainUrl:    "https://hackerone.com/",
		checkerFn: bodyChecker("Page not found"),
	},
	Site{
		name: "HackerRank",
		// errorMsg:   "Something went wrong",
		// checkBy:    "message",
		rank:       3715,
		profileUrl: "https://hackerrank.com/%s",
		// mainUrl:    "https://hackerrank.com/",
		checkerFn: bodyChecker("Something went wrong"),
	},
	Site{
		name: "House-Mixes.com",
		// errorMsg:   "Profile Not Found",
		// checkBy:    "message",
		rank:       375927,
		profileUrl: "https://www.house-mixes.com/profile/%s",
		// mainUrl:    "https://www.house-mixes.com/",
		checkerFn: bodyChecker("Profile Not Found"),
	},
	Site{
		name: "Houzz",
		// errorMsg:   "The page you requested was not found.",
		// checkBy:    "message",
		rank:       2381,
		profileUrl: "https://houzz.com/user/%s",
		// mainUrl:    "https://houzz.com/",
		checkerFn: bodyChecker("The page you requested was not found."),
	},
	Site{
		name: "HubPages",
		// checkBy:    "status_code",
		rank:       9531,
		profileUrl: "https://hubpages.com/@%s",
		// mainUrl:    "https://hubpages.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "IFTTT",
		// errorMsg: "The requested page or file does not exist",
		// checkBy:    "message",
		rank:       6707,
		profileUrl: "https://www.ifttt.com/p/%s",
		// mainUrl:    "https://www.ifttt.com/",
		checkerFn: bodyChecker("The requested page or file does not exist"),
	},
	Site{
		name: "ImageShack",
		// checkBy:    "response_url",
		// errorUrl:   "https://imageshack.us/",
		rank:       58288,
		profileUrl: "https://imageshack.us/user/%s",
		// mainUrl:    "https://imageshack.us/",
		checkerFn: redirectChecker("https://imageshack.us/"),
	},
	Site{
		name: "Imgur",
		// checkBy:    "status_code",
		rank:       54,
		profileUrl: "https://imgur.com/user/%s",
		// mainUrl:    "https://imgur.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Instagram",
		// errorMsg:   "The link you followed may be broken",
		// checkBy:    "message",
		rank:       15,
		profileUrl: "https://www.instagram.com/%s",
		// mainUrl:    "https://www.instagram.com/",
		checkerFn: bodyChecker("The link you followed may be broken"),
	},
	Site{
		name: "Instructables",
		// errorMsg:   "404: We're sorry, things break sometimes",
		// checkBy:    "message",
		rank:       885,
		profileUrl: "https://www.instructables.com/member/%s",
		// mainUrl:    "https://www.instructables.com/",
		checkerFn: bodyChecker("404: We're sorry, things break sometimes"),
	},
	Site{
		name: "Investing.com",
		// checkBy:    "status_code",
		rank:       453,
		profileUrl: "https://www.investing.com/traders/%s",
		// mainUrl:    "https://www.investing.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Issuu",
		// checkBy:    "status_code",
		rank:       569,
		profileUrl: "https://issuu.com/%s",
		// mainUrl:    "https://issuu.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Itch.io",
		// checkBy:    "status_code",
		rank:       1999,
		profileUrl: "https://%s.itch.io/",
		// mainUrl:    "https://itch.io/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Jimdo",
		// checkBy:    "status_code",
		noPeriod:   "True",
		rank:       55600,
		profileUrl: "https://%s.jimdosite.com",
		// mainUrl:    "https://jimdosite.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Kaggle",
		// checkBy:    "status_code",
		rank:       2758,
		profileUrl: "https://www.kaggle.com/%s",
		// mainUrl:    "https://www.kaggle.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "KanoWorld",
		// checkBy:    "status_code",
		rank:       160033,
		profileUrl: "https://api.kano.me/progress/user/%s",
		// mainUrl:    "https://world.kano.me/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Keybase",
		// checkBy:    "status_code",
		rank:       78228,
		profileUrl: "https://keybase.io/%s",
		// mainUrl:    "https://keybase.io/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Kik",
		// errorMsg:   "The page you requested was not found",
		// checkBy:    "message",
		rank:       291593,
		profileUrl: "https://ws2.kik.com/user/%s",
		// mainUrl:    "http://kik.me/",
		checkerFn: bodyChecker("The page you requested was not found"),
	},
	Site{
		name: "Kongregate",
		// errorMsg:   "Sorry, no account with that name was found.",
		// checkBy:    "message",
		rank:       2235,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://www.kongregate.com/accounts/%s",
		// mainUrl:    "https://www.kongregate.com/",
		checkerFn: bodyChecker("Sorry, no account with that name was found."),
	},
	Site{
		name: "Launchpad",
		// checkBy:    "status_code",
		rank:       7887,
		profileUrl: "https://launchpad.net/~%s",
		// mainUrl:    "https://launchpad.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "LeetCode",
		// checkBy:    "status_code",
		rank:       2834,
		profileUrl: "https://leetcode.com/%s",
		// mainUrl:    "https://leetcode.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Letterboxd",
		// errorMsg:   "Sorry, we can\u2019t find the page you\u2019ve requested.",
		// checkBy:    "message",
		rank:       3660,
		profileUrl: "https://letterboxd.com/%s",
		// mainUrl:    "https://letterboxd.com/",
		checkerFn: bodyChecker("Sorry, we can\u2019t find the page you\u2019ve requested."),
	},
	Site{
		name: "LiveJournal",
		// checkBy:    "status_code",
		rank:       237,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.livejournal.com",
		// mainUrl:    "https://www.livejournal.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Mastodon",
		// checkBy:    "status_code",
		rank:       1389227,
		profileUrl: "https://mstdn.io/@%s",
		// mainUrl:    "https://mstdn.io/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Medium",
		// checkBy:    "status_code",
		rank:       119,
		profileUrl: "https://medium.com/@%s",
		// mainUrl:    "https://medium.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "MeetMe",
		//checkBy:    "response_url",
		// errorUrl:   "https://www.meetme.com/",
		rank:       19198,
		profileUrl: "https://www.meetme.com/%s",
		// mainUrl:    "https://www.meetme.com/"
		checkerFn: redirectChecker("https://www.meetme.com/"),
	},
	Site{
		name: "MixCloud",
		// errorMsg:   "Page Not Found",
		// checkBy:    "message",
		rank:       2436,
		profileUrl: "https://www.mixcloud.com/%s",
		// mainUrl:    "https://www.mixcloud.com/",
		checkerFn: bodyChecker("Page Not Found"),
	},
	Site{
		name: "MyAnimeList",
		// checkBy:    "status_code",
		rank:       487,
		profileUrl: "https://myanimelist.net/profile/%s",
		// mainUrl:    "https://myanimelist.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Myspace",
		// checkBy:    "status_code",
		rank:       4638,
		profileUrl: "https://myspace.com/%s",
		// mainUrl:    "https://myspace.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "NameMC (Minecraft.net skins)",
		// errorMsg:   "Profiles: 0 results",
		// checkBy:    "message",
		rank:       5520,
		profileUrl: "https://namemc.com/profile/%s",
		// mainUrl:    "https://namemc.com/",
		checkerFn: bodyChecker("Profiles: 0 results"),
	},
	Site{
		name: "Newgrounds",
		// checkBy:    "status_code",
		rank:       3402,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.newgrounds.com",
		// mainUrl:    "https://newgrounds.com",
		checkerFn: statusChecker,
	},
	Site{
		name: "OK",
		// checkBy:    "status_code",
		rank:       63,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_.-]*$",
		profileUrl: "https://ok.ru/%s",
		// mainUrl:    "https://ok.ru/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Pastebin",
		// checkBy: "response_url",
		// errorUrl:   "https://pastebin.com/index",
		rank:       918,
		profileUrl: "https://pastebin.com/u/%s",
		// mainUrl:    "https://pastebin.com/",
		checkerFn: redirectChecker("https://pastebin.com/index"),
	},
	Site{
		name: "Patreon",
		// checkBy:    "status_code",
		rank:       243,
		profileUrl: "https://www.patreon.com/%s",
		// mainUrl:    "https://www.patreon.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Pexels",
		// errorMsg:   "Ouch, something went wrong!",
		// checkBy:    "message",
		rank:       638,
		profileUrl: "https://www.pexels.com/@%s",
		// mainUrl:    "https://www.pexels.com/",
		checkerFn: bodyChecker("Ouch, something went wrong!"),
	},
	Site{
		name: "Photobucket",
		// checkBy:    "status_code",
		rank:       4037,
		profileUrl: "https://photobucket.com/user/%s/library",
		// mainUrl:    "https://photobucket.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Pinterest",
		// checkBy:    "status_code",
		rank:       73,
		profileUrl: "https://www.pinterest.com/%s/",
		// mainUrl:    "https://www.pinterest.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Pixabay",
		// checkBy:    "status_code",
		rank:       392,
		profileUrl: "https://pixabay.com/en/users/%s",
		// mainUrl:    "https://pixabay.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Plug.DJ",
		// checkBy:    "status_code",
		rank:       31273,
		profileUrl: "https://plug.dj/@/%s",
		// mainUrl:    "https://plug.dj/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Pokemon Showdown",
		// checkBy:    "status_code",
		rank:       3508,
		profileUrl: "https://pokemonshowdown.com/users/%s",
		// mainUrl:    "https://pokemonshowdown.com",
		checkerFn: statusChecker,
	},
	Site{
		name: "ProductHunt",
		// errorMsg:   "Product Hunt is a curation of the best new products",
		// checkBy:    "message",
		rank:       3303,
		profileUrl: "https://www.producthunt.com/@%s",
		// mainUrl:    "https://www.producthunt.com/",
		checkerFn: bodyChecker("Product Hunt is a curation of the best new products"),
	},
	Site{
		name: "Quora",
		// checkBy: "response_url",
		// errorUrl:   "https://www.quora.com/profile/%s",
		rank:       69,
		profileUrl: "https://www.quora.com/profile/%s",
		// mainUrl:    "https://www.quora.com/",
		checkerFn: redirectChecker("https://www.quora.com/profile/%s"),
	},
	Site{
		name: "Rajce.net",
		// errorMsg:   "410",
		// checkBy:    "message",
		rank:       1486,
		profileUrl: "https://%s.rajce.idnes.cz/",
		// mainUrl:    "https://www.rajce.idnes.cz/",
		checkerFn: bodyChecker("410"),
	},
	Site{
		name: "Rate Your Music",
		// checkBy:    "status_code",
		rank:       4200,
		profileUrl: "https://rateyourmusic.com/~%s",
		// mainUrl:    "https://rateyourmusic.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Reddit",
		// checkBy:    "status_code",
		rank:       20,
		profileUrl: "https://www.reddit.com/user/%s",
		// mainUrl:    "https://www.reddit.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Repl.it",
		// errorMsg:   "404",
		// checkBy:    "message",
		rank:       7412,
		profileUrl: "https://repl.it/@%s",
		// mainUrl:    "https://repl.it/",
		checkerFn: bodyChecker("404"),
	},
	Site{
		name: "ResearchGate",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.researchgate.net/directory/profiles",
		rank: 114,

		// TODO: Regex
		regexCheck: "^[a-zA-Z0-9-_]+$",
		// regexCheck: "[w+_\\w+",
		profileUrl: "https://www.researchgate.net/profile/%s",
		// mainUrl:    "https://www.researchgate.net/",
		checkerFn: redirectChecker("https://www.researchgate.net/directory/profiles"),
	},
	Site{
		name: "ReverbNation",
		// errorMsg:   "Sorry, we couldn't find that page",
		// checkBy:    "message",
		rank:       11008,
		profileUrl: "https://www.reverbnation.com/%s",
		// mainUrl:    "https://www.reverbnation.com/",
		checkerFn: bodyChecker("Sorry, we couldn't find that page"),
	},
	Site{
		name: "Roblox",
		// errorMsg:   "Page cannot be found or no longer exists",
		// checkBy:    "message",
		rank:       108,
		profileUrl: "https://www.roblox.com/user.aspx?username=%s",
		// mainUrl:    "https://www.roblox.com/",
		checkerFn: bodyChecker("Page cannot be found or no longer exists"),
	},
	Site{
		name: "Scribd",
		// errorMsg:   "Page not found",
		// checkBy:    "message",
		rank:       159,
		profileUrl: "https://www.scribd.com/%s",
		// mainUrl:    "https://www.scribd.com/",
		checkerFn: bodyChecker("Page not found"),
	},
	Site{
		name: "Signal",
		// errorMsg:   "Oops! That page doesn\u2019t exist or is private.",
		// checkBy:    "message",
		rank:       1539044,
		profileUrl: "https://community.signalusers.org/u/%s",
		// mainUrl:    "https://community.signalusers.org",
		checkerFn: bodyChecker("Oops! That page doesn\u2019t exist or is private."),
	},
	Site{
		name: "Slack",
		// checkBy:    "status_code",
		rank:       179,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.slack.com",
		// mainUrl:    "https://slack.com",
		checkerFn: statusChecker,
	},
	Site{
		name: "SlideShare",
		// checkBy:    "status_code",
		rank:       122,
		profileUrl: "https://slideshare.net/%s",
		// mainUrl:    "https://slideshare.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Smashcast",
		// checkBy:    "status_code",
		rank:       117504,
		profileUrl: "https://www.smashcast.tv/api/media/live/%s",
		// mainUrl:    "https://www.smashcast.tv/",
		checkerFn: statusChecker,
	},
	Site{
		name: "SoundCloud",
		// checkBy:    "status_code",
		rank:       90,
		profileUrl: "https://soundcloud.com/%s",
		// mainUrl:    "https://soundcloud.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "SourceForge",
		// checkBy:    "status_code",
		rank:       363,
		profileUrl: "https://sourceforge.net/u/%s",
		// mainUrl:    "https://sourceforge.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Spotify",
		// checkBy:    "status_code",
		rank:       102,
		profileUrl: "https://open.spotify.com/user/%s",
		// mainUrl:    "https://open.spotify.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Star Citizen",
		// checkBy:    "status_code",
		rank:       8581,
		profileUrl: "https://robertsspaceindustries.com/citizens/%s",
		// mainUrl:    "https://robertsspaceindustries.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Steam",
		// errorMsg: "The specified profile could not be found",
		// checkBy:    "message",
		rank:       126,
		profileUrl: "https://steamcommunity.com/id/%s",
		// mainUrl:    "https://steamcommunity.com/",
		checkerFn: bodyChecker("The specified profile could not be found"),
	},
	Site{
		name: "SteamGroup",
		// errorMsg:   "No group could be retrieved for the given URL",
		// checkBy:    "message",
		rank:       126,
		profileUrl: "https://steamcommunity.com/groups/%s",
		// mainUrl:    "https://steamcommunity.com/",
		checkerFn: bodyChecker("No group could be retrieved for the given URL"),
	},
	Site{
		name: "Taringa",
		// checkBy:    "status_code",
		rank:       945,
		profileUrl: "https://www.taringa.net/%s",
		// mainUrl:    "https://taringa.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Telegram",
		// errorMsg:   "<meta property=\"twitter:title\" content=\"Telegram: Contact",
		// checkBy:    "message",
		rank:       593,
		profileUrl: "https://t.me/%s",
		// mainUrl:    "https://t.me/",
		checkerFn: bodyChecker("<meta property=\"twitter:title\" content=\"Telegram: Contact"),
	},
	Site{
		name: "Tinder",
		// errorMsg:   "Looking for Someone?",
		// checkBy:    "message",
		rank:       1694,
		profileUrl: "https://www.gotinder.com/@%s",
		// mainUrl:    "https://tinder.com/",
		checkerFn: bodyChecker("Looking for Someone?"),
	},
	Site{
		name: "TradingView",
		// checkBy:    "status_code",
		rank:       515,
		profileUrl: "https://www.tradingview.com/u/%s/",
		// mainUrl:    "https://www.tradingview.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Trakt",
		// checkBy:    "status_code",
		rank:       4647,
		profileUrl: "https://www.trakt.tv/users/%s",
		// mainUrl:    "https://www.trakt.tv/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Trip",
		// errorMsg:   "Page not found",
		// checkBy:    "message",
		rank:       3085,
		profileUrl: "https://www.trip.skyscanner.com/user/%s",
		// mainUrl:    "https://www.trip.skyscanner.com/",
		checkerFn: bodyChecker("Page not found"),
	},
	Site{
		name: "TripAdvisor",
		// errorMsg:   "This page is on vacation\u2026",
		// checkBy:    "message",
		rank:       225,
		profileUrl: "https://tripadvisor.com/members/%s",
		// mainUrl:    "https://tripadvisor.com/",
		checkerFn: bodyChecker("This page is on vacation\u2026"),
	},
	Site{
		name: "Twitch",
		// checkBy:    "status_code",
		rank:       25,
		profileUrl: "https://m.twitch.tv/%s",
		// mainUrl:    "https://www.twitch.tv/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Twitter",
		// errorMsg:   "page doesn\u2019t exist",
		// checkBy:    "message",
		rank:       10,
		profileUrl: "https://www.twitter.com/%s",
		// mainUrl:    "https://www.twitter.com/",
		checkerFn: bodyChecker("page doesn\u2019t exist"),
	},
	Site{
		name: "Unsplash",
		// checkBy:    "status_code",
		rank:       519,
		profileUrl: "https://unsplash.com/@%s",
		// mainUrl:    "https://unsplash.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "VK",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.quora.com/profile/%s",
		rank:       19,
		profileUrl: "https://vk.com/%s",
		// mainUrl:    "https://vk.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "VSCO",
		// checkBy:    "status_code",
		rank:       3366,
		profileUrl: "https://vsco.co/%s",
		// mainUrl:    "https://vsco.co/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Venmo",
		// checkBy:    "status_code",
		rank:       5299,
		profileUrl: "https://venmo.com/%s",
		// mainUrl:    "https://venmo.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Vimeo",
		// checkBy:    "status_code",
		rank:       127,
		profileUrl: "https://vimeo.com/%s",
		// mainUrl:    "https://vimeo.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Virgool",
		// checkBy:    "message",
		// errorMsg:   "۴۰۴",
		rank:       8175,
		profileUrl: "https://virgool.io/@%s",
		// mainUrl:    "https://virgool.io/",
		checkerFn: bodyChecker("۴۰۴"),
	},
	Site{
		name: "VirusTotal",
		// errorMsg:   "not found",
		// checkBy:    "message",
		rank:       3934,
		profileUrl: "https://www.virustotal.com/ui/users/%s/trusted_users",
		// mainUrl:    "https://www.virustotal.com/",
		checkerFn: bodyChecker("not found"),
	},
	Site{
		name: "Wattpad",
		// errorMsg:   "userError-404",
		// checkBy:    "message",
		rank:       495,
		profileUrl: "https://www.wattpad.com/user/%s",
		// mainUrl:    "https://www.wattpad.com/",
		checkerFn: bodyChecker("userError-404"),
	},
	Site{
		name: "We Heart It",
		// errorMsg:   "Oops! You've landed on a moving target!",
		// checkBy:    "message",
		rank:       3857,
		profileUrl: "https://weheartit.com/%s",
		// mainUrl:    "https://weheartit.com/",
		checkerFn: bodyChecker("Oops! You've landed on a moving target!"),
	},
	Site{
		name: "WebNode",
		// checkBy:    "status_code",
		rank:       17852,
		profileUrl: "https://%s.webnode.cz/",
		// mainUrl:    "https://www.webnode.cz/",
		checkerFn: statusChecker,
	},
	Site{
		name: "Wikipedia",
		// errorMsg:   "If a page was recently created here, it may not be visible yet because of a delay in updating the database",
		// checkBy:    "message",
		rank:       5,
		profileUrl: "https://www.wikipedia.org/wiki/User:%s",
		// mainUrl:    "https://www.wikipedia.org/",
		checkerFn: bodyChecker("If a page was recently created here, it may not be visible yet because of a delay in updating the database"),
	},
	Site{
		name: "Wix",
		// checkBy:    "status_code",
		rank:       402,
		profileUrl: "https://%s.wix.com",
		// mainUrl:    "https://wix.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "WordPress",
		// checkBy:    "response_url",
		// errorUrl:   "wordpress.com/typo/?subdomain=",
		rank:       47,
		regexCheck: "^[a-zA-Z][a-zA-Z0-9_-]*$",
		profileUrl: "https://%s.wordpress.com/",
		// mainUrl:    "https://wordpress.com",
		checkerFn: redirectChecker("wordpress.com/typo/?subdomain="),
	},
	Site{
		name: "YouNow",
		// errorMsg:   "No users found",
		// checkBy:    "message",
		rank:       17205,
		profileUrl: "https://www.younow.com/%s/",
		// mainUrl:    "https://www.younow.com/",
		probeUrl:  "https://api.younow.com/php/api/broadcast/info/user=%s/",
		checkerFn: bodyChecker("No users found"),
	},
	Site{
		name: "YouPic",
		// checkBy:    "status_code",
		rank:       43585,
		profileUrl: "https://youpic.com/photographer/%s/",
		// mainUrl:    "https://youpic.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "YouTube",
		// errorMsg: "Not Found",
		// checkBy:    "message",
		rank:       2,
		profileUrl: "https://www.youtube.com/%s",
		// mainUrl:    "https://www.youtube.com/",
		checkerFn: bodyChecker("Not Found"),
	},
	Site{
		name: "Zhihu",
		// checkBy:    "response_url",
		// errorUrl:   "https://www.zhihu.com/people/%s",
		rank:       70,
		profileUrl: "https://www.zhihu.com/people/%s",
		// mainUrl:    "https://www.zhihu.com/",
		checkerFn: redirectChecker("https://www.zhihu.com/people/%s"),
	},
	Site{
		name: "boingboing.net",
		// checkBy:    "status_code",
		rank:       5301,
		profileUrl: "https://bbs.boingboing.net/u/%s",
		// mainUrl:    "https://boingboing.net/",
		checkerFn: statusChecker,
	},
	Site{
		name: "devRant",
		//checkBy:    "response_url",
		//errorUrl:   "https://devrant.com/",
		rank:       159249,
		profileUrl: "https://devrant.com/users/%s",
		// mainUrl:    "https://devrant.com/",
		checkerFn: redirectChecker("https://devrant.com/"),
	},
	Site{
		name: "gfycat",
		// checkBy:    "status_code",
		rank:       230,
		profileUrl: "https://gfycat.com/@%s",
		// mainUrl:    "https://gfycat.com/",
		checkerFn: statusChecker,
	},
	Site{
		name: "iMGSRC.RU",
		// checkBy:    "response_url",
		// errorUrl:   "https://imgsrc.ru/",
		rank:       8262,
		profileUrl: "https://imgsrc.ru/main/user.php?user=%s",
		// mainUrl:    "https://imgsrc.ru/",
		checkerFn: redirectChecker("https://imgsrc.ru/"),
	},
	Site{
		name: "last.fm",
		// checkBy:    "status_code",
		rank:       1242,
		profileUrl: "https://last.fm/user/%s",
		// mainUrl:    "https://last.fm/",
		checkerFn: statusChecker,
	},
	Site{
		name: "mixer.com",
		// checkBy:    "status_code",
		rank:       2145,
		profileUrl: "https://mixer.com/%s",
		// mainUrl:    "https://mixer.com/",
		probeUrl:  "https://mixer.com/api/v1/channels/%s",
		checkerFn: statusChecker,
	},
	Site{
		name: "osu!",
		// checkBy:    "status_code",
		rank:       1813,
		profileUrl: "https://osu.ppy.sh/users/%s",
		// mainUrl:    "https://osu.ppy.sh/",
		checkerFn: statusChecker,
	},
	Site{
		name:       "Cent",
		rank:       1813,
		profileUrl: "https://beta.cent.co/@%s",
		checkerFn:  bodyChecker("<title>Cent</title>"),
	},
}
