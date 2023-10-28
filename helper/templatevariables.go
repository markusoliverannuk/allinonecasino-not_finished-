package helper

import "text/template"

var loginPageTMPLT = template.Must(template.ParseFiles("templates/loginpage.html"))
var homePageTMPLT = template.Must(template.ParseFiles("templates/mainpage.html"))
var registerPageTMPLT = template.Must(template.ParseFiles("templates/registerpage.html"))
var coinflipPageTMPLT = template.Must(template.ParseFiles("templates/coinflip.html"))
var roulettePageTMPLT = template.Must(template.ParseFiles("templates/roulettepage.html"))
var coinfliplobbyPageTMPLT = template.Must(template.ParseFiles("templates/coinfliplobby.html"))
