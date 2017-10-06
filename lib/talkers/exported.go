// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package talkers

import (
    // stdlib
    "log"
    // local
    "../appcontext"
    "../talkers/talkersinterface"
)

var (
    c *appcontext.Context
)

type Talkers struct {}

func New(ac *appcontext.Context) {
    c = ac
    m := &Talkers{}
    c.RegisterTalkersInterface(talkersinterface.TalkersInterface(m))
}

func (t *Talkers) Init() {
    log.Printf("Initializing responders...")
}
