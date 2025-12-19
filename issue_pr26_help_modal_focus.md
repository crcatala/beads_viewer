# Issue Title

**Bug: Keyboard navigation broken after dismissing help modal in specialized views**

# Issue Description

First off, thank you for creating bv — it's been incredibly useful for navigating Beads projects!

## Problem

After opening the help overlay (`?`) from specialized views (graph, board, actionable, insights, or full-screen detail view), dismissing it breaks keyboard navigation. The user is unexpectedly returned to the list view instead of their previous view.

**Reproduction steps:**
1. Open graph view (`g`)
2. Navigate to a node with arrow keys (works)
3. Open help (`?`)
4. Dismiss help (press any key)
5. Try to navigate with arrow keys → **navigation broken**, focus is on list

This is confusing because the graph is still visually displayed, but keyboard input no longer controls it.

## Root Cause

The help dismiss logic unconditionally sets `m.focused = focusList`, ignoring which view was active before help was opened.

## Suggested Fix

I understand this repo doesn't accept external contributions, but I've put together a potential fix in PR #26 that adds a `restoreFocusFromHelp()` helper to intelligently return focus to the appropriate view based on current state. Feel free to review it for reference if you'd like to implement a fix yourself.

The approach restores focus to:
- `focusDetail` if in full-screen detail view
- `focusGraph` if in graph view
- `focusBoard` if in board view
- `focusActionable` if in actionable view
- `focusInsights` if in insights view
- `focusList` as default fallback

Thanks again for your work on this tool!
