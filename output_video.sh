#!/bin/bash

fps="30"

ffmpeg -framerate $fps -pattern_type glob -i 'frames/*.png' -c:v ffv1 out_$fps-fps.avi



# facts:
# fps = 60 frames per second
# bpm = 95 beats per minute
# bpp = 8 beats per period

# calc:
# pbf = bpm / (fps * 60) = 95 / (60 * 60) beats per frame

# final:
# fpp = bpp / pbf = bpp / bpm * (fps * 60) = 8 / 95 * (60 * 60) frams per period

# 2*2*2 / (19 * 5) * (10 * 10 * 3 * 3 * 2 * 2)





