package cyoa

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	tt "text/template"
)

var testAdventure string = `{
	"title": "The Little Blue Gopher",
	"story": [
	  "Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
	  "One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
	  "On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
	],
	"options": [
	  {
		"text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
		"arc": "new-york"
	  },
	  {
		"text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
		"arc": "denver"
	  }
	]
  }`

func TestUnmarshalJSON(t *testing.T) {
	jsonData := []byte(testAdventure)
	var a Adventure
	err := json.Unmarshal(jsonData, &a)
	if err != nil {
		t.Error("error decoding desired json into adventure", err)
	}
	if a.Title != "The Little Blue Gopher" {
		t.Errorf("invalid title, should be The Little Blue Gopher, is %v", a.Title)
	}
	if len(a.Story) != 3 {
		t.Errorf("invalid number of paragraphs, should be 3, is %v", len(a.Story))
	}
	if len(a.Options) != 2 {
		t.Errorf("invalid number of options, should be 2, is %v", len(a.Options))
	}
	if a.Options[1].Arc != "denver" {
		t.Errorf("invalid option arc, should be denver, is %v", len(a.Options[1].Arc))
	}
	text := "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York."
	if a.Options[0].Text != text {
		t.Errorf("invalid option text, should be %v, is %v", text, a.Options[1].Text)
	}
}

var testStory string = `{
	"intro": {
	  "title": "The Little Blue Gopher",
	  "story": [
		"Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
		"One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
		"On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
	  ],
	  "options": [
		{
		  "text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
		  "arc": "new-york"
		},
		{
		  "text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
		  "arc": "denver"
		}
	  ]
	},
	"new-york": {
	  "title": "Visiting New York",
	  "story": [
		"Upon arriving in New York you and your furry travel buddy first attempt to hail a cab. Unfortunately nobody wants to give a ride to someone with a \"pet\". They kept saying something about shedding, as if gophers shed.",
		"Unwilling to accept defeat, you pull out your phone and request a ride using <undisclosed-app>. In a few short minutes a car pulls up and the driver helps you load your luggage. He doesn't seem thrilled about your travel companion but he doesn't say anything.",
		"The ride to your hotel is fairly uneventful, with the exception of the driver droning on and on about how he barely breaks even driving around the city and how tips are necessary to make a living. After a while it gets pretty old so you slip in some earbuds and listen to your music.",
		"After arriving at your hotel you check in and walk to the conference center where GothamGo is being held. The friendly man at the desk helped you get your badge and you hurry in to take a seat.",
		"As you head down the aisle you notice a strange man on stage with a mask, cape, and poorly drawn abs on his stomach. Next to him is a man in a... is that a fox outfit? What are these two doing? And what have you gotten yourself into?"
	  ],
	  "options": [
		{
		  "text": "This is getting too weird for me. Let's bail and head back home.",
		  "arc": "home"
		},
		{
		  "text": "Maybe people just dress funny in the big city. Grab a a seat and see what happens.",
		  "arc": "debate"
		}
	  ]
	},
	"debate": {
	  "title": "The Great Debate",
	  "story": [
		"After a bit everyone settles down the two people on stage begin having a debate. You don't recall too many specifics, but for some reason you have a feeling you are supposed to pick sides."
	  ],
	  "options": [
		{
		  "text": "Clearly that man in the fox outfit was the winner.",
		  "arc": "sean-kelly"
		},
		{
		  "text": "I don't think those fake abs would help much in a feat of strength, but our caped friend clearly won this bout. Let's go congratulate him.",
		  "arc": "mark-bates"
		},
		{
		  "text": "Slip out the back before anyone asks us to pick a side.",
		  "arc": "home"
		}
	  ]
	}
}`
var testTemplate string = `============================================================================
{{.Title}}
{{ range .Story }} 
{{ . }}
{{ end }}
{{ range $key, $value := .Options }} 
----------------------------------------------------------------------------
{{ $value.Text }}
press {{ $key }} to choose {{ $value.Arc }}
{{ end }}
`

func Test(t *testing.T) {
	jsonBytes := []byte(testStory)
	var am map[string]Adventure
	err := json.Unmarshal(jsonBytes, &am)
	if err != nil {
		t.Error("error decoding desired json into map of  adventures", err)
	}
	if len(am) != 3 {
		t.Errorf("invalid lenght of adventures map, should be 4, is %v", len(am))
	}
	_, present := am["debate"]
	if !present {
		t.Errorf("debate adventure not present in the map")
	}
}

var intro string = `============================================================================
        The Little Blue Gopher
         
        Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?
         
        One of his friends once recommended going to New York to make friends at this mysterious thing called "GothamGo". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.
         
        On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit.
        
         
        ----------------------------------------------------------------------------
        That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.
        press 0 to choose new-york
         
        ----------------------------------------------------------------------------
        Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.
        press 1 to choose denver`

func TestServeHttp(t *testing.T) {
	//given
	jsonBytes := []byte(testStory)
	var am map[string]Adventure
	err := json.Unmarshal(jsonBytes, &am)
	if err != nil {
		t.Error("error decoding desired json into map of  adventures", err)
	}
	tplPtr := &tt.Template{}
	tplPtr.Parse(testTemplate)

	ah := AdventureHandler{
		AM:       am,
		Template: tplPtr,
	}

	rw := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/intro", nil)

	//when
	ah.ServeHTTP(rw, req)

	//then
	if expected, actual := intro, rw.Body.String(); strings.EqualFold(expected, actual) {
		t.Error("Incorrect response :\n", actual)
	}
}
