package acquisition

// Rules compiled from
// https://adrienne.be/ferengi_rules_of_acquisition/
// https://memory-alpha.fandom.com/wiki/Rules_of_Acquisition
// https://memory-beta.fandom.com/wiki/Ferengi_Rules_of_Acquisition
// https://projectsanctuary.com/the_complete_ferengi_rules_of_acquisition.htm
// http://www.sjtrek.com/trek/rules/
// with some edits for spelling and punctuation
var rules = []string{
	"Once you have their money, never give it back.",                                   //1
	"The best deal is the one that brings the most profit.",                            //2
	"Never spend more for an acquisition than you have to.",                            //3
	"A woman wearing clothes is like a man without profit.",                            //4
	"Always exaggerate your estimates.",                                                //5
	"Never let family stand in the way of opportunity.",                                //6
	"Always keep your ears open.",                                                      //7
	"Small print leads to large risk.",                                                 //8
	"Opportunity plus instinct equals profit.",                                         //9
	"Greed is eternal.",                                                                //10
	"Even if it's free, you can always buy it cheaper.",                                //11
	"Anything worth selling is worth selling twice.",                                   //12
	"Anything worth doing is worth doing for money.",                                   //13
	"Sometimes the quickest way to find profits is to let them find you.",              //14
	"Dead men close no deals.",                                                         //15
	"A deal is a deal... until a better one comes along.",                              //16
	"A contract is a contract is a contract... but only between Ferengi.",              //17
	"A Ferengi without profit is no Ferengi at all.",                                   //18
	"Satisfaction is not guaranteed.",                                                  //19
	"He who dives under the table today lives to profit tomorrow.",                     //20
	"Never place friendship before profit.",                                            //21
	"Wise men can hear profit in the wind.",                                            //22
	"Nothing is more important than your health... except for your money",              //23
	"Latinum can't buy happiness, but you can sure have a blast renting it.",           //24
	"You pay for it, it's your idea!",                                                  //25
	"As the customers go, so goes the wise profiteer.",                                 //26
	"There's nothing more dangerous than an honest businessman.",                       //27
	"Morality is always defined by those in power.",                                    //28
	"What's in it for me?",                                                             //29
	"Confidentiality equals profit.",                                                   //30 (alternatively, "A wise man knows that...")
	"Never make fun of a Ferengi's mother.",                                            //31 (alternatively, "...insult something he cares about instead.")
	"Be careful what you sell. It may do exactly what the customer expects.",           //32
	"It never hurts to suck up to the boss.",                                           //33
	"War is good for business.",                                                        //34 (DS9 Reference has these two reversed)
	"Peace is good for business.",                                                      //35
	"Too many Ferengi can't laugh at themselves anymore.",                              //36*
	"The early investor reaps the most interest.",                                      //37
	"Free advertising is cheap.",                                                       //38
	"Don't tell customers more than they need to know.",                                //39
	"She can touch your lobes but never your latinum.",                                 //40
	"Profit is its own reward.",                                                        //41
	"Only negotiate when you are certain to profit.",                                   //42
	"Feed your greed, but not enough to choke it.",                                     //43
	"Never confuse wisdom with luck.",                                                  //44
	"Expand or die.",                                                                   //45
	"Labor camps are full of people who trusted the wrong person.",                     //46*
	"Never trust a man wearing a better suit than you own.",                            //47
	"The bigger the smile, the sharper the knife.",                                     //48
	"Old age and greed will always overcome youth and talent.",                         //49
	"Gratitude can bring on generosity.",                                               //50
	"Never admit a mistake if there's someone else to blame.",                          //51
	"Never ask when you can take.",                                                     //52
	"Never trust anybody taller than you.",                                             //53
	"Rate divided by time equals profit.",                                              //54 (aka The Velocity of Wealth)
	"Take joy from profit, and profit from joy.",                                       //55
	"Pursue profit; women come later.",                                                 //56*
	"Good customers are as rare as latinum. Treasure them.",                            //57
	"There is no substitute for success.",                                              //58
	"Free advice is seldom cheap.",                                                     //59
	"Keep your lies consistent.",                                                       //60
	"Never buy what can be stolen.",                                                    //61*
	"The riskier the road, the greater the profit.",                                    //62
	"Work is the best therapy... at least for your employees.",                         //63
	"Don't talk shop; talk shopping.",                                                  //64*
	"Win or lose, there's always Hupyrian beetle snuff.",                               //65
	"Someone's always got bigger ears.",                                                //66
	"",                                                                                 //67 (ProjectSanctuary has this as "Enough is never enough", however, other sources have that as 97. Self-referential?)
	"Risk doesn't always equal reward.",                                                //68
	"Ferengi are not responsible for the stupidity of other races.",                    //69
	"Get the money first, then let the buyers worry about collecting the merchandise.", //70
	"Gamble and trade have two things in common: risk and latinum.",                    //71
	"Never trust your customers.",                                                      //72
	"If it gets you profit, sell your own mother.",                                     //73
	"Knowledge is latinum.",                                                            //74 (alternatively, "Knowledge equals profit.")
	"Home is where the heart is, but the stars are made of latinum.",                   //75
	"Every once in a while, declare peace. It confuses the hell out of your enemies.",  //76
	"If you break it, I'll charge you for it!",                                         //77
	"Don't discriminate. The most unlikely species can create the best customers.",     //78*
	"Beware the Vulcan greed for knowledge.",                                           //79
	"If it works, sell it. If it works well, sell it for more. If it doesn't work, quadruple the price and sell it as an antique.", //80
	"", //81 (ProjectSanctuary: "There's nothing more dangerous than an honest businessman.", other wources have that as 27)
	"The flimsier the product, the higher the price.",      //82
	"Revenge is profitless.",                               //83
	"A friend is not a friend if he asks for a discount.",  //84
	"Never let the competition know what you're thinking.", //85
	"", //86
	"Learn the customer's weaknesses, so that you can better take advantage of him.", //87
	"It ain't over til its over.", //88 (memory-beta lists two Rules #88, both sourced - the other is "Vengeance will cost you everything.")
	"Ask not what your profits can do for you; ask what you can do for your profits.", //89 (memory-beta again lists two sourced Rules #89 - the other is "[It is] better to lose some profit and live than lose all profit and die.")
	"Mine is better than ours.",                  //90
	"He who drinks fast pays slow.",              //91
	"There are many paths to profit.",            //92
	"He's a fool who makes his doctor his heir.", //93
	"Females and finances don't mix.",            //94
	"Expand or die.",                             //95
	"For every Rule, there is an equal and opposite Rule (except when there's not).", //96
	"Enough is never enough.",                                //97
	"Every man has his price.",                               //98
	"Trust is the biggest liability of all.",                 //99
	"When it's good for business, tell the truth.",           //100
	"Profit trumps emotion.",                                 //101
	"Nature decays, but latinum lasts forever.",              //102
	"Sleep can interfere with opportunity.",                  //103
	"Faith moves mountains... of inventory.",                 //104
	"Don't trust anyone who trusts you.",                     //105
	"There is no honor in poverty.",                          //106
	"",                                                       //107
	"Hope doesn't keep the lights on.",                       //108 (memory-beta once again has two sourced rules #108 - the other is "A woman wearing clothes is like a man without any profits.")
	"Dignity and an empty sack is worth the sack.",           //109
	"Exploitation begins at home.",                           //110
	"Treat people in your debt like family... exploit them.", //111
	"Never have sex with the boss's sister.",                 //112
	"Always have sex with the boss.",                         //113
	"",                                                       //114 - projectsanctuary has this as "Small print leads to large risk."
	"The best contract always has a lot of fine print.",      //115 - projectsanctuary has this as "Greed is eternal"
	"There's always a way out.",                              //116
	"You can't free a fish from water.",                      //117
	"Never cheat an honest man offering a decent price.",     //118
	"Buy, sell, or get out of the way.",                      //119 - adrienne.be says "Never judge a customer by the size of his wallet... sometimes good things come in small packages.",
	"",                                                       //120 - projectsanctuary says "Even a blind man can recognize the glow of latinum", others list that as 123
	"Everything is for sale, even friendship.",               //121
	"", //122 - projectsanctuary says "As the customers go, so goes the wise profiteer.", others list that as 26
	"Even a blind man can recognize the glow of latinum.",                         //123
	"Friendship is temporary, profit is forever.",                                 //124
	"You can't make a deal if you're dead.",                                       //125
	"A lie isn't a lie, it's just the truth seen from a different point of view.", //126
	"Stay neutral in conflicts so that you can sell supplies to both sides.",      //127
	"",                               //128 - projectsanctuary says "Ferengi are not responsible for the stupidity of other races."
	"",                               //129
	"",                               //130
	"",                               //131
	"",                               //132
	"",                               //133
	"",                               //134
	"",                               //135
	"",                               //136
	"",                               //137
	"",                               //138
	"Wives serve; brothers inherit.", //139
	"",                               //140
	"",                               //141
	"",                               //142
	"",                               //143
	"There's nothing wrong with charity... as long as it winds up in your pocket.", //144
	"", //145
	"", //146
	"", //147
	"", //148
	"", //149
	"", //150
	"", //151
	"", //152
	"", //153
	"", //154
	"", //155
	"", //156
	"", //157
	"", //158
	"", //159
	"", //160
	"", //161
	"Even in the worst of times someone turns a profit.", //162
	"",                             //163
	"",                             //164
	"",                             //165
	"",                             //166
	"",                             //167
	"Whisper your way to success.", //168 (adrienne.be says this is #28)
	"",                             //169
	"",                             //170
	"",                             //171
	"",                             //172
	"",                             //173
	"",                             //174
	"",                             //175
	"",                             //176
	"Know your enemies... but do business with them always.", //177
	"", //178
	"", //179
	"", //180
	"Not even dishonesty can tarnish the shine of profit.", //181
	"", //182
	"When life hands you ungaberries, make detergent.",                      //183
	"A Ferengi waits to bid until his opponents have exhausted themselves.", //184
	"", //185
	"", //186
	"", //187
	"", //188
	"Let others keep their reputation. You keep their money.", //189
	"Hear all, trust nothing.",                                //190
	"",                                                        //191
	"Never cheat a Klingon... unless you're sure you can get away with it.", //192
	"Trouble comes in threes.", //193
	"It's always good business to know about new customers before they walk in your door.", //194
	"",                                       //195
	"",                                       //196
	"",                                       //197
	"",                                       //198
	"Location, location, location.",          //199
	"A Ferengi chooses no side but his own.", //200
	"",                                       //201
	`The justification for profit is profit`, //202
	"New customers are like razor toothed gree worms. They can be succulent, but sometimes they bite back.", //203
	"", //204
	"", //205
	"", //206
	"", //207
	"Sometimes, the only thing more dangerous than a question is the answer.", //208
	"", //209
	"", //210
	"Employees are the rungs on the ladder to success. Don't hesitate to step on them.", //211
	"", //212
	"", //213
	"Never begin a negotiation on an empty stomach.", //214
	"",                                //215
	"",                                //216
	"Always know what you're buying.", //217
	"Sometimes what you get free costs entirely too much.", //218
	"", //219
	"", //220
	"", //221
	"", //222
	"Beware the man who doesn't make time for oo-mox.", //223
	"",                                       //224
	"",                                       //225
	"",                                       //226
	"",                                       //227
	"",                                       //228
	"Latinum lasts longer than lust.",        //229
	"",                                       //230
	"",                                       //231
	"",                                       //232
	"",                                       //233
	"",                                       //234
	"",                                       //235
	"You can't buy fate.",                    //236
	"",                                       //237
	"",                                       //238
	"Never be afraid to mislabel a product.", //239
	"",                                       //240
	"",                                       //241
	"More is good, all is better.",           //242
	"",                                       //243
	"",                                       //244
	"",                                       //245
	"",                                       //246
	"Never question luck.",                   //247
	"",                                       //248
	"",                                       //249
	"",                                       //250
	"",                                       //251
	"",                                       //252
	"",                                       //253
	"",                                       //254
	"A wife is a luxury, a smart accountant a necessity.", //255 (projectsanctuary says this is #86)/
	"", //256
	"", //257
	"", //258
	"", //259
	"", //260
	"A wealthy man can afford everything except a conscience.", //261
	"", //262
	"Never allow doubt to tarnish your lust for latinum.", //263
	"",                                   //264
	"",                                   //265
	"When in doubt, lie.",                //266
	"",                                   //267
	"",                                   //268
	"",                                   //269
	"",                                   //270
	"",                                   //271
	"",                                   //272
	"",                                   //273
	"",                                   //274
	"",                                   //275
	"",                                   //276
	"",                                   //277
	"",                                   //278
	"",                                   //279
	"",                                   //280
	"",                                   //281
	"",                                   //282
	"",                                   //283
	"",                                   //284
	"No good deed ever goes unpunished.", //285
	"When Morn leaves, it's all over.",   //286
}
