package main

import tb "gopkg.in/tucnak/telebot.v2"

type GroupUser struct {
	User     *tb.User
	Warnings int
}

type BotInfo struct {
	Version float64
	Name    string
	Admins  []string
}
