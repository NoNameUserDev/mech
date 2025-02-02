# Docs

## How to save HAR?

Open browser and open a new tab. Click `Open menu`, `Web Developer`, `Network`.
Then click `Network Settings`, `Persist Logs`. Also check `Disable Cache`. Then
browse to the page you want to capture. Once you are ready, click `Pause`, then
click `Network Settings`, `Save All As HAR`. Rename file to JSON.

## How to add a new site?

1. `--proxy 127.0.0.1:8080 --no-check-certificate`
2. try navigating to the target page from home screen, instead of going directly
   to page
3. check media page for JSON requests
4. check HAR file
5. check HTML
6. check JavaScript
7. MITM Proxy APK for JSON requests
8. JADX

## Proxy

https://freeproxylists.net/br.html

## Sites

~~~
11.366 B true YouTube
4.046 B true Instagram
2.017 B true TikTok
1.276 B false Spotify: Music and Podcasts
1.258 B true Twitter
338.228 M true SoundCloud: Play Music & Songs
315.762 M false Deezer: Music & Podcast Player
278.179 M true Pandora - Music & Podcasts
252.302 M false Amazon Music: Discover Songs
167.009 M true IMDb: Your guide to movies, TV shows, celebrities
132.597 M true Tumblr – Culture, Art, Chaos
92.476 M false Reddit
84.791 M false iHeart: Radio, Music, Podcasts
82.606 M false Apple Music
31.354 M true Vimeo
29.674 M true TED
24.683 M false TIDAL Music
21.914 M false Paramount+ | Peak Streaming
21.298 M false Napster Music
13.757 M true The NBC App - Stream TV Shows
4.652 M true MTV
2.613 M true Bandcamp
2.224 M false PBS Video: Live TV & On Demand
687.669 K false Qobuz
~~~

Amazon Music:

First 30 seconds is normal, rest of track is encrypted.

Apple Podcasts:

https://github.com/89z/mech/tree/12876427d9a0629de6131734c3e778daf1a62b7a

BBC:

https://github.com/89z/mech/tree/43f0466455e0204ec3319900def5e1ae638aaa8b

Bleep:

https://github.com/89z/mech/tree/f2e43b8d5557c732c7deaee78a815e6fb22378b5

PBS:

https://github.com/89z/mech/tree/c825743ab7594025b9c70632d934820e2c68d20a

Reddit:

https://github.com/89z/mech/tree/b901af458f09e04b662dbeabc8f5880527f9adfa

Spotify:

Audio is encrypted.
