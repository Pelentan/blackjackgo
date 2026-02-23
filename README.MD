# Blackjack in Go

## Purpose

The initial purpose of this application was for a prospective employer to get a look at my coding style. I hate doing something for just one reason, so I decided to also use it to prove out an alternative Engineer/AI partnering solution than the one I did with this repository:

https://github.com/Pelentan/lora-dataset-prep

The experience was nowhere near as smooth this time around.

## The Experiment

The goal was to prove out the feasibility of a more "pair-programming" dynamic using a local AI rig. The AI rig was a 5090 running Ollama, accessed via Open WebUI. The model used was Qwen2.5-Coder:32b. I created a Workspace and attached a Knowledge Base containing a guidelines file I use to prime AI assistants with a set of conventions. This last may need (a lot) of tweaking for Qwen. The intent is to validate the ability to use AI in a secure environment.

The process was straightforward in concept: I worked out the layout of the application and what I wanted each part to do generally, created the basic stubs, then fed it to Qwen with my first request for guidance. It was then a continuous back and forth: I ask questions and use the responses to code things in.  Sort of like a turbo-charged session of coding and Stackoverflow queries.

Problems quickly arose. It kept forgetting changes I had just made. The interface got iffy when copy-pasting. Loading a file directly was a straight no-go. Not ideal. I'm pretty sure I spent more time trying to get that to work than it would have taken to just code it straight and look things up online. When I had something working, I fed it to Claude.ai and asked it to rate the code. It gave it a C+. Higher than I would have given it. Claude and I tightened it up to what is here. Nothing major, as the framework was solid. Just some minor tweaks that smoothed things out.

I don't consider it a failure or wasted time. I think I know what I need to do to set up the Workspace better. Qwen didn't get anything *wrong* technically, nor did it really hallucinate. It just kept losing track of what I had done, on top of not fully internalizing the guidelines file. It was reading it and following some of the instructions, just not all of them. That will need to be tweaked for the specific model.

I think part of the problem also came from trying to use the "pair-programming" technique. If I had used the technique I used in the other project it would have done better. But I wanted to test this method out, and now I have results.

## Additional Features

- **Betting system** — player starts with $500 and places a bet each hand
- **Deck reshuffling** — automatic reshuffle when the deck runs low
- **Win/loss tracking** — running totals displayed at game over
- **Dramatic pause to determine winner** — because user experience matters

## Tech Stack

- Go 1.22+

## Build and Run

```bash
go build .
./blackjackgo
```

On Windows:

```bash
go build .
./blackjackgo.exe
```