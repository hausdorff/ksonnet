package kubecfg

import "github.com/ksonnet/ksonnet/metadata"

type InitCmd struct {
	name      string
	rootPath  metadata.AbsPath
	spec      metadata.ClusterSpec
	context   *string
	serverURI *string
	namespace *string
}

func NewInitCmd(name string, rootPath metadata.AbsPath, specFlag string, context, serverURI, namespace *string) (*InitCmd, error) {
	// NOTE: We're taking `rootPath` here as an absolute path (rather than a partial path we expand to an absolute path)
	// to make it more testable.

	spec, err := metadata.ParseClusterSpec(specFlag)
	if err != nil {
		return nil, err
	}

	return &InitCmd{name: name, rootPath: rootPath, spec: spec, context: context, serverURI: serverURI, namespace: namespace}, nil
}

func (c *InitCmd) Run() error {
	_, err := metadata.Init(c.name, c.rootPath, c.spec, c.context, c.serverURI, c.namespace)
	return err
}
