package main

import (
	"testing"
)

var serfixTests = []struct {
	in  string
	out string
}{
	{
		// Empty string, no escaped quotes
		in:  `s:9:"";`,
		out: `s:0:"";`,
	},
	{
		// Empty string, with escaped quotes
		in:  `s:9:\"\";`,
		out: `s:0:\"\";`,
	},
	{
		// Easy string, no escaped quotes
		in:  `s:0:"0";`,
		out: `s:1:"0";`,
	},
	{
		// Easy string, with escaped quotes
		in:  `s:0:\"0\";`,
		out: `s:1:\"0\";`,
	},
	{
		// Complex string, no escaped quotes
		in:  `s:00:"http://example.com/image.jpg";`,
		out: `s:28:"http://example.com/image.jpg";`,
	},
	{
		// Complex string, with escaped quotes
		in:  `s:00:\"http://example.com/image.jpg\";`,
		out: `s:28:\"http://example.com/image.jpg\";`,
	},
	{
		// Escaped paths
		in:  `s:00:\".*wp-(atom|rdf|rss|rss2|feed|commentsrss2).php$\";`,
		out: `s:47:\".*wp-(atom|rdf|rss|rss2|feed|commentsrss2).php$\";`,
	},
	{
		// Escaped paths
		in:  `s:00:\".*wp-(atom|rdf|rss|rss2|feed|commentsrss2)\\.php$\";`,
		out: `s:48:\".*wp-(atom|rdf|rss|rss2|feed|commentsrss2)\\.php$\";`,
	},
	{
		// Complex string with escape sequences
		in:  `s:000:\"1234 N. Myspace Road, Suburbn, AL 12345  |  (555) 555-5555\r\n<br /><br />\r\nPerson places things ideas are the fun part of all thisby Special People Town Foundation.\";`,
		out: `s:163:\"1234 N. Myspace Road, Suburbn, AL 12345  |  (555) 555-5555\r\n<br /><br />\r\nPerson places things ideas are the fun part of all thisby Special People Town Foundation.\";`,
	},
	{
		// Spanish characters
		in:  `s:000:\"<br /><h2>Nuestro objetivo es servir a todos los niños del distrito escolar primario <br />sin ningún costo. </h2>\";`,
		out: `s:116:\"<br /><h2>Nuestro objetivo es servir a todos los niños del distrito escolar primario <br />sin ningún costo. </h2>\";`,
	},
	{
		// Timestamp
		in:  `s:00:\"Fri, 27 Jun 2014 18:45:31 +0000\";`,
		out: `s:31:\"Fri, 27 Jun 2014 18:45:31 +0000\";`,
	},
	{
		// Just ridiculous
		in: `s:00:\"\n	\n	\n	\n	\n	\n	\n\";`,
		out: `s:13:\"\n	\n	\n	\n	\n	\n	\n\";`,
	},
	{
		// Just ridiculous
		in:  `s:000:\"<div id=\"v-q5P4Vemb-1\" class=\"video-player\">\n</div><br />  <a rel=\"nofollow\" href=\"http://feeds.wordpress.com/1.0/gocomments/wptv.wordpress.com/36028/\"><img alt=\"\" border=\"0\" src=\"http://feeds.wordpress.com/1.0/comments/wptv.wordpress.com/36028/\" /></a> <img alt=\"\" border=\"0\" src=\"http://stats.wordpress.com/b.gif?host=wordpress.tv&blog=5089392&post=36028&subd=wptv&ref=&feed=1\" width=\"1\" height=\"1\" /><div><a href=\"http://wordpress.tv/2014/06/27/carrie-dils-learning-to-troubleshoot-wordpress/\"><img alt=\"Carrie Dils: Learning to Troubleshoot WordPress\" src=\"http://videos.videopress.com/q5P4Vemb/video-55e3804ddb_scruberthumbnail_0.jpg\" width=\"160\" height=\"120\" /></a></div>\";`,
		out: `s:677:\"<div id=\"v-q5P4Vemb-1\" class=\"video-player\">\n</div><br />  <a rel=\"nofollow\" href=\"http://feeds.wordpress.com/1.0/gocomments/wptv.wordpress.com/36028/\"><img alt=\"\" border=\"0\" src=\"http://feeds.wordpress.com/1.0/comments/wptv.wordpress.com/36028/\" /></a> <img alt=\"\" border=\"0\" src=\"http://stats.wordpress.com/b.gif?host=wordpress.tv&blog=5089392&post=36028&subd=wptv&ref=&feed=1\" width=\"1\" height=\"1\" /><div><a href=\"http://wordpress.tv/2014/06/27/carrie-dils-learning-to-troubleshoot-wordpress/\"><img alt=\"Carrie Dils: Learning to Troubleshoot WordPress\" src=\"http://videos.videopress.com/q5P4Vemb/video-55e3804ddb_scruberthumbnail_0.jpg\" width=\"160\" height=\"120\" /></a></div>\";`,
	},
	{
		// Just ridiculous
		in:  `s:0000:\"<p>WordCamp Europe organizers <a href=\"http://2014.europe.wordcamp.org/2014/06/27/ticket-sales-open-for-wordcamp-europe/\" target=\"_blank\">announced</a> today that tickets are now on sale for this year&#8217;s event, which will be held in Sofia, Bulgaria, on September 27th – 29th. Last year&#8217;s event was by all accounts a smashing success and included diverse attendees from around the world. Approximately 70% of those in attendance flew in from outside the Netherlands.</p>\n<p>The organizers expect 900+ attendees this year, which would make it one of the largest WordPress events of the year. Fortunately, they were able to secure the <a href=\"http://www.ndk.bg/\" target=\"_blank\">National Palace of Culture</a> for the venue, the largest multifunctional congress, conference, convention and exhibition center in Southeastern Europe.</p>\n<p><a href=\"http://i0.wp.com/wptavern.com/wp-content/uploads/2014/03/npc.jpg\" rel=\"prettyphoto[25449]\"><img src=\"http://i0.wp.com/wptavern.com/wp-content/uploads/2014/03/npc.jpg?resize=789%2C379\" alt=\"npc\" class=\"aligncenter size-full wp-image-18647\" /></a></p>\n<p>Due to the success of the previous year, companies are rushing to <a href=\"http://2014.europe.wordcamp.org/sponsor-wordcamp-europe/\" target=\"_blank\">sponsor the event</a>, and the packages are even cheaper because of the lower cost of the location. The organizers reported that all the top tier sponsorship packages were sold out within 24 hours last year.</p>\n<p>In a recent <a href=\"http://joshspeaking.com/matt-mullenweg/\" target=\"_blank\">interview</a>, Matt Mullenweg noted that May 2014 marked the first time that non-English downloads of WordPress surpassed the number of English downloads. Although the software was created by English-speaking people, its user base is rapidly expanding to become more representative of the world&#8217;s population. WordCamp Europe is currently one of the few events that demonstrates the true diversity of the community by bringing together a massive multicultural, multilingual group of WordPress users and professionals.</p>\n<p>If you want to connect with the European WordPress community, Sofia is the place to be at the end of September. The deadline for speaker applications is July 5th, 2014. Last year&#8217;s featured presenters included Matt Mullenweg, Joost de Valk, and Vitaly Friedman, along with many other internationally renowned speakers. Tickets for this highly anticipated event cost 30 Euros and 100 tickets were <a href=\"https://twitter.com/WCEurope/status/482503859429183488\" target=\"_blank\">sold within the first hour</a>. If you plan on going, <a href=\"http://2014.europe.wordcamp.org/2014/06/27/ticket-sales-open-for-wordcamp-europe/\" target=\"_blank\">purchase yours</a> as soon as possible; WordCamp Europe is likely to sell out soon.</p>\";`,
		out: `s:2816:\"<p>WordCamp Europe organizers <a href=\"http://2014.europe.wordcamp.org/2014/06/27/ticket-sales-open-for-wordcamp-europe/\" target=\"_blank\">announced</a> today that tickets are now on sale for this year&#8217;s event, which will be held in Sofia, Bulgaria, on September 27th – 29th. Last year&#8217;s event was by all accounts a smashing success and included diverse attendees from around the world. Approximately 70% of those in attendance flew in from outside the Netherlands.</p>\n<p>The organizers expect 900+ attendees this year, which would make it one of the largest WordPress events of the year. Fortunately, they were able to secure the <a href=\"http://www.ndk.bg/\" target=\"_blank\">National Palace of Culture</a> for the venue, the largest multifunctional congress, conference, convention and exhibition center in Southeastern Europe.</p>\n<p><a href=\"http://i0.wp.com/wptavern.com/wp-content/uploads/2014/03/npc.jpg\" rel=\"prettyphoto[25449]\"><img src=\"http://i0.wp.com/wptavern.com/wp-content/uploads/2014/03/npc.jpg?resize=789%2C379\" alt=\"npc\" class=\"aligncenter size-full wp-image-18647\" /></a></p>\n<p>Due to the success of the previous year, companies are rushing to <a href=\"http://2014.europe.wordcamp.org/sponsor-wordcamp-europe/\" target=\"_blank\">sponsor the event</a>, and the packages are even cheaper because of the lower cost of the location. The organizers reported that all the top tier sponsorship packages were sold out within 24 hours last year.</p>\n<p>In a recent <a href=\"http://joshspeaking.com/matt-mullenweg/\" target=\"_blank\">interview</a>, Matt Mullenweg noted that May 2014 marked the first time that non-English downloads of WordPress surpassed the number of English downloads. Although the software was created by English-speaking people, its user base is rapidly expanding to become more representative of the world&#8217;s population. WordCamp Europe is currently one of the few events that demonstrates the true diversity of the community by bringing together a massive multicultural, multilingual group of WordPress users and professionals.</p>\n<p>If you want to connect with the European WordPress community, Sofia is the place to be at the end of September. The deadline for speaker applications is July 5th, 2014. Last year&#8217;s featured presenters included Matt Mullenweg, Joost de Valk, and Vitaly Friedman, along with many other internationally renowned speakers. Tickets for this highly anticipated event cost 30 Euros and 100 tickets were <a href=\"https://twitter.com/WCEurope/status/482503859429183488\" target=\"_blank\">sold within the first hour</a>. If you plan on going, <a href=\"http://2014.europe.wordcamp.org/2014/06/27/ticket-sales-open-for-wordcamp-europe/\" target=\"_blank\">purchase yours</a> as soon as possible; WordCamp Europe is likely to sell out soon.</p>\";`,
	},
}

func TestSerfixCharacterCounts(t *testing.T) {
	for i, test := range serfixTests {
		actual := Replace(test.in)
		if actual != test.out {
			t.Error("Test", i, "Expected", test.out, "got", actual)
		}
	}
}
