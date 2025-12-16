package export

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Dicklesworthstone/beads_viewer/pkg/analysis"
	"github.com/Dicklesworthstone/beads_viewer/pkg/model"
)

func TestSaveGraphSnapshot_SVG_Escaping(t *testing.T) {
	issues := []model.Issue{
		{ID: "A", Title: "Dangerous <script>", Status: model.StatusOpen},
	}
	analyzer := analysis.NewAnalyzer(issues)
	stats := analyzer.Analyze()

	tmp := t.TempDir()
	out := filepath.Join(tmp, "unsafe.svg")
	
	err := SaveGraphSnapshot(GraphSnapshotOptions{
		Path:     out,
		Format:   "svg",
		Issues:   issues,
		Stats:    &stats,
		DataHash: "hash",
	})
	if err != nil {
		t.Fatalf("SaveGraphSnapshot error: %v", err)
	}

	content, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	svg := string(content)

	// Check if the title text is properly escaped
	// We look for the exact string "Dangerous <script>" inside the SVG.
	// If it's present verbatim inside a <text> tag without escaping, it's invalid XML if it contains < or > or &.
	// svgo writes raw strings. So "Dangerous <script>" will be written as-is.
	// We verify that it IS escaped (e.g. "Dangerous &lt;script&gt;") or at least not present in raw form.
	
	if !strings.Contains(svg, "Dangerous &lt;script&gt;") {
		t.Errorf("SVG does not contain escaped text: %s\nFull SVG:\n%s", "Dangerous &lt;script&gt;", svg)
	}
}
