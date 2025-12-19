# Issue Title

**Bug: Keyboard scrolling doesn't work in full-screen detail view on narrow terminals**

# Issue Description

First off, thank you for creating bv — it's been incredibly useful for navigating Beads projects!

## Problem

On narrow terminals, bv displays a full-screen detail view when pressing Enter on an issue. However, **keyboard scrolling doesn't work** — arrow keys, j/k, Page Up/Down, and other navigation keys have no effect on the viewport content.

This makes it impossible to read long issue descriptions, comments, or dependency trees on small terminal windows (common on laptops, split-screen setups, or when SSH'ing from mobile).

**Reproduction steps:**
1. Resize terminal to a narrow width (triggers full-screen detail view instead of split view)
2. Select an issue and press Enter
3. Try to scroll with arrow keys, j/k, Ctrl+d/u, or Page Up/Down
4. Nothing happens — content doesn't scroll

## Root Cause

When entering the full-screen detail view, the code sets `showDetails = true` but does not update the focus state. Focus remains on `focusList`, so key events continue being routed to the (hidden) list component instead of the visible viewport.

## Suggested Fix

I understand this repo doesn't accept external contributions, but I've put together a potential fix in PR #25 that ensures proper focus management when entering/exiting the detail view. Feel free to review it for reference if you'd like to implement a fix yourself.

The approach:
1. Sets focus to `focusDetail` when entering full-screen detail view from any view (list, board, graph, actionable, insights)
2. Resets the viewport to top when selecting a new issue (prevents stale scroll position)
3. Restores focus to `focusList` when exiting detail view (q/Esc)

Thanks again for your work on this tool!
