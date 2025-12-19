# Issue Title

**Bug: Kanban board rendering overflow/artifacts at certain terminal widths**

# Issue Description

First off, thank you for creating bv — it's been incredibly useful for navigating Beads projects!

## Problem

At certain terminal widths, the Kanban board renders incorrectly with visual artifacts:
- Column borders appear to have newlines every other line (double-spaced appearance)
- Swimlanes with many cards push the status bar off-screen
- Selected card highlighting looks inconsistent with partial gray backgrounds

This seems to affect various terminal widths unpredictably.

## Root Causes

After investigation, I found several contributing factors:

### 1. lipgloss Border Width Not Accounted For
`lipgloss.Border()` adds 2 characters (left + right) **on top of** the specified `Width()`, not inside it. So `Width(baseWidth).Border(...)` actually renders at `baseWidth + 2`, causing overflow.

### 2. Column Gaps Reserved But Never Rendered
The code calculates available width by reserving space for 2-char gaps between columns, but `JoinHorizontal()` just concatenates without actual gaps — the reserved space goes unused, throwing off the math.

### 3. Card Content Not Truncated to Actual Content Area
Long titles and metadata weren't being truncated to fit the card's actual content area (`width - 4` for border + padding), allowing text to wrap and increase card height beyond the expected 6 lines.

## Suggested Fix

I understand this repo doesn't accept external contributions, but I've put together a potential fix in PR #24 that addresses all three issues. Feel free to review it for reference if you'd like to implement a fix yourself.

Key changes in that PR:
- Subtract border width when setting column/card widths
- Add explicit `MarginRight(2)` between columns instead of relying on reserved space
- Truncate card content at `width - 4` to account for border + padding
- Use `DoubleBorder()` for selected cards instead of background highlighting

Thanks again for your work on this tool!
