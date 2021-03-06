
# optwdl: Helpful word sets for solving Wordle puzzles.
Wordle gives you 6 tries to identify a 5 letter word. I decided to look for sets of 5 words that cover as many letters as possible. So I wrote a Go program to search for them.

I've yet to find a set that covers 25 letters, but the program currently returns 607 valid word combos that cover 24 letters. The output for a run using the word list (178681 words) included with the program is shown below.

For those who care about such things, the run took 72 seconds on my Mac Core i5 laptop.  It could surely be further optimized.

Choose a set that tickles your fancy for easier Wordle-ing.

Enjoy!
## Output using AllWords list from `dict.go`
```
542 'a' candidates found
313 'e' candidates found
354 'i' candidates found
345 'o' candidates found
275 'u' candidates found
1: BRUGH, VENDS, FLOCK, JIMPY, WALTZ
2: BRUSK, FJELD, CHOMP, VYING, WALTZ
3: BUCKS, FJELD, MORPH, VYING, WALTZ
4: BUCKS, FJELD, ROWTH, ZINGY, VAMPS
5: BUCKS, FJELD, ROWTH, ZINGY, VAMPY
6: BUCKS, FJELD, THROW, ZINGY, VAMPS
7: BUCKS, FJELD, THROW, ZINGY, VAMPY
8: BUCKS, FJELD, WHOMP, RITZY, VANGS
9: BUCKS, FJELD, WHOMP, VYING, TRANQ
10: BUCKS, FJELD, WHOMP, VYING, TZARS
11: BUCKS, FJELD, WHOMP, ZINGY, TRANQ
12: BUCKS, FJELD, WHORT, ZINGY, VAMPS
13: BUCKS, FJELD, WHORT, ZINGY, VAMPY
14: BUCKS, FJELD, WORTH, ZINGY, VAMPS
15: BUCKS, FJELD, WORTH, ZINGY, VAMPY
16: BUCKS, FJELD, WROTH, ZINGY, VAMPS
17: BUCKS, FJELD, WROTH, ZINGY, VAMPY
18: BUCKS, FREMD, PHLOX, VYING, WALTZ
19: BUCKS, GYVED, WHOMP, FRITZ, XYLAN
20: BUCKS, XYLEM, FJORD, WHIPT, VANGS
21: BULGY, VENDS, WHOMP, FRITZ, JACKS
22: BULGY, VENDS, WHOMP, FRITZ, JACKY
23: BULKS, VETCH, FJORD, ZINGY, SWAMP
24: BULKS, WECHT, FJORD, ZINGY, VAMPS
25: BULKS, WECHT, FJORD, ZINGY, VAMPY
26: BULKY, VETCH, FJORD, ZINGS, SWAMP
27: BULKY, WECHT, FJORD, ZINGS, VAMPS
28: BULKY, WECHT, FJORD, ZINGS, VAMPY
29: BUMFS, DRECK, PHLOX, VYING, WALTZ
30: BUMFS, GYVED, PHLOX, WRICK, JANTY
31: BUMFS, GYVED, PHLOX, WRICK, TRANQ
32: BUMFS, WRECK, PHLOX, DITZY, VANGS
33: BUMPH, DECKS, JOWLY, FRITZ, VANGS
34: BUMPH, DELFT, VROWS, ZINGY, JACKS
35: BUMPH, DELFT, VROWS, ZINGY, JACKY
36: BUMPH, DWELT, JOCKS, VYING, ZARFS
37: BUMPH, FJELD, CORKS, VYING, WALTZ
38: BUMPH, FJELD, GOWKS, ZINCY, TRANQ
39: BUMPH, FJELD, ROCKS, VYING, WALTZ
40: BUMPH, FJELD, VROWS, ZINGY, STACK
41: BUMPH, FJELD, VROWS, ZINGY, TACKS
42: BUMPH, FJELD, VROWS, ZINGY, TACKY
43: BUMPH, FJELD, VROWS, ZINGY, THACK
44: BUMPH, FJELD, VROWS, ZINGY, TRACK
45: BUMPH, GECKS, FJORD, VINYL, WALTZ
46: BUMPH, GYVED, CLOWN, FRITZ, JACKS
47: BUMPH, GYVED, CROFT, JINKS, WALTZ
48: BUMPH, GYVED, JOCKS, FRITZ, BLAWN
49: BUMPH, GYVED, JOCKS, FRITZ, LAWNS
50: BUMPH, GYVED, JOCKS, FRITZ, LAWNY
51: BUMPH, GYVED, JOCKS, FRITZ, WANLY
52: BUMPH, GYVED, JOCKS, FRITZ, XYLAN
53: BUMPH, GYVED, JOWLS, FRITZ, CLANK
54: BUMPH, GYVED, JOWLS, FRITZ, CRANK
55: BUMPH, GYVED, JOWLS, FRITZ, SNACK
56: BUMPH, GYVED, KNOWS, FRITZ, CALYX
57: BUMPH, GYVED, WONKS, FRITZ, CALYX
58: BUMPH, GYVES, FJORD, CLINK, WALTZ
59: BUMPH, NERVY, JOCKS, GLITZ, DWARF
60: BUMPH, VELDT, FROWS, ZINGY, JACKS
61: BUMPH, VELDT, FROWS, ZINGY, JACKY
62: BUMPH, VELDT, JOCKS, WINGY, ZARFS
63: BUMPH, VELDT, JOCKS, WRING, FLAXY
64: BUMPH, VELDT, JOCKS, ZINGY, DWARF
65: BUMPH, VELDT, JOCKS, ZINGY, SWARF
66: BUMPH, VELDT, JOCKS, ZINGY, WHARF
67: BUMPH, VENTS, FJORD, CYLIX, GAWKS
68: BUMPH, VENTS, FJORD, CYLIX, GAWKY
69: BUMPH, WEDGY, JOCKS, FRITZ, XYLAN
70: BUMPH, WYLED, JOCKS, FRITZ, VANGS
71: BUMPS, CHEVY, FJORD, GLITZ, SWANK
72: BUMPS, CHEVY, FJORD, GLITZ, WANKS
73: BUMPS, KETCH, FJORD, VYING, WALTZ
74: BUMPS, VETCH, FJORD, KYLIX, GNAWS
75: BUMPS, VETCH, FJORD, KYLIX, SWANG
76: BUMPS, VETCH, FJORD, KYLIX, TWANG
77: BUMPS, VETCH, FJORD, KYLIX, WHANG
78: BUMPS, VETCH, FJORD, KYLIX, WRANG
79: BUMPS, VETCH, FJORD, ZINGY, WALKS
80: BUMPS, WECHT, FJORD, KYLIX, VANGS
81: BUMPY, FJELD, ZONKS, WITCH, GRAVY
82: BUMPY, VENTS, FJORD, ZILCH, GAWKS
83: BUMPY, VENTS, FJORD, ZILCH, GAWKY
84: BUMPY, VETCH, FJORD, GINKS, WALTZ
85: BUMPY, VETCH, FJORD, KINGS, WALTZ
86: BUMPY, VETCH, FJORD, SWINK, GLAZY
87: BUMPY, VETCH, FJORD, WINKS, GLAZY
88: BUMPY, VETCH, FJORD, ZINGS, WALKS
89: BUNCH, FJELD, GOWKS, RITZY, VAMPS
90: BUNCH, FJELD, GOWKS, RITZY, VAMPY
91: BUNCH, VELDT, GOWKS, JIMPY, ZARFS
92: BUNCH, XYLEM, FJORD, SPITZ, GAWKS
93: BUNCH, XYLEM, FJORD, SPITZ, GAWKY
94: BUNGS, KEMPT, FJORD, CHIVY, WALTZ
95: BUNGS, KEMPT, FJORD, VICHY, WALTZ
96: BUNGS, VETCH, FJORD, KYLIX, SWAMP
97: BUNGS, WECHT, FJORD, KYLIX, VAMPS
98: BUNGS, WECHT, FJORD, KYLIX, VAMPY
99: BUNKS, CHEVY, FJORD, GLITZ, SWAMP
100: BUNKS, CHEWY, FJORD, GLITZ, VAMPS
101: BUNKS, CHEWY, FJORD, GLITZ, VAMPY
102: BUNKS, GYVED, WHOMP, FRITZ, CALYX
103: BUNKS, VETCH, FJORD, GIMPY, WALTZ
104: BUNKS, VETCH, FJORD, PIGMY, WALTZ
105: BUNKS, VETCH, FJORD, WIMPY, GLAZY
106: BURGH, VENDS, FLOCK, JIMPY, WALTZ
107: BUTCH, FJELD, GROWS, ZINKY, VAMPS
108: BUTCH, FJELD, GROWS, ZINKY, VAMPY
109: BUTCH, FJELD, VROWS, ZINKY, GAMPS
110: BUTCH, FJELD, VROWS, ZINKY, GRAMP
111: BUTCH, FJELD, WORKS, ZINGY, VAMPS
112: BUTCH, FJELD, WORKS, ZINGY, VAMPY
113: BUTCH, FJELD, ZONKS, WIMPY, GRAVY
114: BUTCH, KEMPS, FJORD, VYING, WALTZ
115: BUTCH, XYLEM, FJORD, ZINGS, PAWKY
116: CHUBS, KEMPT, FJORD, VYING, WALTZ
117: CHUMP, FJELD, BORKS, VYING, WALTZ
118: CHUMP, FJELD, BORTZ, SWINK, GRAVY
119: CHUMP, FJELD, BORTZ, VYING, DAWKS
120: CHUMP, FJELD, BORTZ, VYING, GAWKS
121: CHUMP, FJELD, BORTZ, VYING, HAWKS
122: CHUMP, FJELD, BORTZ, VYING, SWANK
123: CHUMP, FJELD, BORTZ, VYING, WACKS
124: CHUMP, FJELD, BORTZ, VYING, WALKS
125: CHUMP, FJELD, BORTZ, VYING, WANKS
126: CHUMP, FJELD, BORTZ, VYING, WARKS
127: CHUMP, FJELD, BORTZ, WINKS, GRAVY
128: CHUMP, GYVED, BLOWN, FRITZ, JACKS
129: CHUMP, GYVED, BORTZ, JINKS, FLAWS
130: CHUMP, GYVED, BORTZ, JINKS, FLAWY
131: CHUMP, GYVED, BORTZ, JINKS, FLAXY
132: CHUMP, GYVED, BORTZ, SWINK, FLAXY
133: CHUMP, GYVED, BORTZ, WINKS, FLAXY
134: CHUMP, GYVED, FROWN, BLITZ, JACKS
135: CHUMP, GYVED, JOWLS, FRITZ, BANKS
136: CHUMP, GYVED, JOWLS, FRITZ, BLANK
137: CHUMP, GYVED, JOWLS, FRITZ, BRANK
138: CHUMP, GYVES, FJORD, BLINK, WALTZ
139: CHUMP, GYVES, FJORD, BLITZ, SWANK
140: CHUMP, GYVES, FJORD, BLITZ, WANKS
141: CHUMP, WEDGY, BORTZ, JINKS, FLAXY
142: CHUMS, FJELD, BORTZ, VYING, PAWKY
143: CHUNK, FJELD, BORTZ, WIMPS, GRAVY
144: CHUNK, FJELD, BORTZ, WIMPY, VANGS
145: CHUNK, FJELD, WOMBY, SPITZ, GRAVY
146: CHUNK, GYBED, JOWLS, FRITZ, VAMPS
147: CHUNK, GYBED, JOWLS, FRITZ, VAMPY
148: CHUNK, GYVED, BORTZ, WIMPS, FLAXY
149: CHUNK, GYVED, PLOWS, FRITZ, JAMBS
150: CHUNK, GYVES, FJORD, BLIMP, WALTZ
151: CHUNK, GYVES, FJORD, BLITZ, SWAMP
152: CLUMP, GYVED, BORTZ, JINKS, WHARF
153: CLUNK, GYVED, QOPHS, FRITZ, JAMBS
154: CLUNK, GYVED, WHOMP, FRITZ, JAMBS
155: CLUNK, GYVED, WHOPS, FRITZ, JAMBS
156: CLUNK, WEDGY, QOPHS, FRITZ, JAMBS
157: CRUMB, FJELD, KOPHS, VYING, WALTZ
158: CRUMB, FJELD, QOPHS, VYING, WALTZ
159: CRUMB, FJELD, QOPHS, ZINKY, TWANG
160: CRUMB, FJELD, ZONKS, WHIPT, GRAVY
161: CRUMB, FJELD, ZONKS, WIGHT, VAMPY
162: CRUMB, GYVED, PHLOX, FINKS, WALTZ
163: CRUMB, GYVED, PHLOX, JINKS, WAFTS
164: CRUMB, GYVED, PHLOX, JINKS, WALTZ
165: CRUMB, GYVED, PLOTZ, JINKS, WHARF
166: CRUMB, GYVED, ZONKS, WHIPT, FLAXY
167: CRUMP, FJELD, ZONKS, WIGHT, BRAXY
168: CURVY, FJELD, ZONKS, BIGHT, SWAMP
169: CURVY, FJELD, ZONKS, WHIPT, GAMBS
170: FLUMP, CHEVY, BORTZ, DJINS, GAWKS
171: FLUMP, CHEVY, BORTZ, DJINS, GAWKY
172: FLUMP, GYVED, BORTZ, JINKS, CHAWS
173: FLUMP, GYVED, BORTZ, JINKS, SCHWA
174: FLUMP, GYVED, BORTZ, JINKS, WATCH
175: FLUMP, GYVED, BORTZ, JINKS, WHACK
176: FLUMP, GYVED, BORTZ, WHINS, JACKS
177: FLUMP, GYVED, BORTZ, WHINS, JACKY
178: FLUMP, GYVED, BORTZ, WINCH, JACKS
179: FLUMP, GYVED, BROWN, SIXTH, JACKS
180: FLUMP, GYVED, BROWN, SIXTH, JACKY
181: FLUMP, GYVED, JOHNS, BRICK, WALTZ
182: FLUMP, GYVED, ZONKS, WITCH, BRAXY
183: FLUMP, HERTZ, JOCKS, VYING, BAWDS
184: FLUMP, HERTZ, JOCKS, VYING, BAWDY
185: FLUMP, VETCH, GOWDS, ZINKY, BRAXY
186: FLUMP, WEDGY, BORTZ, JINKS, SCHAV
187: FLUNG, VERBS, WHOMP, DITZY, JACKS
188: FLUNG, VERBS, WHOMP, DITZY, JACKY
189: FLUNG, WRECK, QOPHS, DITZY, JAMBS
190: FLUNK, VETCH, GOWDS, JIMPY, BRAXY
191: FLUNK, VETCH, GROSZ, JIMPY, BAWDS
192: FLUNK, VETCH, GROSZ, JIMPY, BAWDY
193: FRUMP, GYVED, BLOWN, SIXTH, JACKS
194: FRUMP, GYVED, BLOWN, SIXTH, JACKY
195: FRUMP, GYVED, BOTCH, JINKS, WALTZ
196: FRUMP, GYVED, JOCKS, BLITZ, SHAWN
197: FRUMP, GYVED, JOCKS, BLITZ, WHANG
198: FRUMP, GYVED, JOHNS, BLITZ, WACKS
199: FRUMP, GYVED, JOHNS, BLITZ, WACKY
200: FRUMP, GYVED, JOHNS, BLITZ, WHACK
201: FRUMP, GYVED, JOHNS, BLITZ, WRACK
202: FRUMP, GYVED, SHOWN, BLITZ, JACKS
203: FRUMP, GYVED, SHOWN, BLITZ, JACKY
204: FUCKS, GYVED, WHOMP, BLITZ, TRANQ
205: FUCKS, VELDT, WHOMP, ZINGY, BRAXY
206: FURZY, BLEND, JOCKS, WIGHT, VAMPS
207: FURZY, BLEND, JOCKS, WIGHT, VAMPY
208: FURZY, VELDT, GOWKS, PINCH, JAMBS
209: FURZY, VETCH, GOWDS, PLINK, JAMBS
210: GRUMP, FJELD, BONKS, CHIVY, WALTZ
211: GRUMP, FJELD, BONKS, VICHY, WALTZ
212: GRUMP, FJELD, KNOBS, CHIVY, WALTZ
213: GRUMP, FJELD, KNOBS, VICHY, WALTZ
214: GRUMP, FJELD, ZONKS, CHIVY, BAWTY
215: GRUMP, FJELD, ZONKS, VICHY, BAWTY
216: GRUMP, FJELD, ZONKS, WITCH, BRAXY
217: GUCKS, FJELD, BORTZ, WHINY, VAMPS
218: GUCKS, FJELD, BORTZ, WHINY, VAMPY
219: GUNKS, CHEVY, FJORD, BLIMP, WALTZ
220: GUNKS, CHEVY, FJORD, BLITZ, SWAMP
221: GUNKS, CHEWY, FJORD, BLITZ, VAMPS
222: GUNKS, CHEWY, FJORD, BLITZ, VAMPY
223: GUNKS, FJELD, BORTZ, CHIVY, SWAMP
224: GUNKS, FJELD, BORTZ, VICHY, SWAMP
225: GUNKS, FJELD, BORTZ, WIMPY, SCHAV
226: GUNKS, FJELD, PROXY, CHIMB, WALTZ
227: GUNKY, CHEWS, FJORD, BLITZ, VAMPS
228: GUNKY, CHEWS, FJORD, BLITZ, VAMPY
229: GUNKY, FJELD, BORTZ, WIMPS, SCHAV
230: HUCKS, FJELD, BORTZ, VYING, SWAMP
231: HUCKS, FJELD, BORTZ, WIMPY, VANGS
232: HUCKS, FJELD, BORTZ, WINGY, VAMPS
233: HUCKS, FJELD, BORTZ, WINGY, VAMPY
234: HUMPS, FJELD, BORTZ, VYING, WACKS
235: HUMPS, FJELD, BORTZ, VYING, WACKY
236: HUMPS, FJELD, BORTZ, VYING, WHACK
237: HUMPS, FJELD, BORTZ, VYING, WRACK
238: HUMPS, FJELD, BROCK, VYING, WALTZ
239: HUMPS, GYVED, BLOWN, FRITZ, JACKS
240: HUMPS, GYVED, BLOWN, FRITZ, JACKY
241: HUMPS, GYVED, FROWN, BLITZ, JACKS
242: HUMPS, GYVED, FROWN, BLITZ, JACKY
243: HUMPY, FJELD, BORTZ, WICKS, VANGS
244: JUMPS, CHEVY, BLOND, FRITZ, GAWKS
245: JUMPS, CHEVY, BLOND, FRITZ, GAWKY
246: JUMPS, CHEVY, BORTZ, FLING, DAWKS
247: JUMPS, CHEVY, FROND, BLITZ, GAWKS
248: JUMPS, CHEVY, FROND, BLITZ, GAWKY
249: JUMPS, CHEVY, KLONG, FRITZ, BAWDS
250: JUMPS, CHEVY, KLONG, FRITZ, BAWDY
251: JUMPS, GYVED, BLOCK, FRITZ, SHAWN
252: JUMPS, GYVED, BLOCK, FRITZ, WHANG
253: JUMPS, GYVED, BLOWN, FRITZ, CHALK
254: JUMPS, GYVED, BLOWN, FRITZ, CHARK
255: JUMPS, GYVED, BLOWN, FRITZ, HACKS
256: JUMPS, GYVED, BLOWN, FRITZ, SHACK
257: JUMPS, GYVED, BLOWN, FRITZ, THACK
258: JUMPS, GYVED, BLOWN, FRITZ, WHACK
259: JUMPS, GYVED, BLOWN, THICK, ZARFS
260: JUMPS, GYVED, BORTZ, CHINK, FLAWS
261: JUMPS, GYVED, BORTZ, CHINK, FLAWY
262: JUMPS, GYVED, BORTZ, CHINK, FLAXY
263: JUMPS, GYVED, BORTZ, CLINK, WHARF
264: JUMPS, GYVED, BORTZ, FILCH, SWANK
265: JUMPS, GYVED, BORTZ, FILCH, WANKS
266: JUMPS, GYVED, BORTZ, FINCH, WALKS
267: JUMPS, GYVED, BORTZ, FLICK, SHAWN
268: JUMPS, GYVED, BORTZ, FLICK, WHANG
269: JUMPS, GYVED, BORTZ, WINCH, FLACK
270: JUMPS, GYVED, BORTZ, WINCH, FLAKY
271: JUMPS, GYVED, BORTZ, WINCH, FLANK
272: JUMPS, GYVED, BORTZ, WINCH, FLASK
273: JUMPS, GYVED, BORTZ, WINCH, FLAXY
274: JUMPS, GYVED, BROWN, THICK, FLAXY
275: JUMPS, GYVED, BROWN, ZILCH, KRAFT
276: JUMPS, GYVED, CROWN, BLITZ, KHAFS
277: JUMPS, GYVED, FLOWN, THICK, BRAXY
278: JUMPS, GYVED, FROCK, BLITZ, SHAWN
279: JUMPS, GYVED, FROCK, BLITZ, WHANG
280: JUMPS, GYVED, FROWN, BLITZ, CHALK
281: JUMPS, GYVED, FROWN, BLITZ, CHARK
282: JUMPS, GYVED, FROWN, BLITZ, HACKS
283: JUMPS, GYVED, FROWN, BLITZ, SHACK
284: JUMPS, GYVED, FROWN, BLITZ, THACK
285: JUMPS, GYVED, FROWN, BLITZ, WHACK
286: JUMPS, HERTZ, BLOCK, VYING, DWARF
287: JUMPS, HERTZ, FLOCK, VYING, BAWDS
288: JUMPS, HERTZ, FLOCK, VYING, BAWDY
289: JUMPS, VELDT, BROCK, ZINGY, WHARF
290: JUMPS, WEDGY, BORTZ, CHINK, FLAXY
291: JUMPY, BLEND, GOWKS, FRITZ, SCHAV
292: JUMPY, DRECK, HOWFS, BLITZ, VANGS
293: JUMPY, FLECK, BORTZ, WHIDS, VANGS
294: JUMPY, VELDS, BORTZ, FINCH, GAWKS
295: JUMPY, VELDS, BORTZ, FINCH, GAWKY
296: JUMPY, VELDT, BROCK, ZINGS, WHARF
297: JUMPY, VELDT, GOWKS, FINCH, BRAXY
298: JUMPY, VENDS, BLOCK, FRITZ, WHANG
299: JUMPY, VENDS, BLOCK, WIGHT, ZARFS
300: JUMPY, VENDS, BORTZ, FILCH, GAWKS
301: JUMPY, VENDS, BORTZ, FILCH, GAWKY
302: JUMPY, VENDS, BORTZ, FLICK, WHANG
303: JUMPY, VENDS, BROCK, FIGHT, WALTZ
304: JUMPY, VENDS, BROCK, GLITZ, WHARF
305: JUMPY, VENDS, BROCK, WIGHT, FLAXY
306: JUMPY, VENDS, FLOCK, WIGHT, BRAXY
307: JUMPY, VENDS, FROCK, BIGHT, WALTZ
308: JUMPY, VENDS, FROCK, BLITZ, WHANG
309: JUMPY, VETCH, GOWDS, BLINK, ZARFS
310: JUMPY, VETCH, GOWDS, BRINK, FLAXY
311: JUMPY, VETCH, GOWKS, BLIND, ZARFS
312: JUMPY, VETCH, GROSZ, BLINK, DWARF
313: JUNKS, GYBED, WHOMP, FRITZ, CALYX
314: JUNKS, GYVED, BORTZ, CHIMP, FLAWS
315: JUNKS, GYVED, BORTZ, CHIMP, FLAWY
316: JUNKS, GYVED, BORTZ, CHIMP, FLAXY
317: JUNKS, GYVED, BORTZ, FILCH, SWAMP
318: JUNKS, GYVED, CHOMP, BLITZ, DWARF
319: JUNKS, GYVED, CHOMP, BLITZ, SWARF
320: JUNKS, GYVED, CHOMP, BLITZ, WHARF
321: JUNKS, GYVED, CHOMP, FRITZ, BAWLS
322: JUNKS, GYVED, CHOMP, FRITZ, BLAWN
323: JUNKS, GYVED, CHOMP, FRITZ, BLAWS
324: JUNKS, GYVED, CHOMP, FRITZ, BRAWL
325: JUNKS, GYVED, CHOMP, FRITZ, BYLAW
326: JUNKS, GYVED, CLOMB, FRITZ, PSHAW
327: JUNKS, GYVED, CLOMB, FRITZ, WHAPS
328: JUNKS, GYVED, CLOMB, WHIPT, ZARFS
329: JUNKS, GYVED, PLOTZ, CHIMB, DWARF
330: JUNKS, GYVED, PLOTZ, CHIMB, SWARF
331: JUNKS, GYVED, PLOTZ, CHIMB, WHARF
332: JUNKS, GYVED, WHOMP, BLITZ, CRAFT
333: JUNKS, GYVED, WHOMP, BLITZ, FARCY
334: JUNKS, GYVED, WHOMP, BLITZ, FRANC
335: JUNKS, GYVED, WHOMP, BLITZ, SCARF
336: JUNKS, GYVED, WHOMP, CLIFT, BRAXY
337: JUNKS, GYVED, WHOMP, FRITZ, BLACK
338: JUNKS, GYVED, WHOMP, FRITZ, CALYX
339: JUNKS, WEDGY, BORTZ, CHIMP, FLAXY
340: JUNKS, WEDGY, BORTZ, FILCH, VAMPS
341: JUNKS, WEDGY, BORTZ, FILCH, VAMPY
342: JUNKY, BELCH, GOWDS, FRITZ, VAMPS
343: JUNKY, BELCH, GOWDS, FRITZ, VAMPY
344: JUNKY, VERBS, CHOMP, GLITZ, DWARF
345: JUNKY, VETCH, GOWDS, BLIMP, ZARFS
346: JUNKY, VETCH, GROSZ, BLIMP, DWARF
347: KLUTZ, BENCH, FJORD, WIMPS, GRAVY
348: KLUTZ, BENCH, FJORD, WIMPY, VANGS
349: KLUTZ, CHEVY, FJORD, GIMPS, BLAWN
350: KLUTZ, CHEVY, FJORD, GIMPS, BRAWN
351: KLUTZ, CHEVY, FJORD, WIMPS, BANGS
352: KLUTZ, CHEVY, FJORD, WIMPS, BHANG
353: KLUTZ, GYBES, FJORD, WINCH, VAMPS
354: KLUTZ, GYBES, FJORD, WINCH, VAMPY
355: KLUTZ, GYVED, FROWN, CHIMP, JAMBS
356: KLUTZ, GYVED, FROWN, CHIPS, JAMBS
357: KLUTZ, GYVED, FROWS, PINCH, JAMBS
358: KLUTZ, GYVED, PROFS, WINCH, JAMBS
359: KLUTZ, GYVED, PROWS, FINCH, JAMBS
360: KLUTZ, GYVES, FJORD, CHIMB, PAWNS
361: KLUTZ, GYVES, FJORD, CHIMB, PRAWN
362: KLUTZ, GYVES, FJORD, CHIMB, SPAWN
363: KLUTZ, GYVES, FJORD, CHIMP, BLAWN
364: KLUTZ, GYVES, FJORD, CHIMP, BRAWN
365: KLUTZ, PREXY, GOWDS, FINCH, JAMBS
366: KLUTZ, PYREX, GOWDS, FINCH, JAMBS
367: KLUTZ, PYXES, FJORD, CHIMB, GNAWS
368: KLUTZ, PYXES, FJORD, CHIMB, SWANG
369: KLUTZ, PYXES, FJORD, CHIMB, TWANG
370: KLUTZ, PYXES, FJORD, CHIMB, VANGS
371: KLUTZ, PYXES, FJORD, CHIMB, WHANG
372: KLUTZ, PYXES, FJORD, CHIMB, WRANG
373: KLUTZ, PYXES, FJORD, WINCH, GAMBS
374: KLUTZ, WENCH, FJORD, GIMPS, BRAXY
375: KLUTZ, WENCH, FJORD, PYXIS, GAMBS
376: MUCKS, FJELD, BORTZ, VYING, PSHAW
377: MUCKS, FJELD, BORTZ, VYING, WHAPS
378: MUCKS, GYVED, PHLOX, FRITZ, BLAWN
379: MUCKS, GYVED, PHLOX, FRITZ, BRAWN
380: MUCKY, FJELD, BORTZ, WHIPS, VANGS
381: MUCKY, FJELD, QOPHS, BRING, WALTZ
382: MUNCH, FJELD, BORTZ, PYXIS, GAWKS
383: MUNCH, FJELD, BORTZ, PYXIS, GAWKY
384: MUNCH, GYVES, FJORD, BLITZ, PAWKY
385: MUNCH, PYXES, FJORD, BLITZ, GAWKS
386: MUNCH, PYXES, FJORD, BLITZ, GAWKY
387: NUDZH, BECKS, JOWLY, GRIFT, VAMPS
388: NUDZH, BECKS, JOWLY, GRIFT, VAMPY
389: NUDZH, CLEFT, GOWKS, JIMPY, BRAXY
390: NUDZH, CLEFT, GOWKS, PRIVY, JAMBS
391: NUDZH, FLECK, PROXY, TWIGS, JAMBS
392: NUDZH, GYVES, FLOCK, TWIRP, JAMBS
393: NUDZH, PREXY, FLOCK, TWIGS, JAMBS
394: NUDZH, PREXY, GOWKS, CLIFT, JAMBS
395: NUDZH, PYREX, FLOCK, TWIGS, JAMBS
396: NUDZH, PYREX, GOWKS, CLIFT, JAMBS
397: NUDZH, VERBS, FLOCK, JIMPY, TWANG
398: NUDZH, WEFTS, BLOCK, JIMPY, GRAVY
399: NUMBS, CHEVY, FJORD, GLITZ, PAWKY
400: NUMBS, GYVED, PHLOX, FRITZ, JACKS
401: NUMBS, GYVED, PHLOX, FRITZ, JACKY
402: NUMBS, GYVED, PHLOX, FRITZ, WACKS
403: NUMBS, GYVED, PHLOX, FRITZ, WACKY
404: NUMBS, GYVED, PHLOX, FRITZ, WHACK
405: NUMBS, GYVED, PHLOX, FRITZ, WRACK
406: NUMBS, VETCH, FJORD, KYLIX, GAWPS
407: NUMBS, WEDGY, PHLOX, FRITZ, JACKS
408: NUMBS, WEDGY, PHLOX, FRITZ, JACKY
409: PLUCK, GYVED, FROWN, SIXTH, JAMBS
410: PLUCK, GYVED, SHOWN, FRITZ, JAMBS
411: PLUMB, CHEVY, FJORD, GINKS, WALTZ
412: PLUMB, CHEVY, FJORD, KINGS, WALTZ
413: PLUMB, GYVED, FROWN, SIXTH, JACKS
414: PLUMB, GYVED, FROWN, SIXTH, JACKY
415: PLUMB, GYVED, JOCKS, FRITZ, SHAWN
416: PLUMB, GYVED, JOCKS, FRITZ, WHANG
417: PLUMB, GYVED, JOHNS, FRITZ, WACKS
418: PLUMB, GYVED, JOHNS, FRITZ, WACKY
419: PLUMB, GYVED, JOHNS, FRITZ, WHACK
420: PLUMB, GYVED, JOHNS, FRITZ, WRACK
421: PLUMB, GYVED, SHOWN, FRITZ, JACKS
422: PLUMB, GYVED, SHOWN, FRITZ, JACKY
423: PLUMB, GYVES, FJORD, CHINK, WALTZ
424: PLUMB, HECKS, FJORD, VYING, WALTZ
425: PLUMB, HERTZ, JOCKS, VYING, DWARF
426: PLUMB, VETCH, FJORD, SWINK, GLAZY
427: PLUMB, VETCH, FJORD, WINKS, GLAZY
428: PLUMB, VETCH, FJORD, ZINGS, GAWKY
429: PLUMB, VETCH, FJORD, ZINGS, PAWKY
430: PLUMB, VETCH, FJORD, ZINGS, WACKY
431: PLUMB, VETCH, FJORD, ZINGY, DAWKS
432: PLUMB, VETCH, FJORD, ZINGY, GAWKS
433: PLUMB, VETCH, FJORD, ZINGY, HAWKS
434: PLUMB, VETCH, FJORD, ZINGY, SWANK
435: PLUMB, VETCH, FJORD, ZINGY, WACKS
436: PLUMB, VETCH, FJORD, ZINGY, WALKS
437: PLUMB, VETCH, FJORD, ZINGY, WANKS
438: PLUMB, VETCH, FJORD, ZINGY, WARKS
439: PLUMB, VETCH, FJORD, ZINKY, GAWKS
440: PLUMB, VETCH, FJORD, ZINKY, GAWPS
441: PLUMB, VETCH, FJORD, ZINKY, GAWSY
442: PLUMB, VETCH, FJORD, ZINKY, GNAWS
443: PLUMB, VETCH, FJORD, ZINKY, SWANG
444: PLUMB, WECHT, FJORD, ZINKY, VANGS
445: PLUNK, CHEVY, GOWDS, FRITZ, JAMBS
446: PLUNK, GYVED, CHOWS, FRITZ, JAMBS
447: PLUNK, GYVES, FJORD, CHIMB, WALTZ
448: PUCKS, FJELD, BORTZ, VYING, SHAWM
449: PUCKS, FJELD, BORTZ, VYING, WHAMS
450: PUCKS, FJELD, RHOMB, VYING, WALTZ
451: RHUMB, FJELD, POCKS, VYING, WALTZ
452: RHUMB, GYVED, FLOWN, SPITZ, JACKS
453: RHUMB, GYVED, FLOWN, SPITZ, JACKY
454: RUGBY, FJELD, ZONKS, WITCH, VAMPS
455: RUGBY, FJELD, ZONKS, WITCH, VAMPY
456: SHUCK, FJELD, BORTZ, VYING, SWAMP
457: SHUCK, FJELD, BORTZ, WIMPY, VANGS
458: SHUCK, FJELD, BORTZ, WINGY, VAMPS
459: SHUCK, FJELD, BORTZ, WINGY, VAMPY
460: SNUCK, GYVED, PHLOX, FRITZ, JAMBS
461: SNUCK, WEDGY, PHLOX, FRITZ, JAMBS
462: STUCK, FJELD, WHOMP, VYING, BRAXY
463: STUCK, FJELD, WHOMP, ZINGY, BRAXY
464: SWUNG, KEMPT, FJORD, ZILCH, BRAXY
465: THUMB, FJELD, PROXY, WICKS, VANGS
466: THUMB, FJELD, PROXY, ZINCS, GAWKS
467: THUMB, FJELD, PROXY, ZINCS, GAWKY
468: THUMB, FJELD, PROXY, ZINGS, WACKS
469: THUMB, FJELD, PROXY, ZINGS, WACKY
470: THUMB, FJELD, PROXY, ZINGS, WHACK
471: THUMB, FJELD, PROXY, ZINGS, WRACK
472: THUMB, FJELD, VROWS, ZINGY, PACKS
473: THUMB, FJELD, VROWS, ZINGY, PLACK
474: THUMB, PECKS, FJORD, VYING, WALTZ
475: THUMB, SPECK, FJORD, VYING, WALTZ
476: THUMP, BECKS, FJORD, VYING, WALTZ
477: THUMP, FJELD, GOWKS, ZINCY, BRAXY
478: THUMP, FJELD, VROWS, ZINGY, BACKS
479: THUMP, FJELD, VROWS, ZINGY, BLACK
480: TRUCK, FJELD, WHOMP, ZINGS, BRAXY
481: TUCKS, FJELD, WHOMP, VYING, BRAXY
482: TUCKS, FJELD, WHOMP, ZINGY, BRAXY
483: VUGHS, BLEND, FROCK, JIMPY, WALTZ
484: VUGHS, FJELD, BORTZ, MINCY, PAWKY
485: VUGHS, FJELD, BORTZ, WIMPY, CLANK
486: VUGHS, FJELD, BORTZ, WIMPY, CRANK
487: VUGHS, FJELD, BORTZ, WIMPY, SNACK
488: VUGHS, FJELD, BROCK, WIMPY, TRANQ
489: VUGHS, FJELD, COMPT, ZINKY, BRAWL
490: VUGHS, FJELD, COMPT, ZINKY, BRAWN
491: VUGHS, FJELD, COMPT, ZINKY, BRAWS
492: VUGHS, FJELD, COMPT, ZINKY, BRAXY
493: VUGHS, FJELD, WOMBY, PRICK, TRANQ
494: VUGHS, FLECK, BORTZ, JIMPY, DAWNS
495: VUGHS, FLECK, BORTZ, JIMPY, DRAWN
496: VUGHS, FLECK, BORTZ, JIMPY, WANDS
497: VUGHS, KEMPT, FJORD, CYLIX, BLAWN
498: VUGHS, KEMPT, FJORD, CYLIX, BRAWN
499: VUGHS, KEMPT, FJORD, ZINCY, BAWLS
500: VUGHS, KEMPT, FJORD, ZINCY, BLAWN
501: VUGHS, KEMPT, FJORD, ZINCY, BLAWS
502: VUGHS, KEMPT, FJORD, ZINCY, BRAWL
503: VUGHS, KEMPT, FJORD, ZINCY, BYLAW
504: VUGHS, NERTZ, BLOCK, JIMPY, DWARF
505: VUGHS, NERTZ, FLOCK, JIMPY, BAWDS
506: VUGHS, NERTZ, FLOCK, JIMPY, BAWDY
507: WHUMP, BECKS, FJORD, GLITZ, XYLAN
508: WHUMP, BECKS, FJORD, VYING, WALTZ
509: WHUMP, BLEND, JOCKS, FRITZ, GRAVY
510: WHUMP, DECKS, BORTZ, VYING, FLAXY
511: WHUMP, DELFS, BORTZ, VYING, JACKS
512: WHUMP, DELFS, BORTZ, VYING, JACKY
513: WHUMP, DELFT, JOCKS, VYING, BRAXY
514: WHUMP, DELFT, JOCKS, ZINGY, BRAXY
515: WHUMP, FJELD, BOCKS, RITZY, VANGS
516: WHUMP, FJELD, BOCKS, VYING, TRANQ
517: WHUMP, FJELD, BOCKS, VYING, TZARS
518: WHUMP, FJELD, BOCKS, ZINGY, TRANQ
519: WHUMP, FJELD, BORTZ, GINKS, CALYX
520: WHUMP, FJELD, BORTZ, KINGS, CALYX
521: WHUMP, FJELD, BORTZ, NICKS, GRAVY
522: WHUMP, FJELD, BORTZ, SNICK, GRAVY
523: WHUMP, FJELD, BORTZ, VYING, BACKS
524: WHUMP, FJELD, BORTZ, VYING, CALKS
525: WHUMP, FJELD, BORTZ, VYING, CARKS
526: WHUMP, FJELD, BORTZ, VYING, CASKY
527: WHUMP, FJELD, BORTZ, VYING, HACKS
528: WHUMP, FJELD, BORTZ, VYING, JACKS
529: WHUMP, FJELD, BORTZ, VYING, LACKS
530: WHUMP, FJELD, BORTZ, VYING, MACKS
531: WHUMP, FJELD, BORTZ, VYING, PACKS
532: WHUMP, FJELD, BORTZ, VYING, RACKS
533: WHUMP, FJELD, BORTZ, VYING, SHACK
534: WHUMP, FJELD, BORTZ, VYING, SLACK
535: WHUMP, FJELD, BORTZ, VYING, SMACK
536: WHUMP, FJELD, BORTZ, VYING, SNACK
537: WHUMP, FJELD, BORTZ, VYING, STACK
538: WHUMP, FJELD, BORTZ, VYING, TACKS
539: WHUMP, FJELD, BORTZ, VYING, WACKS
540: WHUMP, FJELD, BORTZ, VYING, YACKS
541: WHUMP, FJELD, BROCK, SIXTY, VANGS
542: WHUMP, FJELD, BROCK, VYING, TZARS
543: WHUMP, FJELD, BROCK, XYSTI, VANGS
544: WHUMP, FJELD, BROCK, ZINGS, VASTY
545: WHUMP, FJELD, BROCK, ZINGY, VASTY
546: WHUMP, FJELD, STOCK, VYING, BRAXY
547: WHUMP, FJELD, STOCK, ZINGY, BRAXY
548: WHUMP, FJELD, TROCK, ZINGS, BRAXY
549: WHUMP, FLECK, BORTZ, DJINS, GRAVY
550: WHUMP, GECKS, FJORD, BLITZ, XYLAN
551: WHUMP, GYBED, JOCKS, FRITZ, XYLAN
552: WHUMP, GYVED, BOCKS, FRITZ, XYLAN
553: WHUMP, GYVED, BONKS, FRITZ, CALYX
554: WHUMP, GYVED, BORTZ, FINKS, CALYX
555: WHUMP, GYVED, BORTZ, JINKS, CALFS
556: WHUMP, GYVED, BORTZ, JINKS, CALYX
557: WHUMP, GYVED, BORTZ, JINKS, FLACK
558: WHUMP, GYVED, BORTZ, JINKS, FLAXY
559: WHUMP, GYVED, BORTZ, NICKS, FLAXY
560: WHUMP, GYVED, BORTZ, SNICK, FLAXY
561: WHUMP, GYVED, CLONK, FRITZ, JAMBS
562: WHUMP, GYVED, FRONS, BLITZ, JACKS
563: WHUMP, GYVED, FRONS, BLITZ, JACKY
564: WHUMP, GYVED, JOCKS, BLITZ, FRANC
565: WHUMP, GYVED, JOCKS, BLITZ, FRANK
566: WHUMP, GYVED, JOCKS, BLITZ, SNARF
567: WHUMP, GYVED, JOCKS, BLITZ, TRANQ
568: WHUMP, GYVED, JOCKS, FLINT, BRAXY
569: WHUMP, GYVED, JOCKS, FRITZ, BLAND
570: WHUMP, GYVED, JOCKS, FRITZ, BLANK
571: WHUMP, GYVED, JOCKS, FRITZ, BLAWN
572: WHUMP, GYVED, JOCKS, FRITZ, XYLAN
573: WHUMP, GYVED, KNOBS, FRITZ, CALYX
574: WHUMP, GYVED, ZONKS, CLIFT, BRAXY
575: WHUMP, GYVES, BLOND, FRITZ, JACKS
576: WHUMP, GYVES, BLOND, FRITZ, JACKY
577: WHUMP, GYVES, FJORD, BLITZ, CLANK
578: WHUMP, GYVES, FJORD, BLITZ, CRANK
579: WHUMP, GYVES, FJORD, BLITZ, SNACK
580: WHUMP, GYVES, FROND, BLITZ, JACKS
581: WHUMP, GYVES, FROND, BLITZ, JACKY
582: WHUMP, NECKS, FJORD, BLITZ, GRAVY
583: WHUMP, NECKS, FJORD, GLITZ, BRAXY
584: WHUMP, SNECK, FJORD, BLITZ, GRAVY
585: WHUMP, SNECK, FJORD, GLITZ, BRAXY
586: WHUMP, VELDT, BROCK, ZINGS, FLAXY
587: WHUMP, VELDT, FORBS, ZINGY, JACKS
588: WHUMP, VELDT, FORBS, ZINGY, JACKY
589: WHUMP, VELDT, FORBY, ZINGS, JACKS
590: WHUMP, VELDT, FORBY, ZINGS, JACKY
591: WHUMP, VELDT, FROCK, ZINGS, BRAXY
592: WHUMP, VELDT, FROCK, ZINGY, JAMBS
593: WHUMP, VELDT, JOCKS, BRING, FLAXY
594: WHUMP, VELDT, JOCKS, ZINGY, BARFS
595: WHUMP, VELDT, JOCKS, ZINGY, BRAXY
596: WHUMP, VENDS, BROCK, GLITZ, FLAXY
597: WHUMP, VENDS, FORBY, GLITZ, JACKS
598: WHUMP, VENDS, FORBY, GLITZ, JACKY
599: WHUMP, VENDS, FROCK, GLITZ, BRAXY
600: WHUMP, VERBS, FLONG, DITZY, JACKS
601: WHUMP, VERBS, FLONG, DITZY, JACKY
602: WHUPS, FJELD, BORTZ, VYING, MACKS
603: WHUPS, FJELD, BORTZ, VYING, SMACK
604: WHUPS, GYVED, CLONK, FRITZ, JAMBS
605: WHUPS, VELDT, FROCK, ZINGY, JAMBS
606: WRUNG, FECKS, PHLOX, DITZY, JAMBS
607: WRUNG, FLECK, QOPHS, DITZY, JAMBS
```