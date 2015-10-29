package shared

import (
  "github.com/thoj/go-ircevent"
)

func Setup()  {
  ircobj := irc.IRC("<nick>", "<user>") //Create new ircobj
//Set options
ircobj.UseTLS = true //default is false
//ircobj.TLSOptions //set ssl options
ircobj.Password = "[server password]"
//Commands
ircobj.Connect("irc.someserver.com:6667") //Connect to server
ircobj.SendRaw("<string>") //sends string to server. Adds \r\n
ircobj.SendRawf("<formatstring>", ...) //sends formatted string to server.n
ircobj.Join("<#channel> [password]")
ircobj.Nick("newnick")
ircobj.Privmsg("<nickname | #channel>", "msg")
ircobj.Privmsgf(<nickname | #channel>, "<formatstring>", ...)
ircobj.Notice("<nickname | #channel>", "msg")
ircobj.Noticef("<nickname | #channel>", "<formatstring>", ...)
}
