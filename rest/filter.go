package rest

import (
	"errors"
	"strconv"
	"strings"
)

type Filter struct {
	Ads			bool		`yaml:"ads,omitempty" json:"ads,omitempty"`
	Trackers	bool		`yaml:"trackers,omitempty" json:"trackers,omitempty"`
	Malware		bool		`yaml:"malware,omitempty" json:"malware,omitempty"`
	SafeSearch	bool		`yaml:"safeSearch,omitempty" json:"safeSearch,omitempty"`
	PG			int			`yaml:"PG,omitempty" json:"PG,omitempty"`
	Malicious	bool		`yaml:"malicious,omitempty" json:"malicious,omitempty"`
	Risk		string		`yaml:"risk,omitempty" json:"risk,omitempty"`
	Illegal		string		`yaml:"illegal,omitempty" json:"illegal,omitempty"`
	Categories	[]string	`yaml:"categories,omitempty" json:"categories,omitempty"`
}

func (f *Filter) Empty() bool {
	if f.Ads || f.Trackers || f.Malware || f.Malicious || f.SafeSearch { return false }
	if f.PG > 0 { return false }
	if len(f.Categories) > 0 { return false }
	if len(f.Risk) > 0 { return false }
	if len(f.Illegal) > 0 { return false }
	return true
}

func (f *Filter) String() ( pretty string ) {
	if f.Ads				{ pretty += ", ads" }
	if f.Trackers			{ pretty += ", trackers" }
	if f.Malware			{ pretty += ", malware" }
	if f.SafeSearch			{ pretty += ", safe search" }
	if f.PG > 0				{ pretty += ", pg-" + strconv.Itoa( f.PG ) }
	if f.Malicious			{ pretty += ", malicious" }
	if len(f.Risk) > 0		{ pretty += ", " + f.Risk + " risk" }
	if len(f.Illegal) > 0	{ pretty += ", " + f.Illegal }
	if len(f.Categories) > 0 { pretty += "; " + strings.Join( f.Categories, "," ) }
	pretty = strings.TrimPrefix( pretty, ", " )
	return
}

func (f *Filter) Check() error {
	switch f.PG {
		case 0, 12, 18, 21: break
		default: return errors.New( "unsupported PG" )
	}
	switch f.Risk {
		case "", "possible", "medium", "high": break
		default: return errors.New( "unsupported risk level " + f.Risk )
	}
	switch f.Illegal {
		case "", "content", "warez", "spyware", "copyright": break
		default: return errors.New( "bad illegal category " + f.Illegal )
	}
	return nil
}