---
layout: ../layouts/Layout.astro
title: "Project: Full-Stack Disc Golf Smartwatch App & Analytics Dashboard"
author: "Computer Science Graduate"
description: "A full-stack project connecting an embedded watch application to a data analytics dashboard."
---

# Full-Stack Disc Golf Smartwatch App & Analytics Platform

This project highlights full-stack software development by building an end-to-end data pipeline: capturing scores on low-power wearable devices and syncing that data to a centralized analytics dashboard.

## Architectural Layers

1. Embedded Client Layer (Smartwatch Application)
* Targets Garmin Connect IQ (Monkey C) or WearOS (Kotlin).
* Caches round information locally in memory to handle offline play when course cell coverage is weak.
* Bundles and transmits the scorecard payload via JSON once an internet connection is established.

2. Ingestion Layer (Backend API)
* Built using Python (FastAPI) or Node.js running inside isolated Docker containers.
* Validates incoming score arrays, processes rounds relative to course par, and calculates player handicap changes.

3. Persistence Layer (Relational Database)
* Built using a PostgreSQL instance managed within a custom container network.
* Manages a relational schema mapping players, courses, hole layouts, and historical scorecards.
* Generates standard file exports compatible with manual external handicap tracking platforms.

4. UI Layer (Web Dashboard)
* Engineered using Astro and modern client-side charting libraries.
* Consumes backend endpoints to visualize data trends, tracking score progressions, putting accuracy, and course-specific statistics over time.


