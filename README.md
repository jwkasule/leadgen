# LEADGEN

Leadgen, a cli client for hitting the google places API for a search
term, lat/long coordinates, and radius, and outputting the results 
with phone numbers and emails in a CSV file for every example google 
maps can find for these inputs.

Usage:

Make sure API_KEY environment variable is set first (source env/api.sh)
and run as follows, making sure [query] and [location] are in quotes:

leadgen [filename] "[query]" "[location]" [radius]

Where "filename" saves the results in data/filename.csv, "[query]" is the 
type of entity you are searching for, "[location]" is in the format of 
coordinates of lat,long and radius is an integer measuring meters from 
the location point to search from (No more than 50000 metres/50km allowed).

Resulting CSV file will be in data/[filename].csv

# LICENSE

Copyright 2022 Dmitri DB

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
