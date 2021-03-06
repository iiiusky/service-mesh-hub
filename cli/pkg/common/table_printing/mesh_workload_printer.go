package table_printing

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing/internal"
	smh_discovery "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
)

func NewMeshWorkloadPrinter(tableBuilder TableBuilder) MeshWorkloadPrinter {
	return &meshWorkloadPrinter{
		tableBuilder: tableBuilder,
	}
}

type meshWorkloadPrinter struct {
	tableBuilder TableBuilder
}

func (m *meshWorkloadPrinter) Print(
	out io.Writer,
	meshWorkloads []*smh_discovery.MeshWorkload,
) error {
	if len(meshWorkloads) == 0 {
		return nil
	}

	fmt.Fprintln(out, "\nKubernetes Mesh Workloads:")

	table := m.tableBuilder(out)

	commonHeaderRows := []string{
		"Controller",
		"Mesh",
		"Labels",
		"Identity",
	}

	var preFilteredRows [][]string
	for _, meshWorkload := range meshWorkloads {
		var newRow []string
		kubeController := meshWorkload.Spec.GetKubeController()
		// Build Controller Name
		if kubeController.GetKubeControllerRef() != nil {
			newRow = append(newRow, fmt.Sprintf(
				"Name: %s\nNamespace: %s\nCluster: %s\nType: Deployment\n",
				kubeController.GetKubeControllerRef().GetName(),
				kubeController.GetKubeControllerRef().GetNamespace(),
				kubeController.GetKubeControllerRef().GetCluster(),
			))
		} else {
			newRow = append(newRow, "")
		}
		// Append common metadata
		newRow = append(newRow, meshWorkload.Spec.GetMesh().GetName())

		// Add workload labels
		if len(kubeController.GetLabels()) != 0 {
			var labels []string
			for k, v := range kubeController.GetLabels() {
				labels = append(labels, fmt.Sprintf("%s: %s", k, v))
			}
			// For test idempotence
			sort.Strings(labels)
			newRow = append(newRow, strings.Join(labels, "\n"))
		} else {
			newRow = append(newRow, "")
		}

		// Add Identity data
		newRow = append(newRow, kubeController.GetServiceAccountName())

		preFilteredRows = append(preFilteredRows, newRow)
	}

	filteredHeaders, filteredRows := internal.FilterEmptyColumns(commonHeaderRows, preFilteredRows)

	table.SetHeader(filteredHeaders)
	table.AppendBulk(filteredRows)
	table.Render()
	return nil
}
