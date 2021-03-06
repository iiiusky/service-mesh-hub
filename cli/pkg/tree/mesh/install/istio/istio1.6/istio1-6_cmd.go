package istio1_6

import (
	"context"
	"io"

	"github.com/solo-io/service-mesh-hub/cli/pkg/cliconstants"
	"github.com/solo-io/service-mesh-hub/cli/pkg/common"
	"github.com/solo-io/service-mesh-hub/cli/pkg/options"
	"github.com/solo-io/service-mesh-hub/cli/pkg/tree/mesh/install/istio"
	"github.com/solo-io/service-mesh-hub/pkg/common/container-runtime/docker"
	"github.com/solo-io/service-mesh-hub/pkg/common/filesystem/files"
	"github.com/solo-io/service-mesh-hub/pkg/common/kube/kubeconfig"
	"github.com/solo-io/service-mesh-hub/pkg/common/mesh-installation/istio/operator"
	"github.com/spf13/cobra"
)

type Istio1_6Cmd *cobra.Command

func NewIstio1_6InstallCmd(
	ctx context.Context,
	out io.Writer,
	kubeLoader kubeconfig.KubeLoader,
	opts *options.Options,
	kubeClientsFactory common.KubeClientsFactory,
	clientsFactory common.ClientsFactory,
	imageNameParser docker.ImageNameParser,
	fileReader files.FileReader,
) Istio1_6Cmd {
	cmd := &cobra.Command{
		Use:   cliconstants.Istio1_6Command.Use,
		Short: cliconstants.Istio1_6Command.Short,
		RunE: istio.BuildIstioInstallationRunFunc(
			ctx,
			out,
			operator.IstioVersion1_6,
			kubeLoader,
			opts,
			kubeClientsFactory,
			clientsFactory,
			imageNameParser,
			fileReader,
		),
	}

	options.AddIstioInstallFlags(cmd, opts)

	return cmd
}
