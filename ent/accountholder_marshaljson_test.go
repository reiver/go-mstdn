package ent

import (
	"testing"

	"github.com/reiver/go-json"
)

func TestAccountHolder_MarshalJSON(t *testing.T) {

	tests := []struct{
		AccountHolder AccountHolder
		Expected string
	}{
		{
			AccountHolder: AccountHolder{
				json: ``,
			},
			Expected:`{}`,
		},
		{
			AccountHolder: AccountHolder{
				json: `{}`,
			},
			Expected:`{}`,
		},



		{
			AccountHolder: AccountHolder{
				json: `{"once":1,"twice":"two","thrice":"٣","fource":"𐌠𐌡"}`,
			},
			Expected: `{"once":1,"twice":"two","thrice":"٣","fource":"𐌠𐌡"}`,
		},



		{
			AccountHolder: AccountHolder{
				json: `{`+
					`"id":1234567890123`+
					`,`+
					`"username":"jowblow"`+
					`,`+
					`"acct":"joeblow"`+
					`,`+
					`"display_name":"Joe Blow :-)"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"group":false`+
					`,`+
					`"created_at":"1970-01-01T00:00:00.000Z"`+
				`}`,
			},
			Expected: `{`+
				`"id":1234567890123`+
				`,`+
				`"username":"jowblow"`+
				`,`+
				`"acct":"joeblow"`+
				`,`+
				`"display_name":"Joe Blow :-)"`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"discoverable":true`+
				`,`+
				`"group":false`+
				`,`+
				`"created_at":"1970-01-01T00:00:00.000Z"`+
			`}`,
		},



		{
			AccountHolder: AccountHolder{
				json: `{`+
					`"id":"108116990725247731"`+
					`,`+
					`"username":"reiver"`+
					`,`+
					`"acct":"reiver"`+
					`,`+
					`"display_name":"@reiver ⊼ (Charles) :batman:"`+
					`,`+
					`"locked":false`+
					`,`+
					`"bot":false`+
					`,`+
					`"discoverable":true`+
					`,`+
					`"group":false`+
					`,`+
					`"created_at":"2022-04-12T00:00:00.000Z"`+
					`,`+
					`"note":"\u003cp\u003e\u003ca href=\"https://mastodon.social/tags/Fediverse\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFediverse\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/Tech\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eTech\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/Privacy\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ePrivacy\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/SciFiArt\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eSciFiArt\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e\u003cp\u003e\u003ca href=\"https://mastodon.social/tags/FediverseAcademy\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFediverseAcademy\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/FediverseCity\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFediverseCity\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/FingerProtocol\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFingerProtocol\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/GreatApe\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eGreatApe\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/PostFreely\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ePostFreely\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/SpaceHost\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eSpaceHost\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/StarSeed\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eStarSeed\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e\u003cp\u003eOnce upon a time, was a mathematician, computer scientist, data scientist, software engineer, industrial researcher 🌞 \u003c/p\u003e\u003cp\u003eVlogger 🌞\u003c/p\u003e\u003cp\u003eOther foci \u003ca href=\"https://mastodon.social/tags/anthropology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eanthropology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/artificialLife\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eartificialLife\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/biology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ebiology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/egyptology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eegyptology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/evolutionaryAlgorithms\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eevolutionaryAlgorithms\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/geneticGenealogy\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003egeneticGenealogy\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/humanBehavior\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ehumanBehavior\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/mythology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003emythology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/p2p\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ep2p\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/paleoGenetics\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003epaleoGenetics\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/populationGenetics\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003epopulationGenetics\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/strongAI\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003estrongAI\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e\u003cp\u003eThe meaning of life — first survive, then reproduce 🌞\u003c/p\u003e"`+
					`,`+
					`"url":"https://mastodon.social/@reiver"`+
					`,`+
					`"uri":"https://mastodon.social/users/reiver"`+
					`,`+
					`"avatar":"https://files.mastodon.social/accounts/avatars/108/116/990/725/247/731/original/2e097b7812894201.png"`+
					`,`+
					`"avatar_static":"https://files.mastodon.social/accounts/avatars/108/116/990/725/247/731/original/2e097b7812894201.png"`+
					`,`+
					`"header":"https://files.mastodon.social/accounts/headers/108/116/990/725/247/731/original/7713e4efcfcb9c47.png"`+
					`,`+
					`"header_static":"https://files.mastodon.social/accounts/headers/108/116/990/725/247/731/original/7713e4efcfcb9c47.png"`+
					`,`+
					`"followers_count":1342`+
					`,`+
					`"following_count":845`+
					`,`+
					`"statuses_count":11631`+
					`,`+
					`"last_status_at":"2023-09-27"`+
					`,`+
					`"noindex":false`+
					`,`+
					`"emojis":[{"shortcode":"batman","url":"https://files.mastodon.social/custom_emojis/images/000/005/163/original/8iGbkB7aT.png","static_url":"https://files.mastodon.social/custom_emojis/images/000/005/163/static/8iGbkB7aT.png","visible_in_picker":true}]`+
					`,`+
					`"roles":[]`+
					`,`+
					`"fields":[{"name":"Bio","value":"\u003ca href=\"https://changelog.ca/\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\" translate=\"no\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003echangelog.ca/\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e","verified_at":"2023-04-07T03:44:01.580+00:00"},{"name":"Alt (life-casting)","value":"\u003ca href=\"https://firefish.lol/@reiver\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\" translate=\"no\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003efirefish.lol/@reiver\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e","verified_at":null},{"name":"Alt (#SpaceHost)","value":"\u003ca href=\"https://spacehost.team/@charles\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\" translate=\"no\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003espacehost.team/@charles\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e","verified_at":"2023-09-16T07:20:38.250+00:00"},{"name":"Location","value":"Metro Vancouver 🇨🇦","verified_at":null}]`+
				`}`,
			},
			Expected: `{`+
				`"id":"108116990725247731"`+
				`,`+
				`"username":"reiver"`+
				`,`+
				`"acct":"reiver"`+
				`,`+
				`"display_name":"@reiver ⊼ (Charles) :batman:"`+
				`,`+
				`"locked":false`+
				`,`+
				`"bot":false`+
				`,`+
				`"discoverable":true`+
				`,`+
				`"group":false`+
				`,`+
				`"created_at":"2022-04-12T00:00:00.000Z"`+
				`,`+
				`"note":"\u003cp\u003e\u003ca href=\"https://mastodon.social/tags/Fediverse\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFediverse\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/Tech\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eTech\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/Privacy\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ePrivacy\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/SciFiArt\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eSciFiArt\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e\u003cp\u003e\u003ca href=\"https://mastodon.social/tags/FediverseAcademy\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFediverseAcademy\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/FediverseCity\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFediverseCity\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/FingerProtocol\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eFingerProtocol\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/GreatApe\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eGreatApe\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/PostFreely\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ePostFreely\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/SpaceHost\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eSpaceHost\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/StarSeed\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eStarSeed\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e\u003cp\u003eOnce upon a time, was a mathematician, computer scientist, data scientist, software engineer, industrial researcher 🌞 \u003c/p\u003e\u003cp\u003eVlogger 🌞\u003c/p\u003e\u003cp\u003eOther foci \u003ca href=\"https://mastodon.social/tags/anthropology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eanthropology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/artificialLife\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eartificialLife\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/biology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ebiology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/egyptology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eegyptology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/evolutionaryAlgorithms\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003eevolutionaryAlgorithms\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/geneticGenealogy\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003egeneticGenealogy\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/humanBehavior\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ehumanBehavior\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/mythology\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003emythology\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/p2p\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003ep2p\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/paleoGenetics\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003epaleoGenetics\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/populationGenetics\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003epopulationGenetics\u003c/span\u003e\u003c/a\u003e \u003ca href=\"https://mastodon.social/tags/strongAI\" class=\"mention hashtag\" rel=\"tag\"\u003e#\u003cspan\u003estrongAI\u003c/span\u003e\u003c/a\u003e\u003c/p\u003e\u003cp\u003eThe meaning of life — first survive, then reproduce 🌞\u003c/p\u003e"`+
				`,`+
				`"url":"https://mastodon.social/@reiver"`+
				`,`+
				`"uri":"https://mastodon.social/users/reiver"`+
				`,`+
				`"avatar":"https://files.mastodon.social/accounts/avatars/108/116/990/725/247/731/original/2e097b7812894201.png"`+
				`,`+
				`"avatar_static":"https://files.mastodon.social/accounts/avatars/108/116/990/725/247/731/original/2e097b7812894201.png"`+
				`,`+
				`"header":"https://files.mastodon.social/accounts/headers/108/116/990/725/247/731/original/7713e4efcfcb9c47.png"`+
				`,`+
				`"header_static":"https://files.mastodon.social/accounts/headers/108/116/990/725/247/731/original/7713e4efcfcb9c47.png"`+
				`,`+
				`"followers_count":1342`+
				`,`+
				`"following_count":845`+
				`,`+
				`"statuses_count":11631`+
				`,`+
				`"last_status_at":"2023-09-27"`+
				`,`+
				`"noindex":false`+
				`,`+
				`"emojis":[{"shortcode":"batman","url":"https://files.mastodon.social/custom_emojis/images/000/005/163/original/8iGbkB7aT.png","static_url":"https://files.mastodon.social/custom_emojis/images/000/005/163/static/8iGbkB7aT.png","visible_in_picker":true}]`+
				`,`+
				`"roles":[]`+
				`,`+
				`"fields":[{"name":"Bio","value":"\u003ca href=\"https://changelog.ca/\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\" translate=\"no\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003echangelog.ca/\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e","verified_at":"2023-04-07T03:44:01.580+00:00"},{"name":"Alt (life-casting)","value":"\u003ca href=\"https://firefish.lol/@reiver\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\" translate=\"no\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003efirefish.lol/@reiver\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e","verified_at":null},{"name":"Alt (#SpaceHost)","value":"\u003ca href=\"https://spacehost.team/@charles\" target=\"_blank\" rel=\"nofollow noopener noreferrer me\" translate=\"no\"\u003e\u003cspan class=\"invisible\"\u003ehttps://\u003c/span\u003e\u003cspan class=\"\"\u003espacehost.team/@charles\u003c/span\u003e\u003cspan class=\"invisible\"\u003e\u003c/span\u003e\u003c/a\u003e","verified_at":"2023-09-16T07:20:38.250+00:00"},{"name":"Location","value":"Metro Vancouver 🇨🇦","verified_at":null}]`+
			`}`,
		},
	}

	for testNumber, test := range tests {

		actualBytes, err := json.Marshal(test.AccountHolder)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ACCOUNT-HOLDER: %#v", test.AccountHolder)
			continue
		}

		{
			actual := string(actualBytes)
			expected := test.Expected

			if expected != actual {
				t.Errorf("For test #%d, the actual value is not what was expected", testNumber)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				continue
			}
		}
	}
}
