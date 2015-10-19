package datatypes

import (
	"strings"

	log "github.com/avabot/ava/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

// StructuredInput is generated by Ava and sent to packages. The UserId is
// guaranteed to be unique. The FlexId is used for UserId lookups to maintain
// context, such as a phone number or email address.
type StructuredInput struct {
	Commands StringSlice
	Actors   StringSlice
	Objects  StringSlice
	Times    StringSlice
	Places   StringSlice
}

type WordClass struct {
	Word  string
	Class int
}

func (si *StructuredInput) String() string {
	s := "\n"
	if len(si.Commands) > 0 {
		s += "Command: " + strings.Join(si.Commands, ", ") + "\n"
	}
	if len(si.Actors) > 0 {
		s += "Actors: " + strings.Join(si.Actors, ", ") + "\n"
	}
	if len(si.Objects) > 0 {
		s += "Objects: " + strings.Join(si.Objects, ", ") + "\n"
	}
	if len(si.Times) > 0 {
		s += "Times: " + strings.Join(si.Times, ", ") + "\n"
	}
	if len(si.Places) > 0 {
		s += "Places: " + strings.Join(si.Places, ", ") + "\n"
	}
	return s[:len(s)-1] + "\n"
}

func (si *StructuredInput) All() string {
	var s string
	if len(si.Commands) > 0 {
		s += strings.Join(si.Commands, " ") + " "
	}
	if len(si.Actors) > 0 {
		s += strings.Join(si.Actors, " ") + " "
	}
	if len(si.Objects) > 0 {
		s += strings.Join(si.Objects, " ") + " "
	}
	if len(si.Times) > 0 {
		s += strings.Join(si.Times, " ") + " "
	}
	if len(si.Places) > 0 {
		s += strings.Join(si.Places, " ") + " "
	}
	return s[0 : len(s)-1]
}

// Add pairs of words with their classes to a structured input. Params should
// follow the ("Order", "Command"), ("noon", "Time") form.
func (si *StructuredInput) Add(wc []WordClass) error {
	if len(wc) == 0 {
		return ErrInvalidOddParameter
	}
	for _, w := range wc {
		switch w.Class {
		case CommandI:
			si.Commands = append(si.Commands, w.Word)
		case ActorI:
			si.Actors = append(si.Actors, w.Word)
		case ObjectI:
			si.Objects = append(si.Objects, w.Word)
		case TimeI:
			si.Times = append(si.Times, w.Word)
		case PlaceI:
			si.Places = append(si.Places, w.Word)
		case NoneI:
			// Do nothing
		default:
			log.Error("invalid class: ", w.Class)
			return ErrInvalidClass
		}
	}
	return nil
}

// TODO Optimize by passing back a struct with []string AND int (ActorI,
// ObjectI, etc.)
func (si *StructuredInput) Pronouns() []string {
	p := []string{}
	for _, w := range si.Objects {
		if Pronouns[w] != 0 {
			p = append(p, w)
		}
	}
	for _, w := range si.Actors {
		if Pronouns[w] != 0 {
			p = append(p, w)
		}
	}
	for _, w := range si.Times {
		if Pronouns[w] != 0 {
			p = append(p, w)
		}
	}
	for _, w := range si.Places {
		if Pronouns[w] != 0 {
			p = append(p, w)
		}
	}
	return p
}
