package src

type Fingerprint struct {
	engine        string
	status        string
	fingerprint   string
	discussion    string
	documentation string
}

func Fingerprints() [61]Fingerprint {

	var fingerprints [61]Fingerprint

	fingerprints[0] = Fingerprint{
		"AWS/S3",
		"Vulnerable",
		"The specified bucket does not exist",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/36",
		"Not available",
	}

	fingerprints[1] = Fingerprint{
		"Bitbucket",
		"Vulnerable",
		"Repository not found",
		"Not available",
		"Not available",
	}

	fingerprints[2] = Fingerprint{
		"Cloudfront",
		"Edge case",
		"Bad Request: ERROR: The request could not be satisfied",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/29",
		"Not available",
	}

	fingerprints[3] = Fingerprint{
		"Desk",
		"Vulnerable",
		"this help center no longer exists",
		"not Available",
		"Not available",
	}

	fingerprints[4] = Fingerprint{
		"Fastly",
		"Edge case",
		"Fastly error: unknown domain:",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/22",
		"Not available",
	}

	fingerprints[5] = Fingerprint{
		"Feedpress",
		"Vulnerable",
		"The feed has not been found.",
		"https://hackerone.com/reports/195350",
		"Not available",
	}

	fingerprints[6] = Fingerprint{
		"Ghost",
		"Vulnerable",
		"The thing you were looking for is no longer here, or never was",
		"",
		"Not available",
	}

	fingerprints[7] = Fingerprint{
		"Github",
		"Vulnerable",
		"There isn't a Github Pages site here",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/37",
		"Not available",
	}

	fingerprints[8] = Fingerprint{
		"Help Juice",
		"Vulnerable",
		"We could not find what you're looking for",
		"Not available",
		"https://help.helpjuice.com/34339-getting-started/custom-domain",
	}

	fingerprints[9] = Fingerprint{
		"Help Scout",
		"Vulnerable",
		"No settings were found for this company",
		"Not available",
		"https://docs.helpscout.net/article/42-setup-custom-domain",
	}

	fingerprints[10] = Fingerprint{
		"Heroku",
		"Edge case",
		"No such app",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/38",
		"Not available",
	}

	fingerprints[11] = Fingerprint{
		"JetBrains",
		"Vulnerable",
		"is not a registered InCloud YouTrack",
		"Not available",
		"Not available",
	}

	fingerprints[12] = Fingerprint{
		"Mashery",
		"Not vulnerable",
		"Unrecognized domain",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/14",
		"Not available",
	}

	fingerprints[13] = Fingerprint{
		"Readme.io",
		"Vulnerable",
		"Project doesnt exist... yet!",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/41",
		"Not available",
	}

	fingerprints[14] = Fingerprint{
		"Shopify",
		"Edge Case",
		"Sorry, this shop is currently unavailable",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/32",
		"https://medium.com/@thebuckhacker/how-to-do-55-000-subdomain-takeover-in-a-blink-of-an-eye-a94954c3fc75",
	}

	fingerprints[15] = Fingerprint{
		"Surge.sh",
		"Vulnerable",
		"project not found",
		"Not available",
		"https://surge.sh/help/adding-a-custom-domain",
	}

	fingerprints[16] = Fingerprint{
		"Tumblr",
		"Vulnerable",
		"Whatever you were looking for doesn't currently exist at this address",
		"Not available",
		"Not available",
	}

	fingerprints[17] = Fingerprint{
		"Tilda",
		"Edge Case",
		"Please renew your subscription",
		"https://github.com/EdOverflow/can-i-take-over-xyz/pull/20",
		"Not available",
	}

	fingerprints[18] = Fingerprint{
		"Unbounce",
		"Not vulnerable",
		"The requested URL was not found on this server",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/11",
		"Not available",
	}

	fingerprints[19] = Fingerprint{
		"UserVoice",
		"Vulnerable",
		"This UserVoice subdomain is currently available",
		"Not available",
		"Not available",
	}

	fingerprints[20] = Fingerprint{
		"Wordpress",
		"Vulnerable",
		"Do you want to register",
		"https://hackerone.com/reports/274336",
		"Not available",
	}

	fingerprints[21] = Fingerprint{
		"Zendesk",
		"Not Vulnerable",
		"Help Center Closed",
		"https://github.com/EdOverflow/can-i-take-over-xyz/issues/23",
		"https://support.zendesk.com/hc/en-us/articles/203664356-Changing-the-address-of-your-Help-Center-subdomain-host-mapping-",
	}
	fingerprints[22] = Fingerprint{
		"Aftership",
		"Vulnerable",
		"Oops.</h2><p class="text-muted text-tight">The page you're looking for doesn't exist.",
		"Not available",
		"Not available",
	}
	fingerprints[23] = Fingerprint{
		"Agile CRM",
		"Vulnerable",
		"Sorry, this page is no longer available",
		"Not available",
		"Not available",
	}
	fingerprints[24] = Fingerprint{
		"Aha",
		"Vulnerable",
		"There is no portal here ... sending you back to Aha!",
		"Not available",
		"Not available",
	}
	fingerprints[25] = Fingerprint{
		"Airee.ru",
		"Vulnerable",
		"Ошибка 402. Сервис Айри.рф не оплачен",
		"Not available",
		"Not available",
	}
	fingerprints[26] = Fingerprint{
		"Anima",
		"Vulnerable",
		"If this is your website and you've just created it, try refreshing in a minute",
		"Not available",
		"Not available",
	}
	fingerprints[27] = Fingerprint{
		"Bigcartel",
		"Vulnerable",
		"<h1>Oops! We couldn&#8217;t find that page.</h1>",
		"Not available",
		"Not available",
	}
	fingerprints[28] = Fingerprint{
		"Brightcove",
		"Vulnerable",
		"<p class="bc-gallery-error-code">Error Code: 404</p>",
		"Not available",
		"Not available",
	}
	fingerprints[29] = Fingerprint{
		"Campaign Monitor",
		"Vulnerable",
		"<strong>Trying to access your account?</strong>",
		"Not available",
		"Not available",
	}
	fingerprints[30] = Fingerprint{
		"Fly.io",
		"Vulnerable",
		"404 Not Found",
		"Not available",
		"Not available",
	}
	fingerprints[31] = Fingerprint{
		"Canny",
		"Vulnerable",
		"Company Not Found",
		"Not available",
		"Not available",
	}
	fingerprints[32] = Fingerprint{
		"Cargo",
		"Vulnerable",
		"If you're moving your domain away from Cargo you must make this configuration through your registrar's DNS control panel.",
		"Not available",
		"Not available",
	}
	fingerprints[33] = Fingerprint{
		"Cargo Collective",
		"Vulnerable",
		"<div class="notfound">",
		"Not available",
		"Not available",
	}
	fingerprints[34] = Fingerprint{
		"Digital Ocean",
		"Vulnerable",
		"Domain uses DO name serves with no records in DO",
		"Not available",
		"Not available",
	}
	fingerprints[35] = Fingerprint{
		"Frontify",
		"Vulnerable",
		"Oops… looks like you got lost",
		"Not available",
		"Not available",
	}
	fingerprints[36] = Fingerprint{
		"Gemfury",
		"Vulnerable",
		"404: This page could not be found.",
		"Not available",
		"Not available",
	}
	fingerprints[37] = Fingerprint{
		"Google Cloud Storage",
		"Vulnerable",
		"NoSuchBucket The specified bucket does not exist.",
		"Not available",
		"Not available",
	}
	fingerprints[38] = Fingerprint{
		"Getresponse",
		"Vulnerable",
		"With GetResponse Landing Pages, lead generation has never been easier",
		"Not available",
		"Not available",
	}fingerprints[39] = Fingerprint{
		"HatenaBlog",
		"Vulnerable",
		"404 Blog is not found",
		"Not available",
		"Not available",
	}fingerprints[40] = Fingerprint{
		"Helprace",
		"Vulnerable",
		"Alias not configured!",
		"Not available",
		"Not available",
	}fingerprints[41] = Fingerprint{
		"Hubspot",
		"Vulnerable",
		"Domain not found",
		"Not available",
		"Not available",
	}
	fingerprints[42] = Fingerprint{
		"Intercom",
		"Vulnerable",
		"This page is reserved for artistic dogs.",
		"Not available",
		"Not available",
	}
	fingerprints[43] = Fingerprint{
		"Jazzhr",
		"Vulnerable",
		"This account no longer active",
		"Not available",
		"Not available",
	}
	fingerprints[44] = Fingerprint{
		"Kinsta",
		"Vulnerable",
		"No Site For Domain",
		"Not available",
		"Not available",
	}
	fingerprints[45] = Fingerprint{
		"LaunchRock",
		"Vulnerable",
		"It looks like you may have taken a wrong turn somewhere. Don't worry...it happens to all of us",
		"Not available",
		"Not available",
	}
	fingerprints[46] = Fingerprint{
		"Landingi",
		"Vulnerable",
		"It looks like you're lost",
		"Not available",
		"Not available",
	}
	fingerprints[47] = Fingerprint{
		"Ngrok",
		"Vulnerable",
		"ngrok.io not found",
		"Not available",
		"Not available",
	}
	fingerprints[48] = Fingerprint{
		"Pantheon",
		"Vulnerable",
		"The gods are wise, but do not know of the site which you seek.",
		"Not available",
		"Not available",
	}
	fingerprints[49] = Fingerprint{
		"Pingdom",
		"Vulnerable",
		"Public Report Not Activated",
		"Not available",
		"Not available",
	}
	fingerprints[50] = Fingerprint{
		"Proposify",
		"Vulnerable",
		"If you need immediate assistance, please contact <a href="mailto:support@proposify.biz	",
		"Not available",
		"Not available",
	}
	fingerprints[51] = Fingerprint{
		"Readthedocs",
		"Vulnerable",
		"unknown to Read the Docs",
		"Not available",
		"Not available",
	}
	fingerprints[52] = Fingerprint{
		"SmartJobBoard",
		"Vulnerable",
		"Job Board Is Unavailable",
		"Not available",
		"Not available",
	}
	fingerprints[53] = Fingerprint{
		"Strikingly",
		"Vulnerable",
		"But if you're looking to build your own website",
		"Not available",
		"Not available",
	}
	fingerprints[54] = Fingerprint{
		"Simplebooklet",
		"Vulnerable",
		"We can't find this <a href="https://simplebooklet.com",
		"Not available",
		"Not available",
	}
	fingerprints[55] = Fingerprint{
		"Smugmug",
		"Vulnerable",
		"Smugmug",
		"Not available",
		"Not available",
	}
	fingerprints[56] = Fingerprint{
		"Surveygizmo",
		"Vulnerable",
		"data-html-name",
		"Not available",
		"Not available",
	}
	fingerprints[57] = Fingerprint{
		"Tave",
		"Vulnerable",
		"<h1>Error 404: Page Not Found</h1>",
		"Not available",
		"Not available",
	}
	fingerprints[58] = Fingerprint{
		"Teamwork",
		"Vulnerable",
		"Oops - We didn't find your site.",
		"Not available",
		"Not available",
	}
	fingerprints[59] = Fingerprint{
		"Tictail",
		"Vulnerable",
		"Building a brand of your own?",
		"Not available",
		"Not available",
	}
	fingerprints[60] = Fingerprint{
		"Wufoo",
		"Vulnerable",
		"Profile not found",
		"Not available",
		"Not available",
	}


	return fingerprints
}
