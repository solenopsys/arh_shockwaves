package sorters

import (
	"xs/internal/configs"
	"xs/internal/jobs"
)

type Sorter interface {
	Sort(libs []*configs.XsModule) []jobs.PrintableJob
}
