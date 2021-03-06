// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"

	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Experiments ExperimentList
	Reports     ReportList
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Experiments: s.Experiments.Clone(),
		Reports:     s.Reports.Clone(),
	}
}

func (s ApiSnapshot) Hash() uint64 {
	return hashutils.HashAll(
		s.hashExperiments(),
		s.hashReports(),
	)
}

func (s ApiSnapshot) hashExperiments() uint64 {
	return hashutils.HashAll(s.Experiments.AsInterfaces()...)
}

func (s ApiSnapshot) hashReports() uint64 {
	return hashutils.HashAll(s.Reports.AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	fields = append(fields, zap.Uint64("experiments", s.hashExperiments()))
	fields = append(fields, zap.Uint64("reports", s.hashReports()))

	return append(fields, zap.Uint64("snapshotHash", s.Hash()))
}

type ApiSnapshotStringer struct {
	Version     uint64
	Experiments []string
	Reports     []string
}

func (ss ApiSnapshotStringer) String() string {
	s := fmt.Sprintf("ApiSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Experiments %v\n", len(ss.Experiments))
	for _, name := range ss.Experiments {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Reports %v\n", len(ss.Reports))
	for _, name := range ss.Reports {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	return ApiSnapshotStringer{
		Version:     s.Hash(),
		Experiments: s.Experiments.NamespacesDotNames(),
		Reports:     s.Reports.NamespacesDotNames(),
	}
}
