package table_printing_test

import (
	"bytes"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common/table_printing/test_goldens"
	smh_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	smh_security "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1"
	smh_security_types "github.com/solo-io/service-mesh-hub/pkg/api/security.smh.solo.io/v1alpha1/types"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// if you need to update the golden files programmatically, change this to `true` to write the
// files instead of checking against them
var UPDATE_VMCSR_GOLDENS = false

var _ = Describe("VMCSR Table Printer", func() {
	const tpGoldenDirectory = "vmcsr"

	var runTest = func(fileName string, vmcsrs []*smh_security.VirtualMeshCertificateSigningRequest) {
		goldenFilename := test_goldens.GoldenFilePath(tpGoldenDirectory, fileName)
		goldenContents, err := ioutil.ReadFile(goldenFilename)
		Expect(err).NotTo(HaveOccurred())

		output := &bytes.Buffer{}
		err = table_printing.NewVirtualMeshMCSRPrinter(table_printing.DefaultTableBuilder()).Print(output, vmcsrs)

		if UPDATE_VMCSR_GOLDENS || test_goldens.UpdateGoldens() {
			err = ioutil.WriteFile(goldenFilename, []byte(output.String()), os.ModeAppend)
			Expect(err).NotTo(HaveOccurred(), "Failed to update the golden file")
			Fail("Need to change UPDATE_GOLDENS back to false before committing")
		} else {
			Expect(err).NotTo(HaveOccurred())
			Expect(output.String()).To(Equal(string(goldenContents)))
		}
	}

	DescribeTable("VMCSR printer", runTest,
		Entry(
			"can print different kinds of virtual mesh certificate signing requestss",
			"vmcsr",
			[]*smh_security.VirtualMeshCertificateSigningRequest{
				{
					ObjectMeta: k8s_meta_types.ObjectMeta{
						Name:      "vm-1",
						Namespace: "service-mesh-hub",
					},
					Spec: smh_security_types.VirtualMeshCertificateSigningRequestSpec{
						CsrData: []byte("test-csr"),
						CertConfig: &smh_security_types.VirtualMeshCertificateSigningRequestSpec_CertConfig{
							Hosts:    []string{"host1", "host2"},
							Org:      "my-org",
							MeshType: smh_core_types.MeshType_ISTIO1_5,
						},
						VirtualMeshRef: &smh_core_types.ResourceRef{
							Name:      "name-1",
							Namespace: "namespace-1",
						},
					},
					Status: smh_security_types.VirtualMeshCertificateSigningRequestStatus{
						ThirdPartyApproval: &smh_security_types.VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow{
							ApprovalStatus: smh_security_types.VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_APPROVED,
						},
						ComputedStatus: &smh_core_types.Status{
							State: smh_core_types.Status_ACCEPTED,
						},
					},
				},
				{
					ObjectMeta: k8s_meta_types.ObjectMeta{
						Name:      "vm-2",
						Namespace: "service-mesh-hub",
					},
					Spec: smh_security_types.VirtualMeshCertificateSigningRequestSpec{
						CsrData: nil,
						CertConfig: &smh_security_types.VirtualMeshCertificateSigningRequestSpec_CertConfig{
							Hosts:    []string{"host2", "host4"},
							Org:      "linkerd-org",
							MeshType: smh_core_types.MeshType_LINKERD,
						},
						VirtualMeshRef: &smh_core_types.ResourceRef{
							Name:      "name-2",
							Namespace: "namespace-2",
						},
					},
					Status: smh_security_types.VirtualMeshCertificateSigningRequestStatus{
						ThirdPartyApproval: &smh_security_types.VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow{
							ApprovalStatus: smh_security_types.VirtualMeshCertificateSigningRequestStatus_ThirdPartyApprovalWorkflow_DENIED,
						},
						ComputedStatus: &smh_core_types.Status{
							State:   smh_core_types.Status_CONFLICT,
							Message: "there was a conflict",
						},
						Response: &smh_security_types.VirtualMeshCertificateSigningRequestStatus_Response{},
					},
				},
			},
		),
	)
})
