// Auto-generated by avdl-compiler v1.3.21 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/home.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type HomeScreenItemID string

func (o HomeScreenItemID) DeepCopy() HomeScreenItemID {
	return o
}

type HomeScreenItemType int

const (
	HomeScreenItemType_TODO   HomeScreenItemType = 1
	HomeScreenItemType_PEOPLE HomeScreenItemType = 2
)

func (o HomeScreenItemType) DeepCopy() HomeScreenItemType { return o }

var HomeScreenItemTypeMap = map[string]HomeScreenItemType{
	"TODO":   1,
	"PEOPLE": 2,
}

var HomeScreenItemTypeRevMap = map[HomeScreenItemType]string{
	1: "TODO",
	2: "PEOPLE",
}

func (e HomeScreenItemType) String() string {
	if v, ok := HomeScreenItemTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type HomeScreenItemData struct {
	T__      HomeScreenItemType            `codec:"t" json:"t"`
	Todo__   *HomeScreenTodo               `codec:"todo,omitempty" json:"todo,omitempty"`
	People__ *HomeScreenPeopleNotification `codec:"people,omitempty" json:"people,omitempty"`
}

func (o *HomeScreenItemData) T() (ret HomeScreenItemType, err error) {
	switch o.T__ {
	case HomeScreenItemType_TODO:
		if o.Todo__ == nil {
			err = errors.New("unexpected nil value for Todo__")
			return ret, err
		}
	case HomeScreenItemType_PEOPLE:
		if o.People__ == nil {
			err = errors.New("unexpected nil value for People__")
			return ret, err
		}
	}
	return o.T__, nil
}

func (o HomeScreenItemData) Todo() (res HomeScreenTodo) {
	if o.T__ != HomeScreenItemType_TODO {
		panic("wrong case accessed")
	}
	if o.Todo__ == nil {
		return
	}
	return *o.Todo__
}

func (o HomeScreenItemData) People() (res HomeScreenPeopleNotification) {
	if o.T__ != HomeScreenItemType_PEOPLE {
		panic("wrong case accessed")
	}
	if o.People__ == nil {
		return
	}
	return *o.People__
}

func NewHomeScreenItemDataWithTodo(v HomeScreenTodo) HomeScreenItemData {
	return HomeScreenItemData{
		T__:    HomeScreenItemType_TODO,
		Todo__: &v,
	}
}

func NewHomeScreenItemDataWithPeople(v HomeScreenPeopleNotification) HomeScreenItemData {
	return HomeScreenItemData{
		T__:      HomeScreenItemType_PEOPLE,
		People__: &v,
	}
}

func (o HomeScreenItemData) DeepCopy() HomeScreenItemData {
	return HomeScreenItemData{
		T__: o.T__.DeepCopy(),
		Todo__: (func(x *HomeScreenTodo) *HomeScreenTodo {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Todo__),
		People__: (func(x *HomeScreenPeopleNotification) *HomeScreenPeopleNotification {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.People__),
	}
}

type HomeScreenTodoType int

const (
	HomeScreenTodoType_NONE          HomeScreenTodoType = 0
	HomeScreenTodoType_BIO           HomeScreenTodoType = 1
	HomeScreenTodoType_PROOF         HomeScreenTodoType = 2
	HomeScreenTodoType_DEVICE        HomeScreenTodoType = 3
	HomeScreenTodoType_FOLLOW        HomeScreenTodoType = 4
	HomeScreenTodoType_CHAT          HomeScreenTodoType = 5
	HomeScreenTodoType_PAPERKEY      HomeScreenTodoType = 6
	HomeScreenTodoType_TEAM          HomeScreenTodoType = 7
	HomeScreenTodoType_FOLDER        HomeScreenTodoType = 8
	HomeScreenTodoType_GIT_REPO      HomeScreenTodoType = 9
	HomeScreenTodoType_TEAM_SHOWCASE HomeScreenTodoType = 10
)

func (o HomeScreenTodoType) DeepCopy() HomeScreenTodoType { return o }

var HomeScreenTodoTypeMap = map[string]HomeScreenTodoType{
	"NONE":          0,
	"BIO":           1,
	"PROOF":         2,
	"DEVICE":        3,
	"FOLLOW":        4,
	"CHAT":          5,
	"PAPERKEY":      6,
	"TEAM":          7,
	"FOLDER":        8,
	"GIT_REPO":      9,
	"TEAM_SHOWCASE": 10,
}

var HomeScreenTodoTypeRevMap = map[HomeScreenTodoType]string{
	0:  "NONE",
	1:  "BIO",
	2:  "PROOF",
	3:  "DEVICE",
	4:  "FOLLOW",
	5:  "CHAT",
	6:  "PAPERKEY",
	7:  "TEAM",
	8:  "FOLDER",
	9:  "GIT_REPO",
	10: "TEAM_SHOWCASE",
}

func (e HomeScreenTodoType) String() string {
	if v, ok := HomeScreenTodoTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type HomeScreenTodo struct {
	T__ HomeScreenTodoType `codec:"t" json:"t"`
}

func (o *HomeScreenTodo) T() (ret HomeScreenTodoType, err error) {
	switch o.T__ {
	}
	return o.T__, nil
}

func NewHomeScreenTodoDefault(t HomeScreenTodoType) HomeScreenTodo {
	return HomeScreenTodo{
		T__: t,
	}
}

func (o HomeScreenTodo) DeepCopy() HomeScreenTodo {
	return HomeScreenTodo{
		T__: o.T__.DeepCopy(),
	}
}

type HomeScreenPeopleNotificationType int

const (
	HomeScreenPeopleNotificationType_FOLLOWED       HomeScreenPeopleNotificationType = 0
	HomeScreenPeopleNotificationType_FOLLOWED_MULTI HomeScreenPeopleNotificationType = 1
)

func (o HomeScreenPeopleNotificationType) DeepCopy() HomeScreenPeopleNotificationType { return o }

var HomeScreenPeopleNotificationTypeMap = map[string]HomeScreenPeopleNotificationType{
	"FOLLOWED":       0,
	"FOLLOWED_MULTI": 1,
}

var HomeScreenPeopleNotificationTypeRevMap = map[HomeScreenPeopleNotificationType]string{
	0: "FOLLOWED",
	1: "FOLLOWED_MULTI",
}

func (e HomeScreenPeopleNotificationType) String() string {
	if v, ok := HomeScreenPeopleNotificationTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type HomeScreenPeopleNotificationFollowed struct {
	FollowTime Time        `codec:"followTime" json:"followTime"`
	User       UserSummary `codec:"user" json:"user"`
}

func (o HomeScreenPeopleNotificationFollowed) DeepCopy() HomeScreenPeopleNotificationFollowed {
	return HomeScreenPeopleNotificationFollowed{
		FollowTime: o.FollowTime.DeepCopy(),
		User:       o.User.DeepCopy(),
	}
}

type HomeScreenPeopleNotificationFollowedMulti struct {
	Followers []HomeScreenPeopleNotificationFollowed `codec:"followers" json:"followers"`
	NumOthers int                                    `codec:"numOthers" json:"numOthers"`
}

func (o HomeScreenPeopleNotificationFollowedMulti) DeepCopy() HomeScreenPeopleNotificationFollowedMulti {
	return HomeScreenPeopleNotificationFollowedMulti{
		Followers: (func(x []HomeScreenPeopleNotificationFollowed) []HomeScreenPeopleNotificationFollowed {
			if x == nil {
				return nil
			}
			var ret []HomeScreenPeopleNotificationFollowed
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Followers),
		NumOthers: o.NumOthers,
	}
}

type HomeScreenPeopleNotification struct {
	T__             HomeScreenPeopleNotificationType           `codec:"t" json:"t"`
	Followed__      *HomeScreenPeopleNotificationFollowed      `codec:"followed,omitempty" json:"followed,omitempty"`
	FollowedMulti__ *HomeScreenPeopleNotificationFollowedMulti `codec:"followedMulti,omitempty" json:"followedMulti,omitempty"`
}

func (o *HomeScreenPeopleNotification) T() (ret HomeScreenPeopleNotificationType, err error) {
	switch o.T__ {
	case HomeScreenPeopleNotificationType_FOLLOWED:
		if o.Followed__ == nil {
			err = errors.New("unexpected nil value for Followed__")
			return ret, err
		}
	case HomeScreenPeopleNotificationType_FOLLOWED_MULTI:
		if o.FollowedMulti__ == nil {
			err = errors.New("unexpected nil value for FollowedMulti__")
			return ret, err
		}
	}
	return o.T__, nil
}

func (o HomeScreenPeopleNotification) Followed() (res HomeScreenPeopleNotificationFollowed) {
	if o.T__ != HomeScreenPeopleNotificationType_FOLLOWED {
		panic("wrong case accessed")
	}
	if o.Followed__ == nil {
		return
	}
	return *o.Followed__
}

func (o HomeScreenPeopleNotification) FollowedMulti() (res HomeScreenPeopleNotificationFollowedMulti) {
	if o.T__ != HomeScreenPeopleNotificationType_FOLLOWED_MULTI {
		panic("wrong case accessed")
	}
	if o.FollowedMulti__ == nil {
		return
	}
	return *o.FollowedMulti__
}

func NewHomeScreenPeopleNotificationWithFollowed(v HomeScreenPeopleNotificationFollowed) HomeScreenPeopleNotification {
	return HomeScreenPeopleNotification{
		T__:        HomeScreenPeopleNotificationType_FOLLOWED,
		Followed__: &v,
	}
}

func NewHomeScreenPeopleNotificationWithFollowedMulti(v HomeScreenPeopleNotificationFollowedMulti) HomeScreenPeopleNotification {
	return HomeScreenPeopleNotification{
		T__:             HomeScreenPeopleNotificationType_FOLLOWED_MULTI,
		FollowedMulti__: &v,
	}
}

func (o HomeScreenPeopleNotification) DeepCopy() HomeScreenPeopleNotification {
	return HomeScreenPeopleNotification{
		T__: o.T__.DeepCopy(),
		Followed__: (func(x *HomeScreenPeopleNotificationFollowed) *HomeScreenPeopleNotificationFollowed {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Followed__),
		FollowedMulti__: (func(x *HomeScreenPeopleNotificationFollowedMulti) *HomeScreenPeopleNotificationFollowedMulti {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.FollowedMulti__),
	}
}

type HomeScreenItem struct {
	Badged bool               `codec:"badged" json:"badged"`
	Data   HomeScreenItemData `codec:"data" json:"data"`
}

func (o HomeScreenItem) DeepCopy() HomeScreenItem {
	return HomeScreenItem{
		Badged: o.Badged,
		Data:   o.Data.DeepCopy(),
	}
}

type HomeScreen struct {
	LastViewed        Time             `codec:"lastViewed" json:"lastViewed"`
	Version           int              `codec:"version" json:"version"`
	Items             []HomeScreenItem `codec:"items" json:"items"`
	FollowSuggestions []UserSummary    `codec:"followSuggestions" json:"followSuggestions"`
}

func (o HomeScreen) DeepCopy() HomeScreen {
	return HomeScreen{
		LastViewed: o.LastViewed.DeepCopy(),
		Version:    o.Version,
		Items: (func(x []HomeScreenItem) []HomeScreenItem {
			if x == nil {
				return nil
			}
			var ret []HomeScreenItem
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Items),
		FollowSuggestions: (func(x []UserSummary) []UserSummary {
			if x == nil {
				return nil
			}
			var ret []UserSummary
			for _, v := range x {
				vCopy := v.DeepCopy()
				ret = append(ret, vCopy)
			}
			return ret
		})(o.FollowSuggestions),
	}
}

type HomeGetScreenArg struct {
	MarkViewed bool `codec:"markViewed" json:"markViewed"`
}

type HomeSkipTodoTypeArg struct {
	T HomeScreenTodoType `codec:"t" json:"t"`
}

type HomeActionTakenArg struct {
}

type HomeMarkViewedArg struct {
}

type HomeInterface interface {
	HomeGetScreen(context.Context, bool) (HomeScreen, error)
	HomeSkipTodoType(context.Context, HomeScreenTodoType) error
	HomeActionTaken(context.Context) error
	HomeMarkViewed(context.Context) error
}

func HomeProtocol(i HomeInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.home",
		Methods: map[string]rpc.ServeHandlerDescription{
			"homeGetScreen": {
				MakeArg: func() interface{} {
					ret := make([]HomeGetScreenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]HomeGetScreenArg)
					if !ok {
						err = rpc.NewTypeError((*[]HomeGetScreenArg)(nil), args)
						return
					}
					ret, err = i.HomeGetScreen(ctx, (*typedArgs)[0].MarkViewed)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"homeSkipTodoType": {
				MakeArg: func() interface{} {
					ret := make([]HomeSkipTodoTypeArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]HomeSkipTodoTypeArg)
					if !ok {
						err = rpc.NewTypeError((*[]HomeSkipTodoTypeArg)(nil), args)
						return
					}
					err = i.HomeSkipTodoType(ctx, (*typedArgs)[0].T)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"homeActionTaken": {
				MakeArg: func() interface{} {
					ret := make([]HomeActionTakenArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.HomeActionTaken(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"homeMarkViewed": {
				MakeArg: func() interface{} {
					ret := make([]HomeMarkViewedArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					err = i.HomeMarkViewed(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type HomeClient struct {
	Cli rpc.GenericClient
}

func (c HomeClient) HomeGetScreen(ctx context.Context, markViewed bool) (res HomeScreen, err error) {
	__arg := HomeGetScreenArg{MarkViewed: markViewed}
	err = c.Cli.Call(ctx, "keybase.1.home.homeGetScreen", []interface{}{__arg}, &res)
	return
}

func (c HomeClient) HomeSkipTodoType(ctx context.Context, t HomeScreenTodoType) (err error) {
	__arg := HomeSkipTodoTypeArg{T: t}
	err = c.Cli.Call(ctx, "keybase.1.home.homeSkipTodoType", []interface{}{__arg}, nil)
	return
}

func (c HomeClient) HomeActionTaken(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.home.homeActionTaken", []interface{}{HomeActionTakenArg{}}, nil)
	return
}

func (c HomeClient) HomeMarkViewed(ctx context.Context) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.home.homeMarkViewed", []interface{}{HomeMarkViewedArg{}}, nil)
	return
}
