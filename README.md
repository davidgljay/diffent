#Differential Entropy Analysis Test

This project is an experiment in analyzing the differential entropy of natural language messages and their metadata. 

The package is still under development, and will perform the following steps:

1) Import data from an arbitrary number of raw text files.
2) Identify a timestamp for each file using first date found.
3) Preprocess that text using TFIF.
4) Iterate through n-grams within each file and perform an entropy calculation (see below).
5) Export the resulting timeline of entropy values for analysis.

##Entropy Calculation
Shifts in framing over time are often significant. If the word "tax" is likely to appear in the phrase "tax reform" and 
shifts to the phrase "death tax" it can represent a meaningful shift in the underlying discourse. In order to capture these
high-level shifts, this algorithm calculates the probability distribution of word pairs across an n-gram, then calculates the
entropy of those probability distributions over time. This makes it easy to see when work distributions are relatively stable
(everyone is talking about "tax reform"), when they are unstable (the word "tax" has no clear central framing), and when shifts
happen in between. The hope is that this metric may prove an interesting feature in other forms of analysis.