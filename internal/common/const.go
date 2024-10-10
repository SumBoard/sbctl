package common

const (
	SbBootloaderPodName = "sumboard-sbctl-sumboard-bootloader"
	SbChartName         = "sumboard/sumboard"
	SbChartRelease      = "sumboard-sbctl"
	SbIngress           = "ingress-sbctl"
	SbNamespace         = "sumboard-sbctl"
	SbRepoName          = "sumboard"
	SbRepoURL           = "https://sumboard.github.io/helm-charts"
	NginxChartName      = "nginx/ingress-nginx"
	NginxChartRelease   = "ingress-nginx"
	NginxNamespace      = "ingress-nginx"
	NginxRepoName       = "nginx"
	NginxRepoURL        = "https://kubernetes.github.io/ingress-nginx"

	// DockerAuthSecretName is the name of the secret which holds the docker authentication information.
	DockerAuthSecretName = "docker-auth"
)
