package migrate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

func ApplyDir(ctx context.Context, exec Executor, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("migrate.ApplyDir: read dir: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".up.sql") {
			continue
		}
		files = append(files, filepath.Join(dir, name))
	}

	sort.Strings(files)
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("migrate.ApplyDir: read file %s: %w", file, err)
		}
		if _, err := exec.Exec(ctx, string(content)); err != nil {
			return fmt.Errorf("migrate.ApplyDir: exec %s: %w", file, err)
		}
	}

	return nil
}
